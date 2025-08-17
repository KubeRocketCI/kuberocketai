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
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/KubeRocketCI/kuberocketai/internal/utils"
)

// Directory constants to avoid hardcoded paths
const (
	agentsDir      = "agents"
	tasksDir       = "tasks"
	templatesDir   = "templates"
	dataDir        = "data"
	localDir       = "local"
	embeddedPrefix = "assets/framework/core"
)

// Path pattern constants
const (
	agentFilePattern     = "%s/%s/%s.yaml"
	taskReferencePrefix  = "./.krci-ai/"
	localTasksPrefix     = "./.krci-ai/local/tasks/"
	standardTasksPrefix  = "./.krci-ai/tasks/"
	templatesLinkPattern = "/templates/"
	dataLinkPattern      = "/data/"
)

// Severity levels for validation issues
type Severity int

const (
	SeverityInfo Severity = iota
	SeverityWarning
	SeverityError
	SeverityCritical
)

func (s Severity) String() string {
	switch s {
	case SeverityInfo:
		return "INFO"
	case SeverityWarning:
		return "WARNING"
	case SeverityCritical:
		return "CRITICAL"
	default:
		return "ERROR"
	}
}

// ValidationIssue represents a single validation issue
type ValidationIssue struct {
	Type        string
	Severity    Severity
	File        string
	Line        int
	Message     string
	FixGuidance string
}

// FrameworkAnalyzer provides comprehensive framework validation
type FrameworkAnalyzer struct {
	baseDir     string
	cache       map[string]time.Time
	fileCache   map[string][]byte
	resultCache map[string][]ValidationIssue
}

// NewFrameworkAnalyzer creates a new framework analyzer
func NewFrameworkAnalyzer(baseDir string) *FrameworkAnalyzer {
	return &FrameworkAnalyzer{
		baseDir:     baseDir,
		cache:       make(map[string]time.Time),
		fileCache:   make(map[string][]byte),
		resultCache: make(map[string][]ValidationIssue),
	}
}

// AnalyzeFramework performs comprehensive framework analysis
func (a *FrameworkAnalyzer) AnalyzeFramework() ([]ValidationIssue, *FrameworkInsights, error) {
	var issues []ValidationIssue

	frameworkDir := filepath.Join(a.baseDir, ".krci-ai")
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("no .krci-ai directory found")
	}

	// Phase 1: Critical issue detection
	criticalIssues, err := a.detectCriticalIssues(frameworkDir)
	if err != nil {
		return nil, nil, fmt.Errorf("critical issue detection failed: %w", err)
	}
	issues = append(issues, criticalIssues...)

	// Phase 2: Warning detection
	warningIssues, err := a.detectWarningIssues(frameworkDir)
	if err != nil {
		return nil, nil, fmt.Errorf("warning detection failed: %w", err)
	}
	issues = append(issues, warningIssues...)

	// Phase 3: Generate framework insights
	insights, err := a.generateFrameworkInsights(frameworkDir)
	if err != nil {
		return nil, nil, fmt.Errorf("framework insights generation failed: %w", err)
	}

	return issues, insights, nil
}

// detectCriticalIssues detects critical validation issues
func (a *FrameworkAnalyzer) detectCriticalIssues(frameworkDir string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	// 1. Detect broken internal links
	brokenLinks, err := a.detectBrokenInternalLinks(frameworkDir)
	if err != nil {
		return nil, err
	}
	issues = append(issues, brokenLinks...)

	// 2. Detect missing task files referenced in agents
	missingTasks, err := a.detectMissingTaskFiles(frameworkDir)
	if err != nil {
		return nil, err
	}
	issues = append(issues, missingTasks...)

	// 3. Detect architecture violations
	archViolations, err := a.detectArchitectureViolations(frameworkDir)
	if err != nil {
		return nil, err
	}
	issues = append(issues, archViolations...)

	// 4. Detect local directory violations
	localDirViolations, err := a.detectLocalDirectoryViolations(frameworkDir)
	if err != nil {
		return nil, err
	}
	issues = append(issues, localDirViolations...)

	// 5. Detect invalid file formats
	formatIssues, err := a.detectFormatIssues(frameworkDir)
	if err != nil {
		return nil, err
	}
	issues = append(issues, formatIssues...)

	return issues, nil
}

// detectBrokenInternalLinks detects broken internal framework links
func (a *FrameworkAnalyzer) detectBrokenInternalLinks(frameworkDir string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	// Regex for internal markdown links: [text](./.krci-ai/path/file.ext)
	linkRegex := regexp.MustCompile(`\[([^\]]+)\]\((\./\.krci-ai/(?:tasks|templates|data)/[^)]+\.(md|yaml|yml|json))\)`)

	// Find all markdown files
	markdownFiles, err := a.findMarkdownFiles(frameworkDir)
	if err != nil {
		return nil, err
	}

	for _, file := range markdownFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		lines := strings.Split(string(content), "\n")
		for lineNum, line := range lines {
			matches := linkRegex.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				if len(match) > 2 {
					linkPath := match[2]
					// Skip external links
					if strings.HasPrefix(linkPath, "http://") || strings.HasPrefix(linkPath, "https://") {
						continue
					}

					// Convert to absolute path
					cleanPath := strings.TrimPrefix(linkPath, "./")
					absolutePath := filepath.Join(a.baseDir, cleanPath)

					if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
						relFile, _ := filepath.Rel(a.baseDir, file)
						issues = append(issues, ValidationIssue{
							Type:        "broken_link",
							Severity:    SeverityCritical,
							File:        relFile,
							Line:        lineNum + 1,
							Message:     fmt.Sprintf("Broken link to framework file: %s", linkPath),
							FixGuidance: "Create the missing file or update the reference path",
						})
					}
				}
			}
		}
	}

	return issues, nil
}

// detectMissingTaskFiles detects missing task files referenced in agent YAML
func (a *FrameworkAnalyzer) detectMissingTaskFiles(frameworkDir string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	agentsDir := filepath.Join(frameworkDir, "agents")
	if _, err := os.Stat(agentsDir); os.IsNotExist(err) {
		return issues, nil
	}

	agentFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.yaml"))
	if err != nil {
		return nil, err
	}
	ymlFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.yml"))
	if err != nil {
		return nil, err
	}
	agentFiles = append(agentFiles, ymlFiles...)

	for _, agentFile := range agentFiles {
		content, err := os.ReadFile(agentFile)
		if err != nil {
			continue
		}

		var agent struct {
			Agent struct {
				Tasks []string `yaml:"tasks"`
			} `yaml:"agent"`
		}

		if err := yaml.Unmarshal(content, &agent); err != nil {
			continue // Will be caught by format validation
		}

		for _, taskPath := range agent.Agent.Tasks {
			if !strings.HasPrefix(taskPath, "./.krci-ai/tasks/") {
				continue
			}

			cleanPath := strings.TrimPrefix(taskPath, "./")
			absolutePath := filepath.Join(a.baseDir, cleanPath)

			if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
				relFile, _ := filepath.Rel(a.baseDir, agentFile)
				issues = append(issues, ValidationIssue{
					Type:        "missing_task",
					Severity:    SeverityCritical,
					File:        relFile,
					Line:        0,
					Message:     fmt.Sprintf("Missing task file: %s", taskPath),
					FixGuidance: "Create the missing task file or remove the reference",
				})
			}
		}
	}

	return issues, nil
}

// detectArchitectureViolations detects violations of component separation
func (a *FrameworkAnalyzer) detectArchitectureViolations(frameworkDir string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	templatesDir := filepath.Join(frameworkDir, "templates")
	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		return issues, nil
	}

	templateFiles, err := filepath.Glob(filepath.Join(templatesDir, "*.md"))
	if err != nil {
		return nil, err
	}

	// Templates should not reference data files directly
	dataLinkRegex := regexp.MustCompile(`\[([^\]]+)\]\((\./\.krci-ai/data/[^)]+)\)`)

	for _, templateFile := range templateFiles {
		content, err := os.ReadFile(templateFile)
		if err != nil {
			continue
		}

		lines := strings.Split(string(content), "\n")
		for lineNum, line := range lines {
			if dataLinkRegex.MatchString(line) {
				relFile, _ := filepath.Rel(a.baseDir, templateFile)
				issues = append(issues, ValidationIssue{
					Type:        "architecture_violation",
					Severity:    SeverityCritical,
					File:        relFile,
					Line:        lineNum + 1,
					Message:     "Template contains data reference - violates component separation",
					FixGuidance: "Move data reference from template to calling task",
				})
			}
		}
	}

	return issues, nil
}

// detectLocalDirectoryViolations detects invalid local directory structure
func (a *FrameworkAnalyzer) detectLocalDirectoryViolations(frameworkDir string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	localDir := filepath.Join(frameworkDir, "local")
	if _, err := os.Stat(localDir); os.IsNotExist(err) {
		// No local directory, no violations
		return issues, nil
	}

	// Check if local directory exists and get its contents
	entries, err := os.ReadDir(localDir)
	if err != nil {
		return nil, err
	}

	// Allowed directories in .krci-ai/local/
	allowedDirs := map[string]bool{
		"tasks":     true,
		"templates": true,
		"data":      true,
	}

	for _, entry := range entries {
		if !allowedDirs[entry.Name()] {
			relFile, _ := filepath.Rel(a.baseDir, filepath.Join(localDir, entry.Name()))
			issues = append(issues, ValidationIssue{
				Type:        "local_directory_violation",
				Severity:    SeverityCritical,
				File:        relFile,
				Line:        0,
				Message:     fmt.Sprintf("Invalid directory '%s' in .krci-ai/local/ - only 'tasks', 'templates', and 'data' are allowed", entry.Name()),
				FixGuidance: "Remove invalid directory or move contents to allowed directories (tasks/, templates/, data/)",
			})
		}
	}

	return issues, nil
}

// detectFormatIssues detects invalid file formats
func (a *FrameworkAnalyzer) detectFormatIssues(frameworkDir string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	// Check YAML files in agents directory
	agentsDir := filepath.Join(frameworkDir, "agents")
	if _, err := os.Stat(agentsDir); err == nil {
		agentFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.y*ml"))
		if err != nil {
			return nil, err
		}

		for _, agentFile := range agentFiles {
			content, err := os.ReadFile(agentFile)
			if err != nil {
				continue
			}

			var temp interface{}
			if err := yaml.Unmarshal(content, &temp); err != nil {
				relFile, _ := filepath.Rel(a.baseDir, agentFile)
				issues = append(issues, ValidationIssue{
					Type:        "invalid_format",
					Severity:    SeverityCritical,
					File:        relFile,
					Line:        0,
					Message:     fmt.Sprintf("Invalid YAML syntax: %s", err.Error()),
					FixGuidance: "Fix YAML syntax errors",
				})
			}
		}
	}

	return issues, nil
}

// detectWarningIssues detects warning-level issues
func (a *FrameworkAnalyzer) detectWarningIssues(frameworkDir string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	// 1. Detect orphaned files
	orphanedFiles, err := a.detectOrphanedFiles(frameworkDir)
	if err != nil {
		return nil, err
	}
	issues = append(issues, orphanedFiles...)

	// 2. Detect circular dependencies
	circularDeps, err := a.detectCircularDependencies(frameworkDir)
	if err != nil {
		return nil, err
	}
	issues = append(issues, circularDeps...)

	return issues, nil
}

// findMarkdownFiles finds all markdown files in the framework
func (a *FrameworkAnalyzer) findMarkdownFiles(frameworkDir string) ([]string, error) {
	var files []string
	dirs := []string{agentsDir, tasksDir, templatesDir, dataDir}

	// Include global directories
	for _, dir := range dirs {
		dirPath := filepath.Join(frameworkDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			continue
		}

		// Recursively find all markdown files including subdirectories
		err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	// Include local directories
	localDirs := []string{"local/tasks", "local/templates", "local/data"}
	for _, dir := range localDirs {
		dirPath := filepath.Join(frameworkDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			continue
		}

		dirFiles, err := filepath.Glob(filepath.Join(dirPath, "*.md"))
		if err != nil {
			return nil, err
		}
		files = append(files, dirFiles...)
	}

	return files, nil
}

// AnalyzeEmbeddedDependencies implements EmbeddedAssetAnalyzer interface
func (a *FrameworkAnalyzer) AnalyzeEmbeddedDependencies(ctx context.Context, embeddedAssets fs.FS, agentNames []string) (*FrameworkInsights, error) {
	insights := &FrameworkInsights{
		ComponentCounts: ComponentCounts{},
		Relationships:   make([]ComponentRelationship, 0),
		UsageStatistics: UsageStatistics{},
	}

	// Step 1: Count embedded components
	if err := a.countEmbeddedComponents(embeddedAssets, &insights.ComponentCounts); err != nil {
		return nil, err
	}

	// Step 2: Analyze relationships for selected agents
	if err := a.analyzeEmbeddedRelationships(ctx, embeddedAssets, agentNames, insights); err != nil {
		return nil, fmt.Errorf("analyzing embedded relationships: %w", err)
	}

	// Step 3: Generate usage statistics
	a.generateUsageStatistics(insights)

	return insights, nil
}

// countEmbeddedComponents counts components in embedded filesystem
func (a *FrameworkAnalyzer) countEmbeddedComponents(embeddedAssets fs.FS, counts *ComponentCounts) error {
	// Count agents
	agentFiles, err := fs.Glob(embeddedAssets, embeddedPrefix+"/"+agentsDir+"/*.y*ml")
	if err != nil {
		return err
	}
	counts.Agents = len(agentFiles)

	// Count tasks (both standard and local)
	taskFiles, err := fs.Glob(embeddedAssets, embeddedPrefix+"/"+tasksDir+"/*.md")
	if err != nil {
		return err
	}
	counts.Tasks = len(taskFiles)

	localTaskFiles, err := fs.Glob(embeddedAssets, embeddedPrefix+"/local/"+tasksDir+"/*.md")
	if err == nil {
		counts.Tasks += len(localTaskFiles)
	}

	// Count templates
	templateFiles, err := fs.Glob(embeddedAssets, embeddedPrefix+"/"+templatesDir+"/*.md")
	if err != nil {
		return err
	}
	counts.Templates = len(templateFiles)

	// Count data files recursively including subdirectories
	dataCount := 0
	err = fs.WalkDir(embeddedAssets, embeddedPrefix+"/"+dataDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && (strings.HasSuffix(d.Name(), ".md") || strings.HasSuffix(d.Name(), ".txt")) {
			dataCount++
		}
		return nil
	})
	if err != nil {
		return err
	}
	counts.Data = dataCount

	return nil
}

// analyzeEmbeddedRelationships analyzes agent relationships in embedded filesystem (comprehensive version)
func (a *FrameworkAnalyzer) analyzeEmbeddedRelationships(ctx context.Context, embeddedAssets fs.FS, agentNames []string, insights *FrameworkInsights) error {
	for _, agentName := range agentNames {
		// Check for cancellation
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		relationship, err := a.buildAgentRelationship(embeddedAssets, agentName)
		if err != nil {
			// Log warning but continue with other agents for robustness
			// In a production system, you might want to use a proper logger here
			fmt.Fprintf(os.Stderr, "Warning: failed to build relationship for agent %s: %v\n", agentName, err)
			continue
		}
		insights.Relationships = append(insights.Relationships, *relationship)
	}
	return nil
}

// buildAgentRelationship builds a complete ComponentRelationship for a single agent
func (a *FrameworkAnalyzer) buildAgentRelationship(embeddedAssets fs.FS, agentName string) (*ComponentRelationship, error) {
	agentPath := fmt.Sprintf(agentFilePattern, embeddedPrefix, agentsDir, agentName)

	relationship := ComponentRelationship{
		Agent:      agentName,
		Tasks:      make([]string, 0),
		LocalTasks: make([]string, 0),
		Templates:  make([]string, 0),
		DataFiles:  make([]string, 0),
	}

	// Get tasks referenced by this agent
	taskRefs, err := a.extractEmbeddedYAMLTasks(embeddedAssets, agentPath)
	if err != nil {
		return nil, fmt.Errorf("extracting tasks for agent %s: %w", agentName, err)
	}

	if err := a.processAgentTasks(embeddedAssets, taskRefs, &relationship); err != nil {
		return nil, fmt.Errorf("processing tasks for agent %s: %w", agentName, err)
	}

	// Remove duplicates using utils package
	relationship.Templates = utils.DeduplicateStrings(relationship.Templates)
	relationship.DataFiles = utils.DeduplicateStrings(relationship.DataFiles)

	return &relationship, nil
}

// processAgentTasks processes all tasks for an agent and updates the relationship
func (a *FrameworkAnalyzer) processAgentTasks(embeddedAssets fs.FS, taskRefs []string, relationship *ComponentRelationship) error {
	// Preallocate slices for better memory efficiency
	if cap(relationship.Templates) == 0 {
		relationship.Templates = make([]string, 0, len(taskRefs)*2)
	}
	if cap(relationship.DataFiles) == 0 {
		relationship.DataFiles = make([]string, 0, len(taskRefs)*2)
	}

	for _, taskRef := range taskRefs {
		taskName := filepath.Base(taskRef)

		// Categorize as local or standard task
		if strings.Contains(taskRef, localTasksPrefix) {
			relationship.LocalTasks = append(relationship.LocalTasks, taskName)
		} else {
			relationship.Tasks = append(relationship.Tasks, taskName)
		}

		if err := a.addTaskDependencies(embeddedAssets, taskRef, relationship); err != nil {
			return fmt.Errorf("adding dependencies for task %s: %w", taskRef, err)
		}
	}
	return nil
}

// addTaskDependencies finds and adds template/data dependencies for a task
func (a *FrameworkAnalyzer) addTaskDependencies(embeddedAssets fs.FS, taskRef string, relationship *ComponentRelationship) error {
	// Convert task reference to embedded path
	embeddedTaskPath := strings.Replace(taskRef, taskReferencePrefix, embeddedPrefix+"/", 1)

	// Check if task file exists in embedded assets before analyzing
	if _, err := embeddedAssets.Open(embeddedTaskPath); err != nil {
		// Task file doesn't exist - this is not necessarily an error in embedded assets
		return nil
	}

	templates, dataFiles, err := a.getEmbeddedTaskReferences(embeddedAssets, embeddedTaskPath)
	if err != nil {
		return fmt.Errorf("getting task references from %s: %w", embeddedTaskPath, err)
	}

	relationship.Templates = append(relationship.Templates, templates...)
	relationship.DataFiles = append(relationship.DataFiles, dataFiles...)
	return nil
}

// extractEmbeddedYAMLTasks extracts task references from embedded agent YAML files
func (a *FrameworkAnalyzer) extractEmbeddedYAMLTasks(embeddedAssets fs.FS, agentPath string) ([]string, error) {
	content, err := fs.ReadFile(embeddedAssets, agentPath)
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
		if strings.HasPrefix(task, standardTasksPrefix) || strings.HasPrefix(task, localTasksPrefix) {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

// getEmbeddedTaskReferences extracts template and data references from embedded task file (minimal)
func (a *FrameworkAnalyzer) getEmbeddedTaskReferences(embeddedAssets fs.FS, taskPath string) ([]string, []string, error) {
	var templates, dataFiles []string

	// Extract markdown links with proper error handling
	links, err := a.extractEmbeddedMarkdownLinks(embeddedAssets, taskPath)
	if err != nil {
		return nil, nil, fmt.Errorf("extracting markdown links from %s: %w", taskPath, err)
	}

	for _, link := range links {
		// Preserve full relative path to maintain subdirectory structure
		relativePath := strings.TrimPrefix(link, taskReferencePrefix)
		if strings.Contains(link, templatesLinkPattern) {
			templates = append(templates, strings.TrimPrefix(relativePath, templatesDir+"/"))
		} else if strings.Contains(link, dataLinkPattern) {
			dataFiles = append(dataFiles, strings.TrimPrefix(relativePath, dataDir+"/"))
		}
	}

	return templates, dataFiles, nil
}

// extractEmbeddedMarkdownLinks extracts internal framework links from embedded markdown files
func (a *FrameworkAnalyzer) extractEmbeddedMarkdownLinks(embeddedAssets fs.FS, filePath string) ([]string, error) {
	content, err := fs.ReadFile(embeddedAssets, filePath)
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

// ValidateEmbeddedAgentDependencies implements EmbeddedAssetAnalyzer interface
// Simple validation that ensures all agent dependencies exist in embedded assets
func (a *FrameworkAnalyzer) ValidateEmbeddedAgentDependencies(ctx context.Context, embeddedAssets fs.FS, agentNames []string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	for _, agentName := range agentNames {
		// Check for cancellation
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		agentPath := fmt.Sprintf(agentFilePattern, embeddedPrefix, agentsDir, agentName)

		// Check if agent file exists
		if _, err := embeddedAssets.Open(agentPath); err != nil {
			issues = append(issues, ValidationIssue{
				Type:        "missing_embedded_agent",
				Severity:    SeverityCritical,
				File:        agentPath,
				Line:        0,
				Message:     fmt.Sprintf("Agent not found in embedded assets: %s", agentName),
				FixGuidance: "Ensure agent file exists in embedded framework assets",
			})
			continue
		}

		// Validate task dependencies exist
		taskIssues, err := a.validateEmbeddedAgentTasks(embeddedAssets, agentPath)
		if err != nil {
			return nil, err
		}
		issues = append(issues, taskIssues...)
	}

	return issues, nil
}

// validateEmbeddedAgentTasks validates task dependencies for embedded agent
func (a *FrameworkAnalyzer) validateEmbeddedAgentTasks(embeddedAssets fs.FS, agentPath string) ([]ValidationIssue, error) {
	var issues []ValidationIssue

	taskRefs, err := a.extractEmbeddedYAMLTasks(embeddedAssets, agentPath)
	if err != nil {
		return nil, err
	}

	for _, taskRef := range taskRefs {
		// Convert task reference to embedded path
		embeddedTaskPath := strings.Replace(taskRef, taskReferencePrefix, embeddedPrefix+"/", 1)

		// Check if task file exists in embedded assets
		if _, err := embeddedAssets.Open(embeddedTaskPath); err != nil {
			issues = append(issues, ValidationIssue{
				Type:        "missing_embedded_task",
				Severity:    SeverityCritical,
				File:        agentPath,
				Line:        0,
				Message:     fmt.Sprintf("Missing task file in embedded assets: %s", taskRef),
				FixGuidance: "Ensure task file exists in embedded framework assets",
			})
		}
	}

	return issues, nil
}
