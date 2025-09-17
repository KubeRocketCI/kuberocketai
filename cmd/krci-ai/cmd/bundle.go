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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/bundle"
	"github.com/KubeRocketCI/kuberocketai/internal/cli"
	"github.com/KubeRocketCI/kuberocketai/internal/discovery"
	"github.com/KubeRocketCI/kuberocketai/internal/tokens"
)

// Constants
const (
	bundleDir      = "bundle"
	dirPermissions = 0755
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

// BundleContent represents the aggregated content for bundle generation
type BundleContent struct {
	Agents    []assets.Agent
	Tasks     map[string]string
	Templates map[string]string
	DataFiles map[string]string
}

func init() {
	rootCmd.AddCommand(bundleCmd)

	// Add flags
	bundleCmd.Flags().Bool("all", false, "Generate complete bundle with all agents and dependencies")
	bundleCmd.Flags().String("agent", "", "Generate targeted bundle with specific agents (comma or space separated: 'pm,architect' or 'pm architect')")
	bundleCmd.Flags().Bool("dry-run", false, "Show bundle scope without generating files")
	bundleCmd.Flags().String("output", "", "Custom output filename (creates ./.krci-ai/bundle/filename)")
}

// ParseAgentList parses comma-separated or space-separated agent names
func ParseAgentList(agentStr string) []string {
	if agentStr == "" {
		return nil
	}

	var agents []string

	// First try comma-separated
	if strings.Contains(agentStr, ",") {
		for part := range strings.SplitSeq(agentStr, ",") {
			agent := strings.TrimSpace(part)
			if agent != "" {
				agents = append(agents, agent)
			}
		}

		return agents
	}

	// Space-separated
	for part := range strings.FieldsSeq(agentStr) {
		agents = append(agents, strings.TrimSpace(part))
	}

	return agents
}

// validateAgentNames validates that specified agent names exist in the framework.
func validateAgentNames(ctx context.Context, discovery *assets.Discovery, selectedAgents []string, output *cli.OutputHandler) error {
	if len(selectedAgents) == 0 {
		return nil
	}

	// Get available agents with full info
	agents, err := discovery.GetAgents(ctx)
	if err != nil {
		return fmt.Errorf("failed to get available agents: %w", err)
	}

	// Create map for quick lookup (case-insensitive) - support both full names and short names
	agentMap := make(map[string]string)

	for _, agent := range agents {
		// Add full agent name
		fullNameLower := strings.ToLower(agent.Name)
		agentMap[fullNameLower] = agent.Name

		// Add short name
		shortNameLower := strings.ToLower(agent.ShortName)
		agentMap[shortNameLower] = agent.Name
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
		for _, agent := range agents {
			output.PrintInfo(fmt.Sprintf("  • %s (short: %s)", agent.Name, agent.ShortName))
		}
		return fmt.Errorf("invalid agent names specified")
	}

	return nil
}

// validateBundleFlags validates the bundle command flags for mutual exclusivity and requirements
func validateBundleFlags(cmd *cobra.Command, output *cli.OutputHandler) error {
	// Get flag values
	bundleAll, err := cmd.Flags().GetBool("all")
	if err != nil {
		return fmt.Errorf("failed to read all flag: %w", err)
	}

	bundleAgents, err := cmd.Flags().GetString("agent")
	if err != nil {
		return fmt.Errorf("failed to read agent flag: %w", err)
	}

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

	return nil
}

// runBundle executes the bundle command
func runBundle(cmd *cobra.Command, args []string) error {
	// Create output and error handlers
	output := cli.NewOutputHandler()
	errorHandler := cli.NewErrorHandler()

	// Validate flags
	if err := validateBundleFlags(cmd, output); err != nil {
		return err
	}

	projectRoot, err := discovery.GetProjectRoot()
	if err != nil {
		errorHandler.HandleError(err, "Failed to get project root")
		return err
	}

	// Create discovery service
	discovery := assets.NewEmbeddedDiscovery(GetEmbeddedAssets(), assets.EmbeddedPrefix)

	// Parse and validate agent selection
	selectedAgents, err := parseAndValidateAgents(cmd, discovery, output)
	if err != nil {
		return err
	}

	// Collect bundle content
	output.PrintProgress("Discovering agents and dependencies...")
	bundleContent, err := collectBundleContent(cmd.Context(), discovery, selectedAgents)
	if err != nil {
		errorHandler.HandleError(err, "Failed to collect bundle content")
		return err
	}

	// Check dry-run flag
	bundleDryRun, err := cmd.Flags().GetBool("dry-run")
	if err != nil {
		return fmt.Errorf("failed to read dry-run flag: %w", err)
	}

	// Show dry-run information if requested
	if bundleDryRun {
		return showBundleScope(output, bundleContent)
	}

	// Get output flag for bundle generation
	bundleOutput, err := cmd.Flags().GetString("output")
	if err != nil {
		return fmt.Errorf("failed to read output flag: %w", err)
	}

	// Generate and write bundle
	return generateAndWriteBundle(cmd.Context(), projectRoot, selectedAgents, bundleContent, bundleOutput, discovery, output, errorHandler)
}

// parseAndValidateAgents handles agent parsing and validation logic
func parseAndValidateAgents(cmd *cobra.Command, discovery *assets.Discovery, output *cli.OutputHandler) ([]string, error) {
	var selectedAgents []string

	// Get agent flag
	bundleAgents, err := cmd.Flags().GetString("agent")
	if err != nil {
		return nil, fmt.Errorf("failed to read agent flag: %w", err)
	}

	if bundleAgents == "" {
		return selectedAgents, nil
	}

	selectedAgents = ParseAgentList(bundleAgents)
	if len(selectedAgents) == 0 {
		output.PrintError("No valid agents specified")
		return nil, fmt.Errorf("no valid agents specified")
	}

	// Validate agent names
	if err := validateAgentNames(cmd.Context(), discovery, selectedAgents, output); err != nil {
		return nil, err
	}

	return selectedAgents, nil
}

// collectBundleContent collects all content needed for bundle generation
func collectBundleContent(ctx context.Context, discovery *assets.Discovery, selectedAgents []string) (*BundleContent, error) {
	agents, err := getAgentsForBundle(ctx, discovery, selectedAgents)
	if err != nil {
		return nil, err
	}

	content := &BundleContent{
		Agents:    agents,
		Tasks:     make(map[string]string),
		Templates: make(map[string]string),
		DataFiles: make(map[string]string),
	}

	for _, agent := range agents {
		if err := collectAgentContent(discovery, agent, content); err != nil {
			return nil, err
		}
	}

	return content, nil
}

// getAgentsForBundle retrieves the appropriate agents based on selection
func getAgentsForBundle(ctx context.Context, discovery *assets.Discovery, selectedAgents []string) ([]assets.Agent, error) {
	if len(selectedAgents) == 0 {
		agents, err := discovery.GetAgents(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get agents: %w", err)
		}
		return agents, nil
	}

	agents, err := discovery.GetAgentsByNames(ctx, selectedAgents)
	if err != nil {
		return nil, fmt.Errorf("failed to get agents: %w", err)
	}
	return agents, nil
}

// collectAgentContent collects all content for a specific agent
func collectAgentContent(discovery *assets.Discovery, agent assets.Agent, content *BundleContent) error {
	if err := collectFilePaths(discovery, agent.GetAllTasksPaths(), content.Tasks); err != nil {
		return err
	}

	if err := collectFilePaths(discovery, agent.GetAllReferencedTasksPaths(), content.Tasks); err != nil {
		return err
	}

	if err := collectFilePaths(discovery, agent.GetAllTemplatesPaths(), content.Templates); err != nil {
		return err
	}

	if err := collectFilePaths(discovery, agent.GetAllDataFilesPaths(), content.DataFiles); err != nil {
		return err
	}

	return nil
}

// collectFilePaths collects content for a list of file paths into the target map
func collectFilePaths(discovery *assets.Discovery, paths []string, target map[string]string) error {
	for _, path := range paths {
		if _, exists := target[path]; !exists {
			content, err := readFileContent(discovery, path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", path, err)
			}
			target[path] = content
		}
	}
	return nil
}

// readFileContent reads file content using the discovery filesystem
func readFileContent(discovery *assets.Discovery, filePath string) (string, error) {
	data, err := discovery.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// generateAndWriteBundle handles bundle generation and file writing
func generateAndWriteBundle(
	ctx context.Context,
	projectRoot string,
	selectedAgents []string,
	bundleContent *BundleContent,
	bundleOutput string,
	discovery *assets.Discovery,
	output *cli.OutputHandler,
	errorHandler *cli.ErrorHandler,
) error {
	// Generate bundle filename
	bundleFilename := bundle.GenerateBundleFilename(bundleOutput, selectedAgents)
	bundleDirPath := filepath.Join(projectRoot, assets.KrciAIDir, bundleDir)
	bundlePath := filepath.Join(bundleDirPath, bundleFilename)

	output.PrintInfo(fmt.Sprintf("Generating bundle: %s", bundlePath))

	// Create bundle directory
	if err := os.MkdirAll(bundleDirPath, dirPermissions); err != nil {
		return fmt.Errorf("failed to create bundle directory: %w", err)
	}

	// Generate and write bundle
	bundleMarkdown := generateBundleMarkdown(bundleContent, discovery)

	if err := writeBundleFile(bundlePath, bundleMarkdown); err != nil {
		errorHandler.HandleError(err, "Failed to write bundle file")
		return err
	}

	// Success
	output.PrintSuccess("Bundle generated successfully!")
	output.PrintInfo(fmt.Sprintf("Bundle file: %s", bundlePath))
	output.PrintInfo(fmt.Sprintf("Bundle size: %d bytes", len(bundleMarkdown)))

	// Calculate and display token summary (with graceful error handling)
	if err := displayBundleTokenSummary(ctx, bundleMarkdown, output); err != nil {
		// Token calculation error should not prevent bundle generation success
		output.PrintWarning(fmt.Sprintf("Token analysis failed: %v", err))
		output.PrintInfo("Bundle was created successfully despite token calculation error")
		output.PrintInfo("Use 'krci-ai tokens --bundle <scope>' for detailed token analysis")
	}

	// Show usage instructions
	output.Newline()
	output.PrintInfo("Usage instructions:")
	output.PrintInfo("• Copy the entire bundle content to your web chat tool (ChatGPT, Claude Web, Gemini Pro)")
	output.PrintInfo("• The bundle includes all agents, tasks, templates, and project-specific data")
	output.PrintInfo("• Each section is clearly separated with collision-resistant delimiters")

	return nil
}

// showBundleScope displays bundle scope information for dry-run
func showBundleScope(output *cli.OutputHandler, content *BundleContent) error {
	output.PrintSuccess("Bundle Scope Analysis:")
	output.PrintInfo(fmt.Sprintf("  Agents: %d", len(content.Agents)))
	output.PrintInfo(fmt.Sprintf("  Tasks: %d", len(content.Tasks)))
	output.PrintInfo(fmt.Sprintf("  Templates: %d", len(content.Templates)))
	output.PrintInfo(fmt.Sprintf("  Data Files: %d", len(content.DataFiles)))

	output.Newline()
	output.PrintInfo("Agent Files:")
	for _, agent := range content.Agents {
		output.PrintInfo(fmt.Sprintf("  • %s (%s - %s)", agent.FilePath, agent.Name, agent.Role))
	}

	if len(content.Tasks) > 0 {
		output.Newline()
		output.PrintInfo("Task Files:")
		for taskPath := range content.Tasks {
			output.PrintInfo(fmt.Sprintf("  • %s", taskPath))
		}
	}

	if len(content.Templates) > 0 {
		output.Newline()
		output.PrintInfo("Template Files:")
		for templatePath := range content.Templates {
			output.PrintInfo(fmt.Sprintf("  • %s", templatePath))
		}
	}

	if len(content.DataFiles) > 0 {
		output.Newline()
		output.PrintInfo("Data Files:")
		for dataPath := range content.DataFiles {
			output.PrintInfo(fmt.Sprintf("  • %s", dataPath))
		}
	}

	output.Newline()
	output.PrintInfo("Use 'krci-ai bundle --all' to generate the actual bundle file")

	return nil
}

// generateBundleMarkdown creates the complete bundle markdown content
func generateBundleMarkdown(content *BundleContent, discovery *assets.Discovery) string {
	var result strings.Builder

	// Add bundle header with usage instructions
	addBundleHeader(&result)

	// Add agents with their dependencies grouped together
	for _, agent := range content.Agents {
		addAgentSection(&result, agent, content, discovery)
	}

	// Add shared templates not associated with specific agents
	addSharedTemplates(&result, content)

	// Add shared data files
	addSharedDataFiles(&result, content)

	return result.String()
}

// addBundleHeader adds the comprehensive bundle header with usage instructions
func addBundleHeader(result *strings.Builder) {
	timestamp := time.Now().Format(time.RFC3339)

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
func addAgentSection(result *strings.Builder, agent assets.Agent, content *BundleContent, discovery *assets.Discovery) {
	// Add agent file
	fmt.Fprintf(result, "==== FILE: %s ====\n", agent.FilePath)

	// Read agent file content
	agentContent, err := discovery.ReadFile(agent.FilePath)
	if err != nil {
		fmt.Fprintf(result, "Error reading agent file: %v\n", err)
	} else {
		result.WriteString(string(agentContent))
	}
	result.WriteString("\n==== END FILE ====\n\n")

	// Add agent-specific tasks immediately after the agent definition
	// This improves LLM context by keeping related content together
	agentTasks := agent.GetAllTasksPaths()
	for _, taskPath := range agentTasks {
		if taskContent, exists := content.Tasks[taskPath]; exists {
			fmt.Fprintf(result, "==== FILE: %s ====\n", taskPath)
			result.WriteString(taskContent)
			result.WriteString("\n==== END FILE ====\n\n")
		}
	}
}

// addSharedTemplates adds template files not associated with specific agents
func addSharedTemplates(result *strings.Builder, content *BundleContent) {
	if len(content.Templates) > 0 {
		result.WriteString("# Shared Templates\n\n")

		for templatePath, templateContent := range content.Templates {
			fmt.Fprintf(result, "==== FILE: %s ====\n", templatePath)
			result.WriteString(templateContent)
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
			result.WriteString(dataContent)
			result.WriteString("\n==== END FILE ====\n\n")
		}
	}
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

// displayBundleTokenSummary calculates and displays token summary for the generated bundle
func displayBundleTokenSummary(ctx context.Context, bundleContent string, output *cli.OutputHandler) error {
	engine, err := tokens.NewDefaultEngine()
	if err != nil {
		return fmt.Errorf("failed to create tokens calculator: %w", err)
	}

	totalTokens, err := engine.CalculateTokens(ctx, bundleContent)
	if err != nil {
		return fmt.Errorf("failed to calculate tokens: %w", err)
	}

	// Display token summary
	output.PrintInfo(fmt.Sprintf("Bundle contains %d tokens", totalTokens))

	return nil
}
