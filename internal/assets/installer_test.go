package assets

import (
	"embed"
	"os"
	"path/filepath"
	"testing"
)

//go:embed testdata/*
var testAssets embed.FS

func TestNewInstaller(t *testing.T) {
	installer := NewInstaller("test", testAssets)
	if installer == nil {
		t.Fatal("NewInstaller returned nil")
	}
	if installer.targetDir != "test" {
		t.Errorf("Expected targetDir 'test', got '%s'", installer.targetDir)
	}
}

func TestNewInstallerWithEmptyDir(t *testing.T) {
	installer := NewInstaller("", testAssets)
	if installer.targetDir != "." {
		t.Errorf("Expected targetDir '.', got '%s'", installer.targetDir)
	}
}

func TestGetInstallationPath(t *testing.T) {
	installer := NewInstaller("test", testAssets)
	expected := filepath.Join("test", krciAIDir)
	if installer.GetInstallationPath() != expected {
		t.Errorf("Expected path '%s', got '%s'", expected, installer.GetInstallationPath())
	}
}

func TestGetAgentsPath(t *testing.T) {
	installer := NewInstaller("test", testAssets)
	expected := filepath.Join("test", krciAIDir, agentsDir)
	if installer.GetAgentsPath() != expected {
		t.Errorf("Expected path '%s', got '%s'", expected, installer.GetAgentsPath())
	}
}

func TestGetTasksPath(t *testing.T) {
	installer := NewInstaller("test", testAssets)
	expected := filepath.Join("test", krciAIDir, tasksDir)
	if installer.GetTasksPath() != expected {
		t.Errorf("Expected path '%s', got '%s'", expected, installer.GetTasksPath())
	}
}

func TestGetTemplatesPath(t *testing.T) {
	installer := NewInstaller("test", testAssets)
	expected := filepath.Join("test", krciAIDir, templatesDir)
	if installer.GetTemplatesPath() != expected {
		t.Errorf("Expected path '%s', got '%s'", expected, installer.GetTemplatesPath())
	}
}

func TestGetDataPath(t *testing.T) {
	installer := NewInstaller("test", testAssets)
	expected := filepath.Join("test", krciAIDir, dataDir)
	if installer.GetDataPath() != expected {
		t.Errorf("Expected path '%s', got '%s'", expected, installer.GetDataPath())
	}
}

func TestIsInstalledWhenNotInstalled(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	if installer.IsInstalled() {
		t.Error("Expected IsInstalled to return false for non-existent installation")
	}
}

func TestIsInstalledWhenInstalled(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	// Create the required directories
	krciPath := filepath.Join(tempDir, krciAIDir)
	requiredDirs := []string{agentsDir, tasksDir, templatesDir, dataDir}

	err := os.MkdirAll(krciPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create krci-ai directory: %v", err)
	}

	for _, dir := range requiredDirs {
		err := os.MkdirAll(filepath.Join(krciPath, dir), 0755)
		if err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	if !installer.IsInstalled() {
		t.Error("Expected IsInstalled to return true for existing installation")
	}
}

func TestCreateDirectory(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	testPath := filepath.Join(tempDir, "test-dir")
	err := installer.createDirectory(testPath)
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Error("Directory was not created")
	}
}

func TestCursorIntegration(t *testing.T) {
	tempDir := t.TempDir()
	integration := &CursorIntegration{targetDir: tempDir}

	expectedPath := filepath.Join(tempDir, cursorRulesDir)
	if integration.GetDirectoryPath() != expectedPath {
		t.Errorf("Expected directory path '%s', got '%s'", expectedPath, integration.GetDirectoryPath())
	}

	if integration.GetFileExtension() != ".mdc" {
		t.Errorf("Expected file extension '.mdc', got '%s'", integration.GetFileExtension())
	}

	content := integration.GenerateContent("test", "Test Role", []byte("test: yaml"))
	if content == "" {
		t.Error("GenerateContent returned empty string")
	}

	// Check if content contains expected parts
	if !contains(content, "# Test Agent Activation") {
		t.Error("Content missing agent activation header")
	}
	if !contains(content, "Test Role") {
		t.Error("Content missing role")
	}
	if !contains(content, "test: yaml") {
		t.Error("Content missing YAML content")
	}
}

func TestClaudeIntegration(t *testing.T) {
	tempDir := t.TempDir()
	integration := &ClaudeIntegration{targetDir: tempDir}

	expectedPath := filepath.Join(tempDir, claudeCommandsDir)
	if integration.GetDirectoryPath() != expectedPath {
		t.Errorf("Expected directory path '%s', got '%s'", expectedPath, integration.GetDirectoryPath())
	}

	if integration.GetFileExtension() != ".md" {
		t.Errorf("Expected file extension '.md', got '%s'", integration.GetFileExtension())
	}

	content := integration.GenerateContent("test", "Test Role", []byte("test: yaml"))
	if content == "" {
		t.Error("GenerateContent returned empty string")
	}

	// Check if content contains expected parts
	if !contains(content, "# /test Command") {
		t.Error("Content missing command header")
	}
	if !contains(content, "Test Role") {
		t.Error("Content missing role")
	}
	if !contains(content, "test: yaml") {
		t.Error("Content missing YAML content")
	}
}

func TestVSCodeIntegration(t *testing.T) {
	tempDir := t.TempDir()
	integration := &VSCodeIntegration{targetDir: tempDir}

	expectedPath := filepath.Join(tempDir, vscodeModesDir)
	if integration.GetDirectoryPath() != expectedPath {
		t.Errorf("Expected directory path '%s', got '%s'", expectedPath, integration.GetDirectoryPath())
	}

	if integration.GetFileExtension() != ".chatmode.md" {
		t.Errorf("Expected file extension '.chatmode.md', got '%s'", integration.GetFileExtension())
	}

	content := integration.GenerateContent("test", "Test Role", []byte("test: yaml"))
	if content == "" {
		t.Error("GenerateContent returned empty string")
	}

	// Check if content contains expected parts
	if !contains(content, "---") {
		t.Error("Content missing front matter delimiters")
	}
	if !contains(content, "description: Activate Test Role role for specialized development assistance") {
		t.Error("Content missing description in front matter")
	}
	if !contains(content, "tools: "+GitHubToolsList) {
		t.Error("Content missing tools in front matter")
	}
	if !contains(content, "# Test Role Agent Chat Mode") {
		t.Error("Content missing agent chat mode header")
	}
	if !contains(content, "Test Role") {
		t.Error("Content missing role")
	}
	if !contains(content, "test: yaml") {
		t.Error("Content missing YAML content")
	}
}

func TestValidateInstallationWithoutInstallation(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	err := installer.ValidateInstallation()
	if err == nil {
		t.Error("Expected validation to fail for non-existent installation")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && containsHelper(s, substr)))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
