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
package update

import (
	"fmt"
	"time"

	"github.com/KubeRocketCI/kuberocketai/internal/github"
	"github.com/KubeRocketCI/kuberocketai/internal/version"
)

// Checker handles update detection
type Checker struct {
	githubClient *github.Client
	owner        string
	repo         string
}

// NewChecker creates a new update checker
func NewChecker() *Checker {
	return &Checker{
		githubClient: github.NewClient(),
		owner:        "KubeRocketCI",
		repo:         "kuberocketai",
	}
}

// NewCheckerWithClient creates a new update checker with a custom GitHub client
func NewCheckerWithClient(client *github.Client, owner, repo string) *Checker {
	return &Checker{
		githubClient: client,
		owner:        owner,
		repo:         repo,
	}
}

// UpdateInfo contains information about available updates
type UpdateInfo struct {
	CurrentVersion    string              `json:"current_version"`
	LatestVersion     string              `json:"latest_version"`
	IsUpdateAvailable bool                `json:"is_update_available"`
	LatestRelease     *github.ReleaseInfo `json:"latest_release,omitempty"`
	Error             string              `json:"error,omitempty"`
	FallbackURL       string              `json:"fallback_url,omitempty"`
}

// CheckForUpdates checks for available updates
func (c *Checker) CheckForUpdates(currentVersion string) *UpdateInfo {
	updateInfo := &UpdateInfo{
		CurrentVersion: currentVersion,
		FallbackURL:    github.GetReleasesURL(c.owner, c.repo),
	}

	// Get latest release from GitHub
	release, err := c.githubClient.GetLatestRelease(c.owner, c.repo)
	if err != nil {
		updateInfo.Error = fmt.Sprintf("Failed to fetch latest release: %v", err)
		return updateInfo // Return with error info, don't fail completely
	}

	// Skip draft and prerelease versions
	if !release.IsStable() {
		updateInfo.Error = "Latest release is not stable (draft or prerelease)"
		return updateInfo
	}

	updateInfo.LatestVersion = release.GetVersion()
	updateInfo.LatestRelease = release.ToReleaseInfo()

	// Compare versions
	isNewer, err := version.CompareVersions(currentVersion, updateInfo.LatestVersion)
	if err != nil {
		updateInfo.Error = fmt.Sprintf("Failed to compare versions: %v", err)
		return updateInfo
	}

	updateInfo.IsUpdateAvailable = isNewer
	return updateInfo
}

// CheckForUpdatesWithRetry checks for updates with retry logic
func (c *Checker) CheckForUpdatesWithRetry(currentVersion string, retries int) *UpdateInfo {
	for attempt := 0; attempt <= retries; attempt++ {
		updateInfo := c.CheckForUpdates(currentVersion)
		if updateInfo.Error == "" {
			return updateInfo
		}

		if attempt < retries {
			// Wait before retrying
			time.Sleep(time.Duration(attempt+1) * time.Second)
		}
	}

	// If all retries failed, return info with fallback
	updateInfo := &UpdateInfo{
		CurrentVersion: currentVersion,
		Error:          "All retry attempts failed",
		FallbackURL:    github.GetReleasesURL(c.owner, c.repo),
	}

	return updateInfo
}
