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
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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
	embeddedPrefix = "assets/framework/core"
)

// Path pattern constants
const (
	agentFilePattern    = "%s/%s/%s.yaml"
	localTasksPrefix    = "./.krci-ai/local/tasks/"
	standardTasksPrefix = "./.krci-ai/tasks/"
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

// ValidationIssue represents a single validation issue (legacy compatibility)
type ValidationIssue struct {
	Type        string
	Severity    Severity
	File        string
	Line        int
	Message     string
	FixGuidance string
}

// FrameworkAnalyzer provides comprehensive framework validation using unified system
type FrameworkAnalyzer struct {
	validator UniversalValidator
	baseDir   string
	cache     map[string]time.Time
}

// NewFrameworkAnalyzer creates a new framework analyzer using the unified validation system
func NewFrameworkAnalyzer(baseDir string, embeddedAssets embed.FS) (*FrameworkAnalyzer, error) {
	validator, err := NewValidationSystemBuilder(baseDir).
		WithEmbeddedAssets(embeddedAssets).
		Build()

	if err != nil {
		return nil, fmt.Errorf("failed to create unified validator: %w", err)
	}

	return &FrameworkAnalyzer{
		validator: validator,
		baseDir:   baseDir,
		cache:     make(map[string]time.Time),
	}, nil
}

// AnalyzeFramework performs comprehensive framework analysis
func (a *FrameworkAnalyzer) AnalyzeFramework() ([]ValidationIssue, *FrameworkInsights, error) {
	// Use the unified validator to validate the entire framework
	report := a.validator.ValidateFramework(a.baseDir)

	// Convert unified results back to legacy format for compatibility
	issues := make([]ValidationIssue, 0, len(report.Results))
	for _, result := range report.Results {
		issues = append(issues, ValidationIssue{
			Type:        result.Type,
			Severity:    result.Severity,
			File:        result.File,
			Line:        result.Line,
			Message:     result.Message,
			FixGuidance: result.FixGuidance,
		})
	}

	// Generate framework insights
	insights := a.generateFrameworkInsights()

	return issues, insights, nil
}

// OptimizedAnalyzeFramework provides optimized framework analysis with caching
func (a *FrameworkAnalyzer) OptimizedAnalyzeFramework() ([]ValidationIssue, *FrameworkInsights, error) {
	// Check if analysis is needed based on file modifications
	frameworkDir := filepath.Join(a.baseDir, ".krci-ai")
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("no .krci-ai directory found")
	}

	// For now, always run full analysis - caching optimization can be added later
	return a.AnalyzeFramework()
}

// generateFrameworkInsights generates comprehensive insights about the framework
func (a *FrameworkAnalyzer) generateFrameworkInsights() *FrameworkInsights {
	frameworkDir := filepath.Join(a.baseDir, ".krci-ai")

	insights := &FrameworkInsights{
		ComponentCounts: ComponentCounts{},
		Relationships:   make([]ComponentRelationship, 0),
		UsageStatistics: UsageStatistics{},
	}

	// Count components
	a.countComponents(frameworkDir, &insights.ComponentCounts)

	// Analyze relationships (simplified)
	a.analyzeRelationships(frameworkDir, insights)

	// Generate usage statistics
	a.generateUsageStatistics(insights)

	return insights
}

// countComponents counts framework components
func (a *FrameworkAnalyzer) countComponents(frameworkDir string, counts *ComponentCounts) {
	// Count agents
	agentFiles, err := filepath.Glob(filepath.Join(frameworkDir, agentsDir, "*.y*ml"))
	if err == nil {
		counts.Agents = len(agentFiles)
	}

	// Count tasks
	taskFiles, err := filepath.Glob(filepath.Join(frameworkDir, tasksDir, "*.md"))
	if err == nil {
		counts.Tasks = len(taskFiles)
	}

	// Count local tasks
	localTaskFiles, err := filepath.Glob(filepath.Join(frameworkDir, "local", tasksDir, "*.md"))
	if err == nil {
		counts.Tasks += len(localTaskFiles)
	}

	// Count templates
	templateFiles, err := filepath.Glob(filepath.Join(frameworkDir, templatesDir, "*.md"))
	if err == nil {
		counts.Templates = len(templateFiles)
	}

	// Count data files
	dataFiles, err := filepath.Glob(filepath.Join(frameworkDir, dataDir, "*.md"))
	if err == nil {
		counts.Data = len(dataFiles)
	}
}

// analyzeRelationships analyzes component relationships (simplified)
func (a *FrameworkAnalyzer) analyzeRelationships(frameworkDir string, insights *FrameworkInsights) {
	agentsDir := filepath.Join(frameworkDir, "agents")
	agentFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.y*ml"))
	if err != nil {
		return // No agents directory
	}

	for _, agentFile := range agentFiles {
		agentName := strings.TrimSuffix(filepath.Base(agentFile), filepath.Ext(agentFile))

		relationship := ComponentRelationship{
			Agent:      agentName,
			Tasks:      make([]string, 0),
			LocalTasks: make([]string, 0),
			Templates:  make([]string, 0),
			DataFiles:  make([]string, 0),
		}

		// Extract tasks from agent YAML
		if tasks, err := a.extractAgentTasks(agentFile); err == nil {
			for _, task := range tasks {
				taskName := filepath.Base(task)
				if strings.Contains(task, localTasksPrefix) {
					relationship.LocalTasks = append(relationship.LocalTasks, taskName)
				} else {
					relationship.Tasks = append(relationship.Tasks, taskName)
				}
			}
		}

		insights.Relationships = append(insights.Relationships, relationship)
	}
}

// extractAgentTasks extracts task references from agent YAML
func (a *FrameworkAnalyzer) extractAgentTasks(agentFile string) ([]string, error) {
	content, err := os.ReadFile(agentFile)
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

	var tasks []string
	for _, task := range agent.Agent.Tasks {
		if strings.HasPrefix(task, standardTasksPrefix) || strings.HasPrefix(task, localTasksPrefix) {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

// generateUsageStatistics generates usage statistics
func (a *FrameworkAnalyzer) generateUsageStatistics(insights *FrameworkInsights) {
	insights.UsageStatistics = UsageStatistics{
		MostUsedTemplate:   "",
		MostUsedData:       "",
		TemplateUsageCount: 0,
		DataUsageCount:     0,
	}
}

// AnalyzeEmbeddedDependencies implements EmbeddedAssetAnalyzer interface
func (a *FrameworkAnalyzer) AnalyzeEmbeddedDependencies(ctx context.Context, embeddedAssets fs.FS, agentNames []string) (*FrameworkInsights, error) {
	insights := &FrameworkInsights{
		ComponentCounts: ComponentCounts{},
		Relationships:   make([]ComponentRelationship, 0),
		UsageStatistics: UsageStatistics{},
	}

	// Count embedded components
	a.countEmbeddedComponents(embeddedAssets, &insights.ComponentCounts)

	// Analyze relationships for selected agents
	for _, agentName := range agentNames {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			relationship, err := a.buildEmbeddedAgentRelationship(embeddedAssets, agentName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: failed to build relationship for agent %s: %v\n", agentName, err)
				continue
			}
			insights.Relationships = append(insights.Relationships, *relationship)
		}
	}

	// Generate usage statistics
	a.generateUsageStatistics(insights)

	return insights, nil
}

// countEmbeddedComponents counts components in embedded filesystem
func (a *FrameworkAnalyzer) countEmbeddedComponents(embeddedAssets fs.FS, counts *ComponentCounts) {
	// Count agents
	agentFiles, err := fs.Glob(embeddedAssets, embeddedPrefix+"/"+agentsDir+"/*.y*ml")
	if err == nil {
		counts.Agents = len(agentFiles)
	}

	// Count tasks
	taskFiles, err := fs.Glob(embeddedAssets, embeddedPrefix+"/"+tasksDir+"/*.md")
	if err == nil {
		counts.Tasks = len(taskFiles)
	}

	// Count templates
	templateFiles, err := fs.Glob(embeddedAssets, embeddedPrefix+"/"+templatesDir+"/*.md")
	if err == nil {
		counts.Templates = len(templateFiles)
	}

	// Count data files
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
	if err == nil {
		counts.Data = dataCount
	}
}

// buildEmbeddedAgentRelationship builds a ComponentRelationship for an embedded agent
func (a *FrameworkAnalyzer) buildEmbeddedAgentRelationship(embeddedAssets fs.FS, agentName string) (*ComponentRelationship, error) {
	agentPath := fmt.Sprintf(agentFilePattern, embeddedPrefix, agentsDir, agentName)

	relationship := ComponentRelationship{
		Agent:      agentName,
		Tasks:      make([]string, 0),
		LocalTasks: make([]string, 0),
		Templates:  make([]string, 0),
		DataFiles:  make([]string, 0),
	}

	// Extract tasks from embedded agent YAML
	taskRefs, err := a.extractEmbeddedYAMLTasks(embeddedAssets, agentPath)
	if err != nil {
		return nil, fmt.Errorf("extracting tasks for agent %s: %w", agentName, err)
	}

	// Process tasks
	for _, taskRef := range taskRefs {
		taskName := filepath.Base(taskRef)
		if strings.Contains(taskRef, localTasksPrefix) {
			relationship.LocalTasks = append(relationship.LocalTasks, taskName)
		} else {
			relationship.Tasks = append(relationship.Tasks, taskName)
		}
	}

	// Remove duplicates
	relationship.Templates = utils.DeduplicateStrings(relationship.Templates)
	relationship.DataFiles = utils.DeduplicateStrings(relationship.DataFiles)

	return &relationship, nil
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

	var tasks []string
	for _, task := range agent.Agent.Tasks {
		if strings.HasPrefix(task, standardTasksPrefix) || strings.HasPrefix(task, localTasksPrefix) {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}
