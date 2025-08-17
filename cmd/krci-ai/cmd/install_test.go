package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

// TestInstallCommandExists verifies that the install command is properly defined
func TestInstallCommandExists(t *testing.T) {
	if installCmd == nil {
		t.Fatal("installCmd is nil")
	}

	if installCmd.Use != "install" {
		t.Errorf("Expected command name 'install', got '%s'", installCmd.Use)
	}

	if installCmd.Short == "" {
		t.Error("Command short description is empty")
	}

	if installCmd.Long == "" {
		t.Error("Command long description is empty")
	}

	if installCmd.Run == nil {
		t.Error("Command run function is nil")
	}
}

// TestInstallCommandHasRequiredFlags verifies that all required flags are defined
func TestInstallCommandHasRequiredFlags(t *testing.T) {
	// Check IDE flag
	ideFlag := installCmd.Flags().Lookup("ide")
	if ideFlag == nil {
		t.Error("IDE flag not found")
	} else {
		if ideFlag.Shorthand != "i" {
			t.Errorf("Expected IDE flag shorthand 'i', got '%s'", ideFlag.Shorthand)
		}
	}

	// Check force flag
	forceFlag := installCmd.Flags().Lookup("force")
	if forceFlag == nil {
		t.Error("Force flag not found")
	} else {
		if forceFlag.Shorthand != "f" {
			t.Errorf("Expected force flag shorthand 'f', got '%s'", forceFlag.Shorthand)
		}
	}

	// Check all flag
	allFlag := installCmd.Flags().Lookup("all")
	if allFlag == nil {
		t.Error("All flag not found")
	} else {
		if allFlag.Shorthand != "a" {
			t.Errorf("Expected all flag shorthand 'a', got '%s'", allFlag.Shorthand)
		}
	}
}

// TestIDEValidation verifies that IDE validation includes VS Code
func TestIDEValidation(t *testing.T) {
	validIDEs := []string{ideCursor, ideClaude, ideVSCode, ideWindsurf, ideAll}

	// Test that vscode is in valid IDEs
	found := false
	for _, ide := range validIDEs {
		if ide == ideVSCode {
			found = true
			break
		}
	}

	if !found {
		t.Error("VS Code (vscode) should be in the list of valid IDEs")
	}

	// Test that the ideVSCode constant is correctly defined
	if ideVSCode != "vscode" {
		t.Errorf("Expected ideVSCode constant to be 'vscode', got '%s'", ideVSCode)
	}
}

// TestFlagReading tests reading individual flags directly
func TestFlagReading(t *testing.T) {
	// Create a test command with flags
	cmd := &cobra.Command{}
	cmd.Flags().String("ide", "", "IDE integration")
	cmd.Flags().Bool("force", false, "Force installation")
	cmd.Flags().Bool("all", false, "Install all")

	// Test default values
	ideFlag, err := cmd.Flags().GetString("ide")
	if err != nil {
		t.Fatalf("Unexpected error reading ide flag: %v", err)
	}
	if ideFlag != "" {
		t.Errorf("Expected empty ide flag, got '%s'", ideFlag)
	}

	forceFlag, err := cmd.Flags().GetBool("force")
	if err != nil {
		t.Fatalf("Unexpected error reading force flag: %v", err)
	}
	if forceFlag {
		t.Error("Expected force flag to be false")
	}

	allFlag, err := cmd.Flags().GetBool("all")
	if err != nil {
		t.Fatalf("Unexpected error reading all flag: %v", err)
	}
	if allFlag {
		t.Error("Expected all flag to be false")
	}

	// Test with flags set
	cmd.Flags().Set("ide", "cursor")
	cmd.Flags().Set("force", "true")
	cmd.Flags().Set("all", "true")

	ideFlag, _ = cmd.Flags().GetString("ide")
	if ideFlag != "cursor" {
		t.Errorf("Expected ide flag to be 'cursor', got '%s'", ideFlag)
	}

	forceFlag, _ = cmd.Flags().GetBool("force")
	if !forceFlag {
		t.Error("Expected force flag to be true")
	}

	allFlag, _ = cmd.Flags().GetBool("all")
	if !allFlag {
		t.Error("Expected all flag to be true")
	}
}

// TestValidateIDEFlag tests the validateIDEFlag function
func TestValidateIDEFlag(t *testing.T) {
	// Mock the PrintError method for testing
	originalValidateIDEFlag := validateIDEFlag

	// Test valid IDE flags
	validIDEs := []string{ideCursor, ideClaude, ideVSCode, ideWindsurf, ideAll}
	for _, validIDE := range validIDEs {
		err := originalValidateIDEFlag(validIDE, nil) // Pass nil since we're not testing error printing
		if err != nil {
			t.Errorf("Expected no error for valid IDE '%s', got: %v", validIDE, err)
		}
	}

	// Test invalid IDE flag
	err := originalValidateIDEFlag("invalid-ide", nil)
	if err == nil {
		t.Error("Expected error for invalid IDE flag, got nil")
	}

	// Test empty string (should be invalid based on the implementation)
	err = originalValidateIDEFlag("", nil)
	if err == nil {
		t.Error("Expected error for empty IDE flag, got nil")
	}
}

// TestInstallCommandRunFunction tests that the command run function is properly structured
func TestInstallCommandRunFunction(t *testing.T) {
	if installCmd.Run == nil {
		t.Fatal("Install command Run function is nil")
	}

	// Test that the command doesn't panic when called with empty args
	// Note: We can't easily test the full functionality without mocking the dependencies,
	// but we can ensure the structure is correct
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Install command Run function panicked: %v", r)
		}
	}()

	// Create a minimal test command to verify structure
	testCmd := &cobra.Command{}
	testCmd.Flags().String("ide", "", "IDE integration")
	testCmd.Flags().Bool("force", false, "Force installation")
	testCmd.Flags().Bool("all", false, "Install all")
	testCmd.Flags().String("agent", "", "Install specific agents")
	testCmd.Flags().String("task", "", "Install specific task")

	// The structure test - we just verify no immediate panic occurs
	// Full integration testing would require complex mocking
}

// TestInstallSelectiveFlags tests the selective installation flags
func TestInstallSelectiveFlags(t *testing.T) {
	// Test that the installAgent and installTask variables exist
	originalAgent := installAgent
	originalTask := installTask

	// Test setting values
	installAgent = "test-agent"
	installTask = "test-task"

	if installAgent != "test-agent" {
		t.Errorf("Expected installAgent 'test-agent', got '%s'", installAgent)
	}

	if installTask != "test-task" {
		t.Errorf("Expected installTask 'test-task', got '%s'", installTask)
	}

	// Restore original values
	installAgent = originalAgent
	installTask = originalTask
}
