# Epic 8: Selective Installation

## Status

| Field      | Value                |
|------------|----------------------|
| Status     | Done                 |
| Priority   | P0 (Critical)        |
| Epic Owner | Development Team     |
| Timeline   | 1 week               |

## Overview

### Problem Statement

Users must install the entire KubeRocketAI framework even when they only need specific agents for their workflow, creating unnecessary bloat and complexity that delays adoption for 85% of target users. Enterprise development teams require granular control over installed components for security compliance and resource optimization, but current all-or-nothing installation approach forces teams to either accept complete framework overhead or manually manage partial installations.

### Goal

Deliver selective agent installation capability enabling users to install specific agents without framework bloat while maintaining full compatibility with existing validation, bundling, and IDE integration workflows.

### Target Users

Primary: Enterprise Development Leads (60%) - managing team-specific agent configurations across multiple projects
Secondary: Focused Developers (30%) - implementing single-agent workflows for specialized tasks
Tertiary: System Administrators (10%) - deploying controlled agent configurations in enterprise environments

## Scope

### What's Included

1. Single agent installation with dependency resolution (`krci-ai install --agent developer`) (BR13)
2. Multi-agent installation using comma-separated lists (`krci-ai install --agent pm,architect,developer`) (BR14)
3. IDE integration for selective installations (`krci-ai install --agent developer --ide cursor`) (BR15)

### What's Not Included

1. Task-specific installation filtering (removed from scope)
2. Dynamic agent discovery from remote repositories (deferred to Epic 9)
3. Custom agent creation during installation (out of MVP scope)
4. Selective uninstallation or update capabilities (future roadmap item)

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

- CLI Extension: Enhance `cmd/krci-ai/cmd/install.go` with `--agent` flag
- Dependency Engine: Leverage existing bundle dependency resolution for asset selection
- Asset Management: Use embedded resource system for offline selective installation
- Validation: Ensure selective installations pass all existing validation and bundling tests

## Acceptance Criteria

1. Single agent installation with dependency resolution
   - Scenario: User installs specific agent using selective installation command
   - Expected Behavior: Only requested agent and its required dependencies are installed in .krci-ai directory structure
   - Measurement/Method: Verify installed agents match request exactly; check dependency files exist; confirm no extra agents installed
   - Preconditions/Assumptions: CLI available; target directory writable; framework assets embedded
   - Guardrails: Installation completes without errors; no partial states; maintains directory structure consistency
   - Traceability: BR13

2. Multi-agent installation using comma-separated syntax
   - Scenario: User installs multiple specific agents in single command using comma-separated list
   - Expected Behavior: Exactly specified agents and their dependencies installed; no additional agents added
   - Measurement/Method: Count installed agents matches request; verify each requested agent present; confirm dependency resolution
   - Preconditions/Assumptions: Multiple valid agent names provided; sufficient disk space; no conflicting installations
   - Guardrails: All requested agents installed successfully or rollback on failure; no orphaned dependencies
   - Traceability: BR14

3. IDE integration compatibility
   - Scenario: User performs selective installation with IDE integration flag to configure development environment
   - Expected Behavior: IDE configuration reflects only installed agents; integration works seamlessly; no missing agent references
   - Measurement/Method: Check IDE shows correct agent list; verify agent activation works; confirm no error messages
   - Preconditions/Assumptions: Supported IDE installed; integration feature available; valid agent-IDE combinations
   - Guardrails: IDE remains stable; agent functionality preserved; configuration reversible
   - Traceability: BR15

## User Stories

Implementation Status:

### Phase 1: Selective Installation Implementation (Sprint 4)

- **Story 8.01:** ✅ **COMPLETED** - Selective Agent Installation Core
  - As an Enterprise Development Lead, I want to install specific agents individually or in groups, so that I can customize my team's framework setup without unnecessary bloat
  - Acceptance: Single agent installation (`--agent developer`) and multi-agent installation (`--agent pm,architect,developer`) work correctly with dependency resolution and IDE integration
  - Value: Reduces setup time and resource usage for teams needing specific agent workflows
  - Dependencies: Epic 1 baseline infrastructure and Epic 2 validation engine
