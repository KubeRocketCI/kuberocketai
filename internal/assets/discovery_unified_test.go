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
	"os"
	"strings"
	"testing"
)

// Use the existing embedded assets from the main command for testing

// TestCurrentBehaviorCapture captures current behavior before refactoring
// This ensures our refactoring doesn't break existing functionality
func TestCurrentBehaviorCapture(t *testing.T) {
	t.Run("FilesystemDiscovery_AllAgents", func(t *testing.T) {
		tmpDir := setupTestFramework(t)
		defer os.RemoveAll(tmpDir)

		discovery := NewDiscovery(tmpDir, testAssets)
		agentDeps, err := discovery.DiscoverAgentsWithDependencies()

		// Capture current behavior for comparison after refactoring
		validateFilesystemDiscoveryResult(t, agentDeps, err, "all_agents")
	})

	// TODO: Add embedded discovery tests once we have proper test embedded assets
	t.Run("EmbeddedDiscovery_Placeholder", func(t *testing.T) {
		t.Skip("Embedded discovery tests require proper test setup - will add during implementation")
	})
}

// TestFilesystemVsEmbeddedBehaviorDifferences documents current differences
func TestFilesystemVsEmbeddedBehaviorDifferences(t *testing.T) {
	t.Run("InstallationCheck", func(t *testing.T) {
		// Filesystem requires installation check
		discovery := NewDiscovery("/nonexistent", testAssets)
		_, err := discovery.DiscoverAgentsWithDependencies()
		if err == nil {
			t.Error("Expected error for non-installed framework in filesystem mode")
		}
		if !strings.Contains(err.Error(), "framework not installed") {
			t.Errorf("Expected installation error, got: %v", err)
		}

		// TODO: Test embedded behavior once proper test setup is available
		t.Skip("Embedded behavior testing requires proper test setup")
	})

	t.Run("AgentFiltering", func(t *testing.T) {
		// Document that filesystem returns ALL agents
		tmpDir := setupTestFramework(t)
		defer os.RemoveAll(tmpDir)

		discovery := NewDiscovery(tmpDir, testAssets)
		allAgents, err := discovery.DiscoverAgentsWithDependencies()
		if err != nil {
			t.Fatalf("Filesystem discovery failed: %v", err)
		}

		// Log the behavior - embedded filtering will be tested after refactoring
		t.Logf("Filesystem discovery returned %d agents (all)", len(allAgents))
		t.Logf("Embedded filtering behavior will be validated after unified implementation")
	})
}

// TestInterfaceCompatibility ensures interface implementations work correctly
func TestInterfaceCompatibility(t *testing.T) {
	tmpDir := setupTestFramework(t)
	defer os.RemoveAll(tmpDir)

	discovery := NewDiscovery(tmpDir, testAssets)

	// Test that discovery implements the tokens calculator interface
	var _ DiscoveryInterface = discovery

	// Test DiscoverAgentWithDependencies uses DiscoverAgentsWithDependencies
	agentDep, err := discovery.DiscoverAgentWithDependencies("test-agent")
	if err == nil {
		// Verify it returns single agent data
		if agentDep.Name == "" {
			t.Error("Expected agent data, got empty agent")
		}
	}
}

// TestEdgeCases captures edge case behavior
func TestEdgeCases(t *testing.T) {
	t.Run("FilesystemBasicEdgeCases", func(t *testing.T) {
		tmpDir := setupTestFramework(t)
		defer os.RemoveAll(tmpDir)

		discovery := NewDiscovery(tmpDir, testAssets)

		// Test that filesystem discovery works consistently
		agentDeps, err := discovery.DiscoverAgentsWithDependencies()
		if err != nil {
			t.Errorf("Filesystem discovery should work: %v", err)
		}

		t.Logf("Filesystem discovery returned %d agents", len(agentDeps))
	})

	// TODO: Add embedded edge cases after unified implementation
	t.Run("EmbeddedEdgeCases_Placeholder", func(t *testing.T) {
		t.Skip("Embedded edge cases will be tested after unified implementation")
	})
}

// Helper functions for validation

func validateFilesystemDiscoveryResult(t *testing.T, agentDeps []AgentDependencyInfo, err error, testCase string) {
	t.Helper()

	if err != nil {
		t.Errorf("[%s] Filesystem discovery failed: %v", testCase, err)
		return
	}

	if len(agentDeps) == 0 {
		t.Errorf("[%s] Expected at least one agent", testCase)
		return
	}

	// Validate structure
	for i, agent := range agentDeps {
		if agent.Name == "" {
			t.Errorf("[%s] Agent %d missing name", testCase, i)
		}
		if agent.ShortName == "" {
			t.Errorf("[%s] Agent %d missing short name", testCase, i)
		}
		if agent.FilePath == "" {
			t.Errorf("[%s] Agent %d missing file path", testCase, i)
		}

		// Tasks, Templates, DataFiles can be empty (initialized as empty slices)
		if agent.Tasks == nil {
			t.Errorf("[%s] Agent %d has nil Tasks slice", testCase, i)
		}
		if agent.Templates == nil {
			t.Errorf("[%s] Agent %d has nil Templates slice", testCase, i)
		}
		if agent.DataFiles == nil {
			t.Errorf("[%s] Agent %d has nil DataFiles slice", testCase, i)
		}
	}

	t.Logf("[%s] Filesystem discovery returned %d agents", testCase, len(agentDeps))
}

// TODO: validateEmbeddedDiscoveryResult will be added when implementing embedded tests

// DiscoveryInterface defines the expected interface (from tokens package)
type DiscoveryInterface interface {
	DiscoverAgentWithDependencies(shortName string) (AgentDependencyInfo, error)
	DiscoverAgentsWithDependencies(agentNames ...string) ([]AgentDependencyInfo, error)
	GetAgentByShortName(shortName string) (*AgentInfo, error)
}
