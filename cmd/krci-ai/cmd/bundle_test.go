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
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBundleCommandExists(t *testing.T) {
	if bundleCmd == nil {
		t.Fatal("bundleCmd is nil")
	}

	if bundleCmd.Use != "bundle" {
		t.Errorf("Expected command name 'bundle', got '%s'", bundleCmd.Use)
	}

	if bundleCmd.Short == "" {
		t.Error("Command short description is empty")
	}

	if bundleCmd.Long == "" {
		t.Error("Command long description is empty")
	}

	if bundleCmd.RunE == nil {
		t.Error("Command RunE function is nil")
	}
}

func TestBundleCommandRegistered(t *testing.T) {
	// Test that bundle command is registered
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "bundle" {
			found = true
			break
		}
	}
	if !found {
		t.Error("bundle command should be registered with root command")
	}
}

func TestBundleCommandFlags(t *testing.T) {
	// Test --all flag
	allFlag := bundleCmd.Flags().Lookup("all")
	if allFlag == nil {
		t.Error("--all flag should be defined")
	} else if allFlag.Value.Type() != "bool" {
		t.Errorf("--all should be a boolean flag, got %s", allFlag.Value.Type())
	}

	// Test --dry-run flag
	dryRunFlag := bundleCmd.Flags().Lookup("dry-run")
	if dryRunFlag == nil {
		t.Error("--dry-run flag should be defined")
	} else if dryRunFlag.Value.Type() != "bool" {
		t.Errorf("--dry-run should be a boolean flag, got %s", dryRunFlag.Value.Type())
	}

	// Test --output flag
	outputFlag := bundleCmd.Flags().Lookup("output")
	if outputFlag == nil {
		t.Error("--output flag should be defined")
	} else if outputFlag.Value.Type() != "string" {
		t.Errorf("--output should be a string flag, got %s", outputFlag.Value.Type())
	}
}

func TestGenerateBundleFilename(t *testing.T) {
	tests := []struct {
		name           string
		customOutput   string
		selectedAgents []string
		expected       string
	}{
		{
			name:           "default filename with no agents",
			customOutput:   "",
			selectedAgents: nil,
			expected:       "all.md",
		},
		{
			name:           "custom filename with extension",
			customOutput:   "my-bundle.md",
			selectedAgents: []string{"pm", "architect"},
			expected:       "my-bundle.md",
		},
		{
			name:           "custom filename without extension",
			customOutput:   "my-bundle",
			selectedAgents: []string{"pm"},
			expected:       "my-bundle.md",
		},
		{
			name:           "single agent filename",
			customOutput:   "",
			selectedAgents: []string{"pm"},
			expected:       "pm.md",
		},
		{
			name:           "multiple agents alphabetically sorted",
			customOutput:   "",
			selectedAgents: []string{"pm", "architect", "dev"},
			expected:       "architect-dev-pm.md",
		},
		{
			name:           "agents case insensitive sorting",
			customOutput:   "",
			selectedAgents: []string{"PM", "Architect"},
			expected:       "architect-pm.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateBundleFilename(tt.customOutput, tt.selectedAgents)
			if result != tt.expected {
				t.Errorf("generateBundleFilename(%q, %v) = %q, want %q", tt.customOutput, tt.selectedAgents, result, tt.expected)
			}
		})
	}
}

func TestParseAgentList(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
		{
			name:     "single agent",
			input:    "pm",
			expected: []string{"pm"},
		},
		{
			name:     "comma separated",
			input:    "pm,architect",
			expected: []string{"pm", "architect"},
		},
		{
			name:     "comma separated with spaces",
			input:    "pm, architect, dev",
			expected: []string{"pm", "architect", "dev"},
		},
		{
			name:     "space separated",
			input:    "pm architect dev",
			expected: []string{"pm", "architect", "dev"},
		},
		{
			name:     "mixed spaces and extra whitespace",
			input:    "  pm   architect  ",
			expected: []string{"pm", "architect"},
		},
		{
			name:     "comma with empty parts",
			input:    "pm,,architect,",
			expected: []string{"pm", "architect"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseAgentList(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("ParseAgentList(%q) returned %d items, want %d", tt.input, len(result), len(tt.expected))
				return
			}
			for i, expected := range tt.expected {
				if result[i] != expected {
					t.Errorf("ParseAgentList(%q)[%d] = %q, want %q", tt.input, i, result[i], expected)
				}
			}
		})
	}
}

func TestOptimizeContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no change needed",
			input:    "line1\nline2\nline3",
			expected: "line1\nline2\nline3",
		},
		{
			name:     "remove excessive blank lines",
			input:    "line1\n\n\n\nline2\n\n\n\nline3",
			expected: "line1\n\nline2\n\nline3",
		},
		{
			name:     "preserve single blank lines",
			input:    "line1\n\nline2\n\nline3",
			expected: "line1\n\nline2\n\nline3",
		},
		{
			name:     "handle trailing blank lines",
			input:    "line1\nline2\n\n\n\n",
			expected: "line1\nline2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := optimizeContent(tt.input)
			if result != tt.expected {
				t.Errorf("optimizeContent() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestGetAgentTasks(t *testing.T) {
	content := &BundleContent{
		Agents: []AgentBundleInfo{
			{
				FilePath: ".krci-ai/agents/test.yaml",
				Content: `agent:
  identity:
    name: "Test Agent"
  tasks:
    - ./.krci-ai/tasks/task1.md
    - ./.krci-ai/tasks/task2.md
  commands:
    help: "Show help"`,
			},
		},
	}

	tasks := getAgentTasks(".krci-ai/agents/test.yaml", content)
	expected := []string{"tasks/task1.md", "tasks/task2.md"}

	if len(tasks) != len(expected) {
		t.Errorf("getAgentTasks() returned %d tasks, want %d", len(tasks), len(expected))
		return
	}

	for i, task := range tasks {
		if task != expected[i] {
			t.Errorf("getAgentTasks()[%d] = %q, want %q", i, task, expected[i])
		}
	}
}

func TestGetAgentTasksNoTasks(t *testing.T) {
	content := &BundleContent{
		Agents: []AgentBundleInfo{
			{
				FilePath: ".krci-ai/agents/test.yaml",
				Content: `agent:
  identity:
    name: "Test Agent"
  commands:
    help: "Show help"`,
			},
		},
	}

	tasks := getAgentTasks(".krci-ai/agents/test.yaml", content)
	if len(tasks) != 0 {
		t.Errorf("getAgentTasks() with no tasks should return empty slice, got %v", tasks)
	}
}

func TestGenerateBundleMarkdown(t *testing.T) {
	content := &BundleContent{
		Agents: []AgentBundleInfo{
			{
				FilePath: ".krci-ai/agents/test.yaml",
				Name:     "Test Agent",
				Role:     "Tester",
				Content:  "agent: test content",
			},
		},
		Tasks: map[string]string{
			"tasks/test-task.md": "# Test Task\nTest task content",
		},
		Templates: map[string]string{
			"templates/test-template.md": "# Test Template\nTest template content",
		},
		DataFiles: map[string]string{
			"data/test-data.md": "# Test Data\nTest data content",
		},
	}

	result := generateBundleMarkdown(content)

	// Test bundle structure
	expectedStrings := []string{
		"# KubeRocketAI Framework Bundle",
		"==== FILE: .krci-ai/agents/test.yaml ====",
		"==== END FILE ====",
		"==== FILE: templates/test-template.md ====",
		"==== END FILE ====",
		"==== FILE: data/test-data.md ====",
		"==== END FILE ====",
		"agent: test content",
		"Test template content",
		"Test data content",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(result, expected) {
			t.Errorf("generateBundleMarkdown() should contain %q", expected)
		}
	}
}

func TestBundleCommandRequiresAllOrAgentFlag(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "bundle-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current dir: %v", err)
	}
	defer func() {
		if chErr := os.Chdir(originalDir); chErr != nil {
			t.Logf("Warning: failed to change back to original directory: %v", chErr)
		}
	}()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp dir: %v", err)
	}

	// Reset flags
	bundleAll = false
	bundleDryRun = false
	bundleOutput = ""

	err = runBundle(bundleCmd, []string{})
	if err == nil {
		t.Error("runBundle() should return error when neither --all nor --agent flag is provided")
	} else if !strings.Contains(err.Error(), "either --all or --agent flag is required") {
		t.Errorf("runBundle() error should mention --all or --agent flag requirement, got: %v", err)
	}
}

func TestBundleCommandMissingFramework(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "bundle-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current dir: %v", err)
	}
	defer func() {
		if chErr := os.Chdir(originalDir); chErr != nil {
			t.Logf("Warning: failed to change back to original directory: %v", chErr)
		}
	}()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp dir: %v", err)
	}

	// Set flags
	bundleAll = true
	bundleDryRun = false
	bundleOutput = ""

	err = runBundle(bundleCmd, []string{})
	if err == nil {
		t.Error("runBundle() should return error when framework not installed")
	} else if !strings.Contains(err.Error(), "framework not installed") {
		t.Errorf("runBundle() error should mention framework not installed, got: %v", err)
	}
}

func TestCollectAdditionalFiles(t *testing.T) {
	// Create a temporary directory structure
	tempDir, err := os.MkdirTemp("", "bundle-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test structure
	templatesDir := filepath.Join(tempDir, ".krci-ai", "templates")
	dataDir := filepath.Join(tempDir, ".krci-ai", "data")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		t.Fatalf("Failed to create data dir: %v", err)
	}

	// Create test files
	templateFile := filepath.Join(templatesDir, "test-template.md")
	dataFile := filepath.Join(dataDir, "test-data.md")
	if err := os.WriteFile(templateFile, []byte("template content"), 0644); err != nil {
		t.Fatalf("Failed to create template file: %v", err)
	}
	if err := os.WriteFile(dataFile, []byte("data content"), 0644); err != nil {
		t.Fatalf("Failed to create data file: %v", err)
	}

	// Test collectAdditionalFiles
	content := &BundleContent{
		Templates: make(map[string]string),
		DataFiles: make(map[string]string),
	}

	err = collectAdditionalFiles(tempDir, content)
	if err != nil {
		t.Fatalf("collectAdditionalFiles() failed: %v", err)
	}

	// Verify files were collected
	templatePath := "templates/test-template.md"
	dataPath := "data/test-data.md"

	if _, exists := content.Templates[templatePath]; !exists {
		t.Errorf("Templates should contain %q", templatePath)
	} else if content.Templates[templatePath] != "template content" {
		t.Errorf("Template content = %q, want %q", content.Templates[templatePath], "template content")
	}

	if _, exists := content.DataFiles[dataPath]; !exists {
		t.Errorf("DataFiles should contain %q", dataPath)
	} else if content.DataFiles[dataPath] != "data content" {
		t.Errorf("Data content = %q, want %q", content.DataFiles[dataPath], "data content")
	}
}

func TestWriteBundleFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "bundle-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	bundlePath := filepath.Join(tempDir, "test-bundle.md")
	testContent := "# Test Bundle\nTest content"

	err = writeBundleFile(bundlePath, testContent)
	if err != nil {
		t.Fatalf("writeBundleFile() failed: %v", err)
	}

	// Verify file was created with correct content
	content, err := os.ReadFile(bundlePath)
	if err != nil {
		t.Fatalf("Failed to read bundle file: %v", err)
	}
	if string(content) != testContent {
		t.Errorf("Bundle file content = %q, want %q", string(content), testContent)
	}
}

func TestWriteBundleFileInvalidPath(t *testing.T) {
	// Test writing to invalid path
	invalidPath := "/invalid/path/bundle.md"
	testContent := "test content"

	err := writeBundleFile(invalidPath, testContent)
	if err == nil {
		t.Error("writeBundleFile() should fail with invalid path")
	} else if !strings.Contains(err.Error(), "failed to create bundle file") {
		t.Errorf("writeBundleFile() error should mention file creation failure, got: %v", err)
	}
}

func TestBundleHeaderGeneration(t *testing.T) {
	var result strings.Builder
	addBundleHeader(&result)

	content := result.String()
	expectedStrings := []string{
		"# KubeRocketAI Framework Bundle",
		"**Generated:**",
		"## Usage Instructions",
		"### File Format Guide",
		"### For LLM Understanding",
		"==== FILE: <path> ====",
		"==== END FILE ====",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(content, expected) {
			t.Errorf("Bundle header should contain %q", expected)
		}
	}
}

func TestAddAgentSection(t *testing.T) {
	var result strings.Builder
	agent := AgentBundleInfo{
		FilePath: ".krci-ai/agents/test.yaml",
		Name:     "Test Agent",
		Role:     "Tester",
		Content:  "agent content",
	}
	content := &BundleContent{
		Tasks: map[string]string{
			"tasks/test-task.md": "task content",
		},
	}

	addAgentSection(&result, agent, content)

	output := result.String()
	expectedStrings := []string{
		"==== FILE: .krci-ai/agents/test.yaml ====",
		"agent content",
		"==== END FILE ====",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Agent section should contain %q", expected)
		}
	}
}

func TestAddSharedTemplates(t *testing.T) {
	var result strings.Builder
	content := &BundleContent{
		Templates: map[string]string{
			"templates/test.md": "template content",
		},
	}

	addSharedTemplates(&result, content)

	output := result.String()
	expectedStrings := []string{
		"# Shared Templates",
		"==== FILE: templates/test.md ====",
		"template content",
		"==== END FILE ====",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Shared templates should contain %q", expected)
		}
	}
}

func TestAddSharedTemplatesEmpty(t *testing.T) {
	var result strings.Builder
	content := &BundleContent{
		Templates: make(map[string]string),
	}

	addSharedTemplates(&result, content)

	output := result.String()
	if output != "" {
		t.Errorf("Empty templates should produce no output, got %q", output)
	}
}

func TestAddSharedDataFiles(t *testing.T) {
	var result strings.Builder
	content := &BundleContent{
		DataFiles: map[string]string{
			"data/test.md": "data content",
		},
	}

	addSharedDataFiles(&result, content)

	output := result.String()
	expectedStrings := []string{
		"# Reference Data",
		"==== FILE: data/test.md ====",
		"data content",
		"==== END FILE ====",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Shared data files should contain %q", expected)
		}
	}
}

func TestAddSharedDataFilesEmpty(t *testing.T) {
	var result strings.Builder
	content := &BundleContent{
		DataFiles: make(map[string]string),
	}

	addSharedDataFiles(&result, content)

	output := result.String()
	if output != "" {
		t.Errorf("Empty data files should produce no output, got %q", output)
	}
}
