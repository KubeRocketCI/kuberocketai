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
