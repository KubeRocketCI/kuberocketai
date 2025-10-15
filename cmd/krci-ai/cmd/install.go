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
	"github.com/KubeRocketCI/kuberocketai/internal/discovery"
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

Supported IDE Integrations:
- cursor     → .cursor/rules/krci-ai/*.mdc
- claude     → .claude/commands/krci-ai/*.md
- vscode     → .github/chatmodes/*.chatmode.md
- windsurf   → .windsurf/rules/*.md
- all        → Install all IDE integrations above

Examples:
  # Basic installation
  krci-ai install                              # Install core structure + all agents (./.krci-ai)
  krci-ai install --force                      # Force reinstall core structure + all agents

  # IDE integrations (see supported list above)
  krci-ai install --ide cursor                 # Install core + specific IDE integration
  krci-ai install --all                        # Install core + all IDE integrations
  krci-ai install -a                           # Same as --all (shorthand)

  # Selective installation (core structure + specific agents only)
  krci-ai install --agent dev                  # Install core structure + dev agent + dependencies
  krci-ai install --agents pm,architect        # Install core structure + multiple agents (comma-separated)
  krci-ai install --agent "pm po qa"           # Install core structure + multiple agents (space-separated)

  # Combined selective + IDE
  krci-ai install --agent dev -i cursor        # Install core + dev agent + IDE integration
  krci-ai install --agents pm,po --ide vscode  # Install core + multiple agents + IDE

  # Force operations
  krci-ai install --ide cursor --force         # Add IDE integration to existing install
  krci-ai install --all --force                # Force reinstall everything

  # Sync IDE files from installed agents (instead of embedded assets)
  krci-ai install --sync-ide                   # Sync all existing IDE integrations from installed agents
  krci-ai install --agent dev --sync-ide       # Install dev agent + sync IDE files`,
	Run: func(cmd *cobra.Command, args []string) {
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
			agentsFlag, flagErr := cmd.Flags().GetString("agents")
			if flagErr == nil && agentsFlag != "" {
				agentFlag = agentsFlag
			}
		}

		// Check sync-ide flag
		syncIDEFlag, err := cmd.Flags().GetBool("sync-ide")
		if err != nil {
			errorHandler.HandleError(err, "Failed to read sync-ide flag")
			return
		}

		if agentFlag != "" {
			runSelectiveInstallation(cmd, agentFlag, ideFlag, syncIDEFlag, output, errorHandler)
			return
		}

		// Standard full installation
		runFullInstallation(cmd, ideFlag, syncIDEFlag, output, errorHandler)
	},
}

// runSelectiveInstallation handles installation of specific agents only
func runSelectiveInstallation(_ *cobra.Command, agentFlag string, ideFlag string, syncIDEFlag bool, output *cli.OutputHandler, errorHandler *cli.ErrorHandler) {
	// Parse agent list using existing bundle logic
	agentNames := ParseAgentList(agentFlag)
	if len(agentNames) == 0 {
		errorHandler.PrintError("No valid agent names provided")
		return
	}

	// Start selective installation process
	output.PrintProgress(fmt.Sprintf("Installing selected agents: %v", agentNames))

	projectRoot, err := discovery.GetProjectRoot()
	if err != nil {
		errorHandler.HandleError(err, "Failed to get project root")
		return
	}

	// Create installer and run selective installation
	installer := assets.NewInstaller(
		projectRoot,
		GetEmbeddedAssets(),
		assets.NewEmbeddedDiscovery(GetEmbeddedAssets(), assets.EmbeddedPrefix),
	)
	if err := installer.InstallSelective(agentNames); err != nil {
		errorHandler.HandleError(err, "Failed to install selected agents")
		return
	}

	// Handle IDE integration if requested
	if ideFlag != "" {
		output.PrintInfo(fmt.Sprintf("Setting up %s IDE integration for selected agents...", ideFlag))
		handleIDEIntegration(installer, ideFlag, output, errorHandler)
	}

	// Handle sync-ide flag at the end
	if syncIDEFlag {
		output.PrintInfo("Syncing IDE integration files from installed agents...")
		handleIDESync(installer, output, errorHandler)
	}

	output.PrintSuccess(fmt.Sprintf("Selected agents installed successfully: %v", agentNames))
}

// runFullInstallation handles standard full framework installation
func runFullInstallation(cmd *cobra.Command, ideFlag string, syncIDEFlag bool, output *cli.OutputHandler, errorHandler *cli.ErrorHandler) {
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

	projectRoot, err := discovery.GetProjectRoot()
	if err != nil {
		errorHandler.HandleError(err, "Failed to get project root")
		return
	}

	// Create installer
	installer := assets.NewInstaller(
		projectRoot,
		GetEmbeddedAssets(),
		assets.NewEmbeddedDiscovery(GetEmbeddedAssets(), assets.EmbeddedPrefix),
	)

	// Check installation status
	isAlreadyInstalled := installer.IsInstalled()

	// Special case: If already installed and only syncing IDE files (no reinstall needed)
	if isAlreadyInstalled && syncIDEFlag && !forceFlag {
		output.PrintInfo("Framework already installed. Syncing IDE integration files from installed agents...")
		handleIDESync(installer, output, errorHandler)
		output.PrintSuccess("IDE integration files synced successfully!")
		return
	}

	// Check if installation should proceed
	if isAlreadyInstalled && !forceFlag {
		output.PrintWarning("Framework already installed in current directory")
		output.PrintInfo("If you want to proceed with installation, use the --force flag:")
		output.PrintInfo("  krci-ai install --force")
		return
	}

	if isAlreadyInstalled && forceFlag {
		output.PrintWarning("Framework already installed - proceeding with forced installation")
	}

	// Perform core installation
	if err := performCoreInstallation(installer, output, errorHandler); err != nil {
		return
	}

	// Handle IDE integration
	// If sync-ide is set, use installed agents; otherwise use embedded assets
	if syncIDEFlag {
		output.PrintInfo("Setting up IDE integration from installed agents...")
		handleIDESync(installer, output, errorHandler)
	} else if ideFlag != "" {
		handleIDEIntegration(installer, ideFlag, output, errorHandler)
	}

	// Show success and next steps
	showInstallationSuccess(installer, ideFlag, output)
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
	output.PrintInfo("Framework components installed to: " + installer.GetFrameworkPath())

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
	installCmd.Flags().StringP("ide", "i", "", "IDE integration: cursor, claude, vscode, windsurf, or 'all' for everything")

	// Add force flag
	installCmd.Flags().BoolP("force", "f", false, "Force installation even if framework is already installed")

	// Add all flag
	installCmd.Flags().BoolP("all", "a", false, "Install core framework with all IDE integrations (equivalent to --ide=all)")

	// Add selective installation flags (following bundle command patterns)
	installCmd.Flags().String("agent", "", "Install specific agents (comma or space separated: 'pm,architect' or 'pm architect')")
	installCmd.Flags().String("agents", "", "Alias for --agent flag")

	// Add sync-ide flag
	installCmd.Flags().Bool("sync-ide", false, "Sync IDE integration files from installed agents")
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

// handleIDESync syncs IDE integration files from installed agents
func handleIDESync(installer *assets.Installer, output *cli.OutputHandler, errorHandler *cli.ErrorHandler) {
	// Check which IDE directories exist and sync them
	if installer.HasCursorIntegration() {
		output.PrintInfo("Syncing Cursor IDE integration...")
		if err := installer.SyncCursorIntegration(); err != nil {
			errorHandler.HandleError(err, "Failed to sync Cursor IDE integration")
			return
		}
		output.PrintSuccess("Cursor IDE integration synced successfully!")
	}

	if installer.HasClaudeIntegration() {
		output.PrintInfo("Syncing Claude Code integration...")
		if err := installer.SyncClaudeIntegration(); err != nil {
			errorHandler.HandleError(err, "Failed to sync Claude Code integration")
			return
		}
		output.PrintSuccess("Claude Code integration synced successfully!")
	}

	if installer.HasVSCodeIntegration() {
		output.PrintInfo("Syncing VS Code integration...")
		if err := installer.SyncVSCodeIntegration(); err != nil {
			errorHandler.HandleError(err, "Failed to sync VS Code integration")
			return
		}
		output.PrintSuccess("VS Code integration synced successfully!")
	}

	if installer.HasWindsurfIntegration() {
		output.PrintInfo("Syncing Windsurf IDE integration...")
		if err := installer.SyncWindsurfIntegration(); err != nil {
			errorHandler.HandleError(err, "Failed to sync Windsurf IDE integration")
			return
		}
		output.PrintSuccess("Windsurf IDE integration synced successfully!")
	}
}
