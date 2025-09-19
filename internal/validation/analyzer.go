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
	"os"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
)

// AgentStats holds statistics for a single agent
type AgentStats struct {
	Name          string
	TaskCount     int
	TemplateCount int
	DataFileCount int
}

// UsageStats holds information about most used components
type UsageStats struct {
	Path  string
	Count int
}

// FrameworkInsights provides component statistics and relationship analysis
type FrameworkInsights struct {
	TotalAgents      int
	TotalTasks       int
	TotalTemplates   int
	TotalDataFiles   int
	TotalReferences  int
	AgentStats       []AgentStats
	MostUsedTemplate *UsageStats
	MostUsedTask     *UsageStats
	MostUsedDataFile *UsageStats
}

// ValidationIssue represents a single validation issue
type ValidationIssue struct {
	File    string
	Message string
}

// FrameworkAnalyzer provides comprehensive framework validation
type FrameworkAnalyzer struct {
	discovery *assets.Discovery
}

// NewFrameworkAnalyzer creates a new framework analyzer
func NewFrameworkAnalyzer(discovery *assets.Discovery) *FrameworkAnalyzer {
	return &FrameworkAnalyzer{
		discovery: discovery,
	}
}

// AnalyzeFramework performs comprehensive framework analysis
func (a *FrameworkAnalyzer) AnalyzeFramework() ([]ValidationIssue, *FrameworkInsights, error) {
	agents, err := a.discovery.GetAgents(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get agents: %w", err)
	}

	var issues []ValidationIssue
	var agentStats []AgentStats
	templateUsage := make(map[string]int)
	taskUsage := make(map[string]int)
	dataFileUsage := make(map[string]int)
	totalReferences := 0

	for _, agent := range agents {
		agentIssues := a.validateAgentFiles(agent)
		issues = append(issues, agentIssues...)

		stats, refs := a.collectAgentStats(agent, templateUsage, taskUsage, dataFileUsage)
		agentStats = append(agentStats, stats)
		totalReferences += refs
	}

	insights := a.buildInsights(agents, agentStats, templateUsage, taskUsage, dataFileUsage, totalReferences)
	return issues, insights, nil
}

// validateAgentFiles validates all files referenced by an agent
func (a *FrameworkAnalyzer) validateAgentFiles(agent assets.Agent) []ValidationIssue {
	var issues []ValidationIssue

	// Check agent file exists
	if _, err := os.Stat(agent.FilePath); os.IsNotExist(err) {
		issues = append(issues, ValidationIssue{
			File:    agent.FilePath,
			Message: fmt.Sprintf("Agent file does not exist: %s (agent: %s)", agent.FilePath, agent.ShortName),
		})
	}

	// Check files referenced by each task
	for _, task := range agent.Tasks {
		// Check template files for this task
		for _, template := range task.Dependencies.Templates {
			if _, err := os.Stat(template.Path); os.IsNotExist(err) {
				issues = append(issues, ValidationIssue{
					File:    template.Path,
					Message: fmt.Sprintf("Template file does not exist: %s (agent: %s, task: %s)", template.Path, agent.ShortName, task.Name),
				})
			}
		}

		// Check data files for this task
		for _, dataFile := range task.Dependencies.DataFiles {
			if _, err := os.Stat(dataFile.Path); os.IsNotExist(err) {
				issues = append(issues, ValidationIssue{
					File:    dataFile.Path,
					Message: fmt.Sprintf("Data file does not exist: %s (agent: %s, task: %s)", dataFile.Path, agent.ShortName, task.Name),
				})
			}
		}

		// Check referenced task files for this task
		for _, taskRef := range task.Dependencies.Tasks {
			if _, err := os.Stat(taskRef.Path); os.IsNotExist(err) {
				issues = append(issues, ValidationIssue{
					File:    taskRef.Path,
					Message: fmt.Sprintf("Referenced task file does not exist: %s (agent: %s, task: %s)", taskRef.Path, agent.ShortName, task.Name),
				})
			}
		}
	}

	// Check task files
	for _, taskPath := range agent.GetAllTasksPaths() {
		if _, err := os.Stat(taskPath); os.IsNotExist(err) {
			issues = append(issues, ValidationIssue{
				File:    taskPath,
				Message: fmt.Sprintf("Task file does not exist: %s (agent: %s)", taskPath, agent.ShortName),
			})
		}
	}

	return issues
}

// collectAgentStats collects statistics for an agent and updates usage counters
func (a *FrameworkAnalyzer) collectAgentStats(agent assets.Agent, templateUsage, taskUsage, dataFileUsage map[string]int) (AgentStats, int) {
	templates := agent.GetAllTemplatesPaths()
	dataFiles := agent.GetAllDataFilesPaths()
	tasks := agent.GetAllTasksPaths()
	referencedTasks := agent.GetAllReferencedTasksPaths()

	// Update usage counters
	refs := 0
	for _, template := range templates {
		templateUsage[template]++
		refs++
	}
	for _, task := range tasks {
		taskUsage[task]++
		refs++
	}
	for _, task := range referencedTasks {
		taskUsage[task]++
		refs++
	}
	for _, dataFile := range dataFiles {
		dataFileUsage[dataFile]++
		refs++
	}

	stats := AgentStats{
		Name:          agent.ShortName,
		TaskCount:     len(tasks),
		TemplateCount: len(templates),
		DataFileCount: len(dataFiles),
	}

	return stats, refs
}

// buildInsights creates the final FrameworkInsights structure
func (a *FrameworkAnalyzer) buildInsights(agents []assets.Agent, agentStats []AgentStats, templateUsage, taskUsage, dataFileUsage map[string]int, totalReferences int) *FrameworkInsights {
	uniqueTemplates, uniqueTasks, uniqueDataFiles := a.calculateUniqueCounts(agents)
	mostUsedTemplate := a.findMostUsed(templateUsage)
	mostUsedTask := a.findMostUsed(taskUsage)
	mostUsedDataFile := a.findMostUsed(dataFileUsage)

	return &FrameworkInsights{
		TotalAgents:      len(agents),
		TotalTasks:       len(uniqueTasks),
		TotalTemplates:   len(uniqueTemplates),
		TotalDataFiles:   len(uniqueDataFiles),
		TotalReferences:  totalReferences,
		AgentStats:       agentStats,
		MostUsedTemplate: mostUsedTemplate,
		MostUsedTask:     mostUsedTask,
		MostUsedDataFile: mostUsedDataFile,
	}
}

// calculateUniqueCounts calculates unique file counts across all agents
func (a *FrameworkAnalyzer) calculateUniqueCounts(agents []assets.Agent) (map[string]struct{}, map[string]struct{}, map[string]struct{}) {
	uniqueTemplates := make(map[string]struct{})
	uniqueTasks := make(map[string]struct{})
	uniqueDataFiles := make(map[string]struct{})

	for _, agent := range agents {
		for _, template := range agent.GetAllTemplatesPaths() {
			uniqueTemplates[template] = struct{}{}
		}
		for _, task := range agent.GetAllTasksPaths() {
			uniqueTasks[task] = struct{}{}
		}
		for _, task := range agent.GetAllReferencedTasksPaths() {
			uniqueTasks[task] = struct{}{}
		}
		for _, dataFile := range agent.GetAllDataFilesPaths() {
			uniqueDataFiles[dataFile] = struct{}{}
		}
	}

	return uniqueTemplates, uniqueTasks, uniqueDataFiles
}

// findMostUsed finds the most used component from a usage map
func (a *FrameworkAnalyzer) findMostUsed(usage map[string]int) *UsageStats {
	var mostUsed *UsageStats
	for path, count := range usage {
		if mostUsed == nil || count > mostUsed.Count {
			mostUsed = &UsageStats{Path: path, Count: count}
		}
	}
	return mostUsed
}
