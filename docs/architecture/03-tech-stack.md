# 3. Tech Stack

The KubeRocketAI CLI is built with a focus on rich user experience, cross-platform distribution, and seamless integration. The technology choices are organized around key capabilities that deliver multi-OS support, colorized interfaces, rich menus, and both interactive and non-interactive operation modes.

## 1. Core Implementation Technologies

Foundation technologies for framework management and business logic.

| Technology | Version | Purpose | Rationale |
|------------|---------|---------|-----------|
| Go | 1.24.4 | Primary development language | Fast, statically compiled, cross-platform, excellent CLI ecosystem |
| Cobra | 1.8.0 | CLI framework and command structure | Industry standard for Go CLIs, supports subcommands, extensibility |
| gopkg.in/yaml.v3 | v3.0.1 | Framework component parsing | Parse YAML frontmatter in agents, tasks, templates, data |
| gojsonschema/gojsonschema | v1.2.0 | Component validation | Schema validation for framework integrity and reliability |
| spf13/viper | v1.18.2 | Configuration management | Unified configuration with support for multiple formats and sources |
| embed | 1.24.4 | Asset bundling | Bundle framework components into self-contained binary |

## 2. CLI User Experience Technologies

Rich, interactive, and colorized user interface capabilities.

| Technology | Version | Purpose | User Experience Benefit |
|------------|---------|---------|------------------------|
| fatih/color | v1.17.0 | Colorized terminal output | **Colorized support** - Enhanced readability and visual feedback |
| AlecAivazis/survey | v2.3.7 | Simple interactive prompts | **Basic interactions** - Quick prompts and confirmations |
| charmbracelet/bubbletea | v0.27.0 | Advanced TUI framework (optional) | **Rich TUI interfaces** - Complex, stateful CLI applications |
| charmbracelet/bubbles | v0.18.0 | TUI component library (optional) | **Professional UI components** - Lists, tables, progress bars |
| spf13/pflag | v1.0.5 | Advanced flag parsing | **Non-interactive mode** - Batch processing and automation support |
| mattn/go-isatty | v0.0.20 | Terminal capability detection | Smart color/interactive mode detection across platforms |

## 3. Cross-Platform Distribution Technologies

Multi-OS support and package management integration.

| Technology | Version | Purpose | Platform Benefit |
|------------|---------|---------|-----------------|
| GoReleaser | v1.26.2 | Automated cross-platform builds | **Multi-OS support** - Automated builds for macOS, Linux, Windows |
| Homebrew Formula | N/A | macOS package management | **Brew install on Mac** - Native macOS installation experience |
| Chocolatey Package | N/A | Windows package management | **Multi-OS support** - Native Windows installation experience |
| GitHub Actions | N/A | Release automation | Automated build, test, and release pipeline |

## 4. Integration Technologies

Git, IDE, and file system integration capabilities.

| Technology | Version | Purpose | Integration Benefit |
|------------|---------|---------|-------------------|
| filepath | 1.24.4 | Cross-platform path handling | Consistent file operations across operating systems |
| os/exec | 1.24.4 | Git command integration | Native git integration for version control operations |
| Go Standard Library | 1.24.4 | File system operations | Reliable file I/O for framework component management |
| Testing | 1.24.4 | Unit and integration testing | Built-in testing with table-driven test patterns |

## Technology Selection Rationale

**Core Implementation**: Chosen for reliability, performance, and maintainability. Go provides excellent cross-platform support with static compilation, while Cobra is the industry standard for professional CLI applications used by Kubernetes, Helm, and other major projects.

**CLI User Experience**: Selected to deliver the rich, colorized, interactive experience specified in requirements. The combination of fatih/color, AlecAivazis/survey, and spf13/pflag provides both interactive and non-interactive operation modes with smart terminal detection. Bubble Tea and Bubbles are available as optional enhancements for advanced TUI experiences (complex workflows, dashboard-style interfaces, rich component interactions) while survey handles simple prompts and confirmations.

**Cross-Platform Distribution**: Designed to enable native package manager support across all major platforms. GoReleaser automates the build process while Homebrew and Chocolatey provide familiar installation experiences for users.

**Integration**: Focused on seamless operation within existing developer workflows. Native Git integration and standard file system operations ensure the framework works naturally with version control and IDE environments without requiring proprietary plugins or extensions.
