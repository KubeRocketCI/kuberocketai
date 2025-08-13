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
	"fmt"
)

// TokenCalculator defines the interface for token calculation engines
type TokenCalculator interface {
	// CalculateTokens calculates the number of tokens for the given text
	CalculateTokens(ctx context.Context, text string) (int, error)
}

// AssetTokenInfo represents token information for a specific asset
type AssetTokenInfo struct {
	Path    string `json:"path"`
	Type    string `json:"type"` // "agent", "task", "template", "data"
	Tokens  int    `json:"tokens"`
	Content string `json:"-"` // Content not included in JSON output
}

// AgentTokenInfo represents complete token information for an agent
type AgentTokenInfo struct {
	AgentName      string           `json:"agent_name"`
	AgentFile      string           `json:"agent_file"`
	AgentShortName string           `json:"agent_short_name"`
	TotalTokens    int              `json:"total_tokens"`
	Assets         []AssetTokenInfo `json:"assets"`
	Dependencies   struct {
		Tasks     []AssetTokenInfo `json:"tasks"`
		Templates []AssetTokenInfo `json:"templates"`
		DataFiles []AssetTokenInfo `json:"data_files"`
	} `json:"dependencies"`
}

// ProjectTokenInfo represents token information for an entire project
type ProjectTokenInfo struct {
	TotalTokens int              `json:"total_tokens"`
	Agents      []AgentTokenInfo `json:"agents"`
	Breakdown   TokenBreakdown   `json:"breakdown"`
}

// TokenBreakdown provides token counts by asset type
type TokenBreakdown struct {
	Agents    int `json:"agents"`
	Tasks     int `json:"tasks"`
	Templates int `json:"templates"`
	DataFiles int `json:"data_files"`
}

// Engine manages token calculation with pluggable tokenizer backends
type Engine struct {
	calculator TokenCalculator
}

// NewEngine creates a new token calculation engine with the specified calculator
func NewEngine(calculator TokenCalculator) *Engine {
	return &Engine{
		calculator: calculator,
	}
}

// CalculateTokens calculates tokens for the given text using the configured calculator
func (e *Engine) CalculateTokens(ctx context.Context, text string) (int, error) {
	if e.calculator == nil {
		return 0, fmt.Errorf("no token calculator configured")
	}

	return e.calculator.CalculateTokens(ctx, text)
}

// CalculateAssetTokens calculates token information for a single asset
func (e *Engine) CalculateAssetTokens(ctx context.Context, path, assetType, content string) (AssetTokenInfo, error) {
	tokens, err := e.CalculateTokens(ctx, content)
	if err != nil {
		return AssetTokenInfo{}, fmt.Errorf("failed to calculate tokens for %s: %w", path, err)
	}

	return AssetTokenInfo{
		Path:    path,
		Type:    assetType,
		Tokens:  tokens,
		Content: content,
	}, nil
}
