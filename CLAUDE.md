# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

KubeRocketAI is an "AI-as-Code" framework providing a centralized golden library of AI agents for software development workflows. Built as a Go CLI application (`krci-ai`) that embeds framework assets and supports multi-IDE integration.

## Development Commands

### Build and Testing

```bash
make build          # Build CLI binary
make test           # Run unit tests with coverage
make lint           # Run linting (golangci-lint)
make fmt            # Format code
make ci             # Run full CI pipeline locally
make build-all      # Cross-platform builds
make tools          # Install development tools
```

### CLI Usage

```bash
krci-ai install --ide=cursor  # Install framework with IDE integration
krci-ai list agents           # List available agents
krci-ai validate             # Validate framework components
```

## Framework Installation

### First-Time Setup

```bash
# Build the CLI
make build

# Install framework locally
./dist/krci-ai install --ide=cursor  # For Cursor IDE
./dist/krci-ai install --ide=claude  # For Claude Code
./dist/krci-ai install --all         # All IDE integrations

# Verify installation
./dist/krci-ai validate
```

### Directory Structure Created

- `.krci-ai/agents/` - 6 role-based agent definitions
- `.krci-ai/tasks/` - Common workflow templates
- `.krci-ai/templates/` - Output formatting templates
- `.krci-ai/data/` - Reference data and standards
- `.cursor/rules/` - IDE-specific integration files (if --ide used)

## Architecture

### Core Structure

- **`/cmd/krci-ai/`** - Main CLI entry point with embedded assets
- **`/internal/`** - Core application logic (assets, CLI, engine, validation)
- **`/cmd/krci-ai/assets/framework/core/`** - Embedded framework components

### Key Architectural Patterns

**Embedded Assets Pattern**: Framework components (agents, tasks, templates) are embedded in the binary using Go's `embed.FS` for offline-first operation.

**Role-Based Agent System**: Six SDLC roles defined as YAML configurations:

- `agents/dev.yaml` - Developer implementation assistance
- `agents/architect.yaml` - Technical architecture decisions
- `agents/ba.yaml` - Requirements analysis
- `agents/pm.yaml` - Product strategy
- `agents/po.yaml` - Backlog management
- `agents/qa.yaml` - Testing and quality assurance

**Multi-IDE Integration**: Complete polymorphic IDE integration system:

- Cursor IDE: `.cursor/rules/*.mdc` files
- Claude Code: `.claude/commands/*.md` files
- VS Code: `.github/chatmodes/*.chatmode.md` files
- Windsurf IDE: Support included
- Universal: `--ide=all` installs all integrations

**Template-Driven Output**: Consistent formatting across roles using embedded templates in `templates/` directory.

### Processing Engine

The YAML processing engine (`internal/engine/processor/`) handles:

- Variable substitution in agent definitions
- Template rendering for consistent outputs
- Schema validation using JSON schemas in `schemas/`

## Development Guidelines

- **Go 1.24.4** with minimal dependencies (cobra, yaml, jsonschema)
- **Conventional Commits** required (feat, fix, docs, chore, etc.)
- **golangci-lint** with strict rules - must pass before commits
- **Comprehensive testing** - unit tests with coverage requirements
- **Semantic versioning** with automated releases via goreleaser

## Key Files

- `cmd/krci-ai/main.go` - Application entry point with asset embedding
- `cmd/krci-ai/cmd/install.go` - Complete install command implementation
- `internal/assets/installer.go` - Asset installation and extraction logic
- `internal/assets/discovery.go` - Agent discovery and YAML parsing
- `internal/cli/output.go` - Colorized CLI output utilities
- `Makefile` - Development workflow automation
- `.goreleaser.yml` - Multi-platform release configuration
