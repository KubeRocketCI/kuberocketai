/*
Copyright © 2025 KubeRocketAI Team

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package validation

import (
	"embed"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// EmbeddedSchemaRegistry implements SchemaRegistry with embedded assets
type EmbeddedSchemaRegistry struct {
	schemas map[SchemaType]*jsonschema.Schema
	mutex   sync.RWMutex
}

// NewEmbeddedSchemaRegistry creates a new schema registry
func NewEmbeddedSchemaRegistry() *EmbeddedSchemaRegistry {
	return &EmbeddedSchemaRegistry{
		schemas: make(map[SchemaType]*jsonschema.Schema),
	}
}

// LoadSchemas loads all schemas from embedded assets
func (r *EmbeddedSchemaRegistry) LoadSchemas(embeddedAssets embed.FS) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Load agent schema
	if err := r.loadSchema(embeddedAssets, "assets/schemas/agent-schema.json", SchemaTypeAgent); err != nil {
		return fmt.Errorf("failed to load agent schema: %w", err)
	}

	// Load frontmatter schema
	if err := r.loadSchema(embeddedAssets, "assets/schemas/task-metadata.json", SchemaTypeFrontmatter); err != nil {
		return fmt.Errorf("failed to load frontmatter schema: %w", err)
	}

	return nil
}

// loadSchema loads a single schema from embedded assets
func (r *EmbeddedSchemaRegistry) loadSchema(embeddedAssets embed.FS, path string, schemaType SchemaType) error {
	schemaData, err := embeddedAssets.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read schema file %s: %w", path, err)
	}

	compiler := jsonschema.NewCompiler()
	compiler.DefaultDraft(jsonschema.Draft7)

	var schemaDoc interface{}
	if parseErr := json.Unmarshal(schemaData, &schemaDoc); parseErr != nil {
		return fmt.Errorf("failed to parse schema JSON: %w", parseErr)
	}

	schemaURL := string(schemaType) + ".json"
	if addErr := compiler.AddResource(schemaURL, schemaDoc); addErr != nil {
		return fmt.Errorf("failed to add schema resource: %w", addErr)
	}

	schema, err := compiler.Compile(schemaURL)
	if err != nil {
		return fmt.Errorf("failed to compile schema: %w", err)
	}

	r.schemas[schemaType] = schema
	return nil
}

// GetSchema returns a compiled schema by type
func (r *EmbeddedSchemaRegistry) GetSchema(schemaType SchemaType) (JSONSchema, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	schema, exists := r.schemas[schemaType]
	if !exists {
		return nil, fmt.Errorf("schema not found: %s", schemaType)
	}

	return &JSONSchemaWrapper{schema: schema}, nil
}

// ValidateAgainstSchema validates data against a specific schema type
func (r *EmbeddedSchemaRegistry) ValidateAgainstSchema(data interface{}, schemaType SchemaType) error {
	schema, err := r.GetSchema(schemaType)
	if err != nil {
		return err
	}

	return schema.Validate(data)
}

// JSONSchemaWrapper wraps jsonschema.Schema to implement JSONSchema interface
type JSONSchemaWrapper struct {
	schema *jsonschema.Schema
}

// Validate validates data against the wrapped schema
func (w *JSONSchemaWrapper) Validate(data interface{}) error {
	return w.schema.Validate(data)
}
