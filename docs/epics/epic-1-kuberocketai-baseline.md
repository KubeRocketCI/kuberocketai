# Epic 1: KubeRocketAI Baseline

## Status

| Field                | Value                    |
|----------------------|--------------------------|
| Status               | Complete                 |
| Priority             | P0 (Critical)           |
| Epic Owner           | Development Team        |
| Timeline             | Week 1 (1 week) - DONE  |

## Overview

### Problem Statement

Enterprise development teams face critical workflow inefficiencies when adopting AI assistance tools. Developers spend 5-10 minutes daily adjusting AI-generated code to match project standards due to inconsistent agent configurations across repositories. Context switching between IDE environments and external AI platforms disrupts development flow, while teams lack standardized AI workflows and governance frameworks. This Epic addresses the foundational need for locally-controlled, version-controlled AI agent management that integrates seamlessly with existing development workflows.

### Goal

Enable KubeRocketAI framework foundation for 100% of development teams within 1 week, delivering 5 core agents and CLI tool that reduces manual AI-code fixes by establishing local-first agent management infrastructure with offline installation capabilities.

### Target Users

**Primary User**: Emily - Enterprise Development Lead managing 8-15 developers across 3-4 microservices in Fortune 500 companies with established DevOps culture.

**User Context**: 500+ Active Development Leads globally seeking to standardize AI agent configurations across teams while maintaining project-specific customizations and organizational governance.

## Scope

### What's Included

**PRD Requirements Addressed:**

- **BR1 [P0]**: User can install KubeRocketAI framework with single command (krci-ai install)
- **BR2 [P0]**: User can access 5 core SDLC agent definitions (PM, Architect, Developer, QA, BA) as Markdown files
- **BR4 [P0]**: User can use agents directly in target IDEs (Cursor @agent, Claude Code /agent) without platform switching
- **NFR1 [P0]**: System supports offline installation without network dependencies
- **NFR3 [P0]**: System maintains multi-platform compatibility via Homebrew and GitHub releases

**Core Deliverables:**

- 5 core agent definitions (PM, Architect, Developer, QA, BA) as YAML files
- Basic CLI tool with install and validate commands
- Homebrew tap and cross-platform distribution
- Self-dogfooding development environment

### What's Not Included

- Advanced validation engine (Epic 2)
- Automated framework installation (Epic 3)
- IDE-specific configuration generation (Epic 4)
- Community marketplace features (Post-MVP)

### Dependencies

Epic Dependencies:

- None (Foundation epic)

System Dependencies:

- Go development environment (1.24+)
- GitHub Actions runner capabilities
- Git repository structure and access

External Dependencies:

- Homebrew tap creation access
- GitHub releases API access

## Solution Approach

This epic follows an accelerated value-first approach, merging agent creation with CLI publication for maximum early impact. The solution implements a monolithic CLI architecture with embedded agents, enabling immediate manual usage while establishing the technical foundation for future automation. All components operate locally without requiring external APIs or services, supporting the local-first architectural approach.

## Acceptance Criteria

1. 5 core agents defined and manually usable in IDEs
   - Validation: Test agents work in Cursor (@agent) and Claude Code (/agent)
   - Command: `krci-ai validate --agents --ide=cursor,claude`

2. CLI installation works across platforms via package managers
   - Validation: Install succeeds on macOS, Linux, Windows
   - Command: `brew install kuberocketai/tap/krci-ai && krci-ai --version`

3. CI pipeline produces cross-platform releases automatically
   - Validation: GitHub releases contain binaries for all platforms
   - Command: `gh release list --repo kuberocketai/krci-ai`

4. Framework foundation enables manual AI-code fix reduction (addresses BR4, NFR1)
   - Validation: Emily can use agents without platform switching in air-gapped environments
   - Command: `krci-ai validate --offline --user-scenario=emily`

## User Stories

**Story 01.01**: As Emily, I want to access pre-defined AI agents for my development team so that we can standardize our AI-assisted workflows across projects.

**Story 01.02**: As a developer, I want to install the KubeRocketAI CLI via Homebrew so that I can quickly set up the framework in my development environment.

**Story 01.03**: As Emily, I want agents to work directly in Cursor and Claude Code so that my team doesn't need to switch between different AI platforms.

**Story 01.04**: As a development team, we want to use KubeRocketAI to develop itself so that we validate the framework's effectiveness with real-world usage.

**Story 01.05**: As an enterprise developer, I want cross-platform CLI support so that I can use the framework on Windows, macOS, and Linux environments.

**Story 01.06**: As a Framework User, I want a basic `krci-ai` CLI tool with core commands published via Homebrew that includes embedded asset management, so that I can easily install and use the KubeRocketAI framework with standard CLI commands while the full feature set is being developed.

**Story 01.07**: As a KubeRocketAI Development Team, I want to establish a dogfooding environment where we use our own agents and framework to develop KubeRocketAI itself, so that we validate the framework's effectiveness through real-world usage and identify improvement opportunities early.

**Story 01.08**: As a framework adopter, I want essential framework introduction and asset creation guidance, so that I can quickly get started with the KubeRocketAI framework and begin adding my own assets manually.

**Story 01.09**: As a Framework Developer, I want the krci-ai tool to support the new mandatory `customization` field in agent schemas with validation and embedded asset updates, so that the new version enforces schema compliance while developers can manually migrate existing agents to the enhanced schema.

## Epic Summary

**Status**: Complete - Foundation established for KubeRocketAI framework with 5 core agents and CLI tool successfully delivered. All objectives achieved within planned timeline.

---

*This epic is based on Phase 1 of the Roadmap (docs/prd/roadmap.md) and represents the accelerated value-first approach for maximum early impact.*
