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
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// AgentInfo represents basic information about an agent
type AgentInfo struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Role        string `yaml:"role"`
	Goal        string `yaml:"goal"`
	FilePath    string `yaml:"-"` // Not from YAML, computed
}

// AgentIdentity represents the identity section of an agent YAML
type AgentIdentity struct {
	Name        string `yaml:"name"`
	ID          string `yaml:"id"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	Role        string `yaml:"role"`
	Goal        string `yaml:"goal"`
}

// Agent represents the structure of an agent YAML file
type Agent struct {
	Agent struct {
		Identity AgentIdentity `yaml:"identity"`
	} `yaml:"agent"`
}

// Discovery handles discovery and parsing of framework assets
type Discovery struct {
	installer *Installer
}

// NewDiscovery creates a new asset discovery service
func NewDiscovery(targetDir string, embeddedAssets embed.FS) *Discovery {
	return &Discovery{
		installer: NewInstaller(targetDir, embeddedAssets),
	}
}

// DiscoverAgents discovers and returns information about all installed agents
func (d *Discovery) DiscoverAgents() ([]AgentInfo, error) {
	if !d.installer.IsInstalled() {
		return nil, fmt.Errorf("framework not installed in current directory - run 'krci-ai install'")
	}

	agentsPath := d.installer.GetAgentsPath()
	agentFiles, err := filepath.Glob(filepath.Join(agentsPath, "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to scan agents directory: %w", err)
	}

	if len(agentFiles) == 0 {
		return nil, fmt.Errorf("no agent files found in %s", agentsPath)
	}

	var agents []AgentInfo
	for _, file := range agentFiles {
		agentInfo, err := d.parseAgentFile(file)
		if err != nil {
			// Log warning but continue with other agents
			fmt.Fprintf(os.Stderr, "Warning: failed to parse agent file %s: %v\n", file, err)
			continue
		}
		agentInfo.FilePath = file
		agents = append(agents, *agentInfo)
	}

	return agents, nil
}

// parseAgentFile parses an agent YAML file and extracts basic information
func (d *Discovery) parseAgentFile(filePath string) (*AgentInfo, error) {
	data, err := os.ReadFile(filePath)
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

// GetAgentByName returns information about a specific agent by name
func (d *Discovery) GetAgentByName(name string) (*AgentInfo, error) {
	agents, err := d.DiscoverAgents()
	if err != nil {
		return nil, err
	}

	for _, agent := range agents {
		if agent.Name == name {
			return &agent, nil
		}
	}

	return nil, fmt.Errorf("agent '%s' not found", name)
}

// ListAvailableAgents returns a simple list of agent names
func (d *Discovery) ListAvailableAgents() ([]string, error) {
	agents, err := d.DiscoverAgents()
	if err != nil {
		return nil, err
	}

	var names []string
	for _, agent := range agents {
		names = append(names, agent.Name)
	}

	return names, nil
}

// FormatAgentSummary returns a formatted string summarizing agent information
func (d *Discovery) FormatAgentSummary(agent AgentInfo) string {
	fileName := strings.TrimSuffix(filepath.Base(agent.FilePath), ".yaml")
	return fmt.Sprintf("%-15s | %-25s | %s", fileName, agent.Role, agent.Description)
}

// ValidateAgentStructure performs basic validation of agent file structure
func (d *Discovery) ValidateAgentStructure(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var agent Agent
	if err := yaml.Unmarshal(data, &agent); err != nil {
		return fmt.Errorf("invalid YAML structure: %w", err)
	}

	identity := agent.Agent.Identity

	// Check required fields
	if identity.Name == "" {
		return fmt.Errorf("agent.identity.name is required")
	}
	if identity.ID == "" {
		return fmt.Errorf("agent.identity.id is required")
	}
	if identity.Version == "" {
		return fmt.Errorf("agent.identity.version is required")
	}
	if identity.Description == "" {
		return fmt.Errorf("agent.identity.description is required")
	}
	if identity.Role == "" {
		return fmt.Errorf("agent.identity.role is required")
	}
	if identity.Goal == "" {
		return fmt.Errorf("agent.identity.goal is required")
	}

	return nil
}
