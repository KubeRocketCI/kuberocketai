package cli

import (
	"bytes"
	"testing"
)

func TestNewOutputHandler(t *testing.T) {
	handler := NewOutputHandler()
	if handler == nil {
		t.Fatal("NewOutputHandler returned nil")
	}
}

func TestOutputHandlerMessages(t *testing.T) {
	// Capture output
	var buf bytes.Buffer

	// Create handler and redirect output for testing
	handler := NewOutputHandler()

	// Since the actual implementation uses color output to stdout,
	// we'll test that the methods don't panic and can be called
	tests := []struct {
		name string
		fn   func()
	}{
		{
			name: "PrintProgress",
			fn:   func() { handler.PrintProgress("Test progress message") },
		},
		{
			name: "PrintInfo",
			fn:   func() { handler.PrintInfo("Test info message") },
		},
		{
			name: "PrintSuccess",
			fn:   func() { handler.PrintSuccess("Test success message") },
		},
		{
			name: "PrintWarning",
			fn:   func() { handler.PrintWarning("Test warning message") },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test that the function doesn't panic
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Function %s panicked: %v", tt.name, r)
				}
			}()

			tt.fn()
		})
	}

	// Test output (this is implementation-specific and might need adjustment)
	_ = buf.String() // Consume any captured output
}

func TestErrorHandler(t *testing.T) {
	// Test that error handler can be created
	handler := NewErrorHandler()
	if handler == nil {
		t.Fatal("NewErrorHandler returned nil")
	}
}

func TestErrorHandlerMethods(t *testing.T) {
	handler := NewErrorHandler()

	tests := []struct {
		name string
		fn   func()
	}{
		{
			name: "HandleError",
			fn:   func() { handler.HandleError(nil, "Test context") },
		},
		{
			name: "PrintError",
			fn:   func() { handler.PrintError("Test error message") },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test that the function doesn't panic
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Function %s panicked: %v", tt.name, r)
				}
			}()

			tt.fn()
		})
	}
}
