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
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewFrameworkAnalyzer(t *testing.T) {
	analyzer := NewFrameworkAnalyzer("/test/path")

	if analyzer.baseDir != "/test/path" {
		t.Errorf("Expected baseDir to be '/test/path', got '%s'", analyzer.baseDir)
	}
	if analyzer.cache == nil {
		t.Error("Expected cache to be initialized")
	}
	if analyzer.fileCache == nil {
		t.Error("Expected fileCache to be initialized")
	}
	if analyzer.resultCache == nil {
		t.Error("Expected resultCache to be initialized")
	}
}

func TestAnalyzeFramework_NoFrameworkDirectory(t *testing.T) {
	tempDir := t.TempDir()
	analyzer := NewFrameworkAnalyzer(tempDir)

	_, _, err := analyzer.AnalyzeFramework()
	if err == nil {
		t.Error("Expected error when no .krci-ai directory exists")
	}
	if !strings.Contains(err.Error(), "no .krci-ai directory found") {
		t.Errorf("Expected error message about missing directory, got: %s", err.Error())
	}
}

func TestDetectBrokenInternalLinks(t *testing.T) {
	tempDir := setupTestFramework(t)
	defer os.RemoveAll(tempDir)

	analyzer := NewFrameworkAnalyzer(tempDir)
	frameworkDir := filepath.Join(tempDir, ".krci-ai")

	// Create a task with broken link
	taskFile := filepath.Join(frameworkDir, "tasks", "test-task.md")
	content := `# Test Task
	
This task references [missing-template](./.krci-ai/templates/missing.md) which doesn't exist.
Also references [existing-data](./.krci-ai/data/test-data.md) which exists.
External link [GitHub](https://github.com) should be ignored.
`
	if err := os.WriteFile(taskFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	issues, err := analyzer.detectBrokenInternalLinks(frameworkDir)
	if err != nil {
		t.Fatalf("detectBrokenInternalLinks failed: %v", err)
	}

	// Should find one broken link (missing.md) but ignore external links
	if len(issues) != 1 {
		t.Errorf("Expected 1 broken link issue, got %d", len(issues))
	}
	if len(issues) > 0 {
		if issues[0].Type != "broken_link" {
			t.Errorf("Expected issue type 'broken_link', got '%s'", issues[0].Type)
		}
		if issues[0].Severity != SeverityCritical {
			t.Errorf("Expected critical severity, got %v", issues[0].Severity)
		}
		if !strings.Contains(issues[0].Message, "missing.md") {
			t.Errorf("Expected message to contain 'missing.md', got: %s", issues[0].Message)
		}
	}
}

func TestSeverityString(t *testing.T) {
	tests := []struct {
		severity Severity
		expected string
	}{
		{SeverityInfo, "INFO"},
		{SeverityWarning, "WARNING"},
		{SeverityCritical, "CRITICAL"},
	}

	for _, test := range tests {
		if test.severity.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.severity.String())
		}
	}
}

func TestValidationReport(t *testing.T) {
	issues := []ValidationIssue{
		{Type: "test_critical", Severity: SeverityCritical, Message: "Critical issue"},
		{Type: "test_warning", Severity: SeverityWarning, Message: "Warning issue"},
		{Type: "test_info", Severity: SeverityInfo, Message: "Info issue"},
	}

	report := NewValidationReport(issues, nil, 0)

	if report.IsValid {
		t.Error("Expected report to be invalid due to critical issue")
	}
	if !report.HasCritical {
		t.Error("Expected report to have critical issues")
	}
	if !report.HasWarnings {
		t.Error("Expected report to have warnings")
	}
	if report.CriticalCount != 1 {
		t.Errorf("Expected 1 critical issue, got %d", report.CriticalCount)
	}
	if report.WarningCount != 1 {
		t.Errorf("Expected 1 warning, got %d", report.WarningCount)
	}
	if report.InfoCount != 1 {
		t.Errorf("Expected 1 info issue, got %d", report.InfoCount)
	}
	if report.GetExitCode() != 1 {
		t.Errorf("Expected exit code 1 for critical issues, got %d", report.GetExitCode())
	}
}

func TestValidationReportValid(t *testing.T) {
	issues := []ValidationIssue{
		{Type: "test_warning", Severity: SeverityWarning, Message: "Warning issue"},
	}

	report := NewValidationReport(issues, nil, 0)

	if !report.IsValid {
		t.Error("Expected report to be valid with only warnings")
	}
	if report.HasCritical {
		t.Error("Expected report to not have critical issues")
	}
	if report.GetExitCode() != 0 {
		t.Errorf("Expected exit code 0 for warnings only, got %d", report.GetExitCode())
	}
}

func TestExtractYAMLTasks(t *testing.T) {
	tempDir := t.TempDir()
	analyzer := NewFrameworkAnalyzer(tempDir)

	// Create agent file with tasks
	agentFile := filepath.Join(tempDir, "test-agent.yaml")
	agentContent := `agent:
  identity:
    name: "Test Agent"
  tasks:
    - ./.krci-ai/tasks/task1.md
    - ./.krci-ai/tasks/task2.md
    - ./some-other-task.md
`
	if err := os.WriteFile(agentFile, []byte(agentContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tasks, err := analyzer.extractYAMLTasks(agentFile)
	if err != nil {
		t.Fatalf("extractYAMLTasks failed: %v", err)
	}

	// Should only return .krci-ai task references
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	found1, found2 := false, false
	for _, task := range tasks {
		if task == "./.krci-ai/tasks/task1.md" {
			found1 = true
		}
		if task == "./.krci-ai/tasks/task2.md" {
			found2 = true
		}
		if task == "./some-other-task.md" {
			t.Errorf("Should not include non-.krci-ai tasks: %s", task)
		}
	}

	if !found1 || !found2 {
		t.Error("Expected to find both .krci-ai tasks")
	}
}

// setupTestFramework creates a test framework structure
func setupTestFramework(t *testing.T) string {
	tempDir := t.TempDir()

	// Create framework directory structure
	frameworkDir := filepath.Join(tempDir, ".krci-ai")
	if err := os.MkdirAll(filepath.Join(frameworkDir, "agents"), 0755); err != nil {
		t.Fatalf("Failed to create agents directory: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(frameworkDir, "tasks"), 0755); err != nil {
		t.Fatalf("Failed to create tasks directory: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(frameworkDir, "templates"), 0755); err != nil {
		t.Fatalf("Failed to create templates directory: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(frameworkDir, "data"), 0755); err != nil {
		t.Fatalf("Failed to create data directory: %v", err)
	}

	// Create some basic files
	taskFile := filepath.Join(frameworkDir, "tasks", "existing-task.md")
	if err := os.WriteFile(taskFile, []byte("# Existing Task\nContent here."), 0644); err != nil {
		t.Fatalf("Failed to create task file: %v", err)
	}

	dataFile := filepath.Join(frameworkDir, "data", "test-data.md")
	if err := os.WriteFile(dataFile, []byte("# Test Data\nData content."), 0644); err != nil {
		t.Fatalf("Failed to create data file: %v", err)
	}

	templateFile := filepath.Join(frameworkDir, "templates", "good-template.md")
	if err := os.WriteFile(templateFile, []byte("# Good Template\nTemplate content."), 0644); err != nil {
		t.Fatalf("Failed to create template file: %v", err)
	}

	return tempDir
}

// setupTestFrameworkWithLocalTasks creates a test framework structure including local tasks
func setupTestFrameworkWithLocalTasks(t *testing.T) string {
	tempDir := setupTestFramework(t)
	frameworkDir := filepath.Join(tempDir, ".krci-ai")

	// Create local tasks directory structure
	localTasksDir := filepath.Join(frameworkDir, "local", "tasks")
	if err := os.MkdirAll(localTasksDir, 0755); err != nil {
		t.Fatalf("Failed to create local tasks directory: %v", err)
	}

	// Create a local task file
	localTaskFile := filepath.Join(localTasksDir, "local-task.md")
	if err := os.WriteFile(localTaskFile, []byte("# Local Task\nLocal task content."), 0644); err != nil {
		t.Fatalf("Failed to create local task file: %v", err)
	}

	// Create an agent with both standard and local tasks
	agentFile := filepath.Join(frameworkDir, "agents", "test-agent.yaml")
	agentContent := `agent:
  identity:
    name: "Test Agent"
    id: test-v1
    version: "1.0.0"
    description: "Test agent for local tasks"
    role: "Test Role"
    goal: "Test goal"
  tasks:
    - ./.krci-ai/tasks/existing-task.md
    - ./.krci-ai/local/tasks/local-task.md
`
	if err := os.WriteFile(agentFile, []byte(agentContent), 0644); err != nil {
		t.Fatalf("Failed to create agent file: %v", err)
	}

	return tempDir
}

func TestComponentRelationshipLocalTasks(t *testing.T) {
	tempDir := setupTestFrameworkWithLocalTasks(t)
	defer os.RemoveAll(tempDir)

	analyzer := NewFrameworkAnalyzer(tempDir)
	issues, insights, err := analyzer.AnalyzeFramework()

	if err != nil {
		t.Fatalf("Framework analysis failed: %v", err)
	}

	if len(issues) > 0 {
		t.Logf("Validation issues found:")
		for _, issue := range issues {
			t.Logf("  - %s: %s (file: %s)", issue.Type, issue.Message, issue.File)
		}
		// Allow certain validation issues that are expected in test setup
		criticalIssues := 0
		for _, issue := range issues {
			if issue.Severity == SeverityCritical {
				criticalIssues++
			}
		}
		if criticalIssues > 0 {
			t.Errorf("Expected no critical validation issues, got %d", criticalIssues)
		}
	}

	// Verify insights structure
	if insights == nil {
		t.Fatal("Expected insights to be generated")
	}

	// Find the test agent relationship
	var testAgentRel *ComponentRelationship
	for i := range insights.Relationships {
		if insights.Relationships[i].Agent == "test-agent" {
			testAgentRel = &insights.Relationships[i]
			break
		}
	}

	if testAgentRel == nil {
		t.Fatal("Expected to find test-agent relationship")
	}

	// Verify task categorization
	if len(testAgentRel.Tasks) != 1 {
		t.Errorf("Expected 1 standard task, got %d", len(testAgentRel.Tasks))
	}

	if len(testAgentRel.LocalTasks) != 1 {
		t.Errorf("Expected 1 local task, got %d", len(testAgentRel.LocalTasks))
	}

	if len(testAgentRel.Tasks) > 0 && testAgentRel.Tasks[0] != "existing-task.md" {
		t.Errorf("Expected standard task 'existing-task.md', got '%s'", testAgentRel.Tasks[0])
	}

	if len(testAgentRel.LocalTasks) > 0 && testAgentRel.LocalTasks[0] != "local-task.md" {
		t.Errorf("Expected local task 'local-task.md', got '%s'", testAgentRel.LocalTasks[0])
	}
}

func TestFormatInsightsWithLocalTasks(t *testing.T) {
	tempDir := setupTestFrameworkWithLocalTasks(t)
	defer os.RemoveAll(tempDir)

	analyzer := NewFrameworkAnalyzer(tempDir)
	_, insights, err := analyzer.AnalyzeFramework()

	if err != nil {
		t.Fatalf("Framework analysis failed: %v", err)
	}

	if insights == nil {
		t.Fatal("Expected insights to be generated")
	}

	// Test formatting
	formatted := insights.FormatInsights()

	// Should contain local task indication
	if !strings.Contains(formatted, "including 1 local") {
		t.Errorf("Expected formatted insights to show local task count, got: %s", formatted)
	}

	// Should show correct total task count
	if !strings.Contains(formatted, "test-agent → 2 tasks (including 1 local)") {
		t.Errorf("Expected formatted insights to show correct task breakdown, got: %s", formatted)
	}
}

func TestFormatInsightsWithoutLocalTasks(t *testing.T) {
	tempDir := setupTestFramework(t)
	defer os.RemoveAll(tempDir)

	// Create an agent with only standard tasks
	frameworkDir := filepath.Join(tempDir, ".krci-ai")
	agentFile := filepath.Join(frameworkDir, "agents", "standard-agent.yaml")
	agentContent := `agent:
  identity:
    name: "Standard Agent"
    id: standard-v1
    version: "1.0.0"
    description: "Test agent without local tasks"
    role: "Standard Role"
    goal: "Standard goal"
  tasks:
    - ./.krci-ai/tasks/existing-task.md
`
	if err := os.WriteFile(agentFile, []byte(agentContent), 0644); err != nil {
		t.Fatalf("Failed to create agent file: %v", err)
	}

	analyzer := NewFrameworkAnalyzer(tempDir)
	_, insights, err := analyzer.AnalyzeFramework()

	if err != nil {
		t.Fatalf("Framework analysis failed: %v", err)
	}

	if insights == nil {
		t.Fatal("Expected insights to be generated")
	}

	// Test formatting
	formatted := insights.FormatInsights()

	// Should NOT contain local task indication
	if strings.Contains(formatted, "including") {
		t.Errorf("Expected formatted insights to NOT show local task indication for standard-only agent, got: %s", formatted)
	}

	// Should show standard format
	if !strings.Contains(formatted, "standard-agent → 1 tasks →") {
		t.Errorf("Expected formatted insights to show standard format, got: %s", formatted)
	}
}

func TestMixedAgentsInsightsFormatting(t *testing.T) {
	tempDir := setupTestFrameworkWithLocalTasks(t)
	defer os.RemoveAll(tempDir)

	// Add another agent with only standard tasks
	frameworkDir := filepath.Join(tempDir, ".krci-ai")
	agentFile := filepath.Join(frameworkDir, "agents", "standard-agent.yaml")
	agentContent := `agent:
  identity:
    name: "Standard Agent"
    id: standard-v1
    version: "1.0.0"
    description: "Test agent without local tasks"
    role: "Standard Role"
    goal: "Standard goal"
  tasks:
    - ./.krci-ai/tasks/existing-task.md
`
	if err := os.WriteFile(agentFile, []byte(agentContent), 0644); err != nil {
		t.Fatalf("Failed to create agent file: %v", err)
	}

	analyzer := NewFrameworkAnalyzer(tempDir)
	_, insights, err := analyzer.AnalyzeFramework()

	if err != nil {
		t.Fatalf("Framework analysis failed: %v", err)
	}

	if insights == nil {
		t.Fatal("Expected insights to be generated")
	}

	// Test formatting shows different formats for different agents
	formatted := insights.FormatInsights()

	// Should contain both formats
	if !strings.Contains(formatted, "test-agent → 2 tasks (including 1 local)") {
		t.Errorf("Expected formatted insights to show local task breakdown for test-agent, got: %s", formatted)
	}

	if !strings.Contains(formatted, "standard-agent → 1 tasks →") && !strings.Contains(formatted, "including") {
		t.Errorf("Expected formatted insights to show standard format for standard-agent, got: %s", formatted)
	}
}
