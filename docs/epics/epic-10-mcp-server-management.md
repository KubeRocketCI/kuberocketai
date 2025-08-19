# Epic 10: MCP Server Management

## Status

| Field                | Value                    |
|----------------------|--------------------------|
| Status               | Planning                 |
| Priority             | P1 (High)               |
| Epic Owner           | Product Owner           |
| Timeline             | 1 week                  |
| Epic Key/Link        | epic-10-mcp            |

## Overview

### Problem Statement

Framework users currently have no visibility into MCP server dependencies required for agent/task execution, leading to runtime failures and poor user experience when tasks attempt to execute without required infrastructure. Users spend time troubleshooting MCP-related issues that could be prevented through proactive dependency discovery, and teams lack clear understanding of which MCP servers are needed for specific agent/task combinations across their projects.

### Goal

Enable MCP server dependency discovery framework users within 5 seconds of command execution, delivering comprehensive MCP server visibility through enhanced CLI commands that reduce MCP-related task failures and provide clear dependency mapping for agent/task combinations.

### Target Users

**Primary**: Emily - Enterprise Development Lead managing 8-15 developers across 3-4 microservices, requiring clear visibility into MCP dependencies for team onboarding and project setup

**Secondary**: Software Engineers and Architects using framework agents, needing to understand MCP server requirements before task execution

**Tertiary**: DevOps Engineers setting up framework environments, requiring dependency information for infrastructure planning

## Scope

### What's Included

1. MCP server discovery through `krci-ai list mcp` command showing all referenced servers across framework (BR16)
2. Enhanced agent information display via `krci-ai list agents -v` including MCP dependencies for each task (BR17)
3. Integration of MCP server information into existing `krci-ai validate` command output without additional flags (BR18)
4. Simple MCP metadata format in task files enabling runtime agent discovery and dependency tracking (BR19)
5. 2-second performance standard for all MCP discovery operations (NFR10)
6. Backward-compatible implementation maintaining existing workflows (NFR11)

### What's Not Included

1. MCP server connectivity validation or health checking (runtime agent responsibility)
2. Automated MCP server installation or configuration management (future enhancement)
3. MCP server version compatibility checking beyond metadata display (deferred)
4. Real-time MCP server monitoring or dashboard capabilities (out of scope)

### Dependencies

#### Epic Dependencies

- Epic 2: Core Engine (validation infrastructure required)
- Epic 3: Install Command (CLI framework needed for new commands)

#### System Dependencies

- Existing task file parsing infrastructure from core engine
- CLI command framework for `list` and `validate` commands

#### External Dependencies

- None (no external MCP servers required for discovery functionality)

## Solution Approach

Extend existing CLI infrastructure to provide MCP server discovery capabilities through three integration points: dedicated `list mcp` command for framework-wide server inventory, enhanced `list agents -v` verbose mode showing per-task MCP dependencies, and integration with existing `validate` command to include MCP information in standard output.

Implementation leverages existing file parsing infrastructure to extract MCP metadata from task files, with caching for performance optimization. All MCP server references stored as simple metadata in task YAML frontmatter, enabling runtime agent discovery without framework-level validation complexity.

## Risks & Assumptions

### Risks

- Risk: Task file metadata parsing may impact performance if numerous agents/tasks are processed
- Risk: MCP metadata format changes could require migration of existing task files

### Assumptions

- Assumption: Existing CLI command infrastructure can support new MCP discovery commands without architectural changes
- Assumption: Task file YAML frontmatter can accommodate MCP metadata without breaking existing parsing
- Assumption: 2-second performance target achievable with file-based discovery approach

## Acceptance Criteria

1. **MCP server framework discovery**
   - Scenario: User executes `krci-ai list mcp` command to discover all MCP servers referenced across framework
   - Expected Behavior: Command displays complete list of MCP servers with task references and agent mapping within 2 seconds
   - Measurement/Method: Execute command in test environment with 20 agents and 100 tasks; verify complete server list and performance timing
   - Preconditions/Assumptions: Framework installed with test agent configurations; task files contain MCP metadata
   - Guardrails: Command completes within 2 seconds; no missing server references; clear agent-to-server mapping
   - Traceability: BR16, NFR10

2. **Enhanced agent information with MCP dependencies**
   - Scenario: User executes `krci-ai list agents -v` to view detailed agent information including MCP server dependencies
   - Expected Behavior: Verbose output includes MCP server requirements for each task within agent definitions
   - Measurement/Method: Compare verbose output with standard output; verify MCP information presence and accuracy
   - Preconditions/Assumptions: Multiple agents installed with varied MCP dependencies; verbose flag functionality exists
   - Guardrails: All agent MCP dependencies displayed; no performance degradation; maintains existing output format
   - Traceability: BR17, NFR11

3. **Integrated MCP validation reporting**
   - Scenario: User executes `krci-ai validate` command and receives MCP server information as part of standard validation output
   - Expected Behavior: Validation output includes MCP dependency information without requiring additional flags or commands
   - Measurement/Method: Run validation on project with MCP dependencies; verify MCP information included in output
   - Preconditions/Assumptions: Existing validate command infrastructure; task files with MCP metadata available
   - Guardrails: No additional latency in validation; MCP info clearly distinguishable; existing validation flow preserved
   - Traceability: BR18, NFR11

4. **MCP metadata infrastructure**
   - Scenario: Task files include simple MCP server metadata that enables discovery and runtime agent processing
   - Expected Behavior: MCP metadata stored in standardized format that supports both discovery commands and agent runtime needs
   - Measurement/Method: Verify task files parse correctly; confirm metadata accessible to both CLI and agent runtime
   - Preconditions/Assumptions: Task file YAML frontmatter parser available; agent runtime can access metadata
   - Guardrails: Backward compatibility maintained; metadata format consistent; no breaking changes to existing tasks
   - Traceability: BR19, NFR11

## User Stories

### Single Implementation

- Story 10.01: MCP Discovery Implementation
  - As a Framework User, I want to discover MCP server dependencies through CLI commands, so that I can understand infrastructure requirements for my projects
