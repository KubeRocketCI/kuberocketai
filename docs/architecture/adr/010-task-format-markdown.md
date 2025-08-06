# ADR-010: Task Format Markdown Implementation

## Status

Accepted

## Context

The KubeRocketAI framework originally defined tasks using YAML format with frontmatter structure. However, the implementation evolved to use simple markdown files with YAML frontmatter for better readability, maintainability, and alignment with documentation practices.

### Original Task Format (Legacy)

```yaml
task: analyze-requirements
description: Extract and analyze business requirements from documents
instructions: |
  1. Review the input document thoroughly
  2. Extract functional and non-functional requirements
  3. Identify gaps and missing requirements
```

### Forces

- **Documentation Consistency**: Tasks should align with markdown-first documentation approach
- **Readability**: Markdown provides better formatting for complex instructions
- **Tool Integration**: Better support for markdown in editors and documentation systems
- **Reference Integration**: Inline references work naturally in markdown format
- **Simplicity**: Reduce format complexity while maintaining structure

## Decision

**Implement tasks as simple markdown files with YAML frontmatter**

### Current Task Format

```markdown
---
task: string                      # Task identifier (kebab-case)
description: string               # What this task accomplishes
---

# Task: {task name}

## Description
{What this task accomplishes}

## Instructions
{Step-by-step procedural workflow with inline references}

1. {First step}
2. {Second step}
3. {etc.}
```

### Key Benefits

1. **Markdown-First Approach**: Aligns with documentation ecosystem
2. **Better Formatting**: Support for headers, lists, code blocks, emphasis
3. **Inline References**: Natural integration of `{{ ref:path }}` syntax
4. **Editor Support**: Better syntax highlighting and editing experience
5. **Version Control**: Better diff visualization for task changes
6. **Human Readable**: Tasks are documentation that happens to be executable

### Structure Requirements

- **YAML Frontmatter**: Contains task identifier and description
- **Markdown Headers**: Clear section organization
- **Numbered Instructions**: Step-by-step procedural workflow
- **Inline References**: `{{ ref:./.krci-ai/path/to/file }}` syntax for file references

## Consequences

### Positive

- **Enhanced Readability**: Markdown provides superior formatting for complex instructions
- **Better Tool Support**: Editors provide better support for markdown editing
- **Documentation Alignment**: Tasks become part of the documentation ecosystem
- **Natural Reference Integration**: Inline references flow naturally in markdown
- **Version Control Benefits**: Better diff visualization and merge conflict resolution
- **Accessibility**: Tasks are readable as documentation without special tools

### Negative

- **Format Migration**: Existing YAML tasks need conversion to markdown format
- **Tool Complexity**: Parsing requires both YAML frontmatter and markdown content
- **Size Overhead**: Markdown format may be slightly larger than pure YAML

### Neutral

- **File Extension**: Tasks use `.md` extension instead of `.yaml`
- **Directory Structure**: No change to task organization in `./.krci-ai/tasks/`
- **Reference Syntax**: Inline reference syntax remains unchanged

## Implementation

### Task File Structure

```markdown
<!-- File: tasks/analyze-requirements.md -->
---
task: analyze-requirements
description: Extract and analyze business requirements from documents
---

# Task: Analyze Requirements

## Description
Extract and analyze business requirements from documents

## Instructions

1. Review the input document thoroughly
2. Extract functional and non-functional requirements
3. Identify gaps and missing requirements
4. Generate recommendations for completeness
5. Format output using {{ ref:./.krci-ai/templates/requirements-analysis.md }}
```

### Reference Integration

Tasks can reference templates and data files naturally within markdown:

```markdown
## Instructions

1. Follow API specification in {{ ref:./.krci-ai/data/specs/apis/user-api.yaml }}
2. Validate against schema in {{ ref:./.krci-ai/data/specs/schemas/user.json }}
3. Format output using {{ ref:./.krci-ai/templates/api-documentation.md }}
```

### Agent Task References

Agents reference tasks using markdown file paths:

```yaml
agent:
  tasks:
    - "./.krci-ai/tasks/analyze-requirements.md"
    - "./.krci-ai/tasks/create-system-design.md"
```

### Directory Structure

```bash
.krci-ai/
├── agents/
│   └── architect.yaml
├── tasks/
│   ├── analyze-requirements.md      # Markdown format
│   └── create-system-design.md      # Markdown format
├── templates/
│   └── requirements-analysis.md
└── data/
    └── docs/
        └── principles.md
```

### Migration Guide

**From YAML to Markdown:**

1. **Convert structure**: Wrap YAML fields in frontmatter
2. **Add markdown headers**: Create `# Task: Name` and `## Instructions`
3. **Update file extension**: Change from `.yaml` to `.md`
4. **Enhance formatting**: Use markdown features for better readability
5. **Update references**: Ensure agent task lists use `.md` extensions

## Related ADRs

- [ADR-002: Template Reference Pattern](002-template-reference-pattern.md) - Reference patterns in tasks
- [ADR-004: Inline Reference Pattern](004-inline-reference-pattern.md) - Inline reference syntax
- [ADR-009: Agent Schema v3 Implementation](009-agent-schema-v3-implementation.md) - Current agent structure

---
*Date: July 8, 2025*
*Decision: Task Format Markdown Implementation*
*Status: Implemented in KubeRocketAI framework*
*File Format: Markdown with YAML frontmatter (.md)*