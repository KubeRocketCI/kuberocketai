package processor

import (
	"os"
	"testing"
)

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
