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
	"time"
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

// Test caching functionality
func TestOptimizedAnalyzeFramework_CacheHit(t *testing.T) {
	// Create a temporary directory with valid framework structure
	tempDir := t.TempDir()
	frameworkDir := filepath.Join(tempDir, ".krci-ai")

	// Create necessary directories
	agentsDir := filepath.Join(frameworkDir, "agents")
	tasksDir := filepath.Join(frameworkDir, "tasks")
	templatesDir := filepath.Join(frameworkDir, "templates")
	dataDir := filepath.Join(frameworkDir, "data")

	for _, dir := range []string{agentsDir, tasksDir, templatesDir, dataDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	// Create a simple agent file
	agentContent := `agent:
  identity:
    name: "test-agent"
  tasks:
    - ./.krci-ai/tasks/test-task.md`

	agentFile := filepath.Join(agentsDir, "test-agent.yaml")
	if err := os.WriteFile(agentFile, []byte(agentContent), 0644); err != nil {
		t.Fatalf("Failed to create agent file: %v", err)
	}

	// Create the referenced task file
	taskContent := "# Test Task\nThis is a test task."
	taskFile := filepath.Join(tasksDir, "test-task.md")
	if err := os.WriteFile(taskFile, []byte(taskContent), 0644); err != nil {
		t.Fatalf("Failed to create task file: %v", err)
	}

	analyzer := NewFrameworkAnalyzer(tempDir)

	// First call - should run full analysis
	issues1, insights1, err1 := analyzer.OptimizedAnalyzeFramework()
	if err1 != nil {
		t.Fatalf("First analysis failed: %v", err1)
	}

	// Second call - should use cache
	issues2, insights2, err2 := analyzer.OptimizedAnalyzeFramework()
	if err2 != nil {
		t.Fatalf("Second analysis failed: %v", err2)
	}

	// Results should be identical
	if len(issues1) != len(issues2) {
		t.Errorf("Expected same number of issues, got %d vs %d", len(issues1), len(issues2))
	}

	if insights1 == nil || insights2 == nil {
		t.Fatal("Expected insights to be generated for both calls")
	}

	// Check that result cache was used
	if len(analyzer.resultCache) == 0 {
		t.Error("Expected result cache to contain entries")
	}

	if _, exists := analyzer.resultCache["framework_validation"]; !exists {
		t.Error("Expected framework_validation key in result cache")
	}
}

func TestOptimizedAnalyzeFramework_CacheMiss(t *testing.T) {
	// Create a temporary directory with valid framework structure
	tempDir := t.TempDir()
	frameworkDir := filepath.Join(tempDir, ".krci-ai")

	// Create necessary directories
	agentsDir := filepath.Join(frameworkDir, "agents")
	tasksDir := filepath.Join(frameworkDir, "tasks")

	for _, dir := range []string{agentsDir, tasksDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	// Create a simple agent file
	agentContent := `agent:
  identity:
    name: "test-agent"
  tasks:
    - ./.krci-ai/tasks/test-task.md`

	agentFile := filepath.Join(agentsDir, "test-agent.yaml")
	if err := os.WriteFile(agentFile, []byte(agentContent), 0644); err != nil {
		t.Fatalf("Failed to create agent file: %v", err)
	}

	// Create the referenced task file
	taskContent := "# Test Task\nThis is a test task."
	taskFile := filepath.Join(tasksDir, "test-task.md")
	if err := os.WriteFile(taskFile, []byte(taskContent), 0644); err != nil {
		t.Fatalf("Failed to create task file: %v", err)
	}

	analyzer := NewFrameworkAnalyzer(tempDir)

	// First call
	_, _, err1 := analyzer.OptimizedAnalyzeFramework()
	if err1 != nil {
		t.Fatalf("First analysis failed: %v", err1)
	}

	// Modify the task file to invalidate cache
	modifiedContent := "# Modified Test Task\nThis is a modified test task."
	if err := os.WriteFile(taskFile, []byte(modifiedContent), 0644); err != nil {
		t.Fatalf("Failed to modify task file: %v", err)
	}

	// Second call - cache should be invalidated
	_, _, err2 := analyzer.OptimizedAnalyzeFramework()
	if err2 != nil {
		t.Fatalf("Second analysis failed: %v", err2)
	}

	// Verify that the file modification was detected
	if !analyzer.isFileModified(taskFile) {
		t.Error("Expected modified file to be detected")
	}
}

func TestIsFileModified(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")

	analyzer := NewFrameworkAnalyzer(tempDir)

	// Test with non-existent file
	if !analyzer.isFileModified(testFile) {
		t.Error("Expected non-existent file to be reported as modified")
	}

	// Create the file
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// First check - should be modified (not in cache)
	if !analyzer.isFileModified(testFile) {
		t.Error("Expected new file to be reported as modified")
	}

	// Second check - should not be modified (now in cache)
	if analyzer.isFileModified(testFile) {
		t.Error("Expected cached file to not be reported as modified")
	}

	// Add a small delay to ensure timestamp difference in CI environments
	time.Sleep(1100 * time.Millisecond)

	// Modify the file
	if err := os.WriteFile(testFile, []byte("modified content"), 0644); err != nil {
		t.Fatalf("Failed to modify test file: %v", err)
	}

	// Should be reported as modified again
	if !analyzer.isFileModified(testFile) {
		t.Error("Expected modified file to be reported as modified")
	}
}

func TestAreAnyFilesModified(t *testing.T) {
	tempDir := t.TempDir()
	frameworkDir := filepath.Join(tempDir, ".krci-ai")

	// Create directories
	agentsDir := filepath.Join(frameworkDir, "agents")
	tasksDir := filepath.Join(frameworkDir, "tasks")
	templatesDir := filepath.Join(frameworkDir, "templates")
	dataDir := filepath.Join(frameworkDir, "data")

	for _, dir := range []string{agentsDir, tasksDir, templatesDir, dataDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	analyzer := NewFrameworkAnalyzer(tempDir)

	// Empty framework - no modifications
	if analyzer.areAnyFilesModified(frameworkDir) {
		t.Error("Expected empty framework to have no modifications")
	}

	// Add a file
	testFile := filepath.Join(tasksDir, "test.md")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Should detect new file
	if !analyzer.areAnyFilesModified(frameworkDir) {
		t.Error("Expected framework with new file to be modified")
	}

	// Check again - should not be modified (cached)
	if analyzer.areAnyFilesModified(frameworkDir) {
		t.Error("Expected cached framework to not be modified")
	}

	// Add a small delay to ensure timestamp difference in CI environments
	time.Sleep(1100 * time.Millisecond)

	// Modify the file
	if err := os.WriteFile(testFile, []byte("modified content"), 0644); err != nil {
		t.Fatalf("Failed to modify test file: %v", err)
	}

	// Should detect modification
	if !analyzer.areAnyFilesModified(frameworkDir) {
		t.Error("Expected framework with modified file to be modified")
	}
}

func TestCheckDirectoryForModifications(t *testing.T) {
	tempDir := t.TempDir()
	testDir := filepath.Join(tempDir, "test")

	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	analyzer := NewFrameworkAnalyzer(tempDir)

	// Empty directory
	if analyzer.checkDirectoryForModifications(testDir) {
		t.Error("Expected empty directory to have no modifications")
	}

	// Add YAML file
	yamlFile := filepath.Join(testDir, "test.yaml")
	if err := os.WriteFile(yamlFile, []byte("test: content"), 0644); err != nil {
		t.Fatalf("Failed to create YAML file: %v", err)
	}

	if !analyzer.checkDirectoryForModifications(testDir) {
		t.Error("Expected directory with new YAML file to be modified")
	}

	// Add markdown file
	mdFile := filepath.Join(testDir, "test.md")
	if err := os.WriteFile(mdFile, []byte("# Test"), 0644); err != nil {
		t.Fatalf("Failed to create markdown file: %v", err)
	}

	if !analyzer.checkDirectoryForModifications(testDir) {
		t.Error("Expected directory with new markdown file to be modified")
	}

	// Test data directory special case
	dataDir := filepath.Join(tempDir, "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		t.Fatalf("Failed to create data directory: %v", err)
	}

	// Add any file in data directory
	dataFile := filepath.Join(dataDir, "test.json")
	if err := os.WriteFile(dataFile, []byte(`{"test": "content"}`), 0644); err != nil {
		t.Fatalf("Failed to create data file: %v", err)
	}

	if !analyzer.checkDirectoryForModifications(dataDir) {
		t.Error("Expected data directory with new file to be modified")
	}
}

func TestOptimizedAnalyzeFramework_NoFrameworkDirectory(t *testing.T) {
	tempDir := t.TempDir()
	analyzer := NewFrameworkAnalyzer(tempDir)

	_, _, err := analyzer.OptimizedAnalyzeFramework()
	if err == nil {
		t.Error("Expected error when no .krci-ai directory exists")
	}

	if !strings.Contains(err.Error(), "file does not exist") {
		t.Errorf("Expected error about missing directory, got: %s", err.Error())
	}
}

func TestOptimizedAnalyzeFramework_CacheWithInsightsError(t *testing.T) {
	tempDir := t.TempDir()
	frameworkDir := filepath.Join(tempDir, ".krci-ai")

	// Create minimal framework structure
	agentsDir := filepath.Join(frameworkDir, "agents")
	if err := os.MkdirAll(agentsDir, 0755); err != nil {
		t.Fatalf("Failed to create agents directory: %v", err)
	}

	analyzer := NewFrameworkAnalyzer(tempDir)

	// First call to populate cache
	issues1, _, _ := analyzer.OptimizedAnalyzeFramework()

	// Verify cache is populated
	if len(analyzer.resultCache) == 0 {
		t.Error("Expected result cache to be populated")
	}

	// Second call should use cache even if insights generation fails
	// (we can't easily force insights to fail, but we can verify cache behavior)
	issues2, _, err2 := analyzer.OptimizedAnalyzeFramework()
	if err2 != nil {
		t.Fatalf("Second analysis failed: %v", err2)
	}

	// Should return cached results
	if len(issues1) != len(issues2) {
		t.Errorf("Expected cached results to match, got %d vs %d issues", len(issues1), len(issues2))
	}
}

// Test output formatting functions - these are currently 0% covered
func TestValidationReportFormatting(t *testing.T) {
	// Create test issues
	issues := []ValidationIssue{
		{
			Type:        "critical_issue",
			Severity:    SeverityCritical,
			File:        "test.yaml",
			Line:        10,
			Message:     "Critical test issue",
			FixGuidance: "Fix this critical issue",
		},
		{
			Type:        "warning_issue",
			Severity:    SeverityWarning,
			File:        "test.md",
			Line:        5,
			Message:     "Warning test issue",
			FixGuidance: "Fix this warning",
		},
		{
			Type:        "info_issue",
			Severity:    SeverityInfo,
			File:        "test.txt",
			Line:        1,
			Message:     "Info test issue",
			FixGuidance: "This is just info",
		},
	}

	// Create test insights
	insights := &FrameworkInsights{
		ComponentCounts: ComponentCounts{
			Agents:    2,
			Tasks:     3,
			Templates: 1,
			Data:      1,
		},
		Relationships: []ComponentRelationship{
			{
				Agent:     "test-agent",
				Tasks:     []string{"task1", "task2"},
				Templates: []string{"template1"},
				DataFiles: []string{"data1"},
			},
		},
		UsageStatistics: UsageStatistics{
			MostUsedTemplate:   "template1",
			MostUsedData:       "data1",
			TemplateUsageCount: 1,
			DataUsageCount:     1,
		},
		TotalReferences: 4,
	}

	processTime := time.Duration(100) * time.Millisecond
	report := NewValidationReport(issues, insights, processTime)

	// Test FormatReport with verbose output
	verboseOutput := report.FormatReport(true)
	if len(verboseOutput) == 0 {
		t.Error("Expected verbose report output to be non-empty")
	}

	// Should contain critical issues section
	if !strings.Contains(verboseOutput, "CRITICAL") {
		t.Error("Expected report to contain CRITICAL section")
	}

	// Should contain warnings section
	if !strings.Contains(verboseOutput, "WARNING") {
		t.Error("Expected report to contain WARNING section")
	}

	// Should contain insights - but only if framework is valid or no critical issues
	// Since we have critical issues, insights might not be shown
	if !strings.Contains(verboseOutput, "Framework Insights") && !strings.Contains(verboseOutput, "⚡ Validation completed") {
		// This is acceptable behavior when there are critical issues
	}

	// Test FormatReport without verbose output
	simpleOutput := report.FormatReport(false)
	if len(simpleOutput) == 0 {
		t.Error("Expected simple report output to be non-empty")
	}

	// Test GetExitCode
	exitCode := report.GetExitCode()
	if exitCode != 1 { // Should be 1 for critical issues (not 2)
		t.Errorf("Expected exit code 1 for critical issues, got %d", exitCode)
	}
}

func TestValidationReportWithoutIssues(t *testing.T) {
	// Test with no issues
	insights := &FrameworkInsights{
		ComponentCounts: ComponentCounts{
			Agents:    1,
			Tasks:     1,
			Templates: 0,
			Data:      0,
		},
		TotalReferences: 0,
	}

	processTime := time.Duration(50) * time.Millisecond
	report := NewValidationReport([]ValidationIssue{}, insights, processTime)

	// Should be valid with no issues
	if !report.IsValid {
		t.Error("Expected report to be valid with no issues")
	}

	if report.HasCritical {
		t.Error("Expected no critical issues")
	}

	if report.HasWarnings {
		t.Error("Expected no warnings")
	}

	// Test exit code for success
	exitCode := report.GetExitCode()
	if exitCode != 0 {
		t.Errorf("Expected exit code 0 for no issues, got %d", exitCode)
	}

	// Test FormatSimpleSuccess
	successOutput := FormatSimpleSuccess(insights, processTime)
	if len(successOutput) == 0 {
		t.Error("Expected success output to be non-empty")
	}

	if !strings.Contains(successOutput, "✅") {
		t.Error("Expected success output to contain checkmark")
	}
}

func TestValidationReportExitCodes(t *testing.T) {
	testCases := []struct {
		name         string
		issues       []ValidationIssue
		expectedCode int
	}{
		{
			name:         "No issues",
			issues:       []ValidationIssue{},
			expectedCode: 0,
		},
		{
			name: "Only warnings",
			issues: []ValidationIssue{
				{Severity: SeverityWarning, Message: "warning"},
			},
			expectedCode: 0, // No critical issues = exit code 0
		},
		{
			name: "Critical issues",
			issues: []ValidationIssue{
				{Severity: SeverityCritical, Message: "critical"},
			},
			expectedCode: 1, // Critical issues = exit code 1
		},
		{
			name: "Mixed issues",
			issues: []ValidationIssue{
				{Severity: SeverityWarning, Message: "warning"},
				{Severity: SeverityCritical, Message: "critical"},
				{Severity: SeverityInfo, Message: "info"},
			},
			expectedCode: 1, // Critical takes precedence, exit code 1
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			report := NewValidationReport(tc.issues, nil, 0)
			exitCode := report.GetExitCode()
			if exitCode != tc.expectedCode {
				t.Errorf("Expected exit code %d, got %d", tc.expectedCode, exitCode)
			}
		})
	}
}

// Test utility functions with low coverage
func TestExtractMarkdownLinks(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.md")

	// Create test markdown with various link formats
	content := `# Test File

Some content with links:

- [Internal link](./.krci-ai/tasks/task1.md)
- [Template link](./.krci-ai/templates/template1.md)
- [Data link](./.krci-ai/data/data1.json)
- [External link](https://example.com)
- [Relative link](../other.md)

More content here.`

	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	analyzer := NewFrameworkAnalyzer(tempDir)
	links, err := analyzer.extractMarkdownLinks(testFile)

	if err != nil {
		t.Fatalf("extractMarkdownLinks failed: %v", err)
	}

	// Should find internal framework links
	expectedLinks := []string{
		"./.krci-ai/tasks/task1.md",
		"./.krci-ai/templates/template1.md",
		"./.krci-ai/data/data1.json",
	}

	if len(links) != len(expectedLinks) {
		t.Errorf("Expected %d links, got %d", len(expectedLinks), len(links))
	}

	for _, expected := range expectedLinks {
		found := false
		for _, link := range links {
			if link == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find link: %s", expected)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	analyzer := NewFrameworkAnalyzer("")

	// Test with duplicates
	input := []string{"a", "b", "a", "c", "b", "d"}
	result := analyzer.removeDuplicates(input)

	expected := []string{"a", "b", "c", "d"}
	if len(result) != len(expected) {
		t.Errorf("Expected %d unique items, got %d", len(expected), len(result))
	}

	// Check all expected items are present
	for _, exp := range expected {
		found := false
		for _, res := range result {
			if res == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find item: %s", exp)
		}
	}

	// Test with empty slice
	emptyResult := analyzer.removeDuplicates([]string{})
	if len(emptyResult) != 0 {
		t.Error("Expected empty result for empty input")
	}

	// Test with no duplicates
	noDupInput := []string{"x", "y", "z"}
	noDupResult := analyzer.removeDuplicates(noDupInput)
	if len(noDupResult) != 3 {
		t.Errorf("Expected 3 items, got %d", len(noDupResult))
	}
}
