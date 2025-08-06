# ADR-004: Inline Reference Pattern

## Status

Accepted

## Context

The KubeRocketAI framework needed a way for components to reference other assets (templates, data, etc.). The challenge was choosing between inline references embedded in instructions versus separate reference sections.

### Alternative Approaches Considered

1. **Separate Reference Sections**: Maintain reference lists separate from instructions
2. **Inline References**: Embed references directly where they're used
3. **Hybrid Approach**: Mix of both patterns

### Forces

- **Single Source of Truth**: References should be defined exactly where they're used
- **Context Clarity**: Users should immediately see what each reference is for
- **Maintenance Burden**: Changes should only need to be made in one place
- **Readability**: Instructions should flow naturally with embedded links
- **Simplicity**: Avoid redundant information or complex structures

## Decision

**Framework Rule: Use ONLY inline references `{{ ref:path/to/file }}` - NO separate reference sections**

### Correct Pattern

```yaml
task: create-user-api
description: Create REST API for user management
instructions: |
  1. Follow API specification in {{ ref:./.krci-ai/data/specs/apis/user-api.yaml }}
  2. Validate against schema in {{ ref:./.krci-ai/data/specs/schemas/user.json }}
  3. Apply rate limiting from {{ ref:./.krci-ai/data/specs/validation/rate-limits.yaml }}
  4. Format using {{ ref:./.krci-ai/templates/create-user-api-output.md }}

# NO separate references section needed!
```

### Incorrect Pattern

```yaml
# DON'T DO THIS - Creates redundancy and maintenance overhead
task: create-user-api
instructions: |
  1. Follow the API specification
  2. Validate against the schema
  3. Apply rate limiting rules
  4. Format using the template
references:                              # REDUNDANT!
  - data/specs/apis/user-api.yaml        # Which step uses this?
  - data/specs/schemas/user.json         # Hard to connect to instructions
  - data/specs/validation/rate-limits.yaml
  - templates/tasks/create-user-api/output.md
```

### Reference Syntax

- **Pattern**: `{{ ref:path/to/file }}`
- **Base Path**: All paths relative to `./.krci-ai/`
- **File Types**: Any file within the framework structure

### Best Practices

1. **Keep Paths Concise**: Use short, clear paths
2. **Meaningful Names**: Use self-describing file names
3. **Context Integration**: References flow naturally in instructions
4. **Validation**: Framework validates all referenced files exist

## Consequences

### Positive

- **Single Source of Truth**: Reference defined exactly where it's used
- **Clear Context**: Immediately see what each reference is for
- **Natural Flow**: Instructions read like normal text with embedded links
- **Easier Maintenance**: Change reference in one place only
- **Simpler Structure**: No need to maintain two separate sections
- **Eliminates Redundancy**: No duplicate information between sections

### Negative

- **Path Management**: Need to maintain correct relative paths
- **Reference Discovery**: Can't easily see all references for a component
- **Validation Complexity**: Framework must parse instructions for references

### Neutral

- **Learning Curve**: Users need to learn reference syntax
- **Framework Dependency**: References tied to framework validation

## Implementation

The pattern is implemented through:

1. **Reference Syntax**: `{{ ref:./.krci-ai/path/to/file }}`
2. **Framework Validation**: All references validated at load time
3. **Error Reporting**: Clear messages for broken or missing references
4. **Path Resolution**: Automatic resolution of relative paths

### Example Integration

```yaml
agent:
  behavioral_rules:
    - "Follow principles in {{ ref:./.krci-ai/data/docs/design-principles.md }}"

task: analyze-requirements
instructions: |
  1. Review specification in {{ ref:./.krci-ai/data/specs/business/requirements.yaml }}
  2. Format output using {{ ref:./.krci-ai/templates/analysis-output.md }}
```

This pattern provides the optimal balance of simplicity, clarity, and maintainability for component relationships in the framework.

## Related ADRs

- [ADR-002: Template Reference Pattern](002-template-reference-pattern.md)
- [ADR-006: Data Asset Consolidation Strategy](006-data-consolidation-strategy.md)
