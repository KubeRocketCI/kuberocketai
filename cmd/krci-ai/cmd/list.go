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
	"strings"

	"github.com/spf13/cobra"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/cli"
	"github.com/KubeRocketCI/kuberocketai/internal/discovery"
	"github.com/KubeRocketCI/kuberocketai/internal/utils"
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
		errorHandler := cli.NewErrorHandler()
		outputHandler := cli.NewOutputHandler()

		// Check verbose flag
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			errorHandler.HandleError(err, "Failed to read verbose flag")
			return
		}

		projectRoot, err := discovery.GetProjectRoot()
		if err != nil {
			errorHandler.HandleError(err, "Failed to get project root")
			return
		}

		// Create discovery service
		discovery := assets.NewDiscovery(assets.GetKrciPath(projectRoot))

		// Discover agents
		agents, err := discovery.GetAgents(context.Background())
		if err != nil {
			errorHandler.HandleError(err, "Failed to discover agents")
			return
		}

		// Display results
		if len(agents) == 0 {
			outputHandler.PrintWarning("No agents found")
			outputHandler.PrintInfo("The framework is installed but no agent files were found")
			return
		}

		// Print header
		outputHandler.PrintSuccess(fmt.Sprintf("Found %d agent(s):", len(agents)))
		outputHandler.Newline()

		if verbose {
			// Verbose output with dependency table

			// Show dependency table
			outputHandler.PrintInfo("Agent Dependencies:")
			outputHandler.Newline()
			table := formatAgentDependencyTable(agents)
			fmt.Print(table)
		} else {
			// Simple table format
			summary := formatAgentSummary(agents)
			fmt.Print(summary)
		}

		outputHandler.Newline()
		outputHandler.PrintInfo("Use 'krci-ai list agents -v' for dependency table showing tasks, templates, and data")
	},
}

// formatAgentSummary formats agents for simple table display using lipgloss
func formatAgentSummary(agents []assets.Agent) string {
	rows := make([][]string, 0, len(agents))

	for _, agent := range agents {
		description := cli.TruncateDescription(agent.Description)
		rows = append(rows, []string{agent.Name, agent.Role, description})
	}

	t := cli.CreateStyledTable().
		Headers("NAME", "ROLE", "DESCRIPTION").
		Rows(rows...)

	return t.String()
}

// formatAgentDependencyTable creates a styled table showing agent dependencies
func formatAgentDependencyTable(agents []assets.Agent) string {
	rows := make([][]string, 0, len(agents))

	for _, agent := range agents {
		taskNames := make([]string, 0, len(agent.Tasks))
		templateNames := make([]string, 0)
		dataFileNames := make([]string, 0)

		for _, task := range agent.Tasks {
			taskNames = append(taskNames, task.Name)

			for _, template := range task.Dependencies.Templates {
				templateNames = append(templateNames, template.Name)
			}

			for _, dataFile := range task.Dependencies.DataFiles {
				dataFileNames = append(dataFileNames, dataFile.Name)
			}
		}

		tasksStr := strings.Join(taskNames, "\n")
		if tasksStr == "" {
			tasksStr = cli.NoneValue
		}

		templatesStr := strings.Join(utils.DeduplicateStrings(templateNames), "\n")
		if templatesStr == "" {
			templatesStr = cli.NoneValue
		}

		dataFilesStr := strings.Join(utils.DeduplicateStrings(dataFileNames), "\n")
		if dataFilesStr == "" {
			dataFilesStr = cli.NoneValue
		}

		rows = append(rows, []string{agent.Name, tasksStr, templatesStr, dataFilesStr})
	}

	t := cli.CreateStyledTable().
		Headers("AGENT", "TASKS", "TEMPLATES", "DATA FILES").
		Rows(rows...)

	return t.String()
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listAgentsCmd)

	// Add verbose flag to list agents command
	listAgentsCmd.Flags().BoolP("verbose", "v", false, "Show detailed agent information")
}
