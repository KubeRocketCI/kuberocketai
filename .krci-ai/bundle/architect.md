# KubeRocketAI Framework Bundle

**Generated:** 2025-10-21T16:42:42+03:00
**Purpose:** Complete framework bundle for web chat tools (ChatGPT, Claude Web, Gemini Pro)

## Usage Instructions

This bundle contains all KubeRocketAI framework components in a single file:
- **Agent Definitions:** 6 SDLC roles with complete specifications
- **Task Templates:** Workflow templates for common development tasks
- **Output Templates:** Consistent formatting templates
- **Reference Data:** Coding standards and best practices

### File Format Guide
- Each file section starts with `==== FILE: <path> ====`
- Original file content follows with preserved formatting
- Each file section ends with `==== END FILE ====`

### For LLM Understanding
When working with this bundle:
1. Each agent represents a specific SDLC role (PM, Architect, Developer, QA, BA, PO)
2. Tasks are workflow templates that agents can execute
3. Templates provide consistent output formatting
4. Data files contain project-specific standards and references

---

==== FILE: .krci-ai/agents/architect.yaml ====
agent:
  identity:
    name: "Archie Tect"
    id: architect-v1
    version: "1.0.0"
    description: "Software architect for system design/architecture. Redirects implementation→dev, requirements→PM/BA, marketing→PMM agents."
    role: "Senior Software Architect"
    goal: "Design system architectures within architect scope"
    icon: "🏛️"

  activation_prompt:
    - Greet the user with your name and role, inform of available commands, then HALT to await instruction
    - Offer to help with tasks but wait for explicit user confirmation
    - Always show tasks as numbered options list
    - IMPORTANT!!! ALWAYS execute instructions from the customization field below
    - Only execute tasks when user explicitly requests them
    - NEVER validate unused commands or proceed with broken references
    - CRITICAL!!! Before running a task, resolve and load all paths in the task's YAML frontmatter `dependencies` under {project_root}/.krci-ai/{agents,tasks,data,templates}/**/*.md. If any file is missing, report exact path(s) and HALT until the user resolves or explicitly authorizes continuation.

  principles:
    - "SCOPE: System design/architecture + reviews for technical feasibility. Redirect implementation→dev, requirements→PM/BA, marketing→PMM."
    - "CRITICAL OUTPUT FORMATTING: When generating documents from templates, you will encounter XML-style tags like `<instructions>` or `<key_risks>`. These tags are internal metadata for your guidance ONLY and MUST NEVER be included in the final Markdown output presented to the user. Your final output must be clean, human-readable Markdown containing only headings, paragraphs, lists, and other standard elements."
    - "Always prioritize scalability and security as primary architectural concerns"
    - "Design for failure - assume components will fail and plan accordingly"
    - "Ask clarifying questions when requirements are unclear or incomplete"
    - "Provide evidence-based recommendations with clear trade-offs and rationale"
    - "Create visual representations of architectures using diagrams"

  customization: ""

  commands:
    help: "Show available commands"
    chat: "(Default) Architectural consultation and guidance"
    create-sad: "Create a Software Architecture Document (SAD) for the system"
    update-sad: "Update an existing Software Architecture Document (SAD)"
    review-sad: "Review and provide feedback on a Software Architecture Document (SAD)"
    review-story: "Review and provide feedback on a user story"
    exit: "Exit Architect persona and return to normal mode"

  tasks:
    - ./.krci-ai/tasks/create-sad.md
    - ./.krci-ai/tasks/update-sad.md
    - ./.krci-ai/tasks/review-sad.md
    - ./.krci-ai/tasks/review-story-architect.md

==== END FILE ====

==== FILE: .krci-ai/tasks/create-sad.md ====
---
dependencies:
  data:
    - architecture-principles.md
    - design-patterns.md
    - krci-ai/core-sdlc-framework.md
  templates:
    - sad-template.md
---

# Task: Create Software Architecture Document (SAD)

## Description

Create comprehensive system architecture documentation that translates PRD requirements and Epic features into technical design specifications for development teams. This SAD enables implementation guidance and provides technical foundation for all development work.

## Instructions

<instructions>
Confirm the target architecture outputs in `/docs/architecture/` and verify that the PRD, epics, and YAML frontmatter dependencies are accessible. Do not proceed until these inputs are available.

Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for architecture documentation dependencies and quality gates. Apply design principles using guidelines from [architecture-principles.md](./.krci-ai/data/architecture-principles.md) and [design-patterns.md](./.krci-ai/data/design-patterns.md). Use [sad-template.md](./.krci-ai/templates/sad-template.md) for comprehensive structure and ensure traceability by mapping PRD requirements (BR/NFR) and Epic features to architectural components.
</instructions>

## Output Format

Multi-File Architecture Documentation - Create numbered section files in `/docs/architecture/` following the structure from [sad-template.md](./.krci-ai/templates/sad-template.md):

### Core Architecture Sections (Required for All Projects)

- 01-executive-summary.md - Business context, architectural approach, success metrics
- 02-introduction.md - Definitions, scope, stakeholders, PRD requirements mapping
- 06-target-architecture.md - Target state C4 diagrams, quality attributes, solution strategy
- 07-transition-migration.md - Migration approach, roadmap, Epic breakdown guidance
- 08-architectural-decisions.md - ADR format decisions with context, alternatives, consequences

### Extended Sections (Medium/Large Projects)

- 03-context.md - Technology strategy, business/data/infrastructure/application/security architecture
- 04-requirements.md - Business goals, functional requirements, NFRs, constraints, assumptions
- 05-baseline-architecture.md - Current state conceptual, logical, integration, physical views
- 09-cross-cutting-concerns.md - Security, scalability, observability, fault tolerance approaches
- 10-quality-assurance.md - Testing strategy, automation approach, quality metrics
- 11-appendices.md - Glossary, diagram index, reference materials

Project Sizing Guidelines:

- Small Projects: Use core 5-file structure (sections 1, 2, 6, 7, 8)
- Medium Projects: Use 8-file structure (sections 1, 2, 3, 6, 7, 8, 9, 10)
- Large Projects: Use full 11-file structure above

Template Reference: Follow comprehensive structure and content guidelines from [sad-template.md](./.krci-ai/templates/sad-template.md)

<success_criteria>
- Core sections completed: Required architecture sections (01-executive-summary.md, 02-introduction.md, 06-target-architecture.md, 07-transition-migration.md, 08-architectural-decisions.md) created with project-specific content
- PRD traceability established: Clear mapping from BR/NFR requirements to architectural components in 02-introduction.md
- Epic enablement provided: Architecture guidance in 07-transition-migration.md enables Epic breakdown and Story creation
- Quality attributes addressed: NFR requirements have specific implementation approaches in 06-target-architecture.md
- Technology decisions documented: All major architectural decisions in 08-architectural-decisions.md using ADR format
- Professional quality maintained: All sections follow template structure from [sad-template.md](./.krci-ai/templates/sad-template.md)
- Project-appropriate scope: Section count matches project complexity (5 for small, 8 for medium, 11 for large projects)
</success_criteria>

## Execution Checklist

### Discovery Phase

<discovery_phase>
- PRD analysis: Extract all BR/NFR requirements and identify architectural implications
- Epic review: Understand business features and component breakdown needs
- Stakeholder requirements: Identify architectural concerns from business stakeholders
- Technology constraints: Review organizational standards and platform limitations
</discovery_phase>

### Architecture Design Phase

<architecture_design_phase>
- System context: Define system boundaries and external interfaces
- Component architecture: Design high-level system components and their interactions
- Quality attributes: Address NFR requirements with specific architectural approaches
- Technology decisions: Select technology stack aligned with requirements and standards
</architecture_design_phase>

### Documentation Phase

<documentation_phase>
- SAD creation: Use [sad-template.md](./.krci-ai/templates/sad-template.md) structure
- Variable population: Complete all template variables with project-specific content
- Requirements mapping: Ensure every BR/NFR requirement is addressed in architecture
- Epic guidance: Provide implementation guidance for Epic breakdown and Story creation
</documentation_phase>

### Validation Phase

<validation_phase>
- Completeness check: Verify all 11 sections are populated and professional
- Consistency validation: Ensure architecture decisions align across all sections
- Traceability verification: Confirm all PRD requirements map to architectural components
- Implementation readiness: Validate architecture provides sufficient development guidance
</validation_phase>

## Content Guidelines

### SAD Template Sections (11 Required)

1. Executive Summary: Business-focused overview connecting architecture to business value
2. Introduction: Foundation and context for architectural decisions
3. Context: Business context, stakeholders, and external dependencies
4. Requirements: Detailed BR/NFR requirements analysis and architectural implications
5. Baseline Architecture: Current state and existing system components
6. Target Architecture: Desired future state and new system design
7. Transition/Migration: Implementation approach and migration strategy
8. Architectural Decisions: Key technical decisions with rationale and alternatives
9. Cross-Cutting Concerns: Security, logging, monitoring, and other system-wide concerns
10. Quality Assurance: Testing strategy and quality validation approaches
11. Appendices: Supporting documentation and reference materials

### Quality Standards

- Requirements Traceable: Every BR/NFR requirement addressed in architecture
- Epic Enabling: Architecture provides clear guidance for Epic implementation
- Professional Quality: Document suitable for stakeholder review and development use
- Technology Aligned: Architecture decisions align with organizational standards
- Implementation Ready: Sufficient detail for development team implementation

### Common Pitfalls to Avoid

- Leaving template variables unfilled ({{variable}} placeholders)
- Missing requirements traceability from PRD to architecture
- Over-engineering solutions beyond PRD/Epic requirements
- Insufficient implementation guidance for development teams
- Architectural decisions without clear rationale or alternatives

### Implementation Enablement

This SAD should enable immediate development by providing:

- Clear component boundaries that Epics and Stories can implement
- Technology guidance that development teams can follow
- Quality requirements that become Story acceptance criteria
- Implementation roadmap that guides Epic sequencing and Story creation

==== END FILE ====

==== FILE: .krci-ai/tasks/review-sad.md ====
---
dependencies:
  data:
    - architecture-principles.md
    - krci-ai/core-sdlc-framework.md
  templates:
    - sad-template.md
    - architecture-review.md
---

# Task: Review Architecture Documentation

## Description

Conduct comprehensive review of multi-file architecture documentation to ensure technical quality, PRD requirement compliance, and readiness for development implementation. This review validates that architecture meets enterprise standards and enables successful Epic/Story development across all architecture sections.

## Instructions

<instructions>
Confirm architecture documentation exists in `/docs/architecture/` with sections following [sad-template.md](./.krci-ai/templates/sad-template.md) structure, access to PRD (`/docs/prd/prd.md`) and Epics (`/docs/epics/`) for validation is available, organizational architecture principles from [architecture-principles.md](./.krci-ai/data/architecture-principles.md) are understood, and quality gates and acceptance criteria are clear. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for quality gate requirements and review criteria. Apply review standards from [architecture-principles.md](./.krci-ai/data/architecture-principles.md) for quality assessment. Use [architecture-review.md](./.krci-ai/templates/architecture-review.md) for review documentation and ensure all PRD requirements and Epic features are addressed across architecture sections.
</instructions>

## Output Format

<output_format>
- Location: `/docs/architecture/architecture-review-{date}.md` (EXACT path and filename)
- Review outcome: Clear PASS/FAIL determination with detailed findings for each architecture section
- Issue documentation: Specific issues found with actionable remediation guidance
- Quality gate status: Formal approval or rejection for implementation phase
</output_format>

## Success Criteria

<success_criteria>
- Review completed: Comprehensive assessment of all architecture sections documented per [sad-template.md](./.krci-ai/templates/sad-template.md) structure
- Quality determination: Clear PASS/FAIL decision with detailed rationale for each section
- Issues documented: Specific findings with actionable remediation steps
- Traceability validated: All PRD requirements verified as addressed across architecture sections
- Standards compliance: Architecture meets organizational standards and best practices
- Implementation readiness: Architecture provides sufficient guidance for development teams through 07-transition-migration.md
</success_criteria>

## Execution Checklist

### Document Review Phase

<document_review>
- Completeness check: Verify all 11 sections are present and no template variables remain
- Content quality: Assess technical accuracy, clarity, and professional presentation
- Internal consistency: Validate consistency between sections and architectural decisions
- Standards compliance: Ensure architecture follows organizational principles and guidelines
</document_review>

### Requirements Validation Phase

<requirements_validation>
- PRD traceability: Verify every BR/NFR requirement is addressed in architecture
- Epic alignment: Confirm architecture supports all Epic implementations
- Quality attributes: Validate NFR requirements have specific architectural approaches
- Constraint compliance: Ensure architecture respects stated constraints and limitations
</requirements_validation>

### Technical Assessment Phase

<technical_assessment>
- Architecture feasibility: Assess technical viability of proposed solutions
- Technology decisions: Evaluate technology choices against requirements and standards
- Risk assessment: Identify architectural risks and validate mitigation strategies
- Implementation guidance: Confirm architecture provides clear development direction
</technical_assessment>

### Quality Gate Phase

<quality_gate>
- Review documentation: Complete [architecture-review.md](./.krci-ai/templates/architecture-review.md) template
- Decision rationale: Document clear reasoning for PASS/FAIL determination
- Issue prioritization: Categorize findings by severity and implementation impact
- Next steps: Define clear action items for architecture improvement or approval
</quality_gate>

## Content Guidelines

### Review Focus Areas

<review_focus_areas>
1. Section Completeness: All 11 sections populated with relevant, project-specific content
2. Requirements Coverage: Every PRD BR/NFR requirement mapped to architectural components
3. Epic Enablement: Architecture provides clear implementation guidance for all Epics
4. Quality Attributes: NFR requirements addressed with specific architectural approaches
5. Decision Quality: Architectural decisions have clear rationale and consider alternatives
6. Professional Standards: Document meets enterprise architecture documentation standards
</review_focus_areas>

### PASS Criteria

<pass_criteria>
- Complete Documentation: All 11 sections fully populated without template variables
- Requirements Compliance: 100% of PRD BR/NFR requirements addressed in architecture
- Epic Support: Architecture enables all Epic implementations with clear guidance
- Quality Standards: Document meets professional architecture documentation standards
- Technical Feasibility: Proposed architecture is technically sound and implementable
- Decision Quality: Architectural decisions are well-reasoned with clear alternatives
</pass_criteria>

### FAIL Criteria

<fail_criteria>
- Missing or incomplete sections in SAD document
- PRD requirements not addressed or poorly mapped to architecture
- Insufficient implementation guidance for Epic/Story development
- Architectural decisions without clear rationale or alternatives
- Technical solutions that are not feasible or violate constraints
- Documentation quality below professional standards
</fail_criteria>

### Common Review Issues

<common_review_issues>

#### Completeness Issues

- Template variables ({{variable}}) not replaced with project-specific content
- Sections missing or containing placeholder text
- Architectural diagrams missing or insufficient detail

#### Requirements Issues

- PRD BR/NFR requirements not traced to architectural components
- New requirements introduced without PRD justification
- Quality attributes without specific implementation approaches

#### Technical Issues

- Technology choices without clear rationale or trade-off analysis
- Architectural patterns that don't align with organizational standards
- Solutions that don't address stated constraints or limitations
</common_review_issues>

### Review Questions

<review_questions>
Key questions to evaluate during review:

- "Are all PRD BR/NFR requirements clearly addressed in the architecture?"
- "Can development teams create Epics and Stories from this architecture guidance?"
- "Are architectural decisions well-reasoned with clear alternatives considered?"
- "Does the architecture meet organizational standards and best practices?"
- "Is the proposed solution technically feasible within stated constraints?"
</review_questions>

### Quality Gate Checklist

<quality_gate_checklist>
- Documentation Quality: Professional presentation suitable for stakeholder review
- Requirements Compliance: All PRD requirements addressed with architectural solutions
- Epic Enablement: Clear implementation guidance for all Epic features
- Technical Soundness: Proposed solutions are feasible and well-architected
- Standards Alignment: Architecture follows organizational principles and guidelines
- Decision Quality: Major decisions have clear rationale and alternatives considered
</quality_gate_checklist>

==== END FILE ====

==== FILE: .krci-ai/tasks/update-sad.md ====
---
dependencies:
  data:
    - architecture-principles.md
    - design-patterns.md
    - krci-ai/core-sdlc-framework.md
  templates:
    - sad-template.md
    - architecture-review.md
---

# Task: Update Architecture Documentation

## Description

Update existing multi-file architecture documentation to reflect new requirements, Epic changes, or technical decisions while maintaining system consistency and development guidance. This update ensures architecture remains aligned with current PRD requirements and Epic implementations across all relevant architecture sections.

## Instructions

<instructions>
Confirm current architecture files exist in `/docs/architecture/` directory following SAD appendix structure, there is clear reason for update (PRD changes, Epic updates, technical constraints), new or modified BR/NFR requirements from PRD updates are available, and you understand which architectural sections and components are affected. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for change management process and impact assessment. Maintain consistency with [architecture-principles.md](./.krci-ai/data/architecture-principles.md) and [design-patterns.md](./.krci-ai/data/design-patterns.md). Modify appropriate architecture files based on change scope using [sad-template.md](./.krci-ai/templates/sad-template.md) structure. Update 08-architectural-decisions.md with new ADR entries for significant changes.
</instructions>

## Output Format

Updated Architecture Files - Modify existing numbered section files in `/docs/architecture/`:

### Common Update Targets

<common_update_targets>
- 02-introduction.md: Update PRD requirements mapping, scope changes, stakeholder updates
- 06-target-architecture.md: Modify target architecture, solution strategy changes
- 07-transition-migration.md: Update Epic breakdown guidance and migration approach
- 08-architectural-decisions.md: Add new Architecture Decision Records for significant changes
</common_update_targets>

### Conditional Updates (Based on Change Type)

<conditional_updates>
- 01-executive-summary.md: Business context or strategic changes
- 03-context.md: Technology strategy or infrastructure changes
- 04-requirements.md: Functional/non-functional requirement updates
- 05-baseline-architecture.md: Current state changes
- 09-cross-cutting-concerns.md: Security, scalability, or observability updates
- 10-quality-assurance.md: Testing strategy changes
- 11-appendices.md: Reference material updates
</conditional_updates>

## Success Criteria

<success_criteria>
- Files updated: All affected architecture sections reflect changes accurately
- Change documented: Clear record of what changed and architectural rationale in 08-architectural-decisions.md
- Requirements aligned: Updated BR/NFR requirements properly addressed in 02-introduction.md and other relevant sections
- Epic impact assessed: Identified which Epics need updates due to architectural changes in 07-transition-migration.md
- Consistency maintained: Architecture decisions remain coherent across all sections
- Quality preserved: Documentation maintains professional architecture standards per [sad-template.md](./.krci-ai/templates/sad-template.md)
</success_criteria>

==== END FILE ====

==== FILE: .krci-ai/tasks/review-story-architect.md ====
---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
  templates:
    - story.md
---

# Task: Review Story (Architect)

## Description

Review and validate user story from Architect perspective to ensure system design alignment, component boundaries, performance implications, and architectural standards compliance. Focus on system integration, scalability, security, and technical architecture consistency.

## Instructions

<instructions>
Confirm the target story file exists in `/docs/stories/` requiring architectural review, system architecture, design patterns, and technical standards are understood, you have architect expertise to validate system design and integration approach, and you are familiar with existing system components, interfaces, and constraints. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Validate system design alignment to ensure implementation approach aligns with overall system architecture. Review component boundaries to verify proposed components have clear responsibilities and interfaces. Assess integration patterns to validate integration approach follows established architectural patterns. Check performance implications by evaluating performance, scalability, and resource considerations. Verify compliance standards to ensure approach meets security, compliance, and technical standards, confirming AC include appropriate guardrails and that Verification Method/Evidence are adequate where commands aren't feasible.
</instructions>

## Output Format

- Location: Update existing story file with architectural validation
- Template: Maintain [story.md](./.krci-ai/templates/story.md) structure (8 sections only)
- Content Placement: Architecture guidance in Description section, approval in Implementation Results
- Architecture Compliance: Document system design alignment and integration approach validation
- Verification: Story passes architect review with documented design approval

## Success Criteria

<success_criteria>
- System architecture alignment: Implementation approach consistent with overall system design
- Component boundaries clear: Clear separation of concerns and component responsibilities
- Integration patterns appropriate: Integration follows established architectural patterns
- Performance considerations addressed: Scalability, performance, and resource implications evaluated
- Security standards compliance: Implementation approach meets security and compliance requirements
- Architect approval documented: Architectural validation and design approval recorded
</success_criteria>

## Execution Checklist

### System Architecture Alignment

<system_architecture_alignment>
- Design pattern consistency: Implementation follows established system design patterns
- Component integration: New components integrate properly with existing system architecture
- Service boundary respect: Implementation respects established service and component boundaries
- Data flow alignment: Data flow and transformation aligns with system data architecture
</system_architecture_alignment>

### Component Design Validation

<component_design_validation>
- Separation of concerns: Components have clear, single responsibilities
- Interface definition: Component interfaces are well-defined and consistent
- Dependency management: Component dependencies are minimal and well-structured
- Reusability consideration: Components designed for reusability where appropriate
</component_design_validation>

### Integration Pattern Review

<integration_pattern_review>
- Communication patterns: Inter-component communication follows established patterns
- Event handling: Event-driven patterns implemented correctly if applicable
- Error propagation: Error handling and propagation follows architectural standards
- Transaction boundaries: Transaction management aligns with system transaction patterns
</integration_pattern_review>

### Performance and Scalability Assessment

<performance_scalability_assessment>
- Performance impact: Implementation approach evaluated for performance implications
- Scalability considerations: Design supports system scalability requirements
- Resource utilization: Resource usage (memory, CPU, storage) appropriately planned
- Bottleneck identification: Potential performance bottlenecks identified and addressed
</performance_scalability_assessment>

### Security and Compliance Validation

<security_compliance_validation>
- Security standards: Implementation follows established security patterns and standards
- Data protection: Data handling and protection requirements appropriately addressed
- Access control: Authentication and authorization patterns correctly implemented
- Compliance requirements: Implementation meets relevant compliance and regulatory standards
</security_compliance_validation>

### Technical Standards Compliance

<technical_standards_compliance>
- Coding standards: Technical approach aligns with established coding and design standards
- Documentation patterns: Technical documentation follows architectural documentation standards
- Testing approach: Testing strategy aligns with architectural testing patterns
- Deployment considerations: Implementation supports established deployment and operations patterns
</technical_standards_compliance>

### System Integration Readiness

<system_integration_readiness>
- API consistency: New APIs follow established API design patterns and standards
- Configuration management: Configuration approach aligns with system configuration patterns
- Monitoring integration: Implementation supports established monitoring and observability patterns
- Operational readiness: Implementation approach supports operational requirements
</system_integration_readiness>

## Content Guidelines

### Architectural Validation Principles for LLM Self-Evaluation

- System Consistency: All architectural decisions must align with overall system design and patterns
- Component Clarity: Component boundaries and responsibilities must be clear and well-defined
- Integration Alignment: Integration approaches must follow established architectural patterns
- Standards Compliance: Implementation must meet security, performance, and compliance standards

### LLM Error Prevention Checklist

- Avoid: Architectural decisions that conflict with established system design patterns
- Avoid: Component designs that violate separation of concerns or create tight coupling
- Avoid: Integration approaches that bypass established architectural patterns
- Reference: Ensure architectural alignment with system design standards and [story.md](./.krci-ai/templates/story.md) template

==== END FILE ====

# Shared Templates

==== FILE: .krci-ai/templates/sad-template.md ====
# Solution Architecture Document (SAD): {{system_name}}

## 1. Executive Summary

**Purpose**: Provide business-focused overview connecting architecture to business value and strategic objectives.

{{executive_summary}}

### Key Sections

- **Business Context**: {{business_context}}
- **Architectural Approach**: {{architectural_approach}}
- **Key Benefits**: {{key_benefits}}
- **Success Metrics**: {{success_metrics}}

---

## 2. Introduction

**Purpose**: Establish foundation and context for all architectural decisions throughout the document.

### 2.1 Definitions, Acronyms, Abbreviations

{{definitions_acronyms_abbreviations}}

### 2.2 Scope

{{scope_boundaries}}

**What's Included:**

- {{included_systems}}

**What's Excluded:**

- {{excluded_systems}}

### 2.3 Stakeholders

{{stakeholders}}

| Stakeholder | Role | Key Concerns |
|-------------|------|--------------|
| {{stakeholder_name}} | {{stakeholder_role}} | {{stakeholder_concerns}} |

### 2.4 PRD Requirements Mapping

{{prd_requirements_mapping}}

**Requirements Traceability:**

| PRD Requirement | Architectural Component | Implementation Approach |
|-----------------|------------------------|------------------------|
| {{requirement_id}} ({{requirement_description}}) | {{component_name}} | {{implementation_approach}} |

**Purpose**: Establish direct traceability from PRD business/system requirements to architectural decisions.

---

## 3. Context

**Purpose**: Establish comprehensive environmental understanding across all architectural domains.

### 3.1 Technology Strategy

{{technology_strategy}}

**Organizational Technology Direction:**

- {{tech_strategy_point_1}}
- {{tech_strategy_point_2}}

**Alignment with Architecture:**

- {{alignment_point_1}}
- {{alignment_point_2}}

### 3.2 Business Architecture

{{business_architecture}}

**Key Business Processes:**

- {{business_process_1}}
- {{business_process_2}}

**Business Objectives:**

- {{business_objective_1}}
- {{business_objective_2}}

### 3.3 Data Architecture

{{data_architecture}}

**Data Models:**

- {{data_model_1}}
- {{data_model_2}}

**Data Flows:**

- {{data_flow_1}}
- {{data_flow_2}}

**Data Governance:**

- {{governance_requirement_1}}
- {{governance_requirement_2}}

### 3.4 Infrastructure Strategy

{{infrastructure_strategy}}

**Organizational Infrastructure Direction:**

- {{org_infrastructure_1}}
- {{org_infrastructure_2}}

**Strategic Platform Decisions:**

- {{strategic_platform_1}}
- {{strategic_platform_2}}

### 3.5 Application Architecture

{{application_architecture}}

**Application Landscape:**

- {{application_1}}
- {{application_2}}

**Integration Points:**

- {{integration_point_1}}
- {{integration_point_2}}

### 3.6 Security Architecture

{{security_architecture}}

**Security Posture:**

- {{security_requirement_1}}
- {{security_requirement_2}}

**Compliance Requirements:**

- {{compliance_requirement_1}}
- {{compliance_requirement_2}}

---

## 4. Requirements

**Purpose**: Document comprehensive driving forces behind all architectural decisions.

### 4.1 Business Goals

{{business_goals}}

**Primary Objectives:**

1. {{business_goal_1}}
2. {{business_goal_2}}
3. {{business_goal_3}}

**Success Criteria:**

- {{success_criterion_1}}
- {{success_criterion_2}}

### 4.2 Functional Requirements

{{functional_requirements}}

**Core Functionality:**

| Requirement ID | Description | Priority | Architectural Impact |
|----------------|-------------|----------|---------------------|
| {{req_id}} | {{req_description}} | {{priority}} | {{arch_impact}} |

### 4.3 Non-Functional Requirements

{{non_functional_requirements}}

**Quality Attributes:**

| Quality Attribute | Requirement | Measurement | Architectural Approach |
|-------------------|-------------|-------------|----------------------|
| Performance | {{performance_req}} | {{performance_metric}} | {{performance_approach}} |
| Scalability | {{scalability_req}} | {{scalability_metric}} | {{scalability_approach}} |
| Security | {{security_req}} | {{security_metric}} | {{security_approach}} |
| Availability | {{availability_req}} | {{availability_metric}} | {{availability_approach}} |

### 4.4 Constraints

{{constraints}}

**Technical Constraints:**

- {{technical_constraint_1}}
- {{technical_constraint_2}}

**Business Constraints:**

- {{business_constraint_1}}
- {{business_constraint_2}}

**Resource Constraints:**

- {{resource_constraint_1}}
- {{resource_constraint_2}}

### 4.5 Assumptions

{{assumptions}}

**Technical Assumptions:**

- {{technical_assumption_1}}
- {{technical_assumption_2}}

**Business Assumptions:**

- {{business_assumption_1}}
- {{business_assumption_2}}

---

## 5. Baseline Architecture

**Purpose**: Document current state architecture to understand transformation scope and complexity.

### 5.1 Conceptual View

{{baseline_conceptual_view}}

**Current System Overview:**

- {{current_system_1}}
- {{current_system_2}}

### 5.2 Logical View

{{baseline_logical_view}}

**Current System Components:**

| Component | Current Purpose | Technology Stack | Condition | Migration Notes |
|-----------|----------------|------------------|-----------|-----------------|
| {{current_component}} | {{current_purpose}} | {{current_tech}} | {{condition}} | {{migration_notes}} |

### 5.3 Integration View

{{baseline_integration_view}}

**Current Integration Points:**

- {{integration_1}}
- {{integration_2}}

### 5.4 Physical/Deployment View

{{baseline_deployment_view}}

**Current Infrastructure:**

- {{infrastructure_1}}
- {{infrastructure_2}}

---

## 6. Target Architecture

**Purpose**: Define comprehensive future state architecture with clear implementation guidance.

### 6.1 Conceptual View (C4 Context Level)

{{target_conceptual_view}}

**System Context Diagram:**

```mermaid
C4Context
    title System Context for {{system_name}}

    Person(user, "{{user_persona}}", "{{user_description}}")
    System(mainSystem, "{{system_name}}", "{{system_description}}")
    System_Ext(extSystem, "{{external_system}}", "{{external_description}}")

    Rel(user, mainSystem, "{{relationship_description}}")
    Rel(mainSystem, extSystem, "{{integration_description}}")
```

**External Systems:**

- {{external_system_1}}: {{relationship_1}}
- {{external_system_2}}: {{relationship_2}}

### 6.2 Logical View (C4 Container/Component)

{{target_logical_view}}

**Container Diagram (C4 Level 2):**

```mermaid
C4Container
    title Container Diagram for {{system_name}}

    Person(user, "{{user_persona}}")
    System_Boundary(c1, "{{system_name}}") {
        Container(web, "{{web_container}}", "{{web_tech}}", "{{web_description}}")
        Container(api, "{{api_container}}", "{{api_tech}}", "{{api_description}}")
        Container(db, "{{database_container}}", "{{db_tech}}", "{{db_description}}")
    }

    Rel(user, web, "{{user_web_rel}}")
    Rel(web, api, "{{web_api_rel}}")
    Rel(api, db, "{{api_db_rel}}")
```

**Component Diagram (C4 Level 3):**

```mermaid
C4Component
    title Component Diagram for {{component_name}}

    Container(container, "{{container_name}}", "{{container_tech}}", "{{container_description}}")

    Component(comp1, "{{component_1}}", "{{comp1_tech}}", "{{comp1_description}}")
    Component(comp2, "{{component_2}}", "{{comp2_tech}}", "{{comp2_description}}")

    Rel(comp1, comp2, "{{component_relationship}}")
```

**Key Components:**

| Component | Responsibility | Technology | Interfaces |
|-----------|---------------|------------|------------|
| {{component_name}} | {{responsibility}} | {{technology}} | {{interfaces}} |

### 6.3 Integration View

{{target_integration_view}}

**API Design:**

- {{api_1}}: {{api_purpose_1}}
- {{api_2}}: {{api_purpose_2}}

**Integration Patterns:**

- {{pattern_1}}: {{pattern_usage_1}}
- {{pattern_2}}: {{pattern_usage_2}}

### 6.4 Data View

{{target_data_view}}

**Data Architecture:**

- {{data_component_1}}: {{data_purpose_1}}
- {{data_component_2}}: {{data_purpose_2}}

**Data Flow:**

```mermaid
flowchart TD
    A[{{data_source}}] --> B[{{data_processor}}]
    B --> C[{{data_storage}}]
    C --> D[{{data_consumer}}]
```

### 6.5 Physical/Deployment View

{{target_deployment_view}}

**Deployment Architecture:**

- {{deployment_component_1}}: {{deployment_purpose_1}}
- {{deployment_component_2}}: {{deployment_purpose_2}}

**Infrastructure Mapping:**

```mermaid
graph TB
    subgraph "{{environment_name}}"
        subgraph "{{zone_1}}"
            A[{{component_1}}]
            B[{{component_2}}]
        end
        subgraph "{{zone_2}}"
            C[{{component_3}}]
            D[{{component_4}}]
        end
    end
```

### 6.6 Quality Attributes Implementation

{{quality_implementation}}

**Architecture Quality Approaches:**

| Quality Attribute | Implementation Strategy | Architectural Pattern | Validation Method |
|-------------------|------------------------|---------------------|-------------------|
| Performance | {{performance_strategy}} | {{performance_pattern}} | {{performance_validation}} |
| Scalability | {{scalability_strategy}} | {{scalability_pattern}} | {{scalability_validation}} |
| Security | {{security_strategy}} | {{security_pattern}} | {{security_validation}} |
| Availability | {{availability_strategy}} | {{availability_pattern}} | {{availability_validation}} |

### 6.7 Risks and Mitigations

{{risks_mitigations}}

| Risk | Impact | Probability | Mitigation Strategy |
|------|--------|-------------|-------------------|
| {{risk_1}} | {{impact_1}} | {{probability_1}} | {{mitigation_1}} |
| {{risk_2}} | {{impact_2}} | {{probability_2}} | {{mitigation_2}} |

### 6.8 Solution Strategy

{{solution_strategy}}

**Architectural Principles:**

1. {{principle_1}}: {{principle_description_1}}
2. {{principle_2}}: {{principle_description_2}}

**Technology Decisions:**

| Decision Area | Chosen Technology | Rationale | Trade-offs |
|---------------|-------------------|-----------|------------|
| {{decision_area_1}} | {{technology_1}} | {{rationale_1}} | {{tradeoffs_1}} |
| {{decision_area_2}} | {{technology_2}} | {{rationale_2}} | {{tradeoffs_2}} |

**Architecture Patterns:**

- {{pattern_1}}: {{pattern_rationale_1}}
- {{pattern_2}}: {{pattern_rationale_2}}

---

## 7. Transition/Migration

**Purpose**: Define clear implementation roadmap from current state to target architecture.

### 7.1 Migration Approach

{{migration_approach}}

**Migration Strategy:**

- {{strategy_element_1}}
- {{strategy_element_2}}

**Migration Principles:**

- {{principle_1}}
- {{principle_2}}

### 7.2 Migration Roadmap

{{migration_roadmap}}

**Implementation Phases:**

| Phase | Duration | Scope | Dependencies | Success Criteria |
|-------|----------|-------|--------------|------------------|
| {{phase_1}} | {{duration_1}} | {{scope_1}} | {{dependencies_1}} | {{criteria_1}} |
| {{phase_2}} | {{duration_2}} | {{scope_2}} | {{dependencies_2}} | {{criteria_2}} |

### 7.3 Implementation Guidance

{{implementation_guidance}}

**Epic Breakdown Guidance:**

| Architectural Component | Epic Mapping | Story Creation Focus |
|------------------------|--------------|---------------------|
| {{component_1}} | {{epic_guidance_1}} | {{story_focus_1}} |
| {{component_2}} | {{epic_guidance_2}} | {{story_focus_2}} |

**Development Standards:**

- {{dev_standard_1}}
- {{dev_standard_2}}

**API Design Guidelines:**

- {{api_guideline_1}}
- {{api_guideline_2}}

**Testing Alignment:**

- {{testing_alignment_1}}
- {{testing_alignment_2}}

---

## 8. Architectural Decisions

**Purpose**: Record significant decisions with comprehensive rationale and alternatives analysis.

### 8.1 Decision Log (ADR Format)

{{architectural_decisions}}

**Decision Template:**

#### ADR-001: {{decision_title}}

**Status**: {{status}} (Proposed/Accepted/Superseded)
**Date**: {{decision_date}}
**Deciders**: {{decision_makers}}

**Context:**
{{decision_context}}

**Decision:**
{{decision_made}}

**Consequences:**
**Positive:**

- {{positive_consequence_1}}
- {{positive_consequence_2}}

**Negative:**

- {{negative_consequence_1}}
- {{negative_consequence_2}}

**Alternatives Considered:**

- {{alternative_1}}: {{alternative_rationale_1}}
- {{alternative_2}}: {{alternative_rationale_2}}

---

## 9. Cross-Cutting Concerns

**Purpose**: Address system-wide aspects that affect multiple architectural components.

### 9.1 Security

{{security_concerns}}

**Cross-Cutting Security Implementation:**

- {{cross_cutting_security_1}}
- {{cross_cutting_security_2}}

**Security Integration Points:**

- {{security_integration_1}}
- {{security_integration_2}}

### 9.2 Scalability

{{scalability_concerns}}

**Scalability Strategy:**

- {{scalability_strategy_1}}
- {{scalability_strategy_2}}

**Scaling Patterns:**

- {{scaling_pattern_1}}
- {{scaling_pattern_2}}

### 9.3 Observability

{{observability_concerns}}

**Monitoring Strategy:**

- {{monitoring_approach_1}}
- {{monitoring_approach_2}}

**Logging Strategy:**

- {{logging_approach_1}}
- {{logging_approach_2}}

### 9.4 Fault Tolerance

{{fault_tolerance_concerns}}

**Resilience Patterns:**

- {{resilience_pattern_1}}
- {{resilience_pattern_2}}

**Error Handling Strategy:**

- {{error_handling_1}}
- {{error_handling_2}}

---

## 10. Quality Assurance

**Purpose**: Define comprehensive approach to validating architectural decisions and implementation quality.

### 10.1 Testing Strategy

{{testing_strategy}}

**Architecture Testing Approach:**

- {{arch_testing_1}}
- {{arch_testing_2}}

**Quality Gates:**

- {{quality_gate_1}}
- {{quality_gate_2}}

### 10.2 Test Automation Approach

{{test_automation}}

**Automation Strategy:**

- {{automation_approach_1}}
- {{automation_approach_2}}

**Testing Tools:**

- {{testing_tool_1}}: {{tool_purpose_1}}
- {{testing_tool_2}}: {{tool_purpose_2}}

### 10.3 Quality Metrics

{{quality_metrics}}

**Success Metrics:**

| Quality Aspect | Metric | Target | Measurement Method |
|----------------|---------|--------|-------------------|
| {{quality_1}} | {{metric_1}} | {{target_1}} | {{method_1}} |
| {{quality_2}} | {{metric_2}} | {{target_2}} | {{method_2}} |

---

## 11. Appendices

**Purpose**: Provide supplementary materials and detailed references.

### 11.1 Glossary

{{glossary}}

| Term | Definition |
|------|------------|
| {{term_1}} | {{definition_1}} |
| {{term_2}} | {{definition_2}} |

### 11.2 Diagrams

{{diagrams}}

**Diagram Index:**

- {{diagram_1}}: {{diagram_purpose_1}}
- {{diagram_2}}: {{diagram_purpose_2}}

### 11.3 Reference Materials

{{reference_materials}}

**Standards and Guidelines:**

- {{reference_1}}
- {{reference_2}}

**Related Documentation:**

- {{related_doc_1}}
- {{related_doc_2}}

---

**Document Status**: {{document_status}}
**Version**: {{version}}
**Last Updated**: {{last_updated}}
**Next Review**: {{next_review}}

==== END FILE ====

==== FILE: .krci-ai/templates/story.md ====
# Story {{story_number}}: {{story_title}}

<instructions>
STORY STRUCTURE: Follow this exact section ordering:
1. Status → 2. Dependencies → 3. Story → 4. Acceptance Criteria → 5. Description → 6. Technical Context → 7. Tasks/Subtasks → 8. Implementation Results → 9. QA Checklist
</instructions>

## Status

<status>
| Field                  | Value                       |
|------------------------|-----------------------------|
| Status                 | {{status}}                  |
| Epic                   | {{epic_reference}}          |
| Priority               | {{priority}}                |
| Estimated Story Points | {{story_points}}            |
| Jira                   | {{jira_ticket}}             |
</status>

<instructions>
Status tracking and Epic/PRD traceability. Enables progress monitoring and dependency validation.

Status Options: Draft -> Approved -> In Progress -> Done -> Completed
Epic Reference Example: Epic 1 - Unified Agent Activation
Priority Example: Critical, High, Medium, Low
Story Points Example: 1, 2, 3, 5, 8, 13 (Fibonacci scale)
Jira Example: [EPMDEDP-15497](https://jira.example.com/browse/EPMDEDP-15497) or "None" if not assigned

CRITICAL: Status section contains ONLY these 5 fields. Do NOT add Dependencies or other fields here.
</instructions>

## Dependencies

<dependencies>
**Blocking:**
{{blocking_dependencies}}

**Blocked By:**
{{blocked_by_dependencies}}

**System/Test Dependencies:**
{{system_test_dependencies}}

<instructions>
Define precise dependencies for execution order and validation readiness.

CRITICAL: Dependencies section comes immediately after Status section.

Format:

- Blocking: Items that depend on this story
- Blocked By: Items this story depends on (stories, approvals)
- System/Test: Environments, versions, stubs/mocks, fixtures

Examples:

- Blocking: Story 02.03 (depends on config file produced here)
- Blocked By: Story 02.01 (API contract), Security approval
- System/Test: Local IDP stub v1.2; Env var FEATURE_FLAG=on

DO:

- Use exact story numbers/links and system versions
- Include test doubles (stubs/mocks) to avoid external blockers
- State "None" explicitly if empty

DON'T:

- List vague dependencies (e.g., "backend work")
- Omit versions/links where relevant
- Depend on production-only services for acceptance
- Put dependencies in Status section
</instructions>

</dependencies>

## Story

<user_story>
**As a** {{persona}},
**I want** {{goal}},
**so that** {{business_value}}.

<instructions>
Standard story format focusing on persona, goal, and value. Must align with Epic Target Users and provide specific value.

Story Example:
As a Software Engineer,
I want a unified agent activation command in the IDE,
so that I can start using an agent without extra setup.

DO:

- Use persona from Epic (PRD-defined), not generic labels
- Make the goal specific and implementation-agnostic
- State tangible business/user value

DON'T:

- Invent personas not defined in PRD/Epic
- Use vague goals like "improve experience"
- Include solution details here
</instructions>

</user_story>

## Acceptance Criteria

<acceptance_criteria>
{{acceptance_criteria}}

<instructions>
CRITICAL: Acceptance Criteria section comes immediately after Story section. Create specific, testable conditions that define completion at story level. Story ACs SHOULD include verification commands and expected outputs.

Required Structure for each criterion:

1. Scenario (Given/When/Then): Brief user/system flow to validate
2. Expected Behavior: Clear, observable outcome
3. Verification Command: Non-interactive command(s) with expected exit code/output
4. Files Created/Changed: Exact paths
5. Guardrails: NFR constraints applicable to this story (e.g., perf, security)
6. Test Data/Fixtures: Test users, payloads, seeds
7. Environment/Flags: Env vars, feature flags/toggles, mock/stub switches
8. Evidence: Required artifacts (logs, screenshots, run output)
9. Non-automatable at story-level: true/false (and rationale if true)
10. Traceability: Epic AC id(s), PRD BR/NFR
11. Out of Scope (optional): Clarify non-goals

Acceptance Criteria Example:

1. **OAuth Integration**: System successfully connects to OAuth service and processes authentication requests
2. **Token Processing**: Validate, decode, and process JWT tokens with proper error handling
3. **Session Management**: Implement session timeout, auto-renewal, and security mechanisms
4. **API Endpoints**: Develop comprehensive authentication endpoints for complete workflow
5. **Testing & Validation**: Implement comprehensive test suite with security validation and coverage requirements

Format Guidelines:

- Simple numbered list with bold titles and clear descriptions
- One line per acceptance criterion
- Focus on specific, testable outcomes
- Avoid complex sub-bullets or excessive detail
- Keep descriptions concise but complete

DO:

- Use simple, clear numbered format
- Focus on specific, testable outcomes
- Keep descriptions concise and actionable
- Map each AC to Epic requirements

DON'T:

- Add complex sub-bullet structures
- Use excessive technical detail in AC descriptions
- Create overly long or complex acceptance criteria
- Mix implementation details with acceptance criteria
</instructions>

</acceptance_criteria>

## Description

<description>
{{description}}

<instructions>
Comprehensive context for strategic understanding and implementation guidance. Provide detailed background, technical context, and architectural alignment.

Comprehensive Context Requirements:

- Strategic purpose within the Epic and broader system goals
- Detailed technical background and architectural significance
- Implementation philosophy and approach rationale
- System integration considerations and dependencies
- Quality and architectural compliance requirements
- Technical constraints and design decision context
</instructions>

</description>

## Technical Context

<technical_context>
{{technical_context}}

<instructions>
Detailed technical background and architectural considerations. Provide comprehensive implementation context and design guidance.

Technical Context Content:

- Architectural significance and system integration points
- Technical approach and design pattern rationale
- Implementation constraints and technical dependencies
- Quality considerations and testing strategy overview
- System design principles and architectural alignment
- Performance, security, and scalability considerations
</instructions>

</technical_context>

## Tasks/Subtasks

<tasks_subtasks>
{{tasks_subtasks}}

<instructions>
Create LLM-executable implementation plan with atomic tasks and validation. Each task maps to acceptance criteria with specific commands and file paths.

TASKS/SUBTASKS STRUCTURE REQUIREMENTS:

- [ ] **Task N: Task Description (AC: X, Y)**
  - [ ] Clear, actionable implementation step with specific deliverable
  - [ ] Another implementation step focused on technical execution
  - [ ] Testing and validation step with clear completion criteria
  - [ ] Command: `example command` - Expected: expected result
  - [ ] Integration and deployment verification step

STRUCTURE REQUIREMENTS:

- Tasks start with checkbox and bold title: `- [ ] **Task N: Description (AC: X, Y)**`
- Subtasks indented with 2 spaces: `  - [ ] Subtask description`
- Clean bullet checkboxes for implementation tracking at both task and subtask level
- Focused on actionable implementation steps
- Clear, concise descriptions without excessive formatting
- Each checkbox represents a specific, completable work item
- Map tasks to acceptance criteria clearly
- Include testing and validation steps
- Avoid excessive formatting or complex structure
- Do not mix strategic planning with implementation checkboxes
- Essential implementation checkboxes are mandatory
</instructions>

</tasks_subtasks>

## Implementation Results

<implementation_results>
{{implementation_results}}

<instructions>
Concrete outcomes populated AFTER completion (evidence-first).

Include:

- Evidence per AC (links to logs, screenshots, artifacts)
- Files created/updated with paths
- Commands executed with outputs
- Final validation summary

DO:

- Use past tense (Created, Implemented, Validated)
- Link evidence to specific AC ids
- Include actual command results/exit codes

DON'T:

- Populate before implementation
- Use future/planning language
- Omit evidence for any AC
</instructions>

</implementation_results>

## QA Checklist

<qa_checklist>
{{qa_checklist}}

<instructions>
Create specific verification steps with commands and expected outputs. Enable automated testing and quality validation.

QA Checklist Example:

### Functional

- [ ] Schema Validation: `python hack/validate-agents.py` (Expect: exit 0)
- [ ] File Existence: `ls -la assets/agents/architect.yaml` (Expect: file exists)
- [ ] Content Validation: `grep -q "identity" assets/agents/architect.yaml` (Expect: pattern found)

### Integration

- [ ] IDE Testing: Activate agent in Cursor (Expect: response within 5s)
- [ ] Cross-Platform: Validate on macOS/Linux/Windows (Expect: consistent behavior)

### Security & Privacy

- [ ] Secrets: No tokens/secrets in logs
- [ ] Auth: No P0/P1 security findings in changed scope

### Accessibility (if UI)

- [ ] Keyboard navigation and focus order
- [ ] Contrast ratio meets baseline

DO:

- Group tests by category; include expected outputs
- Keep checks non-interactive where possible
- Align with AC guardrails

DON'T:

- Use subjective testing criteria
- Omit expected outputs or success indicators
- Depend on production-only services
</instructions>

</qa_checklist>

==== END FILE ====

==== FILE: .krci-ai/templates/architecture-review.md ====
# Architecture Review: {{system_name}}

## Review Summary

{{review_summary}}

## Architecture Assessment

### Overall Rating

{{overall_rating}}

### Architecture Strengths

{{architecture_strengths}}

### Areas for Improvement

{{improvement_areas}}

## Detailed Analysis

### Component Design Review

{{component_design_review}}

### Data Architecture Review

{{data_architecture_review}}

### Security Architecture Review

{{security_architecture_review}}

### Scalability Assessment

{{scalability_assessment}}

### Performance Analysis

{{performance_analysis}}

### Maintainability Review

{{maintainability_review}}

## Anti-Patterns Identified

{{antipatterns_found}}

## Technical Debt Assessment

{{technical_debt}}

## Compliance Review

{{compliance_assessment}}

## Risk Analysis

### High Priority Risks

{{high_priority_risks}}

### Medium Priority Risks

{{medium_priority_risks}}

### Low Priority Risks

{{low_priority_risks}}

## Recommendations

### Immediate Actions (Priority 1)

{{immediate_actions}}

### Short-term Improvements (Priority 2)

{{short_term_improvements}}

### Long-term Enhancements (Priority 3)

{{long_term_enhancements}}

## Implementation Roadmap

{{implementation_roadmap}}

## Review Conclusions

{{review_conclusions}}

==== END FILE ====

# Reference Data

==== FILE: .krci-ai/data/architecture-principles.md ====
# Architecture Design Principles

## Core Principles

<architecture_principles>

### Scalability First

Design systems for growth from day one. Consider horizontal and vertical scaling strategies.

- Plan for 10x current load
- Design stateless services where possible
- Use load balancing and auto-scaling

### Security by Design

Implement security at every layer, not as an afterthought.

- Apply defense in depth
- Use principle of least privilege
- Implement secure defaults
- Regular security assessments

### Fault Tolerance

Plan for component failures and graceful degradation.

- Design for failure scenarios
- Implement circuit breakers
- Use redundancy and failover mechanisms
- Plan disaster recovery

### Separation of Concerns

Organize code and components by responsibility.

- Single responsibility principle
- Clear component boundaries
- Minimal coupling between components
- High cohesion within components

### Observability

Build monitoring, logging, and tracing into the system.

- Comprehensive logging strategy
- Real-time monitoring and alerting
- Distributed tracing for complex flows
- Performance metrics and analytics

### Maintainability

Prioritize code and architecture that can be easily understood and modified.

- Clear documentation
- Consistent coding standards
- Automated testing
- Simple deployment processes

### Performance Optimization

Balance performance with maintainability and cost.

- Profile before optimizing
- Focus on bottlenecks
- Consider caching strategies
- Optimize data access patterns

### Evolutionary Architecture

Design for change and adaptation over time.

- Loosely coupled architecture
- API versioning strategy
- Incremental migration paths
- Technology stack flexibility
</architecture_principles>

==== END FILE ====

==== FILE: .krci-ai/data/design-patterns.md ====
# Design Patterns for Architecture

## Architectural Patterns

<architectural_patterns>

### Microservices Pattern

Decompose applications into small, independent services.

- When to use: Large, complex applications with multiple teams
- Benefits: Independent deployment, technology diversity, fault isolation
- Considerations: Network latency, data consistency, operational complexity

### Event-Driven Architecture

Use events to trigger and communicate between decoupled services.

- When to use: Real-time processing, loose coupling requirements
- Benefits: Scalability, flexibility, real-time responsiveness
- Considerations: Event ordering, eventual consistency

### API Gateway Pattern

Centralize cross-cutting concerns for microservices.

- When to use: Multiple client types, microservices architecture
- Benefits: Unified entry point, security, rate limiting
- Considerations: Single point of failure, performance bottleneck

### CQRS (Command Query Responsibility Segregation)

Separate read and write operations for better performance.

- When to use: High-read vs high-write workloads
- Benefits: Optimized read/write models, scalability
- Considerations: Complexity, eventual consistency
</architectural_patterns>

## Integration Patterns

<integration_patterns>

### Circuit Breaker

Prevent cascading failures in distributed systems.

- Implementation: Monitor failure rates, open circuit on threshold
- Benefits: System stability, graceful degradation
- Tools: Hystrix, Resilience4j

### Bulkhead

Isolate resources to prevent total system failure.

- Implementation: Separate thread pools, connection pools
- Benefits: Fault isolation, resource protection

### Retry with Exponential Backoff

Handle transient failures gracefully.

- Implementation: Exponential delays between retries
- Benefits: Improved reliability, reduced system load
</integration_patterns>

## Data Patterns

<data_patterns>

### Database per Service

Each microservice owns its data exclusively.

- Benefits: Service independence, technology flexibility
- Considerations: Data consistency, query complexity

### Saga Pattern

Manage distributed transactions across services.

- Types: Choreography-based, Orchestration-based
- Benefits: Data consistency without distributed transactions

### Event Sourcing

Store events instead of current state.

- Benefits: Complete audit trail, temporal queries
- Considerations: Storage overhead, complexity
</data_patterns>

## Security Patterns

<security_patterns>

### Zero Trust Architecture

Never trust, always verify approach to security.

- Principles: Verify identity, least privilege, assume breach
- Implementation: Multi-factor authentication, micro-segmentation

### OAuth 2.0 / OpenID Connect

Standard patterns for authentication and authorization.

- Use cases: API access, single sign-on, third-party integration
</security_patterns>

## Performance Patterns

<performance_patterns>

### Caching Strategies

- Cache-Aside: Application manages cache
- Write-Through: Write to cache and database simultaneously
- Write-Behind: Write to cache first, database later

### Load Balancing

- Round Robin: Distribute requests evenly
- Least Connections: Route to server with fewest active connections
- Weighted: Route based on server capacity
</performance_patterns>

==== END FILE ====

==== FILE: .krci-ai/data/krci-ai/core-sdlc-framework.md ====
# SDLC Framework Quick Reference

Purpose: AI agents collaborate through structured, filesystem-based artifacts. Agents use this document to understand role, dependencies, and locate templates/standards.

## Framework Principles

- Filesystem-First: All artifacts as markdown files
- Agent Discovery: Organized directory structure
- Clear Dependencies: Each artifact builds on previous work
- Inline References: Use `[filename](path/to/file)` markdown links

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
| Framework Advisor | `advisor-v1` | Framework Components, Validation Reports | None (meta-framework support) | Framework component creation, validation, maintenance guidance |
</roles>

## Artifact Flow

```text
Project Brief → PRD → Epics → Stories → Code → Tests → MVP → Marketing
                  ↓             ↓
              Architecture ← → Code
```

**Flow**: PM(Brief+PRD) → BA(refine) → PO(Epics+Stories) → Arch(design) → Dev(code) → QA(test) → PMM(marketing)

**Dependencies**:

- Project Brief: No dependencies (root)
- PRD: Requires Project Brief
- Epic: Requires PRD completion
- Story: Requires Epic + Architecture
- Code: Requires Stories + Architecture
- Tests: Requires Stories + Code
- Marketing: Requires PRD + MVP

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
| **Framework Components** | Advisor | Framework maintenance | Agents, tasks, templates, validation reports | `/.krci-ai/{agents,tasks,templates}/*.md` |

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
│   └── data/                       # REFERENCE: Standards & guidelines (can have subfolders)
```

<quality_gates>

1. Project Brief Approval → Enables PRD creation
2. PRD Approval → Enables Epic/Architecture creation
3. Architecture Review → Enables implementation
4. Code Review → Enables testing
5. Test Validation → Enables MVP delivery

Enforcement: All `{{variables}}` filled, dependencies satisfied, template format followed.
</quality_gates>

<handoff_points>

1. Idea → PM: Market research
2. PM → BA: PRD refinement
3. BA → PO: Requirements to backlog
4. PO → Arch: Technical feasibility
5. Arch → Dev: Implementation guidance
6. Dev → QA: Quality validation
7. QA → MVP: Deployment readiness
8. PM → PMM: Go-to-market strategy
9. MVP → PMM: Marketing materials

Validation: Upstream artifacts complete, quality gates passed.
</handoff_points>

<common_issues>

| Issue | Resolution | Agent Action |
|-------|------------|--------------|
| Epic creation blocked | PRD incomplete | Return to Product Manager for PRD completion |
| Story blocked by Architecture | Architecture not defined | Complete Architecture Document before Story creation |
| Implementation blocked | Story criteria unclear | Return to Product Owner for Story refinement |
| Testing blocked | Code incomplete | Ensure Developer completes all Story requirements |
| Template variables unfilled | Validation failed | Complete all `{{variables}}` using referenced templates |
| Framework component errors | Validation failed | Consult Framework Advisor for component creation/fixes |

</common_issues>

<agent_workflow>

1. Find role in `<roles>` table (ID, outputs, dependencies)
2. Check dependencies exist and approved
3. Load templates: `.krci-ai/templates/{artifact-template}.md`
4. Execute task: `.krci-ai/tasks/{task-name}.md`
5. Validate: No `{{variables}}` unfilled, dependencies referenced
6. Quality gate: Standards met before handoff

Path examples: `.krci-ai/templates/story.md`
</agent_workflow>

## Template Variables

Templates use `{{variable_name}}` format. Required fields must be populated.
Example: `{{epic_title}}` → "User Authentication System"

### Template Conventions

CRITICAL: XML tags in templates are agent guidance only - exclude from final output.

<success_flow>
Idea → PM (Brief+PRD) → BA (Analysis) → PO (Epics+Stories) → Architect (Design) → Developer/Go Developer (Code) → QA (Tests) → MVP → PMM (Marketing)
</success_flow>

<implementation>
Agent: YAML (identity, commands, tasks)
Task: Markdown workflows with templates/data references
Template: Markdown with `{{variables}}`
Data: Standards and specifications

```yaml
agent:
  identity:
    name: "Role Name"
    id: "role-id-v1"
  commands:
    help: "Show commands"
    chat: "Role consultation"
    exit: "Exit persona"
  tasks:
    - "./.krci-ai/tasks/{task-name}.md"
```

</implementation>

==== END FILE ====

