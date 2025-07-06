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

	"github.com/epam/kuberocketai/internal/cli"
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
  krci-ai install --ide=all          # Install with all IDE integrations`,
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

		// Validate IDE flag value if provided
		if ideFlag != "" {
			validIDEs := []string{"cursor", "claude", "vscode", "windsurf", "all"}
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

		// Placeholder for actual installation logic
		output.PrintWarning("This feature is coming soon!")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Add IDE integration flag
	installCmd.Flags().StringP("ide", "i", "", "IDE integration (cursor, claude, vscode, windsurf, all)")
}
