# SDLC Framework Quick Reference

Purpose: This framework enables AI agents to collaborate in software development through structured, filesystem-based artifacts. Agents reference this document to understand their role, dependencies, and where to find required templates and standards.

## Framework Principles

- Filesystem-First: All artifacts stored as markdown files, no APIs
- Agent Discovery: Agents find artifacts through organized directory structure
- Natural Language: Human-readable artifacts that agents can process
- Clear Dependencies: Each artifact builds on previous work
- Inline References: Use `[filename](path/to/file)` markdown links for all references
- Local Override: Project-specific components in `.krci-ai/local/` automatically override global ones

<roles>
| Role | Outputs | Dependencies |
|------|---------|--------------|
| Product Manager | Project Brief, PRD, Roadmap | None (root artifacts) |
| Product Owner | Epics, Stories | PRD |
| Business Analyst | Refined PRD, Workflows | PRD, Stakeholder inputs |
| Architect | Architecture Documents | PRD, Epics |
| Developer | Code, Implementation | Stories, Architecture |
| Go Developer | Go Code, Implementation | Stories, Architecture |
| QA Engineer | Test Results, Quality Reports | Stories, Code |
| Product Marketing Manager | Marketing Materials, GTM Strategy | PRD, MVP |
</roles>

## Artifact Flow

```text
Project Brief → PRD → Epics → Stories → Code → Tests → MVP → Marketing
                  ↓             ↓
              Architecture ← → Code
```

IMPORTANT: Dependency Rules:

- Project Brief: No dependencies (root artifact)
- PRD: Requires Project Brief approval
- Epic: Requires PRD completion, references specific BR/NFR requirements
- Story: Requires Epic definition, maps to implementation tasks
- Architecture: Requires PRD + Epic context for technical design
- Code: Requires Stories + Architecture for implementation guidance
- Tests: Requires Stories + Code for validation
- Marketing: Requires PRD + MVP for go-to-market strategy

## Directory Structure

```bash
{project_root}/
├── docs/                           # All SDLC artifacts
│   ├── prd/                        # PM: Project Brief, PRD
│   │   ├── project-brief.md        # Vision & strategy
│   │   └── prd.md                  # Business requirements (BR1, BR2...), system requirements (NFR1, NFR2...)
│   ├── epics/                      # PO: High-level features
│   │   └── epic-{number}-{slug}.md # e.g., epic-1-kuberocketai-baseline.md
│   ├── stories/                    # PO: User requirements with tasks
│   │   └── {epic_number}.{story_number}.story.md    # e.g., 01.01.story.md
│   ├── architecture/               # Architect: System design
│   │   ├── adr/                    # Architecture Decision Records
│   │   ├── 01-introduction.md      # System overview
│   │   ├── 02-high-level-architecture.md
│   │   └── [other numbered sections]
│   └── marketing/                  # PMM: Go-to-market materials
│       └── {campaign}-{type}.md    # e.g., launch-pitch-deck.md
├── .krci-ai/                       # Framework assets
│   ├── agents/                     # WHO: Role definitions (YAML files)
│   ├── tasks/                      # WHAT: Procedural workflows (Markdown)
│   ├── templates/                  # HOW: Output formatting (Markdown with {{variables}})
│   ├── data/                       # REFERENCE: Standards & guidelines
│   └── local/                      # CODEBASE-SPECIFIC: Override global components
│       ├── tasks/                  # Custom codebase tasks (override global)
│       ├── templates/              # Custom codebase templates (override global)
│       └── data/                   # Custom codebase data (override global)
```

<quality_gates>

1. Project Brief Approval → Enables PRD creation
2. PRD Approval → Enables Epic/Architecture creation
3. Architecture Review → Enables implementation
4. Code Review → Enables testing
5. Test Validation → Enables MVP delivery
</quality_gates>

<local_override>
Local Override System: Components in `.krci-ai/local/` automatically take precedence over global components:

- Local Tasks: Project-specific workflows override global tasks with same filename
- Local Templates: Custom output formats override global templates with same filename
- Local Data: Project-specific standards override global data with same filename
- Directory Restriction: Only `tasks/`, `templates/`, `data/` folders allowed in `.krci-ai/local/`
</local_override>

## Template Variables

All templates use `{{variable_name}}` format for dynamic content:

- Required fields: Must be populated, no empty `{{}}` allowed
- Dependencies: Reference parent artifacts (Epic references PRD requirements)
- Traceability: Link requirements to implementation (BR1, NFR2, etc.)

<success_flow>
Idea → PM (Brief+PRD) → BA (Analysis) → PO (Epics+Stories) → Architect (Design) → Developer/Go Developer (Code) → QA (Tests) → MVP → PMM (Marketing)
</success_flow>

<implementation>
- Agent Structure: YAML files defining identity, commands, principles, and task references
- Task Structure: Markdown workflows with inline template/data references and step-by-step instructions
- Template Structure: Markdown files with `{{variables}}` for dynamic content and inline guidance
- Data Structure: Markdown files with standards, principles, and technical specifications
</implementation>
