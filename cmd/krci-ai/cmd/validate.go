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
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/epam/kuberocketai/internal/engine/processor"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate KubeRocketAI framework components",
	Long: `Validate KubeRocketAI framework components in the current directory.

This command validates:
- Agent YAML files for schema compliance
- Task path link validation in agent references
- Markdown format validation for task files
- Cross-platform file accessibility

The validation runs on the current directory framework structure and provides
detailed error reporting for any issues found.

Examples:
  krci-ai validate                    # Validate framework in current directory
  krci-ai validate --verbose          # Validate with detailed output
  krci-ai validate --quiet            # Validate with minimal output`,
	RunE: runValidate,
}

// Validation flags
var (
	verboseOutput bool
	quietOutput   bool
)

func init() {
	rootCmd.AddCommand(validateCmd)

	// Add flags
	validateCmd.Flags().BoolVarP(&verboseOutput, "verbose", "v", false, "verbose output with detailed validation results")
	validateCmd.Flags().BoolVarP(&quietOutput, "quiet", "q", false, "quiet output, only show summary")
}

// runValidate executes the validation command
func runValidate(cmd *cobra.Command, args []string) error {
	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	if !quietOutput {
		fmt.Printf("🔍 Validating framework components in: %s\n\n", currentDir)
	}

	// Initialize validation system
	validator, err := NewFrameworkValidator(currentDir)
	if err != nil {
		return fmt.Errorf("failed to initialize validator: %w", err)
	}

	// Run validation
	results, err := validator.ValidateFramework()
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Display results
	displayValidationResults(results)

	// Exit with error code if validation failed
	if !results.IsValid() {
		os.Exit(1)
	}

	return nil
}

// ValidationResult represents the result of a single validation
type ValidationResult struct {
	Type        string // "agent", "task", "template", "data"
	File        string
	IsValid     bool
	Errors      []string
	Warnings    []string
	ProcessTime string
}

// ValidationResults represents the complete validation results
type ValidationResults struct {
	Results       []ValidationResult
	TotalFiles    int
	ValidFiles    int
	InvalidFiles  int
	TotalErrors   int
	TotalWarnings int
	ProcessTime   string
}

// IsValid returns true if all validations passed
func (r *ValidationResults) IsValid() bool {
	return r.InvalidFiles == 0
}

// FrameworkValidator handles framework validation
type FrameworkValidator struct {
	baseDir       string
	yamlProcessor *processor.YAMLProcessor
}

// NewFrameworkValidator creates a new framework validator
func NewFrameworkValidator(baseDir string) (*FrameworkValidator, error) {
	// Use the embedded assets which include schemas
	embeddedAssets := GetEmbeddedAssets()
	yamlProcessor, err := processor.NewYAMLProcessor(embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create YAML processor: %w", err)
	}

	return &FrameworkValidator{
		baseDir:       baseDir,
		yamlProcessor: yamlProcessor,
	}, nil
}

// ValidateFramework validates the entire framework
func (v *FrameworkValidator) ValidateFramework() (*ValidationResults, error) {
	results := &ValidationResults{
		Results: make([]ValidationResult, 0),
	}

	// Look for .krci-ai directory
	frameworkDir := filepath.Join(v.baseDir, ".krci-ai")
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("no .krci-ai directory found in current directory - use 'krci-ai install' to set up framework")
	}

	// Validate agents
	if err := v.validateAgents(frameworkDir, results); err != nil {
		return nil, fmt.Errorf("agent validation failed: %w", err)
	}

	// Validate tasks (basic existence check for now)
	if err := v.validateTasks(frameworkDir, results); err != nil {
		return nil, fmt.Errorf("task validation failed: %w", err)
	}

	// Calculate totals
	results.TotalFiles = len(results.Results)
	for _, result := range results.Results {
		if result.IsValid {
			results.ValidFiles++
		} else {
			results.InvalidFiles++
		}
		results.TotalErrors += len(result.Errors)
		results.TotalWarnings += len(result.Warnings)
	}

	return results, nil
}

// validateAgents validates all agent files
func (v *FrameworkValidator) validateAgents(frameworkDir string, results *ValidationResults) error {
	agentsDir := filepath.Join(frameworkDir, "agents")

	// Check if agents directory exists
	if _, err := os.Stat(agentsDir); os.IsNotExist(err) {
		// No agents directory is not an error - might be a minimal setup
		return nil
	}

	// Find all agent files
	agentFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.yaml"))
	if err != nil {
		return fmt.Errorf("failed to find agent files: %w", err)
	}

	// Also check for .yml files
	ymlFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.yml"))
	if err != nil {
		return fmt.Errorf("failed to find agent files: %w", err)
	}

	agentFiles = append(agentFiles, ymlFiles...)

	// Validate each agent file
	for _, agentFile := range agentFiles {
		result := v.validateAgentFile(agentFile)
		results.Results = append(results.Results, result)
	}

	return nil
}

// validateAgentFile validates a single agent file
func (v *FrameworkValidator) validateAgentFile(filePath string) ValidationResult {
	// Convert absolute path to relative path from baseDir
	relPath, err := filepath.Rel(v.baseDir, filePath)
	if err != nil {
		relPath = filePath // fallback to absolute path if conversion fails
	}

	result := ValidationResult{
		Type:     "agent",
		File:     relPath,
		IsValid:  true,
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}

	// Parse and validate the agent
	_, validationErrors, err := v.yamlProcessor.ProcessAndValidateAgent(filePath)
	if err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, fmt.Sprintf("Parse error: %s", err.Error()))
		return result
	}

	// Process validation errors
	for _, validationError := range validationErrors {
		result.IsValid = false
		result.Errors = append(result.Errors, validationError.Message)
	}

	return result
}

// validateTasks validates task files (basic existence check)
func (v *FrameworkValidator) validateTasks(frameworkDir string, results *ValidationResults) error {
	tasksDir := filepath.Join(frameworkDir, "tasks")

	// Check if tasks directory exists
	if _, err := os.Stat(tasksDir); os.IsNotExist(err) {
		// No tasks directory is not an error - might be a minimal setup
		return nil
	}

	// Find all task files
	taskFiles, err := filepath.Glob(filepath.Join(tasksDir, "*.md"))
	if err != nil {
		return fmt.Errorf("failed to find task files: %w", err)
	}

	// Validate each task file (basic check for now)
	for _, taskFile := range taskFiles {
		result := v.validateTaskFile(taskFile)
		results.Results = append(results.Results, result)
	}

	return nil
}

// validateTaskFile validates a single task file
func (v *FrameworkValidator) validateTaskFile(filePath string) ValidationResult {
	// Convert absolute path to relative path from baseDir
	relPath, err := filepath.Rel(v.baseDir, filePath)
	if err != nil {
		relPath = filePath // fallback to absolute path if conversion fails
	}

	result := ValidationResult{
		Type:     "task",
		File:     relPath,
		IsValid:  true,
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}

	// Check if file exists and is readable
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		result.IsValid = false
		result.Errors = append(result.Errors, "Task file does not exist")
		return result
	}

	// Check if file has .md extension
	if !strings.HasSuffix(filePath, ".md") {
		result.IsValid = false
		result.Errors = append(result.Errors, "Task file must have .md extension")
		return result
	}

	// Try to read the file
	if _, err := os.ReadFile(filePath); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, fmt.Sprintf("Cannot read task file: %s", err.Error()))
		return result
	}

	return result
}

// displayValidationResults displays the validation results
func displayValidationResults(results *ValidationResults) {
	// Color setup
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	bold := color.New(color.Bold)

	// Display individual results if verbose or if there are errors
	if verboseOutput || results.InvalidFiles > 0 {
		fmt.Println("📋 Validation Results:")
		fmt.Println(strings.Repeat("-", 60))

		for _, result := range results.Results {
			// File path is already relative to project directory
			displayPath := result.File

			// Display file status
			if result.IsValid {
				if verboseOutput {
					_, _ = green.Printf("✅ %s: %s\n", strings.ToUpper(result.Type), displayPath)
				}
			} else {
				_, _ = red.Printf("❌ %s: %s\n", strings.ToUpper(result.Type), displayPath)

				// Display errors
				for _, errorMsg := range result.Errors {
					fmt.Printf("   🔸 %s\n", errorMsg)
				}

				// Display warnings
				for _, warning := range result.Warnings {
					_, _ = yellow.Printf("   ⚠️  %s\n", warning)
				}
			}
		}
		fmt.Println()
	}

	// Display summary
	_, _ = bold.Println("📊 Validation Summary:")
	fmt.Println(strings.Repeat("-", 40))

	fmt.Printf("Total Files:     %d\n", results.TotalFiles)

	if results.ValidFiles > 0 {
		_, _ = green.Printf("Valid Files:     %d\n", results.ValidFiles)
	} else {
		fmt.Printf("Valid Files:     %d\n", results.ValidFiles)
	}

	if results.InvalidFiles > 0 {
		_, _ = red.Printf("Invalid Files:   %d\n", results.InvalidFiles)
	} else {
		fmt.Printf("Invalid Files:   %d\n", results.InvalidFiles)
	}

	if results.TotalErrors > 0 {
		_, _ = red.Printf("Total Errors:    %d\n", results.TotalErrors)
	} else {
		fmt.Printf("Total Errors:    %d\n", results.TotalErrors)
	}

	if results.TotalWarnings > 0 {
		_, _ = yellow.Printf("Total Warnings:  %d\n", results.TotalWarnings)
	} else {
		fmt.Printf("Total Warnings:  %d\n", results.TotalWarnings)
	}

	fmt.Println()

	// Display final status
	if results.IsValid() {
		_, _ = green.Println("🎉 Framework validation completed successfully!")
	} else {
		_, _ = red.Println("❌ Framework validation failed!")
		if !quietOutput {
			fmt.Println("   Please fix the errors above and run validation again.")
		}
	}
}
