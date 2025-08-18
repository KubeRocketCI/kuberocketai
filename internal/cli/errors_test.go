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
package cli

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestEnhancedErrorHandler tests the basic ErrorHandler functionality
func TestEnhancedErrorHandler(t *testing.T) {
	handler := NewErrorHandler()
	require.NotNil(t, handler, "NewErrorHandler should not return nil")
}

// TestPrintTaskValidationError tests enhanced task validation error messaging
func TestPrintTaskValidationError(t *testing.T) {
	// Capture stderr output
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	handler := NewErrorHandler()

	tests := []struct {
		name           string
		agentName      string
		invalidTasks   []string
		availableTasks []string
		expectInOutput []string
	}{
		{
			name:           "single invalid task",
			agentName:      "pm",
			invalidTasks:   []string{"invalid-task"},
			availableTasks: []string{"create-prd", "update-prd"},
			expectInOutput: []string{"invalid-task", "pm", "create-prd", "update-prd", "Available tasks"},
		},
		{
			name:           "multiple invalid tasks",
			agentName:      "dev",
			invalidTasks:   []string{"task1", "task2"},
			availableTasks: []string{"implement-feature"},
			expectInOutput: []string{"task1", "task2", "dev", "implement-feature", "Available tasks"},
		},
		{
			name:           "no available tasks",
			agentName:      "qa",
			invalidTasks:   []string{"invalid-task"},
			availableTasks: []string{},
			expectInOutput: []string{"invalid-task", "qa"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear the buffer
			var buf bytes.Buffer

			handler.PrintTaskValidationError(tt.agentName, tt.invalidTasks, tt.availableTasks)

			// Close write end and read from pipe
			w.Close()
			buf.ReadFrom(r)
			output := buf.String()

			// Check that expected strings are in output
			for _, expected := range tt.expectInOutput {
				assert.Contains(t, output, expected, "Expected string should be in output")
			}

			// Reset pipe for next test
			r, w, _ = os.Pipe()
			os.Stderr = w
		})
	}

	// Restore stderr
	w.Close()
	os.Stderr = old
}

// TestPrintAgentValidationError tests enhanced agent validation error messaging
func TestPrintAgentValidationError(t *testing.T) {
	// Capture stderr output
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	handler := NewErrorHandler()

	tests := []struct {
		name            string
		invalidAgents   []string
		availableAgents []string
		expectInOutput  []string
	}{
		{
			name:            "single invalid agent",
			invalidAgents:   []string{"invalid-agent"},
			availableAgents: []string{"pm", "dev", "qa"},
			expectInOutput:  []string{"invalid-agent", "Available agents", "pm", "dev", "qa"},
		},
		{
			name:            "multiple invalid agents",
			invalidAgents:   []string{"agent1", "agent2"},
			availableAgents: []string{"pm", "dev"},
			expectInOutput:  []string{"agent1", "agent2", "Available agents", "pm", "dev"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			handler.PrintAgentValidationError(tt.invalidAgents, tt.availableAgents)

			// Close write end and read from pipe
			w.Close()
			buf.ReadFrom(r)
			output := buf.String()

			// Check that expected strings are in output
			for _, expected := range tt.expectInOutput {
				assert.Contains(t, output, expected, "Expected string should be in output")
			}

			// Reset pipe for next test
			r, w, _ = os.Pipe()
			os.Stderr = w
		})
	}

	// Restore stderr
	w.Close()
	os.Stderr = old
}

// TestPrintInstallationError tests enhanced installation error messaging
func TestPrintInstallationError(t *testing.T) {
	// Capture stderr output
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	handler := NewErrorHandler()

	tests := []struct {
		name           string
		err            error
		phase          string
		expectInOutput []string
	}{
		{
			name:           "permission error",
			err:            os.ErrPermission,
			phase:          "agent installation",
			expectInOutput: []string{"agent installation", "permission denied", "Recovery suggestions", "write permissions"},
		},
		{
			name:           "generic error",
			err:            os.ErrNotExist,
			phase:          "framework setup",
			expectInOutput: []string{"framework setup", "file does not exist", "Recovery suggestions"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			handler.PrintInstallationError(tt.err, tt.phase)

			// Close write end and read from pipe
			w.Close()
			buf.ReadFrom(r)
			output := buf.String()

			// Check that expected strings are in output
			for _, expected := range tt.expectInOutput {
				assert.Contains(t, output, expected, "Expected string should be in output")
			}

			// Reset pipe for next test
			r, w, _ = os.Pipe()
			os.Stderr = w
		})
	}

	// Restore stderr
	w.Close()
	os.Stderr = old
}

// TestPrintErrorWithGuidance tests error messages with recovery guidance
func TestPrintErrorWithGuidance(t *testing.T) {
	// Capture stderr output
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	handler := NewErrorHandler()

	tests := []struct {
		name           string
		message        string
		guidance       []string
		expectInOutput []string
	}{
		{
			name:           "error with guidance",
			message:        "Installation failed",
			guidance:       []string{"Check permissions", "Verify disk space"},
			expectInOutput: []string{"Installation failed", "Recovery guidance", "Check permissions", "Verify disk space"},
		},
		{
			name:           "error without guidance",
			message:        "Simple error",
			guidance:       []string{},
			expectInOutput: []string{"Simple error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			handler.PrintErrorWithGuidance(tt.message, tt.guidance)

			// Close write end and read from pipe
			w.Close()
			buf.ReadFrom(r)
			output := buf.String()

			// Check that expected strings are in output
			for _, expected := range tt.expectInOutput {
				assert.Contains(t, output, expected, "Expected string should be in output")
			}

			// Reset pipe for next test
			r, w, _ = os.Pipe()
			os.Stderr = w
		})
	}

	// Restore stderr
	w.Close()
	os.Stderr = old
}

// TestExistingErrorMethods tests that existing error methods still work
func TestExistingErrorMethods(t *testing.T) {
	handler := NewErrorHandler()

	// Test that these methods don't panic
	defer func() {
		if r := recover(); r != nil {
			assert.Fail(t, "Existing error methods should not panic", "Existing error method panicked: %v", r)
		}
	}()

	handler.PrintError("Test error message")
	handler.PrintWarning("Test warning message")
}

// BenchmarkPrintTaskValidationError benchmarks the task validation error printing
func BenchmarkPrintTaskValidationError(b *testing.B) {
	handler := NewErrorHandler()
	invalidTasks := []string{"task1", "task2", "task3"}
	availableTasks := []string{"create-prd", "update-prd", "delete-prd", "analyze-requirements"}

	// Capture stderr to avoid output during benchmark
	old := os.Stderr
	os.Stderr, _ = os.Open(os.DevNull)
	defer func() { os.Stderr = old }()

	for i := 0; i < b.N; i++ {
		handler.PrintTaskValidationError("pm", invalidTasks, availableTasks)
	}
}

// BenchmarkPrintInstallationError benchmarks the installation error printing
func BenchmarkPrintInstallationError(b *testing.B) {
	handler := NewErrorHandler()
	testErr := os.ErrPermission

	// Capture stderr to avoid output during benchmark
	old := os.Stderr
	os.Stderr, _ = os.Open(os.DevNull)
	defer func() { os.Stderr = old }()

	for i := 0; i < b.N; i++ {
		handler.PrintInstallationError(testErr, "test phase")
	}
}
