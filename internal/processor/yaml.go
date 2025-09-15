package processor

import (
	"bytes"
	"fmt"
	"os"

	"github.com/stretchr/testify/assert/yaml"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

// FileReader interface for abstracting file reading operations
type FileReader interface {
	ReadFile(name string) ([]byte, error)
}

func UnmarshalAgentFile(filePath string) (*AgentYamlRepresentation, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var rawAgent AgentYamlRepresentation
	if err := yaml.Unmarshal(data, &rawAgent); err != nil {
		return nil, fmt.Errorf("failed to unmarshal agent file: %w", err)
	}

	return &rawAgent, nil
}

// UnmarshalAgentFileFromFS unmarshals an agent file using a FileReader interface
func UnmarshalAgentFileFromFS(fs FileReader, filePath string) (*AgentYamlRepresentation, error) {
	data, err := fs.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var rawAgent AgentYamlRepresentation
	if err := yaml.Unmarshal(data, &rawAgent); err != nil {
		return nil, fmt.Errorf("failed to unmarshal agent file: %w", err)
	}

	return &rawAgent, nil
}

func UnmarshalTaskDependenciesFile(filePath string) (*TaskDependenciesYamlRepresentation, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q: %w", filePath, err)
	}

	md := goldmark.New(
		goldmark.WithExtensions(&frontmatter.Extender{
			Mode: frontmatter.SetMetadata,
		}),
	)
	ctx := parser.NewContext()
	var buf bytes.Buffer
	if err := md.Convert(data, &buf, parser.WithContext(ctx)); err != nil {
		return nil, fmt.Errorf("failed to parse markdown file %q: %w", filePath, err)
	}

	meta := frontmatter.Get(ctx)
	if meta == nil {
		return nil, fmt.Errorf("no frontmatter found in file %q - YAML frontmatter is required for task dependencies", filePath)
	}

	var gotData TaskDependenciesYamlRepresentation
	if err := meta.Decode(&gotData); err != nil {
		return nil, fmt.Errorf("failed to decode YAML frontmatter in file %q: %w", filePath, err)
	}

	return &gotData, nil
}

// UnmarshalTaskDependenciesFileFromFS unmarshals task dependencies from a file using a FileReader interface
func UnmarshalTaskDependenciesFileFromFS(fs FileReader, filePath string) (*TaskDependenciesYamlRepresentation, error) {
	data, err := fs.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q: %w", filePath, err)
	}

	md := goldmark.New(
		goldmark.WithExtensions(&frontmatter.Extender{
			Mode: frontmatter.SetMetadata,
		}),
	)
	ctx := parser.NewContext()
	var buf bytes.Buffer
	if err := md.Convert(data, &buf, parser.WithContext(ctx)); err != nil {
		return nil, fmt.Errorf("failed to parse markdown file %q: %w", filePath, err)
	}

	var gotData TaskDependenciesYamlRepresentation

	meta := frontmatter.Get(ctx)
	if meta == nil {
		return &gotData, nil
	}

	if err := meta.Decode(&gotData); err != nil {
		return nil, fmt.Errorf("failed to decode YAML frontmatter in file %q: %w", filePath, err)
	}

	return &gotData, nil
}
