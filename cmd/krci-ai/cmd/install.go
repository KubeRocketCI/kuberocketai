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
	"slices"

	"github.com/spf13/cobra"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/cli"
)

const (
	ideAll      = "all"
	ideCursor   = "cursor"
	ideClaude   = "claude"
	ideVSCode   = "vscode"
	ideWindsurf = "windsurf"
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
  krci-ai install                              # Install core framework components
  krci-ai install --ide=cursor                 # Install with Cursor IDE integration
  krci-ai install --agent developer            # Install only developer agent and dependencies
  krci-ai install --agents pm,architect,dev    # Install specific agents (comma-separated)
  krci-ai install --agent developer --ide cursor # Selective install with IDE integration
  krci-ai install --all              # Install core + all IDE integrations (shortcut)
  krci-ai install --all --force      # Force install everything
  krci-ai install --force            # Force install core components only`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create output and error handlers
		output := cli.NewOutputHandler()
		errorHandler := cli.NewErrorHandler()

		// Validate IDE flag upfront (common to all installation paths)
		ideFlag, err := cmd.Flags().GetString("ide")
		if err != nil {
			errorHandler.HandleError(err, "Failed to read IDE flag")
			return
		}

		// Check --all flag and override IDE if necessary
		allFlag, err := cmd.Flags().GetBool("all")
		if err != nil {
			errorHandler.HandleError(err, "Failed to read all flag")
			return
		}
		if allFlag {
			ideFlag = ideAll
		}

		// Validate IDE flag after potential --all override
		if ideFlag != "" {
			if validateErr := validateIDEFlag(ideFlag, errorHandler); validateErr != nil {
				return
			}
		}

		// Check for selective installation first
		agentFlag, err := cmd.Flags().GetString("agent")
		if err != nil {
			errorHandler.HandleError(err, "Failed to read agent flag")
			return
		}

		// Check aliases flag if main flag is empty
		if agentFlag == "" {
			if agentsFlag, err := cmd.Flags().GetString("agents"); err == nil && agentsFlag != "" {
				agentFlag = agentsFlag
			}
		}

		if agentFlag != "" {
			runSelectiveInstallation(cmd, agentFlag, ideFlag, output, errorHandler)
			return
		}

		// Standard full installation
		runFullInstallation(cmd, ideFlag, output, errorHandler)
	},
}

// runSelectiveInstallation handles installation of specific agents only
func runSelectiveInstallation(_ *cobra.Command, agentFlag string, ideFlag string, output *cli.OutputHandler, errorHandler *cli.ErrorHandler) {
	// Parse agent list using existing bundle logic
	agentNames := ParseAgentList(agentFlag)
	if len(agentNames) == 0 {
		errorHandler.PrintError("No valid agent names provided")
		return
	}

	// Start selective installation process
	output.PrintProgress(fmt.Sprintf("Installing selected agents: %v", agentNames))

	// Create installer and run selective installation
	installer := assets.NewInstaller(".", GetEmbeddedAssets())
	if err := installer.InstallSelective(agentNames); err != nil {
		errorHandler.HandleError(err, "Failed to install selected agents")
		return
	}

	// Handle IDE integration if requested
	if ideFlag != "" {
		output.PrintInfo(fmt.Sprintf("Setting up %s IDE integration for selected agents...", ideFlag))
		handleIDEIntegration(installer, ideFlag, output, errorHandler)
	}

	output.PrintSuccess(fmt.Sprintf("Selected agents installed successfully: %v", agentNames))
}

// runFullInstallation handles standard full framework installation
func runFullInstallation(cmd *cobra.Command, ideFlag string, output *cli.OutputHandler, errorHandler *cli.ErrorHandler) {
	output.PrintProgress("Installing KubeRocketAI framework components...")

	// Get force flag (other flags already processed upfront)
	forceFlag, err := cmd.Flags().GetBool("force")
	if err != nil {
		errorHandler.HandleError(err, "Failed to read force flag")
		return
	}

	// Show IDE integration info if enabled
	if ideFlag != "" {
		output.PrintInfo(fmt.Sprintf("IDE integration: %s", ideFlag))
	}

	// Create installer
	installer := assets.NewInstaller(".", GetEmbeddedAssets())

	// Check installation status and force flag
	if !handleInstallationCheck(installer, forceFlag, output) {
		return
	}

	// Perform core installation
	if err := performCoreInstallation(installer, output, errorHandler); err != nil {
		return
	}

	// Handle IDE integration using existing function
	handleIDEIntegration(installer, ideFlag, output, errorHandler)

	// Show success and next steps
	showInstallationSuccess(installer, ideFlag, output)
}

// handleInstallationCheck checks if installation should proceed
func handleInstallationCheck(installer *assets.Installer, forceFlag bool, output *cli.OutputHandler) bool {
	if installer.IsInstalled() && !forceFlag {
		output.PrintWarning("Framework already installed in current directory")
		output.PrintInfo("If you want to proceed with installation, use the --force flag:")
		output.PrintInfo("  krci-ai install --force")
		return false
	}

	if installer.IsInstalled() && forceFlag {
		output.PrintWarning("Framework already installed - proceeding with forced installation")
	}

	return true
}

// performCoreInstallation handles the core framework installation
func performCoreInstallation(installer *assets.Installer, output *cli.OutputHandler, errorHandler *cli.ErrorHandler) error {
	// Install framework components
	output.PrintInfo("Creating .krci-ai directory structure...")
	if err := installer.Install(); err != nil {
		errorHandler.HandleError(err, "Failed to install framework components")
		return err
	}

	// Validate installation
	if err := installer.ValidateInstallation(); err != nil {
		errorHandler.HandleError(err, "Installation validation failed")
		return err
	}

	return nil
}

// showInstallationSuccess displays success message and next steps
func showInstallationSuccess(installer *assets.Installer, ideFlag string, output *cli.OutputHandler) {
	// Success
	output.PrintSuccess("Framework installation completed successfully!")
	output.PrintInfo("Framework components installed to: " + installer.GetInstallationPath())

	// Show next steps
	output.PrintInfo("\nNext steps:")
	output.PrintInfo("  • Run 'krci-ai list agents' to see available agents")
	output.PrintInfo("  • Run 'krci-ai validate' to verify installation")

	if ideFlag == ideCursor || ideFlag == ideAll {
		output.PrintInfo("  • Cursor IDE rules installed to: " + installer.GetCursorRulesPath() + "/")
	}
	if ideFlag == ideClaude || ideFlag == ideAll {
		output.PrintInfo("  • Claude Code commands installed to: " + installer.GetClaudeCommandsPath() + "/")
	}
	if ideFlag == ideVSCode || ideFlag == ideAll {
		output.PrintInfo("  • VS Code chat modes installed to: " + installer.GetVSCodeChatmodesPath() + "/")
	}
	if ideFlag == ideWindsurf || ideFlag == ideAll {
		output.PrintInfo("  • Windsurf IDE rules installed to: " + installer.GetWindsurfRulesPath() + "/")
	}
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Add IDE integration flag
	installCmd.Flags().StringP("ide", "i", "", "IDE integration (cursor, claude, vscode, windsurf, all)")

	// Add force flag
	installCmd.Flags().BoolP("force", "f", false, "Force installation even if framework is already installed")

	// Add all flag
	installCmd.Flags().BoolP("all", "a", false, "Install core framework with all IDE integrations (equivalent to --ide=all)")

	// Add selective installation flags (following bundle command patterns)
	installCmd.Flags().String("agent", "", "Install specific agents (comma or space separated: 'pm,architect' or 'pm architect')")
	installCmd.Flags().String("agents", "", "Alias for --agent flag")
}

// validateIDEFlag validates the IDE flag value and returns error if invalid
func validateIDEFlag(ideFlag string, errorHandler *cli.ErrorHandler) error {
	validIDEs := []string{ideCursor, ideClaude, ideVSCode, ideWindsurf, ideAll}
	if slices.Contains(validIDEs, ideFlag) {
		return nil
	}

	errorHandler.PrintError(fmt.Sprintf("Invalid IDE '%s'. Valid options: cursor, claude, vscode, windsurf, all", ideFlag))
	return fmt.Errorf("invalid IDE flag")
}

// handleIDEIntegration handles IDE integration setup for any installation type
func handleIDEIntegration(installer *assets.Installer, ideFlag string, output *cli.OutputHandler, errorHandler *cli.ErrorHandler) {
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

	if ideFlag == ideVSCode || ideFlag == ideAll {
		output.PrintInfo("Setting up VS Code integration...")
		if err := installer.InstallVSCodeIntegration(); err != nil {
			errorHandler.HandleError(err, "Failed to install VS Code integration")
			return
		}
		output.PrintSuccess("VS Code integration installed successfully!")
	}

	if ideFlag == ideWindsurf || ideFlag == ideAll {
		output.PrintInfo("Setting up Windsurf IDE integration...")
		if err := installer.InstallWindsurfIntegration(); err != nil {
			errorHandler.HandleError(err, "Failed to install Windsurf IDE integration")
			return
		}
		output.PrintSuccess("Windsurf IDE integration installed successfully!")
	}
}
