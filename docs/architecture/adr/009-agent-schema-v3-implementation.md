# ADR-009: Agent Schema v3 Implementation

## Status

Accepted

## Context

The KubeRocketAI framework has evolved from the initial v2 agent schema (documented in ADR-001) to the current v3 implementation that aligns with the JSON schema used by the CLI tool. This evolution was driven by real-world usage and the need for better validation, structure, and consistency.

### Schema Evolution History

- **v1**: Original flat structure with basic fields
- **v2**: Enhanced with identity fields and behavioral rules (ADR-001)
- **v3**: Current implementation with nested identity structure and array-based fields

### Forces

- **JSON Schema Alignment**: CLI tool uses JSON schema for validation
- **Nested Structure**: Better organization of identity-related fields
- **Array-based Fields**: More flexible activation prompts and principles
- **Required Commands**: Standardized command interface
- **Validation Consistency**: Same structure for validation and documentation

## Decision

**Implement Agent Schema v3 with full alignment to KubeRocketAI JSON schema**

### Current Schema (v3)

```yaml
agent:
  identity:
    name: string                  # Friendly user name (pattern: ^[A-Z][a-zA-Z0-9 .'-]{1,49}$)
    icon: string                  # Optional icon (emoji or short string, max 4 chars)
    id: string                    # Machine-readable unique identifier with version
    version: string               # Schema version (semantic versioning)
    description: string           # Brief description (10-150 characters)
    role: string                  # Job title or function (5-100 characters)
    goal: string                  # Ultimate objective (10-200 characters)
  activation_prompt:              # Array of persona activation instructions
    - string                      # 1-10 items, each 10-300 characters
  principles:                     # Behavioral guidelines and values
    - string                      # 3-10 items, each 10-200 characters
  commands:                       # Available command mappings (3-20 commands)
    help: string                  # Required - "Show available commands"
    chat: string                  # Required - "(Default) [Domain] consultation"
    exit: string                  # Required - "Exit persona command"
    [command_name]: string        # Additional commands (5-200 characters)
  tasks:                          # Optional - List of task file paths
    - ./.krci-ai/tasks/[name].md  # Must match ./.krci-ai/tasks/*.md pattern
```

### Key Changes from v2

1. **Nested Identity Structure**: All identity fields grouped under `identity` object
2. **Icon Field**: Added optional `icon` field for visual representation
3. **Array-based activation_prompt**: Changed from single string to array (1-10 items)
4. **Principles Instead of behavioral_rules**: Renamed for clarity
5. **Required Commands**: Added mandatory `exit` command alongside `help` and `chat`
6. **File Path Tasks**: Tasks reference actual file paths instead of identifiers
7. **Validation Constraints**: Added specific length and count constraints

### Validation Rules

- **Identity Name**: Must match pattern `^[A-Z][a-zA-Z0-9 .'-]{1,49}$`
- **Icon**: Optional, max 4 characters
- **Description**: 10-150 characters
- **Role**: 5-100 characters
- **Goal**: 10-200 characters
- **Activation Prompt**: 1-10 items, each 10-300 characters
- **Principles**: 3-10 items, each 10-200 characters
- **Commands**: 3-20 total, descriptions 5-200 characters
- **Tasks**: Must match `./.krci-ai/tasks/*.md` pattern

## Consequences

### Positive

- **Full JSON Schema Alignment**: Documentation matches CLI tool validation
- **Better Organization**: Nested identity structure groups related fields
- **Flexible Activation**: Array-based prompts allow multi-part persona setup
- **Consistent Validation**: Same rules apply to documentation and CLI
- **Enhanced Usability**: Required commands ensure consistent user experience
- **Clear Constraints**: Explicit validation rules prevent common errors

### Negative

- **Migration Complexity**: Existing v2 agents need updating to v3 structure
- **Schema Complexity**: More nested structure than previous versions
- **Learning Curve**: Users must understand new nested format

### Neutral

- **Backward Compatibility**: v2 agents need manual migration to v3
- **Documentation Update**: All examples updated to reflect v3 structure
- **ADR Preservation**: Previous ADRs remain as historical records

## Implementation

### Documentation Updates

All agent examples in `04_Data_Models.md` have been updated to use the v3 schema structure, including:

- Canonical agent schema
- Progressive complexity examples (Levels 1-4)
- Real-world scenarios
- Constraint examples
- Connection patterns

### Example v3 Agent

```yaml
agent:
  identity:
    name: "Winston"
    icon: "🏗️"
    id: "software-architect-v1"
    version: "1.0.0"
    description: "Designs scalable and secure software architectures"
    role: "Senior Software Architect"
    goal: "Design scalable, secure system architectures"
  activation_prompt:
    - "<SEE RECOMMENDATIONS in 04_Data_Models.md section 4.2.1>"
  principles:
    - "Always prioritize scalability and security as primary concerns"
    - "Ask clarifying questions when requirements are unclear"
    - "Design for failure and plan for component resilience"
  commands:
    help: "Show available commands"
    chat: "(Default) Architectural consultation"
    exit: "Exit architect mode"
    analyze: "Analyze requirements"
    design: "Create system design"
  tasks:
    - "./.krci-ai/tasks/analyze-requirements.md"
    - "./.krci-ai/tasks/create-system-design.md"
```

### Migration Guide

**From v2 to v3:**

1. **Wrap identity fields**: Move `name`, `id`, `version`, `description`, `role`, `goal` under `identity` object
2. **Add icon field**: Optional `icon` field under `identity`
3. **Convert activation_prompt**: Change from string to array format
4. **Rename behavioral_rules**: Change to `principles`
5. **Add exit command**: Add required `exit` command to commands
6. **Update task paths**: Change task identifiers to file paths with `.md` extension

## Related ADRs

- [ADR-001: Agent Data Model Enhancement](001-agent-data-model-enhancement.md) - Historical v2 schema
- [ADR-005: Agent Command Interface Design](005-agent-command-interface.md) - Command structure
- [ADR-007: Agent Rules Hierarchy](007-agent-rules-hierarchy.md) - Principles definition

---
*Date: July 8, 2025*
*Decision: Agent Schema v3 Implementation*
*Status: Implemented in KubeRocketAI framework*
*JSON Schema: https://raw.githubusercontent.com/KubeRocketCI/kuberocketai/refs/heads/main/cmd/krci-ai/assets/schemas/agent-schema.json*