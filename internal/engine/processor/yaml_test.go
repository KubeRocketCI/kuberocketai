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
package processor

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// stringPtr returns a pointer to the given string (helper for tests)
func stringPtr(s string) *string {
	return &s
}

func TestYAMLProcessor_ParseAgentFile(t *testing.T) {
	// For testing, use the file-based constructor
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	require.NoError(t, err)

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
		defer func() {
			err := os.Remove(tmpFile)
			require.NoError(t, err)
		}()

		agent, err := processor.ParseAgentFile(tmpFile)
		require.NoError(t, err)

		// Validate parsed content
		assert.Equal(t, "Test Agent", agent.Agent.Identity.Name)
		assert.Equal(t, "test-agent-v1", agent.Agent.Identity.ID)
		assert.Equal(t, "1.0.0", agent.Agent.Identity.Version)
		assert.Equal(t, "Test Engineer", agent.Agent.Identity.Role)
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
	chat: "Default chat mode" # Invalid indentation
    exit: "Exit test mode"
  tasks:
    - "./.krci-ai/tasks/test-task.md"
`

		tmpFile := createTempFile(t, malformedContent)
		defer func() {
			err := os.Remove(tmpFile)
			require.NoError(t, err)
		}()

		_, err := processor.ParseAgentFile(tmpFile)
		assert.Error(t, err)
	})

	t.Run("non-existent file", func(t *testing.T) {
		_, err := processor.ParseAgentFile("/path/to/nonexistent/file.yaml")
		assert.Error(t, err)
	})
}

func TestYAMLProcessor_ValidateAgent_WithSchema(t *testing.T) {
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	require.NoError(t, err)

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
				Tasks: []string{"./.krci-ai/tasks/test-task.md"},
			},
		}

		errors := processor.ValidateAgent(agent, "test.yaml")
		assert.Empty(t, errors)
	})

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
				Customization: nil,
				Commands: map[string]string{
					"help": "Show available commands",
					"chat": "Default chat mode",
					"exit": "Exit test mode",
				},
				Tasks: []string{"./.krci-ai/tasks/test-task.md"},
			},
		}

		errors := processor.ValidateAgent(agent, "test.yaml")
		assert.NotEmpty(t, errors)
		assert.Len(t, errors, 2)
	})

	t.Run("missing customization field - schema validation", func(t *testing.T) {
		// Create agent with missing customization field entirely
		agentJSON := `{
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
				},
				"tasks": ["./.krci-ai/tasks/test-task.md"]
			}
		}`

		var agent Agent
		err := json.Unmarshal([]byte(agentJSON), &agent)
		require.NoError(t, err)

		errors := processor.ValidateAgent(&agent, "test.yaml")
		assert.NotEmpty(t, errors)
		// The actual error message may vary based on schema validation implementation
	})

	t.Run("invalid agent - schema validation", func(t *testing.T) {
		agent := &Agent{
			Agent: AgentSpec{
				Identity: AgentIdentity{
					Name:        "Test Agent",
					ID:          "test-agent-v1",
					Version:     "invalid-version",
					Description: "Short", // Too short
					Role:        "Test Engineer",
					Goal:        "Test", // Too short
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
				Customization: stringPtr("test customization"),
				Commands: map[string]string{
					"help": "Show available commands",
					"chat": "Default chat mode",
					"exit": "Exit test mode",
				},
				Tasks: []string{"./.krci-ai/tasks/test-task.md"},
			},
		}

		errors := processor.ValidateAgent(agent, "test.yaml")
		assert.NotEmpty(t, errors)
	})
}

func TestYAMLProcessor_WithRealAgentFiles(t *testing.T) {
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	require.NoError(t, err)

	// Test with real agent files
	agentFiles := []string{
		"../../../cmd/krci-ai/assets/framework/core/agents/architect.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/dev.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/pm.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/qa.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/ba.yaml",
		"../../../cmd/krci-ai/assets/framework/core/agents/po.yaml",
	}

	for _, file := range agentFiles {
		t.Run(file, func(t *testing.T) {
			agent, err := processor.ParseAgentFile(file)
			require.NoError(t, err)

			errors := processor.ValidateAgent(agent, file)
			assert.Empty(t, errors, "Validation errors: %v", errors)
		})
	}
}

func TestAgentSpec_GetCustomization(t *testing.T) {
	tests := []struct {
		name           string
		customization  *string
		expectedResult string
	}{
		{
			name:           "nil customization",
			customization:  nil,
			expectedResult: "",
		},
		{
			name:           "empty customization",
			customization:  stringPtr(""),
			expectedResult: "",
		},
		{
			name:           "non-empty customization",
			customization:  stringPtr("test customization"),
			expectedResult: "test customization",
		},
		{
			name:           "whitespace customization",
			customization:  stringPtr("   "),
			expectedResult: "   ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec := AgentSpec{Customization: tt.customization}
			result := spec.GetCustomization()
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestAgentSpec_HasCustomization(t *testing.T) {
	tests := []struct {
		name          string
		customization *string
		expected      bool
	}{
		{
			name:          "nil customization",
			customization: nil,
			expected:      false,
		},
		{
			name:          "empty customization",
			customization: stringPtr(""),
			expected:      false,
		},
		{
			name:          "non-empty customization",
			customization: stringPtr("test customization"),
			expected:      true,
		},
		{
			name:          "whitespace customization",
			customization: stringPtr("   "),
			expected:      false,
		},
		{
			name:          "single space customization",
			customization: stringPtr(" "),
			expected:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec := AgentSpec{Customization: tt.customization}
			result := spec.HasCustomization()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      ValidationError
		expected string
	}{
		{
			name: "basic validation error",
			err: ValidationError{
				Field:   "identity.version",
				Value:   "invalid-version",
				Message: "must match pattern ^v?\\d+\\.\\d+\\.\\d+",
				File:    "agents/dev.yaml",
			},
			expected: "validation failed for field 'identity.version' in file 'agents/dev.yaml': must match pattern ^v?\\d+\\.\\d+\\.\\d+",
		},
		{
			name: "validation error with empty values",
			err: ValidationError{
				Field:   "",
				Value:   "",
				Message: "some validation message",
				File:    "test.yaml",
			},
			expected: "validation failed for field '' in file 'test.yaml': some validation message",
		},
		{
			name: "validation error with complex message",
			err: ValidationError{
				Field:   "identity.version",
				Value:   "invalid-version",
				Message: "must match pattern ^v?\\d+\\.\\d+\\.\\d+",
				File:    "agents/dev.yaml",
			},
			expected: "validation failed for field 'identity.version' in file 'agents/dev.yaml': must match pattern ^v?\\d+\\.\\d+\\.\\d+",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestYAMLProcessor_ParseAgentFileRaw(t *testing.T) {
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	require.NoError(t, err)

	t.Run("valid YAML file", func(t *testing.T) {
		validContent := `agent:
  identity:
    name: "Test Agent"
    id: "test-agent-v1"
    version: "1.0.0"
`
		tmpFile := createTempFile(t, validContent)
		defer func() {
			err := os.Remove(tmpFile)
			require.NoError(t, err)
		}()

		data, err := processor.ParseAgentFileRaw(tmpFile)
		require.NoError(t, err)
		assert.NotEmpty(t, data)
	})

	t.Run("non-existent file", func(t *testing.T) {
		_, err := processor.ParseAgentFileRaw("/path/to/nonexistent/file.yaml")
		assert.Error(t, err)
	})
}

func TestYAMLProcessor_ValidateAgentRaw(t *testing.T) {
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	require.NoError(t, err)

	t.Run("validate agent raw basic functionality", func(t *testing.T) {
		validAgentData := map[string]interface{}{
			"agent": map[string]interface{}{
				"identity": map[string]interface{}{
					"name":        "Test Agent",
					"id":          "test-agent-v1",
					"version":     "1.0.0",
					"description": "A test agent for validation testing with sufficient length",
					"role":        "Test Engineer",
					"goal":        "Test the validation system thoroughly with proper length",
					"icon":        "🧪",
				},
				"activation_prompt": []interface{}{
					"You are a test agent designed for validation testing",
					"Follow test protocols carefully and thoroughly",
				},
				"principles": []interface{}{
					"Always test thoroughly and document results with sufficient detail",
					"Provide clear and actionable feedback to users",
					"Maintain high quality standards throughout the testing process",
				},
				"customization": "",
				"commands": map[string]interface{}{
					"help": "Show available commands",
					"chat": "Default chat mode",
					"exit": "Exit test mode",
				},
				"tasks": []interface{}{"./.krci-ai/tasks/test-task.md"},
			},
		}

		errors := processor.ValidateAgentRaw(validAgentData, "test.yaml")
		assert.Empty(t, errors)
	})

	t.Run("invalid agent data", func(t *testing.T) {
		invalidAgentData := map[string]interface{}{
			"agent": map[string]interface{}{
				"identity": map[string]interface{}{
					"name":        "Test Agent",
					"id":          "test-agent-v1",
					"version":     "invalid-version",
					"description": "Short",
					"role":        "Test Engineer",
					"goal":        "Test",
					"icon":        "🧪",
				},
				"customization": "",
			},
		}

		errors := processor.ValidateAgentRaw(invalidAgentData, "test.yaml")
		assert.NotEmpty(t, errors)
	})
}

func TestYAMLProcessor_GetTaskMetadataSchemaPath(t *testing.T) {
	processor, err := NewYAMLProcessorFromFile("../../../cmd/krci-ai/assets/schemas/agent-schema.json")
	require.NoError(t, err)

	expectedPath := "assets/schemas/task-metadata.json"
	actualPath := processor.GetTaskMetadataSchemaPath()

	assert.Equal(t, expectedPath, actualPath, "GetTaskMetadataSchemaPath should return the correct schema path")
}

// Helper function to create temporary files for testing
func createTempFile(t *testing.T, content string) string {
	tmpFile, err := os.CreateTemp("", "agent_test_*.yaml")
	require.NoError(t, err)

	_, err = tmpFile.WriteString(content)
	require.NoError(t, err)

	err = tmpFile.Close()
	require.NoError(t, err)

	return tmpFile.Name()
}
