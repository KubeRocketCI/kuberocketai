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
package tokens

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/bundle"
)

// DiscoveryInterface defines the interface for asset discovery operations
type DiscoveryInterface interface {
	DiscoverAgentWithDependencies(shortName string) (assets.AgentDependencyInfo, error)
	DiscoverAgentsWithDependencies(agentNames ...string) ([]assets.AgentDependencyInfo, error)
	GetAgentByShortName(shortName string) (*assets.AgentInfo, error)
}

// Calculator provides high-level token calculation functionality
type Calculator struct {
	engine    *Engine
	discovery DiscoveryInterface
	targetDir string
}

// NewCalculator creates a new token calculator with GPT-4 tokenization
// This is the original constructor for backward compatibility
func NewCalculator(targetDir string, embeddedAssets embed.FS) (*Calculator, error) {
	engine, err := NewDefaultEngine()
	if err != nil {
		return nil, fmt.Errorf("failed to create GPT-4 calculator: %w", err)
	}

	discovery := assets.NewDiscovery(targetDir, embeddedAssets)

	return NewCalculatorWithDependencies(engine, discovery, targetDir), nil
}

// NewCalculatorWithDependencies creates a new calculator with injected dependencies
// This is the preferred constructor for testing and when you have custom dependencies
func NewCalculatorWithDependencies(engine *Engine, discovery DiscoveryInterface, targetDir string) *Calculator {
	return &Calculator{
		engine:    engine,
		discovery: discovery,
		targetDir: targetDir,
	}
}

// CalculateAgentTokens calculates token information for a specific agent
func (c *Calculator) CalculateAgentTokens(ctx context.Context, agentShortName string) (*AgentTokenInfo, error) {
	// Get agent information with dependencies
	agentDeps, err := c.discovery.DiscoverAgentWithDependencies(agentShortName)
	if err != nil {
		return nil, fmt.Errorf("failed to discover agents: %w", err)
	}

	return c.calculateSingleAgentTokens(ctx, agentDeps)
}

// CalculateAllTokens calculates token information for the entire project
func (c *Calculator) CalculateAllTokens(ctx context.Context) (*ProjectTokenInfo, error) {
	// Get all agents with dependencies
	agentDeps, err := c.discovery.DiscoverAgentsWithDependencies()
	if err != nil {
		return nil, fmt.Errorf("failed to discover agents: %w", err)
	}

	projectInfo := &ProjectTokenInfo{
		Agents:    make([]AgentTokenInfo, 0, len(agentDeps)),
		Breakdown: TokenBreakdown{},
	}

	// Use goroutines for parallel processing
	g, ctx := errgroup.WithContext(ctx)
	results := make(chan AgentTokenInfo, len(agentDeps))

	// Process each agent concurrently
	for _, agent := range agentDeps {
		g.Go(func() error {
			agentInfo, err := c.calculateSingleAgentTokens(ctx, agent)
			if err != nil {
				return fmt.Errorf("failed to calculate tokens for agent %s: %w", agent.Name, err)
			}

			select {
			case results <- *agentInfo:
			case <-ctx.Done():
				return ctx.Err()
			}

			return nil
		})
	}

	go func() {
		_ = g.Wait() // Error is handled below in the main Wait() call
		close(results)
	}()

	// Collect results
	for result := range results {
		projectInfo.Agents = append(projectInfo.Agents, result)
		projectInfo.TotalTokens += result.TotalTokens
		c.updateProjectBreakdown(&projectInfo.Breakdown, result)
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return projectInfo, nil
}

// updateProjectBreakdown updates the project token breakdown with results from an agent
func (c *Calculator) updateProjectBreakdown(breakdown *TokenBreakdown, result AgentTokenInfo) {
	// Update breakdown from agent assets
	for _, asset := range result.Assets {
		switch asset.Type {
		case "agent":
			breakdown.Agents += asset.Tokens
		case "task":
			breakdown.Tasks += asset.Tokens
		case "template":
			breakdown.Templates += asset.Tokens
		case "data":
			breakdown.DataFiles += asset.Tokens
		}
	}

	// Add dependency tokens to breakdown
	for _, task := range result.Dependencies.Tasks {
		breakdown.Tasks += task.Tokens
	}
	for _, template := range result.Dependencies.Templates {
		breakdown.Templates += template.Tokens
	}
	for _, dataFile := range result.Dependencies.DataFiles {
		breakdown.DataFiles += dataFile.Tokens
	}
}

// calculateSingleAgentTokens calculates tokens for a single agent and its dependencies
func (c *Calculator) calculateSingleAgentTokens(ctx context.Context, agent assets.AgentDependencyInfo) (*AgentTokenInfo, error) {
	agentInfo := &AgentTokenInfo{
		AgentName:      agent.Name,
		AgentShortName: agent.ShortName,
		AgentFile:      agent.FilePath,
		Assets:         make([]AssetTokenInfo, 0),
	}

	// Calculate tokens for the agent file itself
	agentContent, err := os.ReadFile(agent.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read agent file %s: %w", agent.FilePath, err)
	}

	agentAsset, err := c.engine.CalculateAssetTokens(ctx, agent.FilePath, "agent", string(agentContent))
	if err != nil {
		return nil, err
	}

	agentInfo.Assets = append(agentInfo.Assets, agentAsset)
	agentInfo.TotalTokens += agentAsset.Tokens

	// Calculate tokens for task dependencies
	agentInfo.Dependencies.Tasks, err = c.calculateAssetTokens(ctx, agent.Tasks, assets.TasksDir)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate task tokens: %w", err)
	}

	for _, task := range agentInfo.Dependencies.Tasks {
		agentInfo.TotalTokens += task.Tokens
	}

	// Calculate tokens for template dependencies
	agentInfo.Dependencies.Templates, err = c.calculateAssetTokens(ctx, agent.Templates, assets.TemplatesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate template tokens: %w", err)
	}

	for _, template := range agentInfo.Dependencies.Templates {
		agentInfo.TotalTokens += template.Tokens
	}

	// Calculate tokens for data file dependencies
	agentInfo.Dependencies.DataFiles, err = c.calculateAssetTokens(ctx, agent.DataFiles, assets.DataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate data file tokens: %w", err)
	}

	for _, dataFile := range agentInfo.Dependencies.DataFiles {
		agentInfo.TotalTokens += dataFile.Tokens
	}

	return agentInfo, nil
}

// calculateAssetTokens calculates tokens for a list of asset files
func (c *Calculator) calculateAssetTokens(ctx context.Context, files []string, assetType string) ([]AssetTokenInfo, error) {
	if len(files) == 0 {
		return []AssetTokenInfo{}, nil
	}

	var mu sync.Mutex
	g, ctx := errgroup.WithContext(ctx)

	results := make([]AssetTokenInfo, 0, len(files))

	err := filepath.Walk(filepath.Join(c.targetDir, assets.KrciAIDir, assetType), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk %s: %w", path, err)
		}

		if info.IsDir() {
			return nil
		}

		// Check if this file matches any of the expected files (handle subdirectories)
		relPath, err := filepath.Rel(filepath.Join(c.targetDir, assets.KrciAIDir, assetType), path)
		if err != nil {
			return nil
		}

		if !slices.Contains(files, relPath) {
			return nil
		}

		g.Go(func() error {
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read %s file %s: %w", assetType, path, err)
			}

			assetInfo, err := c.engine.CalculateAssetTokens(ctx, path, assetType, string(content))
			if err != nil {
				return err
			}

			mu.Lock()
			results = append(results, assetInfo)
			mu.Unlock()

			if ctx.Err() != nil {
				return ctx.Err()
			}

			return nil
		})

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}

	// Wait for all goroutines to complete
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

// ValidateAgentExists checks if an agent with the given name exists
func (c *Calculator) ValidateAgentExists(agentName string) error {
	_, err := c.discovery.GetAgentByShortName(agentName)
	if err != nil {
		return fmt.Errorf("agent validation failed: %w", err)
	}

	return nil
}

// CalculateBundleTokens calculates tokens for an actual bundle file
func (c *Calculator) CalculateBundleTokens(ctx context.Context, agents []string) (*BundleTokenInfo, error) {
	bundleFilename := bundle.GenerateBundleFilename("", agents, "")
	bundlePath := filepath.Join(c.targetDir, assets.KrciAIDir, assets.BundleDir, bundleFilename)

	// Read bundle file content
	bundleContent, err := os.ReadFile(bundlePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read bundle file %s: %w", bundlePath, err)
	}

	// Calculate tokens for the bundle content
	bundleTokens, err := c.engine.CalculateTokens(ctx, string(bundleContent))
	if err != nil {
		return nil, fmt.Errorf("failed to calculate bundle tokens: %w", err)
	}

	// Create bundle info
	bundleInfo := &BundleTokenInfo{
		BundleScope: strings.Join(agents, ","),
		TotalTokens: bundleTokens,
		BundleFile:  bundleFilename,
	}

	return bundleInfo, nil
}
