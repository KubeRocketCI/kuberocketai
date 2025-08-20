# 13. SDLC Framework

This lightweight SDLC framework provides a structured approach for AI agents to collaborate in software development. The framework is filesystem-based, natural language oriented, and designed for agent-to-agent communication through well-defined artifacts.

## 13.1 Core Principles

- **Filesystem-First**: All artifacts stored in filesystem, no APIs or events
- **Agent Discovery**: Agents know where to find artifacts through organized directory structure and naming conventions
- **Natural Language**: Human-readable artifacts that agents can process
- **Simple Dependencies**: Clear but minimal artifact relationships
- **Validation**: Static checkers ensure artifact quality, and filesystem integrity

## 13.2 Roles

### Core Roles

| Role | ID | Responsibilities | Input | Output |
|------|----|--------------------|--------|--------|
| **Product Owner** | `product-owner` | Define product vision and goals, manage and prioritize backlog, clarify requirements, make acceptance decisions, and align features to business value | PRD, roadmap, feedback | Product roadmap, prioritized backlog, approved user stories, acceptance criteria |
| **Project Manager** | `project-manager` | Lead the entire project lifecycle: initiate, plan, execute, monitor, control, and close projects per PMBoK. Manage scope, schedule, resources, risk, communication, and documentation to ensure successful delivery | Project Charter, SoW, requirements, risks, feedback | Comprehensive project plans, schedules, WBS, risk and status reports, closure documents |
| **Business Analyst** | `business-analyst` | Elicit and analyze requirements, map business processes, identify gaps, propose improvements, and bridge communication between stakeholders and technical teams | Stakeholder inputs, business needs, system analysis | Detailed product requirements document, workflows, gap analysis, business cases |
| **Architect** | `architect` | Design system and solution architecture, ensure technical feasibility and scalability, maintain alignment to standards, and document patterns and interfaces | Requirements, use cases, constraints | Architecture diagrams, technical specifications, integration blueprints |
| **Developer** | `developer` | Implement, integrate, and maintain software components and features based on design and requirements. Perform code reviews, resolve issues, and collaborate throughout the SDLC | User stories, technical specifications, architecture | Source code, pull requests, deployed software, resolved issues |
| **QA Engineer** | `qa-engineer` | Design and execute test plans, validate requirements and functionality, identify and report defects, ensure solution quality and compliance | Code, requirements, test cases | Test reports, defect logs, QA sign-offs, quality metrics |
| **DevOps Engineer** | `devops-engineer` | Manage CI/CD pipelines, automate deployments, maintain infrastructure, monitor systems, ensure scalability and reliability, and support operational excellence | Codebase, architecture, deployment requirements | Deployment scripts, infrastructure as code, monitoring dashboards, release documentation |

### Role Combinations

Roles can be merged in compact teams:

- **PO + PM**: Product Owner also manages project delivery, planning, and stakeholder communication.
- **Dev + QA**: Developer handles coding and basic testing/unit tests.
- **Arch + DevOps**: Architect oversees both solution design and infrastructure/CI/CD decisions.

## 13.3 Artifact System

### Artifact Types

**Format**: All artifacts are **Markdown files** (.md) with YAML frontmatter for metadata. This ensures human readability and natural language processing by AI agents.

#### Artifact Relationship Diagram

The SDLC framework creates a structured dependency chain where each artifact builds upon previous work, ensuring traceability from initial concept to delivered product. Ideas flow through strategic planning (PRD), feature definition (Epics), user requirements (Stories), and implementation tasks, while technical artifacts like Architecture guide the development process. Code implementation draws from multiple sources including Stories, Tasks, and Architecture documents to ensure comprehensive coverage. Quality validation through Testing and infrastructure setup via Deployment configurations complete the cycle before MVP delivery. This dependency model ensures that every artifact has clear inputs and produces valuable outputs for downstream activities.

The following diagram shows how SDLC artifacts depend on each other, forming a clear hierarchy from vision to implementation:

```mermaid
graph TD
    Idea([💡 Idea]) --> Brief[📄 Project Brief<br/>Vision & Strategy]
    Brief --> PRD[📋 PRD<br/>Product Requirements]
    PRD --> Epic[🎯 Epic<br/>High-level Feature]
    Epic --> Story[📖 Story<br/>User Requirement<br/>+ Implementation Tasks<br/>+ Deployment Info]

    PRD --> Arch[🏗️ Architecture<br/>System Design]
    Epic --> Arch
    Story --> Code[💻 Code<br/>Implementation<br/>+ Deployment Config]
    Arch --> Code

    Code --> Test[🧪 Test Result<br/>Quality Validation]
    Story --> Test

    Code --> MVP[🎉 MVP Delivered]
    Test --> MVP

    style Idea fill:#e1f5fe,stroke:#333,color:#000
    style Brief fill:#f0e68c,stroke:#333,color:#000
    style PRD fill:#f3e5f5,stroke:#333,color:#000
    style Epic fill:#e8f5e8,stroke:#333,color:#000
    style Story fill:#fff3e0,stroke:#333,color:#000
    style Arch fill:#e0f2f1,stroke:#333,color:#000
    style Code fill:#f1f8e9,stroke:#333,color:#000
    style Test fill:#fff9c4,stroke:#333,color:#000
    style MVP fill:#ffebee,stroke:#333,color:#000
```

Artifact Dependencies Explained:

- **Project Brief**: Root artifact defining project vision, problem statement, and success metrics (no dependencies)
- **PRD**: Product requirements derived from project vision (depends on Project Brief)
- **Epic**: High-level features derived from PRD requirements (depends on PRD). Files should follow the naming convention `{epic_number}-epic-{slug}.md`, e.g., `01-epic-user-authentication.md`.
- **Story**: Breaks down Epic into user-focused requirements with implementation tasks and deployment info (depends on Epic). Files should follow the naming convention `{epic_number}.{story_number}.story.md`, e.g., `01.01.story.md`.
- **Architecture**: Technical design informed by PRD and Epics (depends on PRD, Epic)
- **Code**: Implementation with deployment configurations guided by Stories and Architecture (depends on Story, Architecture)
- **Test Result**: Quality validation of Stories and Code (depends on Story, Code)
- **MVP**: Final deliverable combining all artifacts (depends on Code, Test)

#### Consolidated Artifact Reference

| Artifact Name | Path | Owner | Depends On   | Next Documents | Purpose |
|---------------|------|-------|--------------|----------------|---------|
| **Project Brief** | `/docs/prd/` | `product-manager` | None | PRD | Define project vision, problem statement, target users, and success metrics |
| **PRD** | `/docs/prd/` | `product-manager` | Project Brief | Epic, Architecture | Define product vision, requirements, and success criteria |
| **Architecture Document** | `/docs/architecture/` | `architect` | PRD, Epic | Code | Comprehensive system design with baseline/target architecture, technical decisions, and Epic/Story implementation guidance |
| **Epic** | `/docs/epics/` | `product-owner` | PRD | Story, Architecture | High-level feature or business objective. Files named as `{epic_number}-epic-{slug}.md` for consistency. |
| **Story** | `/docs/stories/` | `product-owner` | Epic | Code, Test | User-focused requirement with acceptance criteria and implementation tasks. Files named as `{epic_number}.{story_number}.story.md` using two-digit format for proper sorting and Epic-Story traceability (e.g., `01.01.story.md`, `01.02.story.md`). |
| **Code** | `-` | `developer` | Story, Architecture | Test, MVP | Implementation artifacts following coding standards |
| **Test Result** | `/docs/tests/` | `qa-engineer` | Story, Code | MVP | Quality validation with test execution results and metrics |

### Artifact Format

All artifacts are **Markdown files** (.md) with natural language content that agents can process and humans can review.

## 13.4 Business Process Flow

### Primary Project Delivery Flow

```mermaid
flowchart TD
    Idea["Idea / Opportunity"]
    Project_Manager["Project Manager"]
    Business_Analyst["Business Analyst"]
    Product_Owner["Product Owner"]
    Architect["Architect"]
    Developer["Developer"]
    QA_Engineer["QA Engineer"]
    DevOps_Engineer["DevOps Engineer"]
    MVP["MVP Delivered"]

    Idea -->|Business Case / Requirements| Project_Manager
    Project_Manager -->|Project Charter / Scope of Work| Business_Analyst
    Project_Manager -->|Project Plan / Schedule| Product_Owner
    Project_Manager -->|Status Reports / Risk Register| Product_Owner

    Business_Analyst -->|Refined Requirements| Product_Owner
    Product_Owner -->|Prioritized Backlog| Architect
    Product_Owner -->|Epic & Story Breakdown| Architect
    Architect -->|Architecture Documents| Developer
    Developer -->|Code + Pull Requests| QA_Engineer
    Developer -->|Test Results| QA_Engineer
    QA_Engineer -->|Test Validation| DevOps_Engineer
    QA_Engineer -->|Feedback| Developer
    DevOps_Engineer -->|CI/CD + Infrastructure| MVP
    MVP -->|Feedback| Product_Owner
    MVP -->|Feedback + Metrics| Project_Manager
```

### Handoff Points

1. **Idea → Project Manager**: Define project business case, scope, initial requirements, feasibility
2. **Project Manager → Business Analyst**: Handoff Project charter / Scope of Work (SoW) for detailed PRD creation and refinement
3. **Business Analyst → Product Owner**: Translate refined requirements to backlog items for prioritization and planning
4. **Product Owner → Architect**: Deliver prioritized backlog for technical feasibility assessment, solution design, and architectural planning
5. **Architect → Developer**: Provide architecture documents, implementation guidance, and acceptance criteria for development
6. **Developer → QA Engineer**: Handoff developed features/codebase for validation against requirements and quality standards
7. **QA Engineer → DevOps Engineer**: Transfer validated code and test artifacts for deployment configuration and production readiness
8. **DevOps Engineer → MVP**: Deploy solution to live environment/system and confirm successful release

### Quality Gates

Every Quality Gate has the appropriate artifact(s).

| Quality Gate                    | Artifact(s)                                 | Gate Description                                               |
|---------------------------------|-----------------------|---------------------|
| Project Brief/Charter Approval   | Project Charter                             | Business case and strategy validated; project officially initiated |
| Requirements/SoW Approval        | PRD, Scope of Work         | Requirements complete, aligned, and signed off                |
| Architecture Review              | Architecture diagrams, technical specs      | Solution design and feasibility approved                      |
| Code Review                      | Source code, pull requests                  | Implementation adheres to standards; passes peer review        |
| Test Validation                  | Test reports, QA sign-off                   | Product meets quality, functionality, and compliance criteria  |
| Deployment Readiness             | Release checklists, deployment scripts, monitoring plan | Solution meets production criteria; ready for go-live          |

## 13.5 Configuration

### Framework Structure

The SDLC Framework follows a simple directory structure:

```bash
/docs/
├── prd/                          # Product vision & requirements
│   ├── project-brief.md          # Project vision & strategy
│   ├── prd.md                    # Product requirements
│   └── roadmap.md # Implementation roadmap & timeline
├── epics/                        # High-level features
│   └── 01-epic-user-authentication.md # Epic file named as {epic_number}-epic-{slug}.md
├── stories/                      # User stories with implementation tasks and deployment info
│   └── 01.01.story.md            # Story file named as {epic_number}.{story_number}.story.md
├── architecture/                 # Complete system design & documentation
│   ├── adr/                      # Architecture Decision Records
│   │   └── 001-example-decision.md # ADR files for architectural decisions
│   ├── 01-introduction.md        # System overview & scope
│   ├── 02-high-level-architecture.md # Core system design
│   ├── 03-tech-stack.md          # Technology choices & rationale
│   ├── 04-data-models.md         # Data structures & relationships
│   ├── 05-components.md          # System components & interfaces
│   ├── 06-distribution-and-release-management.md # Deployment architecture
│   ├── 07-error-handling-strategy.md # Error management approach
│   ├── 08-security.md            # Security architecture & controls
│   ├── 09-coding-standards.md    # Development standards & practices
│   ├── 11-source-tree.md         # Code organization structure
│   ├── 12-nfr.md                 # Non-functional requirements
│   ├── 13-sdlc-framework.md      # Development process framework
│   └── 99-appendix-sad.md        # Software Architecture Document appendix
└── tests/                        # Quality validation
    └── test-results-001.md
```

## 13.6 Artifact Creation

### Role Responsibilities

Each SDLC role is responsible for creating specific artifacts:

| SDLC Role          | Primary Responsibilities                                                                         | Artifact Outputs                                 |
|--------------------|-------------------------------------------------------------------------------------------------|--------------------------------------------------|
| **Project Manager**| Project initiation, planning, execution, monitoring, closure; manages scope, schedule, risk, communication | Project Charter, Project Brief, Scope of Work (SoW), Project Plan, Status Report, Risk Register |
| **Product Owner**  | Backlog management, story definition, prioritization, stakeholder alignment; defines product requirements                      | PRD (Product Requirements Document), Epics, User Stories, Prioritized Backlog, Acceptance Criteria |
| **Business Analyst**| Requirements analysis, workflow/process modeling, documentation                                | Refined Requirements, Workflows, Use Cases       |
| **Architect**      | System design, solution architecture, technical decisions, documentation                        | Architecture Documents, Technical Specifications  |
| **Developer**      | Implementation, coding, code review, defect resolution                                          | Source Code, Implemented Features, Pull Requests  |
| **QA Engineer**    | Test planning and execution, quality assurance, defect reporting                                | Test Results, QA Reports, Sign-off Documentation  |
| **DevOps Engineer**| Infrastructure automation, CI/CD pipelines, deployment, monitoring                              | Deployment Scripts, Infrastructure as Code, Monitoring Dashboards |

## 13.7 End-to-End Scenario: Idea to Code

This section demonstrates the complete artifact flow from initial idea through to implementation, showing how each role contributes and transfers artifacts to the next stage of the SDLC process.

### Artifact Transfer Sequence

The following sequence diagram illustrates how artifacts are created and transferred between roles in a typical end-to-end scenario:

```mermaid
sequenceDiagram
 participant Stakeholder
 participant PM as Project Manager
 participant BA as Business Analyst
 participant PO as Product Owner
 participant Arch as Architect
 participant Dev as Developer
 participant QA as QA Engineer

 Note over Stakeholder: Initial Business Idea

 Stakeholder->>PM: Business Opportunity & Requirements
 PM->>PM: Project Initiation & Planning
 PM->>PM: Create Project Charter & Project Brief

 PM->>PO: Handover Business Context & Objectives
 PO->>PO: Create PRD (Product Requirements Document)
 PO->>BA: PRD for Requirements Analysis
 BA->>BA: Refine PRD with Detailed Requirements and Workflows
 BA->>PO: Refined PRD

 PO->>PO: Create Epics (based on PRD)
 PO->>Arch: Refined PRD + Epics for Technical Review
 Arch->>Arch: Create Architecture Documents

 PO->>PO: Create Stories (based on Epics)
 PO->>Dev: Stories + Architecture Documents

 Dev->>Dev: Implement Code (based on Stories + Architecture)
 Dev->>QA: Code + Stories for Testing

 QA->>QA: Create Test Results (based on Stories + Code)
 QA->>Dev: Test Results

 alt Tests Pass
 QA->>PM: Test Results (Success)
 Note over PM: MVP Ready for Delivery
 else Tests Fail
 QA->>Dev: Test Results (Issues Found)
 Dev->>Dev: Fix Code Issues
 Dev->>QA: Updated Code
 end
```

This scenario demonstrates how artifacts evolve and transform:

1. **Stakeholder Idea → Project Charter/Brief**: Project Manager captures business needs and strategic goals
2. **Project Charter/Brief → PRD**: Product Owner defines product requirements based on business context
3. **PRD → Refined PRD**: Business Analyst adds detailed workflows and acceptance criteria
4. **Refined PRD → Epics**: Product Owner breaks down requirements into high-level features
5. **Refined PRD + Epics → Architecture Documents**: Architect designs systems to support features
6. **Epics → Stories**: Product Owner creates actionable requirements and tasks
7. **Stories + Architecture → Code**: Developer implements features per specifications
8. **Stories + Code → Test Results**: QA Engineer validates implementation against requirements
9. **Validated Code + Test Results → MVP**: Ready for delivery

### Key Success Factors

- **Clear Handoffs**: Each role has defined inputs and outputs
- **Artifact Traceability**: Every artifact traces back to business requirements
- **Quality Gates**: Testing validates implementation against original stories and acceptance criteria
- **Iterative Refinement**: Feedback loops enable continuous improvement
- **Documentation Trail**: Complete audit trail from idea to deployment

## 13.8 Artifact Templates

The templates are already implemented as a part of framework, so this section has been removed.

## 13.9 Framework Implementation Requirements

Based on the KubeRocketAI 4-component architecture, the following implementation structure supports the SDLC artifact flow:

### Required Components

**Agent Structure** (WHO - SDLC Roles):

```yaml
agent:
  identity:
    name: "Project Manager"
    id: "project-manager-v1"
    role: "Strategic product direction"
    goal: "Define product vision and roadmap"
  commands:
    help: "Show available commands"
    chat: "(Default) Product consultation"
    exit: "Exit persona command"
    create-brief: "Create project brief"
    create-prd: "Create product requirements document"
  tasks:
    - "./.krci-ai/tasks/create-project-brief.md"
    - "./.krci-ai/tasks/create-prd.md"
```

**Task Structure** (WHAT - SDLC Procedures):

```markdown
# Task: Create Project Brief

## Description
Create comprehensive project brief defining vision, problem statement, and success metrics

## Instructions
1. Analyze stakeholder requirements and market opportunity
2. Define problem statement and target users
3. Establish success metrics and constraints
4. Format output using [project-brief-template.md](./.krci-ai/templates/project-brief-template.md)
```

**Template Structure** (HOW - Output Format):

```markdown
# Project Brief: {{project_name}}

## Executive Summary
{{executive_summary}}

## Problem Statement
{{problem_description}}

## Target Users
{{target_users}}

## Success Metrics
{{success_metrics}}
```

**Data Structure** (REFERENCE - Standards):

```yaml
# data/sdlc-standards.yaml
standards:
  artifact_requirements:
    - "All artifacts must be in Markdown format"
    - "Dependencies must be clearly documented"
    - "Acceptance criteria must be testable"
  quality_gates:
    - "PRD approval required before Epic creation"
    - "Architecture review required before implementation"
    - "Test validation required before MVP delivery"
```

### Required Filesystem Structure

```bash
.krci-ai/
├── agents/
│   ├── project-manager.yaml
│   ├── product-owner.yaml
│   ├── architect.yaml
│   ├── developer.yaml
│   └── qa-engineer.yaml
├── tasks/
│   ├── create-project-brief.md
│   ├── create-prd.md
│   ├── create-epic.md
│   ├── create-story.md
│   ├── create-architecture.md
│   └── create-tests.md
├── templates/
│   ├── project-brief-template.md
│   ├── prd-template.md
│   ├── epic-template.md
│   ├── story-template.md
│   ├── architecture-template.md
│   └── test-results-template.md
└── data/
    ├── sdlc-standards.yaml
    ├── coding-standards.yaml
    └── architecture-principles.yaml
```

### Integration Points

- **Agent Commands** map to SDLC artifact creation tasks
- **Task Instructions** follow artifact dependency flow (Brief → PRD → Epic → Story → Code → Test → MVP)
- **Templates** ensure consistent artifact formatting across teams
- **Data Standards** enforce quality gates and validation requirements

This implementation enables AI-assisted SDLC artifact creation while maintaining the dependency flow and quality standards defined in this framework.

## 13.10 Framework Governance

### Quality Gates Enforcement

**Artifact Quality Validation**:

- All templates must be completed with no `{{variable}}` placeholders remaining
- Dependencies must be satisfied before proceeding to next artifact
- Quality criteria in each template must be met

**Escalation Process**:

1. **Artifact Issues**: If artifact is incomplete → Return to owner for completion
2. **Requirement Changes**: If PRD changes → Update dependent artifacts (Epic, Architecture, Story)
3. **Technical Blockers**: If Architecture infeasible → Escalate to Project Manager for scope adjustment

### Change Management

**Requirement Changes**:

- **Minor Changes**: Update affected artifacts and continue
- **Major Changes**: Re-validate entire artifact chain from PRD forward
- **Scope Changes**: Require Project Manager approval and impact assessment

**Version Control**:

- All artifacts tracked through Git version control
- Changes tracked through commit history and directory organization
- Dependencies updated when upstream artifacts change

### Common Issues & Resolutions

| Issue | Cause | Resolution |
|-------|-------|------------|
| **Epic creation blocked** | PRD incomplete or unclear | Return to Product Owner for PRD completion |
| **Story blocked by Architecture** | Architecture not yet defined | Complete Architecture Document before Story creation |
| **Implementation blocked** | Story acceptance criteria unclear | Return to Product Owner for Story refinement |
| **Testing blocked** | Code incomplete | Ensure Developer completes all Story requirements |
| **Agent confusion** | Template variables not filled | Validate all templates using provided checklists |

### Success Metrics

**Framework Effectiveness**:

- **Artifact Completion Rate**: >95% of artifacts completed without rework
- **Dependency Satisfaction**: All dependencies resolved before proceeding
- **Quality Gate Pass Rate**: >90% of artifacts pass quality gates on first review
- **Time to Value**: Consistent delivery timeline from Project Brief to MVP

This governance framework ensures smooth operation while maintaining the lightweight nature of the SDLC process.
