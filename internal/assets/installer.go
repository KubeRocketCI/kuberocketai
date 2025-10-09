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
	"context"
	"embed"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/KubeRocketCI/kuberocketai/internal/utils"
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
	EmbeddedPrefix = "assets/framework/core"

	// File extensions
	mdExtension = ".md"

	// GitHub tools configuration
	GitHubToolsList = "['changes', 'codebase', 'editFiles', 'fetch', 'findTestFiles', 'githubRepo', 'problems', 'runCommands', 'search', 'searchResults', 'terminalLastCommand', 'usages']"

	DefaultAgentIcon = "🤖"

	// File permissions
	DirectoryPermissions = 0755 // Read, write, execute for owner; read, execute for group and others
	FilePermissions      = 0644 // Read, write for owner; read for group and others
)

type IstallerDiscovery interface {
	GetAgents(ctx context.Context) ([]Agent, error)
	GetAgentsByNames(ctx context.Context, names []string) ([]Agent, error)
}

// Installer handles installation of framework assets
type Installer struct {
	projectDir     string
	krciPath       string
	embeddedAssets EmbeddedFileSystem
	discovery      IstallerDiscovery
}

// NewInstaller creates a new asset installer
func NewInstaller(projectDir string, embeddedAssets embed.FS, discovery IstallerDiscovery) *Installer {
	return &Installer{
		projectDir:     projectDir,
		krciPath:       GetKrciPath(projectDir),
		embeddedAssets: EmbeddedFileSystem{fs: embeddedAssets},
		discovery:      discovery,
	}
}

// Install installs framework assets to the target directory
func (i *Installer) Install() error {
	agents, err := i.discovery.GetAgents(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get agents: %w", err)
	}

	prefix := filepath.Clean(EmbeddedPrefix + "/")
	filesFilter := make(map[string]struct{})
	for _, agent := range agents {
		i.addAgentDependencies(agent, filesFilter, prefix)
	}

	return i.copyEmbeddedFiles(context.TODO(), filesFilter)
}

// InstallSelective installs only specified agents and their dependencies using existing bundle logic
func (i *Installer) InstallSelective(agentNames []string) error {
	if len(agentNames) == 0 {
		return fmt.Errorf("no agents specified")
	}

	agents, err := i.discovery.GetAgentsByNames(context.Background(), agentNames)
	if err != nil {
		return fmt.Errorf("failed to get agents: %w", err)
	}

	if err := i.validateAgentsFound(agents, agentNames); err != nil {
		return err
	}

	prefix := filepath.Clean(EmbeddedPrefix + "/")
	filesFilter := make(map[string]struct{})
	for _, agent := range agents {
		i.addAgentDependencies(agent, filesFilter, prefix)
	}

	return i.copyEmbeddedFiles(context.TODO(), filesFilter)
}

// validateAgentsFound checks if all requested agents were found
func (i *Installer) validateAgentsFound(agents []Agent, agentNames []string) error {
	if len(agents) == len(agentNames) {
		return nil
	}

	foundNames := make(map[string]struct{})
	for _, agent := range agents {
		foundNames[agent.Name] = struct{}{}
	}

	var notFound []string
	for _, name := range agentNames {
		if _, exists := foundNames[name]; !exists {
			notFound = append(notFound, name)
		}
	}
	return fmt.Errorf("agents not found: %s", strings.Join(notFound, ", "))
}

// addAgentDependencies adds agent and its dependencies to the files filter
func (i *Installer) addAgentDependencies(agent Agent, filesFilter map[string]struct{}, prefix string) {
	trimPrefix := func(path string) string {
		return strings.TrimPrefix(path, prefix)
	}

	filesFilter[trimPrefix(agent.FilePath)] = struct{}{}
	maps.Insert(filesFilter, maps.All(utils.MapSliceToSet(agent.GetAllDataFilesPaths(), trimPrefix)))
	maps.Insert(filesFilter, maps.All(utils.MapSliceToSet(agent.GetAllTemplatesPaths(), trimPrefix)))
	maps.Insert(filesFilter, maps.All(utils.MapSliceToSet(agent.GetAllTasksPaths(), trimPrefix)))
	maps.Insert(filesFilter, maps.All(utils.MapSliceToSet(agent.GetAllReferencedTasksPaths(), trimPrefix)))
}

// copyEmbeddedFiles copies embedded files to the target directory.
func (i *Installer) copyEmbeddedFiles(ctx context.Context, filepaths map[string]struct{}) error {
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan string)

	g.Go(func() error {
		defer close(paths)
		for path := range filepaths {
			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		return nil
	})

	// Configure worker count based on workload size
	fileCount := len(filepaths)
	workers := min(max(fileCount/10, 1), 10) // 1-10 workers based on file count
	for range workers {
		g.Go(func() error {
			for path := range paths {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
				}

				targetPath := filepath.Join(i.krciPath, path)

				if err := os.MkdirAll(filepath.Dir(targetPath), DirectoryPermissions); err != nil {
					return fmt.Errorf("failed to create parent directory for %s: %w", targetPath, err)
				}

				embeddedPath := filepath.Join(EmbeddedPrefix, path)
				data, err := i.embeddedAssets.ReadFile(embeddedPath)
				if err != nil {
					return fmt.Errorf("failed to read embedded file %s (target: %s): %w", embeddedPath, targetPath, err)
				}

				if err := os.WriteFile(targetPath, data, FilePermissions); err != nil {
					return fmt.Errorf("failed to write file to %s (source: %s): %w", targetPath, embeddedPath, err)
				}
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to copy embedded files: %w", err)
	}

	return nil
}

// createDirectory creates a directory if it doesn't exist
func (i *Installer) createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, DirectoryPermissions)
	}
	return nil
}

// IsInstalled checks if the framework is properly installed in the target directory
func (i *Installer) IsInstalled() bool {
	// Check if main directory exists
	if _, err := os.Stat(i.krciPath); os.IsNotExist(err) {
		return false
	}

	// Check if required subdirectories exist
	requiredDirs := []string{agentsDir, TasksDir}
	for _, dir := range requiredDirs {
		dirPath := filepath.Join(i.krciPath, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			return false
		}
	}

	return true
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

func (i *Installer) GetFrameworkPath() string {
	return i.krciPath
}
