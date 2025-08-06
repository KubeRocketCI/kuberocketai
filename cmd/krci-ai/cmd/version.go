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
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/KubeRocketCI/kuberocketai/internal/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long: `Display detailed version information for the krci-ai CLI.

This command shows:
- CLI version and build information
- Framework version compatibility
- Go version and platform details
- Helpful hints for checking updates

Use the --output flag to format the output as JSON for programmatic use.`,
	RunE: runVersion,
}

func init() {
	versionCmd.Flags().StringP("output", "o", "", "Output format (json)")
}

func runVersion(cmd *cobra.Command, args []string) error {
	outputFormat, _ := cmd.Flags().GetString("output")

	versionInfo := version.GetVersionInfo()

	switch outputFormat {
	case "json":
		return outputJSON(versionInfo)
	default:
		return outputDefault(versionInfo)
	}
}

func outputJSON(versionInfo version.VersionInfo) error {
	jsonBytes, err := json.MarshalIndent(versionInfo, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal version info to JSON: %w", err)
	}

	fmt.Println(string(jsonBytes))
	return nil
}

func outputDefault(versionInfo version.VersionInfo) error {
	// Display formatted version information
	fmt.Println(versionInfo.String())

	// Add helpful hint about update checking
	fmt.Println()
	color.Cyan("💡 Tip: Run 'krci-ai check-updates' to check for newer versions")

	return nil
}
