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
package main

import (
	"embed"

	"github.com/KubeRocketCI/kuberocketai/cmd/krci-ai/cmd"
)

// EmbeddedAssets contains all assets embedded at build time
//
//go:embed assets
var EmbeddedAssets embed.FS

// Build-time variables
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date, builtBy)
	cmd.SetEmbeddedAssets(EmbeddedAssets)
	cmd.Execute()
}
