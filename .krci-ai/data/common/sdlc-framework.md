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
| Role | ID | Outputs | Dependencies | Key Actions |
|------|----|---------|--------------|-----------|
| Product Manager | `pm-v1` | Project Brief, PRD, Roadmap | None (root artifacts) | Market research, strategic vision, requirements validation |
| Project Manager | `prm-v1` | Project Charter, SOW, Project Plan, Risk Register | Project Brief, PRD | Project planning, execution oversight, risk management, stakeholder alignment |
| Business Analyst | `ba-v1` | Refined PRD, Workflows, Business Rules | PRD, Stakeholder inputs | Requirements gathering, workflow design, acceptance criteria, process mapping |
| Product Owner | `po-v1` | Epics, Stories | PRD | Backlog prioritization, sprint planning, story refinement |
| Architect | `architect-v1` | Architecture Documents | PRD, Epics | System design, technical feasibility, architecture patterns |
| Developer | `dev-v1` | Code, Implementation | Stories, Architecture | Feature implementation, code reviews, deployment configs |
| Go Developer | `go-dev-v1` | Go Code, Implementation | Stories, Architecture | Go-specific implementation, testing, performance optimization |
| QA Engineer | `qa-v1` | Test Results, Quality Reports | Stories, Code | Quality validation, test execution, bug reporting |
| Technical Writer | `tw-v1` | Documentation, Media Artifacts | All artifacts | Documentation review, content improvement, presentation enhancement |
| Product Marketing Manager | `pmm-v1` | Marketing Materials, GTM Strategy | PRD, MVP | Go-to-market strategy, sales enablement, launch campaigns |
</roles>

## Artifact Flow

```text
Project Brief → PRD → Epics → Stories → Code → Tests → MVP → Marketing
                  ↓             ↓
              Architecture ← → Code
```

**Detailed Flow with Traceability**:

1. **PM creates Project Brief** → defines business vision, success metrics
2. **PM creates PRD** → references Project Brief, defines BR1, BR2, NFR1, NFR2...
3. **PRM creates Project Charter** → references Project Brief + PRD for scope
4. **BA refines PRD** → adds workflows, detailed acceptance criteria
5. **PO creates Epics** → references specific PRD requirements (BR1, NFR2...)
6. **Architect creates Architecture** → addresses PRD + Epic technical requirements
7. **PO creates Stories** → breaks down Epics, references Architecture for tasks
8. **Dev/Go-Dev implements Code** → fulfills Story acceptance criteria + Architecture specs
9. **QA creates Test Results** → validates Story acceptance criteria against Code
10. **PMM creates Marketing** → leverages PRD positioning + MVP demonstration

IMPORTANT: Dependency Rules:

- Project Brief: No dependencies (root artifact)
- PRD: Requires Project Brief approval
- Epic: Requires PRD completion, references specific BR/NFR requirements
- Story: Requires Epic definition, maps to implementation tasks
- Architecture: Requires PRD + Epic context for technical design
- Code: Requires Stories + Architecture for implementation guidance
- Tests: Requires Stories + Code for validation
- Marketing: Requires PRD + MVP for go-to-market strategy

## Artifact Definitions

| Artifact | Creator | Purpose | Contains | File Location |
|----------|---------|---------|----------|---------------|
| **Project Brief** | PM | Project vision & strategy | Problem statement, success metrics, constraints | `/docs/prd/project-brief.md` |
| **PRD** | PM | Product requirements | Business requirements (BR1, BR2...), system requirements (NFR1, NFR2...) | `/docs/prd/prd.md` |
| **Project Charter** | PRM | Project scope & authorization | Objectives, scope, stakeholders, success criteria | `/docs/prd/project-charter.md` |
| **Epic** | PO | High-level feature definition | Feature description, acceptance criteria, story breakdown | `/docs/epics/epic-{number}-{slug}.md` |
| **Story** | PO | User requirement with implementation | User story, acceptance criteria, tasks, deployment info | `/docs/stories/{epic_number}.{story_number}.story.md` |
| **Architecture** | Arch | System design | Technical specifications, patterns, decisions | `/docs/architecture/*.md` |
| **Code** | Dev/Go-Dev | Implementation | Source code, configs, deployment manifests | Repository root & subdirs |
| **Test Results** | QA | Quality validation | Test execution results, defect reports, metrics | `/docs/tests/test-results-*.md` |
| **Marketing Materials** | PMM | Go-to-market strategy | Pitch decks, sales enablement, campaigns | `/docs/marketing/*.md` |

## Directory Structure

```bash
{project_root}/
├── docs/                           # All SDLC artifacts
│   ├── prd/                        # PM/PRM: Strategic documents
│   │   ├── project-brief.md        # Vision & strategy (PM)
│   │   ├── prd.md                  # Business/system requirements (PM)
│   │   └── project-charter.md      # Project scope & authorization (PRM)
│   ├── epics/                      # PO: High-level features
│   │   └── epic-{number}-{slug}.md # e.g., epic-1-kuberocketai-baseline.md
│   ├── stories/                    # PO: User requirements with tasks
│   │   └── {epic_number}.{story_number}.story.md    # e.g., 01.01.story.md
│   ├── architecture/               # Architect: System design
│   │   ├── adr/                    # Architecture Decision Records
│   │   ├── 01-introduction.md      # System overview
│   │   ├── 02-high-level-architecture.md
│   │   └── [other numbered sections]
│   ├── tests/                      # QA: Quality validation
│   │   └── test-results-*.md       # Test execution results
│   └── marketing/                  # PMM: Go-to-market materials
│       └── {campaign}-{type}.md    # e.g., launch-pitch-deck.md
├── .krci-ai/                       # Framework assets
│   ├── agents/                     # WHO: Role definitions (YAML files)
│   ├── tasks/                      # WHAT: Procedural workflows (Markdown)
│   ├── templates/                  # HOW: Output formatting (Markdown with {{variables}}) (can have subfolders)
│   ├── data/                       # REFERENCE: Standards & guidelines (can have subfolders)
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

**Enforcement**: All templates must be completed with no `{{variable}}` placeholders. Dependencies must be satisfied before proceeding.

**Agent Validation Checklist**:
- [ ] All required `{{variables}}` filled
- [ ] Dependencies exist and approved
- [ ] Output follows template format
- [ ] Quality standards met
- [ ] Next role handoff requirements satisfied
</quality_gates>

<handoff_points>

1. **Idea → PM**: Market research and validation
2. **PM → BA**: PRD creation and refinement
3. **BA → PO**: Requirements to backlog items
4. **PO → Arch**: Technical feasibility assessment
5. **Arch → Dev**: Implementation guidance
6. **Dev → QA**: Quality validation
7. **QA → MVP**: Deployment readiness
8. **PM → PMM**: Go-to-market strategy development
9. **MVP → PMM**: Product demonstration and marketing material creation

**Validation**: Each handoff requires completion of upstream artifacts and approval at quality gates.
</handoff_points>

<common_issues>

| Issue | Resolution | Agent Action |
|-------|------------|--------------|
| Epic creation blocked | PRD incomplete | Return to Product Manager for PRD completion |
| Story blocked by Architecture | Architecture not defined | Complete Architecture Document before Story creation |
| Implementation blocked | Story criteria unclear | Return to Product Owner for Story refinement |
| Testing blocked | Code incomplete | Ensure Developer completes all Story requirements |
| Template variables unfilled | Validation failed | Complete all `{{variables}}` using referenced templates |

</common_issues>

<agent_workflow>

**How Agents Use This Framework**:

1. **Identify Your Role**: Find your role in the `<roles>` table, note your ID, outputs, dependencies
2. **Check Dependencies**: Verify all required upstream artifacts exist and are approved at quality gates
3. **Load Templates**: Use `.krci-ai/templates/{artifact-template}.md` for output format
4. **Execute Task**: Follow `.krci-ai/tasks/{task-name}.md` workflow instructions
5. **Validate Output**: Ensure no `{{variables}}` remain unfilled, all dependencies referenced
6. **Quality Gate**: Confirm artifact meets quality standards before handoff to next role

**Path Resolution Examples**:
- Templates: `{project_root}/.krci-ai/templates/story.md`
- Tasks: `{project_root}/.krci-ai/tasks/create-story.md`
- Data: `{project_root}/.krci-ai/data/coding-standards.md`
- Override: `{project_root}/.krci-ai/local/templates/story.md` (takes precedence)

</agent_workflow>

<local_override>
Local Override System: Components in `.krci-ai/local/` automatically take precedence over global components:

- Local Tasks: Project-specific workflows override global tasks with same filename
- Local Templates: Custom output formats override global templates with same filename
- Local Data: Project-specific standards override global data with same filename
- Directory Restriction: Only `tasks/`, `templates/`, `data/` folders allowed in `.krci-ai/local/`
</local_override>

## Template Variables

All templates use `{{variable_name}}` format for dynamic content:

- **Required fields**: Must be populated, no empty `{{}}` allowed
- **Dependencies**: Reference parent artifacts (Epic references PRD requirements)
- **Traceability**: Link requirements to implementation using identifiers

**Examples**:
- `{{epic_title}}` → "User Authentication System"
- `{{br_references}}` → "BR1, BR3, BR5" (Business Requirements from PRD)
- `{{nfr_references}}` → "NFR2, NFR4" (Non-Functional Requirements from PRD)
- `{{acceptance_criteria}}` → Specific testable conditions
- `{{dependency_artifacts}}` → "epic-1-user-auth.md, 02-high-level-architecture.md"

<success_flow>
Idea → PM (Brief+PRD) → BA (Analysis) → PO (Epics+Stories) → Architect (Design) → Developer/Go Developer (Code) → QA (Tests) → MVP → PMM (Marketing)
</success_flow>

<implementation>
- Agent Structure: YAML files defining identity, commands, principles, and task references
- Task Structure: Markdown workflows with inline template/data references and step-by-step instructions
- Template Structure: Markdown files with `{{variables}}` for dynamic content and inline guidance
- Data Structure: Markdown files with standards, principles, and technical specifications

Agent Requirements:

```yaml
agent:
  identity:
    name: "Role Name"
    id: "role-id-v1"
    role: "Role description"
    goal: "Primary objective"
  commands:
    help: "Show available commands"
    chat: "(Default) Role consultation"
    exit: "Exit persona"
    create-{artifact}: "Create artifact by running the task {task-name}"
  tasks:
    - "./.krci-ai/tasks/{task-name}.md"
  principles:
    - "Scope boundaries and redirections"
    - "Quality standards and validation"
```

</implementation>
