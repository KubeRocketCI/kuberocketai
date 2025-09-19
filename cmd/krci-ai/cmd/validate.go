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
	"time"

	"github.com/spf13/cobra"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/cli"
	"github.com/KubeRocketCI/kuberocketai/internal/discovery"
	"github.com/KubeRocketCI/kuberocketai/internal/validation"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate KubeRocketAI framework components",
	Long: `Validate KubeRocketAI framework components in the current directory.

This command validates:
- Agent YAML files for schema compliance
- Task path link validation in agent references
- Template files structure and accessibility
- Markdown links to framework files ([text](./.krci-ai/path/file.md))
- Markdown format validation for task files
- Cross-platform file accessibility

The validation runs on the current directory framework structure and provides
detailed error reporting for any issues found.

Examples:
  krci-ai validate                    # Validate framework in current directory
  krci-ai validate --quiet            # Validate with minimal output`,
	RunE: runValidate,
}

// Validation flags removed - now using GetBool pattern

func init() {
	rootCmd.AddCommand(validateCmd)

	// Add flags
	validateCmd.Flags().BoolP("quiet", "q", false, "quiet output, only show summary")
}

// runValidate executes the validation command
func runValidate(cmd *cobra.Command, args []string) error {
	// Get flags
	quietOutput, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("failed to get quiet flag: %w", err)
	}

	startTime := time.Now()

	// Initialize enhanced validation system
	projectRoot, err := discovery.GetProjectRoot()
	if err != nil {
		return err
	}

	// Create discovery service for the framework directory
	frameworkDir := assets.GetKrciPath(projectRoot)
	discoveryService := assets.NewDiscovery(frameworkDir)

	// Create analyzer with discovery
	analyzer := validation.NewFrameworkAnalyzer(discoveryService)

	// Run optimized framework analysis with caching
	issues, insights, err := analyzer.AnalyzeFramework()
	if err != nil {
		return fmt.Errorf("framework analysis failed: %w", err)
	}

	processTime := time.Since(startTime)

	// Create output handler
	output := cli.NewOutputHandler()

	// Display validation results
	if !quietOutput {
		if len(issues) > 0 {
			output.PrintError("Framework validation failed")
			for _, issue := range issues {
				output.PrintError(fmt.Sprintf("- %s", issue.Message))
			}
			output.Newline()
		} else {
			output.PrintSuccess("Framework validation passed")
		}

		// Print framework insights
		output.PrintFrameworkInsights(insights, len(issues))

		output.Printf("⚡ Validation completed in %.1fs\n", processTime.Seconds())
	} else if len(issues) > 0 {
		output.Printf("❌ Framework validation failed with %d issues\n", len(issues))
	}

	// Return error for issues. Cobra root will handle exit.
	if len(issues) > 0 {
		return fmt.Errorf("validation failed with %d issues", len(issues))
	}
	return nil
}
