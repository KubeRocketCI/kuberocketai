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
	"sort"
	"strings"
)

// FrameworkInsights provides component statistics and relationship analysis
type FrameworkInsights struct {
	ComponentCounts ComponentCounts         `json:"component_counts"`
	Relationships   []ComponentRelationship `json:"relationships"`
	UsageStatistics UsageStatistics         `json:"usage_statistics"`
	TotalReferences int                     `json:"total_references"`
}

// ComponentCounts tracks the number of each component type
type ComponentCounts struct {
	Agents    int `json:"agents"`
	Tasks     int `json:"tasks"`
	Templates int `json:"templates"`
	Data      int `json:"data"`
}

// ComponentRelationship represents agent → task → template/data flows
type ComponentRelationship struct {
	Agent     string   `json:"agent"`
	Tasks     []string `json:"tasks"`
	Templates []string `json:"templates"`
	DataFiles []string `json:"data_files"`
}

// UsageStatistics provides insights about component usage patterns
type UsageStatistics struct {
	MostUsedTemplate   string `json:"most_used_template"`
	MostUsedData       string `json:"most_used_data"`
	TemplateUsageCount int    `json:"template_usage_count"`
	DataUsageCount     int    `json:"data_usage_count"`
}

// generateFrameworkInsights analyzes the framework and generates insights
func (a *FrameworkAnalyzer) generateFrameworkInsights(frameworkDir string) (*FrameworkInsights, error) {
	insights := &FrameworkInsights{
		ComponentCounts: ComponentCounts{},
		Relationships:   make([]ComponentRelationship, 0),
		UsageStatistics: UsageStatistics{},
	}

	// Step 1: Count components
	if err := a.countComponents(frameworkDir, &insights.ComponentCounts); err != nil {
		return nil, err
	}

	// Step 2: Analyze relationships
	if err := a.analyzeRelationships(frameworkDir, insights); err != nil {
		return nil, err
	}

	// Step 3: Generate usage statistics
	if err := a.generateUsageStatistics(frameworkDir, insights); err != nil {
		return nil, err
	}

	return insights, nil
}

// countComponents counts the number of each component type
func (a *FrameworkAnalyzer) countComponents(frameworkDir string, counts *ComponentCounts) error {
	// Count agents
	agentsDir := filepath.Join(frameworkDir, "agents")
	if _, err := os.Stat(agentsDir); err == nil {
		agentFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.y*ml"))
		if err != nil {
			return err
		}
		counts.Agents = len(agentFiles)
	}

	// Count tasks
	tasksDir := filepath.Join(frameworkDir, "tasks")
	if _, err := os.Stat(tasksDir); err == nil {
		taskFiles, err := filepath.Glob(filepath.Join(tasksDir, "*.md"))
		if err != nil {
			return err
		}
		counts.Tasks = len(taskFiles)
	}

	// Count templates
	templatesDir := filepath.Join(frameworkDir, "templates")
	if _, err := os.Stat(templatesDir); err == nil {
		templateFiles, err := filepath.Glob(filepath.Join(templatesDir, "*.md"))
		if err != nil {
			return err
		}
		counts.Templates = len(templateFiles)
	}

	// Count data files
	dataDir := filepath.Join(frameworkDir, "data")
	if _, err := os.Stat(dataDir); err == nil {
		dataFiles, err := filepath.Glob(filepath.Join(dataDir, "*.*"))
		if err != nil {
			return err
		}
		counts.Data = len(dataFiles)
	}

	return nil
}

// analyzeRelationships analyzes agent → task → template/data relationships
func (a *FrameworkAnalyzer) analyzeRelationships(frameworkDir string, insights *FrameworkInsights) error {
	agentsDir := filepath.Join(frameworkDir, "agents")
	if _, err := os.Stat(agentsDir); os.IsNotExist(err) {
		return nil
	}

	agentFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.y*ml"))
	if err != nil {
		return err
	}

	for _, agentFile := range agentFiles {
		agentName := strings.TrimSuffix(filepath.Base(agentFile), filepath.Ext(agentFile))

		relationship := ComponentRelationship{
			Agent:     agentName,
			Tasks:     make([]string, 0),
			Templates: make([]string, 0),
			DataFiles: make([]string, 0),
		}

		// Get tasks referenced by this agent
		taskRefs, err := a.extractYAMLTasks(agentFile)
		if err != nil {
			continue
		}

		for _, taskRef := range taskRefs {
			taskName := filepath.Base(taskRef)
			relationship.Tasks = append(relationship.Tasks, taskName)

			// For each task, find what templates/data it references
			cleanPath := strings.TrimPrefix(taskRef, "./")
			taskPath := filepath.Join(a.baseDir, cleanPath)

			if _, err := os.Stat(taskPath); err == nil {
				templates, dataFiles := a.getTaskReferences(taskPath)
				relationship.Templates = append(relationship.Templates, templates...)
				relationship.DataFiles = append(relationship.DataFiles, dataFiles...)
			}
		}

		// Remove duplicates
		relationship.Templates = a.removeDuplicates(relationship.Templates)
		relationship.DataFiles = a.removeDuplicates(relationship.DataFiles)

		insights.Relationships = append(insights.Relationships, relationship)
	}

	return nil
}

// getTaskReferences extracts template and data references from a task file
func (a *FrameworkAnalyzer) getTaskReferences(taskFile string) ([]string, []string) {
	var templates, dataFiles []string

	links, err := a.extractMarkdownLinks(taskFile)
	if err != nil {
		return templates, dataFiles
	}

	for _, link := range links {
		fileName := filepath.Base(link)
		if strings.Contains(link, "/templates/") {
			templates = append(templates, fileName)
		} else if strings.Contains(link, "/data/") {
			dataFiles = append(dataFiles, fileName)
		}
	}

	return templates, dataFiles
}

// generateUsageStatistics analyzes usage patterns to identify most/least used components
func (a *FrameworkAnalyzer) generateUsageStatistics(frameworkDir string, insights *FrameworkInsights) error {
	templateUsage := make(map[string]int)
	dataUsage := make(map[string]int)

	// Count references across all relationships
	for _, rel := range insights.Relationships {
		for _, template := range rel.Templates {
			templateUsage[template]++
			insights.TotalReferences++
		}
		for _, dataFile := range rel.DataFiles {
			dataUsage[dataFile]++
			insights.TotalReferences++
		}
	}

	// Find most used template
	maxTemplateCount := 0
	for template, count := range templateUsage {
		if count > maxTemplateCount {
			maxTemplateCount = count
			insights.UsageStatistics.MostUsedTemplate = template
			insights.UsageStatistics.TemplateUsageCount = count
		}
	}

	// Find most used data file
	maxDataCount := 0
	for dataFile, count := range dataUsage {
		if count > maxDataCount {
			maxDataCount = count
			insights.UsageStatistics.MostUsedData = dataFile
			insights.UsageStatistics.DataUsageCount = count
		}
	}

	return nil
}

// removeDuplicates removes duplicate strings from a slice
func (a *FrameworkAnalyzer) removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}

	sort.Strings(result)
	return result
}

// FormatInsights formats the insights for display
func (insights *FrameworkInsights) FormatInsights() string {
	var result strings.Builder

	result.WriteString(fmt.Sprintf("📊 Overview: %d agents, %d tasks, %d templates, %d data files\n",
		insights.ComponentCounts.Agents,
		insights.ComponentCounts.Tasks,
		insights.ComponentCounts.Templates,
		insights.ComponentCounts.Data))

	result.WriteString(fmt.Sprintf("🔗 All internal links resolved (%d references checked)\n",
		insights.TotalReferences))

	if len(insights.Relationships) > 0 {
		result.WriteString("\n💡 FRAMEWORK INSIGHTS:\n")

		for _, rel := range insights.Relationships {
			if len(rel.Tasks) > 0 {
				templateCount := len(rel.Templates)
				result.WriteString(fmt.Sprintf("   • %s → %d tasks → %d templates\n",
					rel.Agent, len(rel.Tasks), templateCount))
			}
		}

		if insights.UsageStatistics.MostUsedTemplate != "" {
			result.WriteString(fmt.Sprintf("   • Most used template: %s (used by %d tasks)\n",
				insights.UsageStatistics.MostUsedTemplate,
				insights.UsageStatistics.TemplateUsageCount))
		}
	}

	return result.String()
}
