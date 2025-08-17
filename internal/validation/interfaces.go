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
	"io/fs"
)

// FrameworkScanner abstracts analyzer behavior for commands
type FrameworkScanner interface {
	OptimizedAnalyzeFramework() ([]ValidationIssue, *FrameworkInsights, error)
}

// EmbeddedAssetAnalyzer provides dependency analysis for embedded assets before installation
type EmbeddedAssetAnalyzer interface {
	// AnalyzeEmbeddedDependencies analyzes dependencies from embedded filesystem
	AnalyzeEmbeddedDependencies(ctx context.Context, embeddedAssets fs.FS, agentNames []string) (*FrameworkInsights, error)

	// ValidateEmbeddedAgentDependencies validates dependencies exist in embedded assets
	ValidateEmbeddedAgentDependencies(ctx context.Context, embeddedAssets fs.FS, agentNames []string) ([]ValidationIssue, error)
}
