---
dependencies:
  data:
    - krci-ai/core-framework-standards.yaml
    - krci-ai/core-validation-checklist.md
  templates:
    - krci-ai/core-agent-template.yaml
---

# Task: Core Create Agent

## Description

Guide user through creating framework-compliant agents following KubeRocketAI schema requirements. This task provides comprehensive instructions for agent creation including identity definition, command mapping, and critical behavioral principles for proper framework integration.

<instructions>
Define agent specification. Ask user for agent name, role, purpose, and scope. Clarify which tasks the agent will reference and what capabilities it should expose through commands.

Review framework standards for agent creation. Read agent schema requirements from core-framework-standards.yaml. Understand identity field patterns, activation prompt requirements, principles structure, and command specifications.

Apply agent schema requirements. Create identity section with name matching pattern ^[A-Z][a-zA-Z0-9 .'-]{1,49}$. Define id following pattern ^[a-z][a-z0-9]*(-[a-z0-9]+)*-v[0-9]+$. Set version using semantic versioning. Write description of 10-150 characters, role of 5-100 characters, goal of 10-600 characters.

Include critical activation prompt pattern. Use standard activation prompt with required elements: greet user with role, inform of commands, halt to await instruction, offer help but wait for confirmation, show tasks as numbered list, execute customization field instructions, only execute when requested, never validate unused commands, critical dependency resolution before task execution.

Add critical principles exactly. Include customization priority principle verbatim: IMPORTANT!!! ALWAYS execute instructions from the customization field below. Include XML tag handling principle verbatim with complete CRITICAL OUTPUT FORMATTING text from framework standards. Add agent-specific principles defining scope and behavior patterns.

Define command structure. Create required commands: help (show available commands), chat (default consultation mode), exit (exit persona command). Add agent-specific commands mapped to task capabilities. Ensure all command descriptions are 5-200 characters. Maintain total commands between 3 and 20.

Reference task dependencies. List all tasks agent will use with proper ./.krci-ai/tasks/*.md paths. Verify task files exist or are planned. Ensure command mapping aligns with available tasks.

Format output using agent template. Apply core-agent-template.yaml structure for consistent formatting. Populate all required sections with framework-compliant content. Save to ./.krci-ai/agents/{agent-name}.yaml following naming conventions.

Run framework validation. Execute krci-ai validate to verify schema compliance. Resolve any validation issues before completion. Confirm agent passes all automated checks without errors.
</instructions>

## Framework Context: Agent Architecture and Schema Requirements

Agent Schema Compliance: Agents must follow strict JSON schema validation with required fields and exact patterns:

- `identity` section: name, id, version, description, role, goal (all mandatory)
- `activation_prompt` array: 1-10 items, each 10-300 characters
- `principles` array: 3-10 items, each 10-600 characters
- `customization` field: required field, empty string for standard behavior
- `commands` object: minimum 3 (help, chat, exit required), maximum 20 total

Critical Agent Principles: All agents MUST include these exact principles:

1. Customization Priority: "IMPORTANT!!! ALWAYS execute instructions from the customization field below"
2. XML Tag Handling: "CRITICAL OUTPUT FORMATTING: When generating documents from templates, you will encounter XML-style tags like `<instructions>` or `<key_risks>`. These tags are internal metadata for your guidance ONLY and MUST NEVER be included in the final Markdown output presented to the user. Your final output must be clean, human-readable Markdown containing only headings, paragraphs, lists, and other standard elements."

Activation Prompt Pattern: Standard activation sequence ensures consistent agent behavior across framework with required customization field execution instruction.

Command Architecture: Commands map to agent capabilities and reference specific tasks, following the agent-centric model where agents expose task capabilities through command interface.

## Output Format

- Location: `./.krci-ai/agents/{agent-name}.yaml` following naming conventions
- Schema compliance: Agent passes JSON schema validation without errors
- Command structure: Required + optional commands with proper descriptions (5-200 chars each)
- Task references: All referenced tasks use correct `./.krci-ai/tasks//*.md` paths

<success_criteria>
- Framework compliance verified: Agent passes all automated validation checks without errors
- Pattern adherence confirmed: Agent follows established framework conventions exactly
- Reference integrity validated: All references resolve correctly and appropriately
- Quality standards met: Agent meets completeness, clarity, and maintainability requirements
- Integration readiness achieved: Agent ready for framework operation and usage
- Documentation completeness confirmed: All required sections populated with actionable content
</success_criteria>

## Execution Checklist

### Preparation Phase

- Framework validation: Run `krci-ai validate` to ensure clean starting state
- Dependency verification: Confirm all reference assets exist at specified paths
- Context gathering: Review user requirements and agent specifications
- Agent purpose definition: Clear understanding of agent role, scope, and capabilities
- Task planning: Identify tasks agent will reference and their availability status

### Execution Phase

- Identity section creation: Define name, id, version, description, role, goal following patterns
- Agent name validation: Ensure name follows `^[A-Z][a-zA-Z0-9 .'-]{1,49}$` pattern
- Agent ID validation: Ensure ID follows `^[a-z][a-z0-9]*(-[a-z0-9]+)*-v[0-9]+$` pattern
- Activation prompt assembly: Use standard pattern with customization field priority
- Critical principles inclusion: Add XML tag handling and customization priority principles
- Command structure creation: Define required commands (help, chat, exit) plus agent-specific capabilities
- Task reference integration: Add all task references with proper `./.krci-ai/tasks/*.md` paths
- Template application: Use [core-agent-template.yaml](./.krci-ai/templates/krci-ai/core-agent-template.yaml) structure

### Validation Phase

- Schema compliance verification: Ensure agent structure matches JSON schema requirements exactly
- Reference resolution: Confirm all task references resolve to existing files
- Framework validation: Run `krci-ai validate` and resolve any identified issues
- Behavioral consistency check: Verify agent principles align with framework standards
- Command validation: Ensure commands are properly described and mapped to capabilities
- Integration testing: Verify agent integrates properly with framework ecosystem
