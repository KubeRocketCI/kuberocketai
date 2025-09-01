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
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KubeRocketCI/kuberocketai/internal/engine/processor"
	"github.com/KubeRocketCI/kuberocketai/internal/metadata"
)

const (
	// File type constants to avoid goconst issues
	fileTypeTask     = "task"
	fileTypeTemplate = "template"
	fileTypeData     = "data"
	fileTypeMarkdown = "markdown"
	fileTypeAgent    = string(SchemaTypeAgent)
)

// IntegratedYAMLParser integrates with existing YAMLProcessor
type IntegratedYAMLParser struct {
	processor      *processor.YAMLProcessor
	embeddedAssets embed.FS
}

// NewIntegratedYAMLParser creates a parser that uses existing YAMLProcessor
func NewIntegratedYAMLParser(embeddedAssets embed.FS) (*IntegratedYAMLParser, error) {
	yamlProcessor, err := processor.NewYAMLProcessor(embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create YAML processor: %w", err)
	}

	return &IntegratedYAMLParser{
		processor:      yamlProcessor,
		embeddedAssets: embeddedAssets,
	}, nil
}

// ParseFile parses a YAML agent file using existing processor
func (p *IntegratedYAMLParser) ParseFile(filePath string) (ParsedFile, error) {
	agent, err := p.processor.ParseAgentFile(filePath)
	if err != nil {
		return ParsedFile{}, fmt.Errorf("failed to parse YAML file: %w", err)
	}

	return ParsedFile{
		Type:     "agent",
		Path:     filePath,
		Metadata: agent,
		Content:  "", // YAML files don't separate content
	}, nil
}

// ParseContent parses YAML content from string
func (p *IntegratedYAMLParser) ParseContent(content string, filePath string) (ParsedFile, error) {
	// Create temporary file for existing processor (limitation of current API)
	tmpFile, err := os.CreateTemp("", "yaml_parse_*.yaml")
	if err != nil {
		return ParsedFile{}, fmt.Errorf("failed to create temp file: %w", err)
	}

	tmpFileName := tmpFile.Name()
	defer func() {
		if closeErr := tmpFile.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close temp file %s: %v\n", tmpFileName, closeErr)
		}
		if removeErr := os.Remove(tmpFileName); removeErr != nil {
			fmt.Printf("Warning: failed to remove temp file %s: %v\n", tmpFileName, removeErr)
		}
	}()

	if _, err := tmpFile.WriteString(content); err != nil {
		return ParsedFile{}, fmt.Errorf("failed to write temp file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return ParsedFile{}, fmt.Errorf("failed to close temp file: %w", err)
	}

	return p.ParseFile(tmpFileName)
}

// GetFileType returns the file type this parser handles
func (p *IntegratedYAMLParser) GetFileType() string {
	return fileTypeAgent
}

// IntegratedMarkdownParser integrates with existing FrontmatterParser
type IntegratedMarkdownParser struct {
	frontmatterParser *metadata.FrontmatterParser
	embeddedAssets    embed.FS
}

// NewIntegratedMarkdownParser creates a parser that uses existing FrontmatterParser
func NewIntegratedMarkdownParser(embeddedAssets embed.FS) (*IntegratedMarkdownParser, error) {
	frontmatterParser, err := metadata.NewFrontmatterParser(embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create frontmatter parser: %w", err)
	}

	return &IntegratedMarkdownParser{
		frontmatterParser: frontmatterParser,
		embeddedAssets:    embeddedAssets,
	}, nil
}

// ParseFile parses a markdown file with potential frontmatter
func (p *IntegratedMarkdownParser) ParseFile(filePath string) (ParsedFile, error) {
	fileType := p.determineFileType(filePath)

	if fileType == fileTypeTask {
		// Use FrontmatterParser for task files
		parsed, err := p.frontmatterParser.ParseTaskFile(filePath)
		if err != nil {
			return ParsedFile{}, fmt.Errorf("failed to parse task file: %w", err)
		}

		return ParsedFile{
			Type:     fileType,
			Path:     filePath,
			Metadata: parsed.Metadata,
			Content:  parsed.Content,
		}, nil
	}

	// For templates and data files, just read content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return ParsedFile{}, fmt.Errorf("failed to read file: %w", err)
	}

	return ParsedFile{
		Type:     fileType,
		Path:     filePath,
		Metadata: nil,
		Content:  string(content),
	}, nil
}

// ParseContent parses markdown content from string
func (p *IntegratedMarkdownParser) ParseContent(content string, filePath string) (ParsedFile, error) {
	fileType := p.determineFileType(filePath)

	if fileType == fileTypeTask {
		// Use FrontmatterParser for task content
		parsed, err := p.frontmatterParser.ParseTaskContent(content, filePath)
		if err != nil {
			return ParsedFile{}, fmt.Errorf("failed to parse task content: %w", err)
		}

		return ParsedFile{
			Type:     fileType,
			Path:     filePath,
			Metadata: parsed.Metadata,
			Content:  parsed.Content,
		}, nil
	}

	return ParsedFile{
		Type:     fileType,
		Path:     filePath,
		Metadata: nil,
		Content:  content,
	}, nil
}

// GetFileType returns the file type this parser handles
func (p *IntegratedMarkdownParser) GetFileType() string {
	return fileTypeMarkdown
}

// determineFileType determines file type based on path
func (p *IntegratedMarkdownParser) determineFileType(filePath string) string {
	if strings.Contains(filePath, "/tasks/") || strings.Contains(filePath, "/local/tasks/") {
		return fileTypeTask
	}
	if strings.Contains(filePath, "/templates/") || strings.Contains(filePath, "/local/templates/") {
		return fileTypeTemplate
	}
	if strings.Contains(filePath, "/data/") || strings.Contains(filePath, "/local/data/") {
		return fileTypeData
	}
	return fileTypeMarkdown
}

// IntegratedParserRegistry creates parsers with existing processors
type IntegratedParserRegistry struct {
	parsers        map[string]FileParser
	embeddedAssets embed.FS
}

// NewIntegratedParserRegistry creates a parser registry with integrated parsers
func NewIntegratedParserRegistry(embeddedAssets embed.FS) (*IntegratedParserRegistry, error) {
	registry := &IntegratedParserRegistry{
		parsers:        make(map[string]FileParser),
		embeddedAssets: embeddedAssets,
	}

	// Create integrated YAML parser
	yamlParser, err := NewIntegratedYAMLParser(embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create YAML parser: %w", err)
	}

	// Create integrated markdown parser
	markdownParser, err := NewIntegratedMarkdownParser(embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create markdown parser: %w", err)
	}

	// Register parsers
	registry.RegisterParser(".yaml", yamlParser)
	registry.RegisterParser(".yml", yamlParser)
	registry.RegisterParser(".md", markdownParser)

	return registry, nil
}

// RegisterParser registers a parser for a file extension
func (r *IntegratedParserRegistry) RegisterParser(extension string, parser FileParser) {
	r.parsers[extension] = parser
}

// GetParser returns a parser for the given file path
func (r *IntegratedParserRegistry) GetParser(filePath string) (FileParser, error) {
	ext := filepath.Ext(filePath)
	parser, exists := r.parsers[ext]
	if !exists {
		return nil, fmt.Errorf("no parser registered for extension: %s", ext)
	}

	return parser, nil
}

// CanParse returns true if a parser exists for the file
func (r *IntegratedParserRegistry) CanParse(filePath string) bool {
	ext := filepath.Ext(filePath)
	_, exists := r.parsers[ext]
	return exists
}

// IntegratedValidatorStrategy validates using existing processors
type IntegratedValidatorStrategy struct {
	yamlProcessor     *processor.YAMLProcessor
	frontmatterParser *metadata.FrontmatterParser
	embeddedAssets    embed.FS
}

// NewIntegratedValidatorStrategy creates a validator using existing processors
func NewIntegratedValidatorStrategy(embeddedAssets embed.FS) (*IntegratedValidatorStrategy, error) {
	yamlProcessor, err := processor.NewYAMLProcessor(embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create YAML processor: %w", err)
	}

	frontmatterParser, err := metadata.NewFrontmatterParser(embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create frontmatter parser: %w", err)
	}

	return &IntegratedValidatorStrategy{
		yamlProcessor:     yamlProcessor,
		frontmatterParser: frontmatterParser,
		embeddedAssets:    embeddedAssets,
	}, nil
}

// Name returns the validator name
func (v *IntegratedValidatorStrategy) Name() string {
	return "integrated_validator"
}

// CanValidate returns true if the file can be validated
func (v *IntegratedValidatorStrategy) CanValidate(filePath string) bool {
	ext := filepath.Ext(filePath)
	return ext == ".yaml" || ext == ".yml" || ext == ".md"
}

// Validate validates a parsed file using appropriate processor
func (v *IntegratedValidatorStrategy) Validate(parsed ParsedFile) []ValidationResult {
	var results []ValidationResult

	switch parsed.Type {
	case fileTypeAgent:
		results = append(results, v.validateAgent(parsed)...)
	case fileTypeTask:
		results = append(results, v.validateTask(parsed)...)
	case fileTypeTemplate, fileTypeData:
		// Basic validation for templates and data files
		results = append(results, v.validateMarkdown(parsed)...)
	}

	return results
}

// validateAgent validates agent files using YAMLProcessor
func (v *IntegratedValidatorStrategy) validateAgent(parsed ParsedFile) []ValidationResult {
	var results []ValidationResult

	// Use existing YAML processor validation
	if agent, ok := parsed.Metadata.(*processor.Agent); ok {
		validationErrors := v.yamlProcessor.ValidateAgent(agent, parsed.Path)

		for _, err := range validationErrors {
			results = append(results, ValidationResult{
				Severity:    SeverityCritical,
				Type:        "agent_validation_error",
				Code:        ErrorCodeFrontmatterSchema,
				File:        parsed.Path,
				Message:     err.Message,
				FixGuidance: "Fix the agent YAML structure according to schema requirements",
			})
		}
	}

	return results
}

// validateTask validates task files using FrontmatterParser
func (v *IntegratedValidatorStrategy) validateTask(parsed ParsedFile) []ValidationResult {
	var results []ValidationResult

	if taskMetadata, ok := parsed.Metadata.(*metadata.TaskMetadata); ok {
		// Validate metadata against schema
		if err := v.frontmatterParser.ValidateMetadata(taskMetadata, parsed.Path); err != nil {
			results = append(results, ValidationResult{
				Severity:    SeverityCritical,
				Type:        "frontmatter_schema_error",
				Code:        ErrorCodeFrontmatterSchema,
				File:        parsed.Path,
				Message:     fmt.Sprintf("Frontmatter validation failed: %s", err.Error()),
				FixGuidance: "Fix frontmatter structure to match schema requirements",
			})
		}
	}

	return results
}

// validateMarkdown validates basic markdown files
func (v *IntegratedValidatorStrategy) validateMarkdown(parsed ParsedFile) []ValidationResult {
	var results []ValidationResult

	// Basic validation - check if content is readable
	if strings.TrimSpace(parsed.Content) == "" {
		results = append(results, ValidationResult{
			Severity:    SeverityWarning,
			Type:        "empty_file",
			Code:        ErrorCodeInvalidFormat,
			File:        parsed.Path,
			Message:     "File appears to be empty",
			FixGuidance: "Add content to the file or remove if not needed",
		})
	}

	return results
}

// GetSeverity returns the validator severity
func (v *IntegratedValidatorStrategy) GetSeverity() Severity {
	return SeverityCritical
}

// IntegratedValidatorRegistry manages integrated validators
type IntegratedValidatorRegistry struct {
	validators     map[string][]ValidatorStrategy
	embeddedAssets embed.FS
}

// NewIntegratedValidatorRegistry creates a validator registry with integrated validators
func NewIntegratedValidatorRegistry(embeddedAssets embed.FS) (*IntegratedValidatorRegistry, error) {
	registry := &IntegratedValidatorRegistry{
		validators:     make(map[string][]ValidatorStrategy),
		embeddedAssets: embeddedAssets,
	}

	// Create integrated validator
	integratedValidator, err := NewIntegratedValidatorStrategy(embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create integrated validator: %w", err)
	}

	// Register for all supported file types
	registry.RegisterValidator("agent", integratedValidator)
	registry.RegisterValidator("task", integratedValidator)
	registry.RegisterValidator("template", integratedValidator)
	registry.RegisterValidator("data", integratedValidator)
	registry.RegisterValidator("markdown", integratedValidator)

	return registry, nil
}

// RegisterValidator registers a validator for a file type
func (r *IntegratedValidatorRegistry) RegisterValidator(fileType string, validator ValidatorStrategy) {
	if r.validators[fileType] == nil {
		r.validators[fileType] = make([]ValidatorStrategy, 0)
	}
	r.validators[fileType] = append(r.validators[fileType], validator)
}

// GetValidators returns all validators for a file type
func (r *IntegratedValidatorRegistry) GetValidators(fileType string) []ValidatorStrategy {
	validators, exists := r.validators[fileType]
	if !exists {
		return []ValidatorStrategy{}
	}

	// Return a copy to avoid concurrent modification
	result := make([]ValidatorStrategy, len(validators))
	copy(result, validators)
	return result
}
