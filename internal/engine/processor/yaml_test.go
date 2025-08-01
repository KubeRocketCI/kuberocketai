package processor

import (
	"encoding/json"
	"os"
	"testing"
)

// stringPtr returns a pointer to the given string (helper for tests)
func stringPtr(s string) *string {
	return &s
}

func TestYAMLProcessor_ParseAgentFile(t *testing.T) {
	// For testing, use the file-based constructor
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	if err != nil {
		t.Fatalf("Failed to create processor: %v", err)
	}

	// Test valid agent file
	t.Run("valid agent file", func(t *testing.T) {
		validAgentContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
    description: "A test agent for validation"
    role: "Test Engineer"
    goal: "Test the validation system"
    icon: "🧪"
  activation_prompt:
    - "You are a test agent"
    - "Follow test protocols"
  principles:
    - "Always test thoroughly"
    - "Provide clear feedback"
    - "Document test results"
  customization: ""
  commands:
    help: "Show available commands"
    chat: "Default chat mode"
    exit: "Exit test mode"
  tasks:
    - "./.krci-ai/tasks/test-task.md"
`

		tmpFile := createTempFile(t, validAgentContent)
		defer os.Remove(tmpFile)

		agent, err := processor.ParseAgentFile(tmpFile)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		// Validate parsed content
		if agent.Agent.Identity.Name != "Test Agent" {
			t.Errorf("Expected name 'Test Agent', got: %s", agent.Agent.Identity.Name)
		}

		if agent.Agent.Identity.ID != "test-agent-v1" {
			t.Errorf("Expected ID 'test-agent-v1', got: %s", agent.Agent.Identity.ID)
		}
	})

	// Test malformed YAML
	t.Run("malformed YAML", func(t *testing.T) {
		malformedContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
    description: "A test agent for validation"
    role: "Test Engineer"
    goal: "Test the validation system"
  activation_prompt:
    - "You are a test agent"
    - "Follow test protocols"
  principles:
    - "Always test thoroughly"
    - "Provide clear feedback"
    - "Document test results"
  commands:
    help: "Show available commands"
    chat: "Default chat mode"
    exit: "Exit test mode"
  invalid_yaml: [unclosed
`

		tmpFile := createTempFile(t, malformedContent)
		defer os.Remove(tmpFile)

		_, err := processor.ParseAgentFile(tmpFile)
		if err == nil {
			t.Error("Expected error for malformed YAML, got nil")
		}
	})

	// Test non-existent file
	t.Run("non-existent file", func(t *testing.T) {
		_, err := processor.ParseAgentFile("/path/to/nonexistent/file.yaml")
		if err == nil {
			t.Error("Expected error for non-existent file, got nil")
		}
	})
}

func TestYAMLProcessor_ValidateAgent_WithSchema(t *testing.T) {
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	if err != nil {
		t.Fatalf("Failed to create processor: %v", err)
	}

	// Test valid agent with schema validation
	t.Run("valid agent - schema validation", func(t *testing.T) {
		agent := &Agent{
			Agent: AgentSpec{
				Identity: AgentIdentity{
					Name:        "Test Agent",
					ID:          "test-agent-v1",
					Version:     "1.0.0",
					Description: "A test agent for validation testing with sufficient length",
					Role:        "Test Engineer",
					Goal:        "Test the validation system thoroughly with proper length",
					Icon:        "🧪",
				},
				ActivationPrompt: []string{
					"You are a test agent designed for validation testing",
					"Follow test protocols carefully and thoroughly",
				},
				Principles: []string{
					"Always test thoroughly and document results with sufficient detail",
					"Provide clear and actionable feedback to users",
					"Maintain high quality standards throughout the testing process",
				},
				Customization: stringPtr(""),
				Commands: map[string]string{
					"help": "Show available commands",
					"chat": "Default chat mode",
					"exit": "Exit test mode",
				},
				Tasks: []string{
					"./.krci-ai/tasks/test-task.md",
				},
			},
		}

		errors := processor.ValidateAgent(agent, "test.yaml")
		if len(errors) != 0 {
			t.Errorf("Expected no validation errors, got: %d errors", len(errors))
			for _, err := range errors {
				t.Logf("Validation error: %s", err.Error())
			}
		}
	})

	// Test customization field validation with nil pointer (missing field)
	t.Run("missing customization field with pointer - schema validation", func(t *testing.T) {
		agent := &Agent{
			Agent: AgentSpec{
				Identity: AgentIdentity{
					Name:        "Test Agent",
					ID:          "test-agent-v1",
					Version:     "1.0.0",
					Description: "A test agent for validation testing with sufficient length",
					Role:        "Test Engineer",
					Goal:        "Test the validation system thoroughly with proper length",
					Icon:        "🧪",
				},
				ActivationPrompt: []string{
					"You are a test agent designed for validation testing",
					"Follow test protocols carefully and thoroughly",
				},
				Principles: []string{
					"Always test thoroughly and document results with sufficient detail",
					"Provide clear and actionable feedback to users",
					"Maintain high quality standards throughout the testing process",
				},
				Customization: nil, // Missing field represented as nil
				Commands: map[string]string{
					"help": "Show available commands",
					"chat": "Default chat mode",
					"exit": "Exit test mode",
				},
			},
		}

		errors := processor.ValidateAgent(agent, "test.yaml")
		if len(errors) == 0 {
			t.Error("Expected validation errors for nil customization field, but got none")
		} else {
			t.Logf("Validation correctly failed with %d errors", len(errors))
			for _, err := range errors {
				t.Logf("Error: %s", err.Error())
			}
		}
	})

	// Test customization field validation with raw JSON
	t.Run("missing customization field - schema validation", func(t *testing.T) {
		// Test with raw JSON data that truly omits the customization field
		rawJSON := `{
			"agent": {
				"identity": {
					"name": "Test Agent",
					"id": "test-agent-v1",
					"version": "1.0.0",
					"description": "A test agent for validation testing with sufficient length",
					"role": "Test Engineer",
					"goal": "Test the validation system thoroughly with proper length",
					"icon": "🧪"
				},
				"activation_prompt": [
					"You are a test agent designed for validation testing",
					"Follow test protocols carefully and thoroughly"
				],
				"principles": [
					"Always test thoroughly and document results with sufficient detail",
					"Provide clear and actionable feedback to users",
					"Maintain high quality standards throughout the testing process"
				],
				"commands": {
					"help": "Show available commands",
					"chat": "Default chat mode",
					"exit": "Exit test mode"
				}
			}
		}`

		var agentData any
		if err := json.Unmarshal([]byte(rawJSON), &agentData); err != nil {
			t.Fatalf("Failed to unmarshal test JSON: %v", err)
		}

		// Validate directly against schema
		if err := processor.schema.Validate(agentData); err != nil {
			t.Logf("Schema validation error (expected): %v", err)
		} else {
			t.Error("Expected schema validation to fail for missing customization field, but it passed")
		}
	})

	// Test invalid agent with schema validation
	t.Run("invalid agent - schema validation", func(t *testing.T) {
		agent := &Agent{
			Agent: AgentSpec{
				Identity: AgentIdentity{
					Name:        "Test Agent",
					ID:          "invalid-id", // Invalid ID format
					Version:     "invalid",    // Invalid version format
					Description: "Short",      // Too short
					Role:        "Test",       // Too short
					Goal:        "Test",       // Too short
				},
				ActivationPrompt: []string{
					"Short", // Too short
				},
				Principles: []string{
					"Short", // Too short
				},
				Customization: stringPtr(""),
				Commands: map[string]string{
					"help": "Show available commands",
					"chat": "Default chat mode",
					"exit": "Exit test mode",
				},
			},
		}

		errors := processor.ValidateAgent(agent, "test.yaml")
		if len(errors) == 0 {
			t.Error("Expected validation errors for invalid agent, but got none")
		}
	})
}

func TestYAMLProcessor_WithRealAgentFiles(t *testing.T) {
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	if err != nil {
		t.Fatalf("Failed to create processor: %v", err)
	}

	// Test with actual agent files from the project
	agentFiles := []string{
		"../../../cmd/krci-ai/assets/framework/core/agents/architect.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/dev.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/pm.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/qa.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/ba.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/po.yaml",
	}

	for _, agentFile := range agentFiles {
		t.Run(agentFile, func(t *testing.T) {
			agent, validationErrors, err := processor.ProcessAndValidateAgent(agentFile)
			if err != nil {
				t.Fatalf("Failed to process agent file %s: %v", agentFile, err)
			}

			if len(validationErrors) > 0 {
				t.Errorf("Agent file %s has validation errors:", agentFile)
				for _, validationError := range validationErrors {
					t.Errorf("  - %s", validationError.Error())
				}
			}

			// Basic sanity check
			if agent.Agent.Identity.Name == "" {
				t.Errorf("Agent file %s has empty name", agentFile)
			}
		})
	}
}

func TestAgentSpec_GetCustomization(t *testing.T) {
	tests := []struct {
		name          string
		customization *string
		want          string
	}{
		{
			name:          "nil customization",
			customization: nil,
			want:          "",
		},
		{
			name:          "empty customization",
			customization: stringPtr(""),
			want:          "",
		},
		{
			name:          "non-empty customization",
			customization: stringPtr("custom behavior"),
			want:          "custom behavior",
		},
		{
			name:          "whitespace customization",
			customization: stringPtr("  custom with spaces  "),
			want:          "  custom with spaces  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := &AgentSpec{
				Customization: tt.customization,
			}
			if got := agent.GetCustomization(); got != tt.want {
				t.Errorf("GetCustomization() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAgentSpec_HasCustomization(t *testing.T) {
	tests := []struct {
		name          string
		customization *string
		want          bool
	}{
		{
			name:          "nil customization",
			customization: nil,
			want:          false,
		},
		{
			name:          "empty customization",
			customization: stringPtr(""),
			want:          false,
		},
		{
			name:          "non-empty customization",
			customization: stringPtr("custom behavior"),
			want:          true,
		},
		{
			name:          "whitespace customization",
			customization: stringPtr("  "),
			want:          true,
		},
		{
			name:          "single space customization",
			customization: stringPtr(" "),
			want:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := &AgentSpec{
				Customization: tt.customization,
			}
			if got := agent.HasCustomization(); got != tt.want {
				t.Errorf("HasCustomization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidationError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  ValidationError
		want string
	}{
		{
			name: "basic validation error",
			err: ValidationError{
				Field:   "name",
				Value:   "test",
				Message: "field is required",
				File:    "test.yaml",
			},
			want: "validation failed for field 'name' in file 'test.yaml': field is required",
		},
		{
			name: "validation error with empty values",
			err: ValidationError{
				Field:   "",
				Value:   nil,
				Message: "",
				File:    "",
			},
			want: "validation failed for field '' in file '': ",
		},
		{
			name: "validation error with complex message",
			err: ValidationError{
				Field:   "identity.version",
				Value:   "invalid-version",
				Message: "must match pattern ^v?\\d+\\.\\d+\\.\\d+",
				File:    "agents/dev.yaml",
			},
			want: "validation failed for field 'identity.version' in file 'agents/dev.yaml': must match pattern ^v?\\d+\\.\\d+\\.\\d+",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("ValidationError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYAMLProcessor_ParseAgentFileRaw(t *testing.T) {
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	if err != nil {
		t.Fatalf("Failed to create processor: %v", err)
	}

	// Test with valid YAML file
	t.Run("valid YAML file", func(t *testing.T) {
		validContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
    description: "A test agent for validation"
    role: "Test Engineer"
    goal: "Test the validation system"
    icon: "🧪"
  activation_prompt:
    - "You are a test agent"
  principles:
    - "Always test thoroughly"
  customization: ""
  commands:
    help: "Show available commands"
  tasks:
    - "./.krci-ai/tasks/test-task.md"
`

		tmpFile := createTempFile(t, validContent)
		defer os.Remove(tmpFile)

		rawData, err := processor.ParseAgentFileRaw(tmpFile)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		// Check if parsed data contains expected structure
		if rawData["agent"] == nil {
			t.Error("Expected 'agent' key in parsed data")
		}

		agentData, ok := rawData["agent"].(map[string]interface{})
		if !ok {
			t.Error("Expected 'agent' to be a map")
		} else if agentData["identity"] == nil {
			t.Error("Expected 'identity' key in agent data")
		}
	})

	// Test with non-existent file
	t.Run("non-existent file", func(t *testing.T) {
		_, err := processor.ParseAgentFileRaw("/path/to/nonexistent/file.yaml")
		if err == nil {
			t.Error("Expected error for non-existent file, got nil")
		}
	})
}

func TestYAMLProcessor_ValidateAgentRaw(t *testing.T) {
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	if err != nil {
		t.Fatalf("Failed to create processor: %v", err)
	}

	// Test that ValidateAgentRaw can be called (basic functionality test)
	t.Run("validate agent raw basic functionality", func(t *testing.T) {
		// Test with minimal data structure to ensure function executes
		testData := map[string]interface{}{
			"agent": map[string]interface{}{
				"identity": map[string]interface{}{
					"name": "Test",
				},
			},
		}

		// We don't care about the validation results, just that the function executes
		_ = processor.ValidateAgentRaw(testData, "test.yaml")
	})

	// Test invalid agent data
	t.Run("invalid agent data", func(t *testing.T) {
		invalidAgentData := map[string]interface{}{
			"agent": map[string]interface{}{
				"identity": map[string]interface{}{
					"name": "Test Agent",
					// Missing required fields
				},
			},
		}

		errors := processor.ValidateAgentRaw(invalidAgentData, "test.yaml")
		if len(errors) == 0 {
			t.Error("Expected validation errors for invalid agent data, but got none")
		}
	})
}

// Helper function to create temporary files for testing
func createTempFile(t *testing.T, content string) string {
	tmpFile, err := os.CreateTemp("", "test_agent_*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpFile.Name()
}
