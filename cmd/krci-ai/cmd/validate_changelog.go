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
	"os"

	"github.com/KubeRocketCI/kuberocketai/internal/changelog"
	"github.com/KubeRocketCI/kuberocketai/internal/cli"
	"github.com/spf13/cobra"
)

// validateChangelogCmd represents the validate-changelog command
var validateChangelogCmd = &cobra.Command{
	Use:   "validate-changelog [changelog-file]",
	Short: "Validate changelog format and structure",
	Long: `Validate changelog format and structure.

This command validates a changelog file for proper format, structure, and content.
It checks for:
- Proper markdown formatting
- Valid version headers
- Consistent section structure
- Date format validation
- Empty section detection

Examples:
  krci-ai validate-changelog CHANGELOG.md
  krci-ai validate-changelog CHANGELOG.md --verbose
  krci-ai validate-changelog CHANGELOG.md --quiet`,
	Args: cobra.ExactArgs(1),
	RunE: runValidateChangelog,
}

var (
	validateVerbose bool
	validateQuiet   bool
)

func init() {
	rootCmd.AddCommand(validateChangelogCmd)

	validateChangelogCmd.Flags().BoolVarP(&validateVerbose, "verbose", "v", false, "Show detailed validation results")
	validateChangelogCmd.Flags().BoolVarP(&validateQuiet, "quiet", "q", false, "Show only errors (suppress warnings)")
}

func runValidateChangelog(cmd *cobra.Command, args []string) error {
	changelogFile := args[0]
	output := cli.NewOutputHandler()

	// Read changelog file
	content, err := os.ReadFile(changelogFile)
	if err != nil {
		return fmt.Errorf("failed to read changelog file: %w", err)
	}

	// Perform detailed validation
	result := changelog.ValidateChangelogDetailed(string(content))

	// Display results based on flags
	if validateQuiet {
		return handleQuietMode(result, output)
	}

	if validateVerbose {
		handleVerboseMode(result, output, changelogFile, string(content))
	} else {
		handleStandardMode(result, output)
	}

	if !result.Valid {
		return fmt.Errorf("changelog validation failed")
	}

	return nil
}

func handleQuietMode(result *changelog.ValidationResult, output *cli.OutputHandler) error {
	if !result.Valid {
		for _, err := range result.Errors {
			output.PrintError(err)
		}
		return fmt.Errorf("changelog validation failed")
	}
	return nil
}

func handleVerboseMode(result *changelog.ValidationResult, output *cli.OutputHandler, file, content string) {
	output.PrintInfo(fmt.Sprintf("Validating changelog: %s", file))
	fmt.Println()

	if result.Valid {
		output.PrintSuccess("Changelog validation passed")
	} else {
		output.PrintError("Changelog validation failed")
	}
	fmt.Println()

	printErrors(result.Errors, output)
	printWarnings(result.Warnings, output)

	if len(result.Errors) == 0 && len(result.Warnings) == 0 {
		output.PrintSuccess("🎉 No issues found!")
	}

	printStructureAnalysis(content, output)
	printVersionAnalysis(content, output)
}

func handleStandardMode(result *changelog.ValidationResult, output *cli.OutputHandler) {
	if result.Valid {
		output.PrintSuccess("Changelog validation passed")
		if len(result.Warnings) > 0 {
			output.PrintWarning(fmt.Sprintf("%d warnings found", len(result.Warnings)))
			for _, warning := range result.Warnings {
				fmt.Printf("  • %s\n", warning)
			}
		}
	} else {
		output.PrintError("Changelog validation failed")
		for _, err := range result.Errors {
			fmt.Printf("  • %s\n", err)
		}
	}
}

func printErrors(errors []string, output *cli.OutputHandler) {
	if len(errors) > 0 {
		output.PrintError("🚨 Errors:")
		for _, err := range errors {
			fmt.Printf("  • %s\n", err)
		}
		fmt.Println()
	}
}

func printWarnings(warnings []string, output *cli.OutputHandler) {
	if len(warnings) > 0 {
		output.PrintWarning("⚠️  Warnings:")
		for _, warning := range warnings {
			fmt.Printf("  • %s\n", warning)
		}
		fmt.Println()
	}
}

func printStructureAnalysis(content string, output *cli.OutputHandler) {
	output.PrintInfo("📊 Structure Analysis:")
	extractor := changelog.NewSectionExtractor(content)
	sections, err := extractor.ExtractAllSections()
	if err == nil {
		fmt.Printf("  • Found %d sections\n", len(sections))
		for _, section := range sections {
			fmt.Printf("    - %s: %d items\n", section.Title, len(section.Items))
		}
	}
}

func printVersionAnalysis(content string, output *cli.OutputHandler) {
	parser := changelog.NewChangelogParser(content)
	versions := parser.GetVersions()
	if len(versions) > 0 {
		fmt.Printf("  • Found %d versions\n", len(versions))
		for _, version := range versions {
			fmt.Printf("    - %s (%s)\n", version.Version, version.Date)
		}
	}
}
