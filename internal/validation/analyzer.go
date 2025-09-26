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
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

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

	// Collect file usage information for deduplication
	fileUsage := a.collectFileUsage(agents)

	for _, agent := range agents {
		agentIssues := a.validateAgentFiles(agent)
		issues = append(issues, agentIssues...)

		stats, refs := a.collectAgentStats(agent, templateUsage, taskUsage, dataFileUsage)
		agentStats = append(agentStats, stats)
		totalReferences += refs
	}

	// Deduplicate XML validation issues
	deduplicatedXMLIssues := a.deduplicateXMLValidationIssues(fileUsage)
	issues = append(issues, deduplicatedXMLIssues...)

	insights := a.buildInsights(agents, agentStats, templateUsage, taskUsage, dataFileUsage, totalReferences)
	return issues, insights, nil
}

// FileReference represents how a file is referenced by agents/tasks
type FileReference struct {
	FilePath   string
	FileType   string // "task", "template", "data file"
	References []AgentTaskRef
}

// AgentTaskRef represents an agent/task combination that references a file
type AgentTaskRef struct {
	AgentName string
	TaskName  string // empty for agent-level files
}

// collectFileUsage collects information about which files are referenced by which agents/tasks
func (a *FrameworkAnalyzer) collectFileUsage(agents []assets.Agent) map[string]*FileReference {
	fileUsage := make(map[string]*FileReference)

	for _, agent := range agents {
		// Collect task files
		for _, taskPath := range agent.GetAllTasksPaths() {
			if fileRef, exists := fileUsage[taskPath]; exists {
				fileRef.References = append(fileRef.References, AgentTaskRef{
					AgentName: agent.ShortName,
					TaskName:  "",
				})
			} else {
				fileUsage[taskPath] = &FileReference{
					FilePath: taskPath,
					FileType: "task",
					References: []AgentTaskRef{{
						AgentName: agent.ShortName,
						TaskName:  "",
					}},
				}
			}
		}

		// Collect files referenced by each task
		for _, task := range agent.Tasks {
			// Template files
			for _, template := range task.Dependencies.Templates {
				if fileRef, exists := fileUsage[template.Path]; exists {
					fileRef.References = append(fileRef.References, AgentTaskRef{
						AgentName: agent.ShortName,
						TaskName:  task.Name,
					})
				} else {
					fileUsage[template.Path] = &FileReference{
						FilePath: template.Path,
						FileType: "template",
						References: []AgentTaskRef{{
							AgentName: agent.ShortName,
							TaskName:  task.Name,
						}},
					}
				}
			}

			// Data files
			for _, dataFile := range task.Dependencies.DataFiles {
				if fileRef, exists := fileUsage[dataFile.Path]; exists {
					fileRef.References = append(fileRef.References, AgentTaskRef{
						AgentName: agent.ShortName,
						TaskName:  task.Name,
					})
				} else {
					fileUsage[dataFile.Path] = &FileReference{
						FilePath: dataFile.Path,
						FileType: "data file",
						References: []AgentTaskRef{{
							AgentName: agent.ShortName,
							TaskName:  task.Name,
						}},
					}
				}
			}

			// Referenced task files
			for _, taskRef := range task.Dependencies.Tasks {
				if fileRef, exists := fileUsage[taskRef.Path]; exists {
					fileRef.References = append(fileRef.References, AgentTaskRef{
						AgentName: agent.ShortName,
						TaskName:  task.Name,
					})
				} else {
					fileUsage[taskRef.Path] = &FileReference{
						FilePath: taskRef.Path,
						FileType: "referenced task",
						References: []AgentTaskRef{{
							AgentName: agent.ShortName,
							TaskName:  task.Name,
						}},
					}
				}
			}
		}
	}

	return fileUsage
}

// validateAgentFiles validates all files referenced by an agent (excluding XML validation which is handled separately)
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

// deduplicateXMLValidationIssues validates each unique file once and creates consolidated error messages
func (a *FrameworkAnalyzer) deduplicateXMLValidationIssues(fileUsage map[string]*FileReference) []ValidationIssue {
	var issues []ValidationIssue

	for filePath, fileRef := range fileUsage {
		// Check if file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// File doesn't exist - skip XML validation (existence validation is handled elsewhere)
			continue
		}

		// Skip XML validation for YAML files
		ext := strings.ToLower(filepath.Ext(filePath))
		if ext == ".yaml" || ext == ".yml" {
			continue
		}

		// Read file content
		file, err := os.Open(filePath)
		if err != nil {
			// Cannot open file - skip XML validation
			continue
		}
		defer func() {
			_ = file.Close()
		}()

		content, err := io.ReadAll(file)
		if err != nil {
			// Cannot read file - skip XML validation
			continue
		}

		// Validate XML tags in content
		xmlIssues := a.validateXMLTags(string(content))
		if len(xmlIssues) == 0 {
			continue // No XML issues in this file
		}

		// Create consolidated error messages for each XML issue
		for _, xmlIssue := range xmlIssues {
			var referencedBy string
			if len(fileRef.References) == 1 {
				// Single reference - use original format
				ref := fileRef.References[0]
				if ref.TaskName != "" {
					referencedBy = fmt.Sprintf("(agent: %s, task: %s, %s)", ref.AgentName, ref.TaskName, fileRef.FileType)
				} else {
					referencedBy = fmt.Sprintf("(agent: %s, %s)", ref.AgentName, fileRef.FileType)
				}
			} else {
				// Multiple references - use consolidated format
				referencedBy = a.formatMultipleReferences(fileRef)
			}

			issues = append(issues, ValidationIssue{
				File:    filePath,
				Message: fmt.Sprintf("XML tag validation error: %s %s - File: %s", xmlIssue, referencedBy, filePath),
			})
		}
	}

	return issues
}

// formatMultipleReferences creates a consolidated reference string for files used by multiple agents/tasks
func (a *FrameworkAnalyzer) formatMultipleReferences(fileRef *FileReference) string {
	agentTaskMap := make(map[string][]string)

	for _, ref := range fileRef.References {
		if ref.TaskName != "" {
			agentTaskMap[ref.AgentName] = append(agentTaskMap[ref.AgentName], ref.TaskName)
		} else {
			// Agent-level reference (task name is empty)
			if _, exists := agentTaskMap[ref.AgentName]; !exists {
				agentTaskMap[ref.AgentName] = []string{}
			}
		}
	}

	// Sort agent names for deterministic output
	var agentNames []string
	for agentName := range agentTaskMap {
		agentNames = append(agentNames, agentName)
	}
	sort.Strings(agentNames)

	var parts []string
	for _, agentName := range agentNames {
		tasks := agentTaskMap[agentName]
		if len(tasks) == 0 {
			parts = append(parts, agentName)
		} else {
			taskList := strings.Join(tasks, ", ")
			parts = append(parts, fmt.Sprintf("%s(%s)", agentName, taskList))
		}
	}

	return fmt.Sprintf("(%s, %s) - Referenced by: %s", fileRef.FileType, "multiple agents/tasks", strings.Join(parts, ", "))
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
