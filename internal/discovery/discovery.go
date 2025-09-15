package discovery

import (
	"fmt"
	"os"
)

// GetProjectRoot returns the project root directory.
// KRCI_AI_PROJECT_DIR is the environment variable that can be set to the project root directory.
// If not set, the current working directory is returned.
func GetProjectRoot() (string, error) {
	if root := os.Getenv("KRCI_AI_PROJECT_DIR"); root != "" {
		return root, nil
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	return pwd, nil
}
