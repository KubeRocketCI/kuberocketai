package cmd

import (
	"testing"
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
	validIDEs := []string{ideCursor, ideClaude, ideVSCode, "windsurf", ideAll}

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
