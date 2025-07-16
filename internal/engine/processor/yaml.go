package processor

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"gopkg.in/yaml.v3"
)

// AgentIdentity represents the identity section of an agent
type AgentIdentity struct {
	Name        string `yaml:"name" json:"name"`
	ID          string `yaml:"id" json:"id"`
	Version     string `yaml:"version" json:"version"`
	Description string `yaml:"description" json:"description"`
	Role        string `yaml:"role" json:"role"`
	Goal        string `yaml:"goal" json:"goal"`
	Icon        string `yaml:"icon,omitempty" json:"icon,omitempty"`
}

// AgentSpec represents the complete agent specification
type AgentSpec struct {
	Identity         AgentIdentity     `yaml:"identity" json:"identity"`
	ActivationPrompt []string          `yaml:"activation_prompt" json:"activation_prompt"`
	Principles       []string          `yaml:"principles" json:"principles"`
	Customization    *string           `yaml:"customization" json:"customization"`
	Commands         map[string]string `yaml:"commands" json:"commands"`
	Tasks            []string          `yaml:"tasks,omitempty" json:"tasks,omitempty"`
}

// GetCustomization returns the customization value or empty string if nil
func (a *AgentSpec) GetCustomization() string {
	if a.Customization == nil {
		return ""
	}
	return *a.Customization
}

// HasCustomization returns true if customization field is present and non-empty
func (a *AgentSpec) HasCustomization() bool {
	return a.Customization != nil && *a.Customization != ""
}

// Agent represents the top-level agent structure
type Agent struct {
	Agent AgentSpec `yaml:"agent" json:"agent"`
}

// ValidationError represents a validation error with details
type ValidationError struct {
	Field   string
	Value   any
	Message string
	File    string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s' in file '%s': %s", e.Field, e.File, e.Message)
}

// YAMLProcessor handles YAML parsing and validation for agents
type YAMLProcessor struct {
	schema *jsonschema.Schema
}

// NewYAMLProcessor creates a new YAML processor with embedded schema
func NewYAMLProcessor(embeddedAssets embed.FS) (*YAMLProcessor, error) {
	processor := &YAMLProcessor{}

	// Load embedded schema - this MUST succeed
	if err := processor.loadEmbeddedSchema(embeddedAssets); err != nil {
		return nil, fmt.Errorf("failed to load embedded schema: %w", err)
	}

	return processor, nil
}

// NewYAMLProcessorFromFile creates a new YAML processor with schema from file (for testing)
func NewYAMLProcessorFromFile(schemaPath string) (*YAMLProcessor, error) {
	if schemaPath == "" {
		return nil, fmt.Errorf("schema path is required")
	}

	processor := &YAMLProcessor{}

	// Load schema from file
	if err := processor.loadSchemaFromFile(schemaPath); err != nil {
		return nil, fmt.Errorf("failed to load schema from file: %w", err)
	}

	return processor, nil
}

// loadEmbeddedSchema loads the JSON schema from embedded assets
func (p *YAMLProcessor) loadEmbeddedSchema(embeddedAssets embed.FS) error {
	schemaData, err := embeddedAssets.ReadFile("assets/schemas/agent-schema.json")
	if err != nil {
		return fmt.Errorf("failed to read embedded schema file: %w", err)
	}

	compiler := jsonschema.NewCompiler()
	compiler.DefaultDraft(jsonschema.Draft7)

	// Parse JSON schema data
	var schemaDoc any
	if parseErr := json.Unmarshal(schemaData, &schemaDoc); parseErr != nil {
		return fmt.Errorf("failed to parse schema JSON: %w", parseErr)
	}

	schemaURL := "agent-schema.json"
	if addErr := compiler.AddResource(schemaURL, schemaDoc); addErr != nil {
		return fmt.Errorf("failed to add schema resource: %w", addErr)
	}

	schema, err := compiler.Compile(schemaURL)
	if err != nil {
		return fmt.Errorf("failed to compile schema: %w", err)
	}

	p.schema = schema
	return nil
}

// loadSchemaFromFile loads the JSON schema from a file (for testing)
func (p *YAMLProcessor) loadSchemaFromFile(schemaPath string) error {
	schemaData, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	compiler := jsonschema.NewCompiler()
	compiler.DefaultDraft(jsonschema.Draft7)

	// Parse JSON schema data
	var schemaDoc any
	if parseErr := json.Unmarshal(schemaData, &schemaDoc); parseErr != nil {
		return fmt.Errorf("failed to parse schema JSON: %w", parseErr)
	}

	if addErr := compiler.AddResource(schemaPath, schemaDoc); addErr != nil {
		return fmt.Errorf("failed to add schema resource: %w", addErr)
	}

	schema, err := compiler.Compile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to compile schema: %w", err)
	}

	p.schema = schema
	return nil
}

// ParseAgentFile parses a YAML agent file and returns the agent structure
func (p *YAMLProcessor) ParseAgentFile(filePath string) (*Agent, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open agent file '%s': %w", filePath, err)
	}
	defer func() {
		_ = file.Close()
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read agent file '%s': %w", filePath, err)
	}

	var agent Agent
	if err := yaml.Unmarshal(content, &agent); err != nil {
		return nil, fmt.Errorf("failed to parse YAML in file '%s': %w", filePath, err)
	}

	return &agent, nil
}

// ParseAgentFileRaw parses a YAML agent file into a raw map to capture all fields
func (p *YAMLProcessor) ParseAgentFileRaw(filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open agent file '%s': %w", filePath, err)
	}
	defer func() {
		_ = file.Close()
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read agent file '%s': %w", filePath, err)
	}

	var rawData map[string]interface{}
	if err := yaml.Unmarshal(content, &rawData); err != nil {
		return nil, fmt.Errorf("failed to parse YAML in file '%s': %w", filePath, err)
	}

	return rawData, nil
}

// ValidateAgent validates an agent against the JSON schema
func (p *YAMLProcessor) ValidateAgent(agent *Agent, filePath string) []ValidationError {
	var errors []ValidationError

	// Schema is guaranteed to be loaded (constructor fails if not)

	// Convert agent to any for schema validation
	var agentData any
	jsonData, err := json.Marshal(agent)
	if err != nil {
		errors = append(errors, ValidationError{
			Field:   "agent",
			Value:   agent,
			Message: fmt.Sprintf("failed to convert agent to JSON: %v", err),
			File:    filePath,
		})
		return errors
	}

	if err := json.Unmarshal(jsonData, &agentData); err != nil {
		errors = append(errors, ValidationError{
			Field:   "agent",
			Value:   agent,
			Message: fmt.Sprintf("failed to unmarshal agent data: %v", err),
			File:    filePath,
		})
		return errors
	}

	// Validate against schema using shared validation logic
	return p.validateAgainstSchema(agentData, agent, filePath)
}

// ValidateAgentRaw validates raw agent data against the JSON schema
func (p *YAMLProcessor) ValidateAgentRaw(rawData map[string]interface{}, filePath string) []ValidationError {
	// Validate against schema using shared validation logic
	return p.validateAgainstSchema(rawData, rawData, filePath)
}

// validateAgainstSchema performs schema validation with shared error handling
func (p *YAMLProcessor) validateAgainstSchema(data any, originalValue any, filePath string) []ValidationError {
	var errors []ValidationError

	// Validate against schema
	if err := p.schema.Validate(data); err != nil {
		// Handle validation error - new library returns structured error
		if validationErr, ok := err.(*jsonschema.ValidationError); ok {
			errors = append(errors, ValidationError{
				Field:   strings.Join(validationErr.InstanceLocation, "."),
				Value:   nil, // ValidationError doesn't expose instance value
				Message: validationErr.Error(),
				File:    filePath,
			})
			// Handle nested errors
			for _, subErr := range validationErr.Causes {
				errors = append(errors, ValidationError{
					Field:   strings.Join(subErr.InstanceLocation, "."),
					Value:   nil, // ValidationError doesn't expose instance value
					Message: subErr.Error(),
					File:    filePath,
				})
			}
		} else {
			errors = append(errors, ValidationError{
				Field:   "agent",
				Value:   originalValue,
				Message: fmt.Sprintf("schema validation failed: %v", err),
				File:    filePath,
			})
		}
	}

	return errors
}

// ProcessAndValidateAgent processes and validates an agent file
func (p *YAMLProcessor) ProcessAndValidateAgent(filePath string) (*Agent, []ValidationError, error) {
	// Parse the file into a structured agent for return
	agent, err := p.ParseAgentFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	// Parse the file as raw data for validation (to capture all fields)
	rawData, err := p.ParseAgentFileRaw(filePath)
	if err != nil {
		return nil, nil, err
	}

	validationErrors := p.ValidateAgentRaw(rawData, filePath)
	return agent, validationErrors, nil
}
