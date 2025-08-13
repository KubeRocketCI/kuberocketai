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
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockTokenCalculator implements TokenCalculator for testing
type MockTokenCalculator struct {
	tokenCount int
	err        error
}

func (m *MockTokenCalculator) CalculateTokens(ctx context.Context, text string) (int, error) {
	if m.err != nil {
		return 0, m.err
	}
	if text == "" {
		return 0, nil
	}
	if m.tokenCount > 0 {
		return m.tokenCount, nil
	}
	// Default behavior: return length of text as token count for testing
	return len(text), nil
}

func TestNewEngine(t *testing.T) {
	calc := &MockTokenCalculator{}
	engine := NewEngine(calc)

	assert.NotNil(t, engine, "NewEngine should not return nil")
	assert.Equal(t, calc, engine.calculator, "Engine calculator should be set correctly")
}

func TestEngine_CalculateTokens(t *testing.T) {
	tests := []struct {
		name           string
		calculator     TokenCalculator
		text           string
		expectedErr    string
		expectedTokens int
	}{
		{
			name:        "nil calculator",
			calculator:  nil,
			text:        "test",
			expectedErr: "no token calculator configured",
		},
		{
			name:           "successful calculation",
			calculator:     &MockTokenCalculator{tokenCount: 42},
			text:           "test text",
			expectedTokens: 42,
		},
		{
			name:        "calculator error",
			calculator:  &MockTokenCalculator{err: errors.New("calculation failed")},
			text:        "test",
			expectedErr: "calculation failed",
		},
		{
			name:           "empty text",
			calculator:     &MockTokenCalculator{},
			text:           "",
			expectedTokens: 0,
		},
		{
			name:           "default behavior",
			calculator:     &MockTokenCalculator{}, // Uses default length-based calculation
			text:           "hello",
			expectedTokens: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewEngine(tt.calculator)
			ctx := context.Background()

			tokens, err := engine.CalculateTokens(ctx, tt.text)

			if tt.expectedErr != "" {
				assert.Error(t, err, "Expected error but got none")
				assert.Equal(t, tt.expectedErr, err.Error(), "Error message mismatch")
				return
			}

			assert.NoError(t, err, "Unexpected error")
			assert.Equal(t, tt.expectedTokens, tokens, "Token count mismatch")
		})
	}
}

func TestEngine_CalculateAssetTokens(t *testing.T) {
	tests := []struct {
		name           string
		calculator     TokenCalculator
		path           string
		assetType      string
		content        string
		expectedTokens int
		expectedErr    string
	}{
		{
			name:           "successful asset calculation",
			calculator:     &MockTokenCalculator{tokenCount: 100},
			path:           "/path/to/asset.yaml",
			assetType:      "agent",
			content:        "name: test-agent\ndescription: Test agent",
			expectedTokens: 100,
		},
		{
			name:        "calculator error",
			calculator:  &MockTokenCalculator{err: errors.New("encoding failed")},
			path:        "/path/to/asset.yaml",
			assetType:   "task",
			content:     "some content",
			expectedErr: "failed to calculate tokens for /path/to/asset.yaml: encoding failed",
		},
		{
			name:           "empty content",
			calculator:     &MockTokenCalculator{},
			path:           "/path/to/empty.yaml",
			assetType:      "template",
			content:        "",
			expectedTokens: 0,
		},
		{
			name:           "data file asset",
			calculator:     &MockTokenCalculator{tokenCount: 250},
			path:           "/path/to/data.md",
			assetType:      "data",
			content:        "# Data File\nThis is a data file with content.",
			expectedTokens: 250,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewEngine(tt.calculator)
			ctx := context.Background()

			asset, err := engine.CalculateAssetTokens(ctx, tt.path, tt.assetType, tt.content)

			if tt.expectedErr != "" {
				assert.Error(t, err, "Expected error but got none")
				assert.Equal(t, tt.expectedErr, err.Error(), "Error message mismatch")
				return
			}

			assert.NoError(t, err, "Unexpected error")

			// Verify asset structure
			assert.Equal(t, tt.path, asset.Path, "Path mismatch")
			assert.Equal(t, tt.assetType, asset.Type, "Type mismatch")
			assert.Equal(t, tt.expectedTokens, asset.Tokens, "Token count mismatch")
			assert.Equal(t, tt.content, asset.Content, "Content mismatch")
		})
	}
}

// Test data structures and JSON serialization
func TestAssetTokenInfo_JSON(t *testing.T) {
	asset := AssetTokenInfo{
		Path:    "/path/to/test.yaml",
		Type:    "agent",
		Tokens:  150,
		Content: "test content",
	}

	// Test JSON marshaling
	data, err := json.Marshal(asset)
	require.NoError(t, err, "Failed to marshal AssetTokenInfo")

	// Verify that Content is not included in JSON (json:"-" tag)
	var jsonMap map[string]interface{}
	err = json.Unmarshal(data, &jsonMap)
	require.NoError(t, err, "Failed to unmarshal JSON")

	assert.NotContains(t, jsonMap, "Content", "Content field should not be included in JSON output")

	// Verify other fields are present
	expectedFields := []string{"path", "type", "tokens"}
	for _, field := range expectedFields {
		assert.Contains(t, jsonMap, field, "Field '%s' missing from JSON output", field)
	}

	// Test JSON unmarshaling
	var unmarshaled AssetTokenInfo
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err, "Failed to unmarshal AssetTokenInfo")

	assert.Equal(t, asset.Path, unmarshaled.Path, "Path mismatch after unmarshal")
	assert.Equal(t, asset.Type, unmarshaled.Type, "Type mismatch after unmarshal")
	assert.Equal(t, asset.Tokens, unmarshaled.Tokens, "Tokens mismatch after unmarshal")
	assert.Empty(t, unmarshaled.Content, "Content should be empty after unmarshal (not included in JSON)")
}

func TestAgentTokenInfo_JSON(t *testing.T) {
	agent := AgentTokenInfo{
		AgentName:      "Test Agent",
		AgentFile:      "/path/to/agent.yaml",
		AgentShortName: "test",
		TotalTokens:    500,
		Assets: []AssetTokenInfo{
			{
				Path:    "/path/to/agent.yaml",
				Type:    "agent",
				Tokens:  100,
				Content: "agent content",
			},
		},
	}

	// Add some dependencies
	agent.Dependencies.Tasks = []AssetTokenInfo{
		{Path: "/path/to/task.md", Type: "task", Tokens: 150},
	}
	agent.Dependencies.Templates = []AssetTokenInfo{
		{Path: "/path/to/template.md", Type: "template", Tokens: 125},
	}
	agent.Dependencies.DataFiles = []AssetTokenInfo{
		{Path: "/path/to/data.md", Type: "data", Tokens: 125},
	}

	// Test JSON marshaling
	data, err := json.Marshal(agent)
	require.NoError(t, err, "Failed to marshal AgentTokenInfo")

	// Test JSON unmarshaling
	var unmarshaled AgentTokenInfo
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err, "Failed to unmarshal AgentTokenInfo")

	// Verify basic fields
	assert.Equal(t, agent.AgentName, unmarshaled.AgentName, "AgentName mismatch")
	assert.Equal(t, agent.TotalTokens, unmarshaled.TotalTokens, "TotalTokens mismatch")

	// Verify assets
	assert.Len(t, unmarshaled.Assets, len(agent.Assets), "Assets count mismatch")

	// Verify dependencies
	assert.Len(t, unmarshaled.Dependencies.Tasks, len(agent.Dependencies.Tasks), "Tasks count mismatch")
	assert.Len(t, unmarshaled.Dependencies.Templates, len(agent.Dependencies.Templates), "Templates count mismatch")
	assert.Len(t, unmarshaled.Dependencies.DataFiles, len(agent.Dependencies.DataFiles), "DataFiles count mismatch")
}

func TestProjectTokenInfo_JSON(t *testing.T) {
	project := ProjectTokenInfo{
		TotalTokens: 1000,
		Agents: []AgentTokenInfo{
			{
				AgentName:   "Agent 1",
				TotalTokens: 500,
			},
			{
				AgentName:   "Agent 2",
				TotalTokens: 500,
			},
		},
		Breakdown: TokenBreakdown{
			Agents:    200,
			Tasks:     300,
			Templates: 250,
			DataFiles: 250,
		},
	}

	// Test JSON marshaling
	data, err := json.Marshal(project)
	require.NoError(t, err, "Failed to marshal ProjectTokenInfo")

	// Test JSON unmarshaling
	var unmarshaled ProjectTokenInfo
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err, "Failed to unmarshal ProjectTokenInfo")

	// Verify fields
	assert.Equal(t, project.TotalTokens, unmarshaled.TotalTokens, "TotalTokens mismatch")
	assert.Len(t, unmarshaled.Agents, len(project.Agents), "Agents count mismatch")
	assert.Equal(t, project.Breakdown.Agents, unmarshaled.Breakdown.Agents, "Breakdown.Agents mismatch")
}

func TestTokenBreakdown_JSON(t *testing.T) {
	breakdown := TokenBreakdown{
		Agents:    100,
		Tasks:     200,
		Templates: 150,
		DataFiles: 75,
	}

	// Test JSON marshaling
	data, err := json.Marshal(breakdown)
	require.NoError(t, err, "Failed to marshal TokenBreakdown")

	// Test JSON unmarshaling
	var unmarshaled TokenBreakdown
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err, "Failed to unmarshal TokenBreakdown")

	// Verify all fields
	assert.Equal(t, breakdown.Agents, unmarshaled.Agents, "Agents mismatch")
	assert.Equal(t, breakdown.Tasks, unmarshaled.Tasks, "Tasks mismatch")
	assert.Equal(t, breakdown.Templates, unmarshaled.Templates, "Templates mismatch")
	assert.Equal(t, breakdown.DataFiles, unmarshaled.DataFiles, "DataFiles mismatch")
}
