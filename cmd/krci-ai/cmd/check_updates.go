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

	"github.com/KubeRocketCI/kuberocketai/internal/changelog"
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

This command performs two main functions:
1. Online version checking: Queries GitHub API to detect available updates
2. Offline changelog display: Shows embedded changelog from current CLI version

The command requires network connectivity for version checking but can display
the embedded changelog offline. If network connectivity fails, it provides
a direct link to the GitHub releases page for manual checking.

Examples:
  krci-ai check-updates              # Check for updates online
  krci-ai check-updates --changelog  # Show embedded changelog offline
  krci-ai check-updates -c           # Short form for changelog`,
	RunE: runCheckUpdates,
}

func init() {
	checkUpdatesCmd.Flags().BoolP("changelog", "c", false, "Show embedded changelog")
}

func runCheckUpdates(cmd *cobra.Command, args []string) error {
	showChangelog, _ := cmd.Flags().GetBool("changelog")

	if showChangelog {
		return displayOfflineChangelog()
	}

	return checkOnlineUpdates()
}

func displayOfflineChangelog() error {
	assets := GetEmbeddedAssets()
	reader := changelog.NewReader(assets)

	color.Cyan("📋 Changelog (Embedded)")
	color.Cyan("========================")
	fmt.Println()

	content, err := reader.ReadChangelog()
	if err != nil {
		color.Yellow("⚠️  Warning: %v", err)
		fmt.Println()
	}

	// Format and display changelog
	formatted, err := changelog.FormatChangelog(content)
	if err != nil {
		// Fallback to raw content if formatting fails
		color.Yellow("⚠️  Could not format changelog, showing raw content:")
		fmt.Println()
		fmt.Println(content)
		return nil
	}

	fmt.Println(formatted)

	// Show status info
	fmt.Println()
	if reader.HasChangelog() {
		color.Green("✅ Embedded changelog available")
	} else {
		color.Yellow("⚠️  Using fallback changelog - embedded version not available")
	}

	color.Cyan("💡 Tip: Run 'krci-ai check-updates' without --changelog to check for newer versions online")

	return nil
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
	updateInfo, err := checker.CheckForUpdatesWithRetry(currentVersion, 2)
	if err != nil {
		return fmt.Errorf("failed to check for updates: %w", err)
	}

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

	fmt.Println()
	color.Cyan("💡 Tip: Run 'krci-ai check-updates --changelog' to see embedded changelog")

	return nil
}
