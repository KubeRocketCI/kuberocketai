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
package validation

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// detectOrphanedFiles implements reverse dependency lookup to find unused files
func (a *FrameworkAnalyzer) detectOrphanedFiles(frameworkDir string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	// Step 1: Collect all template and data files
	allFiles := make(map[string]bool)

	templatesDir := filepath.Join(frameworkDir, "templates")
	if _, err := os.Stat(templatesDir); err == nil {
		templateFiles, err := filepath.Glob(filepath.Join(templatesDir, "*.md"))
		if err != nil {
			return nil, err
		}
		for _, file := range templateFiles {
			relPath, _ := filepath.Rel(a.baseDir, file)
			allFiles[relPath] = false // false means not referenced
		}
	}

	dataDir := filepath.Join(frameworkDir, "data")
	if _, err := os.Stat(dataDir); err == nil {
		dataFiles, err := filepath.Glob(filepath.Join(dataDir, "*.*"))
		if err != nil {
			return nil, err
		}
		for _, file := range dataFiles {
			relPath, _ := filepath.Rel(a.baseDir, file)
			allFiles[relPath] = false
		}
	}

	// Step 2: Scan for referenced files
	referencedFiles := make(map[string]bool)

	if err := a.scanAgentReferences(frameworkDir, referencedFiles); err != nil {
		return nil, err
	}

	if err := a.scanTaskReferences(frameworkDir, referencedFiles); err != nil {
		return nil, err
	}

	// Step 4: Find orphaned files
	for filePath := range allFiles {
		if !referencedFiles[filePath] {
			issues = append(issues, ValidationIssue{
				Type:        "orphaned_file",
				Severity:    SeverityWarning,
				File:        filePath,
				Line:        0,
				Message:     "Unused file in framework - not referenced by any component",
				FixGuidance: "Remove unused file or add reference from appropriate task",
			})
		}
	}

	return issues, nil
}

// detectCircularDependencies detects circular dependency chains
func (a *FrameworkAnalyzer) detectCircularDependencies(frameworkDir string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	// Build dependency graph
	dependencyGraph := make(map[string][]string)

	if err := a.buildAgentTaskDependencies(frameworkDir, dependencyGraph); err != nil {
		return nil, err
	}

	if err := a.buildTaskTemplateDependencies(frameworkDir, dependencyGraph); err != nil {
		return nil, err
	}

	// Step 3: Detect cycles using DFS
	visited := make(map[string]bool)
	visiting := make(map[string]bool)

	for node := range dependencyGraph {
		if !visited[node] {
			if a.hasCycle(node, dependencyGraph, visited, visiting) {
				issues = append(issues, ValidationIssue{
					Type:        "circular_dependency",
					Severity:    SeverityWarning,
					File:        node,
					Line:        0,
					Message:     "Circular dependency detected in component references",
					FixGuidance: "Review dependency chain and remove circular references",
				})
			}
		}
	}

	return issues, nil
}

// hasCycle implements DFS cycle detection
func (a *FrameworkAnalyzer) hasCycle(node string, graph map[string][]string, visited, visiting map[string]bool) bool {
	if visiting[node] {
		return true // Back edge found - cycle detected
	}
	if visited[node] {
		return false // Already processed
	}

	visiting[node] = true
	for _, neighbor := range graph[node] {
		if a.hasCycle(neighbor, graph, visited, visiting) {
			return true
		}
	}
	visiting[node] = false
	visited[node] = true
	return false
}

// extractYAMLTasks extracts task references from agent YAML files
func (a *FrameworkAnalyzer) extractYAMLTasks(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var agent struct {
		Agent struct {
			Tasks []string `yaml:"tasks"`
		} `yaml:"agent"`
	}

	if err := yaml.Unmarshal(content, &agent); err != nil {
		return nil, err
	}

	// Filter for .krci-ai task references (both standard and local)
	var tasks []string
	for _, task := range agent.Agent.Tasks {
		if strings.HasPrefix(task, "./.krci-ai/tasks/") || strings.HasPrefix(task, "./.krci-ai/local/tasks/") {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

// extractMarkdownLinks extracts internal framework links from markdown files
func (a *FrameworkAnalyzer) extractMarkdownLinks(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Regex for internal markdown links: [text](./.krci-ai/path/file.ext)
	linkRegex := regexp.MustCompile(`\[([^\]]+)\]\((\./\.krci-ai/(?:tasks|templates|data)/[^)]+\.(md|yaml|yml|json))\)`)

	var links []string
	matches := linkRegex.FindAllStringSubmatch(string(content), -1)
	for _, match := range matches {
		if len(match) > 2 {
			linkPath := match[2]
			// Skip external links
			if !strings.HasPrefix(linkPath, "http://") && !strings.HasPrefix(linkPath, "https://") {
				links = append(links, linkPath)
			}
		}
	}

	return links, nil
}

// scanAgentReferences scans agent files for task references
//
//nolint:dupl // Similar structure but different purpose than scanTaskReferences
func (a *FrameworkAnalyzer) scanAgentReferences(frameworkDir string, referencedFiles map[string]bool) error {
	return a.processDirectoryFiles(frameworkDir, "agents", "*.y*ml", func(file string) ([]string, error) {
		return a.extractYAMLTasks(file)
	}, func(ref string) {
		cleanPath := strings.TrimPrefix(ref, "./")
		taskPath := filepath.Join(a.baseDir, cleanPath)
		if _, err := os.Stat(taskPath); err == nil {
			relPath, _ := filepath.Rel(a.baseDir, taskPath)
			referencedFiles[relPath] = true
		}
	})
}

// scanTaskReferences scans task files for template/data references (both standard and local tasks)
//
//nolint:dupl // Similar structure but different purpose than scanAgentReferences
func (a *FrameworkAnalyzer) scanTaskReferences(frameworkDir string, referencedFiles map[string]bool) error {
	// Scan standard tasks directory
	if err := a.processDirectoryFiles(frameworkDir, "tasks", "*.md", func(file string) ([]string, error) {
		return a.extractMarkdownLinks(file)
	}, func(ref string) {
		cleanPath := strings.TrimPrefix(ref, "./")
		absolutePath := filepath.Join(a.baseDir, cleanPath)
		if _, err := os.Stat(absolutePath); err == nil {
			relPath, _ := filepath.Rel(a.baseDir, absolutePath)
			referencedFiles[relPath] = true
		}
	}); err != nil {
		return err
	}

	// Scan local tasks directory
	return a.processDirectoryFiles(frameworkDir, "local/tasks", "*.md", func(file string) ([]string, error) {
		return a.extractMarkdownLinks(file)
	}, func(ref string) {
		cleanPath := strings.TrimPrefix(ref, "./")
		absolutePath := filepath.Join(a.baseDir, cleanPath)
		if _, err := os.Stat(absolutePath); err == nil {
			relPath, _ := filepath.Rel(a.baseDir, absolutePath)
			referencedFiles[relPath] = true
		}
	})
}

// buildAgentTaskDependencies builds agent → task dependencies for dependency graph
//
//nolint:dupl // Similar structure but different purpose than buildTaskTemplateDependencies
func (a *FrameworkAnalyzer) buildAgentTaskDependencies(frameworkDir string, dependencyGraph map[string][]string) error {
	dir := filepath.Join(frameworkDir, "agents")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil
	}

	files, err := filepath.Glob(filepath.Join(dir, "*.y*ml"))
	if err != nil {
		return err
	}

	for _, file := range files {
		relPath, _ := filepath.Rel(a.baseDir, file)
		refs, err := a.extractYAMLTasks(file)
		if err != nil {
			continue
		}

		for _, ref := range refs {
			cleanPath := strings.TrimPrefix(ref, "./")
			taskPath := filepath.Join(a.baseDir, cleanPath)
			if _, err := os.Stat(taskPath); err == nil {
				relTaskPath, _ := filepath.Rel(a.baseDir, taskPath)
				dependencyGraph[relPath] = append(dependencyGraph[relPath], relTaskPath)
			}
		}
	}
	return nil
}

// buildTaskTemplateDependencies builds task → template/data dependencies for dependency graph (both standard and local tasks)
//
//nolint:dupl // Similar structure but different purpose than buildAgentTaskDependencies
func (a *FrameworkAnalyzer) buildTaskTemplateDependencies(frameworkDir string, dependencyGraph map[string][]string) error {
	// Process standard tasks directory
	if err := a.processTaskDependenciesInDirectory(frameworkDir, "tasks", dependencyGraph); err != nil {
		return err
	}

	// Process local tasks directory
	return a.processTaskDependenciesInDirectory(frameworkDir, "local/tasks", dependencyGraph)
}

// processTaskDependenciesInDirectory processes task dependencies in a specific directory
func (a *FrameworkAnalyzer) processTaskDependenciesInDirectory(frameworkDir, taskDir string, dependencyGraph map[string][]string) error {
	dir := filepath.Join(frameworkDir, taskDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil
	}

	files, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return err
	}

	for _, file := range files {
		relPath, _ := filepath.Rel(a.baseDir, file)
		refs, err := a.extractMarkdownLinks(file)
		if err != nil {
			continue
		}

		for _, ref := range refs {
			cleanPath := strings.TrimPrefix(ref, "./")
			absolutePath := filepath.Join(a.baseDir, cleanPath)
			if _, err := os.Stat(absolutePath); err == nil {
				relLinkPath, _ := filepath.Rel(a.baseDir, absolutePath)
				dependencyGraph[relPath] = append(dependencyGraph[relPath], relLinkPath)
			}
		}
	}
	return nil
}

// processDirectoryFiles is a generic helper to process files in a directory
func (a *FrameworkAnalyzer) processDirectoryFiles(frameworkDir, subDir, pattern string, extractRefs func(string) ([]string, error), processRef func(string)) error {
	dir := filepath.Join(frameworkDir, subDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil
	}

	files, err := filepath.Glob(filepath.Join(dir, pattern))
	if err != nil {
		return err
	}

	for _, file := range files {
		refs, err := extractRefs(file)
		if err != nil {
			continue
		}

		for _, ref := range refs {
			processRef(ref)
		}
	}
	return nil
}
