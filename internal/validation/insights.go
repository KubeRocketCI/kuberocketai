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
	Agent      string   `json:"agent"`
	Tasks      []string `json:"tasks"`
	LocalTasks []string `json:"local_tasks"` // Local tasks from .krci-ai/local/tasks/
	Templates  []string `json:"templates"`
	DataFiles  []string `json:"data_files"`
}

// UsageStatistics provides insights about component usage patterns
type UsageStatistics struct {
	MostUsedTemplate   string `json:"most_used_template"`
	MostUsedData       string `json:"most_used_data"`
	TemplateUsageCount int    `json:"template_usage_count"`
	DataUsageCount     int    `json:"data_usage_count"`
}

// getTaskReferences moved to unified validation system
// Use the ValidationStrategy interface for link extraction

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
			totalTasks := len(rel.Tasks) + len(rel.LocalTasks)
			if totalTasks > 0 {
				templateCount := len(rel.Templates)
				if len(rel.LocalTasks) > 0 {
					result.WriteString(fmt.Sprintf("   • %s → %d tasks (including %d local) → %d templates\n",
						rel.Agent, totalTasks, len(rel.LocalTasks), templateCount))
				} else {
					result.WriteString(fmt.Sprintf("   • %s → %d tasks → %d templates\n",
						rel.Agent, totalTasks, templateCount))
				}
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
