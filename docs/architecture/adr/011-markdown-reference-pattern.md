# ADR-011: Markdown Reference Pattern Implementation

## Status

Accepted

## Context

The KubeRocketAI framework originally used a custom `{{ ref:path/to/file }}` syntax for referencing files within tasks and agent principles. However, this custom syntax deviated from standard markdown practices and required special parsing logic.

### Original Reference Syntax (Legacy)

```markdown
## Instructions

1. Follow API specification in {{ ref:./.krci-ai/data/specs/apis/user-api.yaml }}
2. Validate against schema in {{ ref:./.krci-ai/data/specs/schemas/user.json }}
3. Format output using {{ ref:./.krci-ai/templates/api-documentation.md }}
```

### Forces

- **Standard Compliance**: Use standard markdown link syntax for better tool support
- **Editor Integration**: Better syntax highlighting and link validation in editors
- **Simplicity**: Reduce custom syntax in favor of markdown standards
- **Tool Ecosystem**: Leverage existing markdown tooling and linters
- **Readability**: Standard markdown links are more familiar to developers

## Decision

**Replace custom `{{ ref:path/to/file }}` syntax with standard markdown links `[filename](path/to/file)`**

### Current Reference Pattern

```markdown
## Instructions

1. Follow API specification in [user-api.yaml](./.krci-ai/data/specs/apis/user-api.yaml)
2. Validate against schema in [user.json](./.krci-ai/data/specs/schemas/user.json)
3. Format output using [api-documentation.md](./.krci-ai/templates/api-documentation.md)
```

### Reference Format Standards

- **Link Text**: Use filename only for brevity (e.g., `[user-api.yaml]`)
- **Path Format**: Relative paths from document location (e.g., `../.krci-ai/data/specs/apis/user-api.yaml`)
- **File Extensions**: Always include file extensions for clarity
- **Path Prefix**: Use `../.krci-ai/` as the base path from documentation to framework directory

### Syntax Migration

| **Original Syntax** | **New Syntax** |
|---------------------|----------------|
| `{{ ref:./.krci-ai/data/specs/api.yaml }}` | `[api.yaml](./.krci-ai/data/specs/api.yaml)` |
| `{{ ref:./.krci-ai/templates/output.md }}` | `[output.md](./.krci-ai/templates/output.md)` |
| `{{ ref:./.krci-ai/data/docs/principles.md }}` | `[principles.md](./.krci-ai/data/docs/principles.md)` |

## Consequences

### Positive

- **Standard Compliance**: Uses standard markdown link syntax
- **Better Tool Support**: Editors provide native link validation and navigation
- **Improved Readability**: Familiar markdown link format for all developers
- **Reduced Complexity**: No custom parsing logic required
- **Editor Features**: Hover previews, click-to-navigate, and syntax highlighting
- **Linting Support**: Standard markdown linters can validate links

### Negative

- **Migration Effort**: Existing `{{ ref: }}` patterns need conversion
- **Path Management**: Relative paths require careful maintenance
- **Tool Dependency**: Framework tools must understand markdown link format

### Neutral

- **Validation Logic**: Static validation must parse markdown links instead of custom syntax
- **Documentation Impact**: All examples updated to reflect new pattern
- **Framework Compatibility**: Both syntaxes could be supported during transition

## Implementation

### Reference Pattern Rules

1. **Inline References Only**: No separate reference sections
2. **Standard Markdown Links**: `[filename](path/to/file)` format
3. **Relative Paths**: Use `../.krci-ai/` prefix from documentation
4. **Descriptive Link Text**: Use filename for link text
5. **File Extension Required**: Always include file extensions

### Usage Examples

#### Agent Behavioral References

```yaml
agent:
  principles:
    - "Follow design principles from [architecture-principles.md](./.krci-ai/data/docs/architecture-principles.md)"
    - "Apply security guidelines from [security-framework.md](./.krci-ai/data/docs/security-framework.md)"
```

#### Task Technical References

```markdown
---
task: create-user-api
description: Create REST API for user management
---

# Task: Create User API

## Instructions

1. Follow API specification in [user-api.yaml](./.krci-ai/data/specs/apis/user-api.yaml)
2. Validate against schema in [user.json](./.krci-ai/data/specs/schemas/user.json)
3. Format output using [api-documentation.md](./.krci-ai/templates/api-documentation.md)
```

### Validation Updates

Framework validation must be updated to:

1. **Parse Markdown Links**: Extract file paths from `[text](path)` syntax
2. **Validate Paths**: Ensure all referenced files exist
3. **Check Accessibility**: Verify files are readable
4. **Report Broken Links**: Provide clear error messages for missing files

### Migration Guide

**From `{{ ref: }}` to Markdown Links:**

1. **Extract Path**: Get path from `{{ ref:path }}`
2. **Extract Filename**: Get filename from end of path
3. **Create Link**: Format as `[filename](path)`
4. **Update Path**: Adjust path relativity if needed
5. **Validate**: Ensure link resolves correctly

**Example Migration:**

```diff
- 1. Follow spec in {{ ref:./.krci-ai/data/specs/api.yaml }}
+ 1. Follow spec in [api.yaml](./.krci-ai/data/specs/api.yaml)

- 2. Format using {{ ref:./.krci-ai/templates/output.md }}
+ 2. Format using [output.md](./.krci-ai/templates/output.md)
```

## Related ADRs

- [ADR-002: Template Reference Pattern](002-template-reference-pattern.md) - Template reference rules
- [ADR-004: Inline Reference Pattern](004-inline-reference-pattern.md) - Inline reference approach
- [ADR-010: Task Format Markdown Implementation](010-task-format-markdown.md) - Markdown task format

---
*Date: July 8, 2025*
*Decision: Markdown Reference Pattern Implementation*
*Status: Implemented in KubeRocketAI framework*
*Syntax: Standard markdown links `[filename](path/to/file)`*
