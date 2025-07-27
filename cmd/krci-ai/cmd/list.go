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

	"github.com/spf13/cobra"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/cli"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List KubeRocketAI framework components",
	Long: `List installed KubeRocketAI framework components in the current directory.

This command helps you discover what agents, tasks, templates, and data
components are available in your local framework installation.

Available subcommands:
  agents     List all installed agents with their roles and descriptions
  tasks      List all installed tasks (coming soon)
  templates  List all installed templates (coming soon)
  data       List all installed data components (coming soon)

Examples:
  krci-ai list agents          # List all available agents
  krci-ai list agents -v       # List agents with dependency table showing tasks, templates, and data`,
}

// listAgentsCmd represents the list agents command
var listAgentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "List all installed agents",
	Long: `List all installed agents in the current directory.

This command scans the .krci-ai/agents/ directory and displays information
about each agent including their name, role, and description.

The agents are read from YAML files that were installed by the 'krci-ai install' command.

Examples:
  krci-ai list agents          # List all agents
  krci-ai list agents -v       # List agents with dependency table showing tasks, templates, and data`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create output and error handlers
		output := cli.NewOutputHandler()
		errorHandler := cli.NewErrorHandler()

		// Check verbose flag
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			errorHandler.HandleError(err, "Failed to read verbose flag")
			return
		}

		// Create discovery service
		discovery := assets.NewDiscovery(".", GetEmbeddedAssets())

		// Discover agents
		agents, err := discovery.DiscoverAgents()
		if err != nil {
			// Check if it's a "not installed" error
			if err.Error() == "framework not installed in current directory - run 'krci-ai install'" {
				output.PrintWarning("Framework not installed in current directory")
				output.PrintInfo("Run 'krci-ai install' to install the framework and then try again")
				return
			}
			errorHandler.HandleError(err, "Failed to discover agents")
			return
		}

		// Display results
		if len(agents) == 0 {
			output.PrintWarning("No agents found")
			output.PrintInfo("The framework is installed but no agent files were found")
			return
		}

		// Print header
		output.PrintSuccess(fmt.Sprintf("Found %d agent(s):", len(agents)))
		fmt.Println()

		if verbose {
			// Verbose output with dependency table
			agentDeps, err := discovery.DiscoverAgentsWithDependencies()
			if err != nil {
				output.PrintWarning("Could not load dependency information, showing basic details only")
				output.PrintError(fmt.Sprintf("Dependency error: %v", err))
				fmt.Println()

				// Fallback to basic information
				for _, agent := range agents {
					fmt.Printf("Agent: %s\n", agent.Name)
					fmt.Printf("  Role: %s\n", agent.Role)
					fmt.Printf("  Description: %s\n", agent.Description)
					if agent.Goal != "" {
						fmt.Printf("  Goal: %s\n", agent.Goal)
					}
					fmt.Printf("  File: %s\n", agent.FilePath)
					fmt.Println()
				}
			} else {
				// Show dependency table
				output.PrintInfo("Agent Dependencies:")
				fmt.Println()
				table := discovery.FormatAgentDependencyTable(agentDeps)
				fmt.Print(table)
			}
		} else {
			// Simple table format
			fmt.Printf("%-15s | %-25s | %s\n", "Name", "Role", "Description")
			fmt.Printf("%-15s | %-25s | %s\n", "---------------", "-------------------------", "-----------------------------------")
			for _, agent := range agents {
				summary := discovery.FormatAgentSummary(agent)
				fmt.Println(summary)
			}
		}

		fmt.Println()
		output.PrintInfo("Use 'krci-ai list agents -v' for dependency table showing tasks, templates, and data")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listAgentsCmd)

	// Add verbose flag to list agents command
	listAgentsCmd.Flags().BoolP("verbose", "v", false, "Show detailed agent information")
}
