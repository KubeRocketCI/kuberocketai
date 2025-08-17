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
package assets

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// AssetSource abstracts different sources of framework assets (filesystem vs embedded)
type AssetSource interface {
	// ListAgentFiles returns a list of agent YAML file paths
	ListAgentFiles() ([]string, error)

	// ReadFile reads a file from the source
	ReadFile(path string) ([]byte, error)

	// Exists checks if a file exists in the source
	Exists(path string) bool

	// GetSourceType returns the type of source for logging/debugging
	GetSourceType() string
}

// EmbeddedAssetSource extends AssetSource with embedded-specific functionality
type EmbeddedAssetSource interface {
	AssetSource
	// GetEmbeddedFS returns the underlying embedded filesystem for validation operations
	GetEmbeddedFS() embed.FS
}

// FilesystemSource implements AssetSource for installed framework files
type FilesystemSource struct {
	baseDir string
}

// NewFilesystemSource creates a new filesystem-based asset source
func NewFilesystemSource(baseDir string) *FilesystemSource {
	return &FilesystemSource{baseDir: baseDir}
}

func (fs *FilesystemSource) ListAgentFiles() ([]string, error) {
	agentsPath := filepath.Join(fs.baseDir, KrciAIDir, agentsDir)
	agentFiles, err := filepath.Glob(filepath.Join(agentsPath, "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to scan agents directory: %w", err)
	}
	return agentFiles, nil
}

func (fs *FilesystemSource) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (fs *FilesystemSource) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (fs *FilesystemSource) GetSourceType() string {
	return FilesystemSourceType
}

// EmbeddedSource implements AssetSource for embedded framework assets
type EmbeddedSource struct {
	embeddedAssets embed.FS
}

// NewEmbeddedSource creates a new embedded asset source
func NewEmbeddedSource(embeddedAssets embed.FS) *EmbeddedSource {
	return &EmbeddedSource{embeddedAssets: embeddedAssets}
}

func (es *EmbeddedSource) ListAgentFiles() ([]string, error) {
	agentFiles, err := fs.Glob(es.embeddedAssets, embeddedPrefix+"/"+agentsDir+"/*.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to scan embedded agents: %w", err)
	}
	return agentFiles, nil
}

func (es *EmbeddedSource) ReadFile(path string) ([]byte, error) {
	return fs.ReadFile(es.embeddedAssets, path)
}

func (es *EmbeddedSource) Exists(path string) bool {
	_, err := es.embeddedAssets.Open(path)
	return err == nil
}

func (es *EmbeddedSource) GetSourceType() string {
	return EmbeddedSourceType
}

// GetEmbeddedFS implements EmbeddedAssetSource interface
func (es *EmbeddedSource) GetEmbeddedFS() embed.FS {
	return es.embeddedAssets
}

// parseAgentFromSource parses an agent YAML file from any asset source
func parseAgentFromSource(source AssetSource, filePath string) (*AgentInfo, error) {
	data, err := source.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var agent Agent
	if err := yaml.Unmarshal(data, &agent); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Extract agent information
	identity := agent.Agent.Identity
	agentInfo := &AgentInfo{
		Name:        identity.Name,
		Description: identity.Description,
		Role:        identity.Role,
		Goal:        identity.Goal,
		Icon:        identity.Icon,
		FilePath:    filePath,
		ShortName:   strings.TrimSuffix(filepath.Base(filePath), ".yaml"),
	}

	// Validate required fields
	if agentInfo.Name == "" {
		return nil, fmt.Errorf("agent name is required")
	}
	if agentInfo.Role == "" {
		return nil, fmt.Errorf("agent role is required")
	}

	return agentInfo, nil
}
