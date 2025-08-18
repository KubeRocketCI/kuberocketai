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
	"embed"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/*
var testIntegrationAssets embed.FS

func TestInstallIDEIntegration(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create a minimal .krci-ai structure with test agent
	krciPath := filepath.Join(tempDir, KrciAIDir)
	agentsPath := filepath.Join(krciPath, agentsDir)
	err := os.MkdirAll(agentsPath, 0755)
	require.NoError(t, err, "Failed to create agents directory")

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
	require.NoError(t, err, "Failed to create test agent file")

	// Test Cursor integration
	cursorIntegration := &CursorIntegration{targetDir: tempDir}
	err = installer.installIDEIntegration(cursorIntegration, "Cursor IDE")
	assert.NoError(t, err, "Failed to install Cursor integration")

	// Verify Cursor files were created
	cursorFile := filepath.Join(tempDir, cursorRulesDir, "test.mdc")
	assert.FileExists(t, cursorFile, "Cursor .mdc file should be created")

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
	krciPath := filepath.Join(tempDir, KrciAIDir)
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
	if !strings.Contains(contentStr, "Dev Agent Activation") {
		t.Error("Content missing agent activation header")
	}
	if !strings.Contains(contentStr, "Senior Developer") {
		t.Error("Content missing role")
	}
}

func TestInstallClaudeIntegration(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create a minimal .krci-ai structure with test agent
	krciPath := filepath.Join(tempDir, KrciAIDir)
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
	if !strings.Contains(contentStr, "# /po Command") {
		t.Error("Content missing command header")
	}
	if !strings.Contains(contentStr, "Product Owner") {
		t.Error("Content missing role")
	}
}

func TestInstallVSCodeIntegration(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create a minimal .krci-ai structure with test agent
	krciPath := filepath.Join(tempDir, KrciAIDir)
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
	if !strings.Contains(contentStr, "---") {
		t.Error("Content missing front matter delimiters")
	}
	if !strings.Contains(contentStr, "description: Activate Software Developer role for specialized development assistance") {
		t.Error("Content missing description in front matter")
	}
	if !strings.Contains(contentStr, "tools: "+GitHubToolsList) {
		t.Error("Content missing tools in front matter")
	}
	if !strings.Contains(contentStr, "# Software Developer Agent Chat Mode") {
		t.Error("Content missing agent chat mode header")
	}
	if !strings.Contains(contentStr, "Software Developer") {
		t.Error("Content missing role")
	}
}

func TestInstallWindsurfIntegration(t *testing.T) {
	tempDir := t.TempDir()
	installer := NewInstaller(tempDir, testIntegrationAssets)

	// Create a minimal .krci-ai structure with test agent
	krciPath := filepath.Join(tempDir, KrciAIDir)
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
	err = installer.InstallWindsurfIntegration()
	if err != nil {
		t.Errorf("InstallWindsurfIntegration failed: %v", err)
	}

	// Verify directory was created
	windsurfDir := filepath.Join(tempDir, windsurfRulesDir)
	if _, err := os.Stat(windsurfDir); os.IsNotExist(err) {
		t.Error("Windsurf rules directory was not created")
	}

	// Verify .md file was created
	mdFile := filepath.Join(windsurfDir, "dev.md")
	if _, err := os.Stat(mdFile); os.IsNotExist(err) {
		t.Error("dev.md file was not created")
	}

	// Verify file content
	content, err := os.ReadFile(mdFile)
	if err != nil {
		t.Fatalf("Failed to read .md file: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "# Software Developer Agent Rule") {
		t.Error("Content missing agent rule header")
	}
	if !strings.Contains(contentStr, "Software Developer") {
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
