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

import "fmt"

// GetReleasesURL returns the GitHub releases page URL for manual checking
func GetReleasesURL(owner, repo string) string {
	return fmt.Sprintf("https://github.com/%s/%s/releases", owner, repo)
}

// GetFallbackMessage returns a user-friendly error message with manual check instructions
func GetFallbackMessage(owner, repo string) string {
	url := GetReleasesURL(owner, repo)
	return fmt.Sprintf("Unable to check for updates. Visit: %s to check manually", url)
}

// GetKuberocketaiReleasesURL returns the specific KubeRocketAI releases URL
func GetKuberocketaiReleasesURL() string {
	return GetReleasesURL("KubeRocketCI", "kuberocketai")
}

// GetKuberocketaiFallbackMessage returns the fallback message for KubeRocketAI
func GetKuberocketaiFallbackMessage() string {
	return GetFallbackMessage("KubeRocketCI", "kuberocketai")
}
