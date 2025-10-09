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
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/KubeRocketCI/kuberocketai/internal/processor"
	"github.com/KubeRocketCI/kuberocketai/internal/utils"
)

// FileSystem interface abstracts file system operations for both OS and embedded filesystems
type FileSystem interface {
	WalkDir(root string, fn fs.WalkDirFunc) error
	ReadFile(name string) ([]byte, error)
}

// OSFileSystem implements FileSystem for the operating system filesystem
type OSFileSystem struct{}

func (OSFileSystem) WalkDir(root string, fn fs.WalkDirFunc) error {
	return filepath.WalkDir(root, fn)
}

func (OSFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// EmbeddedFileSystem implements FileSystem for embedded filesystems
type EmbeddedFileSystem struct {
	fs embed.FS
}

func (efs EmbeddedFileSystem) WalkDir(root string, fn fs.WalkDirFunc) error {
	return fs.WalkDir(efs.fs, filepath.ToSlash(root), fn)
}

func (efs EmbeddedFileSystem) ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(efs.fs, filepath.ToSlash(name))
}

// Agent represents basic information about an agent
type Agent struct {
	Name        string
	Description string
	Role        string
	Goal        string
	Icon        string
	Tasks       []Task
	FilePath    string
	ShortName   string
}

func (a *Agent) GetAllTasksPaths() []string {
	tasksPaths := make([]string, 0, len(a.Tasks))
	for _, task := range a.Tasks {
		tasksPaths = append(tasksPaths, task.Path)
	}

	return tasksPaths
}

func (a *Agent) GetAllTemplatesPaths() []string {
	templatesPaths := make([]string, 0, len(a.Tasks))
	for _, task := range a.Tasks {
		for _, template := range task.Dependencies.Templates {
			templatesPaths = append(templatesPaths, template.Path)
		}
	}

	return utils.DeduplicateStrings(templatesPaths)
}

func (a *Agent) GetAllDataFilesPaths() []string {
	dataFilesPaths := make([]string, 0, len(a.Tasks))
	for _, task := range a.Tasks {
		for _, dataFile := range task.Dependencies.DataFiles {
			dataFilesPaths = append(dataFilesPaths, dataFile.Path)
		}
	}

	return utils.DeduplicateStrings(dataFilesPaths)
}

func (a *Agent) GetAllReferencedTasksPaths() []string {
	tasksPaths := make([]string, 0, len(a.Tasks))
	for _, task := range a.Tasks {
		for _, taskRef := range task.Dependencies.Tasks {
			tasksPaths = append(tasksPaths, taskRef.Path)
		}
	}

	return utils.DeduplicateStrings(tasksPaths)
}

type Task struct {
	Path         string
	Name         string
	Dependencies TaskDependencies
}

type TaskDependencies struct {
	Templates  []Template
	DataFiles  []DataFile
	Tasks      []TaskRef
	McpServers []string
}

type Template struct {
	Path string
	Name string
}

type DataFile struct {
	Path string
	Name string
}

type TaskRef struct {
	Path string
	Name string
}

// Discovery handles discovery and parsing of framework assets from any source
type Discovery struct {
	fs           FileSystem
	frameworkDir string
}

// NewDiscovery creates a new asset discovery service for filesystem assets
func NewDiscovery(frameworkDir string) *Discovery {
	return &Discovery{
		fs:           OSFileSystem{},
		frameworkDir: frameworkDir,
	}
}

// NewEmbeddedDiscovery creates a new asset discovery service for embedded assets
func NewEmbeddedDiscovery(embeddedFS embed.FS, frameworkDir string) *Discovery {
	return &Discovery{
		fs:           EmbeddedFileSystem{fs: embeddedFS},
		frameworkDir: frameworkDir,
	}
}

func (d *Discovery) GetAgents(ctx context.Context) ([]Agent, error) {
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan string)

	g.Go(func() error {
		defer close(paths)
		if err := d.fs.WalkDir(GetAgentsPath(d.frameworkDir), func(agentPath string, dirEntry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !dirEntry.Type().IsRegular() || filepath.Ext(agentPath) != ".yaml" {
				return nil
			}
			select {
			case paths <- filepath.Clean(agentPath):
			case <-ctx.Done():
				return ctx.Err()
			}

			return nil
		}); err != nil {
			return fmt.Errorf("failed to read directory: %w", err)
		}

		return nil
	})

	results := make(chan Agent)
	const workers = 20
	for range workers {
		g.Go(func() error {
			for agentPath := range paths {
				agent, err := processor.UnmarshalAgentFileFromFS(d.fs, agentPath)
				if err != nil {
					return fmt.Errorf("failed to unmarshal agent file: %w", err)
				}

				tasks, err := d.getAgentTasks(ctx, agent)
				if err != nil {
					return err
				}

				select {
				case results <- MakeAgent(agentPath, agent, tasks):
				case <-ctx.Done():
					return ctx.Err()
				}
			}

			return nil
		})
	}

	go func() {
		_ = g.Wait()
		close(results)
	}()

	agents := make([]Agent, 0)
	for result := range results {
		agents = append(agents, result)
	}

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("failed to list agents: %w", err)
	}

	return agents, nil
}

// GetAgent returns an agent by short name.
// TODO: Method can be optimized by reading only one agent instead of all agents.
func (d *Discovery) GetAgent(ctx context.Context, shortName string) (*Agent, error) {
	agents, err := d.GetAgents(ctx)
	if err != nil {
		return nil, err
	}

	for _, agent := range agents {
		if agent.ShortName == shortName {
			return &agent, nil
		}
	}
	return nil, fmt.Errorf("agent %s not found", shortName)
}

func (d *Discovery) GetAgentsByNames(ctx context.Context, names []string) ([]Agent, error) {
	agents, err := d.GetAgents(ctx)
	if err != nil {
		return nil, err
	}

	selectedAgents := make([]Agent, 0, len(names))
	for _, name := range names {
		found := false
		for _, agent := range agents {
			if agent.ShortName == name {
				selectedAgents = append(selectedAgents, agent)
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("agent %s not found", name)
		}
	}

	return selectedAgents, nil
}

// ReadFile reads a file using the discovery's filesystem
func (d *Discovery) ReadFile(filePath string) ([]byte, error) {
	return d.fs.ReadFile(filePath)
}

func (d *Discovery) getAgentTasks(ctx context.Context, rawAgent *processor.AgentYamlRepresentation) ([]Task, error) {
	g, ctx := errgroup.WithContext(ctx)
	tasksPaths := make(chan string)

	g.Go(func() error {
		defer close(tasksPaths)
		for _, taskPath := range rawAgent.Agent.Tasks {
			select {
			case tasksPaths <- taskPath:
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		return nil
	})

	workers := 10
	results := make(chan Task)
	for range workers {
		g.Go(func() error {
			for taskPath := range tasksPaths {
				task, err := d.getAgentTask(taskPath)
				if err != nil {
					return err
				}

				select {
				case results <- task:
				case <-ctx.Done():
					return ctx.Err()
				}
			}

			return nil
		})
	}

	go func() {
		_ = g.Wait()
		close(results)
	}()

	tasks := make([]Task, 0, len(rawAgent.Agent.Tasks))
	for result := range results {
		tasks = append(tasks, result)
	}

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("failed to get agent tasks for agent %s: %w", rawAgent.Agent.Identity.ID, err)
	}

	return tasks, nil
}

func (d *Discovery) getAgentTask(taskRef string) (Task, error) {
	taskName := strings.TrimPrefix(taskRef, "./"+KrciAIDir+"/tasks/")
	taskPath := filepath.Join(GetTasksPath(d.frameworkDir), taskName)
	taskDependencies, err := processor.UnmarshalTaskDependenciesFileFromFS(d.fs, taskPath)
	if err != nil {
		return Task{}, err
	}

	return MakeTask(d.frameworkDir, taskPath, *taskDependencies), nil
}

func GetAgentsPath(frameworkDir string) string {
	return filepath.Join(frameworkDir, agentsDir)
}

func GetTasksPath(frameworkDir string) string {
	return filepath.Join(frameworkDir, TasksDir)
}

func GetTemplatesPath(frameworkDir string) string {
	return filepath.Join(frameworkDir, TemplatesDir)
}

func GetDataPath(frameworkDir string) string {
	return filepath.Join(frameworkDir, DataDir)
}

func GetKrciPath(projectDir string) string {
	return filepath.Join(projectDir, KrciAIDir)
}

func MakeAgent(path string, representation *processor.AgentYamlRepresentation, tasks []Task) Agent {
	return Agent{
		Name:        representation.Agent.Identity.Name,
		Description: representation.Agent.Identity.Description,
		Role:        representation.Agent.Identity.Role,
		Goal:        representation.Agent.Identity.Goal,
		Icon:        representation.Agent.Identity.Icon,
		FilePath:    path,
		ShortName:   strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Tasks:       tasks,
	}
}

func MakeTask(basePath string, taskPath string, dependencies processor.TaskDependenciesYamlRepresentation) Task {
	return Task{
		Path:         taskPath,
		Name:         strings.TrimSuffix(filepath.Base(taskPath), filepath.Ext(taskPath)),
		Dependencies: MakeTaskDependency(basePath, dependencies),
	}
}

func MakeTaskDependency(basePath string, dependency processor.TaskDependenciesYamlRepresentation) TaskDependencies {
	d := TaskDependencies{
		Templates:  make([]Template, 0, len(dependency.Dependencies.Templates)),
		DataFiles:  make([]DataFile, 0, len(dependency.Dependencies.DataFiles)),
		Tasks:      make([]TaskRef, 0, len(dependency.Dependencies.Tasks)),
		McpServers: dependency.Dependencies.McpServers,
	}

	for _, template := range dependency.Dependencies.Templates {
		d.Templates = append(d.Templates, Template{
			Path: filepath.Join(GetTemplatesPath(basePath), template),
			Name: template,
		})
	}

	for _, dataFile := range dependency.Dependencies.DataFiles {
		d.DataFiles = append(d.DataFiles, DataFile{
			Path: filepath.Join(GetDataPath(basePath), dataFile),
			Name: dataFile,
		})
	}

	for _, task := range dependency.Dependencies.Tasks {
		d.Tasks = append(d.Tasks, TaskRef{
			Path: filepath.Join(GetTasksPath(basePath), task),
			Name: task,
		})
	}

	return d
}
