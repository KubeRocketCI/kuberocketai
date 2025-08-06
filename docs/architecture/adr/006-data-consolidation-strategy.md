# ADR-006: Data Asset Consolidation Strategy

## Status

Accepted

## Context

The KubeRocketAI framework initially had separate component types for different kinds of reference materials (docs, specs, schemas). This created complexity in understanding which component type to use for different reference materials and complicated the directory structure.

### Initial Separation

- **Docs**: Documentation and guidelines
- **Specs**: API specifications and schemas
- **Schemas**: Validation rules and constraints
- **Data**: Configuration and preferences

### Forces

- **Simplicity**: Fewer component types easier to understand
- **User Freedom**: Users should organize reference materials as they prefer
- **Clear Boundaries**: Distinction between active components and reference materials
- **Dual Purpose**: Some reference materials serve both behavioral and technical purposes

## Decision

**Consolidate all static reference materials into a single "Data" component type**

### Unified Data Component

| Content Type | Purpose | Usage |
|--------------|---------|-------|
| **API Specs** | Technical constraints | Referenced by tasks |
| **Schemas** | Validation rules | Referenced by tasks |
| **Documentation** | Guidelines and knowledge | Referenced by agents |
| **Configuration** | Settings and preferences | Referenced by any component |

### Directory Structure

```bash
./.krci-ai/
├── agents/     # WHO (personas) - Required structure
├── tasks/      # WHAT (actions) - Required structure
├── templates/  # HOW (formats) - Required structure
└── data/       # REFERENCE - User-defined sub-structure
    ├── specs/        # User choice
    │   ├── apis/
    │   └── schemas/
    ├── docs/         # User choice
    └── config/       # User choice
```

### Dual Reference Pattern

**Agents reference BEHAVIORAL data:**

```yaml
agent:
  behavioral_rules:
    - "Follow principles in {{ ref:./.krci-ai/data/docs/design-principles.md }}"
    - "Apply security framework from {{ ref:./.krci-ai/data/docs/security-guidelines.md }}"
```

**Tasks reference TECHNICAL data:**

```yaml
task: create-user-api
instructions: |
  1. Follow API spec in {{ ref:./.krci-ai/data/specs/apis/user-api.yaml }}
  2. Validate against {{ ref:./.krci-ai/data/specs/schemas/user.json }}
```

### User Organization Freedom

Users can organize the `data/` directory however they prefer:

- By domain: `data/business/`, `data/technical/`
- By type: `data/specs/`, `data/docs/`, `data/config/`
- By project: `data/project-a/`, `data/shared/`

## Consequences

### Positive

- **Simplified Mental Model**: Only 4 component types instead of 6+
- **User Freedom**: Complete flexibility in organizing reference materials
- **Clear Distinction**: Active components (agents, tasks, templates) vs. reference materials (data)
- **Dual Purpose Support**: Same data can serve behavioral and technical needs
- **Reduced Complexity**: Fewer validation rules and directory requirements

### Negative

- **Learning Curve**: Users must understand dual reference pattern
- **Organization Responsibility**: Users must create their own organization within data/

### Neutral

- **Framework Validation**: Only validates that referenced files exist, not organization
- **Migration Impact**: Existing separate components need consolidation

## Implementation

The consolidation is implemented through:

1. **Single Data Directory**: All static reference materials in `./.krci-ai/data/`
2. **User-Defined Organization**: Framework doesn't enforce sub-directory structure
3. **Dual Reference Rules**:
   - Agents reference behavioral data (guidelines, principles)
   - Tasks reference technical data (specs, schemas, validation)
4. **Path Validation**: Framework validates all `{{ ref:./.krci-ai/data/... }}` references

### Migration Examples

**Before (Complex):**

```bash
./.krci-ai/
├── docs/system-design-principles.md
├── specs/apis/user-api.yaml
├── schemas/user.json
└── config/preferences.yaml
```

**After (Consolidated):**

```bash
./.krci-ai/
└── data/
    ├── docs/system-design-principles.md
    ├── specs/
    │   ├── apis/user-api.yaml
    │   └── schemas/user.json
    └── config/preferences.yaml
```

### Framework Benefits

- **4 Component Types**: Agent, Task, Template, Data (down from 6+)
- **Clear Boundaries**: Active vs. reference materials
- **Maximum Flexibility**: Users organize data as they prefer
- **Consistent References**: Single reference pattern for all static materials

This consolidation provides the optimal balance of simplicity and flexibility for managing reference materials in the framework.

## Related ADRs

- [ADR-003: Four Component Separation Model](003-component-separation-model.md)
- [ADR-002: Template Reference Pattern](002-template-reference-pattern.md)
- [ADR-004: Inline Reference Pattern](004-inline-reference-pattern.md)
