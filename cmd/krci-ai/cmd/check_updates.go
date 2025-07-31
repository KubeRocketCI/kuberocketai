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

	"github.com/KubeRocketCI/kuberocketai/internal/update"
	"github.com/KubeRocketCI/kuberocketai/internal/version"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// checkUpdatesCmd represents the check-updates command
var checkUpdatesCmd = &cobra.Command{
	Use:   "check-updates",
	Short: "Check for CLI updates",
	Long: `Check for available updates to the krci-ai CLI.

This command queries the GitHub API to detect available updates and displays
release information. If network connectivity fails, it provides a direct
link to the GitHub releases page for manual checking.

Examples:
  krci-ai check-updates              # Check for updates online`,
	RunE: runCheckUpdates,
}

func init() {
	// No flags needed - simple version checking only
}

func runCheckUpdates(cmd *cobra.Command, args []string) error {
	return checkOnlineUpdates()
}

func checkOnlineUpdates() error {
	currentVersion := version.GetCurrentVersion()

	color.Cyan("🔍 Checking for Updates")
	color.Cyan("=======================")
	fmt.Println()

	color.White("Current version: %s", currentVersion)

	// Create update checker
	checker := update.NewChecker()

	// Check for updates with retry
	updateInfo := checker.CheckForUpdatesWithRetry(currentVersion, 2)

	// Display results
	if updateInfo.Error != "" {
		color.Red("❌ Update check failed: %s", updateInfo.Error)
		fmt.Println()
		color.Yellow("📖 Manual check: %s", updateInfo.FallbackURL)
		return nil
	}

	if updateInfo.IsUpdateAvailable {
		color.Green("🎉 Update available!")
		color.White("Latest version: %s", updateInfo.LatestVersion)

		if updateInfo.LatestRelease != nil {
			fmt.Println()
			color.Cyan("Release information:")
			color.White("• Published: %s", updateInfo.LatestRelease.PublishedAt.Format("2006-01-02"))
			color.White("• Release page: %s", updateInfo.LatestRelease.URL)

			if updateInfo.LatestRelease.Name != "" {
				color.White("• Release name: %s", updateInfo.LatestRelease.Name)
			}
		}

		fmt.Println()
		color.Cyan("💡 Visit the release page above to download the latest version")
	} else {
		color.Green("✅ You are using the latest version!")
		color.White("Latest version: %s", updateInfo.LatestVersion)
	}

	return nil
}
