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
)

const (
	krciAIDir      = ".krci-ai"
	agentsDir      = "agents"
	tasksDir       = "tasks"
	templatesDir   = "templates"
	dataDir        = "data"
	configDir      = "config"
	embeddedPrefix = "assets/framework/core"
)

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
	krciPath := filepath.Join(i.targetDir, krciAIDir)

	// Create main .krci-ai directory
	if err := i.createDirectory(krciPath); err != nil {
		return fmt.Errorf("failed to create .krci-ai directory: %w", err)
	}

	// Create subdirectories
	subdirs := []string{agentsDir, tasksDir, templatesDir, dataDir, configDir}
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
	krciPath := filepath.Join(i.targetDir, krciAIDir)

	// Check if main directory exists
	if _, err := os.Stat(krciPath); os.IsNotExist(err) {
		return false
	}

	// Check if required subdirectories exist
	requiredDirs := []string{agentsDir, tasksDir, templatesDir, dataDir}
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
	return filepath.Join(i.targetDir, krciAIDir)
}

// GetAgentsPath returns the path to the agents directory
func (i *Installer) GetAgentsPath() string {
	return filepath.Join(i.targetDir, krciAIDir, agentsDir)
}

// GetTasksPath returns the path to the tasks directory
func (i *Installer) GetTasksPath() string {
	return filepath.Join(i.targetDir, krciAIDir, tasksDir)
}

// GetTemplatesPath returns the path to the templates directory
func (i *Installer) GetTemplatesPath() string {
	return filepath.Join(i.targetDir, krciAIDir, templatesDir)
}

// GetDataPath returns the path to the data directory
func (i *Installer) GetDataPath() string {
	return filepath.Join(i.targetDir, krciAIDir, dataDir)
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
	taskFiles, err := filepath.Glob(filepath.Join(tasksPath, "*.yaml"))
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
