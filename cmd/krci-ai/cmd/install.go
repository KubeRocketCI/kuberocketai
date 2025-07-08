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

	"github.com/epam/kuberocketai/internal/assets"
	"github.com/epam/kuberocketai/internal/cli"
)

const (
	ideAll    = "all"
	ideCursor = "cursor"
	ideClaude = "claude"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install KubeRocketAI framework components",
	Long: `Install KubeRocketAI framework components to your local project.

This command will set up the AI-as-Code framework by installing:
- Agent definitions for different development roles
- Task templates for common workflows
- Output templates for consistent formatting
- Reference data for coding standards and best practices

Examples:
  krci-ai install                    # Install core framework components
  krci-ai install --ide=cursor       # Install with Cursor IDE integration
  krci-ai install --ide=claude       # Install with Claude Code integration
  krci-ai install --ide=all          # Install with all IDE integrations
  krci-ai install --all              # Install core + all IDE integrations (shortcut)
  krci-ai install --all --force      # Force install everything
  krci-ai install --force            # Force install core components only`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create output and error handlers
		output := cli.NewOutputHandler()
		errorHandler := cli.NewErrorHandler()

		// Start installation process
		output.PrintProgress("Installing KubeRocketAI framework components...")

		// Get and validate IDE flag
		ideFlag, err := cmd.Flags().GetString("ide")
		if err != nil {
			errorHandler.HandleError(err, "Failed to read IDE flag")
		}

		// Get force flag
		forceFlag, err := cmd.Flags().GetBool("force")
		if err != nil {
			errorHandler.HandleError(err, "Failed to read force flag")
		}

		// Get all flag
		allFlag, err := cmd.Flags().GetBool("all")
		if err != nil {
			errorHandler.HandleError(err, "Failed to read all flag")
		}

		// If --all flag is set, override ide flag to "all"
		if allFlag {
			ideFlag = ideAll
			output.PrintInfo("--all flag detected, installing core framework with all IDE integrations")
		}

		// Validate IDE flag value if provided
		if ideFlag != "" {
			validIDEs := []string{ideCursor, ideClaude, "vscode", "windsurf", ideAll}
			isValid := false
			for _, valid := range validIDEs {
				if ideFlag == valid {
					isValid = true
					break
				}
			}

			if !isValid {
				errorHandler.PrintError(fmt.Sprintf("Invalid IDE '%s'. Valid options: cursor, claude, vscode, windsurf, all", ideFlag))
				return
			}

			output.PrintInfo(fmt.Sprintf("IDE integration: %s", ideFlag))
		}

		// Create installer
		installer := assets.NewInstaller(".", GetEmbeddedAssets())

		// Check if already installed
		if installer.IsInstalled() && !forceFlag {
			output.PrintWarning("Framework already installed in current directory")
			output.PrintInfo("If you want to proceed with installation, use the --force flag:")
			output.PrintInfo("  krci-ai install --force")
			return
		}

		// Show force installation message if forcing
		if installer.IsInstalled() && forceFlag {
			output.PrintWarning("Framework already installed - proceeding with forced installation")
		}

		// Install framework components
		output.PrintInfo("Creating .krci-ai directory structure...")
		if err := installer.Install(); err != nil {
			errorHandler.HandleError(err, "Failed to install framework components")
			return
		}

		// Validate installation
		if err := installer.ValidateInstallation(); err != nil {
			errorHandler.HandleError(err, "Installation validation failed")
			return
		}

		// Handle IDE integration
		if ideFlag == ideCursor || ideFlag == ideAll {
			output.PrintInfo("Setting up Cursor IDE integration...")
			if err := installer.InstallCursorIntegration(); err != nil {
				errorHandler.HandleError(err, "Failed to install Cursor IDE integration")
				return
			}
			output.PrintSuccess("Cursor IDE integration installed successfully!")
		}

		if ideFlag == ideClaude || ideFlag == ideAll {
			output.PrintInfo("Setting up Claude Code integration...")
			if err := installer.InstallClaudeIntegration(); err != nil {
				errorHandler.HandleError(err, "Failed to install Claude Code integration")
				return
			}
			output.PrintSuccess("Claude Code integration installed successfully!")
		}

		// Success
		output.PrintSuccess("Framework installation completed successfully!")
		output.PrintInfo("Framework components installed to: " + installer.GetInstallationPath())

		// Show next steps
		output.PrintInfo("\nNext steps:")
		output.PrintInfo("  • Run 'krci-ai list agents' to see available agents")
		output.PrintInfo("  • Run 'krci-ai validate' to verify installation")

		if ideFlag == ideCursor || ideFlag == ideAll {
			output.PrintInfo("  • Cursor IDE rules installed to: .cursor/rules/")
		}
		if ideFlag == ideClaude || ideFlag == ideAll {
			output.PrintInfo("  • Claude Code commands installed to: .claude/commands/")
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Add IDE integration flag
	installCmd.Flags().StringP("ide", "i", "", "IDE integration (cursor, claude, vscode, windsurf, all)")

	// Add force flag
	installCmd.Flags().BoolP("force", "f", false, "Force installation even if framework is already installed")

	// Add all flag
	installCmd.Flags().BoolP("all", "a", false, "Install core framework with all IDE integrations (equivalent to --ide=all)")
}
