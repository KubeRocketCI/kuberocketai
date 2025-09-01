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
	"fmt"
)

// ValidationSystemBuilder provides fluent API for building validation system
type ValidationSystemBuilder struct {
	baseDir        string
	embeddedAssets embed.FS
	config         ValidationConfig
	err            error
}

// NewValidationSystemBuilder creates a new builder
func NewValidationSystemBuilder(baseDir string) *ValidationSystemBuilder {
	return &ValidationSystemBuilder{
		baseDir: baseDir,
		config:  ValidationConfig{},
	}
}

// WithEmbeddedAssets sets the embedded assets for schema loading
func (b *ValidationSystemBuilder) WithEmbeddedAssets(assets embed.FS) *ValidationSystemBuilder {
	if b.err != nil {
		return b
	}

	b.embeddedAssets = assets
	return b
}

// Build creates the complete validation system with all dependencies
func (b *ValidationSystemBuilder) Build() (UniversalValidator, error) {
	if b.err != nil {
		return nil, b.err
	}

	// Check if embedded assets were set by attempting to read from them
	if _, err := b.embeddedAssets.Open("."); err != nil {
		return nil, fmt.Errorf("embedded assets are required and must be valid")
	}

	// Build schema registry
	schemaRegistry := NewEmbeddedSchemaRegistry()
	if err := schemaRegistry.LoadSchemas(b.embeddedAssets); err != nil {
		return nil, fmt.Errorf("failed to load schemas: %w", err)
	}
	b.config.SchemaRegistry = schemaRegistry

	// Build simple dependency validator
	dependencyValidator := NewSimpleDependencyValidator(b.baseDir)
	b.config.AssetResolver = dependencyValidator

	// Build parser registry with integrated parsers
	parserRegistry, err := NewIntegratedParserRegistry(b.embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create parser registry: %w", err)
	}
	b.config.ParserRegistry = parserRegistry

	// Build validator registry with integrated validators
	validatorRegistry, err := NewIntegratedValidatorRegistry(b.embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create validator registry: %w", err)
	}
	b.config.ValidatorRegistry = validatorRegistry

	// Set default validation rules (empty for now)
	b.config.ValidationRules = []ValidationRule{}

	// Create the validator
	return NewFrameworkValidator(b.config)
}
