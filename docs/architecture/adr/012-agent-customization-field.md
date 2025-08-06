# ADR-012: Agent Customization Field Enhancement

## Status

Accepted

## Context

The KubeRocketAI agent framework has established a standardized activation prompt pattern that provides consistent behavior across all agents. However, users have requested the ability to customize the agent bootstrap process for specific use cases while maintaining the core framework structure.

### Current State

All agents follow a critical activation prompt pattern defined in `04_Data_Models.md`:

```yaml
activation_prompt:
  - "Greet the user with your name and role, inform of available commands, then HALT to await instruction"
  - "Offer to help with [domain] tasks but wait for explicit user confirmation"
  - "Only execute tasks when user explicitly requests them"
  - "When loading any asset (task, data, template), always use the project root relative path resolution"
```

### Forces

- **User Flexibility**: Users need ability to customize agent bootstrap behavior for specific contexts
- **Framework Consistency**: Must maintain standardized activation patterns across the framework
- **Backward Compatibility**: Existing agents should continue to work without modification
- **Precedence Control**: Custom instructions should take precedence over standard patterns when present
- **Validation Simplicity**: Schema validation should remain straightforward
- **Clear Semantics**: Clear distinction between standard and custom behavior

## Decision

**Add mandatory `customization` field to agent schema with default empty string value**

### Schema Enhancement

```yaml
agent:
  identity:
    name: string                  # Friendly user name
    icon: string                  # Optional icon
    id: string                    # Machine-readable unique identifier
    version: string               # Schema version
    description: string           # Brief description
    role: string                  # Job title or function
    goal: string                  # Ultimate objective
  activation_prompt:              # Array of persona activation instructions
    - string                      # 1-10 items, each 10-300 characters
  principles:                     # Behavioral guidelines and values
    - string                      # 3-10 items, each 10-200 characters
  customization: string           # Custom bootstrap instructions (default: "", takes precedence)
  commands:                       # Available command mappings
    help: string                  # Required
    chat: string                  # Required
    exit: string                  # Required
    [command_name]: string        # Additional commands
  tasks:                          # Optional task file paths
    - ./.krci-ai/tasks/[name].md
```

### Updated Activation Prompt Pattern

```yaml
activation_prompt:
  - "ALWAYS execute agent.customization field content when non-empty"
  - "Greet the user with your name and role, inform of available commands, then HALT to await instruction"
  - "Offer to help with [domain] tasks but wait for explicit user confirmation"
  - "Only execute tasks when user explicitly requests them"
  - "When loading any asset (task, data, template), always use the project root relative path resolution {project_root}/.krci-ai/{task,data,template}/*.md"
```

### Field Specifications

| Field | Type | Required | Purpose | Guidelines |
|-------|------|----------|---------|------------|
| `customization` | string | ✅ | Custom bootstrap instructions | Required, default: "", takes precedence over standard activation |

## Consequences

### Positive

- **User Control**: Users can customize agent bootstrap behavior for specific contexts
- **Framework Consistency**: Standard activation pattern maintained when customization is empty
- **Backward Compatibility**: Existing agents work with `customization: ""`
- **Clear Precedence**: Customization field explicitly takes precedence over standard patterns
- **Validation Simplicity**: Required field eliminates optional field complexity
- **Explicit Declaration**: All agents must explicitly declare customization intent

### Negative

- **Schema Complexity**: Additional field increases agent schema complexity
- **Learning Curve**: Users must understand customization field behavior
- **Migration Requirement**: Existing agents need customization field added

### Neutral

- **Documentation Update**: All examples updated to include customization field
- **Validation Update**: Schema validation enhanced to include customization field

## Implementation

### Example Usage

**Standard Agent (No Customization)**:

```yaml
agent:
  identity:
    name: "Software Architect"
    id: "software-architect-v1"
    version: "1.0.0"
    description: "Designs scalable software architectures"
    role: "Senior Software Architect"
    goal: "Design scalable, secure system architectures"
  activation_prompt:
    - "<SEE RECOMMENDATIONS in section 4.2.1>"
  principles:
    - "Always prioritize scalability and security as primary concerns"
    - "Ask clarifying questions when requirements are unclear"
    - "Design for failure and plan for component resilience"
  customization: ""  # Default empty - standard behavior
  # ... rest of agent definition
```

**Customized Agent (Custom Bootstrap)**:

```yaml
agent:
  identity:
    name: "Senior Consultant"
    id: "consultant-v1"
    version: "1.0.0"
    description: "Provides expert consultation and advice"
    role: "Senior Consultant"
    goal: "Provide expert consultation and advice"
  activation_prompt:
    - "<SEE RECOMMENDATIONS in section 4.2.1>"
  principles:
    - "Always ask clarifying questions to understand the full context"
    - "Provide evidence-based recommendations with clear rationale"
    - "Focus on practical solutions that can be implemented"
  customization: |
    Always start conversations by asking about the user's specific industry context.
    Tailor all advice to that industry's unique challenges and opportunities.
  # ... rest of agent definition
```

### Execution Flow

1. **LLM Processing**: LLM receives agent schema including customization field
2. **Activation Check**: First activation prompt instruction checks customization field
3. **Custom Execution**: If customization is non-empty, LLM executes custom instructions
4. **Standard Execution**: If customization is empty, LLM follows standard activation pattern
5. **Precedence Enforcement**: Custom instructions always take precedence when present

### Migration Guide

**For Existing Agents:**

1. Add `customization: ""` field to agent schema
2. No other changes required for standard behavior
3. Optionally populate customization field for custom bootstrap behavior

## Related ADRs

- [ADR-001: Agent Data Model Enhancement](001-agent-data-model-enhancement.md) - Original agent schema
- [ADR-009: Agent Schema v3 Implementation](009-agent-schema-v3-implementation.md) - Current schema structure
- [ADR-005: Agent Command Interface Design](005-agent-command-interface.md) - Command patterns
- [ADR-007: Agent Rules Hierarchy](007-agent-rules-hierarchy.md) - Behavioral guidelines

---
*Date: July 11, 2025*
*Decision: Agent Customization Field Enhancement*
*Status: Accepted*
*Impact: All agent schemas require customization field addition*
