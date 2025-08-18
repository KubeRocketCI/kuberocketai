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
package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReleasesURL(t *testing.T) {
	tests := []struct {
		owner string
		repo  string
		want  string
	}{
		{
			owner: "owner1",
			repo:  "repo1",
			want:  "https://github.com/owner1/repo1/releases",
		},
		{
			owner: "KubeRocketCI",
			repo:  "kuberocketai",
			want:  "https://github.com/KubeRocketCI/kuberocketai/releases",
		},
		{
			owner: "",
			repo:  "",
			want:  "https://github.com///releases",
		},
	}

	for _, tt := range tests {
		t.Run(tt.owner+"/"+tt.repo, func(t *testing.T) {
			got := GetReleasesURL(tt.owner, tt.repo)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetFallbackMessage(t *testing.T) {
	tests := []struct {
		owner string
		repo  string
		want  string
	}{
		{
			owner: "test",
			repo:  "repo",
			want:  "Unable to check for updates. Visit: https://github.com/test/repo/releases to check manually",
		},
		{
			owner: "KubeRocketCI",
			repo:  "kuberocketai",
			want:  "Unable to check for updates. Visit: https://github.com/KubeRocketCI/kuberocketai/releases to check manually",
		},
	}

	for _, tt := range tests {
		t.Run(tt.owner+"/"+tt.repo, func(t *testing.T) {
			got := GetFallbackMessage(tt.owner, tt.repo)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetKuberocketaiReleasesURL(t *testing.T) {
	expected := "https://github.com/KubeRocketCI/kuberocketai/releases"
	got := GetKuberocketaiReleasesURL()
	assert.Equal(t, expected, got)
}

func TestGetKuberocketaiFallbackMessage(t *testing.T) {
	expected := "Unable to check for updates. Visit: https://github.com/KubeRocketCI/kuberocketai/releases to check manually"
	got := GetKuberocketaiFallbackMessage()
	assert.Equal(t, expected, got)
}
