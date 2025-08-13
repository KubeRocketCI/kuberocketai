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
	"slices"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/cli"
	"github.com/KubeRocketCI/kuberocketai/internal/validation"
)

// Constants
const (
	krciAIDir = ".krci-ai"
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
  krci-ai bundle --agent pm --task create-prd    # Generate minimal bundle for PM agent with create-prd task
  krci-ai bundle --agent pm --output pm.md       # Generate PM-only bundle with custom filename
  krci-ai bundle --agent "pm architect"          # Space-separated agent names
  krci-ai bundle --agent pm --task update-prd --output custom.md  # Minimal bundle with custom filename
  krci-ai bundle --agent pm,architect --dry-run  # Show targeted bundle scope without generating files
  krci-ai bundle --agent pm --task create-prd --dry-run  # Show minimal bundle scope without generating files
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
	bundleTask   string
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
	bundleCmd.Flags().StringVar(&bundleTask, "task", "", "Generate minimal bundle with specific agent and task (requires --agent flag)")
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

// validateBundleFlags validates the bundle command flags for mutual exclusivity and requirements
func validateBundleFlags(output *cli.OutputHandler) error {
	// Validate flags - ensure mutual exclusivity and required combinations
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

	if bundleAll && bundleTask != "" {
		output.PrintError("Cannot specify both --all and --task flags")
		output.PrintInfo("Use --task only with --agent for minimal single-task bundles")
		return fmt.Errorf("cannot specify both --all and --task flags")
	}

	if bundleTask != "" && bundleAgents == "" {
		output.PrintError("--task flag requires --agent flag")
		output.PrintInfo("Usage: krci-ai bundle --agent pm --task create-prd")
		return fmt.Errorf("--task flag requires --agent flag")
	}

	if bundleTask != "" && len(ParseAgentList(bundleAgents)) > 1 {
		output.PrintError("--task flag requires exactly one agent")
		output.PrintInfo("Usage: krci-ai bundle --agent pm --task create-prd (single agent only)")
		return fmt.Errorf("--task flag requires exactly one agent")
	}

	return nil
}

// runBundle executes the bundle command
func runBundle(cmd *cobra.Command, args []string) error {
	// Create output and error handlers
	output := cli.NewOutputHandler()
	errorHandler := cli.NewErrorHandler()

	// Validate flags
	if err := validateBundleFlags(output); err != nil {
		return err
	}

	// Get current working directory and validate framework
	currentDir, err := setupAndValidateFramework(output, errorHandler)
	if err != nil {
		return err
	}

	// Parse and validate agent selection
	selectedAgents, err := parseAndValidateAgents(currentDir, output)
	if err != nil {
		return err
	}

	// Collect bundle content
	output.PrintProgress("Discovering agents and dependencies...")
	bundleContent, err := collectBundleContent(currentDir, selectedAgents, bundleTask)
	if err != nil {
		errorHandler.HandleError(err, "Failed to collect bundle content")
		return err
	}

	// Show dry-run information if requested
	if bundleDryRun {
		return showBundleScope(output, bundleContent)
	}

	// Generate and write bundle
	return generateAndWriteBundle(currentDir, selectedAgents, bundleContent, output, errorHandler)
}

// setupAndValidateFramework handles directory setup and framework validation
func setupAndValidateFramework(output *cli.OutputHandler, errorHandler *cli.ErrorHandler) (string, error) {
	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	// Check if framework is installed
	frameworkDir := filepath.Join(currentDir, krciAIDir)
	if _, statErr := os.Stat(frameworkDir); os.IsNotExist(statErr) {
		output.PrintError("No " + krciAIDir + " directory found in current directory")
		output.PrintInfo("Run 'krci-ai install' to set up the framework first")
		return "", fmt.Errorf("framework not installed")
	}

	output.PrintProgress("Validating framework integrity...")

	// Run framework validation before bundle generation
	if validationErr := runFrameworkValidation(currentDir); validationErr != nil {
		errorHandler.HandleError(validationErr, "Framework validation failed")
		output.PrintError("Bundle generation requires a valid framework")
		output.PrintInfo("Fix validation errors and try again")
		return "", validationErr
	}

	output.PrintSuccess("Framework validation passed")
	return currentDir, nil
}

// parseAndValidateAgents handles agent parsing and validation logic
func parseAndValidateAgents(currentDir string, output *cli.OutputHandler) ([]string, error) {
	var selectedAgents []string

	if bundleAgents == "" {
		return selectedAgents, nil
	}

	selectedAgents = ParseAgentList(bundleAgents)
	if len(selectedAgents) == 0 {
		output.PrintError("No valid agents specified")
		return nil, fmt.Errorf("no valid agents specified")
	}

	// Validate agent names exist
	if validateErr := validateAgentNames(currentDir, selectedAgents, output); validateErr != nil {
		return nil, validateErr
	}

	// Handle task validation and output
	return handleTaskValidation(currentDir, selectedAgents, output)
}

// handleTaskValidation handles task-specific validation and info output
func handleTaskValidation(currentDir string, selectedAgents []string, output *cli.OutputHandler) ([]string, error) {
	if bundleTask != "" {
		if validateErr := ValidateAgentTaskCombination(currentDir, selectedAgents[0], bundleTask, output); validateErr != nil {
			return nil, validateErr
		}
		output.PrintInfo(fmt.Sprintf("Selected agent-task combination: %s - %s", selectedAgents[0], bundleTask))
	} else {
		output.PrintInfo(fmt.Sprintf("Selected agents: %s", strings.Join(selectedAgents, ", ")))
	}

	return selectedAgents, nil
}

// generateAndWriteBundle handles bundle generation and file writing
func generateAndWriteBundle(currentDir string, selectedAgents []string, bundleContent *BundleContent, output *cli.OutputHandler, errorHandler *cli.ErrorHandler) error {
	// Generate bundle filename
	bundleFilename := generateBundleFilename(bundleOutput, selectedAgents, bundleTask)
	bundleDir := filepath.Join(currentDir, krciAIDir, "bundle")
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

// extractFileReference extracts a file reference from a line starting at a specific prefix
func extractFileReference(line, prefix string) string {
	start := strings.Index(line, prefix)
	if start == -1 {
		return ""
	}

	remaining := line[start+len(prefix):]
	end := len(remaining)
	for i, char := range remaining {
		if char == ' ' || char == ')' || char == ']' || char == ',' {
			end = i
			break
		}
	}
	return remaining[:end]
}

// getTaskDependencies analyzes a task file to extract its dependencies
func getTaskDependencies(currentDir, taskName string) []string {
	taskFilePath := filepath.Join(currentDir, krciAIDir, "tasks", taskName+".md")
	taskContent, err := os.ReadFile(taskFilePath)
	if err != nil {
		return []string{}
	}

	var dependencies []string
	content := string(taskContent)

	// Look for file references in the task content
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		// Find references to data files
		if strings.Contains(line, "./"+krciAIDir+"/data/") {
			dep := extractFileReference(line, "./"+krciAIDir+"/data/")
			if dep != "" && !slices.Contains(dependencies, dep) {
				dependencies = append(dependencies, dep)
			}
		}

		// Find references to template files
		if strings.Contains(line, "./"+krciAIDir+"/templates/") {
			dep := extractFileReference(line, "./"+krciAIDir+"/templates/")
			if dep != "" && !slices.Contains(dependencies, dep) {
				dependencies = append(dependencies, dep)
			}
		}
	}

	return dependencies
}

// isReferencedByTask checks if a file path is referenced by the task dependencies
func isReferencedByTask(filePath string, taskDeps []string) bool {
	fileName := filepath.Base(filePath)
	for _, dep := range taskDeps {
		if strings.Contains(dep, fileName) || strings.Contains(filePath, dep) {
			return true
		}
	}
	// If no specific dependencies found, include common files
	return len(taskDeps) == 0
}

// collectTemplatesAndData collects template and data files for an agent, optionally filtered by task
func collectTemplatesAndData(currentDir string, agent assets.AgentDependencyInfo, taskName string, content *BundleContent) {
	if taskName != "" {
		// Get task-specific dependencies by parsing the task file content
		taskDeps := getTaskDependencies(currentDir, taskName)

		// Only include templates and data referenced by the specific task
		for _, templatePath := range agent.Templates {
			if isReferencedByTask(templatePath, taskDeps) {
				fullTemplatePath := filepath.Join(currentDir, krciAIDir, "templates", templatePath)
				if templateContent, err := os.ReadFile(fullTemplatePath); err == nil {
					relTemplatePath := filepath.Join("templates", templatePath)
					content.Templates[relTemplatePath] = string(templateContent)
				}
			}
		}

		for _, dataPath := range agent.DataFiles {
			if isReferencedByTask(dataPath, taskDeps) {
				fullDataPath := filepath.Join(currentDir, krciAIDir, "data", dataPath)
				if dataContent, err := os.ReadFile(fullDataPath); err == nil {
					relDataPath := filepath.Join("data", dataPath)
					content.DataFiles[relDataPath] = string(dataContent)
				}
			}
		}
	} else {
		// Include all templates and data for non-task-specific bundles
		for _, templatePath := range agent.Templates {
			fullTemplatePath := filepath.Join(currentDir, krciAIDir, "templates", templatePath)
			if templateContent, err := os.ReadFile(fullTemplatePath); err == nil {
				relTemplatePath := filepath.Join("templates", templatePath)
				content.Templates[relTemplatePath] = string(templateContent)
			}
		}

		for _, dataPath := range agent.DataFiles {
			fullDataPath := filepath.Join(currentDir, krciAIDir, "data", dataPath)
			if dataContent, err := os.ReadFile(fullDataPath); err == nil {
				relDataPath := filepath.Join("data", dataPath)
				content.DataFiles[relDataPath] = string(dataContent)
			}
		}
	}
}

// ValidateAgentTaskCombination validates that the specified task exists for the specified agent
func ValidateAgentTaskCombination(currentDir, agentName, taskName string, output *cli.OutputHandler) error {
	// Create discovery service
	discovery := assets.NewDiscovery(currentDir, GetEmbeddedAssets())

	// Get available agents with full info
	agentInfos, err := discovery.DiscoverAgents()
	if err != nil {
		return fmt.Errorf("failed to discover available agents: %w", err)
	}

	// Find the specified agent
	var selectedAgent *assets.AgentInfo
	for _, agent := range agentInfos {
		fileName := strings.TrimSuffix(filepath.Base(agent.FilePath), ".yaml")
		if strings.EqualFold(agent.Name, agentName) || strings.EqualFold(fileName, agentName) {
			selectedAgent = &agent
			break
		}
	}

	if selectedAgent == nil {
		output.PrintError(fmt.Sprintf("Agent '%s' not found", agentName))
		return fmt.Errorf("agent '%s' not found", agentName)
	}

	// Parse the agent YAML file to get tasks
	agentContent, err := os.ReadFile(selectedAgent.FilePath)
	if err != nil {
		return fmt.Errorf("failed to read agent file %s: %w", selectedAgent.FilePath, err)
	}

	// Parse YAML structure to extract tasks
	var agentYAML struct {
		Agent struct {
			Tasks []string `yaml:"tasks"`
		} `yaml:"agent"`
	}

	if err := yaml.Unmarshal(agentContent, &agentYAML); err != nil {
		return fmt.Errorf("failed to parse agent YAML: %w", err)
	}

	// Check if the specified task exists for this agent
	taskFound := false
	var availableTasks []string

	for _, taskPath := range agentYAML.Agent.Tasks {
		// Extract task name from path (e.g., "./.krci-ai/tasks/create-prd.md" -> "create-prd")
		taskFileName := strings.TrimSuffix(filepath.Base(taskPath), ".md")
		availableTasks = append(availableTasks, taskFileName)

		if strings.EqualFold(taskFileName, taskName) {
			taskFound = true
		}
	}

	if !taskFound {
		output.PrintError(fmt.Sprintf("Task '%s' not found for agent '%s'", taskName, agentName))
		output.PrintInfo(fmt.Sprintf("Available tasks for %s:", selectedAgent.Name))
		for _, task := range availableTasks {
			output.PrintInfo(fmt.Sprintf("  • %s", task))
		}
		return fmt.Errorf("task '%s' not found for agent '%s'", taskName, agentName)
	}

	// Verify the task file actually exists
	taskFilePath := filepath.Join(currentDir, krciAIDir, "tasks", taskName+".md")
	if _, err := os.Stat(taskFilePath); os.IsNotExist(err) {
		output.PrintError(fmt.Sprintf("Task file '%s' does not exist", taskFilePath))
		return fmt.Errorf("task file '%s' does not exist", taskFilePath)
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

// collectBundleContent discovers and collects framework content, optionally filtered by agent names and task
func collectBundleContent(currentDir string, selectedAgents []string, taskName string) (*BundleContent, error) {
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

		// Collect task files - filter by specific task if specified
		for _, taskPath := range agent.Tasks {
			// If a specific task is requested, only include that task
			if taskName != "" {
				taskFileName := strings.TrimSuffix(filepath.Base(taskPath), ".md")
				if !strings.EqualFold(taskFileName, taskName) {
					continue // Skip tasks that don't match the specified task
				}
			}

			fullTaskPath := filepath.Join(currentDir, krciAIDir, "tasks", taskPath)
			if taskContent, err := os.ReadFile(fullTaskPath); err == nil {
				relTaskPath := filepath.Join("tasks", taskPath)
				content.Tasks[relTaskPath] = string(taskContent)
			}
		}

		// Collect template and data files using helper function
		collectTemplatesAndData(currentDir, agent, taskName, content)
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
	templatesDir := filepath.Join(currentDir, krciAIDir, "templates")
	if templateFiles, err := filepath.Glob(filepath.Join(templatesDir, "*.md")); err == nil {
		for _, templateFile := range templateFiles {
			relPath, _ := filepath.Rel(currentDir, templateFile)
			relPath = strings.TrimPrefix(relPath, krciAIDir+"/")

			if _, exists := content.Templates[relPath]; !exists {
				if templateContent, err := os.ReadFile(templateFile); err == nil {
					content.Templates[relPath] = string(templateContent)
				}
			}
		}
	}

	// Collect additional data files
	dataDir := filepath.Join(currentDir, krciAIDir, "data")
	if err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip files with errors
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			relPath, _ := filepath.Rel(currentDir, path)
			relPath = strings.TrimPrefix(relPath, krciAIDir+"/")

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
func generateBundleFilename(customOutput string, selectedAgents []string, taskName string) string {
	if customOutput != "" {
		// Ensure .md extension
		if !strings.HasSuffix(customOutput, ".md") {
			customOutput += ".md"
		}
		return customOutput
	}

	// Generate filename based on selected agents and task
	if len(selectedAgents) > 0 {
		// If both agent and task are specified, use agent-task pattern
		if taskName != "" && len(selectedAgents) == 1 {
			agentName := strings.ToLower(selectedAgents[0])
			taskNameLower := strings.ToLower(taskName)
			return fmt.Sprintf("%s-%s.md", agentName, taskNameLower)
		}

		// For multiple agents or agent-only bundles, use existing logic
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
						if strings.HasPrefix(taskRef, "./"+krciAIDir+"/tasks/") {
							taskPath := strings.TrimPrefix(taskRef, "./"+krciAIDir+"/")
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
