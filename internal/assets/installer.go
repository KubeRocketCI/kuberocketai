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
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	// Main framework directory structure
	KrciAIDir    = ".krci-ai"
	BundleDir    = "bundle"
	DataDir      = "data"
	TasksDir     = "tasks"
	TemplatesDir = "templates"

	// Internal directory structure (unexported)
	agentsDir      = "agents"
	embeddedPrefix = "assets/framework/core"

	// IDE integration directories
	cursorRulesDir    = ".cursor/rules/krci-ai"
	claudeCommandsDir = ".claude/commands/krci-ai"
	vscodeModesDir    = ".github/chatmodes"
	windsurfRulesDir  = ".windsurf/rules"

	// File extensions
	mdExtension = ".md"

	// GitHub tools configuration
	GitHubToolsList = "['changes', 'codebase', 'editFiles', 'fetch', 'findTestFiles', 'githubRepo', 'problems', 'runCommands', 'search', 'searchResults', 'terminalLastCommand', 'usages']"
)

// AgentData represents the parsed YAML structure for agent files
type AgentData struct {
	Agent struct {
		Identity struct {
			Role string `yaml:"role"`
		} `yaml:"identity"`
	} `yaml:"agent"`
}

// IDEIntegration defines the interface for IDE-specific integrations
type IDEIntegration interface {
	GetDirectoryPath() string
	GetFileExtension() string
	GenerateContent(agentName, role string, yamlContent []byte) string
}

// CursorIntegration implements IDEIntegration for Cursor IDE
type CursorIntegration struct {
	targetDir string
}

func (c *CursorIntegration) GetDirectoryPath() string {
	return filepath.Join(c.targetDir, cursorRulesDir)
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
	targetDir string
}

func (c *ClaudeIntegration) GetDirectoryPath() string {
	return filepath.Join(c.targetDir, claudeCommandsDir)
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

// Installer handles installation of framework assets
type Installer struct {
	targetDir      string
	embeddedAssets embed.FS
}

// NewInstaller creates a new asset installer
func NewInstaller(targetDir string, embeddedAssets embed.FS) *Installer {
	if targetDir == "" {
		targetDir = "."
	}
	return &Installer{
		targetDir:      targetDir,
		embeddedAssets: embeddedAssets,
	}
}

// Install installs framework assets to the target directory
func (i *Installer) Install() error {
	krciPath := filepath.Join(i.targetDir, KrciAIDir)

	// Create main .krci-ai directory
	if err := i.createDirectory(krciPath); err != nil {
		return fmt.Errorf("failed to create .krci-ai directory: %w", err)
	}

	// Create subdirectories
	subdirs := []string{agentsDir, TasksDir, TemplatesDir, DataDir}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(krciPath, subdir)
		if err := i.createDirectory(dirPath); err != nil {
			return fmt.Errorf("failed to create %s directory: %w", subdir, err)
		}
	}

	// Copy embedded assets
	if err := i.copyEmbeddedAssets(krciPath); err != nil {
		return fmt.Errorf("failed to copy embedded assets: %w", err)
	}

	return nil
}

// createDirectory creates a directory if it doesn't exist
func (i *Installer) createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

// copyEmbeddedAssets copies all embedded assets to the target directory
func (i *Installer) copyEmbeddedAssets(krciPath string) error {
	// Walk through embedded filesystem
	return fs.WalkDir(i.embeddedAssets, embeddedPrefix, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory
		if path == embeddedPrefix {
			return nil
		}

		// Get relative path from embedded prefix
		relPath, err := filepath.Rel(embeddedPrefix, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Target path in the .krci-ai directory
		targetPath := filepath.Join(krciPath, relPath)

		if d.IsDir() {
			return i.createDirectory(targetPath)
		}

		// Copy file
		return i.copyEmbeddedFile(path, targetPath)
	})
}

// copyEmbeddedFile copies a single embedded file to the target location
func (i *Installer) copyEmbeddedFile(embeddedPath, targetPath string) error {
	// Read embedded file
	data, err := i.embeddedAssets.ReadFile(embeddedPath)
	if err != nil {
		return fmt.Errorf("failed to read embedded file %s: %w", embeddedPath, err)
	}

	// Ensure target directory exists
	targetDir := filepath.Dir(targetPath)
	if err := i.createDirectory(targetDir); err != nil {
		return fmt.Errorf("failed to create target directory %s: %w", targetDir, err)
	}

	// Write file to target
	if err := os.WriteFile(targetPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", targetPath, err)
	}

	return nil
}

// IsInstalled checks if the framework is properly installed in the target directory
func (i *Installer) IsInstalled() bool {
	krciPath := filepath.Join(i.targetDir, KrciAIDir)

	// Check if main directory exists
	if _, err := os.Stat(krciPath); os.IsNotExist(err) {
		return false
	}

	// Check if required subdirectories exist
	requiredDirs := []string{agentsDir, TasksDir, TemplatesDir, DataDir}
	for _, dir := range requiredDirs {
		dirPath := filepath.Join(krciPath, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// GetInstallationPath returns the full path to the .krci-ai directory
func (i *Installer) GetInstallationPath() string {
	return filepath.Join(i.targetDir, KrciAIDir)
}

// GetAgentsPath returns the path to the agents directory
func (i *Installer) GetAgentsPath() string {
	return filepath.Join(i.targetDir, KrciAIDir, agentsDir)
}

// GetTasksPath returns the path to the tasks directory
func (i *Installer) GetTasksPath() string {
	return filepath.Join(i.targetDir, KrciAIDir, TasksDir)
}

// GetTemplatesPath returns the path to the templates directory
func (i *Installer) GetTemplatesPath() string {
	return filepath.Join(i.targetDir, KrciAIDir, TemplatesDir)
}

// GetDataPath returns the path to the data directory
func (i *Installer) GetDataPath() string {
	return filepath.Join(i.targetDir, KrciAIDir, DataDir)
}

// GetCursorRulesPath returns the path to the Cursor rules directory
func (i *Installer) GetCursorRulesPath() string {
	return filepath.Join(i.targetDir, cursorRulesDir)
}

// GetClaudeCommandsPath returns the path to the Claude commands directory
func (i *Installer) GetClaudeCommandsPath() string {
	return filepath.Join(i.targetDir, claudeCommandsDir)
}

// GetVSCodeChatmodesPath returns the path to the VS Code chatmodes directory
func (i *Installer) GetVSCodeChatmodesPath() string {
	return filepath.Join(i.targetDir, vscodeModesDir)
}

// GetWindsurfRulesPath returns the path to the Windsurf rules directory
func (i *Installer) GetWindsurfRulesPath() string {
	return filepath.Join(i.targetDir, windsurfRulesDir)
}

// ValidateInstallation performs basic validation of the installation
func (i *Installer) ValidateInstallation() error {
	if !i.IsInstalled() {
		return fmt.Errorf("framework not installed - run 'krci-ai install' to install")
	}

	// Check that agents directory has files
	agentsPath := i.GetAgentsPath()
	agentFiles, err := filepath.Glob(filepath.Join(agentsPath, "*.yaml"))
	if err != nil {
		return fmt.Errorf("failed to check agent files: %w", err)
	}
	if len(agentFiles) == 0 {
		return fmt.Errorf("no agent files found in %s", agentsPath)
	}

	// Check that tasks directory has files
	tasksPath := i.GetTasksPath()
	taskFiles, err := filepath.Glob(filepath.Join(tasksPath, "*.md"))
	if err != nil {
		return fmt.Errorf("failed to check task files: %w", err)
	}
	if len(taskFiles) == 0 {
		return fmt.Errorf("no task files found in %s", tasksPath)
	}

	return nil
}

// ListInstalledAgents returns a list of installed agent file names
func (i *Installer) ListInstalledAgents() ([]string, error) {
	if !i.IsInstalled() {
		return nil, fmt.Errorf("framework not installed")
	}

	agentsPath := i.GetAgentsPath()
	agentFiles, err := filepath.Glob(filepath.Join(agentsPath, "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to list agent files: %w", err)
	}

	var agents []string
	for _, file := range agentFiles {
		agentName := strings.TrimSuffix(filepath.Base(file), ".yaml")
		agents = append(agents, agentName)
	}

	return agents, nil
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

	// Parse YAML to get agent identity for role
	var agent AgentData
	if err := yaml.Unmarshal(agentData, &agent); err != nil {
		return fmt.Errorf("failed to parse agent YAML: %w", err)
	}

	// Generate output file path
	agentName := strings.TrimSuffix(filepath.Base(agentFile), ".yaml")
	outputPath := filepath.Join(integration.GetDirectoryPath(), agentName+integration.GetFileExtension())

	// Generate content using the integration-specific logic
	content := integration.GenerateContent(agentName, agent.Agent.Identity.Role, agentData)

	// Write file
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outputPath, err)
	}

	return nil
}

// InstallCursorIntegration creates .cursor/rules directory and generates .mdc files for agents
func (i *Installer) InstallCursorIntegration() error {
	integration := &CursorIntegration{targetDir: i.targetDir}
	return i.installIDEIntegration(integration, "Cursor IDE")
}

// InstallClaudeIntegration creates .claude/commands directory and generates .md files for agents
func (i *Installer) InstallClaudeIntegration() error {
	integration := &ClaudeIntegration{targetDir: i.targetDir}
	return i.installIDEIntegration(integration, "Claude Code")
}

// InstallVSCodeIntegration creates .github/chatmodes directory and generates .chatmode.md files for agents
func (i *Installer) InstallVSCodeIntegration() error {
	integration := &VSCodeIntegration{targetDir: i.targetDir}
	return i.installIDEIntegration(integration, "VS Code")
}

// InstallWindsurfIntegration creates .windsurf/rules directory and generates .md files for agents
func (i *Installer) InstallWindsurfIntegration() error {
	integration := &WindsurfIntegration{targetDir: i.targetDir}
	return i.installIDEIntegration(integration, "Windsurf IDE")
}

// InstallSelective installs only specified agents and their dependencies using existing bundle logic
func (i *Installer) InstallSelective(agentNames []string) error {
	// Create embedded source and discovery following SOLID principles
	embeddedSource := NewEmbeddedSource(i.embeddedAssets)
	discovery := NewDiscoveryWithSource(i.targetDir, i.embeddedAssets, embeddedSource)

	// Validate agent names using unified validation
	if err := discovery.ValidateEmbeddedAgentNames(i.embeddedAssets, agentNames); err != nil {
		return fmt.Errorf("agent validation failed: %w", err)
	}

	// Get dependencies using unified discovery method
	agentDeps, err := discovery.DiscoverAgentsWithDependencies(agentNames...)
	if err != nil {
		return fmt.Errorf("dependency analysis failed: %w", err)
	}

	// Use simplified installation approach
	return i.installWithFilter(agentDeps)
}

// installWithFilter installs assets with optional filtering for selective installation
func (i *Installer) installWithFilter(agentDeps []AgentDependencyInfo) error {
	krciPath := filepath.Join(i.targetDir, KrciAIDir)

	// Create main .krci-ai directory
	if err := i.createDirectory(krciPath); err != nil {
		return fmt.Errorf("failed to create .krci-ai directory: %w", err)
	}

	// Create subdirectories
	subdirs := []string{agentsDir, TasksDir, TemplatesDir, DataDir}
	for _, subdir := range subdirs {
		dirPath := filepath.Join(krciPath, subdir)
		if err := i.createDirectory(dirPath); err != nil {
			return fmt.Errorf("failed to create %s directory: %w", subdir, err)
		}
	}

	// Install filtered assets
	return i.installFilteredAssets(krciPath, agentDeps)
}

// installFilteredAssets installs only the specified agents and their dependencies
func (i *Installer) installFilteredAssets(krciPath string, agentDeps []AgentDependencyInfo) error {
	// Track all files that need to be copied to avoid duplicates
	filesToCopy := make(map[string]bool)

	// Add agent files using constants
	for _, agent := range agentDeps {
		agentPath := fmt.Sprintf("%s/%s/%s.yaml", embeddedPrefix, agentsDir, agent.ShortName)
		filesToCopy[agentPath] = true
	}

	// Add task files using constants (no local tasks for embedded assets)
	for _, agent := range agentDeps {
		for _, task := range agent.Tasks {
			taskPath := fmt.Sprintf("%s/%s/%s", embeddedPrefix, TasksDir, task)
			filesToCopy[taskPath] = true
		}
	}

	// Add template files using constants
	for _, agent := range agentDeps {
		for _, template := range agent.Templates {
			templatePath := fmt.Sprintf("%s/%s/%s", embeddedPrefix, TemplatesDir, template)
			filesToCopy[templatePath] = true
		}
	}

	// Add data files using constants
	for _, agent := range agentDeps {
		for _, dataFile := range agent.DataFiles {
			dataPath := fmt.Sprintf("%s/%s/%s", embeddedPrefix, DataDir, dataFile)
			filesToCopy[dataPath] = true
		}
	}

	// Copy all required files using unified copy method
	for embeddedPath := range filesToCopy {
		if err := i.copyAssetFile(embeddedPath, krciPath); err != nil {
			return fmt.Errorf("failed to copy %s: %w", embeddedPath, err)
		}
	}

	return nil
}

// copyAssetFile copies a file from embedded assets to target directory
func (i *Installer) copyAssetFile(embeddedPath, krciPath string) error {
	// Check if file exists in embedded assets
	if _, err := i.embeddedAssets.Open(embeddedPath); err != nil {
		// File doesn't exist, skip silently (might be optional dependency)
		return nil
	}

	// Get relative path from embedded prefix using constant
	relPath, err := filepath.Rel(embeddedPrefix, embeddedPath)
	if err != nil {
		return fmt.Errorf("failed to get relative path for %s: %w", embeddedPath, err)
	}

	// Target path in the .krci-ai directory
	targetPath := filepath.Join(krciPath, relPath)

	// Read embedded file
	data, err := i.embeddedAssets.ReadFile(embeddedPath)
	if err != nil {
		return fmt.Errorf("failed to read embedded file %s: %w", embeddedPath, err)
	}

	// Ensure target directory exists
	targetDir := filepath.Dir(targetPath)
	if err := i.createDirectory(targetDir); err != nil {
		return fmt.Errorf("failed to create target directory %s: %w", targetDir, err)
	}

	// Write file to target
	if err := os.WriteFile(targetPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", targetPath, err)
	}

	return nil
}
