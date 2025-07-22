package assets

import (
	"embed"
	"os"
	"path/filepath"
	"testing"
)

//go:embed testdata/*
var testIntegrationAssets embed.FS

func TestInstallIDEIntegration(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create a minimal .krci-ai structure with test agent
	krciPath := filepath.Join(tempDir, krciAIDir)
	agentsPath := filepath.Join(krciPath, agentsDir)
	err := os.MkdirAll(agentsPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create agents directory: %v", err)
	}

	// Create a test agent file
	testAgentContent := `agent:
  identity:
    name: TestAgent
    id: test-v1
    version: "1.0.0"
    description: "Test agent for unit testing"
    role: "Test Role"
    goal: "Test agent functionality"
    icon: "🧪"

  activation_prompt:
    - Test activation prompt

  principles:
    - "Test principle"

  commands:
    help: "Show test commands"
    test: "Execute test command"

  tasks:
    - ./test-task.md`

	agentFile := filepath.Join(agentsPath, "test.yaml")
	err = os.WriteFile(agentFile, []byte(testAgentContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test agent file: %v", err)
	}

	// Test Cursor integration
	cursorIntegration := &CursorIntegration{targetDir: tempDir}
	err = installer.installIDEIntegration(cursorIntegration, "Cursor IDE")
	if err != nil {
		t.Errorf("Failed to install Cursor integration: %v", err)
	}

	// Verify Cursor files were created
	cursorFile := filepath.Join(tempDir, cursorRulesDir, "test.mdc")
	if _, err := os.Stat(cursorFile); os.IsNotExist(err) {
		t.Error("Cursor .mdc file was not created")
	}

	// Test Claude integration
	claudeIntegration := &ClaudeIntegration{targetDir: tempDir}
	err = installer.installIDEIntegration(claudeIntegration, "Claude Code")
	if err != nil {
		t.Errorf("Failed to install Claude integration: %v", err)
	}

	// Verify Claude files were created
	claudeFile := filepath.Join(tempDir, claudeCommandsDir, "test.md")
	if _, err := os.Stat(claudeFile); os.IsNotExist(err) {
		t.Error("Claude .md file was not created")
	}

	// Test VS Code integration
	vscodeIntegration := &VSCodeIntegration{targetDir: tempDir}
	err = installer.installIDEIntegration(vscodeIntegration, "VS Code")
	if err != nil {
		t.Errorf("Failed to install VS Code integration: %v", err)
	}

	// Verify VS Code files were created
	vscodeFile := filepath.Join(tempDir, vscodeModesDir, "test.chatmode.md")
	if _, err := os.Stat(vscodeFile); os.IsNotExist(err) {
		t.Error("VS Code .chatmode.md file was not created")
	}
}

func TestInstallCursorIntegration(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create a minimal .krci-ai structure with test agent
	krciPath := filepath.Join(tempDir, krciAIDir)
	agentsPath := filepath.Join(krciPath, agentsDir)
	err := os.MkdirAll(agentsPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create agents directory: %v", err)
	}

	// Create a test agent file
	testAgentContent := `agent:
  identity:
    role: "Senior Developer"`

	agentFile := filepath.Join(agentsPath, "dev.yaml")
	err = os.WriteFile(agentFile, []byte(testAgentContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test agent file: %v", err)
	}

	// Test installation
	err = installer.InstallCursorIntegration()
	if err != nil {
		t.Errorf("InstallCursorIntegration failed: %v", err)
	}

	// Verify directory was created
	cursorDir := filepath.Join(tempDir, cursorRulesDir)
	if _, err := os.Stat(cursorDir); os.IsNotExist(err) {
		t.Error("Cursor rules directory was not created")
	}

	// Verify .mdc file was created
	mdcFile := filepath.Join(cursorDir, "dev.mdc")
	if _, err := os.Stat(mdcFile); os.IsNotExist(err) {
		t.Error("Cursor .mdc file was not created")
	}

	// Verify file content
	content, err := os.ReadFile(mdcFile)
	if err != nil {
		t.Fatalf("Failed to read .mdc file: %v", err)
	}

	contentStr := string(content)
	if !contains(contentStr, "Dev Agent Activation") {
		t.Error("Content missing agent activation header")
	}
	if !contains(contentStr, "Senior Developer") {
		t.Error("Content missing role")
	}
}

func TestInstallClaudeIntegration(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create a minimal .krci-ai structure with test agent
	krciPath := filepath.Join(tempDir, krciAIDir)
	agentsPath := filepath.Join(krciPath, agentsDir)
	err := os.MkdirAll(agentsPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create agents directory: %v", err)
	}

	// Create a test agent file
	testAgentContent := `agent:
  identity:
    role: "Product Owner"`

	agentFile := filepath.Join(agentsPath, "po.yaml")
	err = os.WriteFile(agentFile, []byte(testAgentContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test agent file: %v", err)
	}

	// Test installation
	err = installer.InstallClaudeIntegration()
	if err != nil {
		t.Errorf("InstallClaudeIntegration failed: %v", err)
	}

	// Verify directory was created
	claudeDir := filepath.Join(tempDir, claudeCommandsDir)
	if _, err := os.Stat(claudeDir); os.IsNotExist(err) {
		t.Error("Claude commands directory was not created")
	}

	// Verify .md file was created
	mdFile := filepath.Join(claudeDir, "po.md")
	if _, err := os.Stat(mdFile); os.IsNotExist(err) {
		t.Error("Claude .md file was not created")
	}

	// Verify file content
	content, err := os.ReadFile(mdFile)
	if err != nil {
		t.Fatalf("Failed to read .md file: %v", err)
	}

	contentStr := string(content)
	if !contains(contentStr, "# /po Command") {
		t.Error("Content missing command header")
	}
	if !contains(contentStr, "Product Owner") {
		t.Error("Content missing role")
	}
}

func TestInstallVSCodeIntegration(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create a minimal .krci-ai structure with test agent
	krciPath := filepath.Join(tempDir, krciAIDir)
	agentsPath := filepath.Join(krciPath, agentsDir)
	err := os.MkdirAll(agentsPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create agents directory: %v", err)
	}

	// Create a test agent file
	testAgentContent := `agent:
  identity:
    role: "Software Developer"`

	agentFile := filepath.Join(agentsPath, "dev.yaml")
	err = os.WriteFile(agentFile, []byte(testAgentContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test agent file: %v", err)
	}

	// Test installation
	err = installer.InstallVSCodeIntegration()
	if err != nil {
		t.Errorf("InstallVSCodeIntegration failed: %v", err)
	}

	// Verify directory was created
	vscodeDir := filepath.Join(tempDir, vscodeModesDir)
	if _, err := os.Stat(vscodeDir); os.IsNotExist(err) {
		t.Error("VS Code chatmodes directory was not created")
	}

	// Verify .chatmode.md file was created
	chatmodeFile := filepath.Join(vscodeDir, "dev.chatmode.md")
	if _, err := os.Stat(chatmodeFile); os.IsNotExist(err) {
		t.Error("VS Code .chatmode.md file was not created")
	}

	// Verify file content
	content, err := os.ReadFile(chatmodeFile)
	if err != nil {
		t.Fatalf("Failed to read .chatmode.md file: %v", err)
	}

	contentStr := string(content)
	if !contains(contentStr, "---") {
		t.Error("Content missing front matter delimiters")
	}
	if !contains(contentStr, "description: Activate Software Developer role for specialized development assistance") {
		t.Error("Content missing description in front matter")
	}
	if !contains(contentStr, "tools: "+GitHubToolsList) {
		t.Error("Content missing tools in front matter")
	}
	if !contains(contentStr, "# Software Developer Agent Chat Mode") {
		t.Error("Content missing agent chat mode header")
	}
	if !contains(contentStr, "Software Developer") {
		t.Error("Content missing role")
	}
}

func TestGenerateIDEFile(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create test agent file
	testAgentContent := `agent:
  identity:
    role: "Test Role"`

	agentFile := filepath.Join(tempDir, "test.yaml")
	err := os.WriteFile(agentFile, []byte(testAgentContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test agent file: %v", err)
	}

	// Test with Cursor integration
	cursorIntegration := &CursorIntegration{targetDir: tempDir}
	err = os.MkdirAll(cursorIntegration.GetDirectoryPath(), 0755)
	if err != nil {
		t.Fatalf("Failed to create cursor directory: %v", err)
	}

	err = installer.generateIDEFile(agentFile, cursorIntegration)
	if err != nil {
		t.Errorf("generateIDEFile failed for Cursor: %v", err)
	}

	// Verify file was created
	outputFile := filepath.Join(cursorIntegration.GetDirectoryPath(), "test.mdc")
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Error("IDE file was not created")
	}

	// Test with Claude integration
	claudeIntegration := &ClaudeIntegration{targetDir: tempDir}
	err = os.MkdirAll(claudeIntegration.GetDirectoryPath(), 0755)
	if err != nil {
		t.Fatalf("Failed to create claude directory: %v", err)
	}

	err = installer.generateIDEFile(agentFile, claudeIntegration)
	if err != nil {
		t.Errorf("generateIDEFile failed for Claude: %v", err)
	}

	// Verify file was created
	outputFile = filepath.Join(claudeIntegration.GetDirectoryPath(), "test.md")
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Error("IDE file was not created")
	}

	// Test with VS Code integration
	vscodeIntegration := &VSCodeIntegration{targetDir: tempDir}
	err = os.MkdirAll(vscodeIntegration.GetDirectoryPath(), 0755)
	if err != nil {
		t.Fatalf("Failed to create vscode directory: %v", err)
	}

	err = installer.generateIDEFile(agentFile, vscodeIntegration)
	if err != nil {
		t.Errorf("generateIDEFile failed for VS Code: %v", err)
	}

	// Verify file was created
	outputFile = filepath.Join(vscodeIntegration.GetDirectoryPath(), "test.chatmode.md")
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Error("VS Code IDE file was not created")
	}
}

func TestGenerateIDEFileWithInvalidYAML(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create invalid agent file
	invalidContent := `invalid: yaml: content: [unclosed`

	agentFile := filepath.Join(tempDir, "invalid.yaml")
	err := os.WriteFile(agentFile, []byte(invalidContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid agent file: %v", err)
	}

	cursorIntegration := &CursorIntegration{targetDir: tempDir}
	err = os.MkdirAll(cursorIntegration.GetDirectoryPath(), 0755)
	if err != nil {
		t.Fatalf("Failed to create cursor directory: %v", err)
	}

	err = installer.generateIDEFile(agentFile, cursorIntegration)
	if err == nil {
		t.Error("Expected error for invalid YAML but got none")
	}
}

func TestGenerateIDEFileWithNonExistentFile(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	nonExistentFile := filepath.Join(tempDir, "nonexistent.yaml")
	cursorIntegration := &CursorIntegration{targetDir: tempDir}

	err := installer.generateIDEFile(nonExistentFile, cursorIntegration)
	if err == nil {
		t.Error("Expected error for non-existent file but got none")
	}
}
