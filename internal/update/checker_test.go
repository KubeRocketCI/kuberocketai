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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KubeRocketCI/kuberocketai/internal/github"
)

func TestNewChecker(t *testing.T) {
	checker := NewChecker()

	assert.Equal(t, "KubeRocketCI", checker.owner)
	assert.Equal(t, "kuberocketai", checker.repo)
	assert.NotNil(t, checker.githubClient)
}

func TestNewCheckerWithClient(t *testing.T) {
	client := github.NewClient()
	checker := NewCheckerWithClient(client, "test-owner", "test-repo")

	assert.Equal(t, "test-owner", checker.owner)
	assert.Equal(t, "test-repo", checker.repo)
	assert.Equal(t, client, checker.githubClient)
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
		err := json.NewEncoder(w).Encode(release)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with older version
	updateInfo := checker.CheckForUpdates("v1.0.0")

	assert.Empty(t, updateInfo.Error)
	assert.True(t, updateInfo.IsUpdateAvailable)
	assert.Equal(t, "2.0.0", updateInfo.LatestVersion)
	assert.Equal(t, "v1.0.0", updateInfo.CurrentVersion)
	assert.NotNil(t, updateInfo.LatestRelease)
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
		err := json.NewEncoder(w).Encode(release)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with same version
	updateInfo := checker.CheckForUpdates("v1.0.0")

	assert.Empty(t, updateInfo.Error)
	assert.False(t, updateInfo.IsUpdateAvailable)
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
		err := json.NewEncoder(w).Encode(release)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with newer local version
	updateInfo := checker.CheckForUpdates("v2.0.0")

	assert.Empty(t, updateInfo.Error)
	assert.False(t, updateInfo.IsUpdateAvailable)
}

func TestCheckForUpdatesError(t *testing.T) {
	// Create test server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Server Error"))
		require.NoError(t, err)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	updateInfo := checker.CheckForUpdates("v1.0.0")
	assert.NotEmpty(t, updateInfo.Error)
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
		err := json.NewEncoder(w).Encode(release)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with invalid version format
	updateInfo := checker.CheckForUpdates("invalid-version")
	assert.NotEmpty(t, updateInfo.Error)
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

	assert.True(t, updateInfo.IsUpdateAvailable)
	assert.Equal(t, "v1.0.0", updateInfo.CurrentVersion)
	assert.Equal(t, "v2.0.0", updateInfo.LatestVersion)
	assert.Equal(t, releaseInfo, updateInfo.LatestRelease)
	assert.Equal(t, "https://github.com/test/repo/releases", updateInfo.FallbackURL)
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
		err := json.NewEncoder(w).Encode(release)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test successful request (should not need retries)
	updateInfo := checker.CheckForUpdatesWithRetry("v1.0.0", 3)

	assert.Empty(t, updateInfo.Error)
	assert.True(t, updateInfo.IsUpdateAvailable)
}

func TestCheckForUpdatesWithRetryAllFail(t *testing.T) {
	// Create test server that always returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Server Error"))
		require.NoError(t, err)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test retry mechanism
	startTime := time.Now()
	updateInfo := checker.CheckForUpdatesWithRetry("v1.0.0", 2)
	elapsed := time.Since(startTime)

	assert.NotEmpty(t, updateInfo.Error)
	assert.GreaterOrEqual(t, elapsed, 3*time.Second) // Should wait for retries
	assert.NotEmpty(t, updateInfo.FallbackURL)
}

func TestCheckForUpdatesWithRetrySucceedAfterFailure(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts <= 2 {
			// Fail first two attempts
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Server Error"))
			require.NoError(t, err)
		} else {
			// Succeed on third attempt
			release := github.Release{
				TagName:     "v2.0.0",
				PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
				HTMLURL:     "https://github.com/test/repo/releases/tag/v2.0.0",
				Body:        "New version available",
			}

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(release)
			require.NoError(t, err)
		}
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test retry mechanism with eventual success
	updateInfo := checker.CheckForUpdatesWithRetry("v1.0.0", 3)

	assert.Empty(t, updateInfo.Error)
	assert.True(t, updateInfo.IsUpdateAvailable)
	assert.Equal(t, 3, attempts) // Should have made 3 attempts
}

func TestCheckForUpdatesPrerelease(t *testing.T) {
	// Create test server with prerelease
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := github.Release{
			TagName:     "v2.0.0-beta.1",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v2.0.0-beta.1",
			Body:        "Prerelease version",
			Prerelease:  true,
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(release)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with prerelease version - should be rejected
	updateInfo := checker.CheckForUpdates("v1.0.0")

	assert.NotEmpty(t, updateInfo.Error)
	assert.Contains(t, updateInfo.Error, "not stable")
	assert.False(t, updateInfo.IsUpdateAvailable)
}

func TestCheckForUpdatesDraft(t *testing.T) {
	// Create test server with draft release
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := github.Release{
			TagName:     "v2.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v2.0.0",
			Body:        "Draft version",
			Draft:       true,
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(release)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	// Test with draft release - should be rejected
	updateInfo := checker.CheckForUpdates("v1.0.0")

	assert.NotEmpty(t, updateInfo.Error)
	assert.Contains(t, updateInfo.Error, "not stable")
	assert.False(t, updateInfo.IsUpdateAvailable)
}
