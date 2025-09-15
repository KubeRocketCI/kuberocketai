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
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/cli"
	"github.com/KubeRocketCI/kuberocketai/internal/discovery"
	"github.com/KubeRocketCI/kuberocketai/internal/tokens"
)

// output is initialized in init function for CLI formatting
var output *cli.OutputHandler

// tokensCmd represents the tokens command
var tokensCmd = &cobra.Command{
	Use:   "tokens",
	Short: "Analyze token usage for agents and project configurations",
	Long: `Analyze token usage for AI agent configurations using GPT-4 tokenization.

This command helps you understand token consumption patterns and avoid context 
limit violations by providing accurate token counts for agents and their dependencies.

Examples:
  # Analyze tokens for a specific agent
  krci-ai tokens --agent pm
  
  # Analyze tokens for all agents in the project
  krci-ai tokens --all
  
  # Analyze tokens for bundle configurations
  krci-ai tokens --bundle dev
  krci-ai tokens --bundle pm,architect
  krci-ai tokens --bundle all
  
  # Get detailed breakdown with JSON output
  krci-ai tokens --all --json --verbose
  
  # Set custom timeout for large projects
  krci-ai tokens --all --timeout 10s`,
	RunE: runTokensCommand,
}

func init() {
	// Initialize output handler
	output = cli.NewOutputHandler()

	// Add tokens command to root
	rootCmd.AddCommand(tokensCmd)

	// Define flags
	tokensCmd.Flags().String("agent", "", "Analyze tokens for a specific agent")
	tokensCmd.Flags().Bool("all", false, "Analyze tokens for all agents in the project")
	tokensCmd.Flags().String("bundle", "", "Analyze tokens for bundle configuration (agent names: 'pm', 'pm,architect', 'all')")
	tokensCmd.Flags().Bool("json", false, "Output results in JSON format")
	tokensCmd.Flags().Duration("timeout", 30*time.Second, "Timeout for token analysis")

	// Mark flags as mutually exclusive
	tokensCmd.MarkFlagsMutuallyExclusive("agent", "all", "bundle")
}

func runTokensCommand(cmd *cobra.Command, args []string) error {
	// Get flags
	tokenAgent, err := cmd.Flags().GetString("agent")
	if err != nil {
		return fmt.Errorf("failed to get agent flag: %w", err)
	}

	tokenAll, err := cmd.Flags().GetBool("all")
	if err != nil {
		return fmt.Errorf("failed to get all flag: %w", err)
	}

	tokenBundle, err := cmd.Flags().GetString("bundle")
	if err != nil {
		return fmt.Errorf("failed to get bundle flag: %w", err)
	}

	tokenJSON, err := cmd.Flags().GetBool("json")
	if err != nil {
		return fmt.Errorf("failed to get json flag: %w", err)
	}

	tokenTimeout, err := cmd.Flags().GetDuration("timeout")
	if err != nil {
		return fmt.Errorf("failed to get timeout flag: %w", err)
	}

	// Validate flags
	if tokenAgent == "" && !tokenAll && tokenBundle == "" {
		return fmt.Errorf("either --agent, --all, or --bundle flag must be specified")
	}

	projectRoot, err := discovery.GetProjectRoot()
	if err != nil {
		return handleTokenError(fmt.Errorf("failed to initialize token calculator: %w", err), tokenJSON)
	}

	// Create token calculator
	calculator, err := tokens.NewCalculator(assets.GetKrciPath(projectRoot))
	if err != nil {
		return handleTokenError(fmt.Errorf("failed to initialize token calculator: %w", err), tokenJSON)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), tokenTimeout)
	defer cancel()

	// Execute the appropriate command
	if tokenAgent != "" {
		return runAgentTokenAnalysis(ctx, calculator, tokenAgent, tokenJSON)
	} else if tokenBundle != "" {
		return runBundleTokenAnalysis(ctx, calculator, tokenBundle, tokenJSON)
	} else {
		return runProjectTokenAnalysis(ctx, calculator, tokenJSON)
	}
}

func runAgentTokenAnalysis(ctx context.Context, calculator *tokens.Calculator, agentName string, jsonOutput bool) error {
	// Show progress indicator
	if !jsonOutput {
		output.PrintProgress(fmt.Sprintf("Analyzing tokens for agent '%s'...", agentName))
	}

	// Calculate tokens
	agentInfo, err := calculator.CalculateAgentTokens(ctx, agentName)
	if err != nil {
		return handleTokenError(err, jsonOutput)
	}

	// Output results
	if jsonOutput {
		return outputAgentTokensJSON(agentInfo)
	} else {
		return outputAgentTokensTable(agentInfo)
	}
}

func runProjectTokenAnalysis(ctx context.Context, calculator *tokens.Calculator, jsonOutput bool) error {
	// Show progress indicator
	if !jsonOutput {
		output.PrintProgress("Analyzing tokens for entire project...")
	}

	// Calculate tokens
	projectInfo, err := calculator.CalculateAllTokens(ctx)
	if err != nil {
		return handleTokenError(err, jsonOutput)
	}

	// Output results
	if jsonOutput {
		return outputProjectTokensJSON(projectInfo)
	} else {
		return outputProjectTokensTable(projectInfo)
	}
}

func runBundleTokenAnalysis(ctx context.Context, calculator *tokens.Calculator, bundleScope string, jsonOutput bool) error {
	// Show progress indicator
	if !jsonOutput {
		output.PrintProgress(fmt.Sprintf("Analyzing tokens for bundle scope '%s'...", bundleScope))
	}

	// Parse bundle scope to get agent list
	var selectedAgents []string
	if bundleScope == "all" {
		selectedAgents = nil // nil means all agents
	} else {
		selectedAgents = ParseAgentList(bundleScope)
		if len(selectedAgents) == 0 {
			return handleTokenError(fmt.Errorf("no valid agents specified in bundle scope"), jsonOutput)
		}
	}

	// Calculate bundle tokens using the tokens package
	bundleTokenInfo, err := calculator.CalculateBundleTokens(ctx, selectedAgents)
	if err != nil {
		return handleTokenError(fmt.Errorf("bundle token analysis failed: %w", err), jsonOutput)
	}

	// Output results
	if jsonOutput {
		return outputBundleTokensJSON(bundleTokenInfo)
	} else {
		return outputBundleTokensTable(bundleTokenInfo, bundleScope)
	}
}

func outputAgentTokensJSON(agentInfo *tokens.AgentTokenInfo) error {
	output, err := json.MarshalIndent(agentInfo, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON output: %w", err)
	}

	fmt.Println(string(output))
	return nil
}

func outputProjectTokensJSON(projectInfo *tokens.ProjectTokenInfo) error {
	output, err := json.MarshalIndent(projectInfo, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON output: %w", err)
	}

	fmt.Println(string(output))
	return nil
}

func outputAgentTokensTable(agentInfo *tokens.AgentTokenInfo) error {
	// Header
	output.Printf("\n%s\n", output.Bold("📊 Token Analysis Results"))
	output.Printf("%s %s\n", output.PrintCyan("Agent:"), agentInfo.AgentName)
	output.Printf("%s %s\n", output.PrintCyan("File:"), agentInfo.AgentFile)
	output.Printf("%s %d tokens\n", output.PrintCyan("Total:"), agentInfo.TotalTokens)

	// Agent file tokens
	if len(agentInfo.Assets) > 0 {
		output.PrintBold("Agent Configuration:")
		for _, asset := range agentInfo.Assets {
			output.Printf("  %s %s (%d tokens)\n",
				output.PrintGreenBold("✓"), asset.Path, asset.Tokens)
		}
		output.Newline()
	}

	// Dependencies breakdown
	if len(agentInfo.Dependencies.Tasks) > 0 ||
		len(agentInfo.Dependencies.Templates) > 0 ||
		len(agentInfo.Dependencies.DataFiles) > 0 {

		output.PrintBold("Dependencies:")

		// Tasks
		if len(agentInfo.Dependencies.Tasks) > 0 {
			output.Printf("  %s Tasks:\n", output.PrintYellow("📋"))
			for _, task := range agentInfo.Dependencies.Tasks {
				output.Printf("    • %s (%d tokens)\n", task.Path, task.Tokens)
			}
		}

		// Templates
		if len(agentInfo.Dependencies.Templates) > 0 {
			output.Printf("  %s Templates:\n", output.PrintMagenta("📄"))
			for _, template := range agentInfo.Dependencies.Templates {
				output.Printf("    • %s (%d tokens)\n", template.Path, template.Tokens)
			}
		}

		// Data files
		if len(agentInfo.Dependencies.DataFiles) > 0 {
			output.Printf("  %s Data Files:\n", output.PrintCyan("📊"))
			for _, dataFile := range agentInfo.Dependencies.DataFiles {
				output.Printf("    • %s (%d tokens)\n", dataFile.Path, dataFile.Tokens)
			}
		}

		// Add disclaimer about token approximation
		output.Newline()
		output.Printf("%s  Token counts are approximate and best aligned with GPT-4 models.\n",
			output.PrintCyan("ℹ️"))

		return nil
	}

	output.Printf("  %s No dependencies found\n", output.PrintCyan("ℹ️"))

	// Add disclaimer about token approximation
	output.Newline()
	output.Printf("%s  Token counts are approximate and best aligned with GPT-4 models.\n",
		output.PrintCyan("ℹ️"))

	return nil
}

func outputProjectTokensTable(projectInfo *tokens.ProjectTokenInfo) error {
	// Header
	output.Printf("\n%s\n", output.Bold("📊 Project Token Analysis"))
	output.Printf("%s %d agents analyzed\n", output.PrintCyan("Agents:"), len(projectInfo.Agents))
	output.Printf("%s %d total tokens\n", output.PrintCyan("Total:"), projectInfo.TotalTokens)

	// Token breakdown
	output.PrintBold("Token Breakdown by Asset Type:")
	output.Printf("  %s Agents: %d tokens\n", output.PrintGreenBold("🤖"), projectInfo.Breakdown.Agents)
	output.Printf("  %s Tasks: %d tokens\n", output.PrintYellow("📋"), projectInfo.Breakdown.Tasks)
	output.Printf("  %s Templates: %d tokens\n", output.PrintMagenta("📄"), projectInfo.Breakdown.Templates)
	output.Printf("  %s Data Files: %d tokens\n", output.PrintBlue("📊"), projectInfo.Breakdown.DataFiles)
	output.Newline()

	output.PrintBold("Individual Agent Token Counts:")
	for _, agent := range projectInfo.Agents {
		output.Printf("  %s %s: %d tokens\n",
			output.PrintGreenBold("✓"), agent.AgentName, agent.TotalTokens)
	}

	// Add disclaimer about token approximation
	output.Newline()
	output.Printf("%s  Token counts are approximate and best aligned with GPT-4 models.\n",
		output.PrintCyan("ℹ️"))

	return nil
}

func handleTokenError(err error, jsonOutput bool) error {
	if jsonOutput {
		// For JSON output, return the raw error
		return err
	}

	// For interactive output, show user-friendly error with suggestions
	fmt.Fprintf(os.Stderr, "%s\n", tokens.FormatUserFriendlyError(err))

	suggestions := tokens.SuggestSolutions(err)
	if len(suggestions) > 0 {
		fmt.Fprintf(os.Stderr, "\n💡 Suggestions:\n")
		for _, suggestion := range suggestions {
			fmt.Fprintf(os.Stderr, "%s\n", suggestion)
		}
	}

	return err
}

// outputBundleTokensJSON outputs bundle token information in JSON format
func outputBundleTokensJSON(bundleInfo *tokens.BundleTokenInfo) error {
	output, err := json.MarshalIndent(bundleInfo, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON output: %w", err)
	}

	fmt.Println(string(output))
	return nil
}

// outputBundleTokensTable outputs bundle token information in table format
func outputBundleTokensTable(bundleInfo *tokens.BundleTokenInfo, bundleScope string) error {
	// Header
	output.Printf("\n%s\n", output.Bold("📊 Bundle Token Analysis"))
	output.Printf("%s %s\n", output.PrintCyan("Bundle Scope:"), bundleScope)
	output.Printf("%s %s\n", output.PrintCyan("Bundle File:"), bundleInfo.BundleFile)
	output.Printf("%s %d tokens\n", output.PrintCyan("Total Tokens:"), bundleInfo.TotalTokens)

	// Add disclaimer about token approximation
	output.Newline()
	output.Printf("%s  Token count is from the actual bundle file content.\n",
		output.PrintCyan("ℹ️"))
	output.Printf("%s  Use 'krci-ai bundle %s' to generate/update this bundle.\n",
		output.PrintCyan("ℹ️"), generateBundleCommandFromScope(bundleScope))
	output.Printf("%s  Token counts are approximate and best aligned with GPT-4 models.\n",
		output.PrintCyan("ℹ️"))

	return nil
}

// generateBundleCommandFromScope generates the appropriate bundle command for the scope
func generateBundleCommandFromScope(bundleScope string) string {
	if bundleScope == "all" {
		return "--all"
	}
	return fmt.Sprintf("--agent %s", bundleScope)
}
