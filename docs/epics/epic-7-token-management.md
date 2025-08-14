# Epic 7: Token Management

## Status

| Field                | Value                    |
|----------------------|--------------------------|
| Status               | Planning                 |
| Priority             | Critical                 |
| Epic Owner           | Product Owner            |
| Timeline             | 1 week                   |

## Overview

### Problem Statement

Users struggle with understanding AI platform context limits and token consumption patterns when creating agent bundles and configurations. Current agent composition lacks visibility into token usage, leading to runtime failures when bundles exceed AI platform context limits (GPT-4: 128k, Claude: 200k, Gemini: 2M tokens). Teams report 20% of agent deployments warning due to unexpected token limit violations, requiring manual trial-and-error optimization that consumes 2-3 hours per configuration cycle.

### Goal

Enable transparent token analysis and optimization for 95% of users within 30 seconds of bundle creation, reducing token-related deployment failures from 20% to under 5% while providing token feedback on early agent design phase.

### Target Users

Primary: Agile Team Members (100%) - managing agent configurations and bundle deployments across multiple AI platforms, designing agent compositions and dependency structures for optimal token efficiency

## Scope

### What's Included

1. Token calculation engine with platform-specific tokenization algorithms (BR10, NFR7, NFR8, NFR9)
2. Real-time token analysis during validation and bundling workflows (BR11) (Out of MVP scope)
3. Detailed dependency breakdown showing agent, task, template, and data contributions (BR12)
4. Context limit warnings and optimization recommendations for major AI platforms (BR11) (Out of MVP scope)
5. CLI commands for token analysis with granular scope options (`--task`, `--agent`, `--all`) (BR10)
6. Integration with existing validation and bundling commands for seamless token awareness (BR11)

### What's Not Included

1. Dynamic token optimization or automatic bundle reduction (future enhancement)
2. Real-time token monitoring during AI platform execution (out of scope)
3. Custom tokenization algorithms beyond major platforms (GPT, Claude, Gemini)
4. Historical token usage analytics or trend analysis

### Dependencies

Epic Dependencies:

- Epic 2: Core Engine (validation framework and dependency tracking)
- Epic 5: Bundle Management (bundle creation and dependency resolution)

System Dependencies:

- Golang tiktoken library for OpenAI tokenization
- Platform-specific tokenization libraries (anthropic, google-generativeai)
- Existing dependency tracking mechanisms from validation engine

External Dependencies:

- AI platform API documentation for context limit specifications
- Tokenization library updates and compatibility maintenance

## Solution Approach

Implementation Strategy:

1. Modular tokenization framework with pluggable platform adapters for extensibility
2. Dependency-aware calculation leveraging existing asset resolution from validation engine
3. Progressive analysis with caching for performance optimization during repeated operations
4. Integration hooks in validation and bundling workflows for transparent token reporting

Technical Approach:

- Use [library](https://github.com/tiktoken-go/tokenizer) or any other Golang-based tokenization library
- CLI: Extended command interface with token analysis options and detailed reporting
- Integration: Token awareness injected into existing validation and bundling workflows

## Acceptance Criteria

1. Token calculation completes within 3 seconds for typical project configurations with 20+ agents
   - Validation: Performance testing with realistic project sizes
   - Command: `time krci-ai tokens --all`

2. Token size estimation accuracy remains within 10% variance of actual AI platform consumption
   - Validation: Cross-platform testing against actual AI platform token counts
   - Command: `pytest tests/token/accuracy_test.py`

3. CLI provides granular token analysis with task, agent, and project scope options
   - Validation: All CLI options return appropriate token breakdowns
   - Command: `krci-ai tokens --agent pm && krci-ai tokens --all`

4. Validation and bundling commands include automatic token size reporting and warnings 
   - Validation: Token information appears in command output with context limit warnings `krci-ai validate | grep -E "Token|Warning"` (Out of MVP scope)
   - Bundle: `krci-ai bundle --all | grep -E "Token|Warning"`

5. Detailed dependency breakdown shows contribution from agents, tasks, templates, and data
   - Validation: Token breakdown identifies specific asset contributions
   - Command: `krci-ai tokens --all`

6. Platform-specific context limit warnings for GPT-4, Claude, and Gemini Pro (Calculate for GPT-4, other platforms out of MVP scope)
   - Validation: Warnings appear when approaching platform-specific limits
   - Command: `krci-ai tokens --model gpt4 --warn-threshold 0.8`

## User Stories

Planned Stories for Implementation:

### Phase 1: Foundation (Sprint 4)

- Story 7.01: Core Token Calculation Engine with CLI
  - As a Developer, I want to analyze token usage for agent configurations using CLI commands
  - Acceptance: Token calculation works for individual agents within 1 second with CLI support for GPT-4 model (`krci-ai tokens --agent pm`, `krci-ai tokens --all`)
  - Dependencies: Epic 2 dependency tracking infrastructure

- Story 7.02: Platform-Specific Tokenization (Multi-Platform) (Out of MVP scope)
  - As a Software Architect, I want accurate token estimates for Claude and Gemini platforms beyond GPT-4
  - Acceptance: Token calculations match actual platform consumption within 10% variance for all major platforms
  - Dependencies: Story 7.01 completion

### Phase 2: Bundle Integration (Sprint 4)

- Story 7.03: Bundle Integration and Optimization
  - As a Developer, I want token analysis integrated with bundle creation
  - Acceptance: Bundle commands include token reporting and context limit warnings(Out of MVP scope)
  - Dependencies: Epic 5 bundle management, Story 7.01

- Story 7.04: Validation Workflow Integration
  - As a DevOps Engineer, I want token information included in validation workflows
  - Acceptance: Validation commands automatically report token usage and warnings
  - Dependencies: Epic 2 validation engine, Story 7.01

  ### Phase 3: Advanced Analytics (Sprint 5)

- Story 7.05: Dependency Breakdown Analysis
  - As a Software Architect, I want detailed token breakdown by asset type and dependency
  - Acceptance: Token breakdown shows agent, task, template, and data contributions
  - Dependencies: Story 7.01 completion
