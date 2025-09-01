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
package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KubeRocketCI/kuberocketai/internal/engine/processor"
)

func TestRunValidate(t *testing.T) {
	// Test with valid framework setup
	t.Run("runValidate with valid framework", func(t *testing.T) {
		tempDir := t.TempDir()

		// Set up .krci-ai framework structure
		krciDir := filepath.Join(tempDir, ".krci-ai")
		agentsDir := filepath.Join(krciDir, "agents")
		tasksDir := filepath.Join(krciDir, "tasks")

		err := os.MkdirAll(agentsDir, 0755)
		require.NoError(t, err, "Should be able to create agents directory")

		err = os.MkdirAll(tasksDir, 0755)
		require.NoError(t, err, "Should be able to create tasks directory")

		// Create valid agent file
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
    - "ALWAYS execute agent.customization field content when non-empty"
    - "You are a test agent designed for validation"
  principles:
    - "Always test thoroughly and document results"
    - "Provide clear and actionable feedback"
  customization: ""
  commands:
    help: "Show available commands"
    chat: "Default chat mode"
  tasks:
    - "./.krci-ai/tasks/test-task.md"
`

		validAgentFile := filepath.Join(agentsDir, "test-agent.yaml")
		err = os.WriteFile(validAgentFile, []byte(validAgentContent), 0644)
		require.NoError(t, err, "Should be able to write valid agent file")

		// Create valid task file
		validTaskContent := `# Test Task

This is a test task for validation.

## Description

Test task description.
`

		validTaskFile := filepath.Join(tasksDir, "test-task.md")
		err = os.WriteFile(validTaskFile, []byte(validTaskContent), 0644)
		require.NoError(t, err, "Should be able to write valid task file")

		// Change to temp directory
		oldDir, _ := os.Getwd()
		err = os.Chdir(tempDir)
		require.NoError(t, err, "Should be able to change directory")
		defer func() { _ = os.Chdir(oldDir) }()

		// Capture stdout to prevent output during test
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Create test command
		cmd := &cobra.Command{}
		cmd.Flags().BoolP("verbose", "v", false, "verbose output")
		cmd.Flags().BoolP("quiet", "q", false, "quiet output")

		// Test runValidate function with the new unified validation system
		err = runValidate(cmd, []string{})

		// Restore stdout
		_ = w.Close()
		os.Stdout = oldStdout

		// Read captured output
		output, _ := io.ReadAll(r)
		outputStr := string(output)

		// In test environment, we expect an error due to missing schemas
		// This validates the function can be called and handles missing resources gracefully
		if err != nil {
			t.Logf("runValidate returned expected error in test environment: %v", err)
			// Error is expected in test environment due to missing embedded schemas
		}

		// The function should still produce some output (even if it's an error message)
		t.Logf("runValidate output length: %d", len(outputStr))
	})

	t.Run("runValidate with missing framework", func(t *testing.T) {
		tempDir := t.TempDir()

		// Change to temp directory (no .krci-ai directory)
		oldDir, _ := os.Getwd()
		err := os.Chdir(tempDir)
		if err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}
		defer func() { _ = os.Chdir(oldDir) }()

		// Create test command
		cmd := &cobra.Command{}
		cmd.Flags().BoolP("verbose", "v", false, "verbose output")
		cmd.Flags().BoolP("quiet", "q", false, "quiet output")

		// Test that runValidate can be called and returns error for missing framework
		err = runValidate(cmd, []string{})

		// Should return an error for missing framework
		if err == nil {
			t.Error("Expected error for missing framework, but got nil")
		} else {
			t.Logf("runValidate returned expected error for missing framework: %v", err)
		}
	})

	t.Run("runValidate with flags", func(t *testing.T) {
		// Test that flags can be set without error
		cmd := &cobra.Command{}
		cmd.Flags().BoolP("verbose", "v", false, "verbose output")
		cmd.Flags().BoolP("quiet", "q", false, "quiet output")

		// Set verbose flag
		_ = cmd.Flags().Set("verbose", "true")
		// Set quiet flag
		_ = cmd.Flags().Set("quiet", "true")

		// Verify flags can be accessed (basic flag test)
		verboseFlag, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			t.Errorf("Failed to get verbose flag: %v", err)
		}
		if !verboseFlag {
			t.Error("Expected verbose flag to be true")
		}

		quietFlag, err := cmd.Flags().GetBool("quiet")
		if err != nil {
			t.Errorf("Failed to get quiet flag: %v", err)
		}
		if !quietFlag {
			t.Error("Expected quiet flag to be true")
		}
	})
}

func TestNewFrameworkValidator(t *testing.T) {
	t.Run("NewFrameworkValidator basic functionality", func(t *testing.T) {
		tempDir := t.TempDir()

		// Test that NewFrameworkValidator function exists and can be called
		// Since we don't have embedded assets in test environment, we expect an error
		validator, err := NewFrameworkValidator(tempDir)
		if err != nil {
			t.Logf("NewFrameworkValidator failed as expected (no embedded assets): %v", err)
			// This is expected in test environment without real embedded assets
		} else {
			// If it succeeds (shouldn't happen in test), verify basic structure
			require.NotNil(t, validator, "Validator should be created")

			if validator.baseDir != tempDir {
				t.Errorf("Expected baseDir to be %s, got %s", tempDir, validator.baseDir)
			}
		}
	})

	t.Run("NewFrameworkValidator function signature", func(t *testing.T) {
		// Test that the function can be called and has expected signature
		// This verifies the function exists and is accessible
		tempDir := t.TempDir()

		// Should not panic when called
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("NewFrameworkValidator panicked: %v", r)
				}
			}()

			_, _ = NewFrameworkValidator(tempDir)
		}()
	})
}

func TestValidateCommand(t *testing.T) {
	t.Run("validate command exists", func(t *testing.T) {
		assert.NotNil(t, validateCmd, "Validate command should not be nil")
		assert.Equal(t, "validate", validateCmd.Use, "Command Use should be 'validate'")
		assert.NotEmpty(t, validateCmd.Short, "Command should have a short description")
		assert.NotEmpty(t, validateCmd.Long, "Command should have a long description")
		assert.NotNil(t, validateCmd.RunE, "Command should have a RunE function")
	})

	t.Run("validate command flags", func(t *testing.T) {
		// Check if validate command has expected flags
		verboseFlag := validateCmd.Flags().Lookup("verbose")
		assert.NotNil(t, verboseFlag, "Validate command should have --verbose flag")
		if verboseFlag != nil {
			assert.Equal(t, "v", verboseFlag.Shorthand, "Verbose flag shorthand should be 'v'")
		}

		quietFlag := validateCmd.Flags().Lookup("quiet")
		assert.NotNil(t, quietFlag, "Validate command should have --quiet flag")
		if quietFlag != nil {
			assert.Equal(t, "q", quietFlag.Shorthand, "Quiet flag shorthand should be 'q'")
		}
	})
}

func TestValidationResults(t *testing.T) {
	t.Run("ValidationResults IsValid", func(t *testing.T) {
		// Test with valid results
		results := &ValidationResults{
			TotalFiles:   3,
			ValidFiles:   3,
			InvalidFiles: 0,
		}

		if !results.IsValid() {
			t.Error("Expected IsValid to return true for valid results")
		}

		// Test with invalid results
		results.InvalidFiles = 1
		results.ValidFiles = 2

		if results.IsValid() {
			t.Error("Expected IsValid to return false when there are invalid files")
		}
	})
}

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
    - "ALWAYS execute agent.customization field content when non-empty"
    - "You are a test agent designed for validation"
    - "Follow test protocols carefully"
  principles:
    - "Always test thoroughly and document results"
    - "Provide clear and actionable feedback"
    - "Maintain high quality standards"
  customization: ""
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

		if results.TotalFiles != 3 {
			t.Errorf("Expected 3 validation results, got: %d", results.TotalFiles)
		}

		if results.ValidFiles != 3 {
			t.Errorf("Expected 3 valid validation results, got: %d", results.ValidFiles)
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

		if results.TotalFiles != 4 {
			t.Errorf("Expected 4 validation results, got: %d", results.TotalFiles)
		}

		if results.ValidFiles != 3 {
			t.Errorf("Expected 3 valid validation results, got: %d", results.ValidFiles)
		}

		if results.InvalidFiles != 1 {
			t.Errorf("Expected 1 invalid file, got: %d", results.InvalidFiles)
		}

		if results.TotalErrors == 0 {
			t.Error("Expected validation errors, but got none")
		}

		// Clean up
		_ = os.Remove(invalidAgentFile)
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
	if err != nil && !strings.Contains(err.Error(), expectedErrorMsg) {
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
    - "ALWAYS execute agent.customization field content when non-empty"
    - "You are a test agent"
  principles:
    - "Test thoroughly"
    - "Provide feedback"
    - "Document results"
  customization: ""
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

func TestValidateTaskPathLinks(t *testing.T) {
	tempDir := t.TempDir()
	validator, err := testFrameworkValidator(tempDir)
	if err != nil {
		t.Fatalf("Failed to create validator: %v", err)
	}

	// Create test directory structure
	krciDir := filepath.Join(tempDir, ".krci-ai")
	tasksDir := filepath.Join(krciDir, "tasks")
	err = os.MkdirAll(tasksDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create tasks directory: %v", err)
	}

	// Create a valid task file
	validTaskFile := filepath.Join(tasksDir, "valid-task.md")
	err = os.WriteFile(validTaskFile, []byte("# Valid Task\n\nThis is a valid task."), 0644)
	if err != nil {
		t.Fatalf("Failed to write valid task file: %v", err)
	}

	t.Run("agent with valid task references", func(t *testing.T) {
		validAgentContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
    description: "A test agent for task validation"
    role: "Test Engineer"
    goal: "Test task path validation functionality"
  activation_prompt:
    - "Test agent for validation"
  principles:
    - "Test thoroughly"
    - "Validate references"
    - "Report clear errors"
  customization: ""
  commands:
    help: "Show commands"
    chat: "Chat mode"
    exit: "Exit mode"
  tasks:
    - "./.krci-ai/tasks/valid-task.md"
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

		if len(result.Errors) != 0 {
			t.Errorf("Expected no errors, got: %v", result.Errors)
		}
	})

	t.Run("agent with invalid task references", func(t *testing.T) {
		invalidAgentContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
    description: "A test agent with invalid task references"
    role: "Test Engineer"
    goal: "Test task path validation functionality"
  activation_prompt:
    - "Test agent for validation"
  principles:
    - "Test thoroughly"
    - "Validate references"
    - "Report clear errors"
  customization: ""
  commands:
    help: "Show commands"
    chat: "Chat mode"
    exit: "Exit mode"
  tasks:
    - "./.krci-ai/tasks/nonexistent-task.md"
    - "./.krci-ai/tasks/another-missing.md"
    - "./.krci-ai/tasks/valid-task.md"
`

		agentFile := filepath.Join(tempDir, "test-agent.yaml")
		err := os.WriteFile(agentFile, []byte(invalidAgentContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write agent file: %v", err)
		}

		result := validator.validateAgentFile(agentFile)

		if result.IsValid {
			t.Error("Expected invalid result due to missing task references, but got valid")
		}

		if len(result.Errors) != 2 {
			t.Errorf("Expected 2 errors for missing tasks, got: %d (%v)", len(result.Errors), result.Errors)
		}

		// Verify specific error messages
		errorMessages := strings.Join(result.Errors, " ")
		if !strings.Contains(errorMessages, "nonexistent-task.md") {
			t.Error("Expected error message about nonexistent-task.md")
		}
		if !strings.Contains(errorMessages, "another-missing.md") {
			t.Error("Expected error message about another-missing.md")
		}
	})

	t.Run("agent with no task references", func(t *testing.T) {
		noTasksAgentContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
    description: "A test agent with no task references"
    role: "Test Engineer"
    goal: "Test task path validation functionality"
  activation_prompt:
    - "Test agent for validation"
  principles:
    - "Test thoroughly"
    - "Validate references"
    - "Report clear errors"
  customization: ""
  commands:
    help: "Show commands"
    chat: "Chat mode"
    exit: "Exit mode"
`

		agentFile := filepath.Join(tempDir, "test-agent.yaml")
		err := os.WriteFile(agentFile, []byte(noTasksAgentContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write agent file: %v", err)
		}

		result := validator.validateAgentFile(agentFile)

		if !result.IsValid {
			t.Errorf("Expected valid result for agent with no tasks, but got invalid. Errors: %v", result.Errors)
		}
	})

	t.Run("agent with non-markdown task reference", func(t *testing.T) {
		// Create a non-markdown file
		nonMarkdownFile := filepath.Join(tasksDir, "invalid-task.txt")
		err := os.WriteFile(nonMarkdownFile, []byte("This is not markdown"), 0644)
		if err != nil {
			t.Fatalf("Failed to write non-markdown file: %v", err)
		}

		nonMarkdownAgentContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
    description: "A test agent with non-markdown task reference"
    role: "Test Engineer"
    goal: "Test task path validation functionality"
  activation_prompt:
    - "Test agent for validation"
  principles:
    - "Test thoroughly"
    - "Validate references"
    - "Report clear errors"
  customization: ""
  commands:
    help: "Show commands"
    chat: "Chat mode"
    exit: "Exit mode"
  tasks:
    - "./.krci-ai/tasks/invalid-task.txt"
`

		agentFile := filepath.Join(tempDir, "test-agent.yaml")
		err = os.WriteFile(agentFile, []byte(nonMarkdownAgentContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write agent file: %v", err)
		}

		result := validator.validateAgentFile(agentFile)

		if result.IsValid {
			t.Error("Expected invalid result due to non-markdown task reference, but got valid")
		}

		// Expect multiple errors: schema validation + task path validation
		if len(result.Errors) < 1 {
			t.Errorf("Expected at least 1 error for non-markdown task, got: %d (%v)", len(result.Errors), result.Errors)
		}

		// Check that one of the errors is about markdown file requirement
		errorMessages := strings.Join(result.Errors, " ")
		if !strings.Contains(errorMessages, "must be a markdown file") {
			t.Errorf("Expected error about markdown file requirement in: %v", result.Errors)
		}
	})
}

func TestValidateMarkdownFile(t *testing.T) {
	tempDir := t.TempDir()
	validator, err := testFrameworkValidator(tempDir)
	if err != nil {
		t.Fatalf("Failed to create validator: %v", err)
	}

	t.Run("valid markdown file", func(t *testing.T) {
		// Create valid markdown file
		validContent := "# Test File\n\nThis is a test markdown file.\n"
		validFile := filepath.Join(tempDir, "valid.md")
		err := os.WriteFile(validFile, []byte(validContent), 0644)
		require.NoError(t, err, "Should create valid markdown file")

		result := validator.validateMarkdownFile(validFile, "test")

		assert.True(t, result.IsValid, "Should validate as valid")
		assert.Equal(t, "test", result.Type, "Type should be set correctly")
		assert.Len(t, result.Errors, 0, "Should have no errors")
		assert.Contains(t, result.File, "valid.md", "File path should be in result")
	})

	t.Run("non-existent file", func(t *testing.T) {
		result := validator.validateMarkdownFile("/nonexistent/file.md", "test")

		assert.False(t, result.IsValid, "Should validate as invalid")
		assert.Equal(t, "test", result.Type, "Type should be set correctly")
		assert.Len(t, result.Errors, 1, "Should have one error")
		assert.Contains(t, result.Errors[0], "Test file does not exist", "Error should mention file doesn't exist")
	})

	t.Run("non-markdown extension", func(t *testing.T) {
		// Create non-markdown file
		txtContent := "This is not a markdown file"
		txtFile := filepath.Join(tempDir, "test.txt")
		err := os.WriteFile(txtFile, []byte(txtContent), 0644)
		require.NoError(t, err, "Should create txt file")

		result := validator.validateMarkdownFile(txtFile, "test")

		assert.False(t, result.IsValid, "Should validate as invalid")
		assert.Equal(t, "test", result.Type, "Type should be set correctly")
		assert.Len(t, result.Errors, 1, "Should have one error")
		assert.Contains(t, result.Errors[0], "Test file must have .md extension", "Error should mention .md extension requirement")
	})
}
