package metadata

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"gopkg.in/yaml.v3"
)

// TaskDependencies represents the dependency structure in YAML frontmatter
type TaskDependencies struct {
	Templates []string `yaml:"templates,omitempty" json:"templates,omitempty"`
	Data      []string `yaml:"data,omitempty" json:"data,omitempty"`
	MCP       []string `yaml:"mcp,omitempty" json:"mcp,omitempty"`
}

// TaskMetadata represents the complete frontmatter structure
type TaskMetadata struct {
	Dependencies *TaskDependencies `yaml:"dependencies,omitempty" json:"dependencies,omitempty"`
}

// ParsedTaskFile represents a parsed task file with separated frontmatter and content
type ParsedTaskFile struct {
	Metadata *TaskMetadata
	Content  string
	FilePath string
}

// FrontmatterParser handles parsing and validation of task frontmatter
type FrontmatterParser struct {
	schema *jsonschema.Schema
}

// NewFrontmatterParser creates a new frontmatter parser with embedded schema
func NewFrontmatterParser(embeddedAssets embed.FS) (*FrontmatterParser, error) {
	parser := &FrontmatterParser{}

	// Load embedded schema
	if err := parser.loadEmbeddedSchema(embeddedAssets); err != nil {
		return nil, fmt.Errorf("failed to load embedded schema: %w", err)
	}

	return parser, nil
}

// NewFrontmatterParserFromFile creates a new parser with schema from file (for testing)
func NewFrontmatterParserFromFile(schemaPath string) (*FrontmatterParser, error) {
	if schemaPath == "" {
		return nil, fmt.Errorf("schema path is required")
	}

	parser := &FrontmatterParser{}

	// Load schema from file
	if err := parser.loadSchemaFromFile(schemaPath); err != nil {
		return nil, fmt.Errorf("failed to load schema from file: %w", err)
	}

	return parser, nil
}

// loadEmbeddedSchema loads the JSON schema from embedded assets
func (p *FrontmatterParser) loadEmbeddedSchema(embeddedAssets embed.FS) error {
	schemaData, err := embeddedAssets.ReadFile("assets/schemas/task-metadata.json")
	if err != nil {
		return fmt.Errorf("failed to read embedded schema file: %w", err)
	}

	return p.compileSchema(schemaData, "task-metadata.json")
}

// loadSchemaFromFile loads the JSON schema from a file
func (p *FrontmatterParser) loadSchemaFromFile(schemaPath string) error {
	schemaData, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	return p.compileSchema(schemaData, schemaPath)
}

// compileSchema compiles the JSON schema
func (p *FrontmatterParser) compileSchema(schemaData []byte, schemaURL string) error {
	compiler := jsonschema.NewCompiler()
	compiler.DefaultDraft(jsonschema.Draft7)

	// Parse JSON schema data
	var schemaDoc any
	if err := json.Unmarshal(schemaData, &schemaDoc); err != nil {
		return fmt.Errorf("failed to parse schema JSON: %w", err)
	}

	if err := compiler.AddResource(schemaURL, schemaDoc); err != nil {
		return fmt.Errorf("failed to add schema resource: %w", err)
	}

	schema, err := compiler.Compile(schemaURL)
	if err != nil {
		return fmt.Errorf("failed to compile schema: %w", err)
	}

	p.schema = schema
	return nil
}

// ParseTaskFile parses a task file with YAML frontmatter
func (p *FrontmatterParser) ParseTaskFile(filePath string) (*ParsedTaskFile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open task file '%s': %w", filePath, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			// Log the error but don't override the main error
			fmt.Printf("Warning: failed to close file %s: %v\n", filePath, closeErr)
		}
	}()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read task file '%s': %w", filePath, err)
	}

	return p.parseLines(lines, filePath)
}

// ParseTaskContent parses task content directly from string
func (p *FrontmatterParser) ParseTaskContent(content, filePath string) (*ParsedTaskFile, error) {
	lines := strings.Split(content, "\n")
	return p.parseLines(lines, filePath)
}

// parseLines parses lines to extract frontmatter and content
func (p *FrontmatterParser) parseLines(lines []string, filePath string) (*ParsedTaskFile, error) {
	result := &ParsedTaskFile{
		FilePath: filePath,
	}

	// Check if file has frontmatter (starts with ---)
	if len(lines) == 0 || lines[0] != "---" {
		// No frontmatter, return content as-is with empty metadata
		result.Content = strings.Join(lines, "\n")
		result.Metadata = &TaskMetadata{}
		return result, nil
	}

	// Find end of frontmatter
	frontmatterEnd := -1
	for i := 1; i < len(lines); i++ {
		if lines[i] == "---" {
			frontmatterEnd = i
			break
		}
	}

	if frontmatterEnd == -1 {
		return nil, fmt.Errorf("frontmatter not properly closed with '---' in file '%s'", filePath)
	}

	// Extract frontmatter YAML
	frontmatterLines := lines[1:frontmatterEnd]
	frontmatterYAML := strings.Join(frontmatterLines, "\n")

	// Parse YAML frontmatter
	var metadata TaskMetadata
	if len(strings.TrimSpace(frontmatterYAML)) > 0 {
		if err := yaml.Unmarshal([]byte(frontmatterYAML), &metadata); err != nil {
			return nil, fmt.Errorf("failed to parse YAML frontmatter in file '%s': %w", filePath, err)
		}
	}

	// Extract content (everything after frontmatter)
	var contentLines []string
	if frontmatterEnd+1 < len(lines) {
		contentLines = lines[frontmatterEnd+1:]
	}
	result.Content = strings.Join(contentLines, "\n")
	result.Metadata = &metadata

	return result, nil
}

// ValidateMetadata validates task metadata against the JSON schema
func (p *FrontmatterParser) ValidateMetadata(metadata *TaskMetadata, filePath string) error {
	if p.schema == nil {
		return fmt.Errorf("schema not loaded")
	}

	// Convert metadata to interface{} for schema validation
	jsonData, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata for validation: %w", err)
	}

	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return fmt.Errorf("failed to unmarshal metadata for validation: %w", err)
	}

	// Validate against schema
	if err := p.schema.Validate(data); err != nil {
		return fmt.Errorf("metadata validation failed for file '%s': %w", filePath, err)
	}

	return nil
}

// ParseAndValidateTaskFile parses and validates a task file
func (p *FrontmatterParser) ParseAndValidateTaskFile(filePath string) (*ParsedTaskFile, error) {
	parsed, err := p.ParseTaskFile(filePath)
	if err != nil {
		return nil, err
	}

	// Validate metadata if present
	if parsed.Metadata != nil {
		if err := p.ValidateMetadata(parsed.Metadata, filePath); err != nil {
			return parsed, err
		}
	}

	return parsed, nil
}

// HasDependencies returns true if the metadata has any dependencies declared
func (m *TaskMetadata) HasDependencies() bool {
	if m.Dependencies == nil {
		return false
	}

	return len(m.Dependencies.Templates) > 0 ||
		len(m.Dependencies.Data) > 0 ||
		len(m.Dependencies.MCP) > 0
}

// GetAllDependencyPaths returns all dependency paths with their types
func (m *TaskMetadata) GetAllDependencyPaths() map[string][]string {
	result := make(map[string][]string)

	if m.Dependencies == nil {
		return result
	}

	if len(m.Dependencies.Templates) > 0 {
		result["templates"] = make([]string, len(m.Dependencies.Templates))
		copy(result["templates"], m.Dependencies.Templates)
	}

	if len(m.Dependencies.Data) > 0 {
		result["data"] = make([]string, len(m.Dependencies.Data))
		copy(result["data"], m.Dependencies.Data)
	}

	if len(m.Dependencies.MCP) > 0 {
		result["mcp"] = make([]string, len(m.Dependencies.MCP))
		copy(result["mcp"], m.Dependencies.MCP)
	}

	return result
}
