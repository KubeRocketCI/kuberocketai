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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/KubeRocketCI/kuberocketai/internal/github"
)

func TestNewChecker(t *testing.T) {
	checker := NewChecker()

	if checker.owner != "KubeRocketCI" {
		t.Errorf("Expected owner 'KubeRocketCI', got %s", checker.owner)
	}

	if checker.repo != "kuberocketai" {
		t.Errorf("Expected repo 'kuberocketai', got %s", checker.repo)
	}

	if checker.githubClient == nil {
		t.Error("GitHub client should not be nil")
	}
}

func TestNewCheckerWithClient(t *testing.T) {
	client := github.NewClient()
	checker := NewCheckerWithClient(client, "test-owner", "test-repo")

	if checker.owner != "test-owner" {
		t.Errorf("Expected owner 'test-owner', got %s", checker.owner)
	}

	if checker.repo != "test-repo" {
		t.Errorf("Expected repo 'test-repo', got %s", checker.repo)
	}

	if checker.githubClient != client {
		t.Error("GitHub client should be the same instance")
	}
}

func TestCheckForUpdates(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := github.Release{
			TagName:     "v2.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v2.0.0",
			Body:        "New version available",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with older version
	updateInfo := checker.CheckForUpdates("v1.0.0")
	if updateInfo.Error != "" {
		t.Fatalf("CheckForUpdates failed: %v", updateInfo.Error)
	}

	if !updateInfo.IsUpdateAvailable {
		t.Error("Expected update to be available")
	}

	if updateInfo.LatestVersion != "2.0.0" {
		t.Errorf("Expected latest version 2.0.0, got %s", updateInfo.LatestVersion)
	}

	if updateInfo.CurrentVersion != "v1.0.0" {
		t.Errorf("Expected current version v1.0.0, got %s", updateInfo.CurrentVersion)
	}

	if updateInfo.LatestRelease == nil {
		t.Error("Expected latest release info to be present")
	}
}

func TestCheckForUpdatesNoUpdate(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := github.Release{
			TagName:     "v1.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v1.0.0",
			Body:        "Current version",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with same version
	updateInfo := checker.CheckForUpdates("v1.0.0")
	if updateInfo.Error != "" {
		t.Fatalf("CheckForUpdates failed: %v", updateInfo.Error)
	}

	if updateInfo.IsUpdateAvailable {
		t.Error("Expected no update to be available")
	}

	if updateInfo.LatestVersion != "1.0.0" {
		t.Errorf("Expected latest version 1.0.0, got %s", updateInfo.LatestVersion)
	}
}

func TestCheckForUpdatesNewerLocal(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := github.Release{
			TagName:     "v1.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v1.0.0",
			Body:        "Older version",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with newer local version
	updateInfo := checker.CheckForUpdates("v2.0.0")
	if updateInfo.Error != "" {
		t.Fatalf("CheckForUpdates failed: %v", updateInfo.Error)
	}

	if updateInfo.IsUpdateAvailable {
		t.Error("Expected no update to be available when local version is newer")
	}
}

func TestCheckForUpdatesError(t *testing.T) {
	// Create test server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	updateInfo := checker.CheckForUpdates("v1.0.0")

	if updateInfo.Error == "" {
		t.Error("Expected error message in updateInfo when server returns 500")
	}
}

func TestCheckForUpdatesInvalidVersion(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := github.Release{
			TagName:     "v1.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v1.0.0",
			Body:        "Valid version",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with invalid version format
	updateInfo := checker.CheckForUpdates("invalid-version")

	if updateInfo.Error == "" {
		t.Error("Expected error message in updateInfo for invalid version format")
	}
}

func TestUpdateInfoStruct(t *testing.T) {
	releaseInfo := &github.ReleaseInfo{
		Version:     "2.0.0",
		TagName:     "v2.0.0",
		Name:        "Release v2.0.0",
		Body:        "New features",
		URL:         "https://github.com/test/repo/releases/tag/v2.0.0",
		PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
		IsStable:    true,
	}

	updateInfo := &UpdateInfo{
		IsUpdateAvailable: true,
		CurrentVersion:    "v1.0.0",
		LatestVersion:     "v2.0.0",
		LatestRelease:     releaseInfo,
		FallbackURL:       "https://github.com/test/repo/releases",
	}

	if !updateInfo.IsUpdateAvailable {
		t.Error("Expected IsUpdateAvailable to be true")
	}

	if updateInfo.CurrentVersion != "v1.0.0" {
		t.Errorf("Expected CurrentVersion v1.0.0, got %s", updateInfo.CurrentVersion)
	}

	if updateInfo.LatestVersion != "v2.0.0" {
		t.Errorf("Expected LatestVersion v2.0.0, got %s", updateInfo.LatestVersion)
	}

	if updateInfo.LatestRelease == nil {
		t.Error("Expected LatestRelease to be present")
	} else {
		if updateInfo.LatestRelease.URL != "https://github.com/test/repo/releases/tag/v2.0.0" {
			t.Errorf("Expected release URL https://github.com/test/repo/releases/tag/v2.0.0, got %s", updateInfo.LatestRelease.URL)
		}

		expectedTime := time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC)
		if !updateInfo.LatestRelease.PublishedAt.Equal(expectedTime) {
			t.Errorf("Expected PublishedAt %v, got %v", expectedTime, updateInfo.LatestRelease.PublishedAt)
		}

		if updateInfo.LatestRelease.Body != "New features" {
			t.Errorf("Expected Body 'New features', got %s", updateInfo.LatestRelease.Body)
		}
	}

	if updateInfo.FallbackURL != "https://github.com/test/repo/releases" {
		t.Errorf("Expected FallbackURL https://github.com/test/repo/releases, got %s", updateInfo.FallbackURL)
	}
}

func TestCheckForUpdatesWithRetrySuccess(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := github.Release{
			TagName:     "v2.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v2.0.0",
			Body:        "New version available",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with successful first attempt
	updateInfo := checker.CheckForUpdatesWithRetry("v1.0.0", 3)
	if updateInfo.Error != "" {
		t.Fatalf("CheckForUpdatesWithRetry failed: %v", updateInfo.Error)
	}

	if !updateInfo.IsUpdateAvailable {
		t.Error("Expected update to be available")
	}

	if updateInfo.LatestVersion != "2.0.0" {
		t.Errorf("Expected latest version 2.0.0, got %s", updateInfo.LatestVersion)
	}

	if updateInfo.Error != "" {
		t.Errorf("Expected no error, got: %s", updateInfo.Error)
	}
}

func TestCheckForUpdatesWithRetryAllFail(t *testing.T) {
	// Create test server that always returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test retry mechanism
	startTime := time.Now()
	updateInfo := checker.CheckForUpdatesWithRetry("v1.0.0", 2)
	elapsed := time.Since(startTime)

	if updateInfo.Error == "" {
		t.Error("Expected error message in updateInfo when all retries fail")
	}

	// Should have waited for retries (at least 1 second for first retry + 2 seconds for second retry)
	if elapsed < 3*time.Second {
		t.Errorf("Expected to wait at least 3 seconds for retries, but took %v", elapsed)
	}

	if updateInfo.FallbackURL == "" {
		t.Error("Expected fallback URL to be set when all retries fail")
	}
}

func TestCheckForUpdatesWithRetrySucceedAfterFailure(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts <= 2 {
			// Fail first two attempts
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Server Error"))
			return
		}

		// Succeed on third attempt
		release := github.Release{
			TagName:     "v2.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v2.0.0",
			Body:        "New version available",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test retry succeeding after failures
	updateInfo := checker.CheckForUpdatesWithRetry("v1.0.0", 3)
	if updateInfo.Error != "" {
		t.Fatalf("CheckForUpdatesWithRetry failed: %v", updateInfo.Error)
	}

	if !updateInfo.IsUpdateAvailable {
		t.Error("Expected update to be available")
	}

	if updateInfo.Error != "" {
		t.Errorf("Expected no error after successful retry, got: %s", updateInfo.Error)
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

func TestCheckForUpdatesPrerelease(t *testing.T) {
	// Create test server that returns a prerelease
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := github.Release{
			TagName:     "v2.0.0-beta.1",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v2.0.0-beta.1",
			Body:        "Beta release",
			Prerelease:  true, // This is a prerelease
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with prerelease version (should be skipped)
	updateInfo := checker.CheckForUpdates("v1.0.0")

	if updateInfo.IsUpdateAvailable {
		t.Error("Expected no update to be available when latest release is prerelease")
	}

	if updateInfo.Error == "" {
		t.Error("Expected error message about prerelease")
	}
}

func TestCheckForUpdatesDraft(t *testing.T) {
	// Create test server that returns a draft
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := github.Release{
			TagName:     "v2.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v2.0.0",
			Body:        "Draft release",
			Draft:       true, // This is a draft
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with draft version (should be skipped)
	updateInfo := checker.CheckForUpdates("v1.0.0")

	if updateInfo.IsUpdateAvailable {
		t.Error("Expected no update to be available when latest release is draft")
	}

	if updateInfo.Error == "" {
		t.Error("Expected error message about draft release")
	}
}
