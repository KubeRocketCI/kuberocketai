package assets

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KubeRocketCI/kuberocketai/internal/processor"
)

const (
	// IDE integration directories
	cursorRulesDir    = ".cursor/rules/krci-ai"
	claudeCommandsDir = ".claude/commands/krci-ai"
	vscodeModesDir    = ".github/chatmodes"
	windsurfRulesDir  = ".windsurf/rules"
)

// IDEIntegration defines the interface for IDE-specific integrations
type IDEIntegration interface {
	GetDirectoryPath() string
	GetFileExtension() string
	GenerateContent(agentName, role string, yamlContent []byte) string
}

// CursorIntegration implements IDEIntegration for Cursor IDE
type CursorIntegration struct {
	projectDir string
}

func (c *CursorIntegration) GetDirectoryPath() string {
	return filepath.Join(c.projectDir, cursorRulesDir)
}

func (c *CursorIntegration) GetFileExtension() string {
	return ".mdc"
}

func (c *CursorIntegration) GenerateContent(agentName, role string, yamlContent []byte) string {
	// Simple title case for agent name (capitalize first letter)
	titleCaseAgentName := strings.ToUpper(agentName[:1]) + agentName[1:]

	return fmt.Sprintf(`---
description:
globs: []
alwaysApply: false
---

# %s Agent Activation

CRITICAL: Carefully read the YAML agent definition below. Immediately activate the %s persona by following the activation instructions, and remain in this persona until you receive an explicit command to exit.

`+"```yaml\n%s```\n",
		titleCaseAgentName,
		role,
		string(yamlContent))
}

// ClaudeIntegration implements IDEIntegration for Claude Code
type ClaudeIntegration struct {
	projectDir string
}

func (c *ClaudeIntegration) GetDirectoryPath() string {
	return filepath.Join(c.projectDir, claudeCommandsDir)
}

func (c *ClaudeIntegration) GetFileExtension() string {
	return mdExtension
}

func (c *ClaudeIntegration) GenerateContent(agentName, role string, yamlContent []byte) string {
	return fmt.Sprintf(`# /%s Command

CRITICAL: Carefully read the YAML agent definition below. Immediately activate the %s persona by following the activation instructions, and remain in this persona until you receive an explicit command to exit.

`+"```yaml\n%s```\n",
		agentName,
		role,
		string(yamlContent))
}

// VSCodeIntegration implements IDEIntegration for VS Code
type VSCodeIntegration struct {
	targetDir string
}

func (v *VSCodeIntegration) GetDirectoryPath() string {
	return filepath.Join(v.targetDir, vscodeModesDir)
}

func (v *VSCodeIntegration) GetFileExtension() string {
	return ".chatmode.md"
}

func (v *VSCodeIntegration) GenerateContent(agentName, role string, yamlContent []byte) string {
	return fmt.Sprintf(`---
description: Activate %s role for specialized development assistance
tools: %s
---

# %s Agent Chat Mode

CRITICAL: Carefully read the YAML agent definition below. Immediately activate the %s persona by following the activation instructions, and remain in this persona until you receive an explicit command to exit.

`+"```yaml\n%s```\n",
		role,
		GitHubToolsList,
		role,
		role,
		string(yamlContent))
}

// WindsurfIntegration implements IDEIntegration for Windsurf IDE
type WindsurfIntegration struct {
	targetDir string
}

func (w *WindsurfIntegration) GetDirectoryPath() string {
	return filepath.Join(w.targetDir, windsurfRulesDir)
}

func (w *WindsurfIntegration) GetFileExtension() string {
	return mdExtension
}

func (w *WindsurfIntegration) GenerateContent(agentName, role string, yamlContent []byte) string {
	return fmt.Sprintf(`# %s Agent Rule

Activate the %s persona by following the agent definition below. This rule provides specialized development assistance for %s-related tasks.

## Agent Definition

`+"```yaml\n%s```\n",
		role,
		role,
		strings.ToLower(role),
		string(yamlContent))
}

// installIDEIntegration is a generic method for installing IDE integrations
func (i *Installer) installIDEIntegration(integration IDEIntegration, ideName string) error {
	// Create IDE-specific directory
	integrationPath := integration.GetDirectoryPath()
	if err := i.createDirectory(integrationPath); err != nil {
		return fmt.Errorf("failed to create %s directory: %w", ideName, err)
	}

	// Get list of agent files
	agentsPath := i.GetAgentsPath()
	agentFiles, err := filepath.Glob(filepath.Join(agentsPath, "*.yaml"))
	if err != nil {
		return fmt.Errorf("failed to find agent files: %w", err)
	}

	// Generate IDE-specific files for each agent
	for _, agentFile := range agentFiles {
		if err := i.generateIDEFile(agentFile, integration); err != nil {
			agentName := strings.TrimSuffix(filepath.Base(agentFile), ".yaml")
			return fmt.Errorf("failed to generate %s file for %s: %w", ideName, agentName, err)
		}
	}

	return nil
}

// generateIDEFile creates an IDE-specific file from an agent YAML file
func (i *Installer) generateIDEFile(agentFile string, integration IDEIntegration) error {
	// Read agent YAML file
	agentData, err := os.ReadFile(agentFile)
	if err != nil {
		return fmt.Errorf("failed to read agent file %s: %w", agentFile, err)
	}

	rawAgent, err := processor.UnmarshalAgentFile(agentFile)
	if err != nil {
		return err
	}

	agent := MakeAgent(agentFile, rawAgent, []Task{})

	// Generate output file path
	outputPath := filepath.Join(integration.GetDirectoryPath(), agent.ShortName+integration.GetFileExtension())

	// Generate content using the integration-specific logic
	content := integration.GenerateContent(agent.ShortName, agent.Role, agentData)

	// Write file
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outputPath, err)
	}

	return nil
}

// InstallCursorIntegration creates .cursor/rules directory and generates .mdc files for agents
func (i *Installer) InstallCursorIntegration() error {
	integration := &CursorIntegration{projectDir: i.projectDir}
	return i.installIDEIntegration(integration, "Cursor IDE")
}

// InstallClaudeIntegration creates .claude/commands directory and generates .md files for agents
func (i *Installer) InstallClaudeIntegration() error {
	integration := &ClaudeIntegration{projectDir: i.projectDir}
	return i.installIDEIntegration(integration, "Claude Code")
}

// InstallVSCodeIntegration creates .github/chatmodes directory and generates .chatmode.md files for agents
func (i *Installer) InstallVSCodeIntegration() error {
	integration := &VSCodeIntegration{targetDir: i.projectDir}
	return i.installIDEIntegration(integration, "VS Code")
}

// InstallWindsurfIntegration creates .windsurf/rules directory and generates .md files for agents
func (i *Installer) InstallWindsurfIntegration() error {
	integration := &WindsurfIntegration{targetDir: i.projectDir}
	return i.installIDEIntegration(integration, "Windsurf IDE")
}

// GetCursorRulesPath returns the path to the Cursor rules directory
func (i *Installer) GetCursorRulesPath() string {
	return filepath.Join(i.projectDir, cursorRulesDir)
}

// GetClaudeCommandsPath returns the path to the Claude commands directory
func (i *Installer) GetClaudeCommandsPath() string {
	return filepath.Join(i.projectDir, claudeCommandsDir)
}

// GetVSCodeChatmodesPath returns the path to the VS Code chatmodes directory
func (i *Installer) GetVSCodeChatmodesPath() string {
	return filepath.Join(i.projectDir, vscodeModesDir)
}

// GetWindsurfRulesPath returns the path to the Windsurf rules directory
func (i *Installer) GetWindsurfRulesPath() string {
	return filepath.Join(i.projectDir, windsurfRulesDir)
}

func (i *Installer) GetAgentsPath() string {
	return GetAgentsPath(i.krciPath)
}

func (i *Installer) GetTasksPath() string {
	return GetTasksPath(i.krciPath)
}

func (i *Installer) GetTemplatesPath() string {
	return GetTemplatesPath(i.krciPath)
}

func (i *Installer) GetDataPath() string {
	return GetDataPath(i.krciPath)
}
