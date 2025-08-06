# ADR-002: Template Reference Pattern

## Status

Accepted

## Context

The KubeRocketAI framework needed to establish clear rules about where and how templates should be referenced to maintain separation of concerns and enable component reusability.

### Initial Challenge

Without clear reference rules, templates could be referenced from any component, leading to:

- **Tight Coupling**: Agents locked into specific output formats
- **Poor Reusability**: Same agent couldn't produce different outputs for different tasks
- **Unclear Ownership**: Confusion about which component controls output formatting

### Forces

- **Separation of Concerns**: Each component type should have distinct responsibilities
- **Composability**: Any agent should be able to execute any task
- **Flexibility**: Same agent should produce different formats for different tasks
- **Framework Consistency**: Clear, enforceable rules for component relationships

## Decision

**Framework Rule: Template references MUST be in Tasks, NEVER in Agents**

### Component Responsibilities

| Component | Responsibility | Template Usage |
|-----------|----------------|----------------|
| **Agent** | Identity and behavior ("WHO you are, HOW you think") | ❌ No templates |
| **Task** | Actions and output ("WHAT to do, HOW to format") | ✅ Templates allowed |

### Correct Pattern

```yaml
# Agent defines WHO and HOW to behave
agent:
  name: software-architect
  id: software-architect-v1
  role: "Senior Software Architect"
  behavioral_rules:
    - "Always consider scalability and security first"
    # NO template references in agent!

---
# Task defines WHAT to do and HOW to format output
task: analyze-requirements
description: Extract and analyze project requirements
instructions: |
  1. Read the input document thoroughly
  2. Extract functional and non-functional requirements
  3. Format output using {{ ref:./.krci-ai/templates/analyze-requirements-output.md }}
```

### Incorrect Pattern

```yaml
# DON'T DO THIS - Agent controlling output format
agent:
  name: software-architect
  behavioral_rules:
    - "Always use {{ ref:./.krci-ai/templates/architecture-doc.md }}"
    # This locks agent into one format for all tasks - BAD!
```

## Consequences

### Positive

- **Clean Separation**: Agents define behavior, tasks define output formatting
- **Maximum Composability**: Any agent can execute any task without format conflicts
- **Flexible Outputs**: Same agent can produce different formats for different tasks
- **Scalable Growth**: Framework grows by adding tasks, not modifying agents

### Example Benefits

- **Architect Agent** can:
  - Write requirements (using requirements template)
  - Create API specs (using API template)
  - Document decisions (using decision template)
- Each task controls its own appropriate formatting

### Negative

- **Learning Curve**: Developers must understand the reference pattern
- **Template Discovery**: Need to find templates through task definitions

### Neutral

- **CI/CD Analogy**: Follows familiar pattern where jobs control their output format
- **Enforcement Required**: Framework must validate reference locations

## Implementation

This rule is enforced through:

1. **Documentation**: Clear examples in 04_Data_Models.md
2. **Validation**: Framework checks reference locations
3. **Examples**: All agent examples demonstrate correct pattern

The pattern enables the framework's core philosophy: components that can be mixed and matched for maximum flexibility.

## Related ADRs

- [ADR-001: Agent Data Model Enhancement](001-agent-data-model-enhancement.md)
- [ADR-004: Inline Reference Pattern](004-inline-reference-pattern.md)
