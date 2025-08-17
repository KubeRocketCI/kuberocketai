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
package tokens

import (
	"context"
	"embed"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testEmbedFS embed.FS

// MockDiscovery implements DiscoveryInterface for testing
type MockDiscovery struct {
	agents           []assets.AgentInfo
	agentDeps        []assets.AgentDependencyInfo
	getAgentError    error
	discoverError    error
	validateError    error
	targetAgent      *assets.AgentDependencyInfo
	targetAgentError error
}

func (m *MockDiscovery) DiscoverAgents() ([]assets.AgentInfo, error) {
	if m.discoverError != nil {
		return nil, m.discoverError
	}
	return m.agents, nil
}

func (m *MockDiscovery) DiscoverAgentsWithDependencies(agentNames ...string) ([]assets.AgentDependencyInfo, error) {
	if m.discoverError != nil {
		return nil, m.discoverError
	}
	return m.agentDeps, nil
}

func (m *MockDiscovery) DiscoverAgentWithDependencies(shortName string) (assets.AgentDependencyInfo, error) {
	if m.targetAgentError != nil {
		return assets.AgentDependencyInfo{}, m.targetAgentError
	}
	if m.targetAgent != nil {
		return *m.targetAgent, nil
	}
	return assets.AgentDependencyInfo{}, errors.New("agent not found")
}

func (m *MockDiscovery) GetAgentByShortName(shortName string) (*assets.AgentInfo, error) {
	if m.getAgentError != nil {
		return nil, m.getAgentError
	}
	for _, agent := range m.agents {
		if agent.ShortName == shortName {
			return &agent, nil
		}
	}
	return nil, errors.New("agent not found")
}

func (m *MockDiscovery) ValidateAgentStructure(filePath string) error {
	return m.validateError
}

// createTestCalculator creates a Calculator with mocked dependencies for testing
func createTestCalculator(discovery *MockDiscovery) *Calculator {
	calc := &MockTokenCalculator{tokenCount: 100}
	engine := NewEngine(calc)

	return NewCalculatorWithDependencies(engine, discovery, "/test/dir")
}

// setupTestFiles creates temporary test files for testing file operations
func setupTestFiles(t *testing.T) string {
	tempDir := t.TempDir()

	// Create test agent file
	agentContent := `name: Test Agent
description: A test agent for unit testing
role: tester
goal: Test the token calculation functionality
icon: 🧪
`
	agentFile := filepath.Join(tempDir, "test-agent.yaml")
	err := os.WriteFile(agentFile, []byte(agentContent), 0644)
	require.NoError(t, err, "Failed to create test agent file")

	// Create krci-ai directory structure
	krciDir := filepath.Join(tempDir, assets.KrciAIDir)
	tasksDir := filepath.Join(krciDir, assets.TasksDir)
	templatesDir := filepath.Join(krciDir, assets.TemplatesDir)
	dataDir := filepath.Join(krciDir, assets.DataDir)

	for _, dir := range []string{krciDir, tasksDir, templatesDir, dataDir} {
		err := os.MkdirAll(dir, 0755)
		require.NoError(t, err, "Failed to create directory %s", dir)
	}

	// Create test task file
	taskContent := `# Test Task
This is a test task for unit testing token calculation.
`
	taskFile := filepath.Join(tasksDir, "test-task.md")
	err = os.WriteFile(taskFile, []byte(taskContent), 0644)
	require.NoError(t, err, "Failed to create test task file")

	// Create test template file
	templateContent := `# Test Template
This is a test template for unit testing.
`
	templateFile := filepath.Join(templatesDir, "test-template.md")
	err = os.WriteFile(templateFile, []byte(templateContent), 0644)
	require.NoError(t, err, "Failed to create test template file")

	// Create test data file
	dataContent := `# Test Data
This is test data for unit testing.
`
	dataFile := filepath.Join(dataDir, "test-data.md")
	err = os.WriteFile(dataFile, []byte(dataContent), 0644)
	require.NoError(t, err, "Failed to create test data file")

	return tempDir
}

func TestNewCalculator(t *testing.T) {
	tempDir := setupTestFiles(t)

	calc, err := NewCalculator(tempDir, testEmbedFS)
	require.NoError(t, err, "NewCalculator() should not fail")
	assert.NotNil(t, calc, "NewCalculator() should not return nil")
	assert.NotNil(t, calc.engine, "Calculator engine should not be nil")
	assert.NotNil(t, calc.discovery, "Calculator discovery should not be nil")
	assert.Equal(t, tempDir, calc.targetDir, "Calculator targetDir mismatch")
}

func TestNewCalculatorWithDependencies(t *testing.T) {
	// Test dependency injection constructor
	mockDiscovery := &MockDiscovery{
		agents: []assets.AgentInfo{
			{ShortName: "test", Name: "Test Agent"},
		},
	}
	mockTokenizer := &MockTokenCalculator{tokenCount: 42}
	engine := NewEngine(mockTokenizer)

	calc := NewCalculatorWithDependencies(engine, mockDiscovery, "/test/dir")

	assert.NotNil(t, calc, "NewCalculatorWithDependencies should not return nil")
	assert.Equal(t, "/test/dir", calc.targetDir, "Target directory should be set correctly")
	assert.Equal(t, engine, calc.engine, "Engine should be set correctly")
	assert.Equal(t, mockDiscovery, calc.discovery, "Discovery should be set correctly")
}

func TestCalculator_CalculateAgentTokens(t *testing.T) {
	tempDir := setupTestFiles(t)

	tests := []struct {
		name              string
		agentShortName    string
		mockDiscovery     *MockDiscovery
		expectedError     string
		expectedMinTokens int
	}{
		{
			name:           "successful calculation",
			agentShortName: "test",
			mockDiscovery: &MockDiscovery{
				targetAgent: &assets.AgentDependencyInfo{
					AgentInfo: assets.AgentInfo{
						Name:      "Test Agent",
						ShortName: "test",
						FilePath:  filepath.Join(tempDir, "test-agent.yaml"),
					},
					Tasks:     []string{"test-task.md"},
					Templates: []string{"test-template.md"},
					DataFiles: []string{"test-data.md"},
				},
			},
			expectedMinTokens: 100, // At least 100 tokens from mock
		},
		{
			name:           "agent discovery error",
			agentShortName: "nonexistent",
			mockDiscovery: &MockDiscovery{
				targetAgentError: errors.New("agent not found"),
			},
			expectedError: "failed to discover agents: agent not found",
		},
		{
			name:           "file read error",
			agentShortName: "test",
			mockDiscovery: &MockDiscovery{
				targetAgent: &assets.AgentDependencyInfo{
					AgentInfo: assets.AgentInfo{
						Name:      "Test Agent",
						ShortName: "test",
						FilePath:  "/nonexistent/path/agent.yaml",
					},
				},
			},
			expectedError: "failed to read agent file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := createTestCalculator(tt.mockDiscovery)
			calc.targetDir = tempDir

			ctx := context.Background()
			result, err := calc.CalculateAgentTokens(ctx, tt.agentShortName)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected error but got none")
				assert.Contains(t, err.Error(), tt.expectedError, "Error message should contain expected text")
				return
			}

			assert.NoError(t, err, "Unexpected error")
			require.NotNil(t, result, "Result should not be nil")
			assert.GreaterOrEqual(t, result.TotalTokens, tt.expectedMinTokens, "Token count should meet minimum")
			assert.Equal(t, tt.mockDiscovery.targetAgent.Name, result.AgentName, "AgentName mismatch")
		})
	}
}

func TestCalculator_CalculateAllTokens(t *testing.T) {
	tempDir := setupTestFiles(t)

	tests := []struct {
		name              string
		mockDiscovery     *MockDiscovery
		expectedError     string
		expectedMinTokens int
		expectedAgents    int
	}{
		{
			name: "successful calculation",
			mockDiscovery: &MockDiscovery{
				agentDeps: []assets.AgentDependencyInfo{
					{
						AgentInfo: assets.AgentInfo{
							Name:      "Agent 1",
							ShortName: "agent1",
							FilePath:  filepath.Join(tempDir, "test-agent.yaml"),
						},
						Tasks:     []string{"test-task.md"},
						Templates: []string{"test-template.md"},
						DataFiles: []string{"test-data.md"},
					},
					{
						AgentInfo: assets.AgentInfo{
							Name:      "Agent 2",
							ShortName: "agent2",
							FilePath:  filepath.Join(tempDir, "test-agent.yaml"),
						},
						Tasks:     []string{"test-task.md"},
						Templates: []string{},
						DataFiles: []string{},
					},
				},
			},
			expectedMinTokens: 200, // At least 200 tokens (100 per agent)
			expectedAgents:    2,
		},
		{
			name: "discovery error",
			mockDiscovery: &MockDiscovery{
				discoverError: errors.New("discovery failed"),
			},
			expectedError: "failed to discover agents: discovery failed",
		},
		{
			name: "empty project",
			mockDiscovery: &MockDiscovery{
				agentDeps: []assets.AgentDependencyInfo{},
			},
			expectedMinTokens: 0,
			expectedAgents:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := createTestCalculator(tt.mockDiscovery)
			calc.targetDir = tempDir

			ctx := context.Background()
			result, err := calc.CalculateAllTokens(ctx)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected error but got none")
				assert.Contains(t, err.Error(), tt.expectedError, "Error message should contain expected text")
				return
			}

			assert.NoError(t, err, "Unexpected error")
			require.NotNil(t, result, "Result should not be nil")
			assert.GreaterOrEqual(t, result.TotalTokens, tt.expectedMinTokens, "Token count should meet minimum")
			assert.Len(t, result.Agents, tt.expectedAgents, "Agent count mismatch")

			// Verify breakdown is calculated
			expectedBreakdownTotal := result.Breakdown.Agents + result.Breakdown.Tasks +
				result.Breakdown.Templates + result.Breakdown.DataFiles
			assert.LessOrEqual(t, expectedBreakdownTotal, result.TotalTokens, "Breakdown total should not exceed total tokens")
		})
	}
}

func TestCalculator_ValidateAgentExists(t *testing.T) {
	tests := []struct {
		name          string
		agentName     string
		mockDiscovery *MockDiscovery
		expectedError string
	}{
		{
			name:      "existing agent",
			agentName: "test",
			mockDiscovery: &MockDiscovery{
				agents: []assets.AgentInfo{
					{ShortName: "test", Name: "Test Agent"},
				},
			},
		},
		{
			name:      "non-existent agent",
			agentName: "nonexistent",
			mockDiscovery: &MockDiscovery{
				getAgentError: errors.New("agent not found"),
			},
			expectedError: "agent validation failed: agent not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := createTestCalculator(tt.mockDiscovery)

			err := calc.ValidateAgentExists(tt.agentName)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected error but got none")
				assert.Equal(t, tt.expectedError, err.Error(), "Error message mismatch")
				return
			}

			assert.NoError(t, err, "Unexpected error")
		})
	}
}
