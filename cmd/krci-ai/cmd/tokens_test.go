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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KubeRocketCI/kuberocketai/internal/tokens"
)

// TestRunTokensCommandValidation tests flag validation in runTokensCommand
func TestRunTokensCommandValidation(t *testing.T) {
	tests := []struct {
		name        string
		agent       string
		all         bool
		expectError bool
		errorMsg    string
	}{
		{
			name:        "no flags provided",
			agent:       "",
			all:         false,
			expectError: true,
			errorMsg:    "either --agent, --all, or --bundle flag must be specified",
		},
		{
			name:        "agent flag provided",
			agent:       "pm",
			all:         false,
			expectError: false,
		},
		{
			name:        "all flag provided",
			agent:       "",
			all:         true,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create command with flags
			cmd := &cobra.Command{}
			cmd.Flags().String("agent", "", "Analyze tokens for a specific agent")
			cmd.Flags().Bool("all", false, "Analyze tokens for all agents in the project")
			cmd.Flags().String("bundle", "", "Analyze tokens for bundle configuration")
			cmd.Flags().Bool("json", false, "Output results in JSON format")
			cmd.Flags().Duration("timeout", 30*time.Second, "Timeout for token analysis")

			// Set flag values
			cmd.Flags().Set("agent", tt.agent)
			cmd.Flags().Set("all", fmt.Sprintf("%t", tt.all))

			// Mock the calculator creation to avoid actual file system operations
			// Note: In a real test, we would need to mock the NewCalculator function

			err := runTokensCommand(cmd, []string{})

			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				// For valid flags, we expect a different error (framework not installed)
				// since we don't have a real framework setup in tests
				require.Error(t, err)
				// The error could be about framework not being installed or calculator initialization
				assert.True(t,
					strings.Contains(err.Error(), "failed to initialize token calculator") ||
						strings.Contains(err.Error(), "framework not installed") ||
						strings.Contains(err.Error(), "failed to discover agents"),
					"Expected error about calculator or framework, got: %s", err.Error())
			}
		})
	}
}

// Note: runAgentTokenAnalysis and runProjectTokenAnalysis functions are tightly coupled
// to the concrete tokens.Calculator type and would require dependency injection
// to be properly unit tested. These functions are better tested through integration tests.

// TestOutputAgentTokensJSON tests JSON output for agent tokens
func TestOutputAgentTokensJSON(t *testing.T) {
	agentInfo := &tokens.AgentTokenInfo{
		AgentName:   "pm",
		AgentFile:   "pm.yaml",
		TotalTokens: 1500,
		Assets: []tokens.AssetTokenInfo{
			{Path: "pm.yaml", Tokens: 500},
		},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputAgentTokensJSON(agentInfo)

	w.Close()
	os.Stdout = oldStdout
	output, _ := io.ReadAll(r)

	assert.NoError(t, err)

	// Verify JSON structure
	var result tokens.AgentTokenInfo
	err = json.Unmarshal(output, &result)
	assert.NoError(t, err)
	assert.Equal(t, "pm", result.AgentName)
	assert.Equal(t, 1500, result.TotalTokens)
}

// TestOutputProjectTokensJSON tests JSON output for project tokens
func TestOutputProjectTokensJSON(t *testing.T) {
	projectInfo := &tokens.ProjectTokenInfo{
		TotalTokens: 5000,
		Agents: []tokens.AgentTokenInfo{
			{AgentName: "pm", TotalTokens: 2000},
		},
		Breakdown: tokens.TokenBreakdown{
			Agents: 1000,
		},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputProjectTokensJSON(projectInfo)

	w.Close()
	os.Stdout = oldStdout
	output, _ := io.ReadAll(r)

	assert.NoError(t, err)

	// Verify JSON structure
	var result tokens.ProjectTokenInfo
	err = json.Unmarshal(output, &result)
	assert.NoError(t, err)
	assert.Equal(t, 5000, result.TotalTokens)
	assert.Len(t, result.Agents, 1)
}

// TestOutputAgentTokensTable tests table output for agent tokens
func TestOutputAgentTokensTable(t *testing.T) {
	agentInfo := &tokens.AgentTokenInfo{
		AgentName:   "pm",
		AgentFile:   "pm.yaml",
		TotalTokens: 1500,
		Assets: []tokens.AssetTokenInfo{
			{Path: "pm.yaml", Tokens: 500},
		},
		Dependencies: struct {
			Tasks     []tokens.AssetTokenInfo `json:"tasks"`
			Templates []tokens.AssetTokenInfo `json:"templates"`
			DataFiles []tokens.AssetTokenInfo `json:"data_files"`
			TasksRef  []tokens.AssetTokenInfo `json:"tasks_ref"`
		}{
			Tasks: []tokens.AssetTokenInfo{
				{Path: "create-story.md", Tokens: 300},
			},
			Templates: []tokens.AssetTokenInfo{
				{Path: "story.md", Tokens: 200},
			},
			DataFiles: []tokens.AssetTokenInfo{
				{Path: "best-practices.md", Tokens: 500},
			},
			TasksRef: []tokens.AssetTokenInfo{
				{Path: "create-story.md", Tokens: 300},
			},
		},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputAgentTokensTable(agentInfo)

	w.Close()
	os.Stdout = oldStdout
	output, _ := io.ReadAll(r)

	assert.NoError(t, err)
	outputStr := string(output)
	assert.Contains(t, outputStr, "pm")
	assert.Contains(t, outputStr, "1500")
	assert.Contains(t, outputStr, "create-story.md")
	assert.Contains(t, outputStr, "Tasks:")
	assert.Contains(t, outputStr, "Templates:")
	assert.Contains(t, outputStr, "Data Files:")
}

// TestOutputAgentTokensTableNoDependencies tests table output with no dependencies
func TestOutputAgentTokensTableNoDependencies(t *testing.T) {
	agentInfo := &tokens.AgentTokenInfo{
		AgentName:   "simple",
		AgentFile:   "simple.yaml",
		TotalTokens: 500,
		Assets: []tokens.AssetTokenInfo{
			{Path: "simple.yaml", Tokens: 500},
		},
		Dependencies: struct {
			Tasks     []tokens.AssetTokenInfo `json:"tasks"`
			Templates []tokens.AssetTokenInfo `json:"templates"`
			DataFiles []tokens.AssetTokenInfo `json:"data_files"`
			TasksRef  []tokens.AssetTokenInfo `json:"tasks_ref"`
		}{}, // No dependencies
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputAgentTokensTable(agentInfo)

	w.Close()
	os.Stdout = oldStdout
	output, _ := io.ReadAll(r)

	assert.NoError(t, err)
	outputStr := string(output)
	assert.Contains(t, outputStr, "simple")
	assert.Contains(t, outputStr, "500")
	assert.Contains(t, outputStr, "No dependencies found")
}

// TestOutputProjectTokensTable tests table output for project tokens
func TestOutputProjectTokensTable(t *testing.T) {
	projectInfo := &tokens.ProjectTokenInfo{
		TotalTokens: 5000,
		Agents: []tokens.AgentTokenInfo{
			{AgentName: "pm", TotalTokens: 2000},
			{AgentName: "dev", TotalTokens: 3000},
		},
		Breakdown: tokens.TokenBreakdown{
			Agents:    1000,
			Tasks:     1500,
			Templates: 1000,
			DataFiles: 1500,
		},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputProjectTokensTable(projectInfo)

	w.Close()
	os.Stdout = oldStdout
	output, _ := io.ReadAll(r)

	assert.NoError(t, err)
	outputStr := string(output)
	assert.Contains(t, outputStr, "5000")
	assert.Contains(t, outputStr, "pm")
	assert.Contains(t, outputStr, "dev")
	assert.Contains(t, outputStr, "2000")
	assert.Contains(t, outputStr, "3000")
	assert.Contains(t, outputStr, "Agents: 1000")
	assert.Contains(t, outputStr, "Tasks: 1500")
}

// TestHandleTokenError tests error handling with different error types
func TestHandleTokenError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		jsonMode bool
		expected string
	}{
		{
			name:     "agent not found error in interactive mode",
			err:      errors.New("agent 'nonexistent' not found"),
			jsonMode: false,
			expected: "Agent not found",
		},
		{
			name:     "framework not installed error",
			err:      errors.New("framework not installed"),
			jsonMode: false,
			expected: "Framework not installed",
		},
		{
			name:     "generic error in JSON mode",
			err:      errors.New("some generic error"),
			jsonMode: true,
			expected: "some generic error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stderr for interactive mode
			if !tt.jsonMode {
				oldStderr := os.Stderr
				r, w, _ := os.Pipe()
				os.Stderr = w

				err := handleTokenError(tt.err, tt.jsonMode)

				w.Close()
				os.Stderr = oldStderr
				output, _ := io.ReadAll(r)

				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
				if tt.expected != "" {
					assert.Contains(t, string(output), tt.expected)
				}
			} else {
				// JSON mode should return the raw error
				err := handleTokenError(tt.err, tt.jsonMode)
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

// TestHandleTokenErrorWithSuggestions tests error handling with suggestions
func TestHandleTokenErrorWithSuggestions(t *testing.T) {
	err := errors.New("agent 'test' not found")

	// Capture stderr
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	resultErr := handleTokenError(err, false) // false = not JSON mode

	w.Close()
	os.Stderr = oldStderr
	output, _ := io.ReadAll(r)

	assert.Error(t, resultErr)
	assert.Equal(t, err, resultErr)
	outputStr := string(output)
	assert.Contains(t, outputStr, "Suggestions:")
	assert.Contains(t, outputStr, "krci-ai list")
}

// TestTokensCommandIntegration tests the command integration
func TestTokensCommandIntegration(t *testing.T) {
	// Test command exists in root
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "tokens" {
			found = true
			break
		}
	}
	assert.True(t, found, "tokens command should be added to root command")
}
