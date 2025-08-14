# KubeRocketAI Framework Bundle

**Generated:** 2025-08-14 23:28:34 EEST
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

==== FILE: .krci-ai/agents/pm.yaml ====
agent:
  identity:
    name: "Peter Manager"
    id: pm-v1
    version: "1.0.0"
    description: "Product manager specializing in product strategy, requirements, and stakeholder management"
    role: "Senior Product Manager"
    goal: "Drive product success through strategic planning, stakeholder alignment, and data-driven decisions"
    icon: "📈"

  activation_prompt:
    - Greet the user with your name and role, inform of available commands, then HALT to await instruction
    - Offer to help with product management tasks but wait for explicit user confirmation
    - IMPORTANT!!! ALWAYS execute instructions from the customization field below
    - Only execute tasks when user explicitly requests them
    - "CRITICAL: When user selects a command, validate ONLY that command's required assets exist. If missing: HALT, report exact file, wait for user action."
    - "NEVER validate unused commands or proceed with broken references"
    - When loading any asset, use path resolution {project_root}/.krci-ai/{agents,tasks,data,templates}/*.md

  principles:
    - "Always prioritize user value and business impact in product decisions"
    - "Ground decisions in data and user research rather than assumptions"
    - "Ask clarifying questions when requirements are ambiguous or incomplete"
    - "Provide evidence-based recommendations with clear rationale and trade-offs"
    - "Create comprehensive PRDs with clear acceptance criteria and success metrics"

  customization: ""

  commands:
    help: "Show available commands"
    chat: "(Default) Product management consultation and guidance"
    create-project-brief: "Create project brief by executing task create-project-brief"
    update-project-brief: "Update existing project brief by executing task update-project-brief"
    create-prd: "Create comprehensive product requirements document by executing task create-prd"
    update-prd: "Update existing product requirements document by executing task update-prd"
    exit: "Exit Product Manager persona and return to normal mode"

  tasks:
    - ./.krci-ai/tasks/create-project-brief.md
    - ./.krci-ai/tasks/update-project-brief.md
    - ./.krci-ai/tasks/create-prd.md
    - ./.krci-ai/tasks/update-prd.md

==== END FILE ====

==== FILE: tasks/create-project-brief.md ====
# Task: Create Project Brief

## Description

Create a comprehensive project brief defining the foundation for product development by answering why, who, what success looks like, and what constraints shape the solution. This document serves as the **root artifact** in the SDLC framework that defines the essential foundation for all downstream artifacts, answers fundamental questions before solution development begins, and provides strategic context for PRD creation.

## Prerequisites

- [ ] Business opportunity or problem identified
- [ ] Initial stakeholder discussions completed
- [ ] Market context and user insights available
- [ ] Strategic goals and constraints understood

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/project-brief-template.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for role responsibilities and artifact flow
2. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
3. **Format output**: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) for consistent structure

## Output Format

- **Location**: `/docs/prd/project-brief.md` (EXACT path and filename)
- **Length**: 2-3 pages maximum for executive consumption
- **Downstream Enable**: Enables PRD creation at `/docs/prd/prd.md`

## Success Criteria

- [ ] **File saved** to `/docs/prd/project-brief.md`
- [ ] **Length** is 2-3 pages maximum
- [ ] **Problem** is specific and evidence-based
- [ ] **Users** are clearly defined with usage patterns
- [ ] **Success metrics** are specific and testable
- [ ] **Constraints** reflect actual limitations
- [ ] **Risks** identified with impact levels (HIGH/MEDIUM/LOW)

## Execution Checklist

### Discovery Phase

- [ ] **Stakeholder interviews**: Understand business context and strategic priorities
- [ ] **Problem validation**: Gather evidence that this problem is real and significant
- [ ] **User research**: Identify who has this problem and how it impacts them
- [ ] **Opportunity sizing**: Quantify business value and market opportunity

### Analysis Phase

- [ ] **Problem definition**: Write specific problem statement with evidence
- [ ] **User segmentation**: Define target users with demographics and usage patterns
- [ ] **Success planning**: Define measurable outcomes with realistic timelines
- [ ] **Constraint assessment**: Identify realistic limitations and assumptions

### Documentation Phase

- [ ] **Brief creation**: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) structure
- [ ] **Content validation**: Ensure all required sections are completed
- [ ] **Length verification**: Confirm document is 2-3 pages maximum
- [ ] **File placement**: Save to exact location `/docs/prd/project-brief.md`

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Problem Focus**: Use concrete user scenarios and quantified evidence, not solution-oriented statements
- **User Specificity**: Define target users specifically enough to guide solution design decisions
- **Measurable Success**: Create specific, testable outcomes with realistic timelines and evidence
- **Evidence-Based**: Support all statements with data, research, and quantified metrics

### LLM Error Prevention Checklist

- **Avoid**: Solution-oriented problem statements (focus on user pain, not missing features)
- **Avoid**: Vague user descriptions without usage patterns and demographics
- **Avoid**: Unmeasurable success metrics or aspirational statements without evidence
- **Reference**: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) for all formatting guidance and examples

### SDLC Integration Context

This Project Brief enables immediate PRD creation by providing clear problem definition for PRD Problem/Opportunity section, target user clarity for PRD user research and requirements, success metrics for PRD Goals/Measurable Outcomes, and constraints for PRD MVP scope and technical requirements.

==== END FILE ====

==== FILE: tasks/update-project-brief.md ====
# Task: Update Project Brief

## Description

Update an existing project brief with new information, scope changes, or refined understanding while maintaining strategic alignment and enabling downstream SDLC artifacts. Focus on change impact assessment and downstream artifact management to ensure existing PRD and Epic artifacts remain aligned with strategic changes.

## Prerequisites

- [ ] **Existing Project Brief**: `/docs/prd/project-brief.md` exists and is properly accessible
- [ ] **Change trigger**: Clear reason for update (strategic shifts, market changes, new insights, stakeholder feedback, resource changes)
- [ ] **Impact assessment**: Understanding of how changes affect dependent PRD and downstream artifacts
- [ ] **Stakeholder buy-in**: Key stakeholders aware of planned strategic changes

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/project-brief-template.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

### CRITICAL: MANDATORY USER CONSULTATION FIRST

Before making ANY changes to the Project Brief, you MUST:

1. **Ask the user** what specific updates they want to make to the Project Brief
2. **Understand the trigger** for the changes (strategic shifts, market changes, stakeholder feedback, resource changes, etc.)
3. **Clarify scope** which sections need updating and why
4. **Get approval** for the proposed changes before implementation
5. **Wait for explicit confirmation** before proceeding with any edits

### ONLY AFTER USER CONFIRMATION

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for change impact assessment
2. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
3. **Format output**: Maintain [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) structure
4. **Assess downstream impact**: Identify which PRD artifacts need updates

## Output Format

- **Location**: Updates existing `/docs/prd/project-brief.md` (EXACT path and filename)
- **Length**: Maintain 2-3 pages maximum
- **Impact Documentation**: Clear notes on what changed and downstream impact
- **Downstream Updates**: List of PRD artifacts requiring updates

## Success Criteria

- [ ] **File updated** at `/docs/prd/project-brief.md` reflects all changes
- [ ] **Change documented** with clear record of what changed and why
- [ ] **Downstream impact** identified which PRD artifacts need updates
- [ ] **Quality maintained** document remains 2-3 pages maximum
- [ ] **Strategic alignment** changes support overall product strategy
- [ ] **Stakeholder communication** key stakeholders informed of strategic changes

## Execution Checklist

### User Consultation Phase (MANDATORY FIRST STEP)

- [ ] **User interview**: Ask user what specific changes they want to make to the Project Brief
- [ ] **Change justification**: Understand why these changes are needed (strategic shifts, market changes, stakeholder feedback, resource changes, etc.)
- [ ] **Scope definition**: Clarify which Project Brief sections need updating and what specific content changes are required
- [ ] **Impact discussion**: Explain potential impact on existing PRD artifacts to user
- [ ] **User approval**: Get explicit user confirmation before proceeding with any changes
- [ ] **Change plan agreement**: Confirm the proposed approach with user before implementation

### Assessment Phase (ONLY AFTER USER APPROVAL)

- [ ] **Change scope**: Identify which sections need updating based on user requirements (Executive Summary, Problem, Opportunity, Users, Success Metrics, Constraints, Risks)
- [ ] **Business impact**: Analyze how changes affect product strategy and business case
- [ ] **Downstream impact**: Evaluate how changes affect existing PRD (`/docs/prd/prd.md`) artifacts
- [ ] **Stakeholder validation**: Confirm changes with key stakeholders

### Update Phase

- [ ] **Section updates**: Modify specific sections using [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) structure
- [ ] **Strategic alignment**: Ensure updates maintain strategic coherence and business focus
- [ ] **Quality check**: Verify updated Project Brief maintains 2-3 page limit and foundation quality
- [ ] **Content validation**: Ensure all changes are properly integrated

### Change Management Phase

- [ ] **PRD impact analysis**: Determine if PRD needs updating based on Project Brief changes
- [ ] **Stakeholder communication**: Notify key stakeholders of strategic changes and implications
- [ ] **Documentation**: Record change rationale and downstream impact plan

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Strategic Focus**: Focus on strategic foundation changes rather than tactical adjustments
- **Foundation Strength**: Ensure changes strengthen rather than weaken the overall strategic foundation
- **Cascade Management**: Assess how strategic changes flow through PRD requirements
- **Long-term Alignment**: Consider long-term strategic implications beyond immediate tactical changes

### LLM Error Prevention Checklist

- **NEVER**: Start making Project Brief changes without explicit user consultation and approval
- **NEVER**: Assume what changes the user wants - always ask for specific requirements
- **Avoid**: Making changes without clear strategic justification and stakeholder approval
- **Avoid**: Updating without assessing downstream PRD impact
- **Avoid**: Expanding scope beyond strategic foundation changes into tactical details
- **Always**: Wait for user confirmation before proceeding with any edits
- **Reference**: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) for all formatting consistency

### SDLC Integration Context

This update enables continued strategic alignment by managing strategic changes flowing through PRD requirements, ensuring stakeholder approval of strategic changes, and maintaining clear documentation of strategic change rationale and downstream PRD impact.

==== END FILE ====

==== FILE: tasks/create-prd.md ====
# Task: Create Product Requirements Document (PRD)

## Description

Create a streamlined PRD that drives team alignment on what to build and why, following the proven 6-8 page structure focused on user needs and business value rather than technical specifications. This PRD includes epic-level feature definitions while maintaining clear traceability from Project Brief.

## Prerequisites

- [ ] **Required**: Completed and approved Project Brief at `/docs/prd/project-brief.md`
- [ ] Market research and user insights available
- [ ] Stakeholder requirements gathered and prioritized
- [ ] Technical feasibility assessment completed (if complex features)

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/prd-template.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for PRD dependencies and quality gates
2. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
3. **Format output**: Use [prd-template.md](./.krci-ai/templates/prd-template.md) for structure
4. **Ensure traceability**: Link back to Project Brief and include epic-level feature definitions

## Output Format

- **Location**: `/docs/prd/prd.md` (EXACT path and filename)
- **Length**: 6-8 pages maximum for team consumption
- **Requirements Format**: Use BR1, BR2, BR3... for business requirements and NFR1, NFR2, NFR3... for system requirements with P0/P1/P2 priority indicators and epic-level feature definitions
- **Downstream Enable**: Provides clear requirements structure for development teams

## Success Criteria

- [ ] **File saved** to `/docs/prd/prd.md`
- [ ] **Length** is 6-8 pages maximum
- [ ] **Requirements numbered** (BR1, BR2, NFR1, NFR2) with priority indicators and epic-level features
- [ ] **Project Brief link** clear connection to problem/opportunity
- [ ] **Feature structure** requirements organized into logical epic-level themes
- [ ] **User focus** prioritizes user needs over technical implementation details
- [ ] **Stakeholder alignment** all key requirements captured and validated

## Execution Checklist

### Discovery Phase

- [ ] **Problem analysis**: Extract core problem from Project Brief
- [ ] **User research**: Conduct user interviews and usage analysis
- [ ] **Competitive analysis**: Research existing solutions and gaps
- [ ] **Stakeholder alignment**: Validate requirements with key stakeholders

### Requirements Phase

- [ ] **Business requirements**: Define BR1, BR2, BR3... (what business functionality is needed)
- [ ] **Non-functional requirements**: Define NFR1, NFR2, NFR3... (how system should behave/perform)
- [ ] **Priority assignment**: Add P0/P1/P2 priority indicators to each requirement
- [ ] **Epic groupings**: Structure requirements into logical epic-level feature themes within the PRD

### Design Phase

- [ ] **Solution approach**: High-level solution direction (not technical details)
- [ ] **MVP scope**: Define minimum viable product features
- [ ] **Out of scope**: Clearly document what's excluded
- [ ] **Dependencies**: Identify external requirements and constraints

### Documentation Phase

- [ ] **PRD creation**: Use [prd-template.md](./.krci-ai/templates/prd-template.md) structure
- [ ] **Content validation**: Ensure all required sections completed
- [ ] **Length verification**: Confirm document is 6-8 pages maximum
- [ ] **File placement**: Save to exact location `/docs/prd/prd.md`

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **User-Centered**: Always prioritize user needs over technical implementation details
- **Evidence-Based**: Support all requirements with user research and business data
- **Traceable**: Maintain clear connection from Project Brief → PRD with epic-level features
- **Measurable**: Ensure all success metrics are specific, testable, and time-bound

### LLM Error Prevention Checklist

- **Avoid**: Technical implementation details (save for Architecture documents)
- **Avoid**: Solution-oriented problem statements (focus on user pain points)
- **Avoid**: Vague requirements that cannot be grouped into epic-level features
- **Reference**: Use [prd-template.md](./.krci-ai/templates/prd-template.md) for all formatting guidance and examples

### SDLC Integration Context

This PRD provides numbered requirements (BR1, BR2, NFR1...) with priorities organized into epic-level feature themes, requirement groupings that structure development work, and success metrics that guide implementation decisions.

==== END FILE ====

==== FILE: tasks/update-prd.md ====
# Task: Update Product Requirements Document

## Description

Update an existing PRD with new requirements, scope changes, or refined business needs while maintaining traceability to Project Brief. Focus on change impact assessment and clear documentation to ensure requirements remain aligned with strategic objectives while defining epic-level features within the PRD.

## Prerequisites

- [ ] **Existing PRD**: `/docs/prd/prd.md` exists and is properly accessible
- [ ] **Change trigger**: Clear reason for update (Project Brief changes, user research, business priorities, technical constraints, stakeholder feedback)
- [ ] **Stakeholder input**: Understanding of what specifically needs to change and why
- [ ] **Epic/Story review**: Current understanding of feature groupings and requirements structure

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/prd-template.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

### CRITICAL: MANDATORY USER CONSULTATION FIRST

Before making ANY changes to the PRD, you MUST:

1. **Ask the user** what specific updates they want to make to the PRD
2. **Understand the trigger** for the changes (new requirements, stakeholder feedback, market changes, etc.)
3. **Clarify scope** which sections need updating and why
4. **Get approval** for the proposed changes before implementation
5. **Wait for explicit confirmation** before proceeding with any edits

### ONLY AFTER USER CONFIRMATION

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for change management process
2. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
3. **Format output**: Maintain [prd-template.md](./.krci-ai/templates/prd-template.md) structure
4. **Maintain traceability**: Update BR/NFR numbering and include epic-level feature definitions

## Output Format

- **Location**: Updates existing `/docs/prd/prd.md` (EXACT path and filename)
- **Length**: Maintain 6-8 pages maximum
- **Requirements Format**: Maintain BR1, BR2, BR3... and NFR1, NFR2, NFR3... numbering with P0/P1/P2 priority indicators and epic-level feature definitions
- **Impact Documentation**: Clear notes on what changed and feature impact
- **Downstream Updates**: List of feature areas requiring updates

## Success Criteria

- [ ] **File updated** at `/docs/prd/prd.md` reflects all changes
- [ ] **Requirements numbered** BR/NFR structure maintained with priority indicators and epic-level features
- [ ] **Change documented** clear record of what changed and why
- [ ] **Feature impact** identified which feature areas need updates
- [ ] **Quality maintained** document remains 6-8 pages maximum
- [ ] **Project Brief alignment** changes align with Project Brief updates (if any)
- [ ] **Stakeholder approval** key stakeholders have approved requirement changes

## Execution Checklist

### User Consultation Phase (MANDATORY FIRST STEP)

- [ ] **User interview**: Ask user what specific changes they want to make to the PRD
- [ ] **Change justification**: Understand why these changes are needed (stakeholder feedback, new requirements, market changes, etc.)
- [ ] **Scope definition**: Clarify which PRD sections need updating and what specific content changes are required
- [ ] **Impact discussion**: Explain potential impact on existing features to user
- [ ] **User approval**: Get explicit user confirmation before proceeding with any changes
- [ ] **Change plan agreement**: Confirm the proposed approach with user before implementation

### Assessment Phase (ONLY AFTER USER APPROVAL)

- [ ] **Change scope**: Identify which sections need updating based on user requirements
- [ ] **Impact analysis**: Evaluate how changes affect existing feature definitions and requirements structure
- [ ] **Stakeholder review**: Confirm who needs to approve these changes before implementation
- [ ] **Requirements mapping**: Understand which BR/NFR numbers and priorities are affected

### Requirements Phase

- [ ] **Business requirements**: Update BR1, BR2, BR3... with new business functionality needs
- [ ] **Non-functional requirements**: Update NFR1, NFR2, NFR3... with new system behavior/performance needs
- [ ] **Priority assessment**: Review and update P0/P1/P2 priority indicators as needed
- [ ] **Epic groupings**: Ensure updated requirements can be organized into logical epic-level features within the PRD

### Update Phase

- [ ] **Section updates**: Modify specific sections using [prd-template.md](./.krci-ai/templates/prd-template.md) structure
- [ ] **Content integration**: Ensure changes are properly integrated without breaking flow
- [ ] **Length verification**: Confirm document remains 6-8 pages maximum
- [ ] **Quality validation**: Verify all changes maintain PRD quality standards

### Change Management Phase

- [ ] **Feature impact assessment**: Determine which feature areas need updating based on requirement changes
- [ ] **Team communication**: Notify development teams of requirement changes
- [ ] **Documentation**: Record change rationale and feature impact plan

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Change Impact Focused**: Always assess feature impact before implementing PRD changes
- **Requirement Versioning**: Maintain BR/NFR numbering and priority consistency with epic-level feature definitions
- **Stakeholder Aligned**: Ensure all requirement changes have proper approval before implementation
- **Quality Preserved**: Keep updates within 6-8 page limit while maintaining user-centered focus

### LLM Error Prevention Checklist

- **NEVER**: Start making PRD changes without explicit user consultation and approval
- **NEVER**: Assume what changes the user wants - always ask for specific requirements
- **Avoid**: Breaking existing BR/NFR numbering that features depend on
- **Avoid**: Making changes without assessing feature impact
- **Avoid**: Updating requirements without proper stakeholder approval process
- **Always**: Wait for user confirmation before proceeding with any edits
- **Reference**: Use [prd-template.md](./.krci-ai/templates/prd-template.md) for all formatting consistency

### SDLC Integration Context

This update enables continued development by maintaining requirement traceability, preserving BR/NFR numbering with epic-level features, and communicating changes to development teams with clear impact assessment and timeline considerations.

==== END FILE ====

# Shared Templates

==== FILE: templates/prd-template.md ====
# Product Requirements Document: {{product_name}}

## 1. Problem/Opportunity

<!-- Be crisp and clear about what user or business problem you're solving -->
<!-- AVOID: "User can't use [solution]" - this is NOT a problem statement -->
<!-- FOCUS: What issues are caused when functionality is missing? -->

{{problem_statement}}

**Evidence:**
{{supporting_evidence}}

<!-- Template Guidance:
Problem Example: "Users frequently forget passwords leading to 15% of support tickets and account lockouts (2,500/month). This creates frustration and consumes 15% of support resources."

Evidence Example: "User research shows 65% of users reset passwords monthly. Support ticket analysis reveals 2,500 password-related tickets costing $50K annually."
-->

---

## 2. Target Users & Use Cases

<!-- Always focus on the user - this aligns building with go-to-market -->
<!-- Be specific about users and use cases, ensure team alignment on definitions -->

**Primary Users:**
{{primary_users}}

**Key Use Cases:**
{{key_use_cases}}

<!-- Template Guidance:
Primary Users Example: "SaaS platform users (10,000+ monthly active) who access platform 3+ times weekly. Primary segment: business professionals aged 25-45 accessing from mobile (60%) and desktop (40%)."

Key Use Cases Example:
1. Daily login for work tasks (highest frequency)
2. Password recovery when locked out (highest pain)
3. Multi-device access synchronization (growing need)
-->

---

## 3. Current Journeys/Landscape *(Optional)*

<!-- Give context on what users do today or how competitors solve this -->
<!-- Quick summary + links to detailed materials -->

**Current User Journey:**
{{current_journey}}

**Competitive Landscape:**
{{competitive_analysis}}

<!-- Template Guidance:
Current Journey Example: "Users must remember complex passwords, leading to frequent lockouts. Recovery process takes 5-10 minutes via email verification."

Competitive Analysis: "Auth0, Okta provide enterprise solutions. Consumer apps use Face ID/Touch ID. Gap exists for SMB-focused authentication."

Links: "[Detailed user journey flow](link)" or "[Competitive analysis doc](link)"
-->

---

## 4. Proposed Solution/Elevator Pitch

<!-- Standard 2-3 liner in plain English -->
<!-- Include top 3 MVP value props + conceptual model -->

**Elevator Pitch:**
{{elevator_pitch}}

**Top 3 MVP Value Props:**

1. {{value_prop_1}}
2. {{value_prop_2}}
3. {{value_prop_3}}

**Conceptual Model:**
{{conceptual_model}}

<!-- Template Guidance:
Elevator Pitch Example: "Enable users to login using biometric authentication (fingerprint/face) and social login options, reducing password dependency by 80% while maintaining enterprise security standards."

Value Props Example:
1. 3-second biometric login eliminates password frustration
2. Social login reduces new user signup friction
3. Enterprise security maintains compliance requirements

Conceptual Model: "[Include simple diagram or description of how users will interact with the solution]"
-->

---

## 5. Goals/Measurable Outcomes

<!-- Literally 2-3 bullets, no more -->
<!-- Measurable outcomes defining success or non-failure -->

**Success Metrics:**

1. {{success_metric_1}}
2. {{success_metric_2}}
3. {{success_metric_3}}

<!-- Template Guidance:
Success Metrics Example:
1. Reduce password-related support tickets by 80% within 3 months
2. Achieve 70% user adoption of new auth methods within 6 months
3. Improve login success rate from 85% to 95%

AVOID vague statements like "improve user experience" or "increase engagement"
-->

---

## 6. MVP/Functional Requirements

<!-- Critical: Focus on required functionality, save the rest for future phases -->
<!-- Question: What's the "min-viable" set of functionality for target user adoption? -->

### Business Requirements (BR)

**BR1**: {{business_requirement_1}}
**BR2**: {{business_requirement_2}}
**BR3**: {{business_requirement_3}}

### Non-Functional Requirements (NFR)

**NFR1**: {{system_requirement_1}}
**NFR2**: {{system_requirement_2}}
**NFR3**: {{system_requirement_3}}

<!-- Template Guidance:

Format: Focus on functionality, not implementation
✅ DO: "First-time user must accept privacy policy to use product"
✅ DO: "Product team can monitor and visualize user engagement"
✅ DO: Link to UX sketches for quick visualization
✅ DO: Include priorities: [P0] [P1] [P2] where P0 = truly required for MVP
✅ DO: Bucket by use case/user journey for Epic creation
✅ DO: Consider all critical user journeys (CUJs) - create, maintain, retire, navigate
✅ DO: Limit to 3 phases/milestones maximum

❌ DON'T: Performance metrics unless required for adoption
❌ DON'T: Design details like "blue 'Continue' button"
❌ DON'T: Technical implementation specifics

Business Requirements (BR) Examples:
BR1 [P0]: User can login using biometric authentication with <3 second response
BR2 [P1]: User can view login history with timestamps and device info
BR3 [P2]: Admin can configure password complexity requirements

Non-Functional Requirements (NFR) Examples:
NFR1 [P0]: System supports 1000 concurrent users with <2 second response time
NFR2 [P1]: System maintains 99.9% uptime during business hours
NFR3 [P2]: System integrates with enterprise SSO solutions

Use Case Buckets for Epic Creation:
### Epic 1: Authentication & Security
- BR1: Biometric authentication implementation
- NFR1: Performance and scalability requirements

### Epic 2: User Management
- BR2: User history and account features
- NFR2: System reliability requirements

Each bucket should map to an Epic following SDLC naming: {epic_number}-epic-{slug}.md
-->

==== END FILE ====

==== FILE: templates/project-brief-template.md ====
# Project Brief: {{project_name}}

> **Target Length**: 2-3 pages maximum
> **Framework**: Root artifact in SDLC framework
> **File Location**: MUST be saved as `/docs/prd/project-brief.md` (exact path)

---

## Executive Summary

{{executive_summary}}

<!-- Template Guidance:
Write a compelling 3-4 sentence overview combining problem, solution approach, and expected outcome.

Example: "Our SaaS platform experiences 2,500 password-related support tickets monthly, consuming 15% of support resources and frustrating users. We will implement biometric authentication and social login options to reduce password dependency, targeting 80% reduction in support tickets and $50K annual savings. This 3-month initiative serves 10,000+ monthly active users and requires Auth0 integration with a $25K budget."

Key Elements:
- What problem are we solving? (specific and quantified)
- How will we solve it? (high-level approach)
- What's the expected outcome? (business value)
- What's the scope? (timeline, users, constraints)
-->

---

## Problem Statement

{{problem_statement}}

<!-- Template Guidance:
Define the specific pain point driving this project with clear scope boundaries.

Example: "Users frequently forget passwords leading to 15% of support tickets and account lockouts (2,500/month). Focus on authentication workflow only, excluding password policy management or user registration processes."

Best Practices:
- Start with user scenarios, not business needs
- Use concrete numbers and evidence
- Define what's included and excluded
- Avoid solution-oriented language
- Focus on pain points and their impact

Evidence to Include:
- Support ticket volumes
- User research findings
- Productivity impact data
- Cost of current workarounds
-->

---

## Opportunity

{{opportunity}}

<!-- Template Guidance:
Quantified business value plus high-level solution approach.

Example: "Reducing password-related support tickets by 80% would save $50K annually and improve user satisfaction scores by 25%. Implement biometric authentication and social login options to reduce password dependency."

Key Elements:
- Business value (cost savings, revenue, efficiency)
- User value (time savings, satisfaction, productivity)
- Market opportunity (competitive advantage, growth)
- High-level solution direction (not detailed implementation)

Quantification Examples:
- Cost reduction: "$50K annual savings"
- Time savings: "15 minutes per user per month"
- Satisfaction: "25% improvement in user satisfaction"
- Efficiency: "80% reduction in support tickets"
-->

---

## Target Users

{{target_users}}

<!-- Template Guidance:
Specific user segments who have this problem with usage patterns and demographics.

Example: "SaaS platform users (10,000+ monthly active users) who access the platform 3+ times per week. Primary segment: business professionals aged 25-45 accessing from mobile devices (60%) and desktop (40%)."

Include:
- User volume and growth trends
- Demographics (age, role, industry)
- Usage patterns (frequency, device, context)
- Segment prioritization (primary vs secondary)
- Geographic distribution if relevant

User Segment Examples:
- "10,000+ monthly active users"
- "Business professionals aged 25-45"
- "Mobile-first users (60% mobile, 40% desktop)"
- "Access platform 3+ times weekly"
- "Located primarily in North America and Europe"
-->

---

## Success Metrics

{{success_metrics}}

<!-- Template Guidance:
How we'll measure if we've solved the problem with specific timelines.

Example: "Reduce password-related support tickets by 80% within 3 months, maintain 99.9% uptime, achieve 70% user adoption of new auth methods within 6 months, improve login success rate from 85% to 95%."

Success Criteria Format:
- Specific: Exactly what will be measured
- Measurable: Numbers, percentages, timelines
- Achievable: Realistic given constraints
- Relevant: Directly tied to problem and opportunity
- Time-bound: Clear deadlines

Metric Categories:
- Problem Resolution: "80% reduction in support tickets"
- User Adoption: "70% user adoption within 6 months"
- Quality: "99.9% uptime maintained"
- User Experience: "Login success rate 85% → 95%"
- Business Impact: "$50K annual cost savings"
-->

---

## Constraints

{{constraints}}

<!-- Template Guidance:
Resource, technical, and assumption factors that limit the solution.

Example: "Must integrate with existing Auth0 setup, 3-month timeline, $25K budget, maximum 2 developers assigned. Assumes current mobile app architecture supports biometric APIs and users have compatible devices."

Constraint Categories:

### Resource Constraints:
- Budget: "$25K maximum budget"
- Timeline: "3-month delivery deadline"
- Team: "Maximum 2 developers available"
- Skills: "No iOS development expertise on team"

### Technical Constraints:
- Integration: "Must integrate with existing Auth0"
- Architecture: "Cannot modify core database schema"
- Performance: "Must maintain current response times"
- Security: "Must meet enterprise security standards"

### Business Constraints:
- Compliance: "Must maintain SOC 2 compliance"
- User Impact: "Zero downtime deployment required"
- Support: "Cannot increase support complexity"
- Branding: "Must align with current UI/UX standards"

### Key Assumptions:
- "Users have biometric-capable devices"
- "Auth0 API will remain stable"
- "No major iOS/Android changes during development"
-->

---

## Key Risks

{{key_risks}}

<!-- Template Guidance:
Major risks that could derail the project with impact assessment.

Example: "User adoption resistance (HIGH): Users may prefer familiar passwords. Auth0 API changes (MEDIUM): Potential breaking changes during integration. Biometric compatibility (MEDIUM): Older devices may not support all features. Timeline risk (HIGH): Integration complexity may exceed estimates."

Risk Assessment Format:
[Risk Name] ([Impact Level]): [Description and potential impact]

Impact Levels:
- HIGH: Could significantly delay or derail project
- MEDIUM: Could cause delays or require scope changes
- LOW: Minor impact, manageable workarounds available

Risk Categories:

### User Adoption Risks:
- "User resistance to change (HIGH)"
- "Learning curve for new features (MEDIUM)"
- "Device compatibility issues (MEDIUM)"

### Technical Risks:
- "Integration complexity (HIGH)"
- "Third-party API changes (MEDIUM)"
- "Performance impact (LOW)"

### Business Risks:
- "Timeline overrun (HIGH)"
- "Budget overrun (MEDIUM)"
- "Resource unavailability (MEDIUM)"

### Market Risks:
- "Competitive response (LOW)"
- "Regulatory changes (MEDIUM)"
- "Technology shifts (LOW)"
-->

---

## SDLC Framework Information

**Dependencies**: None (root artifact)
**Output Location**: This Project Brief MUST be saved as `/docs/prd/project-brief.md`
**Downstream Enablement**: Enables PRD creation at `/docs/prd/prd.md`

<!-- SDLC Framework Integration:
This Project Brief serves as the foundation for:
- PRD Problem/Opportunity section
- PRD Target Users & Use Cases
- PRD Goals/Measurable Outcomes
- PRD scope and constraint definition

Directory Structure:
/docs/
├── prd/                          # Product vision & requirements
│   ├── project-brief.md          # Project vision & strategy (THIS FILE)
│   └── prd.md                    # Product requirements (ENABLED BY THIS)
├── epics/                        # High-level features
├── stories/                      # User stories
├── architecture/                 # System design
└── tests/                        # Quality validation
-->

---

<!-- QUALITY CHECKLIST
✅ Document is 2-3 pages maximum
✅ Executive summary captures complete project essence
✅ Problem statement is specific and evidence-based
✅ Opportunity is quantified with business value
✅ Target users are specific with usage patterns
✅ Success metrics are measurable with timelines
✅ Constraints are realistic and comprehensive
✅ Key risks identified with impact assessment
✅ File saved exactly as /docs/prd/project-brief.md
✅ Ready to enable PRD creation
-->
==== END FILE ====

# Reference Data

==== FILE: data/business-frameworks.md ====
# Business Analysis Frameworks and Models

## Requirements Analysis Frameworks

### BABOK (Business Analysis Body of Knowledge)

Comprehensive framework for business analysis practices and techniques.

- **Knowledge Areas**: Business Analysis Planning, Elicitation, Requirements Management, Solution Assessment
- **Techniques**: Interviews, Workshops, Document Analysis, Observation, Surveys
- **Deliverables**: Requirements Documentation, Stakeholder Analysis, Solution Assessment
- **Application**: Use for structured requirements gathering and analysis projects

### MoSCoW Prioritization

Framework for prioritizing requirements based on business importance.

- **Must Have**: Critical requirements without which the solution fails
- **Should Have**: Important requirements that add significant value
- **Could Have**: Desirable requirements that enhance the solution
- **Won't Have**: Requirements that are out of scope for current iteration

### Kano Model

Framework for understanding customer satisfaction with product features.

- **Must-be Quality**: Basic expectations that cause dissatisfaction if missing
- **One-dimensional Quality**: Features that increase satisfaction linearly
- **Attractive Quality**: Unexpected features that delight customers
- **Indifferent Quality**: Features that don't significantly impact satisfaction

## Process Analysis Frameworks

### Six Sigma DMAIC

Data-driven process improvement methodology.

- **Define**: Define project goals and customer requirements
- **Measure**: Measure current process performance and collect data
- **Analyze**: Analyze data to identify root causes of problems
- **Improve**: Implement solutions to address root causes
- **Control**: Monitor and control the improved process

### Lean Process Analysis

Framework focused on eliminating waste and optimizing value flow.

- **Value Stream Mapping**: Visualize entire process flow and identify waste
- **Waste Identification**: Identify and eliminate non-value-added activities
- **Flow Optimization**: Improve process flow and reduce cycle time
- **Pull Systems**: Implement demand-driven process execution

### SIPOC Analysis

Framework for understanding process scope and context.

- **Suppliers**: Entities that provide inputs to the process
- **Inputs**: Materials, information, and resources entering the process
- **Process**: Activities that transform inputs into outputs
- **Outputs**: Products, services, or information produced by the process
- **Customers**: Recipients or users of the process outputs

### Value Stream Mapping

Visual tool for analyzing and improving process flow.

- **Current State Map**: Document existing process flow and identify waste
- **Future State Map**: Design improved process with waste elimination
- **Implementation Plan**: Roadmap for transitioning to future state
- **Continuous Improvement**: Regular review and optimization cycles

## Stakeholder Analysis Frameworks

### RACI Matrix

Framework for defining roles and responsibilities in processes and projects.

- **Responsible**: Person who performs the activity or task
- **Accountable**: Person who is ultimately answerable for the activity
- **Consulted**: People who provide input and expertise
- **Informed**: People who need to be kept informed of progress

### Power-Interest Grid

Framework for stakeholder analysis and engagement strategy.

- **High Power, High Interest**: Key stakeholders requiring active management
- **High Power, Low Interest**: Stakeholders to keep satisfied
- **Low Power, High Interest**: Stakeholders to keep informed
- **Low Power, Low Interest**: Stakeholders requiring minimal effort

### Stakeholder Onion Diagram

Visual framework for mapping stakeholder relationships and influence.

- **Core**: Direct users and beneficiaries of the solution
- **Direct**: Stakeholders directly impacted by the solution
- **Indirect**: Stakeholders indirectly affected by the solution
- **External**: External stakeholders with potential influence

## Problem Analysis Frameworks

### Root Cause Analysis (5 Whys)

Systematic approach to identifying underlying causes of problems.

- **Problem Statement**: Clear definition of the observed problem
- **Why Analysis**: Repeatedly ask "why" to drill down to root causes
- **Cause Verification**: Validate identified causes with evidence
- **Solution Development**: Address root causes rather than symptoms

### Fishbone Diagram (Ishikawa)

Visual tool for systematic problem analysis and cause identification.

- **Problem Definition**: Clear statement of the effect or problem
- **Cause Categories**: People, Process, Technology, Environment, Materials
- **Brainstorming**: Generate potential causes within each category
- **Analysis**: Investigate and validate the most likely causes

### Force Field Analysis

Framework for analyzing forces supporting and opposing change.

- **Driving Forces**: Factors that support the desired change
- **Restraining Forces**: Factors that resist or oppose the change
- **Force Assessment**: Evaluate strength and impact of each force
- **Strategy Development**: Strengthen driving forces and mitigate restraining forces

## Solution Design Frameworks

### Design Thinking

Human-centered approach to innovation and solution development.

- **Empathize**: Understand user needs and perspectives
- **Define**: Synthesize observations into problem statements
- **Ideate**: Generate creative solution alternatives
- **Prototype**: Build testable representations of solutions
- **Test**: Validate solutions with users and stakeholders

### Jobs-to-be-Done Framework

Framework for understanding customer motivation and solution design.

- **Job Definition**: Identify the fundamental job customers are trying to accomplish
- **Job Mapping**: Break down the job into sequential steps
- **Outcome Identification**: Define desired outcomes for each job step
- **Solution Design**: Create solutions that help customers complete jobs better

## Business Model Analysis

### Business Model Canvas

Visual framework for analyzing and designing business models.

- **Value Propositions**: Benefits delivered to customers
- **Customer Segments**: Groups of customers with common needs
- **Channels**: How value propositions reach customers
- **Customer Relationships**: Types of relationships with customer segments
- **Revenue Streams**: How the business generates income
- **Key Resources**: Assets required to deliver value
- **Key Activities**: Critical activities for business success
- **Key Partnerships**: Network of suppliers and partners
- **Cost Structure**: Costs involved in operating the business

### Value Proposition Canvas

Framework for designing and testing value propositions.

- **Customer Profile**: Customer jobs, pains, and gains
- **Value Map**: Products/services, pain relievers, and gain creators
- **Fit Analysis**: Alignment between customer needs and value offering

==== END FILE ====

