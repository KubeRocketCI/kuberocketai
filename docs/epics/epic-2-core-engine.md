# Epic 2: Core Engine (Week 2-3)

## Status

| Field                | Value                    |
|----------------------|--------------------------|
| Status               | Completed                |
| Priority             | P0 (Critical)            |
| Epic Owner           | Development Team         |
| Timeline             | Week 2-3 (2 weeks) - DONE |

## Overview

### Problem Statement

Development teams using AI assistance tools struggle with unreliable agent configurations that generate code misaligned with project standards. Without validation capabilities, developers spend 5-10 minutes daily manually fixing AI-generated code due to inconsistent agent definitions, missing project context, and lack of error detection. Additionally, teams need comprehensive link validation to ensure all internal framework references are resolvable and dependency relationships are clear for effective framework management. Teams need real-time validation feedback to ensure agent configurations are correct before deployment, preventing downstream manual fixes and improving development velocity.

### Goal

Enable reliable asset processing and validation for 100% of KubeRocketAI framework components within 2 seconds, delivering validation engine that reduces manual AI-code fixes by 85% through real-time error detection, design-time configuration validation, and comprehensive internal link resolution with dependency analysis.

### Target Users

**Primary User**: Emily - Enterprise Development Lead requiring reliable agent configurations across 3-4 microservices with consistent validation standards and clear dependency visibility.

**User Context**: 500+ Active Development Leads globally who need real-time feedback on agent configuration quality and framework link integrity to prevent manual fixes and ensure team productivity.

## Scope

### What's Included

**PRD Requirements Addressed:**

- **BR4 [P0]**: User can validate agent configurations and templates using built-in validation engine with design-time and runtime checks
- **NFR1 [P0]**: Framework validates agent configurations in under 2 seconds for real-time feedback during development

**Core Deliverables:**

- Markdown + YAML frontmatter parser for framework assets
- Schema validation system for agents, tasks, templates, data
- CLI validation command (`krci-ai validate`) with clear error reporting
- Two-tier validation architecture (design-time and runtime)
- Performance optimization for large frameworks
- **NEW**: Internal link resolution validator for comprehensive `./.krci-ai/` reference validation
- **NEW**: Dependency graph generation and visualization capabilities
- **NEW**: Agent|Tasks|Templates|Data relationship table generation
- **NEW**: Enhanced CLI commands for dependency analysis and link validation

### What's Not Included

- Framework installation functionality (Epic 3)
- IDE-specific configuration generation (Epic 4)
- User interface beyond CLI commands
- Remote validation services (local-first approach)

### Dependencies

Epic Dependencies:

- Epic 1: KubeRocketAI Baseline (Go module, CLI framework, agent schemas)

System Dependencies:

- Go development environment (1.19+)
- YAML parsing libraries (gopkg.in/yaml.v3)
- JSON Schema validation libraries
- File system operations capabilities

External Dependencies:

- JSON Schema specification compliance
- Markdown parsing standards compatibility

## Solution Approach

This epic implements a two-tier validation architecture with static validation through CLI and runtime validation for critical issues only. The solution uses Go's robust YAML parsing with JSON Schema validation, providing comprehensive error reporting with actionable guidance. Performance optimization ensures validation completes in under 1 second (exceeding NFR2 requirement of <2 seconds) while maintaining accuracy and user experience quality.

## Acceptance Criteria

1. Framework validation prevents runtime errors and broken agent configurations for development teams
   - Validation: Zero agent configuration failures occur in production deployments after validation passes
   - Command: `krci-ai validate`

2. Link integrity validation eliminates broken references within framework components
   - Validation: All internal framework links resolve successfully preventing reference failures
   - Command: `krci-ai validate`

3. Architecture compliance validates component separation per framework constraints C1-C7
   - Validation: Agent defines behavior, Task defines workflow, Templates referenced only in Tasks, Data referenced by both
   - Command: `krci-ai validate`

4. Framework validation completes within 2 seconds enabling real-time development feedback (exceeds NFR2)
   - Validation: Validation response time under 2 seconds for frameworks up to 100 components
   - Command: `time krci-ai validate`

5. Development teams receive clear, actionable guidance for resolving framework issues
   - Validation: Error messages provide specific fix instructions reducing resolution time by 80%
   - Command: `krci-ai validate`

6. Framework component relationships are visible and analyzable for development planning
   - Validation: Component dependency visualization shows complete Agent-Task-Template-Data relationships
   - Command: `krci-ai validate`

## User Stories

**Story 02.01**: As Emily, I want to validate agent configurations before deployment so that my team doesn't experience runtime errors or misaligned code generation.

**Story 02.02**: As a developer, I want the validation command to complete in under 2 seconds so that I can get immediate feedback during development without workflow disruption.

**Story 02.03**: As Emily, I want clear error messages with actionable guidance so that my team can quickly resolve configuration issues without extensive debugging.

**Story 02.04**: As Emily, I want comprehensive internal link validation and architecture constraint enforcement (C1-C7) so that my team can ensure framework integrity, understand component relationships, and prevent constraint violations before deployment.

## Implementation Status

**In Progress**: Core validation engine development (Week 2-3)

- YAML schema validation and error reporting system
- Template/content validation with performance optimization
- Two-tier validation architecture (static + runtime)
- Risk mitigation: Performance caching, iterative complexity management

---

*This epic is based on Phase 2 of the Roadmap (docs/prd/roadmap.md) and establishes the technical foundation for framework reliability.*
