## KubeRocketAI CLI Release

Welcome to this release of KubeRocketAI CLI!

## Installation

### Homebrew (macOS)

```bash
brew tap KubeRocketCI/homebrew-tap
brew install krci-ai
```

{{ if .Versions -}}
<a name="unreleased"></a>

### [Unreleased]

{{ if .Unreleased.CommitGroups -}}
{{ range .Unreleased.CommitGroups -}}

#### {{ .Title }}

{{ range .Commits -}}

* [{{ .Hash.Short }}]({{ $.Info.RepositoryURL }}/commit/{{ .Hash.Long }}) {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}

{{ range .Versions -}}
<a name="{{ .Tag.Name }}"></a>

### {{ if .Tag.Previous }}[{{ .Tag.Name }}]{{ else }}{{ .Tag.Name }}{{ end }} - {{ datetime "2006-01-02" .Tag.Date }}

{{ if .CommitGroups -}}
{{ range .CommitGroups -}}

#### {{ .Title }}

{{ range .Commits -}}

* [{{ .Hash.Short }}]({{ $.Info.RepositoryURL }}/commit/{{ .Hash.Long }}) {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }}{{ if .Notes }} ({{ range .Notes }}{{ .Body }}{{ end }}){{ end }}
{{ end }}
{{ end -}}
{{ end -}}

{{- if .RevertCommits -}}

#### Reverts

{{ range .RevertCommits -}}

* {{ .Revert.Header }}
{{ end }}
{{ end -}}

{{- if .MergeCommits -}}

#### Pull Requests

{{ range .MergeCommits -}}

* {{ .Header }}
{{ end }}
{{ end -}}

{{- if .NoteGroups -}}
{{ range .NoteGroups -}}

#### {{ .Title }}

{{ range .Notes }}
{{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}

{{- if .Versions }}
[Unreleased]: {{ .Info.RepositoryURL }}/compare/{{ $latest := index .Versions 0 }}{{ $latest.Tag.Name }}...HEAD
{{ range .Versions -}}
{{ if .Tag.Previous -}}
[{{ .Tag.Name }}]: {{ $.Info.RepositoryURL }}/compare/{{ .Tag.Previous.Name }}...{{ .Tag.Name }}
{{ end -}}
{{ end -}}
{{ end -}}
