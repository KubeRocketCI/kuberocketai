# ADR-001: Agent Data Model Enhancement

## Status

Accepted

## Context

The original Agent data model in the KubeRocketAI framework had several limitations that hindered effective AI agent behavior:

1. **Identity Management**: Agents lacked stable, machine-readable identifiers for reliable referencing
2. **Behavioral Structure**: Simple `instructions` field was insufficient for rich persona definition
3. **Command Simplicity**: Commands needed cleaner structure while maintaining functionality
4. **Evolution Support**: No versioning mechanism for schema evolution

### Forces

- **Simplicity**: Framework must remain simple and maintainable
- **AI Agent Effectiveness**: Agents need proper identity, behavioral rules, and persona context
- **Backwards Compatibility**: Changes should not break existing implementations unnecessarily
- **Clean Design**: Avoid over-engineering while providing essential functionality

## Decision

We enhanced the Agent data model with minimal essential extensions:

### Original Schema (v1)

```yaml
agent:
  name: string
  description: string
  role: string
  goal: string
instructions: string
commands:
  [command_name]: string → uses: task_name
tasks: array[string]
```

### Enhanced Schema (v2)

```yaml
agent:
  name: string                    # Agent identifier
  id: string                      # Machine-readable unique identifier
  version: string                 # Schema version for capability evolution
  description: string             # Brief description of agent's purpose
  role: string                    # Job title or function
  goal: string                    # Ultimate objective or mission
activation_prompt: string         # Detailed LLM context for persona embodiment
behavioral_rules: array[string]   # Hard constraints that cannot be violated
commands:                         # Available command mappings
  help: string                    # Always present
  chat: string                    # Always present - default mode
  [command_name]: string → uses: task_name  # Clean one-liner format
tasks: array[string]              # List of available task identifiers
```

### Key Changes

1. **Added Identity Fields**:
   - `id`: Stable machine-readable identifier
   - `version`: Schema version for evolution tracking

2. **Structured Behavioral Control**:
   - `activation_prompt`: Rich LLM context replacing simple instructions
   - `behavioral_rules`: Array of hard constraints that cannot be violated

3. **Command Simplicity**:
   - Maintained clean one-liner format: `command: "description" → uses: task-name`
   - Rejected complex nested structures for maximum simplicity

4. **Removed Legacy**:
   - `instructions` field replaced with structured approach

## Consequences

### Positive

- **Better Agent Identity**: Stable references and version-aware evolution
- **Behavioral Precision**: Rich persona context with enforcement rules
- **Clean Commands**: Simple, readable command structure
- **Future-Proof**: Schema designed for extensibility

### Negative

- **Migration Required**: Existing agents need updating to new schema
- **Slight Complexity**: More fields than original simple version

### Neutral

- **No Resource Duplication**: Resources managed at appropriate levels (agent/task)
- **Maintained Philosophy**: Essential enhancements without over-engineering

## Implementation

All agent examples in `04_Data_Models.md` have been updated to use the enhanced schema consistently. The framework now provides robust agent identity and behavioral control while maintaining clean design principles.

## Related ADRs

- [ADR-005: Agent Command Interface Design](005-agent-command-interface.md)
- [ADR-003: Four Component Separation Model](003-component-separation-model.md)
