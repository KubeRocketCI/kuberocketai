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
	"runtime"

	"github.com/Masterminds/semver/v3"
)

// Build variables (set during build time)
var (
	Version   = "dev"
	Commit    = "unknown"
	Date      = "unknown"
	BuiltBy   = "unknown"
	Framework = "1.0.0"
)

// ParseVersion parses a version string and returns a semver.Version
func ParseVersion(version string) (*semver.Version, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return nil, fmt.Errorf("invalid version format '%s': %w", version, err)
	}
	return v, nil
}

// CompareVersions compares current version with latest version
// Returns true if latest is newer than current, false otherwise
func CompareVersions(current, latest string) (bool, error) {
	currentVer, err := ParseVersion(current)
	if err != nil {
		return false, fmt.Errorf("failed to parse current version: %w", err)
	}

	latestVer, err := ParseVersion(latest)
	if err != nil {
		return false, fmt.Errorf("failed to parse latest version: %w", err)
	}

	return latestVer.GreaterThan(currentVer), nil
}

// GetCurrentVersion returns the current CLI version
func GetCurrentVersion() string {
	return Version
}

// GetVersionInfo returns detailed version information
func GetVersionInfo() VersionInfo {
	return VersionInfo{
		Version:   Version,
		Commit:    Commit,
		Date:      Date,
		BuiltBy:   BuiltBy,
		Framework: Framework,
		GoVersion: runtime.Version(),
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// VersionInfo contains detailed version information
type VersionInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	Date      string `json:"date"`
	BuiltBy   string `json:"built_by"`
	Framework string `json:"framework"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

// String returns a formatted version string
func (v VersionInfo) String() string {
	return fmt.Sprintf("krci-ai version %s\nCommit: %s\nBuilt: %s by %s\nFramework: %s\nGo: %s\nPlatform: %s",
		v.Version, v.Commit, v.Date, v.BuiltBy, v.Framework, v.GoVersion, v.Platform)
}
