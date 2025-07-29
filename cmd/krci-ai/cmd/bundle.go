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
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/cli"
	"github.com/KubeRocketCI/kuberocketai/internal/validation"
)

// bundleCmd represents the bundle command
var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Generate complete agent bundles for web chat tools",
	Long: `Generate complete agent bundles for web chat tools (ChatGPT, Gemini Pro, Claude Web).

This command creates a single markdown file containing all KubeRocketAI framework 
components optimized for web chat consumption. The bundle includes:
- All 6 core SDLC agents (PM, Architect, Developer, QA, BA, PO) with complete definitions
- All agent dependencies, tasks, and templates in consolidated format
- Project-specific context from .krci-ai/data/ and .krci-ai/templates/ directories
- System prompt structure with clear agent separation and role definitions

Examples:
  krci-ai bundle --all                           # Generate complete bundle at ./.krci-ai/bundle/all.md
  krci-ai bundle --agent pm,architect            # Generate targeted bundle with PM and Architect agents
  krci-ai bundle --agent pm --output pm.md       # Generate PM-only bundle with custom filename
  krci-ai bundle --agent "pm architect"          # Space-separated agent names
  krci-ai bundle --agent pm,architect --dry-run  # Show targeted bundle scope without generating files
  krci-ai bundle --all --output my-framework.md  # Generate bundle with custom filename
  krci-ai bundle --all --dry-run                 # Show bundle scope without generating files
  krci-ai bundle --help                          # Show comprehensive usage information`,
	RunE: runBundle,
}

// Bundle generation flags
var (
	bundleAll    bool
	bundleDryRun bool
	bundleOutput string
	bundleAgents string
)

// BundleContent represents the aggregated content for bundle generation
type BundleContent struct {
	Agents    []AgentBundleInfo
	Tasks     map[string]string
	Templates map[string]string
	DataFiles map[string]string
}

// AgentBundleInfo represents agent information with content for bundling
type AgentBundleInfo struct {
	FilePath string
	Content  string
	Name     string
	Role     string
}

func init() {
	rootCmd.AddCommand(bundleCmd)

	// Add flags
	bundleCmd.Flags().BoolVar(&bundleAll, "all", false, "Generate complete bundle with all agents and dependencies")
	bundleCmd.Flags().StringVar(&bundleAgents, "agent", "", "Generate targeted bundle with specific agents (comma or space separated: 'pm,architect' or 'pm architect')")
	bundleCmd.Flags().BoolVar(&bundleDryRun, "dry-run", false, "Show bundle scope without generating files")
	bundleCmd.Flags().StringVar(&bundleOutput, "output", "", "Custom output filename (creates ./.krci-ai/bundle/filename)")
}

// ParseAgentList parses comma-separated or space-separated agent names
func ParseAgentList(agentStr string) []string {
	if agentStr == "" {
		return nil
	}

	var agents []string

	// First try comma-separated
	if strings.Contains(agentStr, ",") {
		parts := strings.Split(agentStr, ",")
		for _, part := range parts {
			agent := strings.TrimSpace(part)
			if agent != "" {
				agents = append(agents, agent)
			}
		}
	} else {
		// Space-separated
		parts := strings.Fields(agentStr)
		for _, part := range parts {
			agents = append(agents, strings.TrimSpace(part))
		}
	}

	return agents
}

// validateAgentNames validates that specified agent names exist in the framework
func validateAgentNames(currentDir string, selectedAgents []string, output *cli.OutputHandler) error {
	// Create discovery service
	discovery := assets.NewDiscovery(currentDir, GetEmbeddedAssets())

	// Get available agents with full info
	agentInfos, err := discovery.DiscoverAgents()
	if err != nil {
		return fmt.Errorf("failed to discover available agents: %w", err)
	}

	// Create map for quick lookup (case-insensitive) - support both full names and file names
	agentMap := make(map[string]string)

	for _, agent := range agentInfos {
		// Add full agent name
		fullNameLower := strings.ToLower(agent.Name)
		agentMap[fullNameLower] = agent.Name

		// Add file-based name (e.g., "pm" from "pm.yaml")
		fileName := strings.TrimSuffix(filepath.Base(agent.FilePath), ".yaml")
		fileNameLower := strings.ToLower(fileName)
		agentMap[fileNameLower] = agent.Name
	}

	// Validate each selected agent
	var invalidAgents []string

	for _, selected := range selectedAgents {
		selectedLower := strings.ToLower(selected)
		if _, exists := agentMap[selectedLower]; !exists {
			invalidAgents = append(invalidAgents, selected)
		}
	}

	// Report invalid agents with suggestions
	if len(invalidAgents) > 0 {
		output.PrintError(fmt.Sprintf("Invalid agent names: %s", strings.Join(invalidAgents, ", ")))
		output.PrintInfo("Available agent names (use either full name or short name):")
		for _, agent := range agentInfos {
			fileName := strings.TrimSuffix(filepath.Base(agent.FilePath), ".yaml")
			output.PrintInfo(fmt.Sprintf("  • %s (short: %s)", agent.Name, fileName))
		}
		return fmt.Errorf("invalid agent names specified")
	}

	return nil
}

// runBundle executes the bundle command
func runBundle(cmd *cobra.Command, args []string) error {
	// Create output and error handlers
	output := cli.NewOutputHandler()
	errorHandler := cli.NewErrorHandler()

	// Validate flags - either --all or --agent must be specified
	if !bundleAll && bundleAgents == "" {
		output.PrintError("Bundle generation requires either --all or --agent flag")
		output.PrintInfo("Usage: krci-ai bundle --all  OR  krci-ai bundle --agent pm,architect")
		return fmt.Errorf("either --all or --agent flag is required")
	}

	if bundleAll && bundleAgents != "" {
		output.PrintError("Cannot specify both --all and --agent flags")
		output.PrintInfo("Use either --all for complete bundle OR --agent for targeted bundle")
		return fmt.Errorf("cannot specify both --all and --agent flags")
	}

	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	// Check if framework is installed
	frameworkDir := filepath.Join(currentDir, ".krci-ai")
	if _, statErr := os.Stat(frameworkDir); os.IsNotExist(statErr) {
		output.PrintError("No .krci-ai directory found in current directory")
		output.PrintInfo("Run 'krci-ai install' to set up the framework first")
		return fmt.Errorf("framework not installed")
	}

	output.PrintProgress("Validating framework integrity...")

	// Run framework validation before bundle generation
	if validationErr := runFrameworkValidation(currentDir); validationErr != nil {
		errorHandler.HandleError(validationErr, "Framework validation failed")
		output.PrintError("Bundle generation requires a valid framework")
		output.PrintInfo("Fix validation errors and try again")
		return validationErr
	}

	output.PrintSuccess("Framework validation passed")

	// Parse and validate agent selection if specified
	var selectedAgents []string
	if bundleAgents != "" {
		selectedAgents = ParseAgentList(bundleAgents)
		if len(selectedAgents) == 0 {
			output.PrintError("No valid agents specified")
			return fmt.Errorf("no valid agents specified")
		}

		// Validate agent names exist
		if validateErr := validateAgentNames(currentDir, selectedAgents, output); validateErr != nil {
			return validateErr
		}

		output.PrintInfo(fmt.Sprintf("Selected agents: %s", strings.Join(selectedAgents, ", ")))
	}

	// Collect bundle content
	output.PrintProgress("Discovering agents and dependencies...")

	bundleContent, err := collectBundleContent(currentDir, selectedAgents)
	if err != nil {
		errorHandler.HandleError(err, "Failed to collect bundle content")
		return err
	}

	// Show dry-run information if requested
	if bundleDryRun {
		return showBundleScope(output, bundleContent)
	}

	// Generate bundle filename
	bundleFilename := generateBundleFilename(bundleOutput, selectedAgents)
	bundleDir := filepath.Join(currentDir, ".krci-ai", "bundle")
	bundlePath := filepath.Join(bundleDir, bundleFilename)

	output.PrintInfo(fmt.Sprintf("Generating bundle: %s", bundlePath))

	// Create bundle directory
	if err := os.MkdirAll(bundleDir, 0755); err != nil {
		return fmt.Errorf("failed to create bundle directory: %w", err)
	}

	// Generate and write bundle
	bundleMarkdown := generateBundleMarkdown(bundleContent)

	if err := writeBundleFile(bundlePath, bundleMarkdown); err != nil {
		errorHandler.HandleError(err, "Failed to write bundle file")
		return err
	}

	// Success
	output.PrintSuccess("Bundle generated successfully!")
	output.PrintInfo(fmt.Sprintf("Bundle file: %s", bundlePath))
	output.PrintInfo(fmt.Sprintf("Bundle size: %d bytes", len(bundleMarkdown)))

	// Show usage instructions
	output.PrintInfo("\nUsage instructions:")
	output.PrintInfo("• Copy the entire bundle content to your web chat tool (ChatGPT, Claude Web, Gemini Pro)")
	output.PrintInfo("• The bundle includes all agents, tasks, templates, and project-specific data")
	output.PrintInfo("• Each section is clearly separated with collision-resistant delimiters")

	return nil
}

// runFrameworkValidation runs the existing validation logic
func runFrameworkValidation(currentDir string) error {
	// Initialize enhanced validation system (reuse existing validation)
	analyzer := validation.NewFrameworkAnalyzer(currentDir)

	// Run optimized framework analysis with caching
	issues, _, err := analyzer.OptimizedAnalyzeFramework()
	if err != nil {
		return fmt.Errorf("framework analysis failed: %w", err)
	}

	// Check for critical issues that would prevent bundle generation
	for _, issue := range issues {
		if issue.Severity == validation.SeverityCritical {
			return fmt.Errorf("critical validation error: %s", issue.Message)
		}
	}

	return nil
}

// FilterAgentsByNames filters agents by their names (case-insensitive), supporting both full names and file names
func FilterAgentsByNames(agentDeps []assets.AgentDependencyInfo, selectedAgents []string) []assets.AgentDependencyInfo {
	if len(selectedAgents) == 0 {
		return agentDeps
	}

	// Create case-insensitive lookup map
	selectedMap := make(map[string]bool)
	for _, name := range selectedAgents {
		selectedMap[strings.ToLower(name)] = true
	}

	var filtered []assets.AgentDependencyInfo
	for _, agent := range agentDeps {
		// Check both full name and file name
		fullNameMatch := selectedMap[strings.ToLower(agent.Name)]
		fileName := strings.TrimSuffix(filepath.Base(agent.FilePath), ".yaml")
		fileNameMatch := selectedMap[strings.ToLower(fileName)]

		if fullNameMatch || fileNameMatch {
			filtered = append(filtered, agent)
		}
	}

	return filtered
}

// collectBundleContent discovers and collects framework content, optionally filtered by agent names
func collectBundleContent(currentDir string, selectedAgents []string) (*BundleContent, error) {
	// Create discovery service
	discovery := assets.NewDiscovery(currentDir, GetEmbeddedAssets())

	// Discover agents with dependencies
	agentDeps, err := discovery.DiscoverAgentsWithDependencies()
	if err != nil {
		return nil, fmt.Errorf("failed to discover agents: %w", err)
	}

	// Filter agents if specific agents are selected
	if len(selectedAgents) > 0 {
		agentDeps = FilterAgentsByNames(agentDeps, selectedAgents)
	}

	content := &BundleContent{
		Agents:    make([]AgentBundleInfo, 0),
		Tasks:     make(map[string]string),
		Templates: make(map[string]string),
		DataFiles: make(map[string]string),
	}

	// Collect agent files
	for _, agent := range agentDeps {
		agentContent, err := os.ReadFile(agent.FilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read agent file %s: %w", agent.FilePath, err)
		}

		relPath, _ := filepath.Rel(currentDir, agent.FilePath)
		content.Agents = append(content.Agents, AgentBundleInfo{
			FilePath: relPath,
			Content:  string(agentContent),
			Name:     agent.Name,
			Role:     agent.Role,
		})

		// Collect task files
		for _, taskPath := range agent.Tasks {
			fullTaskPath := filepath.Join(currentDir, ".krci-ai", "tasks", taskPath)
			if taskContent, err := os.ReadFile(fullTaskPath); err == nil {
				relTaskPath := filepath.Join("tasks", taskPath)
				content.Tasks[relTaskPath] = string(taskContent)
			}
		}

		// Collect template files
		for _, templatePath := range agent.Templates {
			fullTemplatePath := filepath.Join(currentDir, ".krci-ai", "templates", templatePath)
			if templateContent, err := os.ReadFile(fullTemplatePath); err == nil {
				relTemplatePath := filepath.Join("templates", templatePath)
				content.Templates[relTemplatePath] = string(templateContent)
			}
		}

		// Collect data files
		for _, dataPath := range agent.DataFiles {
			fullDataPath := filepath.Join(currentDir, ".krci-ai", "data", dataPath)
			if dataContent, err := os.ReadFile(fullDataPath); err == nil {
				relDataPath := filepath.Join("data", dataPath)
				content.DataFiles[relDataPath] = string(dataContent)
			}
		}
	}

	// Collect any additional templates and data files not referenced by agents
	if err := collectAdditionalFiles(currentDir, content); err != nil {
		return nil, fmt.Errorf("failed to collect additional files: %w", err)
	}

	return content, nil
}

// collectAdditionalFiles collects templates and data files not referenced by agents
func collectAdditionalFiles(currentDir string, content *BundleContent) error {
	// Collect additional templates
	templatesDir := filepath.Join(currentDir, ".krci-ai", "templates")
	if templateFiles, err := filepath.Glob(filepath.Join(templatesDir, "*.md")); err == nil {
		for _, templateFile := range templateFiles {
			relPath, _ := filepath.Rel(currentDir, templateFile)
			relPath = strings.TrimPrefix(relPath, ".krci-ai/")

			if _, exists := content.Templates[relPath]; !exists {
				if templateContent, err := os.ReadFile(templateFile); err == nil {
					content.Templates[relPath] = string(templateContent)
				}
			}
		}
	}

	// Collect additional data files
	dataDir := filepath.Join(currentDir, ".krci-ai", "data")
	if err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip files with errors
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			relPath, _ := filepath.Rel(currentDir, path)
			relPath = strings.TrimPrefix(relPath, ".krci-ai/")

			if _, exists := content.DataFiles[relPath]; !exists {
				if dataContent, err := os.ReadFile(path); err == nil {
					content.DataFiles[relPath] = string(dataContent)
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// showBundleScope displays bundle scope information for dry-run
func showBundleScope(output *cli.OutputHandler, content *BundleContent) error {
	output.PrintSuccess("Bundle Scope Analysis:")
	output.PrintInfo(fmt.Sprintf("  Agents: %d", len(content.Agents)))
	output.PrintInfo(fmt.Sprintf("  Tasks: %d", len(content.Tasks)))
	output.PrintInfo(fmt.Sprintf("  Templates: %d", len(content.Templates)))
	output.PrintInfo(fmt.Sprintf("  Data Files: %d", len(content.DataFiles)))

	fmt.Println()
	output.PrintInfo("Agent Files:")
	for _, agent := range content.Agents {
		fmt.Printf("  • %s (%s - %s)\n", agent.FilePath, agent.Name, agent.Role)
	}

	if len(content.Tasks) > 0 {
		fmt.Println()
		output.PrintInfo("Task Files:")
		for taskPath := range content.Tasks {
			fmt.Printf("  • %s\n", taskPath)
		}
	}

	if len(content.Templates) > 0 {
		fmt.Println()
		output.PrintInfo("Template Files:")
		for templatePath := range content.Templates {
			fmt.Printf("  • %s\n", templatePath)
		}
	}

	if len(content.DataFiles) > 0 {
		fmt.Println()
		output.PrintInfo("Data Files:")
		for dataPath := range content.DataFiles {
			fmt.Printf("  • %s\n", dataPath)
		}
	}

	fmt.Println()
	output.PrintInfo("Use 'krci-ai bundle --all' to generate the actual bundle file")

	return nil
}

// generateBundleFilename creates the bundle filename
func generateBundleFilename(customOutput string, selectedAgents []string) string {
	if customOutput != "" {
		// Ensure .md extension
		if !strings.HasSuffix(customOutput, ".md") {
			customOutput += ".md"
		}
		return customOutput
	}

	// Generate filename based on selected agents
	if len(selectedAgents) > 0 {
		// Sort agent names alphabetically for consistent naming
		sortedAgents := make([]string, len(selectedAgents))
		copy(sortedAgents, selectedAgents)

		// Convert to lowercase for consistent filenames
		for i, agent := range sortedAgents {
			sortedAgents[i] = strings.ToLower(agent)
		}

		// Sort alphabetically
		for i := 0; i < len(sortedAgents); i++ {
			for j := i + 1; j < len(sortedAgents); j++ {
				if sortedAgents[i] > sortedAgents[j] {
					sortedAgents[i], sortedAgents[j] = sortedAgents[j], sortedAgents[i]
				}
			}
		}

		return strings.Join(sortedAgents, "-") + ".md"
	}

	return "all.md"
}

// generateBundleMarkdown creates the complete bundle markdown content
func generateBundleMarkdown(content *BundleContent) string {
	var result strings.Builder

	// Add bundle header with usage instructions
	addBundleHeader(&result)

	// Add agents with their dependencies grouped together
	for _, agent := range content.Agents {
		addAgentSection(&result, agent, content)
	}

	// Add shared templates not associated with specific agents
	addSharedTemplates(&result, content)

	// Add shared data files
	addSharedDataFiles(&result, content)

	return result.String()
}

// addBundleHeader adds the comprehensive bundle header with usage instructions
func addBundleHeader(result *strings.Builder) {
	timestamp := time.Now().Format("2006-01-02 15:04:05 MST")

	result.WriteString("# KubeRocketAI Framework Bundle\n\n")
	result.WriteString("**Generated:** " + timestamp + "\n")
	result.WriteString("**Purpose:** Complete framework bundle for web chat tools (ChatGPT, Claude Web, Gemini Pro)\n\n")

	result.WriteString("## Usage Instructions\n\n")
	result.WriteString("This bundle contains all KubeRocketAI framework components in a single file:\n")
	result.WriteString("- **Agent Definitions:** 6 SDLC roles with complete specifications\n")
	result.WriteString("- **Task Templates:** Workflow templates for common development tasks\n")
	result.WriteString("- **Output Templates:** Consistent formatting templates\n")
	result.WriteString("- **Reference Data:** Coding standards and best practices\n\n")

	result.WriteString("### File Format Guide\n")
	result.WriteString("- Each file section starts with `==== FILE: <path> ====`\n")
	result.WriteString("- Original file content follows with preserved formatting\n")
	result.WriteString("- Each file section ends with `==== END FILE ====`\n\n")

	result.WriteString("### For LLM Understanding\n")
	result.WriteString("When working with this bundle:\n")
	result.WriteString("1. Each agent represents a specific SDLC role (PM, Architect, Developer, QA, BA, PO)\n")
	result.WriteString("2. Tasks are workflow templates that agents can execute\n")
	result.WriteString("3. Templates provide consistent output formatting\n")
	result.WriteString("4. Data files contain project-specific standards and references\n\n")

	result.WriteString("---\n\n")
}

// addAgentSection adds an agent section with its related tasks grouped together
func addAgentSection(result *strings.Builder, agent AgentBundleInfo, content *BundleContent) {
	// Add agent file
	fmt.Fprintf(result, "==== FILE: %s ====\n", agent.FilePath)
	result.WriteString(optimizeContent(agent.Content))
	result.WriteString("\n==== END FILE ====\n\n")

	// Add agent-specific tasks immediately after the agent definition
	// This improves LLM context by keeping related content together
	agentTasks := getAgentTasks(agent.FilePath, content)
	for _, taskPath := range agentTasks {
		if taskContent, exists := content.Tasks[taskPath]; exists {
			fmt.Fprintf(result, "==== FILE: %s ====\n", taskPath)
			result.WriteString(optimizeContent(taskContent))
			result.WriteString("\n==== END FILE ====\n\n")
		}
	}
}

// getAgentTasks extracts task references from agent YAML content
func getAgentTasks(agentFilePath string, content *BundleContent) []string {
	// Simple extraction - look for tasks in the agent content
	var tasks []string

	// Find the agent content
	for _, agent := range content.Agents {
		if agent.FilePath == agentFilePath {
			// Parse tasks from YAML content (simple string matching)
			lines := strings.Split(agent.Content, "\n")
			inTasksSection := false

			for _, line := range lines {
				trimmed := strings.TrimSpace(line)
				if strings.Contains(trimmed, "tasks:") {
					inTasksSection = true
					continue
				}

				if inTasksSection {
					// Check if we've left the tasks section
					if strings.HasPrefix(trimmed, "commands:") || strings.HasPrefix(trimmed, "principles:") {
						break
					}

					// Extract task path
					if taskRef, found := strings.CutPrefix(trimmed, "- "); found {
						taskRef = strings.TrimSpace(taskRef)

						// Convert to relative path format
						if strings.HasPrefix(taskRef, "./.krci-ai/tasks/") {
							taskPath := strings.TrimPrefix(taskRef, "./.krci-ai/")
							tasks = append(tasks, taskPath)
						}
					}
				}
			}
			break
		}
	}

	return tasks
}

// addSharedTemplates adds template files not associated with specific agents
func addSharedTemplates(result *strings.Builder, content *BundleContent) {
	if len(content.Templates) > 0 {
		result.WriteString("# Shared Templates\n\n")

		for templatePath, templateContent := range content.Templates {
			fmt.Fprintf(result, "==== FILE: %s ====\n", templatePath)
			result.WriteString(optimizeContent(templateContent))
			result.WriteString("\n==== END FILE ====\n\n")
		}
	}
}

// addSharedDataFiles adds data files
func addSharedDataFiles(result *strings.Builder, content *BundleContent) {
	if len(content.DataFiles) > 0 {
		result.WriteString("# Reference Data\n\n")

		for dataPath, dataContent := range content.DataFiles {
			fmt.Fprintf(result, "==== FILE: %s ====\n", dataPath)
			result.WriteString(optimizeContent(dataContent))
			result.WriteString("\n==== END FILE ====\n\n")
		}
	}
}

// optimizeContent optimizes file content by removing excessive blank lines while preserving readability
func optimizeContent(content string) string {
	lines := strings.Split(content, "\n")
	var optimizedLines []string
	consecutiveEmptyLines := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			consecutiveEmptyLines++
			// Only allow up to 1 consecutive empty line
			if consecutiveEmptyLines <= 1 {
				optimizedLines = append(optimizedLines, line)
			}
		} else {
			consecutiveEmptyLines = 0
			optimizedLines = append(optimizedLines, line)
		}
	}

	return strings.Join(optimizedLines, "\n")
}

// writeBundleFile writes the bundle content to file
func writeBundleFile(bundlePath, content string) error {
	file, err := os.Create(bundlePath)
	if err != nil {
		return fmt.Errorf("failed to create bundle file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			// Log the error but don't override the main error
			fmt.Fprintf(os.Stderr, "Warning: failed to close bundle file: %v\n", closeErr)
		}
	}()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write bundle content: %w", err)
	}

	return nil
}
