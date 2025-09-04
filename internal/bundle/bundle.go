package bundle

import (
	"fmt"
	"slices"
	"strings"
)

// GenerateBundleFilename creates the bundle filename
func GenerateBundleFilename(customOutput string, selectedAgents []string, taskName string) string {
	if customOutput != "" {
		// Ensure .md extension
		if !strings.HasSuffix(customOutput, ".md") {
			customOutput += ".md"
		}
		return customOutput
	}

	if len(selectedAgents) == 0 {
		return "all.md"
	}

	// If both agent and task are specified, use agent-task pattern
	if taskName != "" && len(selectedAgents) == 1 {
		agentName := strings.ToLower(selectedAgents[0])
		taskNameLower := strings.ToLower(taskName)
		return fmt.Sprintf("%s-%s.md", agentName, taskNameLower)
	}

	// For multiple agents or agent-only bundles, use existing logic
	sortedAgents := slices.Clone(selectedAgents)

	// Convert to lowercase for consistent filenames
	for i, agent := range sortedAgents {
		sortedAgents[i] = strings.ToLower(agent)
	}

	// Sort alphabetically
	slices.Sort(sortedAgents)

	return strings.Join(sortedAgents, "-") + ".md"
}
