# Epic 3: Install Command and Update Management (Week 4)

## Status

| Field                | Value                    |
|----------------------|--------------------------|
| Status               | Completed                |
| Priority             | P1 (High)               |
| Epic Owner           | Development Team        |
| Timeline             | Week 4 (1 week) - DONE |

## Overview

### Problem Statement

Development teams face significant friction when adopting new AI frameworks due to complex installation processes, dependency management, and version compatibility issues. Setup overhead and context switching between installation tools disrupts development flow, while lack of standardized update management creates configuration drift across team members. Teams in air-gapped or security-sensitive environments require offline installation capabilities that current AI tools cannot provide.

### Goal

Enable offline framework installation for 100% of development teams within 30 seconds, delivering single-command setup (`krci-ai install`) that works in air-gapped environments while providing automated update management with version mismatch detection and guided package manager instructions.

### Target Users

**Primary User**: Emily - Enterprise Development Lead needing to standardize framework adoption across 8-15 developers with consistent installation and update processes.

**User Context**: 500+ Active Development Leads globally, including teams in air-gapped environments requiring offline installation capabilities and enterprise teams needing guided update management.

## Scope

### What's Included

**PRD Requirements Addressed:**

- **BR1 [P0]**: User can install KubeRocketAI framework with single command (krci-ai install) that extracts embedded Agent Playbook locally
- **NFR2 [P0]**: System supports offline installation without network dependencies, enabling air-gapped development environments
- **NFR4 [P1]**: Framework handles bundle-based update management with version mismatch detection and graceful degradation

**Core Deliverables:**

- `krci-ai install` command for offline framework installation
- Local playbook directory management and organization
- `krci-ai check-updates` command for version detection
- `krci-ai update` command for guided update process
- Cross-platform compatibility for all supported environments

**Note:**

- Backup and restore functionality is **not included** in this epic. For backup and possible reverts, users are advised to use git or their preferred version control system.

### What's Not Included

- Agent creation functionality (Epic 1)
- Framework validation capabilities (Epic 2)
- IDE-specific configuration generation (Epic 4)
- Remote repository fetching (Post-MVP feature)
- Community marketplace integration (Post-MVP)
- Backup and restore system (see note above)

### Dependencies

Epic Dependencies:

- Epic 2: Core engine and validation system (asset processing, validation commands)

System Dependencies:

- Go embed functionality for asset extraction
- File system operations and local asset validation
- Cross-platform file path management
- Package manager integration capabilities (brew, chocolatey)

External Dependencies:

- GitHub releases API for update detection
- Package manager distribution channels (Homebrew, Chocolatey)
- Network access for update checking (optional, graceful degradation)

## Solution Approach

This epic implements a bundle-based distribution model where CLI and framework assets are versioned together and distributed via package managers. The solution uses Go's embed functionality for offline asset extraction while providing optional online features for update detection. The system maintains local-first architecture with graceful online degradation, ensuring core operations work offline while enabling intelligent update management when network access is available.

## Acceptance Criteria

1. Install command successfully extracts and installs embedded framework assets
   - Validation: Framework assets extracted to correct local directory structure
   - Command: `krci-ai install && find .krci-ai -type f | wc -l`

2. Installation works completely offline without network dependencies
   - Validation: Install succeeds in air-gapped environment without internet access
   - Command: `sudo iptables -A OUTPUT -j DROP && krci-ai install --offline`

3. Update detection identifies available CLI versions via GitHub releases
   - Validation: Check-updates command detects newer versions when available
   - Command: `krci-ai check-updates --verbose`

4. Guided update process provides clear package manager instructions
   - Validation: Update command shows platform-specific upgrade instructions
   - Command: `krci-ai update --dry-run --show-commands`

5. Version mismatch detection between CLI and local framework assets
   - Validation: System detects and reports CLI vs framework version mismatches
   - Command: `krci-ai validate --version-check --verbose`

6. Single-command framework setup for Emily's teams (addresses BR1, NFR2, NFR4)
   - Validation: 30-second installation in air-gapped environments with update management
   - Command: `time krci-ai install --complete --validate-offline`

- [x] Streamlined setup contributes to 15-minute user experience target

## User Stories

**Story 03.01**: As Emily, I want to install the KubeRocketAI framework with a single command so that my team can quickly set up standardized AI agent configurations.

**Story 03.02**: As Emily, I want automatic update detection so that my team stays current with framework improvements without manual version checking.

## Implementation Status

**Completed** (Week 4 - 100% Complete): Offline installation and bundle-based update management

- Asset extraction with Go embed and local directory management
- GitHub releases API integration with graceful offline degradation
- Bundle-based distribution: CLI + framework assets versioned together
- Risk mitigation: Asset validation, version mismatch detection
- **Backup/restore functionality is not included; users should use git for backup and reverts.**

---

*This epic is based on Phase 3 of the Roadmap (docs/prd/roadmap.md) and enables the framework distribution and update management model.*
