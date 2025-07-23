# SDLC Framework Quick Reference

**Purpose**: This framework enables AI agents to collaborate in software development through structured, filesystem-based artifacts. Agents reference this document to understand their role, dependencies, and where to find required templates and standards.

## Framework Principles

- **Filesystem-First**: All artifacts stored as markdown files, no APIs
- **Agent Discovery**: Agents find artifacts through organized directory structure
- **Natural Language**: Human-readable artifacts that agents can process
- **Clear Dependencies**: Each artifact builds on previous work
- **Inline References**: Use `[filename](path/to/file)` markdown links for all references

## Role Responsibilities

| Role | Outputs | Dependencies |
|------|---------|--------------|
| **Product Manager** | Project Brief, PRD, Roadmap | None (root artifacts) |
| **Product Owner** | Epics, Stories | PRD |
| **Business Analyst** | Refined PRD, Workflows | PRD, Stakeholder inputs |
| **Architect** | Architecture Documents | PRD, Epics |
| **Developer** | Code, Implementation | Stories, Architecture |
| **QA Engineer** | Test Results, Quality Reports | Stories, Code |

## Artifact Flow

```text
Project Brief → PRD → Epics → Stories → Code → Tests → MVP
                  ↓
              Architecture ← → Code
```

**Dependency Rules:**

- **Project Brief**: No dependencies (root artifact)
- **PRD**: Requires Project Brief approval
- **Epic**: Requires PRD completion, references specific BR/NFR requirements
- **Story**: Requires Epic definition, maps to implementation tasks
- **Architecture**: Requires PRD + Epic context for technical design
- **Code**: Requires Stories + Architecture for implementation guidance
- **Tests**: Requires Stories + Code for validation

## Directory Structure

```bash
{project_root}/
├── docs/                           # All SDLC artifacts
│   ├── prd/                        # PM: Project Brief, PRD, Roadmap
│   │   ├── project-brief.md        # Vision & strategy
│   │   ├── prd.md                  # Business requirements (BR1, BR2...), system requirements (NFR1, NFR2...)
│   │   └── roadmap.md              # Implementation timeline
│   ├── epics/                      # PO: High-level features
│   │   └── {epic_number}-epic-{slug}.md    # e.g., 01-epic-user-authentication.md
│   ├── stories/                    # PO: User requirements with tasks
│   │   └── {epic_number}.{story_number}.story.md    # e.g., 01.01.story.md
│   ├── architecture/               # Architect: System design
│   │   ├── adr/                    # Architecture Decision Records
│   │   ├── 01-introduction.md      # System overview
│   │   ├── 02-high-level-architecture.md
│   │   └── [other numbered sections]
│   └── tests/                      # QA: Quality validation
│       └── test-results-{story}.md
└── .krci-ai/                       # Framework assets
    ├── agents/                     # WHO: Role definitions
    ├── tasks/                      # WHAT: Procedural workflows
    ├── templates/                  # HOW: Output formatting
    └── data/                       # REFERENCE: Standards & guidelines
```

## Quality Gates

1. **Project Brief Approval** → Enables PRD creation
2. **PRD Approval** → Enables Epic/Architecture creation
3. **Architecture Review** → Enables implementation
4. **Code Review** → Enables testing
5. **Test Validation** → Enables MVP delivery

## Template Variables

All templates use `{{variable_name}}` format for dynamic content:

- **Required fields**: Must be populated, no empty `{{}}` allowed
- **Dependencies**: Reference parent artifacts (Epic references PRD requirements)
- **Traceability**: Link requirements to implementation (BR1, NFR2, etc.)

## Success Flow

**Idea** → **PM (Brief+PRD)** → **PO (Epics+Stories)** → **Architect (Design)** → **Developer (Code)** → **QA (Tests)** → **MVP**

## Framework Implementation

**Agent Structure**: Each role has dedicated agent with commands mapping to tasks
**Task Structure**: Procedural workflows with inline template/data references
**Template Structure**: Markdown files with `{{variables}}` for dynamic content
**Data Structure**: Standards, principles, and technical specifications
