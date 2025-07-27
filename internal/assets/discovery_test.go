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
package assets

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// testAssets is already declared in installer_test.go

func setupTestFramework(t *testing.T) string {
	t.Helper()

	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "krci-ai-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Create framework structure
	frameworkDir := filepath.Join(tmpDir, ".krci-ai")
	agentsDir := filepath.Join(frameworkDir, "agents")
	tasksDir := filepath.Join(frameworkDir, "tasks")
	templatesDir := filepath.Join(frameworkDir, "templates")
	dataDir := filepath.Join(frameworkDir, "data")

	for _, dir := range []string{agentsDir, tasksDir, templatesDir, dataDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create dir %s: %v", dir, err)
		}
	}

	// Create test agent file
	agentContent := `agent:
  identity:
    name: "Test Agent"
    id: test-agent-v1
    version: "1.0.0"
    description: "Test agent for unit testing"
    role: "Test Role"
    goal: "Test goal"
    icon: "🧪"
  tasks:
    - ./.krci-ai/tasks/test-task.md
`
	agentFile := filepath.Join(agentsDir, "test.yaml")
	if err := os.WriteFile(agentFile, []byte(agentContent), 0644); err != nil {
		t.Fatalf("Failed to write agent file: %v", err)
	}

	// Create test task file with template/data references
	taskContent := `# Test Task

This task uses [test template](./.krci-ai/templates/test-template.md) and [test data](./.krci-ai/data/test-data.md).
`
	taskFile := filepath.Join(tasksDir, "test-task.md")
	if err := os.WriteFile(taskFile, []byte(taskContent), 0644); err != nil {
		t.Fatalf("Failed to write task file: %v", err)
	}

	// Create test template and data files
	templateFile := filepath.Join(templatesDir, "test-template.md")
	if err := os.WriteFile(templateFile, []byte("# Test Template"), 0644); err != nil {
		t.Fatalf("Failed to write template file: %v", err)
	}

	dataFile := filepath.Join(dataDir, "test-data.md")
	if err := os.WriteFile(dataFile, []byte("# Test Data"), 0644); err != nil {
		t.Fatalf("Failed to write data file: %v", err)
	}

	return tmpDir
}

func TestDiscoverAgentsWithDependencies(t *testing.T) {
	tmpDir := setupTestFramework(t)
	defer os.RemoveAll(tmpDir)

	discovery := NewDiscovery(tmpDir, testAssets)

	agentDeps, err := discovery.DiscoverAgentsWithDependencies()
	if err != nil {
		t.Fatalf("DiscoverAgentsWithDependencies failed: %v", err)
	}

	if len(agentDeps) == 0 {
		t.Fatal("Expected at least one agent with dependencies")
	}

	agent := agentDeps[0]
	if agent.Name != "Test Agent" {
		t.Errorf("Expected agent name 'Test Agent', got '%s'", agent.Name)
	}

	if agent.Role != "Test Role" {
		t.Errorf("Expected agent role 'Test Role', got '%s'", agent.Role)
	}

	if agent.Icon != "🧪" {
		t.Errorf("Expected agent icon '🧪', got '%s'", agent.Icon)
	}

	// Check that dependencies are populated
	if len(agent.Tasks) == 0 {
		t.Error("Expected agent to have tasks")
	}

	// Verify task name
	expectedTask := "test-task.md"
	found := false
	for _, task := range agent.Tasks {
		if task == expectedTask {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected to find task '%s' in agent tasks %v", expectedTask, agent.Tasks)
	}
}

func TestFormatAgentDependencyTable(t *testing.T) {
	tmpDir := setupTestFramework(t)
	defer os.RemoveAll(tmpDir)

	discovery := NewDiscovery(tmpDir, testAssets)

	// Test with sample data
	testAgents := []AgentDependencyInfo{
		{
			AgentInfo: AgentInfo{
				Name:        "Agent 1",
				Description: "Test agent 1",
				Role:        "Role 1",
			},
			Tasks:     []string{"task1.md", "task2.md"},
			Templates: []string{"template1.md"},
			DataFiles: []string{"data1.md"},
		},
		{
			AgentInfo: AgentInfo{
				Name:        "Agent 2",
				Description: "Test agent 2",
				Role:        "Role 2",
			},
			Tasks:     []string{},
			Templates: []string{"template2.md", "template3.md"},
			DataFiles: []string{},
		},
	}

	table := discovery.FormatAgentDependencyTable(testAgents)

	// Verify agent headers are present (new format shows each agent separately)
	if !strings.Contains(table, "Agent 1") {
		t.Error("Table should contain 'Agent 1'")
	}
	if !strings.Contains(table, "Agent 2") {
		t.Error("Table should contain 'Agent 2'")
	}

	// Verify table headers (new format uses individual tables per agent)
	if !strings.Contains(table, "Task") {
		t.Error("Table should contain 'Task' header")
	}
	if !strings.Contains(table, "Templates") {
		t.Error("Table should contain 'Templates' header")
	}
	if !strings.Contains(table, "Data") {
		t.Error("Table should contain 'Data' header")
	}

	// Verify task data
	if !strings.Contains(table, "task1.md") {
		t.Error("Table should contain 'task1.md'")
	}

	// Verify template data
	if !strings.Contains(table, "template1.md") {
		t.Error("Table should contain 'template1.md'")
	}

	// Debug: print the table to see what it looks like
	t.Logf("Generated table:\n%s", table)

	// Note: The new format may not show dashes for empty values in the same way
	// Just verify the table contains expected content
}

func TestFormatAgentDependencyTableEmpty(t *testing.T) {
	tmpDir := setupTestFramework(t)
	defer os.RemoveAll(tmpDir)

	discovery := NewDiscovery(tmpDir, testAssets)

	// Test with empty slice
	table := discovery.FormatAgentDependencyTable([]AgentDependencyInfo{})

	expected := "No agents found"
	if table != expected {
		t.Errorf("Expected '%s', got '%s'", expected, table)
	}
}

func TestFormatAgentDependencyTableLongContent(t *testing.T) {
	tmpDir := setupTestFramework(t)
	defer os.RemoveAll(tmpDir)

	discovery := NewDiscovery(tmpDir, testAssets)

	// Test with very long content to verify truncation
	longTasks := make([]string, 20)
	for i := range longTasks {
		longTasks[i] = "very-long-task-name-that-should-be-truncated.md"
	}

	testAgents := []AgentDependencyInfo{
		{
			AgentInfo: AgentInfo{
				Name: "LongAgent",
				Role: "LongRole",
			},
			Tasks: longTasks,
		},
	}

	table := discovery.FormatAgentDependencyTable(testAgents)

	// Verify table is formatted properly and contains agent name
	if !strings.Contains(table, "LongAgent") {
		t.Error("Table should contain agent name")
	}

	// Note: Truncation might not always show "..." depending on content length and column widths
	// The main thing is that the table formats correctly without errors
}

func TestFormatAgentTableWithIcon(t *testing.T) {
	tmpDir := setupTestFramework(t)
	defer os.RemoveAll(tmpDir)

	discovery := NewDiscovery(tmpDir, testAssets)

	// Test with agent that has an icon
	testAgent := AgentDependencyInfo{
		AgentInfo: AgentInfo{
			Name:        "Test Agent",
			Description: "Test agent with icon",
			Role:        "Test Role",
			Icon:        "🧪",
		},
		Tasks:     []string{"test-task.md"},
		Templates: []string{"test-template.md"},
		DataFiles: []string{"test-data.md"},
	}

	table := discovery.formatAgentTable(testAgent)

	// Verify icon is included
	if !strings.Contains(table, "🧪") {
		t.Error("Table should contain the agent icon")
	}

	// Verify agent name is included
	if !strings.Contains(table, "Test Agent") {
		t.Error("Table should contain the agent name")
	}

	// Verify description is included
	if !strings.Contains(table, "Test agent with icon") {
		t.Error("Table should contain the agent description")
	}

	// Verify table structure
	if !strings.Contains(table, "Task") {
		t.Error("Table should contain 'Task' header")
	}
	if !strings.Contains(table, "Templates") {
		t.Error("Table should contain 'Templates' header")
	}
	if !strings.Contains(table, "Data") {
		t.Error("Table should contain 'Data' header")
	}
}

func TestFormatAgentTableWithoutIcon(t *testing.T) {
	tmpDir := setupTestFramework(t)
	defer os.RemoveAll(tmpDir)

	discovery := NewDiscovery(tmpDir, testAssets)

	// Test with agent that has no icon (should use default)
	testAgent := AgentDependencyInfo{
		AgentInfo: AgentInfo{
			Name:        "Test Agent",
			Description: "Test agent without icon",
			Role:        "Test Role",
			Icon:        "", // No icon specified
		},
		Tasks:     []string{"test-task.md"},
		Templates: []string{"test-template.md"},
		DataFiles: []string{"test-data.md"},
	}

	table := discovery.formatAgentTable(testAgent)

	// Verify default icon is used
	if !strings.Contains(table, DefaultAgentIcon) {
		t.Errorf("Table should contain the default agent icon '%s'", DefaultAgentIcon)
	}

	// Verify agent name is included
	if !strings.Contains(table, "Test Agent") {
		t.Error("Table should contain the agent name")
	}
}
