# Epic 8: Selective Installation

## Status

| Field      | Value                |
|------------|----------------------|
| Status     | Planning             |
| Priority   | P0 (Critical)        |
| Epic Owner | Development Team     |
| Timeline   | 1 week               |

## Overview

### Problem Statement

Users must install the entire KubeRocketAI framework even when they only need specific agents for their workflow, creating unnecessary bloat and complexity that delays adoption for 85% of target users. Enterprise development teams require granular control over installed components for security compliance and resource optimization, but current all-or-nothing installation approach forces teams to either accept complete framework overhead or manually manage partial installations.

### Goal

Enable selective agent installation for 90% of target users within 30 seconds of decision-making, reducing initial setup time from 2-3 minutes to under 30 seconds while maintaining 100% compatibility with existing validation, bundling, and IDE integration workflows.

### Target Users

Primary: Enterprise Development Leads (60%) - managing team-specific agent configurations across multiple projects
Secondary: Focused Developers (30%) - implementing single-agent workflows for specialized tasks
Tertiary: System Administrators (10%) - deploying controlled agent configurations in enterprise environments

## Scope

### What's Included

1. Single agent installation with dependency resolution (`krci-ai install --agent developer`) (BR13)
2. Multi-agent installation using comma-separated lists (`krci-ai install --agents pm,architect,developer`) (BR14)
3. Task-specific installation with agent context (`krci-ai install --agent pm --task create-prd,update-prd`) (BR15)
4. IDE integration for selective installations (`krci-ai install --agent developer --ide cursor`) (BR16)

### What's Not Included

1. Dynamic agent discovery from remote repositories (deferred to Epic 9)
2. Custom agent creation during installation (out of MVP scope)
3. Selective uninstallation or update capabilities (future roadmap item)

### Dependencies

Epic Dependencies:

- Epic 1: KubeRocketAI Baseline (core CLI infrastructure and embedded asset system)
- Epic 2: Core Engine (dependency resolution and validation capabilities)
- Epic 5: Bundle Management (flag patterns and dependency tracking logic)

System Dependencies:

- Golang 1.24+ runtime environment for CLI compilation
- Embedded asset management system for offline installation
- File system write permissions for `.krci-ai/` directory structure

External Dependencies:

- No network dependencies (maintains offline capability)
- No external service integrations required
- Compatible with existing enterprise security policies

## Solution Approach

Implementation Strategy:

1. Reuse existing bundle command dependency resolution engine for consistency and reliability
2. Extend CLI argument parsing to support identical flag patterns as bundle commands
3. Implement selective asset extraction from embedded framework resources
4. Maintain identical directory structure patterns for full compatibility

Technical Approach:

- CLI Extension: Enhance `cmd/krci-ai/cmd/install.go` with `--agent`, `--task` flags
- Dependency Engine: Leverage existing bundle dependency resolution for asset selection
- Asset Management: Use embedded resource system for offline selective installation
- Validation: Ensure selective installations pass all existing validation and bundling tests

## Acceptance Criteria

1. Single agent installation completes successfully with dependency resolution (BR13)
   - Validation: `krci-ai install --agent developer` creates only developer agent and required dependencies
   - Command: `test -f .krci-ai/agents/developer.md && test $(ls .krci-ai/agents/ | wc -l) -eq 1`

2. Multi-agent installation using comma-separated syntax matches bundle patterns (BR14)
   - Validation: `krci-ai install --agents pm,architect,developer` installs exactly 3 specified agents
   - Command: `krci-ai install --agents pm,architect,developer && test $(ls .krci-ai/agents/ | wc -l) -eq 3`

3. Task-specific installation filters correctly without mixing multiple agents (BR15)
   - Validation: `krci-ai install --agent pm --task create-prd,update-prd` installs only specified tasks
   - Command: `test -f .krci-ai/tasks/create-prd.md && test -f .krci-ai/tasks/update-prd.md`

4. IDE integration works seamlessly with selective installations (BR16)
   - Validation: Selective installation with IDE flag configures only installed agents
   - Command: `krci-ai install --agent developer --ide cursor && cursor --list-agents | grep -c developer`

## User Stories

Planned Stories for Implementation:

### Phase 1: Complete Selective Installation (Sprint 4)

- Story 8.01: Single and Multi-Agent Installation
  - As an Enterprise Development Lead, I want to install specific agents using `krci-ai install --agent developer` or `krci-ai install --agents pm,architect,developer`
  - Acceptance: Installation completes successfully with only specified agents and dependencies, supporting both single and multiple agent selection
  - Dependencies: Epic 1 baseline infrastructure and Epic 2 validation engine

- Story 8.02: Task-Specific Installation
  - As a Focused Developer, I want to install specific tasks using `krci-ai install --agent pm --task create-prd,update-prd`
  - Acceptance: Only specified tasks and their dependencies installed correctly
  - Dependencies: Story 8.01 completion and task dependency mapping

- Story 8.03: IDE Integration for Selective Installation
  - As an Enterprise Development Lead, I want selective installations to work with IDE integration
  - Acceptance: `krci-ai install --agent developer --ide cursor` configures IDE correctly
  - Dependencies: Epic 4 IDE integration and Story 8.01 completion
