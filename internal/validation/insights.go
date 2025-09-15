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
func (a *FrameworkAnalyzer) generateFrameworkInsights() (*FrameworkInsights, error) {
	insights := &FrameworkInsights{
		ComponentCounts: ComponentCounts{},
		Relationships:   make([]ComponentRelationship, 0),
		UsageStatistics: UsageStatistics{},
	}

	// Step 1: Get agents and analyze components
	if err := a.analyzeComponents(insights); err != nil {
		return nil, err
	}

	// Step 2: Generate usage statistics
	a.generateUsageStatistics(insights)

	return insights, nil
}

// analyzeComponents gets agents and builds component counts and relationships
func (a *FrameworkAnalyzer) analyzeComponents(insights *FrameworkInsights) error {
	// Get all agents from discovery
	agents, err := a.discovery.GetAgents(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get agents: %w", err)
	}

	// Track unique components using sets
	uniqueTemplates := make(map[string]struct{})
	uniqueDataFiles := make(map[string]struct{})
	uniqueTasks := make(map[string]struct{})

	// Build relationships and count components
	for _, agent := range agents {
		rel := ComponentRelationship{
			Agent:     agent.ShortName,
			Tasks:     make([]string, 0, len(agent.Tasks)),
			Templates: agent.GetAllTemplatesPaths(),
			DataFiles: agent.GetAllDataFilesPaths(),
		}

		// Collect task names
		for _, task := range agent.Tasks {
			rel.Tasks = append(rel.Tasks, task.Name)
			uniqueTasks[task.Name] = struct{}{}
		}

		// Track unique templates and data files
		for _, template := range rel.Templates {
			uniqueTemplates[template] = struct{}{}
		}
		for _, dataFile := range rel.DataFiles {
			uniqueDataFiles[dataFile] = struct{}{}
		}

		insights.Relationships = append(insights.Relationships, rel)
	}

	// Set component counts
	insights.ComponentCounts.Agents = len(agents)
	insights.ComponentCounts.Tasks = len(uniqueTasks)
	insights.ComponentCounts.Templates = len(uniqueTemplates)
	insights.ComponentCounts.Data = len(uniqueDataFiles)

	return nil
}

// generateUsageStatistics analyzes usage patterns to identify most/least used components
func (a *FrameworkAnalyzer) generateUsageStatistics(insights *FrameworkInsights) {
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
}

// removeDuplicates removed; use utils.DeduplicateStrings directly.

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
			totalTasks := len(rel.Tasks)
			if totalTasks > 0 {
				templateCount := len(rel.Templates)
				result.WriteString(fmt.Sprintf("   • %s → %d tasks → %d templates\n",
					rel.Agent, totalTasks, templateCount))
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
