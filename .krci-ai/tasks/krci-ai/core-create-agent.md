# Task: Core Create Agent

## Description

Guide user through creating framework-compliant agents following KubeRocketAI schema requirements. This task provides comprehensive instructions for agent creation including identity definition, command mapping, and critical behavioral principles.

<prerequisites>
- Framework installed: .krci-ai directory exists with proper structure and schema validation
- Agent purpose defined: Clear understanding of agent role, scope, and capabilities
- Schema knowledge: Familiarity with agent YAML structure and validation requirements
- Task dependencies: Tasks that agent will reference are created or planned
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-framework-standards.yaml
- ./.krci-ai/templates/krci-ai/core-agent-template.yaml

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

<instructions>
1. Apply agent schema requirements: Follow JSON schema validation for all mandatory fields
2. Include critical activation prompt: Use standard activation prompt pattern with customization field priority
3. Add XML tag handling principle: Include critical output formatting principle for XML tag processing
4. Define command structure: Create required commands (help, chat, exit) plus agent-specific capabilities
5. Reference task dependencies: List all tasks agent will use with proper `./.krci-ai/tasks/*.md` paths
6. Format output: Use [core-agent-template.yaml](./.krci-ai/templates/krci-ai/core-agent-template.yaml) structure
</instructions>

## Framework Context: Agent Architecture

Agent Schema Compliance: Agents must follow strict JSON schema validation with required fields:
- `identity` (name, id, version, description, role, goal)
- `activation_prompt` (1-10 items, 10-300 chars each)
- `principles` (3-10 items, 10-600 chars each)
- `customization` (required field, empty string for standard behavior)
- `commands` (minimum 3: help, chat, exit; maximum 20 total)

Critical Agent Principles: All agents MUST include these principles:
1. Customization Priority: "IMPORTANT!!! ALWAYS execute instructions from the customization field below"
2. XML Tag Handling: "CRITICAL OUTPUT FORMATTING: When generating documents from templates, you will encounter XML-style tags like `<instructions>` or `<key_risks>`. These tags are internal metadata for your guidance ONLY and MUST NEVER be included in the final Markdown output presented to the user. Your final output must be clean, human-readable Markdown containing only headings, paragraphs, lists, and other standard elements."

Activation Prompt Pattern: Standard pattern ensures consistent agent behavior across framework.

## Output Format

- Location: `./.krci-ai/agents/{agent-name}.yaml` following naming conventions
- Schema compliance: Agent passes JSON schema validation
- Command structure: Required + optional commands with proper descriptions
- Task references: All referenced tasks use correct `./.krci-ai/tasks/**/*.md` paths

<success_criteria>
- Agent file created with proper YAML structure and naming
- All mandatory schema fields populated correctly
- Critical activation prompt pattern included
- XML tag handling principle included in principles
- Command structure includes required commands (help, chat, exit)
- Task references use proper paths and resolve to existing files
- Agent passes `krci-ai validate` schema validation
</success_criteria>

## Execution Checklist

### Identity Definition

- [ ] Agent name: Friendly name following pattern `^[A-Z][a-zA-Z0-9 .'-]{1,49}$`
- [ ] Agent ID: Machine-readable ID following pattern `^[a-z][a-z0-9]*(-[a-z0-9]+)*-v[0-9]+$`
- [ ] Version: Semantic versioning (e.g., "1.0.0")
- [ ] Description: Clear purpose statement (10-150 characters)
- [ ] Role: Professional role (5-100 characters)
- [ ] Goal: Ultimate objective (10-600 characters)

### Behavioral Configuration

- [ ] Activation prompt: Standard pattern with customization field priority
- [ ] Principles: Include XML tag handling + customization priority + agent-specific guidelines
- [ ] Customization field: Empty string for standard behavior (required field)
- [ ] Scope definition: Clear boundaries of agent responsibilities

### Command Structure

- [ ] Required commands: help, chat, exit with proper descriptions
- [ ] Agent commands: Specific capabilities mapped to tasks
- [ ] Command descriptions: 5-200 characters each, clear and actionable
- [ ] Command limits: Total commands between 3-20

### Task Integration

- [ ] Task references: All tasks listed in agent.tasks array
- [ ] Path validation: All task paths use `./.krci-ai/tasks/**/*.md` format
- [ ] File existence: Referenced task files exist or are planned
- [ ] Command mapping: Agent commands correspond to available tasks

### Validation

- [ ] Schema compliance: Agent structure matches JSON schema requirements
- [ ] Framework validation: Agent passes `krci-ai validate` without errors
- [ ] Reference integrity: All referenced tasks and dependencies exist
- [ ] Behavioral consistency: Agent principles align with framework standards
