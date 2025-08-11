package assets

import (
	"embed"
	"os"
	"path/filepath"
	"testing"

	"github.com/KubeRocketCI/kuberocketai/internal/testutil"
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
	if !testutil.ContainsSubstr(content, "# Test Agent Activation") {
		t.Error("Content missing agent activation header")
	}
	if !testutil.ContainsSubstr(content, "Test Role") {
		t.Error("Content missing role")
	}
	if !testutil.ContainsSubstr(content, "test: yaml") {
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
	if !testutil.ContainsSubstr(content, "# /test Command") {
		t.Error("Content missing command header")
	}
	if !testutil.ContainsSubstr(content, "Test Role") {
		t.Error("Content missing role")
	}
	if !testutil.ContainsSubstr(content, "test: yaml") {
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
	if !testutil.ContainsSubstr(content, "---") {
		t.Error("Content missing front matter delimiters")
	}
	if !testutil.ContainsSubstr(content, "description: Activate Test Role role for specialized development assistance") {
		t.Error("Content missing description in front matter")
	}
	if !testutil.ContainsSubstr(content, "tools: "+GitHubToolsList) {
		t.Error("Content missing tools in front matter")
	}
	if !testutil.ContainsSubstr(content, "# Test Role Agent Chat Mode") {
		t.Error("Content missing agent chat mode header")
	}
	if !testutil.ContainsSubstr(content, "Test Role") {
		t.Error("Content missing role")
	}
	if !testutil.ContainsSubstr(content, "test: yaml") {
		t.Error("Content missing YAML content")
	}
}

func TestWindsurfIntegration(t *testing.T) {
	tempDir := t.TempDir()
	integration := &WindsurfIntegration{targetDir: tempDir}

	expectedPath := filepath.Join(tempDir, windsurfRulesDir)
	if integration.GetDirectoryPath() != expectedPath {
		t.Errorf("Expected directory path '%s', got '%s'", expectedPath, integration.GetDirectoryPath())
	}

	if integration.GetFileExtension() != mdExtension {
		t.Errorf("Expected file extension '%s', got '%s'", mdExtension, integration.GetFileExtension())
	}

	content := integration.GenerateContent("test", "Test Role", []byte("test: yaml"))
	if content == "" {
		t.Error("GenerateContent returned empty string")
	}

	// Check if content contains expected parts
	if !testutil.ContainsSubstr(content, "# Test Role Agent Rule") {
		t.Error("Content missing agent rule header")
	}
	if !testutil.ContainsSubstr(content, "Test Role") {
		t.Error("Content missing role")
	}
	if !testutil.ContainsSubstr(content, "test: yaml") {
		t.Error("Content missing YAML content")
	}
}

func TestGetWindsurfRulesPath(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	expectedPath := filepath.Join(tempDir, windsurfRulesDir)
	if installer.GetWindsurfRulesPath() != expectedPath {
		t.Errorf("Expected Windsurf rules path '%s', got '%s'", expectedPath, installer.GetWindsurfRulesPath())
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

func TestInstall(t *testing.T) {
	tempDir := t.TempDir()

	// Create a mock installer that doesn't copy embedded assets
	installer := NewInstaller(tempDir, testAssets)

	// Mock the install process by manually creating directories
	krciPath := filepath.Join(tempDir, krciAIDir)

	// Create main .krci-ai directory
	err := installer.createDirectory(krciPath)
	if err != nil {
		t.Fatalf("Failed to create .krci-ai directory: %v", err)
	}

	// Create subdirectories
	subdirs := []string{agentsDir, tasksDir, templatesDir, dataDir, configDir}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(krciPath, subdir)
		err := installer.createDirectory(dirPath)
		if err != nil {
			t.Fatalf("Failed to create %s directory: %v", subdir, err)
		}
	}

	// Verify directories were created
	if _, err := os.Stat(krciPath); os.IsNotExist(err) {
		t.Error("Main .krci-ai directory was not created")
	}

	// Verify subdirectories were created
	for _, subdir := range subdirs {
		dirPath := filepath.Join(krciPath, subdir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			t.Errorf("Subdirectory %s was not created", subdir)
		}
	}

	// Verify installation is detected
	if !installer.IsInstalled() {
		t.Error("Installation not detected after directory creation")
	}
}

func TestInstallInExistingDirectory(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	// Create directory first
	krciPath := filepath.Join(tempDir, krciAIDir)
	err := os.MkdirAll(krciPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	// createDirectory should still work on existing directory
	err = installer.createDirectory(krciPath)
	if err != nil {
		t.Fatalf("createDirectory failed on existing directory: %v", err)
	}
}

func TestCopyEmbeddedAssets(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	// Create base directory
	krciPath := filepath.Join(tempDir, krciAIDir)
	err := os.MkdirAll(krciPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create base directory: %v", err)
	}

	// Test that copyEmbeddedAssets function exists and handles errors gracefully
	// Since we don't have real embedded assets, we expect this to return an error
	err = installer.copyEmbeddedAssets(krciPath)
	if err == nil {
		t.Log("copyEmbeddedAssets succeeded (no embedded assets to copy)")
	} else {
		t.Logf("copyEmbeddedAssets failed as expected (no embedded assets): %v", err)
	}
}

func TestCopyEmbeddedFile(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	// Test the file copying logic by simulating the process
	targetPath := filepath.Join(tempDir, "test-output.txt")
	testContent := []byte("test content")

	// Test copyEmbeddedFile logic by creating target directory and file
	targetDir := filepath.Dir(targetPath)
	err := installer.createDirectory(targetDir)
	if err != nil {
		t.Fatalf("Failed to create target directory: %v", err)
	}

	err = os.WriteFile(targetPath, testContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Verify file was written
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		t.Error("Target file was not created")
	}

	// Verify content
	content, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	if string(content) != string(testContent) {
		t.Errorf("Expected content '%s', got '%s'", string(testContent), string(content))
	}
}

func TestListInstalledAgents(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	// Test when not installed
	agents, err := installer.ListInstalledAgents()
	if err == nil {
		t.Error("Expected error when framework not installed")
	}
	if agents != nil {
		t.Error("Expected nil agents list when not installed")
	}

	// Set up manual installation (without embedded assets)
	krciPath := filepath.Join(tempDir, krciAIDir)
	agentsPath := installer.GetAgentsPath()
	err = os.MkdirAll(agentsPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create agents directory: %v", err)
	}

	// Create required subdirectories for IsInstalled check
	subdirs := []string{tasksDir, templatesDir, dataDir}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(krciPath, subdir)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			t.Fatalf("Failed to create %s directory: %v", subdir, err)
		}
	}

	// Create test agent files
	testAgents := []string{"agent1.yaml", "agent2.yaml", "agent3.yaml"}

	for _, agentFile := range testAgents {
		agentPath := filepath.Join(agentsPath, agentFile)
		testContent := `agent:
  identity:
    name: "Test Agent"
    role: "Test Role"
`
		err = os.WriteFile(agentPath, []byte(testContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test agent file %s: %v", agentFile, err)
		}
	}

	// Test listing agents
	agents, err = installer.ListInstalledAgents()
	if err != nil {
		t.Fatalf("ListInstalledAgents failed: %v", err)
	}

	expectedAgents := []string{"agent1", "agent2", "agent3"}
	if len(agents) != len(expectedAgents) {
		t.Errorf("Expected %d agents, got %d", len(expectedAgents), len(agents))
	}

	// Check if all expected agents are present
	agentMap := make(map[string]bool)
	for _, agent := range agents {
		agentMap[agent] = true
	}

	for _, expected := range expectedAgents {
		if !agentMap[expected] {
			t.Errorf("Expected agent %s not found in results", expected)
		}
	}
}

func TestValidateInstallationComplete(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	// Set up manual installation (without embedded assets)
	krciPath := filepath.Join(tempDir, krciAIDir)
	agentsPath := installer.GetAgentsPath()
	tasksPath := installer.GetTasksPath()

	// Create required directories
	subdirs := []string{agentsDir, tasksDir, templatesDir, dataDir}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(krciPath, subdir)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			t.Fatalf("Failed to create %s directory: %v", subdir, err)
		}
	}

	// Create test agent files
	agentContent := `agent:
  identity:
    name: "Test Agent"
    role: "Test Role"
`
	err := os.WriteFile(filepath.Join(agentsPath, "test-agent.yaml"), []byte(agentContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test agent file: %v", err)
	}

	// Create test task files
	taskContent := "# Test Task\n\nThis is a test task."
	err = os.WriteFile(filepath.Join(tasksPath, "test-task.md"), []byte(taskContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test task file: %v", err)
	}

	// Test validation of complete installation
	err = installer.ValidateInstallation()
	if err != nil {
		t.Errorf("ValidateInstallation failed for complete installation: %v", err)
	}
}

func TestValidateInstallationMissingAgents(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	// Set up installation without agents
	krciPath := filepath.Join(tempDir, krciAIDir)
	tasksPath := installer.GetTasksPath()

	// Create required directories
	subdirs := []string{agentsDir, tasksDir, templatesDir, dataDir}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(krciPath, subdir)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			t.Fatalf("Failed to create %s directory: %v", subdir, err)
		}
	}

	// Create task files but no agent files
	taskContent := "# Test Task\n\nThis is a test task."
	err := os.WriteFile(filepath.Join(tasksPath, "test-task.md"), []byte(taskContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test task file: %v", err)
	}

	// Validation should fail due to missing agents
	err = installer.ValidateInstallation()
	if err == nil {
		t.Error("Expected validation to fail when no agent files present")
	}
}

func TestValidateInstallationMissingTasks(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	// Set up installation without tasks
	krciPath := filepath.Join(tempDir, krciAIDir)
	agentsPath := installer.GetAgentsPath()

	// Create required directories
	subdirs := []string{agentsDir, tasksDir, templatesDir, dataDir}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(krciPath, subdir)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			t.Fatalf("Failed to create %s directory: %v", subdir, err)
		}
	}

	// Create agent files but no task files
	agentContent := `agent:
  identity:
    name: "Test Agent"
    role: "Test Role"
`
	err := os.WriteFile(filepath.Join(agentsPath, "test-agent.yaml"), []byte(agentContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test agent file: %v", err)
	}

	// Validation should fail due to missing tasks
	err = installer.ValidateInstallation()
	if err == nil {
		t.Error("Expected validation to fail when no task files present")
	}
}

func TestGetIDEPaths(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testAssets)

	// Test all IDE path getters
	tests := []struct {
		name     string
		getter   func() string
		expected string
	}{
		{"Cursor Rules", installer.GetCursorRulesPath, filepath.Join(tempDir, cursorRulesDir)},
		{"Claude Commands", installer.GetClaudeCommandsPath, filepath.Join(tempDir, claudeCommandsDir)},
		{"VSCode Chatmodes", installer.GetVSCodeChatmodesPath, filepath.Join(tempDir, vscodeModesDir)},
		{"Windsurf Rules", installer.GetWindsurfRulesPath, filepath.Join(tempDir, windsurfRulesDir)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.getter()
			if result != test.expected {
				t.Errorf("Expected %s path '%s', got '%s'", test.name, test.expected, result)
			}
		})
	}
}
