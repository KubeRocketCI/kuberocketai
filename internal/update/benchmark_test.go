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

func BenchmarkNewChecker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewChecker()
	}
}

func BenchmarkCheckerWithClient(b *testing.B) {
	client := github.NewClient()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewCheckerWithClient(client, "test", "repo")
	}
}

func BenchmarkCheckForUpdates(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = checker.CheckForUpdates("v1.0.0")
	}
}

func BenchmarkCheckForUpdatesWithRetry(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = checker.CheckForUpdatesWithRetry("v1.0.0", 2)
	}
}

func BenchmarkUpdateInfoCreation(b *testing.B) {
	releaseInfo := &github.ReleaseInfo{
		Version:     "2.0.0",
		TagName:     "v2.0.0",
		Name:        "Release v2.0.0",
		Body:        "New features",
		URL:         "https://github.com/test/repo/releases/tag/v2.0.0",
		PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
		IsStable:    true,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = &UpdateInfo{
			IsUpdateAvailable: true,
			CurrentVersion:    "v1.0.0",
			LatestVersion:     "v2.0.0",
			LatestRelease:     releaseInfo,
			FallbackURL:       "https://github.com/test/repo/releases",
		}
	}
}

// Benchmark memory allocations
func BenchmarkCheckForUpdatesMemory(b *testing.B) {
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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		updateInfo, _ := checker.CheckForUpdates("v1.0.0")
		_ = updateInfo
	}
}

// Concurrent update checking
func BenchmarkConcurrentUpdateChecking(b *testing.B) {
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

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = checker.CheckForUpdates("v1.0.0")
		}
	})
}

// Benchmark different version comparison scenarios
func BenchmarkVersionComparisons(b *testing.B) {
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

	versions := []string{
		"v1.0.0",
		"v1.5.0",
		"v2.0.0",
		"v3.0.0",
		"v1.0.0-alpha.1",
		"v1.0.0-beta.2",
		"v1.0.0-rc.1",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		version := versions[i%len(versions)]
		_, _ = checker.CheckForUpdates(version)
	}
}

// Benchmark error handling scenarios
func BenchmarkErrorHandling(b *testing.B) {
	// Create test server that returns errors
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server Error"))
	}))
	defer server.Close()

	client := github.NewClient()
	client.BaseURL = server.URL
	checker := NewCheckerWithClient(client, "test", "repo")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = checker.CheckForUpdates("v1.0.0")
	}
}
