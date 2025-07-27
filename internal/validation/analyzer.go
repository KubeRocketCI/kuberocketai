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
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
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

	// 4. Detect invalid file formats
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
	dirs := []string{"agents", "tasks", "templates", "data"}

	for _, dir := range dirs {
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
