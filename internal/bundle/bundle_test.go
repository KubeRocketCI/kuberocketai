package bundle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBundleFilenameWithTask(t *testing.T) {
	tests := []struct {
		name           string
		customOutput   string
		selectedAgents []string
		taskName       string
		expected       string
	}{
		{
			name:           "single agent with task",
			customOutput:   "",
			selectedAgents: []string{"pm"},
			taskName:       "create-prd",
			expected:       "pm-create-prd.md",
		},
		{
			name:           "single agent with task - mixed case",
			customOutput:   "",
			selectedAgents: []string{"PM"},
			taskName:       "Create-PRD",
			expected:       "pm-create-prd.md",
		},
		{
			name:           "multiple agents with task - should ignore task",
			customOutput:   "",
			selectedAgents: []string{"pm", "architect"},
			taskName:       "create-prd",
			expected:       "architect-pm.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateBundleFilename(tt.customOutput, tt.selectedAgents, tt.taskName)
			assert.Equal(t, tt.expected, result)
		})
	}
}
