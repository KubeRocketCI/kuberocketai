package cmd

import (
	"github.com/epam/kuberocketai/internal/engine/processor"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// testFrameworkValidator creates a validator for testing using file-based schema
func testFrameworkValidator(baseDir string) (*FrameworkValidator, error) {
	// Use file-based schema for testing
	yamlProcessor, err := processor.NewYAMLProcessorFromFile("../assets/schemas/agent-schema.json")
	if err != nil {
		return nil, err
	}

	return &FrameworkValidator{
		baseDir:       baseDir,
		yamlProcessor: yamlProcessor,
	}, nil
}

func TestFrameworkValidator_ValidateFramework(t *testing.T) {
	// Test with a temporary directory structure
	tempDir := t.TempDir()

	// Create .krci-ai directory structure
	krciDir := filepath.Join(tempDir, ".krci-ai")
	agentsDir := filepath.Join(krciDir, "agents")
	tasksDir := filepath.Join(krciDir, "tasks")

	err := os.MkdirAll(agentsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create agents directory: %v", err)
	}

	err = os.MkdirAll(tasksDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create tasks directory: %v", err)
	}

	// Create a valid agent file
	validAgentContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
    description: "A test agent for validation testing"
    role: "Test Engineer"
    goal: "Test the validation system thoroughly"
    icon: "🧪"
  activation_prompt:
    - "You are a test agent designed for validation"
    - "Follow test protocols carefully"
  principles:
    - "Always test thoroughly and document results"
    - "Provide clear and actionable feedback"
    - "Maintain high quality standards"
  commands:
    help: "Show available commands"
    chat: "Default chat mode"
    exit: "Exit test mode"
  tasks:
    - "./.krci-ai/tasks/test-task.md"
`

	validAgentFile := filepath.Join(agentsDir, "test-agent.yaml")
	err = os.WriteFile(validAgentFile, []byte(validAgentContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write valid agent file: %v", err)
	}

	// Create a valid task file
	validTaskContent := `# Test Task

This is a test task for validation.

## Description

Test task description.
`

	validTaskFile := filepath.Join(tasksDir, "test-task.md")
	err = os.WriteFile(validTaskFile, []byte(validTaskContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write valid task file: %v", err)
	}

	// Test validation with valid framework
	t.Run("valid framework", func(t *testing.T) {
		validator, err := testFrameworkValidator(tempDir)
		if err != nil {
			t.Fatalf("Failed to create validator: %v", err)
		}
		results, err := validator.ValidateFramework()

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !results.IsValid() {
			t.Error("Expected valid results, but validation failed")
		}

		if results.TotalFiles != 2 {
			t.Errorf("Expected 2 files, got: %d", results.TotalFiles)
		}

		if results.ValidFiles != 2 {
			t.Errorf("Expected 2 valid files, got: %d", results.ValidFiles)
		}

		if results.InvalidFiles != 0 {
			t.Errorf("Expected 0 invalid files, got: %d", results.InvalidFiles)
		}
	})

	// Test validation with invalid agent file
	t.Run("invalid agent file", func(t *testing.T) {
		// Create an invalid agent file
		invalidAgentContent := `agent:
  identity:
    name: "X"  # Too short
    id: "invalid-id"  # Invalid pattern
    version: "invalid"  # Invalid version
    description: "Short"  # Too short
    role: "X"  # Too short
    goal: "Short"  # Too short
  activation_prompt: []  # Empty
  principles: []  # Empty
  commands: {}  # Empty
`

		invalidAgentFile := filepath.Join(agentsDir, "invalid-agent.yaml")
		err = os.WriteFile(invalidAgentFile, []byte(invalidAgentContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write invalid agent file: %v", err)
		}

		validator, err := testFrameworkValidator(tempDir)
		if err != nil {
			t.Fatalf("Failed to create validator: %v", err)
		}
		results, err := validator.ValidateFramework()

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if results.IsValid() {
			t.Error("Expected invalid results, but validation passed")
		}

		if results.TotalFiles != 3 {
			t.Errorf("Expected 3 files, got: %d", results.TotalFiles)
		}

		if results.ValidFiles != 2 {
			t.Errorf("Expected 2 valid files, got: %d", results.ValidFiles)
		}

		if results.InvalidFiles != 1 {
			t.Errorf("Expected 1 invalid file, got: %d", results.InvalidFiles)
		}

		if results.TotalErrors == 0 {
			t.Error("Expected validation errors, but got none")
		}

		// Clean up
		os.Remove(invalidAgentFile)
	})
}

func TestFrameworkValidator_NoFrameworkDirectory(t *testing.T) {
	// Test with directory that has no .krci-ai directory
	tempDir := t.TempDir()

	validator, err := testFrameworkValidator(tempDir)
	if err != nil {
		t.Fatalf("Failed to create validator: %v", err)
	}
	_, err = validator.ValidateFramework()

	if err == nil {
		t.Error("Expected error for missing .krci-ai directory, got nil")
	}

	expectedErrorMsg := "no .krci-ai directory found"
	if err != nil && !containsString(err.Error(), expectedErrorMsg) {
		t.Errorf("Expected error message to contain '%s', got: %s", expectedErrorMsg, err.Error())
	}
}

func TestValidateAgentFile(t *testing.T) {
	tempDir := t.TempDir()
	validator, err := testFrameworkValidator(tempDir)
	if err != nil {
		t.Fatalf("Failed to create validator: %v", err)
	}

	// Test with valid agent file
	t.Run("valid agent file", func(t *testing.T) {
		validAgentContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
    description: "A test agent for validation testing"
    role: "Test Engineer"
    goal: "Test the validation system thoroughly"
  activation_prompt:
    - "You are a test agent"
  principles:
    - "Test thoroughly"
    - "Provide feedback"
    - "Document results"
  commands:
    help: "Show commands"
    chat: "Chat mode"
    exit: "Exit mode"
`

		agentFile := filepath.Join(tempDir, "test-agent.yaml")
		err := os.WriteFile(agentFile, []byte(validAgentContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write agent file: %v", err)
		}

		result := validator.validateAgentFile(agentFile)

		if !result.IsValid {
			t.Errorf("Expected valid result, but got invalid. Errors: %v", result.Errors)
		}

		if result.Type != "agent" {
			t.Errorf("Expected type 'agent', got: %s", result.Type)
		}

		if len(result.Errors) != 0 {
			t.Errorf("Expected no errors, got: %v", result.Errors)
		}
	})

	// Test with non-existent file
	t.Run("non-existent file", func(t *testing.T) {
		result := validator.validateAgentFile("/path/to/nonexistent/file.yaml")

		if result.IsValid {
			t.Error("Expected invalid result for non-existent file, but got valid")
		}

		if len(result.Errors) == 0 {
			t.Error("Expected errors for non-existent file, but got none")
		}
	})
}

func TestValidateTaskFile(t *testing.T) {
	tempDir := t.TempDir()
	validator, err := testFrameworkValidator(tempDir)
	if err != nil {
		t.Fatalf("Failed to create validator: %v", err)
	}

	// Test with valid task file
	t.Run("valid task file", func(t *testing.T) {
		taskContent := `# Test Task

This is a test task.

## Description

Test task description.
`

		taskFile := filepath.Join(tempDir, "test-task.md")
		err := os.WriteFile(taskFile, []byte(taskContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write task file: %v", err)
		}

		result := validator.validateTaskFile(taskFile)

		if !result.IsValid {
			t.Errorf("Expected valid result, but got invalid. Errors: %v", result.Errors)
		}

		if result.Type != "task" {
			t.Errorf("Expected type 'task', got: %s", result.Type)
		}

		if len(result.Errors) != 0 {
			t.Errorf("Expected no errors, got: %v", result.Errors)
		}
	})

	// Test with non-markdown file
	t.Run("non-markdown file", func(t *testing.T) {
		taskFile := filepath.Join(tempDir, "test-task.txt")
		err := os.WriteFile(taskFile, []byte("test content"), 0644)
		if err != nil {
			t.Fatalf("Failed to write task file: %v", err)
		}

		result := validator.validateTaskFile(taskFile)

		if result.IsValid {
			t.Error("Expected invalid result for non-markdown file, but got valid")
		}

		if len(result.Errors) == 0 {
			t.Error("Expected errors for non-markdown file, but got none")
		}
	})

	// Test with non-existent file
	t.Run("non-existent file", func(t *testing.T) {
		result := validator.validateTaskFile("/path/to/nonexistent/file.md")

		if result.IsValid {
			t.Error("Expected invalid result for non-existent file, but got valid")
		}

		if len(result.Errors) == 0 {
			t.Error("Expected errors for non-existent file, but got none")
		}
	})
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return strings.Contains(s, substr)
}
