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
package validation

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ErrorCode provides programmatic error identification
type ErrorCode string

const (
	// Critical errors
	ErrorCodeMissingTask       ErrorCode = "MISSING_TASK"
	ErrorCodeBrokenLink        ErrorCode = "BROKEN_LINK"
	ErrorCodeInvalidFormat     ErrorCode = "INVALID_FORMAT"
	ErrorCodeArchViolation     ErrorCode = "ARCH_VIOLATION"
	ErrorCodeFrontmatterParse  ErrorCode = "FRONTMATTER_PARSE"
	ErrorCodeFrontmatterSchema ErrorCode = "FRONTMATTER_SCHEMA"
	ErrorCodeMissingDependency ErrorCode = "MISSING_DEPENDENCY"
	ErrorCodeFrameworkMissing  ErrorCode = "FRAMEWORK_MISSING"
	ErrorCodeParseError        ErrorCode = "PARSE_ERROR"

	// Warning errors
	ErrorCodeOrphanedFile       ErrorCode = "ORPHANED_FILE"
	ErrorCodeCircularDependency ErrorCode = "CIRCULAR_DEPENDENCY"
	ErrorCodeInvalidExtension   ErrorCode = "INVALID_EXTENSION"
)

// ValidationResult represents a unified validation result
type ValidationResult struct {
	Severity    Severity  `json:"severity"`
	Type        string    `json:"type"`
	Code        ErrorCode `json:"code"`
	File        string    `json:"file"`
	Line        int       `json:"line"`
	Message     string    `json:"message"`
	Context     any       `json:"context,omitempty"`
	FixGuidance string    `json:"fix_guidance"`
}

// UnifiedValidationReport contains comprehensive validation results
type UnifiedValidationReport struct {
	Results     []ValidationResult `json:"results"`
	Summary     ValidationSummary  `json:"summary"`
	Performance PerformanceMetrics `json:"performance"`
}

// ValidationSummary provides aggregated validation statistics
type ValidationSummary struct {
	TotalIssues   int `json:"total_issues"`
	CriticalCount int `json:"critical_count"`
	ErrorCount    int `json:"error_count"`
	WarningCount  int `json:"warning_count"`
	InfoCount     int `json:"info_count"`
}

// PerformanceMetrics tracks validation performance
type PerformanceMetrics struct {
	ValidationTimeMs int64   `json:"validation_time_ms"`
	FilesProcessed   int     `json:"files_processed"`
	CacheHitRate     float64 `json:"cache_hit_rate"`
}

// SchemaType represents different schema types
type SchemaType string

const (
	SchemaTypeAgent       SchemaType = "agent"
	SchemaTypeFrontmatter SchemaType = "frontmatter"
	SchemaTypeTemplate    SchemaType = "template"
	SchemaTypeData        SchemaType = "data"
)

// ValidationConfig holds all validation dependencies
type ValidationConfig struct {
	SchemaRegistry    SchemaRegistry
	ParserRegistry    ParserRegistry
	ValidatorRegistry ValidatorRegistry
	AssetResolver     AssetResolver
	ValidationRules   []ValidationRule
}

// Validate ensures all required dependencies are present
func (c ValidationConfig) Validate() error {
	if c.SchemaRegistry == nil {
		return fmt.Errorf("schemaRegistry is required")
	}
	if c.ParserRegistry == nil {
		return fmt.Errorf("parserRegistry is required")
	}
	if c.ValidatorRegistry == nil {
		return fmt.Errorf("validatorRegistry is required")
	}
	if c.AssetResolver == nil {
		return fmt.Errorf("assetResolver is required")
	}
	return nil
}

// UniversalValidator defines the unified validation interface
type UniversalValidator interface {
	ValidateAgent(filePath string) ValidationResult
	ValidateTask(filePath string) ValidationResult
	ValidateTemplate(filePath string) ValidationResult
	ValidateFramework(frameworkPath string) UnifiedValidationReport
	ValidateBatch(filePaths []string) UnifiedValidationReport
}

// SchemaRegistry provides centralized schema management
type SchemaRegistry interface {
	GetSchema(schemaType SchemaType) (JSONSchema, error)
	ValidateAgainstSchema(data any, schemaType SchemaType) error
	LoadSchemas(embeddedAssets embed.FS) error
}

// JSONSchema represents a compiled JSON schema
type JSONSchema interface {
	Validate(data any) error
}

// ParserRegistry manages different file type parsers
type ParserRegistry interface {
	GetParser(filePath string) (FileParser, error)
	RegisterParser(extension string, parser FileParser)
	CanParse(filePath string) bool
}

// FileParser provides unified parsing interface
type FileParser interface {
	ParseFile(filePath string) (ParsedFile, error)
	ParseContent(content string, filePath string) (ParsedFile, error)
	GetFileType() string
}

// ParsedFile represents a parsed file with structured data
type ParsedFile struct {
	Type     string `json:"type"`
	Path     string `json:"path"`
	Metadata any    `json:"metadata"`
	Content  string `json:"content"`
}

// ValidatorRegistry manages different validation strategies
type ValidatorRegistry interface {
	GetValidators(fileType string) []ValidatorStrategy
	RegisterValidator(fileType string, validator ValidatorStrategy)
}

// ValidatorStrategy implements specific validation logic
type ValidatorStrategy interface {
	Name() string
	CanValidate(filePath string) bool
	Validate(parsed ParsedFile) []ValidationResult
	GetSeverity() Severity
}

// AssetResolver handles cross-asset dependency resolution
type AssetResolver interface {
	ResolvePath(relativePath string) (string, error)
	ValidateDependencies(dependencies map[string][]string) ([]string, error)
	SetSearchPaths(paths []string)
}

// ValidationRule defines configurable validation rules
type ValidationRule interface {
	Name() string
	Apply(parsed ParsedFile) []ValidationResult
	IsEnabled() bool
	GetPriority() int
}

// FrameworkValidator implements UniversalValidator with dependency injection
type FrameworkValidator struct {
	config         ValidationConfig
	embeddedAssets embed.FS
}

// NewFrameworkValidator creates a new validator with explicit dependencies
func NewFrameworkValidator(config ValidationConfig) (*FrameworkValidator, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid validation config: %w", err)
	}

	// Extract embedded assets from parser registry
	var embeddedAssets embed.FS
	if integratedRegistry, ok := config.ParserRegistry.(*IntegratedParserRegistry); ok {
		embeddedAssets = integratedRegistry.embeddedAssets
	}

	return &FrameworkValidator{
		config:         config,
		embeddedAssets: embeddedAssets,
	}, nil
}

// ValidateAgent validates a single agent file
func (v *FrameworkValidator) ValidateAgent(filePath string) ValidationResult {
	parser, err := v.config.ParserRegistry.GetParser(filePath)
	if err != nil {
		return ValidationResult{
			Severity:    SeverityError,
			Type:        "parser_error",
			Code:        ErrorCodeInvalidFormat,
			File:        filePath,
			Message:     fmt.Sprintf("Failed to get parser: %s", err.Error()),
			FixGuidance: "Ensure file has a supported extension and format",
		}
	}

	parsed, err := parser.ParseFile(filePath)
	if err != nil {
		return ValidationResult{
			Severity:    SeverityCritical,
			Type:        "parse_error",
			Code:        ErrorCodeInvalidFormat,
			File:        filePath,
			Message:     fmt.Sprintf("Failed to parse file: %s", err.Error()),
			FixGuidance: "Fix file format and syntax errors",
		}
	}

	validators := v.config.ValidatorRegistry.GetValidators(parsed.Type)
	var results []ValidationResult

	for _, validator := range validators {
		if validator.CanValidate(filePath) {
			validationResults := validator.Validate(parsed)
			results = append(results, validationResults...)
		}
	}

	// Return the most severe result, or success if no issues
	if len(results) == 0 {
		return ValidationResult{
			Severity:    SeverityInfo,
			Type:        "validation_success",
			File:        filePath,
			Message:     "File validation passed",
			FixGuidance: "",
		}
	}

	// Return the most critical result
	mostCritical := results[0]
	for _, result := range results[1:] {
		if result.Severity > mostCritical.Severity {
			mostCritical = result
		}
	}

	return mostCritical
}

// ValidateTask validates a single task file
func (v *FrameworkValidator) ValidateTask(filePath string) ValidationResult {
	return v.ValidateAgent(filePath) // Use same logic for now
}

// ValidateTemplate validates a single template file
func (v *FrameworkValidator) ValidateTemplate(filePath string) ValidationResult {
	return v.ValidateAgent(filePath) // Use same logic for now
}

// ValidateFramework validates an entire framework
func (v *FrameworkValidator) ValidateFramework(frameworkPath string) UnifiedValidationReport {
	start := time.Now()
	var allResults []ValidationResult

	// Check if framework directory exists
	frameworkDir := filepath.Join(frameworkPath, ".krci-ai")
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		allResults = append(allResults, ValidationResult{
			Severity:    SeverityCritical,
			Type:        "framework_missing",
			Code:        ErrorCodeFrameworkMissing,
			File:        ".krci-ai",
			Message:     "Framework directory not found",
			FixGuidance: "Run 'krci-ai install' to set up the framework",
		})
		return v.generateReport(allResults, start)
	}

	// Validate agents
	agentsDir := filepath.Join(frameworkDir, "agents")
	if agentFiles, err := v.findFiles(agentsDir, []string{".yaml", ".yml"}); err == nil {
		for _, agentFile := range agentFiles {
			result := v.ValidateAgent(agentFile)
			allResults = append(allResults, result)
		}
	}

	// Validate tasks with frontmatter only
	tasksDir := filepath.Join(frameworkDir, "tasks")
	if taskFiles, err := v.findFiles(tasksDir, []string{".md"}); err == nil {
		for _, taskFile := range taskFiles {
			results := v.validateTaskWithFrontmatter(taskFile)
			allResults = append(allResults, results...)
		}
	}

	// Validate local tasks
	localTasksDir := filepath.Join(frameworkDir, "local", "tasks")
	if taskFiles, err := v.findFiles(localTasksDir, []string{".md"}); err == nil {
		for _, taskFile := range taskFiles {
			results := v.validateTaskWithFrontmatter(taskFile)
			allResults = append(allResults, results...)
		}
	}

	return v.generateReport(allResults, start)
}

// findFiles finds all files with given extensions in a directory
func (v *FrameworkValidator) findFiles(dir string, extensions []string) ([]string, error) {
	var files []string

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return files, nil // Directory doesn't exist, return empty slice
	}

	for _, ext := range extensions {
		pattern := filepath.Join(dir, "*"+ext)
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}
		files = append(files, matches...)
	}

	return files, nil
}

// validateTaskWithFrontmatter validates a task file using frontmatter validation only
func (v *FrameworkValidator) validateTaskWithFrontmatter(taskPath string) []ValidationResult {
	// Parse the file first
	parser, err := v.config.ParserRegistry.GetParser(taskPath)
	if err != nil {
		return []ValidationResult{{
			Severity:    SeverityCritical,
			Type:        "parser_error",
			Code:        ErrorCodeParseError,
			File:        taskPath,
			Message:     fmt.Sprintf("Failed to get parser for task: %v", err),
			FixGuidance: "Ensure the validation system is properly configured",
		}}
	}

	parsed, err := parser.ParseFile(taskPath)
	if err != nil {
		return []ValidationResult{{
			Severity:    SeverityCritical,
			Type:        "parse_error",
			Code:        ErrorCodeFrontmatterParse,
			File:        taskPath,
			Message:     fmt.Sprintf("Failed to parse task file: %v", err),
			FixGuidance: "Check YAML frontmatter syntax and structure",
		}}
	}

	// Get validators for markdown files (tasks)
	validators := v.config.ValidatorRegistry.GetValidators("markdown")
	if len(validators) == 0 {
		return []ValidationResult{{
			Severity:    SeverityWarning,
			Type:        "no_validator",
			Code:        ErrorCodeParseError,
			File:        taskPath,
			Message:     "No validators found for markdown files",
			FixGuidance: "Ensure the validation system is properly configured with markdown validators",
		}}
	}

	var allResults []ValidationResult
	for _, validator := range validators {
		if validator.CanValidate(taskPath) {
			results := validator.Validate(parsed)
			allResults = append(allResults, results...)
		}
	}

	return allResults
}

// generateReport generates a unified validation report from results
func (v *FrameworkValidator) generateReport(results []ValidationResult, start time.Time) UnifiedValidationReport {
	return UnifiedValidationReport{
		Results: results,
		Summary: v.generateSummary(results),
		Performance: PerformanceMetrics{
			FilesProcessed:   len(results),
			ValidationTimeMs: time.Since(start).Milliseconds(),
		},
	}
}

// ValidateBatch validates multiple files in parallel
func (v *FrameworkValidator) ValidateBatch(filePaths []string) UnifiedValidationReport {
	ctx := context.Background()
	results := make([]ValidationResult, 0, len(filePaths))

	// For now, validate sequentially - parallel processing in Phase 3
	for _, filePath := range filePaths {
		select {
		case <-ctx.Done():
			return UnifiedValidationReport{
				Results: results,
				Summary: v.generateSummary(results),
				Performance: PerformanceMetrics{
					FilesProcessed: len(results),
				},
			}
		default:
			result := v.ValidateAgent(filePath)
			results = append(results, result)
		}
	}

	return UnifiedValidationReport{
		Results: results,
		Summary: v.generateSummary(results),
		Performance: PerformanceMetrics{
			FilesProcessed: len(results),
		},
	}
}

// generateSummary creates validation summary from results
func (v *FrameworkValidator) generateSummary(results []ValidationResult) ValidationSummary {
	summary := ValidationSummary{
		TotalIssues: len(results),
	}

	for _, result := range results {
		switch result.Severity {
		case SeverityCritical:
			summary.CriticalCount++
		case SeverityError:
			summary.ErrorCount++
		case SeverityWarning:
			summary.WarningCount++
		case SeverityInfo:
			summary.InfoCount++
		}
	}

	return summary
}
