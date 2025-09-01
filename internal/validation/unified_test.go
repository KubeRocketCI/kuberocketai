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
	"os"
	"path/filepath"
	"testing"
)

// Test implementations for null registries
type NullSchemaRegistry struct{}

func (r *NullSchemaRegistry) LoadSchemas(embeddedAssets embed.FS) error {
	return nil
}

func (r *NullSchemaRegistry) GetSchema(schemaType SchemaType) (JSONSchema, error) {
	return &NullJSONSchema{}, nil
}

func (r *NullSchemaRegistry) ValidateAgainstSchema(data interface{}, schemaType SchemaType) error {
	return nil
}

type NullJSONSchema struct{}

func (s *NullJSONSchema) Validate(data interface{}) error {
	return nil
}

type NullParserRegistry struct{}

func (r *NullParserRegistry) RegisterParser(extension string, parser FileParser) {}

func (r *NullParserRegistry) GetParser(filePath string) (FileParser, error) {
	return &NullFileParser{}, nil
}

func (r *NullParserRegistry) CanParse(filePath string) bool {
	return true
}

type NullValidatorRegistry struct{}

func (r *NullValidatorRegistry) RegisterValidator(fileType string, validator ValidatorStrategy) {}

func (r *NullValidatorRegistry) GetValidators(fileType string) []ValidatorStrategy {
	return []ValidatorStrategy{}
}

type NullFileParser struct{}

func (p *NullFileParser) ParseFile(filePath string) (ParsedFile, error) {
	return ParsedFile{
		Type:     "test",
		Path:     filePath,
		Metadata: map[string]interface{}{},
		Content:  "",
	}, nil
}

func (p *NullFileParser) ParseContent(content string, filePath string) (ParsedFile, error) {
	return ParsedFile{
		Type:     "test",
		Path:     filePath,
		Metadata: map[string]interface{}{},
		Content:  content,
	}, nil
}

func (p *NullFileParser) GetFileType() string {
	return "test"
}

func TestValidationConfig_Validate(t *testing.T) {
	// Create temp dir for testing
	tmpDir, err := os.MkdirTemp("", "validation_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	tests := []struct {
		name    string
		config  ValidationConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: ValidationConfig{
				SchemaRegistry:    &NullSchemaRegistry{},
				ParserRegistry:    &NullParserRegistry{},
				ValidatorRegistry: &NullValidatorRegistry{},
				AssetResolver:     NewSimpleDependencyValidator(tmpDir),
				ValidationRules:   []ValidationRule{},
			},
			wantErr: false,
		},
		{
			name: "missing schema registry",
			config: ValidationConfig{
				ParserRegistry:    &NullParserRegistry{},
				ValidatorRegistry: &NullValidatorRegistry{},
				AssetResolver:     NewSimpleDependencyValidator(tmpDir),
			},
			wantErr: true,
		},
		{
			name: "missing parser registry",
			config: ValidationConfig{
				SchemaRegistry:    &NullSchemaRegistry{},
				ValidatorRegistry: &NullValidatorRegistry{},
				AssetResolver:     NewSimpleDependencyValidator(tmpDir),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidationConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewFrameworkValidator(t *testing.T) {
	// Create temp dir for testing
	tmpDir, err := os.MkdirTemp("", "validator_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	validConfig := ValidationConfig{
		SchemaRegistry:    &NullSchemaRegistry{},
		ParserRegistry:    &NullParserRegistry{},
		ValidatorRegistry: &NullValidatorRegistry{},
		AssetResolver:     NewSimpleDependencyValidator(tmpDir),
		ValidationRules:   []ValidationRule{},
	}

	validator, err := NewFrameworkValidator(validConfig)
	if err != nil {
		t.Fatalf("NewFrameworkValidator() error = %v, want nil", err)
	}

	if validator == nil {
		t.Error("NewFrameworkValidator() returned nil validator")
	}
}

func TestFrameworkValidator_ValidateAgent(t *testing.T) {
	// Create temp dir for testing
	tmpDir, err := os.MkdirTemp("", "agent_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	config := ValidationConfig{
		SchemaRegistry:    &NullSchemaRegistry{},
		ParserRegistry:    &NullParserRegistry{},
		ValidatorRegistry: &NullValidatorRegistry{},
		AssetResolver:     NewSimpleDependencyValidator(tmpDir),
		ValidationRules:   []ValidationRule{},
	}

	validator, err := NewFrameworkValidator(config)
	if err != nil {
		t.Fatalf("NewFrameworkValidator() error = %v", err)
	}

	// Create a temporary test file

	testFile := filepath.Join(tmpDir, "test.yaml")
	content := `agent:
  identity:
    name: "Test Agent"
    id: test-v1
    version: "1.0.0"
    description: "Test agent"
    role: "Test Role"
    goal: "Test goal"
`
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	result := validator.ValidateAgent(testFile)

	// With null implementations, we expect validation to pass
	if result.Severity != SeverityInfo {
		t.Errorf("ValidateAgent() severity = %v, want %v", result.Severity, SeverityInfo)
	}
}

func TestValidationSystemBuilder(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "validation_builder_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	builder := NewValidationSystemBuilder(tmpDir)
	if builder == nil {
		t.Fatal("NewValidationSystemBuilder() returned nil")
	}

	// Test builder without embedded assets (should fail)
	_, err = builder.Build()
	if err == nil {
		t.Error("Build() without embedded assets should fail")
	}
}

func TestValidationSystemBuilderWithoutAssets(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "analyzer_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create a minimal embedded FS for testing
	validator, err := NewValidationSystemBuilder(tmpDir).Build()
	if err == nil {
		t.Error("Expected error when building without embedded assets")
	}

	// Test that we can create a validator (will fail without real embedded assets, but that's OK)
	if validator != nil {
		t.Error("Expected nil validator when build fails")
	}
}

func TestSeverity_String(t *testing.T) {
	tests := []struct {
		severity Severity
		want     string
	}{
		{SeverityInfo, "INFO"},
		{SeverityWarning, "WARNING"},
		{SeverityError, "ERROR"},
		{SeverityCritical, "CRITICAL"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.severity.String(); got != tt.want {
				t.Errorf("Severity.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorCode_Constants(t *testing.T) {
	// Test that error codes are properly defined
	codes := []ErrorCode{
		ErrorCodeMissingTask,
		ErrorCodeBrokenLink,
		ErrorCodeInvalidFormat,
		ErrorCodeArchViolation,
		ErrorCodeFrontmatterParse,
		ErrorCodeFrontmatterSchema,
		ErrorCodeMissingDependency,
		ErrorCodeOrphanedFile,
		ErrorCodeCircularDependency,
		ErrorCodeInvalidExtension,
	}

	for i, code := range codes {
		if string(code) == "" {
			t.Errorf("Error code %d is empty", i)
		}
	}
}

func TestValidationResult_Structure(t *testing.T) {
	result := ValidationResult{
		Severity:    SeverityCritical,
		Type:        "test_error",
		Code:        ErrorCodeMissingTask,
		File:        "test.yaml",
		Line:        42,
		Message:     "Test validation message",
		Context:     map[string]string{"test": "context"},
		FixGuidance: "Fix the test issue",
	}

	if result.Severity != SeverityCritical {
		t.Error("ValidationResult severity not set correctly")
	}
	if result.Type != "test_error" {
		t.Error("ValidationResult type not set correctly")
	}
	if result.File != "test.yaml" {
		t.Error("ValidationResult file not set correctly")
	}
	if result.Message != "Test validation message" {
		t.Error("ValidationResult message not set correctly")
	}
	if result.FixGuidance != "Fix the test issue" {
		t.Error("ValidationResult fix guidance not set correctly")
	}
	if ctx, ok := result.Context.(map[string]string); !ok || ctx["test"] != "context" {
		t.Error("ValidationResult context not set correctly")
	}
	if result.Code != ErrorCodeMissingTask {
		t.Error("ValidationResult code not set correctly")
	}
	if result.Line != 42 {
		t.Error("ValidationResult line not set correctly")
	}
}
