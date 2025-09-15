package bundle

import (
	"slices"
	"strings"
)

// GenerateBundleFilename creates the bundle filename
func GenerateBundleFilename(customOutput string, selectedAgents []string) string {
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

	agents := make([]string, len(selectedAgents))

	// Convert to lowercase for consistent filenames
	for i, agent := range selectedAgents {
		agents[i] = strings.ToLower(agent)
	}

	// Sort alphabetically
	slices.Sort(agents)

	return strings.Join(agents, "-") + ".md"
}
