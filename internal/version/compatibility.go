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
package version

import (
	"fmt"
)

// CompatibilityMatrix defines CLI-framework version compatibility
var CompatibilityMatrix = map[string][]string{
	"1.0.0": {"1.0.0", "1.0.1", "1.0.2"},
	"1.1.0": {"1.0.0", "1.0.1", "1.0.2", "1.1.0", "1.1.1"},
	"1.2.0": {"1.1.0", "1.1.1", "1.2.0", "1.2.1"},
	"2.0.0": {"2.0.0", "2.0.1", "2.1.0"},
}

// ValidateCompatibility checks if CLI and framework versions are compatible
func ValidateCompatibility(cliVersion, frameworkVersion string) error {
	cliVer, err := ParseVersion(cliVersion)
	if err != nil {
		return fmt.Errorf("invalid CLI version: %w", err)
	}

	frameworkVer, err := ParseVersion(frameworkVersion)
	if err != nil {
		return fmt.Errorf("invalid framework version: %w", err)
	}

	// Get compatible framework versions for this CLI version
	compatibleVersions, exists := CompatibilityMatrix[cliVer.String()]
	if !exists {
		return fmt.Errorf("no compatibility information for CLI version %s", cliVer.String())
	}

	// Check if framework version is in the compatible list
	for _, compatibleVer := range compatibleVersions {
		if frameworkVer.String() == compatibleVer {
			return nil
		}
	}

	return fmt.Errorf("CLI version %s is not compatible with framework version %s. Compatible framework versions: %v",
		cliVer.String(), frameworkVer.String(), compatibleVersions)
}

// GetCompatibilityMatrix returns the current compatibility matrix
func GetCompatibilityMatrix() map[string][]string {
	// Return a copy to prevent external modification
	matrix := make(map[string][]string)
	for k, v := range CompatibilityMatrix {
		versions := make([]string, len(v))
		copy(versions, v)
		matrix[k] = versions
	}
	return matrix
}

// GetCompatibleFrameworkVersions returns compatible framework versions for given CLI version
func GetCompatibleFrameworkVersions(cliVersion string) ([]string, error) {
	cliVer, err := ParseVersion(cliVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid CLI version: %w", err)
	}

	compatibleVersions, exists := CompatibilityMatrix[cliVer.String()]
	if !exists {
		return nil, fmt.Errorf("no compatibility information for CLI version %s", cliVer.String())
	}

	// Return a copy to prevent external modification
	versions := make([]string, len(compatibleVersions))
	copy(versions, compatibleVersions)
	return versions, nil
}

// IsFrameworkVersionCompatible checks if a specific framework version is compatible with CLI
func IsFrameworkVersionCompatible(cliVersion, frameworkVersion string) bool {
	err := ValidateCompatibility(cliVersion, frameworkVersion)
	return err == nil
}
