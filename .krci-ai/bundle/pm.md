# KubeRocketAI Framework Bundle

**Generated:** 2025-10-21T16:42:51+03:00
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
    description: "Product manager for strategy/PRDs/roadmaps. Redirects implementation→dev, architecture→architect, stories→PO agents."
    role: "Senior Product Manager"
    goal: "Drive product success through strategic planning within PM scope"
    icon: "📈"

  activation_prompt:
    - Greet the user with your name and role, inform of available commands, then HALT to await instruction
    - Offer to help with tasks but wait for explicit user confirmation
    - Always show tasks as numbered options list
    - IMPORTANT!!! ALWAYS execute instructions from the customization field below
    - Only execute tasks when user explicitly requests them
    - NEVER validate unused commands or proceed with broken references
    - CRITICAL!!! Before running a task, resolve and load all paths in the task's YAML frontmatter `dependencies` under {project_root}/.krci-ai/{agents,tasks,data,templates}/**/*.md. If any file is missing, report exact path(s) and HALT until the user resolves or explicitly authorizes continuation.

  principles:
    - "SCOPE: Strategy/PRD/roadmap creation only. Redirect implementation→dev, architecture→architect, stories→PO."
    - "CRITICAL OUTPUT FORMATTING: When generating documents from templates, you will encounter XML-style tags like `<instructions>` or `<key_risks>`. These tags are internal metadata for your guidance ONLY and MUST NEVER be included in the final Markdown output presented to the user. Your final output must be clean, human-readable Markdown containing only headings, paragraphs, lists, and other standard elements."
    - "Always prioritize user value and business impact in product decisions"
    - "Ground decisions in data and user research rather than assumptions"
    - "Ask clarifying questions when requirements are ambiguous or incomplete"
    - "Provide evidence-based recommendations with clear rationale and trade-offs"
    - "Create comprehensive PRDs with clear acceptance criteria and success metrics"

  customization: ""

  commands:
    help: "Show available commands with numbered options"
    chat: "(Default) Product management consultation and guidance"

    # Project Brief Creation Commands
    create-project-brief: "Create project brief using standard workflow (2-3 pages, business framework based)"
    create-project-brief-advanced: "Create project brief using advanced validation flow (evidence-based, comprehensive validation)"

    # Project Brief Management Commands
    enhance-project-brief: "Upgrade existing standard brief to advanced validation flow"
    update-project-brief: "Update existing project brief by executing task update-project-brief"

    # Context Gathering Commands (Enhanced Flow)
    gather-context: "Collect structured project inputs using business frameworks by executing task gather-project-context"

    # Validation Commands
    validate-problem: "Validate problem statement using Lean Startup Problem-Solution Fit framework"
    validate-users: "Validate target users using Jobs-to-be-Done framework"
    validate-metrics: "Validate success metrics using SMART criteria and OKR alignment framework"
    validate-value: "Validate business value using Value Proposition Canvas framework"

    # Brief Enhancement Commands
    refine-brief: "Incorporate validation feedback and update project brief sections"
    finalize-brief: "Complete project brief when all validations are satisfied"

    # PRD Commands
    create-prd: "Create comprehensive product requirements document by executing task create-prd"
    update-prd: "Update existing product requirements document by executing task update-prd"

    exit: "Exit Product Manager persona and return to normal mode"

  tasks:
    # Project Brief Creation Tasks
    - ./.krci-ai/tasks/create-project-brief.md
    - ./.krci-ai/tasks/create-project-brief-advanced.md
    - ./.krci-ai/tasks/update-project-brief.md
    - ./.krci-ai/tasks/enhance-project-brief.md

    # Context Gathering (Enhanced Flow)
    - ./.krci-ai/tasks/gather-project-context.md

    # Validation Tasks
    - ./.krci-ai/tasks/validate-problem-statement.md
    - ./.krci-ai/tasks/validate-target-users.md
    - ./.krci-ai/tasks/validate-success-metrics.md
    - ./.krci-ai/tasks/validate-business-value.md

    # Brief Enhancement Tasks
    - ./.krci-ai/tasks/refine-project-brief.md
    - ./.krci-ai/tasks/finalize-project-brief.md

    # PRD Tasks
    - ./.krci-ai/tasks/create-prd.md
    - ./.krci-ai/tasks/update-prd.md

==== END FILE ====

==== FILE: .krci-ai/tasks/refine-project-brief.md ====
---
dependencies:
  data:
    - validation-frameworks.md
  templates:
    - project-brief-template.md
    - assumption-tracker-template.md
---

# Task: Refine Project Brief

## Description

Incorporate validation feedback and evidence into the project brief, updating sections with validated information, improved confidence levels, and refined assumptions. This task synthesizes multiple validation results to create an enhanced, evidence-based project brief.

This task uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and may utilize the [assumption tracker template](./.krci-ai/templates/assumption-tracker-template.md) for tracking refinements.

## Instructions

<instructions>
Confirm the original project brief exists at `/docs/prd/project-brief.md`, one or more validation reports are completed, assumption tracker is updated with validation results, and stakeholder feedback on validation results is available. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Collect and analyze all completed validation results from `/docs/prd/brief-validation-*.md`. Review validation reports to extract key findings that impact project brief sections. Identify confidence level changes based on validation evidence and document required brief updates. Aggregate evidence quality scores across all validations, identify strongest and weakest evidence areas, update overall brief confidence based on evidence assessment, and flag areas requiring additional validation.

Update project brief sections systematically with validation results. Enhance executive summary by updating problem description with validated evidence, refining solution approach based on validation findings, updating business value projections with validated metrics, and enhancing scope definition with validated constraints. Refine problem statement by integrating root cause analysis findings, adding quantified problem evidence from validation, updating problem scope based on validation boundaries, and enhancing impact assessment with validated metrics.

Enhance target users section by updating user segments with Jobs-to-be-Done validation, adding validated demographic and behavioral data, including opportunity scores and user priorities, and refining user context with journey validation insights. Improve success metrics by updating metrics with SMART criteria validation results, adding baseline data discovered during validation, refining targets based on evidence and benchmarks, and balancing leading and lagging indicators appropriately.

Update constraints and risks by adding constraints discovered during validation, updating risk assessment with validation insights, including evidence-based risk probability estimates, and enhancing mitigation strategies with validated approaches. For each section, update validation checkpoints by marking completed validations as verified, updating confidence levels based on evidence quality, documenting validation methods used, and including evidence sources and quality assessments.

Integrate assumption tracker updates by updating assumption status based on validation results, adding new assumptions discovered during validation, removing or modifying disproven assumptions, and prioritizing remaining assumptions for future validation. Ensure 2-3 page length limit maintained, verify executive-ready language and structure, confirm all sections have appropriate evidence support, and validate SDLC integration requirements met.

Prepare stakeholder review by creating summary of changes made based on validation, highlighting areas of increased or decreased confidence, documenting remaining uncertainties and validation needs, and creating stakeholder presentation of refined brief.
</instructions>

## Output Format

<output_format>
- Primary Output: Updated `/docs/prd/project-brief.md` with validation integration
- Secondary Outputs:
  - `/docs/prd/brief-refinement-summary.md` documenting changes made
  - Updated `/docs/prd/brief-assumptions.md` with current status
  - `/docs/prd/validation-dashboard.md` with overall validation status
- Length: Project brief remains 2-3 pages maximum
</output_format>

## Success Criteria

<success_criteria>
- All validation results integrated into appropriate brief sections
- Confidence levels updated based on evidence quality assessment
- Assumption tracker synchronized with brief content and validation results
- Brief length maintained within 2-3 page executive limit
- Evidence sources documented with quality assessments
- Remaining validation needs identified and prioritized
</success_criteria>

## Execution Checklist

<execution_checklist>

### Validation Integration

- Collect validation reports: Gather all completed validation documents
- Extract key findings: Identify findings that require brief updates
- Assess evidence quality: Evaluate strength of validation evidence
- Plan section updates: Determine which sections need refinement using [project brief template](./.krci-ai/templates/project-brief-template.md) structure

### Brief Refinement

- Update executive summary: Integrate high-level validation insights
- Refine problem statement: Add validated evidence and refined scope
- Enhance user sections: Include validated user research and insights
- Improve success metrics: Update with validated baselines and targets
- Update constraints/risks: Include validation-discovered factors

### Quality Assurance

- Length verification: Ensure brief remains within 2-3 page limit
- Evidence documentation: Verify all claims have supporting evidence cited
- Assumption alignment: Ensure brief content aligns with assumption tracker
- Stakeholder readiness: Confirm brief is executive-ready and decision-enabling

</execution_checklist>

## Content Guidelines

### Integration Principles

- Evidence Primacy: Prioritize validated evidence over original assumptions
- Transparency: Clearly indicate confidence levels and evidence sources
- Conciseness: Maintain executive brevity while including essential validation insights
- Actionability: Ensure refined brief enables clear next steps and decisions

### Quality Standards

- Professional Presentation: Maintain executive-level language and structure
- Evidence Attribution: Clearly cite validation sources and evidence quality
- Balanced Perspective: Include both confirming and challenging validation results
- Future-Focused: Identify areas needing continued validation and monitoring

==== END FILE ====

==== FILE: .krci-ai/tasks/finalize-project-brief.md ====
---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
    - validation-frameworks.md
---

# Task: Finalize Project Brief

## Description

Complete the project brief creation process by ensuring all validations are satisfied, stakeholders have approved content, and the brief meets all quality standards for downstream SDLC artifact creation.

This task may reference [validation frameworks](./.krci-ai/data/validation-frameworks.md) for final quality checks.

## Instructions

<instructions>
Confirm the project brief exists with validation checkpoints completed, all required validations are completed with satisfactory results, stakeholder review and feedback is incorporated, and assumption tracker is updated with current status. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Complete the project brief creation process by ensuring all validations are satisfied, stakeholders have approved content, and the brief meets all quality standards for downstream SDLC artifact creation.

### Phase 1: Quality Gate Assessment

Verify all quality gates have been satisfied:

#### Validation Completeness Check

- All validation checkpoints marked as completed
- Validation confidence levels meet minimum thresholds
- Critical assumptions validated with strong evidence
- High-impact risks assessed and mitigated

#### Stakeholder Approval Verification

- Key stakeholders have reviewed and approved content
- Feedback incorporated and documented
- No blocking concerns remain unaddressed
- Authorization to proceed to PRD creation obtained

### Phase 2: Final Brief Quality Review

#### Content Quality Assessment

- Executive summary captures complete project essence
- Problem statement is specific and evidence-based
- Target users are clearly defined with validated research
- Success metrics are measurable with realistic timelines
- Constraints reflect actual limitations and assumptions
- Risks identified with appropriate mitigation strategies

#### Technical Quality Check

- Document length is 2-3 pages maximum
- File saved exactly as `/docs/prd/project-brief.md`
- All required sections completed per template
- Evidence sources properly cited and accessible
- Validation status clearly documented

### Phase 3: SDLC Readiness Verification

#### Downstream Enablement Check

- Brief provides sufficient context for PRD creation
- Success metrics enable measurable outcomes
- User definitions support Epic and Story creation
- Technical considerations inform Architecture planning
- Business value justifies resource investment

#### Documentation Package Completion

- Project brief finalized and approved
- Supporting validation reports completed
- Assumption tracker current and maintained
- Evidence library documented and accessible
</instructions>

## Output Format

<output_format>
- Primary Output: Final approved `/docs/prd/project-brief.md`
- Status Update: Mark brief status as "APPROVED - Ready for PRD"
- Handoff Package: Complete documentation set for next phase
- Length: 2-3 pages maximum maintained
</output_format>

## Success Criteria

<success_criteria>
- All quality gates satisfied with documented evidence
- Stakeholder approval obtained with sign-off documentation
- Brief quality verified against all template requirements
- SDLC readiness confirmed for downstream artifact creation
- Documentation package complete for PRD development phase
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Finalization Review

- Validation status check: Verify all required validations completed
- Quality assessment: Review brief against all quality criteria
- Stakeholder coordination: Ensure all approvals obtained
- Documentation review: Verify supporting materials complete

### Final Quality Assurance

- Content verification: Confirm all sections meet quality standards
- Technical check: Verify format, length, and file location requirements
- Evidence validation: Ensure all claims properly supported and cited
- SDLC integration: Confirm readiness for next phase activities

### Finalization and Handoff

- Status update: Mark brief as approved and ready for PRD
- Package completion: Ensure all supporting documents finalized
- Handoff preparation: Ready documentation for next phase team
- Success communication: Notify stakeholders of completion

</execution_checklist>

## Finalization Criteria

### Minimum Validation Requirements

- Problem statement validated with confidence >70%
- Target users validated with evidence from >10 user interactions
- Success metrics validated using SMART criteria
- Business value validated with quantified ROI analysis

### Quality Thresholds

- Document length: 2-3 pages (strict requirement)
- Evidence citation: >80% of claims supported by documented sources
- Stakeholder approval: All key decision makers signed off
- Assumption risk: <3 high-risk assumptions with validation plans

### SDLC Integration Standards

- Enables PRD creation without additional research
- Provides clear success criteria for downstream teams
- Includes sufficient user context for Story development
- Contains technical considerations for Architecture planning

==== END FILE ====

==== FILE: .krci-ai/tasks/create-prd.md ====
---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
    - business-frameworks.md
  templates:
    - prd-template.md
---

# Task: Create Product Requirements Document (PRD)

## Description

Create a streamlined PRD that drives team alignment on what to build and why, following the proven 6-8 page structure focused on user needs and business value rather than technical specifications. This PRD includes epic-level feature definitions while maintaining clear traceability from Project Brief.

## Instructions

<instructions>
Confirm the exact output path `/docs/prd/prd.md` you will create or update. Verify that the Project Brief at `/docs/prd/project-brief.md` is accessible, along with market research, user insights, stakeholder requirements, and dependencies declared in the YAML frontmatter. Do not proceed if required inputs are missing.

Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for PRD workflow and quality gates. Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md). Use [prd-template.md](./.krci-ai/templates/prd-template.md) and populate all variables precisely, maintaining traceability to the Project Brief and including epic-level feature definitions with BR/NFR numbering and P0/P1/P2 priorities.
</instructions>

## Output Format

- Location: `/docs/prd/prd.md` (EXACT path and filename)
- Length: 6-8 pages maximum for team consumption
- Requirements Format: Use BR1, BR2, BR3... for business requirements and NFR1, NFR2, NFR3... for system requirements with P0/P1/P2 priority indicators and epic-level feature definitions
- Downstream Enable: Provides clear requirements structure for development teams

## Success Criteria

<success_criteria>
- File saved to `/docs/prd/prd.md`
- Length is 6-8 pages maximum
- Requirements numbered (BR1, BR2, NFR1, NFR2) with priority indicators and epic-level features
- Project Brief link clear connection to problem/opportunity
- Feature structure requirements organized into logical epic-level themes
- User focus prioritizes user needs over technical implementation details
- Stakeholder alignment all key requirements captured and validated
</success_criteria>

## Execution Checklist

### Discovery Phase

<discovery_phase>
- Problem analysis: Extract core problem from Project Brief
- User research: Conduct user interviews and usage analysis
- Competitive analysis: Research existing solutions and gaps
- Stakeholder alignment: Validate requirements with key stakeholders
</discovery_phase>

### Requirements Phase

<requirements_phase>
- Business requirements: Define BR1, BR2, BR3... (what business functionality is needed)
- Non-functional requirements: Define NFR1, NFR2, NFR3... (how system should behave/perform)
- Priority assignment: Add P0/P1/P2 priority indicators to each requirement
- Epic groupings: Structure requirements into logical epic-level feature themes within the PRD
</requirements_phase>

### Design Phase

<design_phase>
- Solution approach: High-level solution direction (not technical details)
- MVP scope: Define minimum viable product features
- Out of scope: Clearly document what's excluded
- Dependencies: Identify external requirements and constraints
</design_phase>

### Documentation Phase

<documentation_phase>
- PRD creation: Use [prd-template.md](./.krci-ai/templates/prd-template.md) structure
- Content validation: Ensure all required sections completed
- Length verification: Confirm document is 6-8 pages maximum
- File placement: Save to exact location `/docs/prd/prd.md`
</documentation_phase>

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- User-Centered: Always prioritize user needs over technical implementation details
- Evidence-Based: Support all requirements with user research and business data
- Traceable: Maintain clear connection from Project Brief → PRD with epic-level features
- Measurable: Ensure all success metrics are specific, testable, and time-bound

### LLM Error Prevention Checklist

- Avoid: Technical implementation details (save for Architecture documents)
- Avoid: Solution-oriented problem statements (focus on user pain points)
- Avoid: Vague requirements that cannot be grouped into epic-level features
- Reference: Use [prd-template.md](./.krci-ai/templates/prd-template.md) for all formatting guidance and examples

### SDLC Integration Context

This PRD provides numbered requirements (BR1, BR2, NFR1...) with priorities organized into epic-level feature themes, requirement groupings that structure development work, and success metrics that guide implementation decisions.

==== END FILE ====

==== FILE: .krci-ai/tasks/update-prd.md ====
---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
    - business-frameworks.md
  templates:
    - prd-template.md
---

# Task: Update Product Requirements Document

## Description

Update an existing PRD with new requirements, scope changes, or refined business needs while maintaining traceability to Project Brief. Focus on change impact assessment and clear documentation to ensure requirements remain aligned with strategic objectives while defining epic-level features within the PRD.

## Instructions

<instructions>
Confirm `/docs/prd/prd.md` exists and is properly accessible, there is clear reason for update (Project Brief changes, user research, business priorities, technical constraints, stakeholder feedback), you understand what specifically needs to change and why, and you have current understanding of feature groupings and requirements structure. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

CRITICAL: MANDATORY USER CONSULTATION FIRST - Before making ANY changes to the PRD, you MUST ask the user what specific updates they want to make, understand the trigger for the changes (new requirements, stakeholder feedback, market changes, etc.), clarify scope which sections need updating and why, get approval for the proposed changes before implementation, and wait for explicit confirmation before proceeding with any edits.

ONLY AFTER USER CONFIRMATION: Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for change management process. Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md). Maintain [prd-template.md](./.krci-ai/templates/prd-template.md) structure. Update BR/NFR numbering and include epic-level feature definitions.
</instructions>

## Output Format

- Location: Updates existing `/docs/prd/prd.md` (EXACT path and filename)
- Length: Maintain 6-8 pages maximum
- Requirements Format: Maintain BR1, BR2, BR3... and NFR1, NFR2, NFR3... numbering with P0/P1/P2 priority indicators and epic-level feature definitions
- Impact Documentation: Clear notes on what changed and feature impact
- Downstream Updates: List of feature areas requiring updates

## Success Criteria

<success_criteria>
- File updated at `/docs/prd/prd.md` reflects all changes
- Requirements numbered BR/NFR structure maintained with priority indicators and epic-level features
- Change documented clear record of what changed and why
- Feature impact identified which feature areas need updates
- Quality maintained document remains 6-8 pages maximum
- Project Brief alignment changes align with Project Brief updates (if any)
- Stakeholder approval key stakeholders have approved requirement changes
</success_criteria>

## Execution Checklist

### User Consultation Phase (MANDATORY FIRST STEP)

- User interview: Ask user what specific changes they want to make to the PRD
- Change justification: Understand why these changes are needed (stakeholder feedback, new requirements, market changes, etc.)
- Scope definition: Clarify which PRD sections need updating and what specific content changes are required
- Impact discussion: Explain potential impact on existing features to user
- User approval: Get explicit user confirmation before proceeding with any changes
- Change plan agreement: Confirm the proposed approach with user before implementation

### Assessment Phase (ONLY AFTER USER APPROVAL)

- Change scope: Identify which sections need updating based on user requirements
- Impact analysis: Evaluate how changes affect existing feature definitions and requirements structure
- Stakeholder review: Confirm who needs to approve these changes before implementation
- Requirements mapping: Understand which BR/NFR numbers and priorities are affected

### Requirements Phase

<requirements_validation>
- Business requirements: Update BR1, BR2, BR3... with new business functionality needs
- Non-functional requirements: Update NFR1, NFR2, NFR3... with new system behavior/performance needs
- Priority assessment: Review and update P0/P1/P2 priority indicators as needed
- Epic groupings: Ensure updated requirements can be organized into logical epic-level features within the PRD
</requirements_validation>

### Update Phase

- Section updates: Modify specific sections using [prd-template.md](./.krci-ai/templates/prd-template.md) structure
- Content integration: Ensure changes are properly integrated without breaking flow
- Length verification: Confirm document remains 6-8 pages maximum
- Quality validation: Verify all changes maintain PRD quality standards

### Change Management Phase

<change_management>
- Feature impact assessment: Determine which feature areas need updating based on requirement changes
- Team communication: Notify development teams of requirement changes
- Documentation: Record change rationale and feature impact plan
</change_management>

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

<quality_principles>
- Change Impact Focused: Always assess feature impact before implementing PRD changes
- Requirement Versioning: Maintain BR/NFR numbering and priority consistency with epic-level feature definitions
- Stakeholder Aligned: Ensure all requirement changes have proper approval before implementation
- Quality Preserved: Keep updates within 6-8 page limit while maintaining user-centered focus
</quality_principles>

### LLM Error Prevention Checklist

- NEVER: Start making PRD changes without explicit user consultation and approval
- NEVER: Assume what changes the user wants - always ask for specific requirements
- Avoid: Breaking existing BR/NFR numbering that features depend on
- Avoid: Making changes without assessing feature impact
- Avoid: Updating requirements without proper stakeholder approval process
- Always: Wait for user confirmation before proceeding with any edits
- Reference: Use [prd-template.md](./.krci-ai/templates/prd-template.md) for all formatting consistency

### SDLC Integration Context

This update enables continued development by maintaining requirement traceability, preserving BR/NFR numbering with epic-level features, and communicating changes to development teams with clear impact assessment and timeline considerations.

==== END FILE ====

==== FILE: .krci-ai/tasks/update-project-brief.md ====
---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
    - business-frameworks.md
  templates:
    - project-brief-template.md
---

# Task: Update Project Brief

## Description

Update an existing project brief with new information, scope changes, or refined understanding while maintaining strategic alignment and enabling downstream SDLC artifacts. Focus on change impact assessment and downstream artifact management to ensure existing PRD and Epic artifacts remain aligned with strategic changes.

## Instructions

<instructions>
Confirm `/docs/prd/project-brief.md` exists and is properly accessible, there is clear reason for update (strategic shifts, market changes, new insights, stakeholder feedback, resource changes), you understand how changes affect dependent PRD and downstream artifacts, and key stakeholders are aware of planned strategic changes. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

CRITICAL: MANDATORY USER CONSULTATION FIRST - Before making ANY changes to the Project Brief, you MUST ask the user what specific updates they want to make, understand the trigger for the changes (strategic shifts, market changes, stakeholder feedback, resource changes, etc.), clarify scope which sections need updating and why, get approval for the proposed changes before implementation, and wait for explicit confirmation before proceeding with any edits.

ONLY AFTER USER CONFIRMATION: Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for change impact assessment. Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md). Maintain [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) structure. Identify which PRD artifacts need updates.
</instructions>

## Output Format

- Location: Updates existing `/docs/prd/project-brief.md` (EXACT path and filename)
- Length: Maintain 2-3 pages maximum
- Impact Documentation: Clear notes on what changed and downstream impact
- Downstream Updates: List of PRD artifacts requiring updates

## Success Criteria

<success_criteria>
- File updated at `/docs/prd/project-brief.md` reflects all changes
- Change documented with clear record of what changed and why
- Downstream impact identified which PRD artifacts need updates
- Quality maintained document remains 2-3 pages maximum
- Strategic alignment changes support overall product strategy
- Stakeholder communication key stakeholders informed of strategic changes
</success_criteria>

## Execution Checklist

### User Consultation Phase (MANDATORY FIRST STEP)

- User interview: Ask user what specific changes they want to make to the Project Brief
- Change justification: Understand why these changes are needed (strategic shifts, market changes, stakeholder feedback, resource changes, etc.)
- Scope definition: Clarify which Project Brief sections need updating and what specific content changes are required
- Impact discussion: Explain potential impact on existing PRD artifacts to user
- User approval: Get explicit user confirmation before proceeding with any changes
- Change plan agreement: Confirm the proposed approach with user before implementation

### Assessment Phase (ONLY AFTER USER APPROVAL)

<strategic_assessment>
- Change scope: Identify which sections need updating based on user requirements (Executive Summary, Problem, Opportunity, Users, Success Metrics, Constraints, Risks)
- Business impact: Analyze how changes affect product strategy and business case
- Downstream impact: Evaluate how changes affect existing PRD (`/docs/prd/prd.md`) artifacts
- Stakeholder validation: Confirm changes with key stakeholders
</strategic_assessment>

### Update Phase

- Section updates: Modify specific sections using [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) structure
- Strategic alignment: Ensure updates maintain strategic coherence and business focus
- Quality check: Verify updated Project Brief maintains 2-3 page limit and foundation quality
- Content validation: Ensure all changes are properly integrated

### Change Management Phase

<downstream_impact>
- PRD impact analysis: Determine if PRD needs updating based on Project Brief changes
- Stakeholder communication: Notify key stakeholders of strategic changes and implications
- Documentation: Record change rationale and downstream impact plan
</downstream_impact>

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

<strategic_principles>
- Strategic Focus: Focus on strategic foundation changes rather than tactical adjustments
- Foundation Strength: Ensure changes strengthen rather than weaken the overall strategic foundation
- Cascade Management: Assess how strategic changes flow through PRD requirements
- Long-term Alignment: Consider long-term strategic implications beyond immediate tactical changes
</strategic_principles>

### LLM Error Prevention Checklist

- NEVER: Start making Project Brief changes without explicit user consultation and approval
- NEVER: Assume what changes the user wants - always ask for specific requirements
- Avoid: Making changes without clear strategic justification and stakeholder approval
- Avoid: Updating without assessing downstream PRD impact
- Avoid: Expanding scope beyond strategic foundation changes into tactical details
- Always: Wait for user confirmation before proceeding with any edits
- Reference: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) for all formatting consistency

### SDLC Integration Context

This update enables continued strategic alignment by managing strategic changes flowing through PRD requirements, ensuring stakeholder approval of strategic changes, and maintaining clear documentation of strategic change rationale and downstream PRD impact.

==== END FILE ====

==== FILE: .krci-ai/tasks/validate-target-users.md ====
---
dependencies:
  data:
    - validation-frameworks.md
    - business-frameworks.md
  templates:
    - validation-report-template.md
---

# Task: Validate Target Users

## Description

Apply Jobs-to-be-Done framework to validate target user segments, their motivations, and the value proposition alignment. This validation ensures user segments are accurately defined, their needs are properly understood, and the solution approach resonates with their actual jobs, pains, and gains.

This validation uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and outputs results using the [validation report template](./.krci-ai/templates/validation-report-template.md).

## Instructions

<instructions>
Confirm project brief with target users section exists, access to target users for interviews or surveys is available, user analytics or behavioral data is accessible, and market segmentation data or competitive user research is available. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Apply Jobs-to-be-Done framework from [validation-frameworks.md](./.krci-ai/data/validation-frameworks.md) to validate target user segments, their motivations, and the value proposition alignment. Extract and structure current user segment definitions from project brief including primary user segment (demographics, behaviors, needs), secondary user segments, user context (when, where), and user goals (objectives, success criteria). Use [validation-report-template.md](./.krci-ai/templates/validation-report-template.md) for output.
</instructions>

## User Hypothesis Structure

```markdown
Primary User Segment Hypothesis: {{user_demographic}} who {{user_behavior}} because they need to {{user_goal}}.

Specific Claims to Validate:
- User segment size: {{segment_size_claim}}
- User behavior patterns: {{behavior_pattern_claims}}
- User needs and pain points: {{needs_pain_claims}}
- User willingness to adopt solutions: {{adoption_claims}}
```

### Phase 2: Jobs-to-be-Done Framework Application

#### Step 1: Job Statement Definition and Validation

Job Statement Construction:
For each user segment, create job statements in the format:
"When I [situation], I want to [motivation], so I can [expected outcome]"

Job Dimensions Analysis:
- Functional Job: The practical task the user is trying to accomplish
- Emotional Job: How the user wants to feel or avoid feeling
- Social Job: How the user wants to be perceived by others

Job Statement Validation Process:
- Conduct user interviews to validate job statements
- Assess job importance (1-10) through user ranking exercises
- Evaluate current satisfaction (1-10) with existing solutions
- Calculate opportunity score: Importance + max(Importance - Satisfaction, 0)

#### Step 2: User Journey and Touchpoint Validation

User Journey Mapping:
- Map complete user journey from problem awareness to solution adoption
- Identify all touchpoints where users interact with current solutions
- Document user actions, thoughts, emotions at each stage
- Validate journey accuracy through user observation or interviews

Pain Point Identification and Validation:
- Identify specific pain points at each journey stage
- Quantify pain point impact (time lost, money spent, frustration level)
- Validate pain point significance through user prioritization exercises
- Assess pain point frequency and universality across user segment

Opportunity Identification:
- Identify improvement opportunities at each journey stage
- Assess opportunity impact on user satisfaction and business value
- Validate opportunity desirability through user feedback
- Prioritize opportunities by impact and feasibility

#### Step 3: User Segment Validation

Demographic and Behavioral Validation:
- Validate user segment demographics through market research data
- Confirm behavioral patterns through analytics and user observation
- Assess segment size and growth trends through industry data
- Validate segment accessibility and reachability for marketing

Needs and Goals Validation:
- Confirm user needs through structured interviews and surveys
- Validate goal priorities through user ranking and trade-off exercises
- Assess need intensity and urgency through user behavior analysis
- Confirm alignment between stated and revealed preferences

### Phase 3: Evidence Collection and Analysis

#### Primary Evidence Collection

User Interview Program:
- Conduct minimum 8-10 interviews per primary user segment
- Use structured interview guide based on Jobs-to-be-Done methodology
- Document specific examples and quantified impacts
- Record user quotes and insights for validation report

User Survey Deployment:
- Design survey to quantify job importance and satisfaction ratings
- Include demographic and behavioral questions for segmentation
- Achieve statistically significant sample size (minimum 100 responses)
- Analyze results for segment patterns and validation insights

User Analytics Analysis:
- Analyze existing user behavior data to validate segment assumptions
- Identify usage patterns that support or challenge segment definitions
- Quantify user engagement and retention metrics by segment
- Validate user journey stages through behavioral flow analysis

#### Secondary Evidence Collection

Market Research Integration:
- Leverage existing market segmentation studies and reports
- Identify industry benchmarks for similar user segments
- Validate segment size estimates through multiple data sources
- Assess competitive positioning targeting similar segments

Competitive User Analysis:
- Analyze competitor user bases and targeting strategies
- Identify overlapping segments and positioning differences
- Assess competitive user satisfaction and switching behavior
- Validate user needs through competitive feature analysis

### Phase 4: Jobs-to-be-Done Analysis Framework

#### Job Importance and Satisfaction Scoring

Importance Assessment (1-10 scale):
- How important is this job for the user's success?
- How frequently does the user need to complete this job?
- What's the impact on the user if this job isn't completed well?
- How much time/resources does the user invest in this job?

Satisfaction Assessment (1-10 scale):
- How satisfied is the user with current solution options?
- How well do current solutions complete the job?
- What gaps exist in current solution capabilities?
- How much effort is required to use current solutions?

Opportunity Score Calculation:

```
Opportunity Score = Importance + max(Importance - Satisfaction, 0)

Interpretation:
- 15+: High opportunity (underserved job with high importance)
- 12-15: Moderate opportunity (room for improvement exists)
- <12: Low opportunity (adequately served or low importance)
```

#### User Segment Validation Assessment

Segment Definition Quality:
- Specificity: Can you clearly identify who belongs to this segment?
- Measurability: Can you quantify segment size and characteristics?
- Accessibility: Can you reach this segment through marketing channels?
- Actionability: Can you create distinct value propositions for this segment?

Evidence Quality Assessment:
- Primary Research Percentage: How much validation comes from direct user research?
- Sample Size Adequacy: Is the research sample representative and significant?
- Evidence Recency: How current is the supporting evidence?
- Cross-Source Validation: Are findings consistent across multiple evidence types?

### Phase 5: Validation Results and Recommendations

#### User Segment Validation Results

Validated Segments:
- Segments supported by strong evidence across multiple sources
- Clear job definitions with high importance and low satisfaction scores
- Measurable segment characteristics and accessible through known channels
- Strong alignment between user needs and proposed solution approach

Challenged Segments:
- Segments with weak or conflicting evidence
- Unclear or low-importance job definitions
- Difficult to measure or access through available channels
- Poor alignment between stated needs and solution approach

New Segments Identified:
- Additional user groups discovered during validation research
- Segments with high opportunity scores not previously considered
- Adjacent segments with similar jobs but different contexts
- Segments that could be served with minor solution modifications

#### Project Brief Update Recommendations

User Section Enhancements:
- Update user segment definitions with validated evidence
- Add quantified user data (segment sizes, behavior patterns, satisfaction levels)
- Include validated job statements and opportunity scores
- Enhance user journey understanding with pain point and opportunity details

Confidence Level Updates:
- Assign confidence levels to all user segment claims based on evidence quality
- Identify areas requiring additional validation or research
- Update assumption tracker with validated, challenged, and new assumptions
- Prioritize remaining user validation activities

## Output Format

<output_format>
- Primary Output: `/docs/prd/brief-validation-users.md` using validation report template
- Secondary Outputs:
  - Updated `/docs/prd/project-brief.md` with validated user information
  - Updated `/docs/prd/brief-assumptions.md` with user assumption status
  - `/docs/prd/user-research-summary.md` with detailed research findings
- Length: Validation report 2-3 pages, supporting documents as needed
</output_format>

## Success Criteria

<success_criteria>

### Phase 1: User Hypothesis Definition

- User segment hypotheses extracted from current project brief
- Specific validation claims identified with measurable criteria
- User context and goals documented for validation activities
- Primary and secondary segments prioritized for validation focus

### Phase 2: Jobs-to-be-Done Framework Application

- Job statements created for each user segment using proper format
- Job dimensions identified (functional, emotional, social) with evidence
- User journey mapped with touchpoints, pain points, and opportunities
- Opportunity scores calculated using framework methodology

### Phase 3: Evidence Collection

- User interviews completed with minimum 8-10 per primary segment
- User survey deployed with statistically significant sample size
- Analytics analysis conducted validating behavioral assumptions
- Market research integrated for segment sizing and validation

### Phase 4: Jobs-to-be-Done Analysis

- Importance and satisfaction scored for all job statements
- Opportunity scores calculated and prioritized by potential impact
- Segment quality assessed using specificity, measurability, accessibility, actionability
- Evidence quality evaluated with confidence levels assigned

### Phase 5: Results and Recommendations

- Validation report completed using standardized template
- Project brief updated with validated user segment information
- Assumption tracker updated with user validation results
- Recommendations provided for user-focused solution development
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Validation Preparation

- Current user analysis: Review existing user segment definitions and evidence
- Research planning: Design interview guide, survey, and analytics analysis plan
- Participant recruitment: Identify and recruit representative users for research
- Tools preparation: Set up interview recording, survey platform, analytics access

### User Research Execution

- Interview conduct: Complete structured interviews with diverse user representatives
- Survey deployment: Launch survey and monitor response rates and quality
- Analytics analysis: Extract relevant user behavior data and patterns
- Competitive research: Analyze competitor user bases and positioning

### Jobs-to-be-Done Analysis

- Job statement validation: Confirm job statements with user feedback
- Importance/satisfaction scoring: Collect quantified ratings from users
- Opportunity calculation: Apply framework methodology to identify high-opportunity jobs
- User journey validation: Confirm journey accuracy through user input

### Results Synthesis and Documentation

- Evidence synthesis: Combine findings from multiple research methods
- Validation scoring: Apply segment quality and evidence quality assessments
- Report creation: Complete validation report with findings and recommendations
- Stakeholder communication: Present results to product team and stakeholders

</execution_checklist>

## Content Guidelines

### User Research Quality Standards

- Representative Sampling: Ensure research participants represent target segment diversity
- Unbiased Research: Use open-ended questions and avoid leading participants toward expected answers
- Quantified Insights: Collect measurable data wherever possible (time, frequency, cost, satisfaction)
- Multiple Method Validation: Validate key findings through multiple research methods
- Evidence Documentation: Properly document and cite all research sources and methodologies

### Jobs-to-be-Done Best Practices

- Outcome Focus: Focus on outcomes users want to achieve, not features they want
- Context Importance: Document when, where, and why users have specific jobs
- Emotional Dimensions: Don't ignore emotional and social aspects of user jobs
- Job Evolution: Understand how user jobs change over time or in different contexts
- Solution Agnostic: Define jobs independent of current solution approaches

### User Segment Definition Standards

- Actionable Segments: Define segments that enable different marketing or product approaches
- Measurable Characteristics: Include quantifiable attributes for segment tracking
- Accessible Segments: Ensure segments can be reached through available marketing channels
- Stable Definitions: Create segment definitions that remain valid over time
- Evidence-Based: Ground all segment characteristics in validated research findings

## Framework Integration Notes

- SDLC Integration: User validation results inform Epic prioritization and Story creation
- Business Framework Usage: Leverages Jobs-to-be-Done methodology for systematic user understanding
- Evidence Standards: Maintains quantified, multi-source validation approach
- Quality Assurance: Built-in scoring and validation ensures reliable user insights
- Professional Output: Structured documentation supports product and marketing decisions

==== END FILE ====

==== FILE: .krci-ai/tasks/create-project-brief.md ====
---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
    - business-frameworks.md
  templates:
    - project-brief-template.md
---

# Task: Create Project Brief

## Description

Create a comprehensive project brief defining the foundation for product development by answering why, who, what success looks like, and what constraints shape the solution. This document serves as the root artifact in the SDLC framework that defines the essential foundation for all downstream artifacts, answers fundamental questions before solution development begins, and provides strategic context for PRD creation.

## Instructions

<instructions>
Confirm the exact output path `/docs/prd/project-brief.md` you will create or update and ensure dependencies declared in the YAML frontmatter are accessible. Do not proceed if required inputs are missing.

Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) to align with artifact flow. Apply relevant methods from [business-frameworks.md](./.krci-ai/data/business-frameworks.md). Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) and populate all variables precisely.
</instructions>

## Output Format

- Location: `/docs/prd/project-brief.md` (EXACT path and filename)
- Length: 2-3 pages maximum for executive consumption
- Downstream Enable: Enables PRD creation at `/docs/prd/prd.md`

## Success Criteria

<success_criteria>
- File saved to `/docs/prd/project-brief.md`
- Length is 2-3 pages maximum
- Problem is specific and evidence-based
- Users are clearly defined with usage patterns
- Success metrics are specific and testable
- Constraints reflect actual limitations
- Risks identified with impact levels (HIGH/MEDIUM/LOW)
</success_criteria>

## Execution Checklist

### Discovery Phase

<discovery_phase>
- Stakeholder interviews: Understand business context and strategic priorities
- Problem validation: Gather evidence that this problem is real and significant
- User research: Identify who has this problem and how it impacts them
- Opportunity sizing: Quantify business value and market opportunity
</discovery_phase>

### Analysis Phase

<analysis_phase>
- Problem definition: Write specific problem statement with evidence
- User segmentation: Define target users with demographics and usage patterns
- Success planning: Define measurable outcomes with realistic timelines
- Constraint assessment: Identify realistic limitations and assumptions
</analysis_phase>

### Documentation Phase

<documentation_phase>
- Brief creation: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) structure
- Content validation: Ensure all required sections are completed
- Length verification: Confirm document is 2-3 pages maximum
- File placement: Save to exact location `/docs/prd/project-brief.md`
</documentation_phase>

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- Problem Focus: Use concrete user scenarios and quantified evidence, not solution-oriented statements
- User Specificity: Define target users specifically enough to guide solution design decisions
- Measurable Success: Create specific, testable outcomes with realistic timelines and evidence
- Evidence-Based: Support all statements with data, research, and quantified metrics

### LLM Error Prevention Checklist

- Avoid: Solution-oriented problem statements (focus on user pain, not missing features)
- Avoid: Vague user descriptions without usage patterns and demographics
- Avoid: Unmeasurable success metrics or aspirational statements without evidence
- Reference: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) for all formatting guidance and examples

### SDLC Integration Context

This Project Brief enables immediate PRD creation by providing clear problem definition for PRD Problem/Opportunity section, target user clarity for PRD user research and requirements, success metrics for PRD Goals/Measurable Outcomes, and constraints for PRD MVP scope and technical requirements.

==== END FILE ====

==== FILE: .krci-ai/tasks/validate-business-value.md ====
---
dependencies:
  data:
    - business-frameworks.md
    - validation-frameworks.md
  templates:
    - validation-report-template.md
---
# Task: Validate Business Value

## Description

Apply Value Proposition Canvas and ROI calculation frameworks to validate business value proposition, financial justification, and market positioning. This validation ensures the project creates meaningful customer and business value with credible return on investment.

This validation uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and outputs results using the [validation report template](./.krci-ai/templates/validation-report-template.md).

## Instructions

<instructions>
Confirm project brief with opportunity/business value section exists, access to customer research or feedback data is available, financial data for cost and benefit estimation is accessible, and competitive analysis or market positioning data is available. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Apply Value Proposition Canvas and ROI calculation frameworks from [validation-frameworks.md](./.krci-ai/data/validation-frameworks.md) to validate business value proposition, financial justification, and market positioning. Extract and structure current value proposition from project brief including customer value benefits, business value financial and strategic benefits, market value competitive advantage, and solution differentiation key differentiators from alternatives. Use [validation-report-template.md](./.krci-ai/templates/validation-report-template.md) for output.
</instructions>

## Value Hypothesis Structure

```markdown
Business Value Hypothesis: {{project_name}} creates value by delivering {{customer_benefit}} to {{target_customers}}, resulting in {{business_outcome}} for the organization through {{value_mechanism}}.

Specific Claims to Validate:
- Customer value: {{customer_value_claims}}
- Business impact: {{business_impact_claims}}
- Financial returns: {{financial_return_claims}}
- Competitive advantage: {{competitive_advantage_claims}}
```

### Phase 2: Value Proposition Canvas Framework Application

#### Customer Profile Validation

Customer Jobs Analysis:
- Functional Jobs: Practical tasks customers are trying to accomplish
- Emotional Jobs: Feelings customers want to achieve or avoid
- Social Jobs: How customers want to be perceived by others

Customer Pains Identification:
- Problem Pains: Undesired outcomes, problems, and characteristics
- Obstacle Pains: Things preventing customers from getting jobs done
- Risk Pains: Potential negative consequences customers fear

Customer Gains Discovery:
- Required Gains: Basic benefits customers expect
- Expected Gains: Benefits customers expect but would function without
- Desired Gains: Benefits customers want but aren't expecting
- Unexpected Gains: Benefits that go beyond expectations

#### Value Map Validation

Products and Services:
- List specific products/services the value proposition builds on
- Validate alignment with customer jobs, pains, and gains
- Assess completeness and competitive differentiation

Pain Relievers:
- How products/services alleviate specific customer pains
- Validate pain relief effectiveness with customer evidence
- Assess pain relief priority and impact

Gain Creators:
- How products/services create customer gains
- Validate gain creation effectiveness with customer evidence
- Assess gain importance and differentiation

#### Fit Assessment Framework

Problem-Solution Fit:

```markdown
Customer Pain: {{identified_pain}}
Pain Intensity: {{pain_score}}/10 (Evidence: {{pain_evidence}})
Solution Pain Relief: {{pain_reliever_description}}
Relief Effectiveness: {{relief_score}}/10 (Evidence: {{relief_evidence}})
Pain-Solution Fit: {{fit_score}}/10

Overall Problem-Solution Fit: {{average_fit_score}}/10
```

Product-Market Fit:

```markdown
Customer Gain: {{desired_gain}}
Gain Importance: {{gain_score}}/10 (Evidence: {{gain_evidence}})
Solution Gain Creation: {{gain_creator_description}}
Creation Effectiveness: {{creation_score}}/10 (Evidence: {{creation_evidence}})
Gain-Solution Fit: {{fit_score}}/10

Overall Product-Market Fit: {{average_fit_score}}/10
```

### Phase 3: ROI Calculation Validation Framework

#### Investment Cost Analysis

Development Costs:
- Personnel costs (developers, designers, product managers)
- Technology costs (tools, infrastructure, licenses)
- Third-party costs (contractors, consultants, services)
- Opportunity costs (alternative investments foregone)

Implementation Costs:
- Deployment and rollout expenses
- Training and change management costs
- Integration and setup costs
- Testing and quality assurance costs

Operational Costs:
- Ongoing maintenance and support costs
- Infrastructure and hosting costs
- Personnel costs for ongoing operations
- Upgrade and enhancement costs

#### Benefit Quantification

Revenue Benefits:
- New revenue from new customers or markets
- Increased revenue from existing customers
- Revenue protection from competitive threats
- Premium pricing from differentiation

Cost Savings Benefits:
- Process efficiency improvements
- Automation and labor savings
- Reduced error and rework costs
- Vendor consolidation savings

Strategic Benefits:
- Market share gains
- Customer retention improvements
- Brand value enhancement
- Competitive positioning advantages

#### ROI Calculation Methods

Simple ROI Calculation:

```markdown
Total Investment: ${{total_investment}}
Annual Benefits: ${{annual_benefits}}
Simple ROI: {{((annual_benefits - annual_costs) / total_investment * 100)}}%
Payback Period: {{total_investment / (annual_benefits - annual_costs)}} years
```

Net Present Value (NPV):

```markdown
Discount Rate: {{discount_rate}}%
Time Horizon: {{time_horizon}} years
Year 0 Investment: ${{initial_investment}}
Year 1-N Cash Flows: ${{cash_flow_by_year}}
NPV: ${{calculated_npv}}
NPV Interpretation: {{npv_positive_negative_assessment}}
```

Internal Rate of Return (IRR):

```markdown
IRR: {{calculated_irr}}%
Company Hurdle Rate: {{hurdle_rate}}%
IRR Assessment: {{irr_vs_hurdle_assessment}}
```

### Phase 4: Evidence Collection and Validation

#### Customer Evidence Collection

Primary Customer Research:
- Customer interviews validating jobs, pains, and gains
- User testing validating pain relief and gain creation
- Customer willingness-to-pay research
- Customer switching behavior analysis

Secondary Market Research:
- Industry reports on similar value propositions
- Competitive analysis of alternative solutions
- Market research on customer segments and needs
- Expert opinions on market trends and opportunities

#### Financial Evidence Collection

Internal Financial Data:
- Historical cost data for similar projects
- Revenue data for existing customer segments
- Operational cost data for baseline comparison
- Investment criteria and hurdle rates

External Benchmark Data:
- Industry benchmark data for similar initiatives
- Vendor quotes and cost estimates
- Market pricing data for competitive solutions
- Economic data affecting discount rates and projections

### Phase 5: Validation Results Analysis

#### Value Proposition Fit Scoring

Customer Profile Validation Score:

```markdown
Jobs Validation: {{jobs_score}}/10
- Evidence quality: {{jobs_evidence_quality}}
- Completeness: {{jobs_completeness}}
- Customer confirmation: {{jobs_customer_confirmation}}

Pains Validation: {{pains_score}}/10
- Evidence quality: {{pains_evidence_quality}}
- Pain intensity: {{pains_intensity_average}}
- Customer confirmation: {{pains_customer_confirmation}}

Gains Validation: {{gains_score}}/10
- Evidence quality: {{gains_evidence_quality}}
- Gain importance: {{gains_importance_average}}
- Customer confirmation: {{gains_customer_confirmation}}
```

Value Map Validation Score:

```markdown
Products/Services Fit: {{products_score}}/10
Pain Relievers Fit: {{pain_relievers_score}}/10
Gain Creators Fit: {{gain_creators_score}}/10

Overall Value Proposition Fit: {{overall_vp_fit_score}}/10
```

#### Financial Validation Assessment

ROI Validation Quality:

```markdown
Cost Estimation Quality: {{cost_quality_score}}/10
- Estimation method: {{cost_estimation_method}}
- Data sources: {{cost_data_sources}}
- Confidence level: {{cost_confidence_level}}%

Benefit Estimation Quality: {{benefit_quality_score}}/10
- Estimation method: {{benefit_estimation_method}}
- Data sources: {{benefit_data_sources}}
- Confidence level: {{benefit_confidence_level}}%

Overall ROI Confidence: {{roi_confidence_level}}%
Investment Decision Recommendation: {{investment_recommendation}}
```

## Output Format

- Primary Output: `/docs/prd/brief-validation-value.md` using validation report template
- Secondary Outputs:
  - Updated `/docs/prd/project-brief.md` with validated value proposition
  - Updated `/docs/prd/brief-assumptions.md` with value assumptions
  - `/docs/prd/value-proposition-canvas.md` with detailed analysis
  - `/docs/prd/roi-analysis.md` with financial projections
- Length: Validation report 2-3 pages, supporting documents as needed

## Success Criteria

<success_criteria>

### Value Proposition Validation

- Customer profile validated with evidence from customer research
- Value map confirmed with customer feedback on pain relief and gain creation
- Problem-solution fit assessed with quantified fit scores >7/10
- Product-market fit evaluated with evidence of customer demand

### Financial Validation

- ROI calculations completed using multiple methods (Simple, NPV, IRR)
- Cost estimates validated with vendor quotes and historical data
- Benefit projections supported with customer evidence and benchmarks
- Investment recommendation provided based on validated analysis

### Evidence Quality

- Customer evidence collected from primary research with target users
- Financial evidence validated with internal and external benchmark data
- Competitive evidence gathered for market positioning validation
- Assumption tracking updated with value proposition validation results

### Strategic Alignment

- Business case validated with quantified value and credible ROI
- Competitive differentiation confirmed with market analysis
- Strategic fit assessed with organizational goals and capabilities
- Risk assessment completed with value realization dependencies
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Validation Preparation

- Current value analysis: Review existing opportunity/business value sections
- Customer access planning: Arrange customer interviews and research access
- Financial data gathering: Collect internal cost and revenue data
- Competitive research: Analyze alternative solutions and market positioning

### Value Proposition Canvas Application

- Customer profile mapping: Define jobs, pains, and gains with customer input
- Value map creation: Detail products/services, pain relievers, gain creators
- Fit assessment: Evaluate problem-solution and product-market fit
- Evidence validation: Confirm fit assessment with customer feedback

### ROI Analysis

- Cost estimation: Calculate development, implementation, and operational costs
- Benefit quantification: Estimate revenue, cost savings, and strategic benefits
- ROI calculation: Apply multiple calculation methods for validation
- Sensitivity analysis: Test ROI under different scenarios and assumptions

### Evidence Collection and Analysis

- Customer research: Conduct interviews and surveys with target users
- Financial validation: Verify costs and benefits with internal and external data
- Competitive analysis: Assess differentiation and market positioning
- Assumption testing: Validate key value proposition assumptions

### Documentation and Communication

- Validation report creation: Complete comprehensive value validation report
- Brief updates: Update project brief with validated value proposition
- Supporting documentation: Create detailed value proposition canvas and ROI analysis
- Stakeholder presentation: Communicate validation results and investment recommendation

</execution_checklist>

## Content Guidelines

### Value Validation Quality Standards

- Customer-Centric: Ground all value claims in validated customer research
- Evidence-Based: Support all value propositions with quantified evidence
- Financially Rigorous: Use multiple ROI methods with conservative assumptions
- Competitively Aware: Position value relative to alternative solutions
- Risk Transparent: Clearly document assumptions and uncertainty factors

### Customer Research Standards

- Representative Sampling: Include diverse customer perspectives within target segments
- Unbiased Methodology: Use open-ended questions avoiding confirmation bias
- Quantified Insights: Collect measurable data on pain intensity and gain importance
- Switching Behavior: Understand customer willingness to change from current solutions
- Willingness to Pay: Validate value with economic decision-making behavior

### Financial Analysis Standards

- Conservative Assumptions: Use realistic rather than optimistic projections
- Multiple Scenarios: Include best case, worst case, and most likely scenarios
- Time Value Consideration: Apply appropriate discount rates for NPV analysis
- Risk Adjustment: Consider implementation risk and market uncertainty
- Peer Benchmarking: Compare ROI expectations with similar initiatives

## Framework Integration Notes

- SDLC Integration: Validated value proposition informs Epic prioritization and Story value statements
- Business Framework Usage: Leverages [business frameworks](./.krci-ai/data/business-frameworks.md) including Value Proposition Canvas and financial analysis methodologies
- Evidence Standards: Maintains customer-validated and financially rigorous approach
- Quality Assurance: Built-in scoring ensures credible value proposition and ROI analysis
- Professional Output: Investment-grade analysis suitable for executive decision-making

==== END FILE ====

==== FILE: .krci-ai/tasks/create-project-brief-advanced.md ====
---
dependencies:
  data:
    - validation-frameworks.md
    - business-frameworks.md
    - krci-ai/core-sdlc-framework.md
  templates:
    - project-brief-template-advanced.md
    - context-gathering-guide-template.md
    - assumption-tracker-template.md
    - validation-report-template.md
  tasks:
    - gather-project-context.md
    - validate-problem-statement.md
    - validate-target-users.md
    - validate-success-metrics.md
    - validate-business-value.md
    - refine-project-brief.md
    - finalize-project-brief.md
    - enhance-project-brief.md
---

# Task: Create Project Brief (Advanced)

## Description

Create a comprehensive project brief using advanced validation workflow with business framework methodology, evidence collection, and stakeholder validation. This advanced version extends the standard project brief creation with systematic validation, assumption tracking, and quality assurance for high-stakes projects requiring executive approval or comprehensive stakeholder buy-in.

This task leverages [validation frameworks](./.krci-ai/data/validation-frameworks.md) and the [advanced project brief template](./.krci-ai/templates/project-brief-template-advanced.md) for comprehensive project analysis.

Use Advanced Flow When:
- High-stakes projects (>$100K budget, >6 month timeline)
- Strategic initiatives requiring executive approval
- Market uncertainty or unvalidated assumptions
- Stakeholder validation and evidence-based decision making required
- Competitive or complex market environment

## Instructions

<instructions>
Confirm this advanced validation flow is justified (high-stakes project >$100K, executive approval needed, market uncertainty, or stakeholder validation required). Ensure 2-4 weeks are available for the multi-session validation process, stakeholders are accessible for interviews, and dependencies declared in the YAML frontmatter are readable before proceeding.

Execute comprehensive context gathering using business frameworks by running [gather-project-context](./.krci-ai/tasks/gather-project-context.md) task for systematic evidence collection. Conduct structured stakeholder interviews using the [context gathering guide](./.krci-ai/templates/context-gathering-guide-template.md) and build comprehensive evidence base with confidence assessments. Create detailed assumption tracking using the [assumption tracker template](./.krci-ai/templates/assumption-tracker-template.md) and output results to `/docs/prd/project-context.md` with supporting evidence and assumption documentation.

Create initial brief using advanced template with validation checkpoints including built-in validation status tracking for each section, confidence levels with evidence source documentation, systematic assumption identification and risk assessment, methodology citations with validation approach, and professional standards with executive-ready formatting. Output initial `/docs/prd/project-brief.md` with validation placeholders and checkpoint system.

Apply systematic validation using established business methodologies by executing [validate-problem-statement](./.krci-ai/tasks/validate-problem-statement.md) using Lean Startup Problem-Solution Fit, then execute [validate-target-users](./.krci-ai/tasks/validate-target-users.md) using Jobs-to-be-Done framework, followed by [validate-success-metrics](./.krci-ai/tasks/validate-success-metrics.md) using SMART criteria and OKR alignment, and complete with [validate-business-value](./.krci-ai/tasks/validate-business-value.md) using Value Proposition Canvas. Output validation reports to `/docs/prd/brief-validation-*.md` using the [validation report template](./.krci-ai/templates/validation-report-template.md) with evidence and confidence assessments.

Integrate validation results into enhanced project brief by running [refine-project-brief](./.krci-ai/tasks/refine-project-brief.md) to integrate validation findings. Update section confidence based on validation evidence quality, modify assumption tracker with validation results, and incorporate validated evidence into brief sections. Output enhanced `/docs/prd/project-brief.md` with validated evidence and reduced assumptions.

Complete quality gates and stakeholder approval for production readiness by running [finalize-project-brief](./.krci-ai/tasks/finalize-project-brief.md) for comprehensive quality review. Obtain sign-off from key decision makers, confirm all advanced flow standards are met, and verify enablement for downstream PRD development. Output production-ready project brief with comprehensive validation documentation.
</instructions>

## Output Format

<output_format>
- Primary Output: `/docs/prd/project-brief.md` (same location as standard flow)
- Supporting Outputs:
  - `/docs/prd/project-context.md` (context gathering results)
  - `/docs/prd/brief-assumptions.md` (assumption tracking)
  - `/docs/prd/brief-validation-*.md` (validation reports)
  - `/docs/prd/brief-evidence.md` (evidence library)
- Length: 2-3 pages for main brief + supporting validation documentation
- Quality: Enterprise-grade with evidence-based validation
</output_format>

## Success Criteria

<success_criteria>

### Context Foundation

- Project context gathered using business framework methodologies
- Stakeholder interviews completed with structured approach
- Evidence library created with quality and confidence assessments
- Assumption inventory developed with impact and risk prioritization

### Validation Completion

- Problem validated using Lean Startup Problem-Solution Fit methodology
- Users validated using Jobs-to-be-Done framework with evidence
- Metrics validated using SMART criteria and OKR alignment assessment
- Value validated using Value Proposition Canvas and ROI analysis

### Quality Assurance

- Brief enhanced with validation evidence and confidence levels
- Assumptions tracked with validation status and evidence sources
- Stakeholder approval obtained from key decision makers
- SDLC integration verified for downstream artifact creation

### Professional Standards

- Executive readiness confirmed with 2-3 page length and professional format
- Evidence documentation complete with source attribution and quality assessment
- Risk mitigation addressed through assumption validation and stakeholder alignment
- Investment grade analysis suitable for budget approval and resource allocation
</success_criteria>

## Execution Checklist

<execution_checklist>

### Advanced Flow Preparation

- Flow justification: Confirm project meets criteria for advanced validation
- Resource planning: Allocate 2-4 weeks for comprehensive validation process
- Stakeholder coordination: Arrange access to key stakeholders and subject matter experts
- Evidence preparation: Identify available research, data, and documentation

### Context and Validation Execution

- Context gathering: Execute systematic context collection using business frameworks
- Brief creation: Generate initial brief with advanced template and validation checkpoints
- Validation cycle: Complete business framework validation for high-priority areas
- Evidence integration: Incorporate validation findings into enhanced brief

### Quality Assurance and Finalization

- Refinement completion: Integrate all validation results and evidence
- Quality verification: Confirm advanced flow standards and professional presentation
- Stakeholder approval: Obtain sign-off from key decision makers and stakeholders
- Documentation package: Complete all supporting documentation for handoff
</execution_checklist>

## Advanced Flow vs Standard Flow

### When to Use Standard Flow (`create-project-brief`)

- Low-to-medium risk projects
- Internal projects with known stakeholders
- Tight timelines requiring rapid project initiation
- Prototypes or experimental projects
- Limited budget for extensive validation

### When to Use Advanced Flow (`create-project-brief-advanced`)

- High-stakes or strategic projects
- External stakeholder validation required
- Market uncertainty or unvalidated assumptions
- Executive approval and evidence-based decision making needed
- Budget and timeline support comprehensive validation

### Upgrade Path

Standard project briefs can be upgraded to advanced using [enhance-project-brief](./.krci-ai/tasks/enhance-project-brief.md) command when project importance or requirements change.

## Framework Integration Notes

- SDLC Integration: Advanced brief provides stronger foundation for complex PRD development
- Business Framework Usage: Leverages established methodologies for professional validation
- Evidence Standards: Maintains quantified, multi-source validation approach
- Quality Assurance: Built-in validation ensures enterprise-grade output
- Professional Output: Investment-grade documentation suitable for executive and board presentation

==== END FILE ====

==== FILE: .krci-ai/tasks/enhance-project-brief.md ====
---
dependencies:
  data:
    - validation-frameworks.md
  templates:
    - project-brief-template-advanced.md
  tasks:
    - gather-project-context.md
    - validate-problem-statement.md
    - validate-target-users.md
    - validate-success-metrics.md
    - validate-business-value.md
    - refine-project-brief.md
---

# Task: Enhance Project Brief

## Description

Upgrade an existing standard project brief to the advanced validation flow, adding business framework validation, evidence collection, and assumption tracking. This task bridges standard rapid creation with comprehensive validation when project importance or risk increases.

This task uses the [advanced project brief template](./.krci-ai/templates/project-brief-template-advanced.md) and [validation frameworks](./.krci-ai/data/validation-frameworks.md) for comprehensive enhancement.

## Instructions

<instructions>
Confirm the existing project brief at `/docs/prd/project-brief.md` is accessible and created with standard flow. Verify the project has increased in strategic importance, budget, or risk level, stakeholders are requesting evidence-based validation, and time is available for comprehensive validation process (2-4 weeks). Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Follow this systematic enhancement process to upgrade the brief to advanced validation flow with business framework validation, evidence collection, and assumption tracking:

### Phase 1: Enhancement Assessment

Evaluate whether enhancement is appropriate and beneficial:

#### Enhancement Triggers Assessment

Strategic Importance:
- Project budget increased above $100K
- Timeline extended beyond 6 months
- Executive stakeholder interest increased
- Competitive or market pressure identified

Risk Level Changes:
- Key assumptions challenged by stakeholders
- Market uncertainty or user segment questions emerged
- Technical feasibility concerns raised
- Business case questioned or ROI disputed

Stakeholder Requirements:
- Evidence-based validation requested
- Detailed business case required for approval
- Comprehensive user research needed
- Risk assessment and mitigation planning required

### Phase 2: Simple Brief Analysis

Extract and analyze existing simple brief content:

#### Content Extraction

1. Existing Problem Statement: Extract current problem definition
2. Current User Definition: Identify stated target users
3. Stated Success Metrics: Capture existing success criteria
4. Documented Constraints: List known limitations and assumptions
5. Identified Risks: Extract mentioned risks and concerns

#### Gap Analysis

```markdown
Simple Brief Analysis: {{project_name}}

Strengths of Current Brief:
- {{strength_1}}
- {{strength_2}}
- {{strength_3}}

Areas Requiring Validation:
- {{validation_need_1}} (Evidence gap: {{evidence_gap_1}})
- {{validation_need_2}} (Evidence gap: {{evidence_gap_2}})
- {{validation_need_3}} (Evidence gap: {{evidence_gap_3}})

High-Risk Assumptions Identified:
- {{assumption_1}} (Risk level: {{risk_level_1}})
- {{assumption_2}} (Risk level: {{risk_level_2}})

Enhancement Priority Areas:
1. {{priority_area_1}} - {{rationale_1}}
2. {{priority_area_2}} - {{rationale_2}}
3. {{priority_area_3}} - {{rationale_3}}
```

### Phase 3: Enhanced Template Migration

Migrate content to enhanced template with validation checkpoints:

#### Template Enhancement Process

1. Backup Simple Brief: Create backup of original simple brief
2. Load Enhanced Template: Apply enhanced template with validation checkpoints
3. Migrate Content: Transfer existing content to enhanced structure
4. Add Validation Placeholders: Insert validation checkpoints and assumption tracking
5. Mark for Enhancement: Flag brief as enhanced flow in progress

#### Enhanced Structure Addition

```markdown
Enhanced Elements Added:
- Validation checkpoints for each section
- Assumption tracking with confidence levels
- Evidence quality assessment placeholders
- Business framework methodology references
- Stakeholder validation tracking
```

### Phase 4: Prioritized Validation Planning

Develop validation plan based on highest-risk and highest-impact areas:

#### Validation Prioritization

High Priority (Immediate validation required):
- Problem validation if core assumptions challenged
- User validation if market uncertainty exists
- Value validation if ROI questioned
- Metrics validation if success criteria disputed

Medium Priority (Planned validation):
- Areas with moderate evidence gaps
- Secondary assumptions with medium impact
- Competitive positioning questions
- Technical feasibility concerns

Low Priority (Monitor and validate if needed):
- Well-supported assumptions
- Low-impact variables
- Stable market conditions
- Proven technical approaches

#### Validation Sequence Planning

```markdown
Validation Execution Plan:

Week 1-2: Context Gathering
- Execute gather-context task
- Stakeholder interviews
- Evidence library creation
- Assumption inventory completion

Week 3-4: High-Priority Validations
- {{high_priority_validation_1}}
- {{high_priority_validation_2}}
- Brief refinement with results

Week 5-6: Medium-Priority Validations (if needed)
- {{medium_priority_validation_1}}
- {{medium_priority_validation_2}}
- Final brief enhancement

Week 7: Finalization
- Quality gate verification
- Stakeholder approval
- Enhanced brief completion
```

### Phase 5: Enhancement Execution

Execute the planned enhancement with validation:

#### Context Gathering Execution

- Run [gather-project-context](./.krci-ai/tasks/gather-project-context.md) task to build evidence foundation
- Conduct stakeholder interviews for gap-filling
- Create comprehensive assumption inventory
- Establish evidence library with confidence assessments

#### Selective Validation Execution

Based on priority assessment, execute relevant validation tasks:
- [validate-problem-statement](./.krci-ai/tasks/validate-problem-statement.md) for problem statement validation
- [validate-target-users](./.krci-ai/tasks/validate-target-users.md) for user segment validation
- [validate-success-metrics](./.krci-ai/tasks/validate-success-metrics.md) for success criteria validation
- [validate-business-value](./.krci-ai/tasks/validate-business-value.md) for business case validation

#### Brief Integration and Refinement

- Run [refine-project-brief](./.krci-ai/tasks/refine-project-brief.md) to integrate validation results
- Update confidence levels based on evidence quality
- Enhance assumptions tracking with validation outcomes
- Prepare enhanced brief for finalization
</instructions>

## Output Format

- Primary Output: Enhanced `/docs/prd/project-brief.md` with validation integration
- Supporting Outputs:
  - `/docs/prd/project-context.md` (if context gathering executed)
  - `/docs/prd/brief-assumptions.md` (enhanced assumption tracking)
  - `/docs/prd/brief-validation-*.md` (validation reports as executed)
  - `/docs/prd/enhancement-summary.md` (summary of changes made)
- Length: 2-3 pages enhanced brief (expanded from 1-2 page simple)

## Success Criteria

<success_criteria>

### Enhancement Planning

- Enhancement justification documented with clear triggers and rationale
- Gap analysis completed identifying areas requiring validation
- Validation plan created with prioritized sequence and timeline
- Resource requirements assessed for validation execution

### Template Migration

- Simple brief content preserved with backup created
- Enhanced template applied with validation checkpoints added
- Assumption tracking integrated with confidence level placeholders
- Flow marker updated from SIMPLE_FLOW to ENHANCED_FLOW

### Validation Execution

- High-priority validations completed based on risk and impact assessment
- Evidence integrated with appropriate confidence levels
- Assumption status updated based on validation results
- Brief quality enhanced with validated evidence and reduced assumptions

### Finalization

- Enhanced brief completed meeting enhanced flow standards
- Stakeholder approval obtained for enhanced version
- Enhancement documentation complete for future reference
- SDLC integration verified for downstream PRD creation
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Enhancement Assessment

- Enhancement triggers evaluated: Verify justification for enhancement
- Current brief analyzed: Extract content and identify gaps
- Stakeholder input gathered: Confirm enhancement requirements
- Resource planning: Assess time and effort required for enhancement

### Enhancement Planning

- Gap analysis conducted: Identify specific areas needing validation
- Validation prioritization: Rank validations by impact and urgency
- Timeline creation: Develop realistic schedule for enhancement process
- Success criteria definition: Define what enhanced brief should achieve

### Migration Execution

- Content backup: Preserve original simple brief
- Template application: Apply enhanced template structure
- Content migration: Transfer existing content to enhanced format
- Validation setup: Add checkpoints and assumption tracking

### Validation Execution

- Context gathering: Execute if comprehensive foundation needed
- Priority validations: Complete high-priority validation tasks
- Evidence integration: Incorporate validation results into brief
- Quality assessment: Evaluate enhanced brief against standards

### Enhancement Completion

- Final integration: Complete refine-project-brief process
- Quality verification: Ensure enhanced brief meets all standards
- Stakeholder review: Obtain approval for enhanced version
- Documentation: Complete enhancement summary and lessons learned
</execution_checklist>

## Content Guidelines

### Enhancement Quality Standards

- Value-Added: Enhancement should significantly improve brief quality and confidence
- Evidence-Based: All enhancements supported by validated evidence or analysis
- Stakeholder-Focused: Enhancement addresses specific stakeholder needs or concerns
- Risk-Mitigation: Enhanced brief reduces project risk through better understanding
- Professional Grade: Enhanced brief suitable for high-stakes decision making

### Validation Selection Standards

- Risk-Based Prioritization: Focus validation effort on highest-risk assumptions
- Impact-Driven: Validate areas with highest potential impact on project success
- Evidence-Gap Closure: Target validations that close most significant evidence gaps
- Stakeholder-Requested: Include validations specifically requested by stakeholders
- Resource-Appropriate: Balance validation depth with available time and resources

### Documentation Standards

- Enhancement Rationale: Clearly document why enhancement was pursued
- Change Tracking: Document all changes from simple to enhanced brief
- Evidence Integration: Show how validation evidence improved brief quality
- Future Reference: Document process and learnings for future enhancements
- Stakeholder Communication: Prepare clear summary of enhancements for stakeholders

## Framework Integration Notes

- SDLC Integration: Enhanced brief provides stronger foundation for downstream artifacts
- Flow Flexibility: Supports both planned enhancement and reactive enhancement triggers
- Quality Assurance: Built-in assessment ensures enhancement adds genuine value
- Resource Management: Prioritized approach manages validation effort efficiently
- Professional Standards: Enhanced output meets enterprise decision-making requirements

==== END FILE ====

==== FILE: .krci-ai/tasks/validate-success-metrics.md ====
---
dependencies:
  data:
    - validation-frameworks.md
    - business-frameworks.md
  templates:
    - validation-report-template.md
---

# Task: Validate Success Metrics

## Description

Apply SMART criteria and OKR alignment frameworks to validate success metrics quality, achievability, and strategic alignment. This validation ensures metrics are specific, measurable, achievable, relevant, and time-bound while supporting organizational objectives.

This validation uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and outputs results using the [validation report template](./.krci-ai/templates/validation-report-template.md).

## Instructions

<instructions>
Confirm project brief with success metrics section exists, baseline data or historical performance data is available, organizational OKR or strategic goals are documented, and industry benchmark data is accessible. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Apply SMART criteria and OKR alignment frameworks from [validation-frameworks.md](./.krci-ai/data/validation-frameworks.md) to validate success metrics quality, achievability, and strategic alignment. Extract current success metrics from project brief including business metrics (revenue, cost savings, market share), user metrics (adoption, engagement, satisfaction), performance metrics (system performance, reliability), and operational metrics (efficiency, productivity). Use [validation-report-template.md](./.krci-ai/templates/validation-report-template.md) for output.
</instructions>

## Metrics Hypothesis Structure

```markdown
Success Metrics Hypothesis: {{project_name}} success measured by achieving:
- Business Impact: {{business_metric}} from {{baseline}} to {{target}} by {{date}}
- User Impact: {{user_metric}} from {{baseline}} to {{target}} by {{date}}
- Performance Impact: {{performance_metric}} from {{baseline}} to {{target}} by {{date}}

Specific Claims to Validate:
- Baseline measurements: {{baseline_claims}}
- Target achievability: {{target_claims}}
- Timeline feasibility: {{timeline_claims}}
- Measurement capability: {{measurement_claims}}
```

### Phase 2: SMART Criteria Validation Framework

#### SMART Assessment Process

For each success metric, evaluate against SMART criteria:

Specific Validation (1-5 scale):
- Is the metric clearly defined without ambiguity?
- Can stakeholders understand what exactly will be measured?
- Is the metric scope bounded and focused?
- Does the metric avoid vague or subjective language?

Measurable Validation (1-5 scale):
- Can the metric be quantified with specific numbers?
- Are measurement tools and data sources identified?
- Is the measurement process documented and reliable?
- Can progress be tracked throughout project lifecycle?

Achievable Validation (1-5 scale):
- Is the target realistic given available resources?
- Are there precedents for similar achievements?
- Have constraints and limitations been considered?
- Is the metric within team/organization control or influence?

Relevant Validation (1-5 scale):
- Does the metric align with business objectives?
- Will achieving this metric create meaningful value?
- Is the metric important to key stakeholders?
- Does the metric contribute to strategic goals?

Time-bound Validation (1-5 scale):
- Is there a clear deadline or timeframe?
- Are there intermediate milestones defined?
- Is the timeline realistic for metric achievement?
- Are dependencies and risks to timeline identified?

#### SMART Score Calculation

```
Total SMART Score = Specific + Measurable + Achievable + Relevant + Time-bound

Interpretation:
- 20-25: Excellent metrics, high confidence to proceed
- 15-19: Good metrics, minor refinement recommended
- 10-14: Moderate metrics, significant improvement needed
- <10: Poor metrics, fundamental revision required
```

### Phase 3: OKR Alignment Validation

#### Organizational Alignment Assessment

Objective Alignment:
- Does project objective align with company/department OKRs?
- Is the objective inspirational and directionally correct?
- Does the objective contribute to higher-level strategic goals?
- Is the objective clear enough to guide decision-making?

Key Result Quality:
- Are key results outcome-focused rather than activity-focused?
- Do key results collectively indicate objective achievement?
- Are key results ambitious yet achievable (60-70% confidence)?
- Do key results avoid measuring inputs or activities?

#### OKR Integration Analysis

```markdown
Company/Department OKR: {{higher_level_okr}}
Project Objective: {{project_objective}}
Alignment Score: {{alignment_score}}/10

Key Results Analysis:
KR1: {{key_result_1}}
- Outcome Focus: {{outcome_score_1}}/5
- Measurability: {{measurability_score_1}}/5
- Ambition Level: {{ambition_score_1}}/5
- Achievement Confidence: {{confidence_score_1}}%

KR2: {{key_result_2}}
- Outcome Focus: {{outcome_score_2}}/5
- Measurability: {{measurability_score_2}}/5
- Ambition Level: {{ambition_score_2}}/5
- Achievement Confidence: {{confidence_score_2}}%

Overall OKR Quality Score: {{okr_total_score}}/50
```

### Phase 4: Leading/Lagging Indicator Analysis

#### Indicator Balance Assessment

Leading Indicators (Predictive metrics):
- Identify metrics that predict future outcomes
- Validate correlation with historical data
- Assess controllability and actionability
- Define measurement frequency and reporting

Lagging Indicators (Outcome metrics):
- Identify metrics that measure final results
- Validate as true measures of success
- Assess measurement accuracy and reliability
- Define measurement and reporting process

#### Balanced Scorecard Creation

```markdown
Leading Indicators (Early Warning System):
1. {{leading_indicator_1}}: {{description}} (Frequency: {{frequency}})
2. {{leading_indicator_2}}: {{description}} (Frequency: {{frequency}})

Lagging Indicators (Success Measurement):
1. {{lagging_indicator_1}}: {{description}} (Frequency: {{frequency}})
2. {{lagging_indicator_2}}: {{description}} (Frequency: {{frequency}})

Predictive Relationships:
- {{leading_indicator}} correlates with {{lagging_indicator}} (R² = {{correlation}})
- Time lag: {{time_lag_period}}
- Confidence in relationship: {{relationship_confidence}}%
```

### Phase 5: Evidence Collection and Validation

#### Baseline Data Validation

Data Collection Requirements:
- Historical performance data for achievability assessment
- Industry benchmark data for target calibration
- Measurement system reliability validation
- Data source accessibility and cost assessment

#### Benchmark Analysis

Industry Benchmarking:
- Identify relevant industry performance benchmarks
- Compare proposed targets with industry standards
- Assess competitive positioning implications
- Validate market expectation alignment

#### Measurement Feasibility Assessment

Technical Validation:
- Verify measurement tools and systems capability
- Assess data collection cost and complexity
- Identify measurement gaps and solutions
- Validate measurement accuracy and reliability

## Output Format

<output_format>
- Primary Output: `/docs/prd/brief-validation-metrics.md` using validation report template
- Secondary Outputs:
  - Updated `/docs/prd/project-brief.md` with validated metrics
  - Updated `/docs/prd/brief-assumptions.md` with metrics assumptions
  - `/docs/prd/metrics-baseline-data.md` with supporting data
- Length: Validation report 2-3 pages, supporting documents as needed
</output_format>

## Success Criteria

<success_criteria>

### SMART Criteria Validation

- All metrics assessed using SMART criteria with scores documented
- SMART scores calculated with improvement recommendations for scores <15
- Baseline data identified for all measurable metrics
- Measurement systems validated for feasibility and reliability

### OKR Alignment Validation

- Strategic alignment confirmed with organizational/departmental OKRs
- Key result quality assessed using outcome-focus criteria
- Ambition level calibrated for 60-70% achievement confidence
- OKR hierarchy documented from company to project level

### Leading/Lagging Balance

- Indicator balance achieved with both leading and lagging metrics
- Predictive relationships validated with historical correlation data
- Measurement frequency defined for each indicator type
- Early warning system established through leading indicators

### Evidence and Benchmarking

- Baseline data collected and quality assessed
- Industry benchmarks identified and targets calibrated
- Measurement feasibility confirmed with cost and complexity assessment
- Success thresholds defined with clear criteria
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Validation Preparation

- Current metrics review: Analyze existing success metrics in project brief
- Data source identification: Identify available baseline and benchmark data
- Stakeholder coordination: Connect with teams responsible for measurement
- OKR documentation: Gather organizational and departmental OKR information

### SMART Criteria Assessment

- Individual metric evaluation: Apply SMART criteria to each metric
- Scoring documentation: Record scores and rationale for each criterion
- Gap identification: Identify metrics failing SMART criteria
- Improvement recommendations: Suggest specific metric enhancements

### OKR Alignment Validation

- Alignment mapping: Document connection to higher-level OKRs
- Key result analysis: Evaluate each key result for quality
- Ambition calibration: Assess and adjust target ambition levels
- Confidence assessment: Rate achievement probability for each metric

### Evidence Collection

- Baseline data gathering: Collect historical performance data
- Benchmark research: Identify relevant industry performance standards
- Measurement validation: Verify measurement system capabilities
- Cost assessment: Evaluate measurement cost and complexity

### Documentation and Reporting

- Validation report creation: Complete comprehensive validation report
- Brief updates: Update project brief with validated metrics
- Assumption tracking: Update assumption tracker with metrics validation results
- Stakeholder communication: Present validation results to project team
</execution_checklist>

## Content Guidelines

### Metrics Quality Standards

- Quantified Focus: Prioritize metrics with specific numerical targets
- Evidence-Based Targets: Ground all targets in baseline data and benchmarks
- Balanced Portfolio: Include both leading (predictive) and lagging (outcome) indicators
- Stakeholder Relevance: Ensure metrics matter to key decision makers
- Measurement Feasibility: Verify realistic measurement capability and cost

### Validation Methodology Standards

- Systematic Assessment: Apply SMART criteria consistently across all metrics
- Evidence Requirements: Support all assessments with documented evidence
- Benchmark Integration: Use industry data to calibrate targets appropriately
- Stakeholder Validation: Confirm metrics relevance with key stakeholders
- Continuous Monitoring: Establish ongoing measurement and review processes

### Documentation Standards

- Clear Scoring: Document all SMART and OKR scores with explicit rationale
- Evidence Citation: Properly attribute all baseline data and benchmarks
- Improvement Recommendations: Provide specific guidance for metric enhancement
- Measurement Planning: Document how each metric will be tracked and reported
- Success Criteria: Define clear thresholds for success and failure

## Framework Integration Notes

- SDLC Integration: Validated metrics inform Epic success criteria and Story acceptance criteria
- Business Framework Usage: Leverages SMART methodology and OKR framework for systematic validation
- Evidence Standards: Maintains quantified, benchmark-driven approach for credible targets
- Quality Assurance: Built-in scoring ensures metrics support effective project management
- Professional Output: Structured documentation supports stakeholder decision-making and tracking

==== END FILE ====

==== FILE: .krci-ai/tasks/validate-problem-statement.md ====
---
dependencies:
  data:
    - validation-frameworks.md
    - business-frameworks.md
  templates:
    - validation-report-template.md
---

# Task: Validate Problem Statement

## Description

Apply Lean Startup Problem-Solution Fit Assessment framework to validate problem statement accuracy, evidence quality, and strategic alignment. This validation ensures the problem is real, urgent, and worth solving while identifying gaps that could impact project success.

This validation uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and outputs results using the [validation report template](./.krci-ai/templates/validation-report-template.md).

## Instructions

<instructions>
Confirm project brief with problem statement section exists, basic problem research or stakeholder input is available, access to potential users or customer data is accessible, and competitive landscape understanding is available. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Apply Lean Startup Problem-Solution Fit Assessment framework from [validation-frameworks.md](./.krci-ai/data/validation-frameworks.md) to validate problem statement accuracy, evidence quality, and strategic alignment. Extract and structure the problem hypothesis from the current project brief including core problem statement, problem scope (included/excluded), affected user segments, and problem context (when, where, why). Use [validation-report-template.md](./.krci-ai/templates/validation-report-template.md) for output.
</instructions>

## Problem Hypothesis Structure

```markdown
Problem Hypothesis: {{user_segment}} experiences {{problem_description}} when {{situation_context}}, resulting in {{negative_impact}}.

Specific Claims to Validate:
- Problem occurs for {{user_segment}} with frequency of {{frequency_claim}}
- Problem causes {{impact_type}} impact measured by {{impact_metric}}
- Current solutions fail because {{solution_gap_reason}}
- Problem urgency driven by {{urgency_factors}}
```

### Phase 2: Lean Startup Validation Framework Application

#### Step 1: Problem Validation Evidence Collection

Customer Interview Validation:
- Conduct minimum 10 target user interviews
- Use structured interview guide focusing on problem experiences
- Document specific problem scenarios and pain points
- Quantify problem frequency and impact per user

Quantified Problem Metrics:
- Collect measurable indicators of problem existence
- Support ticket data, time spent on workarounds, cost of current solutions
- User behavior analytics showing problem manifestation
- Business impact data (revenue loss, efficiency reduction)

Competitive Analysis:
- Document existing solutions and their limitations
- Identify solution gaps that validate problem persistence
- Assess market demand through competitive pricing and adoption
- Analyze competitor positioning to validate problem importance

#### Step 2: Problem-Solution Fit Assessment

Problem Validation Scoring (1-10 scale):
- Problem Intensity: How painful is the problem for users?
- Problem Frequency: How often do users experience this problem?
- Problem Reach: How many users/companies experience this problem?
- Problem Urgency: How urgent is solving this problem for users?

Evidence Quality Assessment:
- Primary evidence percentage (direct user research vs secondary sources)
- Evidence recency (how recent is the supporting data)
- Evidence diversity (multiple source types and perspectives)
- Evidence reliability (credible sources and methodologies)

#### Step 3: Solution Hypothesis Validation

Solution Approach Assessment:
- Does the proposed solution address identified root causes?
- Is the solution technically and commercially feasible?
- Does the solution create meaningful differentiation vs alternatives?
- Can the solution be delivered within stated constraints?

Solution Desirability Testing:
- User reaction to proposed solution approach
- Willingness to pay or switch from current solutions
- Feature prioritization based on problem-solving value
- Solution-market fit assessment

### Phase 3: Evidence Analysis and Validation

#### Evidence Quality Framework

Primary Evidence Sources (Highest confidence):
- Direct customer interviews with problem validation
- First-party usage data showing problem manifestation
- Customer support data quantifying problem frequency
- User research studies with problem-focused findings

Secondary Evidence Sources (Moderate confidence):
- Industry reports mentioning related problems
- Competitive analysis revealing solution gaps
- Expert opinions on problem significance
- Market research highlighting related pain points

Evidence Confidence Scoring:
- High (80-100%): Multiple primary sources, quantified data, recent research
- Medium (60-79%): Mix of primary/secondary, some quantification, reasonably recent
- Low (40-59%): Primarily secondary sources, limited quantification, dated research
- Very Low (<40%): Assumptions without validation, anecdotal evidence only

#### Validation Results Synthesis

Problem-Solution Fit Score Calculation:

```txt
Problem Score = (Intensity + Frequency + Reach + Urgency) / 4
Solution Score = (Root Cause Fit + Feasibility + Differentiation + Deliverability) / 4
Overall Fit = (Problem Score + Solution Score) / 2
```

Fit Assessment Thresholds:
- Strong Fit (8-10): High confidence to proceed with current problem/solution
- Moderate Fit (6-7.9): Some refinement needed, additional validation recommended
- Weak Fit (4-5.9): Significant issues identified, major refinement required
- Poor Fit (<4): Problem or solution hypothesis needs fundamental revision

### Phase 4: Validation Results Documentation

#### Create Validation Report

Use validation report template to document:
- Framework methodology and process followed
- Evidence collected with quality assessment
- Problem-solution fit scores and rationale
- Key findings and insights
- Assumptions validated, challenged, or identified
- Recommendations for project brief updates

#### Update Project Brief

Based on validation results:
- Enhance problem statement with validated evidence
- Update confidence levels and assumption tracking
- Revise problem scope or focus based on findings
- Add validation checkpoint completion status

#### Update Assumption Tracker

Document assumption changes:
- Mark validated assumptions with evidence sources
- Update confidence levels based on validation results
- Add new assumptions identified during validation
- Prioritize remaining assumptions for future validation

## Output Format

<output_format>
- Primary Output: `/docs/prd/brief-validation-problem.md` using validation report template
- Secondary Outputs:
  - Updated `/docs/prd/project-brief.md` with validation results
  - Updated `/docs/prd/brief-assumptions.md` with assumption status
- Length: Validation report 2-3 pages, brief updates minimal
- Evidence Documentation: All sources properly cited and quality-assessed
</output_format>

## Success Criteria

<success_criteria>

### Phase 1: Problem Hypothesis

- Problem hypothesis extracted from current project brief
- Specific validation claims identified with measurable criteria
- Problem scope clearly defined with inclusion/exclusion boundaries
- Target user segments specified for validation activities

### Phase 2: Framework Application

- Customer interviews completed with minimum 10 target users
- Quantified metrics collected showing problem existence and impact
- Competitive analysis conducted validating solution gaps
- Problem-solution fit scores calculated using framework methodology

### Phase 3: Evidence Analysis

- Evidence quality assessed with confidence levels assigned
- Primary evidence prioritized over secondary sources
- Multiple source validation completed for key claims
- Evidence gaps identified with recommendations for closure

### Phase 4: Documentation

- Validation report completed using standardized template
- Project brief updated with validated evidence and confidence levels
- Assumption tracker updated with validation results
- Recommendations provided for next validation steps
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Validation Setup

- Current brief analysis: Review existing problem statement and evidence
- Validation planning: Identify specific claims to test and evidence needed
- Stakeholder coordination: Arrange access to customers, data, and subject matter experts
- Interview guide preparation: Develop structured questions for customer interviews

### Problem Validation Execution

- Customer interviews: Conduct minimum 10 interviews with structured approach
- Data collection: Gather quantified evidence from internal and external sources
- Competitive research: Analyze existing solutions and market positioning
- Expert consultation: Validate findings with internal subject matter experts

### Analysis and Synthesis

- Evidence analysis: Assess quality and reliability of collected evidence
- Scoring calculation: Apply framework scoring methodology consistently
- Pattern identification: Identify themes and insights across evidence sources
- Gap identification: Document areas needing additional validation

### Documentation and Communication

- Report creation: Complete validation report using standardized template
- Brief updates: Update project brief with validation results and evidence
- Assumption updates: Revise assumption tracker based on validation findings
- Stakeholder communication: Present findings to key project stakeholders

</execution_checklist>

## Content Guidelines

### Validation Quality Standards

- Evidence-Based: All conclusions supported by documented evidence with source attribution
- Quantified Focus: Prioritize measurable evidence over qualitative opinions
- Multiple Source Validation: Key findings validated across multiple evidence sources
- Bias Mitigation: Actively seek disconfirming evidence and diverse perspectives
- Transparency: Clearly document limitations, assumptions, and confidence levels

### Interview Best Practices

- Open-ended Questions: Avoid leading questions that bias toward expected answers
- Specific Scenarios: Ask for concrete examples rather than general opinions
- Quantification: Request specific numbers (frequency, time, cost) when possible
- Problem Focus: Focus on problems and pain points before discussing solutions
- Context Gathering: Understand when, where, and why problems occur

### Evidence Documentation Standards

- Source Attribution: Clearly identify all evidence sources with dates and methodology
- Quality Assessment: Rate evidence quality and reliability consistently
- Confidence Levels: Assign and document confidence levels for all key findings
- Assumptions Tracking: Explicitly identify and track all assumptions made during analysis
- Update Requirements: Document what evidence updates would change conclusions

## Framework Integration Notes

- SDLC Integration: Validation results feed into Epic creation and solution architecture
- Business Framework Usage: Leverages Lean Startup methodology from [validation frameworks](./.krci-ai/data/validation-frameworks.md) for systematic validation
- Evidence Standards: Maintains KubeRocketAI focus on quantified, evidence-based approach
- Quality Assurance: Built-in scoring and confidence assessment ensure reliable results
- Professional Output: Structured documentation suitable for stakeholder decision-making

==== END FILE ====

==== FILE: .krci-ai/tasks/gather-project-context.md ====
---
dependencies:
  data:
    - validation-frameworks.md
    - business-frameworks.md
    - krci-ai/core-sdlc-framework.md
  templates:
    - context-gathering-guide-template.md
---

# Task: Gather Project Context

## Description

Systematically collect and structure project context using business frameworks and established methodologies. This task provides multiple input methods to gather comprehensive information needed for evidence-based project brief creation, ensuring all assumptions are tracked and validated with appropriate evidence.

This task leverages [validation frameworks](./.krci-ai/data/validation-frameworks.md) and the [context gathering guide template](./.krci-ai/templates/context-gathering-guide-template.md) for systematic context collection.

## Instructions

<instructions>
Confirm business opportunity or problem area is identified, stakeholder access is available for interviews and input, initial project scope or direction is defined, and access to existing research, documentation, or data sources is available. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Systematically collect and structure project context using business frameworks and established methodologies. This provides comprehensive information needed for evidence-based project brief creation.

### Phase 1: Input Method Selection

Choose appropriate input collection methods based on available resources and project context:

#### Method 1: Document Analysis

When to use: Existing research, competitive analysis, user studies, or PRDs available
Process:
1. Inventory all available documents and research
2. Categorize documents by type (user research, market data, technical specs, competitive analysis)
3. Extract key insights using structured analysis framework
4. Identify evidence gaps and assumption areas

Document Types to Analyze:
- Market research reports
- Competitive analysis
- User research and interviews
- Existing PRDs or product specs
- Technical feasibility studies
- Stakeholder feedback
- Customer support data
- Analytics and usage data

#### Method 2: Stakeholder Interview Process

When to use: Key stakeholders available for structured input gathering
Process:
1. Use BABOK elicitation techniques for structured interviews
2. Apply Design Thinking empathy mapping methods
3. Document interviews using standardized templates
4. Validate findings across multiple stakeholders

Stakeholder Categories:
- Business stakeholders (product owners, executives)
- User representatives (customer success, support)
- Technical stakeholders (architects, lead developers)
- Market experts (sales, marketing, customer success)

#### Method 3: Assumption Inventory Creation

When to use: Limited existing research, need to identify knowledge gaps
Process:
1. Brainstorm all assumptions about problem, users, solution, market
2. Categorize assumptions by type and confidence level
3. Prioritize assumptions by impact and uncertainty
4. Create validation plan for high-priority assumptions

Assumption Categories:
- Problem assumptions (what, who, why, when)
- User assumptions (segments, needs, behaviors, preferences)
- Solution assumptions (feasibility, desirability, viability)
- Market assumptions (size, competition, trends)
- Business assumptions (model, metrics, resources)

#### Method 4: Evidence Library Creation

When to use: Quantified data available, need systematic evidence collection
Process:
1. Identify available quantified data sources
2. Structure data using evidence classification framework
3. Assess data quality and confidence levels
4. Identify additional data needs

Evidence Types:
- Usage analytics and user behavior data
- Support ticket analysis and customer feedback
- Market sizing and industry benchmark data
- Technical performance and feasibility data
- Financial data (costs, revenue, pricing)

### Phase 2: Business Framework Application

Apply relevant [business frameworks](./.krci-ai/data/business-frameworks.md) to structure collected information:

#### Problem Context Analysis

Frameworks to Apply:
- Six Sigma 5 Whys: For root cause analysis of identified problems
- Impact-Frequency Matrix: For problem prioritization and urgency assessment
- SIPOC Analysis: For process understanding and stakeholder mapping

Output: Structured problem context with root causes, stakeholders, and priority assessment

#### User Context Analysis

Frameworks to Apply:
- Jobs-to-be-Done: For user motivation and outcome identification
- Design Thinking Empathy Maps: For user perspective and experience mapping
- User Journey Mapping: For touchpoint and pain point identification

Output: Comprehensive user context with jobs, pains, gains, and journey insights

#### Market Context Analysis

Frameworks to Apply:
- Business Model Canvas: For value proposition and business model analysis
- Porter's Five Forces: For competitive environment assessment
- TAM/SAM/SOM Analysis: For market sizing and opportunity quantification

Output: Market context with competitive landscape, opportunity size, and positioning

#### Business Context Analysis

Frameworks to Apply:
- SWOT Analysis: For internal capabilities and external factors
- OKR Alignment: For strategic alignment and goal hierarchy
- Resource Assessment: For capability and constraint analysis

Output: Business context with strategic alignment, resources, and constraints

### Phase 3: Context Synthesis and Validation

Synthesize collected information into structured project context:

#### Context Synthesis Process

1. Integrate Multi-Source Findings: Combine insights from all input methods
2. Identify Patterns and Themes: Look for consistent insights across sources
3. Highlight Contradictions: Note conflicting information for further validation
4. Assess Evidence Quality: Rate confidence levels for each insight

#### Validation Checkpoints

1. Stakeholder Validation: Review synthesized context with key stakeholders
2. Evidence Validation: Verify quantified claims with additional sources
3. Assumption Validation: Test high-impact assumptions with appropriate methods
4. Internal Consistency: Ensure context elements align logically

### Phase 4: Context Documentation

Document project context using structured templates:

#### Context Summary Structure

```markdown
# Project Context Summary: {{project_name}}

## Context Collection Methodology
- Input methods used: {{methods_list}}
- Business frameworks applied: {{frameworks_list}}
- Evidence sources: {{sources_list}}
- Confidence assessment: {{overall_confidence_level}}

## Problem Context
- Root problem identification using {{validation_method}}
- Problem scope and boundaries
- Stakeholder impact assessment
- Priority and urgency analysis

## User Context
- Target user segments with validation evidence
- User jobs, pains, and gains analysis
- User journey and touchpoint mapping
- User needs prioritization

## Market Context
- Market opportunity sizing and validation
- Competitive landscape analysis
- Value proposition positioning
- Market trend and timing analysis

## Business Context
- Strategic alignment and goal hierarchy
- Resource and capability assessment
- Constraint and limitation identification
- Success criteria and measurement approach

## Evidence Library
- Primary evidence sources and confidence levels
- Secondary research and benchmark data
- Quantified metrics and baseline measurements
- Evidence gaps and validation needs

## Assumption Inventory
- High-impact assumptions requiring validation
- Medium-impact assumptions for monitoring
- Low-impact assumptions acceptable as-is
- Assumption validation plan and timeline
```

</instructions>

## Output Format

<output_format>

- Location: `/docs/prd/project-context.md` (EXACT path and filename)
- Supporting Files:
  - `/docs/prd/evidence-library.md`
  - `/docs/prd/assumption-inventory.md`
- Length: 3-4 pages maximum for executive consumption
- Downstream Enablement: Enables enhanced project brief creation
</output_format>

## Success Criteria

<success_criteria>

### Phase 1: Input Collection

- Input methods selected based on available resources and project needs
- Document analysis completed with structured insight extraction
- Stakeholder interviews conducted using business framework methods
- Assumption inventory created with prioritization and confidence levels

### Phase 2: Framework Application

- Problem context analyzed using appropriate business frameworks
- User context mapped with jobs, pains, gains, and journey insights
- Market context assessed with competitive and opportunity analysis
- Business context documented with strategic alignment and constraints

### Phase 3: Context Validation

- Multi-source integration completed with pattern identification
- Stakeholder validation conducted with key project stakeholders
- Evidence quality assessed with confidence levels assigned
- Assumption validation plan created for high-impact assumptions

### Phase 4: Documentation

- Context summary created using structured template
- Evidence library documented with sources and confidence levels
- Assumption inventory finalized with validation priorities
- Files saved to exact locations `/docs/prd/project-context.md`
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Context Gathering

- Stakeholder identification: Map all relevant stakeholders and their roles
- Resource assessment: Identify available documents, data, and research
- Timeline planning: Set realistic timeline for context gathering activities
- Framework selection: Choose appropriate business frameworks for analysis

### Context Collection Phase

- Document inventory: Catalog all available research and documentation
- Stakeholder interviews: Conduct structured interviews using BABOK methods
- Data collection: Gather quantified evidence from available sources
- Assumption brainstorming: Identify assumptions across all project dimensions

### Analysis Phase

- Framework application: Apply selected business frameworks to collected data
- Pattern recognition: Identify themes and insights across sources
- Gap identification: Identify evidence gaps and validation needs
- Confidence assessment: Rate evidence quality and assumption confidence

### Validation Phase

- Stakeholder review: Validate synthesized context with key stakeholders
- Evidence verification: Cross-check quantified claims with additional sources
- Assumption testing: Plan validation for high-impact assumptions
- Internal consistency: Ensure logical alignment across context elements

### Documentation Phase

- Context summary: Create comprehensive context documentation
- Evidence library: Document all sources with confidence assessments
- Assumption inventory: List assumptions with validation priorities
- Quality review: Verify completeness and accuracy of documentation
</execution_checklist>

## Content Guidelines

### Quality Principles for Context Gathering

- Evidence-Based: Prioritize quantified evidence over opinions or assumptions
- Multi-Source Validation: Validate key insights across multiple sources
- Assumption Transparency: Clearly identify and track all assumptions
- Confidence Assessment: Rate evidence quality and assumption confidence levels
- Business Framework Grounding: Use established methodologies for analysis

### Context Collection Best Practices

- Structure Over Volume: Focus on structured insights rather than comprehensive data
- Validation Over Confirmation: Seek to test assumptions rather than confirm biases
- Quantification Over Description: Prioritize measurable insights over qualitative observations
- Stakeholder Diversity: Include diverse perspectives to avoid single viewpoint bias
- Evidence Traceability: Maintain clear links from insights back to source evidence

### Documentation Standards

- Executive Summary Focus: Keep context summary concise and decision-focused
- Evidence Transparency: Clearly document sources and confidence levels
- Assumption Accountability: Make assumptions explicit and track validation status
- Action Orientation: Structure context to enable project brief creation
- Professional Format: Use consistent formatting and structure for stakeholder consumption

## Integration with Project Brief Creation

### Context-to-Brief Mapping

Project Context → Project Brief Sections:
- Problem Context → Problem Statement and Opportunity sections
- User Context → Target Users section with validated segments
- Market Context → Opportunity section with market sizing and positioning
- Business Context → Success Metrics, Constraints, and Risk sections
- Evidence Library → Supporting evidence for all brief assertions
- Assumption Inventory → Transparent assumption tracking throughout brief

### Enhanced Brief Quality Through Context

- Evidence-Based Assertions: All brief statements backed by context evidence
- Validated Assumptions: Brief assumptions tested through context gathering
- Stakeholder Alignment: Brief content pre-validated with stakeholders during context phase
- Risk Mitigation: Context gaps identified and addressed before brief creation
- Strategic Alignment: Brief grounded in business context and strategic goals

## Framework Integration Notes

- SDLC Integration: Project context serves as foundation for all downstream [SDLC framework](./.krci-ai/data/krci-ai/core-sdlc-framework.md) artifacts
- Business Framework Usage: Leverages established methodologies for professional analysis
- Evidence Standards: Maintains KubeRocketAI focus on quantified, evidence-based approach
- Quality Assurance: Built-in validation checkpoints ensure context accuracy and completeness
- Professional Output: Structured documentation suitable for executive and stakeholder consumption

==== END FILE ====

# Shared Templates

==== FILE: .krci-ai/templates/project-brief-template-advanced.md ====
# Project Brief: {{project_name}} (Advanced Validation)

> **Target Length**: 2-3 pages maximum
> **Framework**: Advanced validation flow with business framework validation
> **File Location**: MUST be saved as `/docs/prd/project-brief.md` (exact path)
> **Flow Type**: ADVANCED_VALIDATION | **Evidence Level**: {{evidence_confidence}} | **Assumptions**: {{tracked_assumptions}}

---

## Executive Summary

<executive_summary>
{{executive_summary}}

Validation Checkpoint: Problem-Solution Fit Assessment
- Problem validated using {{validation_method}}
- Solution approach validated with {{business_framework}}
- Business value quantified with {{confidence_level}} confidence

Key Assumptions:
- {{assumption_1}} (Confidence: {{confidence_1}}, Evidence: {{evidence_source_1}})
- {{assumption_2}} (Confidence: {{confidence_2}}, Evidence: {{evidence_source_2}})

Validation Methods Applied:
1. Lean Startup Problem-Solution Fit Analysis
2. Business Model Canvas Value Proposition Validation
3. Market Opportunity Assessment with TAM/SAM/SOM

<instructions>
Provide a compelling 3-4 sentence overview combining validated problem, solution approach, and quantified expected outcome.

Example: Our SaaS platform experiences 2,500 password-related support tickets monthly (validated through 6-month support data analysis), consuming 15% of support resources and causing user frustration (confirmed through 25 user interviews). We will implement biometric authentication and social login options to reduce password dependency, targeting 80% reduction in support tickets and $50K annual savings (validated through similar industry implementations). This 3-month initiative serves 10,000+ monthly active users and requires Auth0 integration with a $25K budget (confirmed through vendor quotes and technical assessment).

Advanced Elements:
- All numbers backed by documented evidence sources
- Solution approach validated through business frameworks
- Expected outcomes confirmed through benchmark analysis
- Scope validated through stakeholder interviews and constraint analysis
</instructions>
</executive_summary>

---

## Problem Statement

<problem_statement>
{{problem_statement}}

Validation Checkpoint: Root Cause Analysis
- Problem validated using {{problem_validation_method}}
- Root causes identified with {{root_cause_method}}
- Problem scope boundaries confirmed through {{scope_validation}}
- Impact quantification verified with {{impact_evidence}}

Problem Evidence:
- Primary Evidence: {{primary_evidence}} (Confidence: {{evidence_confidence_1}})
- Supporting Data: {{supporting_data}} (Confidence: {{evidence_confidence_2}})
- Stakeholder Validation: {{stakeholder_input}} (Confidence: {{evidence_confidence_3}})

Validation Methods Available:
1. Six Sigma 5 Whys Root Cause Analysis
2. Impact-Frequency Problem Assessment
3. Stakeholder Problem Validation Interviews

<instructions>
Define the specific pain point driving this project with validated scope boundaries and quantified evidence.

Example: Development teams spend 3-4 hours daily on manual code reviews (validated through time-tracking analysis of 50 developers across 5 teams over 3 months), leading to 15% of total development time lost and delayed feature releases (confirmed through project delivery data analysis). This problem affects 85% of development teams according to industry survey data and costs organizations $150K annually per 20-developer team (calculation validated through comparable industry benchmarks). Focus on code review workflow optimization only, excluding code quality policy management or development process changes (scope confirmed through stakeholder interviews).

Advanced Evidence Requirements:
- Quantified problem metrics with data source attribution
- User research findings with sample sizes and methodologies
- Industry benchmark comparisons with credible sources
- Stakeholder interview validation with diverse role representation
- Historical data analysis showing problem persistence and trends
</instructions>
</problem_statement>

---

## Opportunity

<opportunity>
{{opportunity}}

Validation Checkpoint: Value Proposition Analysis
- Business value validated using {{value_validation_method}}
- Market opportunity sized with {{market_analysis_method}}
- Competitive advantage confirmed through {{competitive_analysis}}
- ROI calculations validated with {{roi_validation_method}}

Value Evidence:
- Financial Analysis: {{financial_evidence}} (Method: {{financial_method}}, Confidence: {{financial_confidence}})
- Market Research: {{market_evidence}} (Source: {{market_source}}, Confidence: {{market_confidence}})
- Competitive Analysis: {{competitive_evidence}} (Scope: {{competitive_scope}}, Confidence: {{competitive_confidence}})

Validation Methods Available:
1. Value Proposition Canvas Fit Assessment
2. Business Model Canvas Value Validation
3. ROI Calculation with Multi-Method Verification

<instructions>
Provide quantified business value plus validated solution approach with evidence-based projections.

Example: Reducing manual code review time by 70% would save $105K annually per 20-developer team (calculation validated through industry salary benchmarks and productivity studies), improve feature delivery speed by 40% (confirmed through similar tool implementations at comparable organizations), and increase developer satisfaction by 25% (based on peer organization case studies). Implement AI-powered code review assistant integrating with existing GitHub workflows to provide automated security analysis, code quality recommendations, and reviewer prioritization (solution approach validated through technical feasibility assessment and vendor demonstrations).

Advanced Value Requirements:
- Business value quantified with multiple calculation methods
- Cost savings validated through industry benchmarks and internal data
- Revenue impact projections based on comparable implementations
- Competitive advantage analysis with market positioning assessment
- Solution feasibility confirmed through technical and vendor validation
</instructions>
</opportunity>

---

## Target Users

<target_users>
{{target_users}}

Validation Checkpoint: User Segment Analysis
- User segments validated using {{user_validation_method}}
- Jobs-to-be-Done analysis completed with {{jobs_evidence}}
- User journey mapping verified through {{journey_validation}}
- Market sizing confirmed with {{sizing_evidence}}

User Evidence:
- Primary Research: {{user_research}} (Sample: {{sample_size}}, Method: {{research_method}}, Confidence: {{user_confidence}})
- Usage Analytics: {{analytics_data}} (Period: {{data_period}}, Sample: {{analytics_sample}}, Confidence: {{analytics_confidence}})
- Market Data: {{market_research}} (Source: {{market_source}}, Date: {{market_date}}, Confidence: {{market_confidence}})

Validation Methods Available:
1. Jobs-to-be-Done User Validation with Importance/Satisfaction Scoring
2. Design Thinking User Journey Mapping with Pain Point Quantification
3. Persona Demographic Validation with Market Sizing

<instructions>
Provide specific user segments validated through research with usage patterns, demographics, and evidence-based sizing.

Example: Software development teams (validated market size: 450,000 teams globally based on Stack Overflow Developer Survey and GitHub Enterprise data) with 20-500 developers who conduct code reviews 3+ times daily (confirmed through usage pattern analysis of 100 development teams). Primary segment: senior developers and team leads aged 28-45 who spend 25-30% of their time on code review activities (validated through time-tracking study of 200 developers across 15 organizations). User jobs-to-be-done analysis reveals core job: When reviewing code, I want to quickly identify security issues and quality problems so I can maintain high standards while not becoming a development bottleneck (importance: 9.2/10, current satisfaction: 4.1/10, opportunity score: 14.3/15 based on survey of 150 senior developers).

Advanced User Requirements:
- User volume quantified with credible market research sources
- Demographics validated through multiple research sources
- Usage patterns confirmed through analytics and user observation
- Segment prioritization based on revenue potential and solution fit
- Jobs-to-be-Done analysis with importance and satisfaction scoring
- User journey validation with pain point impact quantification
</instructions>
</target_users>

---

## Success Metrics

<success_metrics>
{{success_metrics}}

Validation Checkpoint: Metrics Quality Assessment
- SMART criteria validation completed for all metrics
- OKR alignment confirmed with {{okr_alignment_evidence}}
- Leading/lagging indicator balance verified
- Measurement feasibility validated with {{measurement_validation}}

Metrics Evidence:
- Baseline Data: {{baseline_metrics}} (Source: {{baseline_source}}, Period: {{baseline_period}}, Confidence: {{baseline_confidence}})
- Industry Benchmarks: {{benchmark_data}} (Source: {{benchmark_source}}, Date: {{benchmark_date}}, Confidence: {{benchmark_confidence}})
- Historical Performance: {{historical_data}} (Period: {{historical_period}}, Sample: {{historical_sample}}, Confidence: {{historical_confidence}})

Validation Methods Available:
1. SMART Criteria Assessment with Scoring Matrix
2. OKR Alignment Validation with Strategic Goal Mapping
3. Leading/Lagging Indicator Analysis with Predictive Relationship Assessment

<instructions>
Provide validated success metrics with SMART criteria assessment and evidence-based targets.

Example: Reduce code review time by 70% within 6 months (Specific: manual review time reduction, Measurable: current average 3.2 hours/day to <1 hour/day based on time-tracking analysis, Achievable: confirmed through similar AI tool implementations showing 60-80% reduction, Relevant: aligns with developer productivity OKR and engineering efficiency goals, Time-bound: 6-month implementation with monthly measurement). Improve security vulnerability detection by 60% measured through monthly security audit results (baseline: 23 vulnerabilities/month missed in manual reviews based on 6-month retrospective analysis, target: <10 vulnerabilities/month based on AI-assisted tool benchmarks). Achieve 80% developer adoption within 3 months measured through GitHub integration usage analytics (baseline: 0%, industry benchmark: 70-85% adoption for development productivity tools based on adoption studies).

Advanced Metrics Requirements:
- All metrics assessed using SMART criteria with documented scoring
- Baseline data collected and validated with multiple sources
- Targets calibrated using industry benchmarks and comparable implementations
- Measurement systems identified and validated for feasibility
- Leading indicators defined to predict lagging outcome achievement
- OKR alignment confirmed with organizational strategic goals
</instructions>
</success_metrics>

---

## Constraints

<constraints>
{{constraints}}

Validation Checkpoint: Constraint Analysis
- Resource constraints validated with {{resource_validation}}
- Technical constraints confirmed through {{technical_validation}}
- Business constraints verified with {{business_validation}}
- Timeline feasibility assessed using {{timeline_validation}}

Constraint Evidence:
- Resource Analysis: {{resource_evidence}} (Method: {{resource_method}}, Confidence: {{resource_confidence}})
- Technical Assessment: {{technical_evidence}} (Source: {{technical_source}}, Confidence: {{technical_confidence}})
- Business Validation: {{business_evidence}} (Stakeholder: {{business_stakeholder}}, Confidence: {{business_confidence}})

<instructions>
Resource, technical, and business factors that limit the solution, validated with evidence and stakeholder confirmation.

Example: Must integrate with existing Auth0 setup (technical constraint confirmed through architecture review and vendor confirmation), 6-month timeline driven by regulatory compliance deadline (business constraint validated through legal team consultation), $150K budget allocation confirmed through financial planning process (resource constraint validated through budget approval documentation), maximum 3 developers assigned based on team capacity analysis (resource constraint confirmed through engineering management review). Assumes current mobile app architecture supports biometric APIs based on technical assessment (assumption validated through proof-of-concept implementation) and 75% of users have biometric-capable devices based on user analytics analysis (assumption validated through device usage data from past 12 months).

Advanced Constraint Requirements:
- Resource constraints validated through capacity planning and budget analysis
- Technical constraints confirmed through architecture review and feasibility assessment
- Business constraints verified with stakeholder interviews and compliance review
- Timeline constraints validated through project estimation and dependency analysis
- Assumptions documented with evidence sources and confidence levels
</instructions>
</constraints>

---

## Key Risks

<key_risks>
{{key_risks}}

Validation Checkpoint: Risk Assessment Analysis
- Risk impact assessment completed using {{risk_assessment_method}}
- Risk probability analysis validated with {{probability_evidence}}
- Mitigation strategies identified and validated
- Risk monitoring plan established with {{monitoring_approach}}

Risk Evidence:
- Historical Risk Data: {{risk_history}} (Source: {{risk_source}}, Period: {{risk_period}}, Confidence: {{risk_confidence}})
- Industry Risk Analysis: {{industry_risks}} (Source: {{industry_source}}, Scope: {{industry_scope}}, Confidence: {{industry_confidence}})
- Expert Assessment: {{expert_input}} (Expert: {{expert_source}}, Method: {{expert_method}}, Confidence: {{expert_confidence}})

Risk Assessment Methods Available:
1. Risk Impact-Probability Matrix with Quantified Scoring
2. Industry Benchmark Risk Analysis with Peer Comparison
3. Expert Panel Risk Assessment with Structured Evaluation

<instructions>
Major risks validated through evidence and expert assessment with quantified impact and probability estimates.

Example: User adoption resistance (HIGH impact, MEDIUM probability): Senior developers may resist AI-assisted tools based on past automation concerns. Evidence: Survey of 100 senior developers shows 35% initial resistance to AI coding tools, but 78% adoption after 3-month trial period in comparable organizations. Mitigation: Implement gradual rollout with developer champion program and productivity showcase (strategy validated through change management consultant review). Integration complexity (MEDIUM impact, HIGH probability): GitHub Enterprise API integration may require custom authentication handling based on initial technical assessment. Evidence: Similar integrations required 2-4 weeks additional development time in 60% of comparable implementations according to vendor case studies. Mitigation: Allocate 4-week buffer for integration testing and custom development (timeline validated through technical lead estimation).

Advanced Risk Requirements:
- Risk impact quantified with specific metrics and evidence sources
- Risk probability based on historical data and expert assessment
- Mitigation strategies validated through case studies and expert consultation
- Risk monitoring plan with specific triggers and measurement approaches
- Evidence sources documented with quality assessment and confidence levels
</instructions>
</key_risks>

---

## Advanced Validation Summary

<validation_summary>

### Validation Framework Applied

- Primary Frameworks: {{primary_frameworks_used}}
- Validation Confidence: {{overall_validation_confidence}}%
- Evidence Quality Score: {{evidence_quality_score}}/100
- Assumption Risk Level: {{assumption_risk_level}}

### Business Framework Integration

- Problem Validation: {{problem_framework}} (Score: {{problem_score}}/10)
- User Validation: {{user_framework}} (Score: {{user_score}}/10)
- Metrics Validation: {{metrics_framework}} (Score: {{metrics_score}}/10)
- Value Validation: {{value_framework}} (Score: {{value_score}}/10)
</validation_summary>

<supporting_documentation>

### Supporting Documentation

- Project Context: `/docs/prd/project-context.md`
- Assumption Tracker: `/docs/prd/brief-assumptions.md`
- Evidence Library: `/docs/prd/brief-evidence.md`
- Validation Reports: `/docs/prd/brief-validation-*.md`
</supporting_documentation>

---

## SDLC Framework Information

<sdlc_framework>
Dependencies: Project context gathering and validation completion
Output Location: This Project Brief MUST be saved as `/docs/prd/project-brief.md`
Downstream Enablement: Enables enhanced PRD creation with validated foundation
Flow Type: ADVANCED_VALIDATION for evidence-based decision making

<instructions>
This Advanced Project Brief provides comprehensive foundation for:
- PRD Problem/Opportunity section (validated with evidence)
- PRD Target Users & Use Cases (research-backed segments)
- PRD Goals/Measurable Outcomes (SMART-validated metrics)
- PRD scope and constraint definition (stakeholder-validated)

Directory Structure (Advanced Flow):
/docs/
├── prd/                          # Product vision & requirements (enhanced)
│   ├── project-context.md        # Context gathering results
│   ├── project-brief.md          # Project vision & strategy (THIS FILE - ADVANCED)
│   ├── brief-assumptions.md      # Assumption tracking and validation status
│   ├── brief-evidence.md         # Evidence library with confidence levels
│   ├── brief-validation-*.md     # Business framework validation reports
│   └── prd.md                    # Product requirements (ENHANCED BY VALIDATION)
├── epics/                        # High-level features
├── stories/                      # User stories
├── architecture/                 # System design
└── tests/                        # Quality validation
</instructions>
</sdlc_framework>

---

## Upgrade from Standard Flow

<upgrade_summary>
Migration Process: This advanced brief was {{migration_status}}:
- Original Brief: {{original_brief_status}}
- Enhancement Trigger: {{enhancement_reason}}
- Validation Added: {{validation_components_added}}
- Confidence Improvement: {{confidence_change}}

Advanced Value: {{advanced_value_summary}}
</upgrade_summary>

==== END FILE ====

==== FILE: .krci-ai/templates/project-brief-template.md ====
# Project Brief: {{project_name}}

> **Target Length**: 2-3 pages maximum
> **Framework**: Root artifact in SDLC framework
> **File Location**: MUST be saved as `/docs/prd/project-brief.md` (exact path)

---

## Executive Summary

<executive_summary>
{{executive_summary}}

<instructions>
Write a compelling 3-4 sentence overview combining problem, solution approach, and expected outcome.

Example: Our SaaS platform experiences 2,500 password-related support tickets monthly, consuming 15% of support resources and frustrating users. We will implement biometric authentication and social login options to reduce password dependency, targeting 80% reduction in support tickets and $50K annual savings. This 3-month initiative serves 10,000+ monthly active users and requires Auth0 integration with a $25K budget.

Key Elements:
- What problem are we solving? (specific and quantified)
- How will we solve it? (high-level approach)
- What's the expected outcome? (business value)
- What's the scope? (timeline, users, constraints)
</instructions>
</executive_summary>

---

## Problem Statement

<problem_statement>
{{problem_statement}}

<instructions>
Define the specific pain point driving this project with clear scope boundaries.

Example: Users frequently forget passwords leading to 15% of support tickets and account lockouts (2,500/month). Focus on authentication workflow only, excluding password policy management or user registration processes.

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
</instructions>
</problem_statement>

---

## Opportunity

<opportunity>
{{opportunity}}

<instructions>
Quantified business value plus high-level solution approach.

Example: Reducing password-related support tickets by 80% would save $50K annually and improve user satisfaction scores by 25%. Implement biometric authentication and social login options to reduce password dependency.

Key Elements:
- Business value (cost savings, revenue, efficiency)
- User value (time savings, satisfaction, productivity)
- Market opportunity (competitive advantage, growth)
- High-level solution direction (not detailed implementation)

Quantification Examples:
- Cost reduction: $50K annual savings
- Time savings: 15 minutes per user per month
- Satisfaction: 25% improvement in user satisfaction
- Efficiency: 80% reduction in support tickets
</instructions>
</opportunity>

---

## Target Users

<target_users>
{{target_users}}

<instructions>
Specific user segments who have this problem with usage patterns and demographics.

Example: SaaS platform users (10,000+ monthly active users) who access the platform 3+ times per week. Primary segment: business professionals aged 25-45 accessing from mobile devices (60%) and desktop (40%).

Include:
- User volume and growth trends
- Demographics (age, role, industry)
- Usage patterns (frequency, device, context)
- Segment prioritization (primary vs secondary)
- Geographic distribution if relevant

User Segment Examples:
- 10,000+ monthly active users
- Business professionals aged 25-45
- Mobile-first users (60% mobile, 40% desktop)
- Access platform 3+ times weekly
- Located primarily in North America and Europe
</instructions>
</target_users>

---

## Success Metrics

<success_metrics>
{{success_metrics}}

<instructions>
How we'll measure if we've solved the problem with specific timelines.

Example: Reduce password-related support tickets by 80% within 3 months, maintain 99.9% uptime, achieve 70% user adoption of new auth methods within 6 months, improve login success rate from 85% to 95%.

Success Criteria Format:
- Specific: Exactly what will be measured
- Measurable: Numbers, percentages, timelines
- Achievable: Realistic given constraints
- Relevant: Directly tied to problem and opportunity
- Time-bound: Clear deadlines

Metric Categories:
- Problem Resolution: 80% reduction in support tickets
- User Adoption: 70% user adoption within 6 months
- Quality: 99.9% uptime maintained
- User Experience: Login success rate 85% → 95%
- Business Impact: $50K annual cost savings
</instructions>
</success_metrics>

---

## Constraints

<constraints>
{{constraints}}

<instructions>
Resource, technical, and assumption factors that limit the solution.

Example: Must integrate with existing Auth0 setup, 3-month timeline, $25K budget, maximum 2 developers assigned. Assumes current mobile app architecture supports biometric APIs and users have compatible devices.

Constraint Categories:

Resource Constraints:
- Budget: $25K maximum budget
- Timeline: 3-month delivery deadline
- Team: Maximum 2 developers available
- Skills: No iOS development expertise on team

Technical Constraints:
- Integration: Must integrate with existing Auth0
- Architecture: Cannot modify core database schema
- Performance: Must maintain current response times
- Security: Must meet enterprise security standards

Business Constraints:
- Compliance: Must maintain SOC 2 compliance
- User Impact: Zero downtime deployment required
- Support: Cannot increase support complexity
- Branding: Must align with current UI/UX standards

Key Assumptions:
- Users have biometric-capable devices
- Auth0 API will remain stable
- No major iOS/Android changes during development
</instructions>
</constraints>

---

## Key Risks

<key_risks>
{{key_risks}}

<instructions>
Major risks that could derail the project with impact assessment.

Example: User adoption resistance (HIGH): Users may prefer familiar passwords. Auth0 API changes (MEDIUM): Potential breaking changes during integration. Biometric compatibility (MEDIUM): Older devices may not support all features. Timeline risk (HIGH): Integration complexity may exceed estimates.

Risk Assessment Format:
[Risk Name] ([Impact Level]): [Description and potential impact]

Impact Levels:
- HIGH: Could significantly delay or derail project
- MEDIUM: Could cause delays or require scope changes
- LOW: Minor impact, manageable workarounds available

Risk Categories:

User Adoption Risks:
- User resistance to change (HIGH)
- Learning curve for new features (MEDIUM)
- Device compatibility issues (MEDIUM)

Technical Risks:
- Integration complexity (HIGH)
- Third-party API changes (MEDIUM)
- Performance impact (LOW)

Business Risks:
- Timeline overrun (HIGH)
- Budget overrun (MEDIUM)
- Resource unavailability (MEDIUM)

Market Risks:
- Competitive response (LOW)
- Regulatory changes (MEDIUM)
- Technology shifts (LOW)
</instructions>
</key_risks>

---

## SDLC Framework Information

**Dependencies**: None (root artifact)
**Output Location**: This Project Brief MUST be saved as `/docs/prd/project-brief.md`
**Downstream Enablement**: Enables PRD creation at `/docs/prd/prd.md`

<instructions>
SDLC Framework Integration:
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
</instructions>

---

<instructions>
QUALITY CHECKLIST:
- Document is 2-3 pages maximum
- Executive summary captures complete project essence
- Problem statement is specific and evidence-based
- Opportunity is quantified with business value
- Target users are specific with usage patterns
- Success metrics are measurable with timelines
- Constraints are realistic and comprehensive
- Key risks identified with impact assessment
- File saved exactly as /docs/prd/project-brief.md
- Ready to enable PRD creation
</instructions>
==== END FILE ====

==== FILE: .krci-ai/templates/validation-report-template.md ====
# Validation Report: {{project_name}}

> File Location: `/docs/prd/brief-validation.md`
> Validation Date: {{validation_date}}
> Report Type: {{report_type}} | Framework Used: {{validation_framework}}
> Overall Confidence: {{overall_confidence}}% | Status: {{validation_status}}

---

## Validation Summary

<validation_summary>

### Validation Scope

- Section Validated: {{section_name}}
- Business Framework: {{business_framework}}
- Validation Method: {{validation_method}}
- Evidence Sources: {{evidence_sources}}
- Stakeholders Involved: {{stakeholders_involved}}

### Key Findings

{{#each key_findings}}
- {{finding_type}}: {{finding_description}}
{{/each}}

### Overall Assessment

- Pre-Validation Confidence: {{pre_validation_confidence}}%
- Post-Validation Confidence: {{post_validation_confidence}}%
- Confidence Change: {{confidence_change}} ({{confidence_direction}})
- Validation Status: {{final_status}}

<instructions>
This section provides a concise overview of the validation process and results.

Validation Scope: Document exactly what was validated and how
Key Findings: List the most important discoveries from the validation
Overall Assessment: Show before/after confidence levels and final status

CRITICAL: Keep summary to 3-4 bullet points maximum for executive readability.
</instructions>
</validation_summary>

---

## Detailed Validation Results

<detailed_results>

### {{validation_framework}} Framework Application

#### Framework Methodology

Framework: {{framework_name}}
Source: {{framework_source}}
Application Context: {{application_context}}

#### Validation Process Followed

{{#each validation_steps}}
{{step_number}}. {{step_name}}
- Method: {{step_method}}
- Evidence Collected: {{step_evidence}}
- Findings: {{step_findings}}
- Confidence Level: {{step_confidence}}%
{{/each}}

<instructions>
Document the specific business framework used and how it was applied systematically.

Framework Application: Show which methodology was used and why it was chosen
Process Steps: Document each step taken with evidence and confidence levels

CRITICAL: Include confidence percentages for each step to track validation quality.
</instructions>
</detailed_results>

#### Framework-Specific Results

{{#if is_problem_validation}}

##### Lean Startup Problem-Solution Fit Results

- Problem Hypothesis: {{problem_hypothesis}}
- Problem Validation Score: {{problem_score}}/10
- Solution Hypothesis: {{solution_hypothesis}}
- Solution Validation Score: {{solution_score}}/10
- Overall Problem-Solution Fit: {{problem_solution_fit}}

Evidence Summary:
- Customer Interview Data: {{customer_interviews}} interviews ({{interview_confidence}}% confidence)
- Quantified Problem Metrics: {{problem_metrics}} ({{metrics_confidence}}% confidence)
- Competitive Analysis: {{competitive_analysis}} ({{competitive_confidence}}% confidence)
- Technical Feasibility: {{technical_feasibility}} ({{technical_confidence}}% confidence)
{{/if}}

{{#if is_user_validation}}

##### Jobs-to-be-Done Validation Results

- Primary Job Statement: {{job_statement}}
- Job Importance Score: {{job_importance}}/10
- Current Satisfaction Score: {{job_satisfaction}}/10
- Opportunity Score: {{opportunity_score}}

Job Analysis:
- Functional Job: {{functional_job}}
- Emotional Job: {{emotional_job}}
- Social Job: {{social_job}}

Evidence Summary:
- User Interview Data: {{user_interviews}} interviews ({{user_confidence}}% confidence)
- Usage Analytics: {{usage_analytics}} (Period: {{analytics_period}}, {{analytics_confidence}}% confidence)
- Market Research: {{market_research}} (Source: {{market_source}}, {{market_confidence}}% confidence)
{{/if}}

{{#if is_metrics_validation}}

##### SMART Criteria Assessment Results

{{#each smart_metrics}}
Metric: {{metric_name}}
- Specific: {{specific_score}}/5 - {{specific_assessment}}
- Measurable: {{measurable_score}}/5 - {{measurable_assessment}}
- Achievable: {{achievable_score}}/5 - {{achievable_assessment}}
- Relevant: {{relevant_score}}/5 - {{relevant_assessment}}
- Time-bound: {{timebound_score}}/5 - {{timebound_assessment}}
- Overall SMART Score: {{smart_total_score}}/25

Evidence Summary:
- Baseline Data: {{baseline_data}} ({{baseline_confidence}}% confidence)
- Industry Benchmarks: {{industry_benchmarks}} ({{benchmark_confidence}}% confidence)
- Historical Performance: {{historical_data}} ({{historical_confidence}}% confidence)
{{/each}}
{{/if}}

{{#if is_value_validation}}

##### Value Proposition Canvas Results

Customer Profile Validation:
- Customer Jobs: {{customer_jobs}} (Validation: {{jobs_validation}}%)
- Customer Pains: {{customer_pains}} (Validation: {{pains_validation}}%)
- Customer Gains: {{customer_gains}} (Validation: {{gains_validation}}%)

Value Map Validation:
- Products & Services: {{products_services}}
- Pain Relievers: {{pain_relievers}} (Fit: {{pain_fit}}%)
- Gain Creators: {{gain_creators}} (Fit: {{gain_fit}}%)

Fit Assessment:
- Problem-Solution Fit: {{problem_solution_fit}}%
- Product-Market Fit: {{product_market_fit}}%
- Overall Value Proposition Fit: {{value_prop_fit}}%

Evidence Summary:
- Customer Research: {{customer_research}} ({{research_confidence}}% confidence)
- Solution Testing: {{solution_testing}} ({{testing_confidence}}% confidence)
- Willingness to Pay: {{willingness_to_pay}} ({{wtp_confidence}}% confidence)
{{/if}}

---

## Evidence Analysis

<evidence_analysis>

### Primary Evidence Sources

{{#each primary_evidence}}

#### {{evidence_type}}: {{evidence_name}}

- Source: {{evidence_source}}
- Collection Method: {{collection_method}}
- Sample Size: {{sample_size}}
- Time Period: {{time_period}}
- Reliability: {{reliability_score}}/10
- Relevance: {{relevance_score}}/10
- Recency: {{recency_score}}/10
- Overall Quality Score: {{quality_score}}/10

Key Insights:
{{#each insights}}
- {{insight_description}}
{{/each}}

Supporting Quotes/Data:
{{#each supporting_data}}
- {{data_point}}
{{/each}}
{{/each}}

<instructions>
Document all evidence sources with quality assessments and key findings.

Primary Evidence: Direct research, customer interviews, first-party data
Quality Scoring: Rate each source on reliability, relevance, and recency (1-10)
Key Insights: Extract the most important findings from each evidence source

CRITICAL: Primary evidence should be weighted more heavily than secondary sources.
</instructions>
</evidence_analysis>

### Secondary Evidence Sources

{{#each secondary_evidence}}

#### {{evidence_type}}: {{evidence_name}}

- Source: {{evidence_source}}
- Publication Date: {{publication_date}}
- Relevance: {{relevance_score}}/10
- Credibility: {{credibility_score}}/10
- Overall Quality Score: {{quality_score}}/10

Key Insights:
{{#each insights}}
- {{insight_description}}
{{/each}}
{{/each}}

### Evidence Gaps Identified

{{#each evidence_gaps}}
- Gap: {{gap_description}}
- Impact on Validation: {{gap_impact}}
- Recommended Action: {{gap_recommendation}}
- Priority: {{gap_priority}}
{{/each}}

---

## Assumption Updates

<assumption_updates>

### Assumptions Validated

{{#each validated_assumptions}}

#### {{assumption_id}}: {{assumption_statement}}

- Pre-Validation Confidence: {{pre_confidence}}%
- Post-Validation Confidence: {{post_confidence}}%
- Status Change: {{old_status}} → {{new_status}}
- Supporting Evidence: {{validation_evidence}}
{{/each}}

### Assumptions Challenged/Modified

{{#each challenged_assumptions}}

#### {{assumption_id}}: {{original_assumption}}

- Issue Identified: {{challenge_description}}
- Revised Assumption: {{revised_assumption}}
- Confidence Change: {{confidence_change}}
- Impact on Project Brief: {{brief_impact}}
{{/each}}

<instructions>
Track changes to project assumptions based on validation findings.

Validated Assumptions: Show confidence level improvements with supporting evidence
Challenged Assumptions: Document what was wrong and how assumptions were revised
Status Changes: Use clear before/after format to show assumption evolution

CRITICAL: All assumption changes must link back to specific evidence sources.
</instructions>
</assumption_updates>

### New Assumptions Identified

{{#each new_assumptions}}

#### {{assumption_id}}: {{assumption_statement}}

- Source: {{assumption_source}}
- Initial Confidence: {{initial_confidence}}%
- Impact Level: {{impact_level}}
- Validation Priority: {{validation_priority}}
- Recommended Validation Method: {{recommended_method}}
{{/each}}

---

## Validation Quality Assessment

### Methodology Quality

- Framework Appropriateness: {{methodology_appropriateness}}/10
- Process Adherence: {{process_adherence}}/10
- Evidence Rigor: {{evidence_rigor}}/10
- Stakeholder Involvement: {{stakeholder_involvement}}/10
- Overall Methodology Score: {{methodology_total}}/40

### Evidence Quality

- Source Diversity: {{source_diversity}}/10
- Sample Size Adequacy: {{sample_adequacy}}/10
- Data Recency: {{data_recency}}/10
- Evidence Relevance: {{evidence_relevance}}/10
- Overall Evidence Score: {{evidence_total}}/40

### Results Reliability

- Internal Consistency: {{internal_consistency}}/10
- Cross-Source Validation: {{cross_validation}}/10
- Stakeholder Agreement: {{stakeholder_agreement}}/10
- Bias Mitigation: {{bias_mitigation}}/10
- Overall Reliability Score: {{reliability_total}}/40

Combined Validation Quality Score: {{total_quality_score}}/120

---

## Impact on Project Brief

### Required Updates

{{#each brief_updates}}

#### {{section_name}} Section

- Current Content: {{current_content}}
- Recommended Change: {{recommended_change}}
- Rationale: {{change_rationale}}
- Priority: {{update_priority}}
{{/each}}

### Confidence Level Changes

{{#each confidence_changes}}
- Section: {{section_name}}
- Previous Confidence: {{previous_confidence}}%
- Updated Confidence: {{updated_confidence}}%
- Change Reason: {{change_reason}}
{{/each}}

### New Risks/Considerations Identified

{{#each new_risks}}
- Risk: {{risk_description}}
- Impact: {{risk_impact}}
- Probability: {{risk_probability}}
- Mitigation: {{risk_mitigation}}
- Brief Section Affected: {{affected_section}}
{{/each}}

---

## Recommendations and Next Steps

<recommendations>
### Immediate Actions Required
{{#each immediate_actions}}
- {{action_item}}
  - Rationale: {{action_rationale}}
  - Owner: {{action_owner}}
  - Due Date: {{due_date}}
  - Impact if Not Done: {{inaction_impact}}
{{/each}}

### Additional Validation Recommended

{{#each additional_validations}}
- Area: {{validation_area}}
- Recommended Framework: {{recommended_framework}}
- Priority: {{validation_priority}}
- Timeline: {{validation_timeline}}
- Expected Outcome: {{expected_outcome}}
{{/each}}

<instructions>
Provide clear, actionable recommendations based on validation findings.

Immediate Actions: List specific tasks that must be completed with owners and dates
Additional Validation: Identify areas needing further validation with recommended methods
Priority Levels: Use High/Medium/Low to help stakeholders focus efforts

CRITICAL: Every recommendation must have a clear rationale tied to validation findings.
</instructions>
</recommendations>

### Project Brief Refinement Suggestions

{{#each refinement_suggestions}}
- Section: {{section_name}}
- Suggestion: {{suggestion_description}}
- Benefit: {{suggestion_benefit}}
- Effort Required: {{effort_level}}
{{/each}}

---

## Validation Timeline and Planning

### Completed Validation Activities

{{#each completed_activities}}
- {{activity_date}}: {{activity_description}}
  - Outcome: {{activity_outcome}}
  - Time Invested: {{time_invested}}
  - Quality Score: {{activity_quality}}/10
{{/each}}

### Planned Follow-up Validations

{{#each planned_validations}}
- Activity: {{planned_activity}}
- Scheduled Date: {{scheduled_date}}
- Expected Duration: {{expected_duration}}
- Responsible Party: {{responsible_party}}
- Success Criteria: {{success_criteria}}
{{/each}}

### Validation Backlog

{{#each validation_backlog}}
- Priority {{priority_level}}: {{backlog_item}}
  - Estimated Effort: {{estimated_effort}}
  - Expected Value: {{expected_value}}
  - Dependencies: {{dependencies}}
{{/each}}

---

## Stakeholder Feedback

### Validation Process Feedback

{{#each process_feedback}}
{{stakeholder_name}} ({{stakeholder_role}}):
> "{{feedback_quote}}"

- Process Rating: {{process_rating}}/10
- Value Rating: {{value_rating}}/10
- Suggestions: {{stakeholder_suggestions}}
{{/each}}

### Results Feedback

{{#each results_feedback}}
{{stakeholder_name}} ({{stakeholder_role}}):
- Agreement with Findings: {{agreement_level}}/10
- Confidence in Results: {{confidence_level}}/10
- Key Concerns: {{stakeholder_concerns}}
- Additional Insights: {{additional_insights}}
{{/each}}

---

## Lessons Learned

### What Worked Well

{{#each what_worked}}
- {{success_area}}: {{success_description}}
{{/each}}

### What Could Be Improved

{{#each improvements}}
- {{improvement_area}}: {{improvement_description}}
  - Suggested Solution: {{suggested_solution}}
{{/each}}

### Recommendations for Future Validations

{{#each future_recommendations}}
- {{recommendation_category}}: {{recommendation_description}}
{{/each}}

---

## Appendices

### A. Detailed Evidence Library

{{#each detailed_evidence}}

#### {{evidence_id}}: {{evidence_title}}

Source: {{source_details}}
Collection Date: {{collection_date}}
Methodology: {{collection_methodology}}

Full Evidence:
{{evidence_content}}

Analysis Notes:
{{analysis_notes}}
{{/each}}

### B. Stakeholder Interview Transcripts

{{#each interview_transcripts}}

#### Interview: {{interviewee_name}} ({{interviewee_role}})

Date: {{interview_date}}
Duration: {{interview_duration}}
Interviewer: {{interviewer_name}}

Key Quotes:
{{#each key_quotes}}
- "{{quote_text}}" (Context: {{quote_context}})
{{/each}}

Full Transcript: [Link to {{transcript_file}}]
{{/each}}

### C. Validation Framework Resources

- Framework Documentation: {{framework_documentation_link}}
- Methodology References: {{methodology_references}}
- Best Practices: {{best_practices_link}}
- Training Materials: {{training_materials_link}}

---

<!-- REPORT METADATA
Report ID: {{report_id}}
Generated By: {{report_generator}}
Framework Version: {{framework_version}}
Template Version: {{template_version}}
Last Modified: {{last_modified}}
Next Review Date: {{next_review_date}}
-->

## Template Usage Guidelines

### When to Use This Template

- After completing any business framework validation
- When validation results significantly impact project brief confidence
- For stakeholder communication about validation outcomes
- As input for project brief refinement decisions

### Report Types

- **INITIAL**: First validation of a project brief section
- **FOLLOW_UP**: Additional validation to address gaps or concerns
- **COMPREHENSIVE**: Complete validation using multiple frameworks
- **TARGETED**: Focused validation of specific assumptions or evidence gaps

### Quality Scoring

- **1-3**: Poor quality, unreliable for decision making
- **4-6**: Moderate quality, useable with caveats
- **7-8**: Good quality, reliable for most decisions
- **9-10**: Excellent quality, high confidence for all decisions

### Evidence Classification

- **PRIMARY**: Direct research, first-party data, customer feedback
- **SECONDARY**: Industry reports, competitive analysis, expert opinions
- **QUANTIFIED**: Numerical data, measurements, statistical analysis
- **QUALITATIVE**: Descriptive data, observations, subjective assessments

==== END FILE ====

==== FILE: .krci-ai/templates/assumption-tracker-template.md ====
# Assumption Tracker: {{project_name}}

> **File Location**: `/docs/prd/brief-assumptions.md`
> **Last Updated**: {{last_updated_date}}
> **Total Assumptions**: {{total_assumptions}} | **Validated**: {{validated_count}} | **Pending**: {{pending_count}}

---

## Assumption Summary Dashboard

<summary_dashboard>

| Category | Total | High Risk | Medium Risk | Low Risk | Validated | Pending |
|----------|-------|-----------|-------------|----------|-----------|----------|
| Problem | {{problem_total}} | {{problem_high}} | {{problem_medium}} | {{problem_low}} | {{problem_validated}} | {{problem_pending}} |
| Users | {{users_total}} | {{users_high}} | {{users_medium}} | {{users_low}} | {{users_validated}} | {{users_pending}} |
| Solution | {{solution_total}} | {{solution_high}} | {{solution_medium}} | {{solution_low}} | {{solution_validated}} | {{solution_pending}} |
| Market | {{market_total}} | {{market_high}} | {{market_medium}} | {{market_low}} | {{market_validated}} | {{market_pending}} |
| Business | {{business_total}} | {{business_high}} | {{business_medium}} | {{business_low}} | {{business_validated}} | {{business_pending}} |

<instructions>
Summary dashboard provides high-level metrics across all assumption categories. Track total assumptions, risk distribution, and validation progress for each category.

Categories: Problem, Users, Solution, Market, Business
Risk Levels: High (impact HIGH + confidence <70%), Medium, Low
Status: Validated (confirmed with evidence), Pending (requires validation)
</instructions>
</summary_dashboard>

---

## Problem Assumptions

### High Impact Assumptions

<high_impact_assumptions>
{{#each problem_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

<instructions>
High impact assumptions require immediate attention and detailed validation plans. These assumptions significantly affect project success if proven false.

Required fields for high impact: rationale, evidence, validation method, timeline, owner, risk assessment, validation plan, success criteria.

CRITICAL: High impact assumptions must have confidence levels tracked and validation deadlines enforced.
</instructions>
</high_impact_assumptions>

---
{{/each}}

### Medium Impact Assumptions

<medium_impact_assumptions>
{{#each problem_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

<instructions>
Medium impact assumptions cause delays or scope changes but project remains viable if they fail. Focus on basic validation without extensive risk planning.

Required fields: evidence description, validation method, owner assignment.
</instructions>
</medium_impact_assumptions>

---
{{/each}}

### Low Impact Assumptions

<low_impact_assumptions>
{{#each problem_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
Low impact assumptions have minimal effect on project outcomes. Monitor only, validate when convenient.

Required fields: basic evidence description, confidence level tracking.
</instructions>
</low_impact_assumptions>

---
{{/each}}

## User Assumptions

<user_assumptions>

### High Impact Assumptions

{{#each user_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

### Medium Impact Assumptions

{{#each user_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

### Low Impact Assumptions

{{#each user_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
User assumptions focus on target user behavior, needs, and characteristics. Critical for product-market fit validation.

Validation Methods: user interviews, surveys, behavioral analytics, usability testing, persona validation.

CRITICAL: User assumptions should be validated through direct user contact whenever possible.
</instructions>
</user_assumptions>

## Solution Assumptions

<solution_assumptions>

### High Impact Assumptions

{{#each solution_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

### Medium Impact Assumptions

{{#each solution_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

### Low Impact Assumptions

{{#each solution_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
Solution assumptions cover technical feasibility, implementation approach, and solution effectiveness.

Validation Methods: prototyping, technical spikes, proof of concepts, architecture review, performance testing.

CRITICAL: Solution assumptions should be validated through hands-on technical validation before full implementation.
</instructions>
</solution_assumptions>

## Market Assumptions

<market_assumptions>

### High Impact Assumptions

{{#each market_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

### Medium Impact Assumptions

{{#each market_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

### Low Impact Assumptions

{{#each market_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
Market assumptions cover market size, competition, timing, and market conditions that affect product success.

Validation Methods: market research, competitive analysis, industry reports, customer development, market testing.

CRITICAL: Market assumptions should be validated through external market data and competitive intelligence.
</instructions>
</market_assumptions>

## Business Assumptions

<business_assumptions>

### High Impact Assumptions

{{#each business_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

### Medium Impact Assumptions

{{#each business_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

### Low Impact Assumptions

{{#each business_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
Business assumptions cover revenue model, business viability, resource requirements, and strategic alignment.

Validation Methods: financial modeling, stakeholder interviews, business case analysis, ROI projections, pilot programs.

CRITICAL: Business assumptions should be validated through quantitative financial analysis and stakeholder validation.
</instructions>
</business_assumptions>

## Validation Pipeline

<validation_pipeline>

### High Priority (Immediate Validation Required)

{{#each high_priority_validations}}
- {{assumption_id}}: {{assumption_statement}}
  - Method: {{validation_method}}
  - Timeline: {{validation_deadline}}
  - Owner: {{assumption_owner}}
  - Status: {{validation_status}}
{{/each}}

### Medium Priority (Planned Validation)

{{#each medium_priority_validations}}
- {{assumption_id}}: {{assumption_statement}}
  - Method: {{validation_method}}
  - Timeline: {{validation_deadline}}
  - Owner: {{assumption_owner}}
  - Status: {{validation_status}}
{{/each}}

### Low Priority (Monitor Only)

{{#each low_priority_validations}}
- {{assumption_id}}: {{assumption_statement}}
  - Method: {{validation_method}}
  - Status: {{validation_status}}
{{/each}}

<instructions>
Validation pipeline organizes assumptions by urgency and validation requirements. High priority requires immediate action with deadlines.

Priority Levels:
- High: Immediate validation required, critical to project success
- Medium: Planned validation, affects timeline but not core viability
- Low: Monitor only, minimal project impact

CRITICAL: High priority items must have assigned owners and specific deadlines.
</instructions>
</validation_pipeline>

## Validation Results Log

### Recently Validated Assumptions

{{#each recent_validations}}

#### {{assumption_id}}. {{assumption_statement}}

- **Validation Date**: {{validation_date}}
- **Method Used**: {{validation_method}}
- **Result**: {{validation_result}}
- **Confidence Level**: {{new_confidence_level}}%
- **Evidence**: {{validation_evidence}}
- **Impact on Project Brief**: {{brief_impact}}

**Key Findings**: {{key_findings}}
**Action Items**: {{action_items}}

---
{{/each}}

### Validation Failures/Challenges

{{#each validation_failures}}

#### {{assumption_id}}. {{assumption_statement}}

- **Validation Date**: {{validation_date}}
- **Method Used**: {{validation_method}}
- **Issue**: {{validation_issue}}
- **Next Steps**: {{next_steps}}
- **Revised Timeline**: {{revised_timeline}}

---
{{/each}}

## Assumption Risk Analysis

### High Risk Assumptions (Impact: HIGH, Confidence: <70%)

{{#each high_risk_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Current Confidence**: {{confidence_level}}%
- **Risk Level**: {{risk_level}}
- **Mitigation Plan**: {{mitigation_plan}}
{{/each}}

### Assumptions Requiring Immediate Attention

{{#each immediate_attention_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Issue**: {{attention_reason}}
- **Deadline**: {{attention_deadline}}
- **Owner**: {{assumption_owner}}
{{/each}}

## Evidence Quality Assessment

### Strong Evidence (Confidence 80-100%)

{{#each strong_evidence_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Evidence Type**: {{evidence_type}}
- **Source**: {{evidence_source}}
- **Confidence**: {{confidence_level}}%
{{/each}}

### Moderate Evidence (Confidence 60-79%)

{{#each moderate_evidence_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Evidence Type**: {{evidence_type}}
- **Source**: {{evidence_source}}
- **Confidence**: {{confidence_level}}%
- **Improvement Needed**: {{improvement_plan}}
{{/each}}

### Weak Evidence (Confidence <60%)

{{#each weak_evidence_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Evidence Type**: {{evidence_type}}
- **Source**: {{evidence_source}}
- **Confidence**: {{confidence_level}}%
- **Validation Priority**: HIGH
- **Validation Plan**: {{validation_plan}}
{{/each}}

## Next Steps and Action Items

<action_items>

### Immediate Actions (This Week)

{{#each immediate_actions}}
- {{action_item}}
  - Owner: {{action_owner}}
  - Due Date: {{due_date}}
  - Related Assumptions: {{related_assumptions}}
{{/each}}

### Short-term Actions (Next 2 Weeks)

{{#each short_term_actions}}
- {{action_item}}
  - Owner: {{action_owner}}
  - Due Date: {{due_date}}
  - Related Assumptions: {{related_assumptions}}
{{/each}}

### Long-term Actions (Next Month)

{{#each long_term_actions}}
- {{action_item}}
  - Owner: {{action_owner}}
  - Due Date: {{due_date}}
  - Related Assumptions: {{related_assumptions}}
{{/each}}

<instructions>
Action items drive assumption validation forward. Organize by timeline urgency.

Timeline Categories:
- Immediate (This Week): Critical validation activities
- Short-term (Next 2 Weeks): Planned validation work
- Long-term (Next Month): Strategic validation initiatives

Required fields: clear action description, assigned owner, specific due date, related assumption IDs.
</instructions>
</action_items>

---

## Template Usage Notes

<usage_notes>

### Assumption Impact Levels

- HIGH: Assumption failure significantly affects project success or direction
- MEDIUM: Assumption failure causes delays or scope changes but project remains viable
- LOW: Assumption failure has minimal impact on project outcomes

### Confidence Levels

- 80-100%: Strong evidence supporting assumption
- 60-79%: Moderate evidence, some validation needed
- 40-59%: Limited evidence, validation required
- 0-39%: Weak or no evidence, immediate validation critical

### Validation Status

- VALIDATED: Assumption confirmed with strong evidence
- IN_PROGRESS: Validation activities underway
- PLANNED: Validation scheduled but not started
- DEFERRED: Validation postponed due to low priority/impact
- FAILED: Assumption proven false, project brief update needed

### Evidence Types

- PRIMARY: Direct user research, customer interviews, first-party data
- SECONDARY: Industry reports, competitive analysis, expert opinions
- QUANTIFIED: Numerical data, measurements, statistical analysis
- QUALITATIVE: User feedback, expert insights, observational data

<instructions>
Template usage notes provide consistent standards for assumption tracking and validation.

CRITICAL: Use these definitions consistently across all assumption entries to ensure accurate risk assessment and prioritization.

Impact levels determine validation urgency. Confidence levels guide evidence collection needs. Status tracking enables progress monitoring.
</instructions>
</usage_notes>

==== END FILE ====

==== FILE: .krci-ai/templates/context-gathering-guide-template.md ====
# Context Gathering Guide: {{project_name}}

> **Purpose**: Structured guide for systematic project context collection
> **Target Output**: `/docs/prd/project-context.md`
> **Framework Integration**: Business methodology-driven approach
> **Quality Standard**: Evidence-based with confidence assessments

---

## Context Collection Methodology Selection

### Available Input Methods

Select appropriate methods based on available resources and project needs:

#### Method 1: Document Analysis

Choose when: Existing research, competitive analysis, user studies, or PRDs available
Required Resources:
- Market research reports
- Competitive analysis documents
- User research and interview data
- Existing PRDs or product specifications
- Technical feasibility studies
- Customer support tickets and feedback
- Analytics and user behavior data

<instructions>
Document analysis method is most effective when substantial existing research is available.

Use this method for:
- Projects with existing market research or user studies
- Follow-up initiatives with prior analysis
- Competitive landscape analysis projects
- Data-rich environments with analytics available

CRITICAL: Always assess document relevance and recency before including in analysis.
</instructions>

**📊 Analysis Framework**:

```markdown
**Document Type**: {{document_type}}
**Source**: {{document_source}}
**Date**: {{document_date}}
**Relevance Score**: {{relevance_score}}/10
**Reliability Score**: {{reliability_score}}/10

**Key Insights Extracted**:
- {{insight_1}}
- {{insight_2}}
- {{insight_3}}

**Supporting Evidence**:
- {{evidence_1}}
- {{evidence_2}}

**Confidence Level**: {{confidence_percentage}}%
**Gaps Identified**: {{knowledge_gaps}}
```

#### Method 2: Stakeholder Interview Process

Choose when: Key stakeholders available for structured input gathering
Required Stakeholders:
- Business stakeholders (product owners, executives)
- User representatives (customer success, support)
- Technical stakeholders (architects, lead developers)
- Market experts (sales, marketing, customer success)

<instructions>
Stakeholder interviews provide primary evidence and diverse perspectives.

Use this method for:
- New initiatives requiring broad stakeholder input
- Complex projects with multiple stakeholder perspectives
- Situations where assumptions need validation
- Projects lacking existing documentation

CRITICAL: Use structured interview guides to ensure consistent data collection across stakeholders.
</instructions>

**🎯 Interview Guide Structure**:

```markdown
**Stakeholder**: {{stakeholder_name}} ({{role}})
**Interview Date**: {{interview_date}}
**Duration**: {{interview_duration}}
**Framework Used**: {{business_framework}} (e.g., BABOK, Design Thinking)

**Opening Context Questions**:
1. How would you describe the current problem/opportunity?
2. What evidence do you have that this problem exists?
3. Who is most affected by this problem?

**Problem Deep Dive** (using 5 Whys framework):
1. What is the core problem?
   - Why does this problem exist?
   - Why does [previous answer] occur?
   - Continue for 5 levels...

**User Context Questions** (using Jobs-to-be-Done framework):
1. When users encounter this problem, what are they trying to accomplish?
2. What is their desired outcome?
3. What current solutions do they use?
4. What gaps exist in current solutions?

**Business Context Questions**:
1. What business impact does this problem have?
2. How do you measure this impact?
3. What constraints exist for solving this problem?
4. What success would look like?

**Evidence and Validation**:
1. What data supports your perspective?
2. How confident are you in this assessment? (1-10)
3. What would change your mind about this?
4. Who else should we talk to about this?

**Key Quotes**:
- "{{important_quote_1}}"
- "{{important_quote_2}}"

**Action Items**:
- {{follow_up_item_1}}
- {{follow_up_item_2}}
```

#### Method 3: Assumption Inventory Creation

Choose when: Limited existing research, need to identify knowledge gaps
Assumption Categories Framework:

<instructions>
Assumption inventory method helps identify and track knowledge gaps systematically.

Use this method for:
- Early-stage projects with limited validation
- High-risk initiatives with many unknowns
- Innovation projects in new markets
- Situations requiring explicit risk identification

CRITICAL: Assign confidence levels and validation methods to all assumptions for effective tracking.
</instructions>

```markdown
## Problem Assumptions
**Category**: Problem Definition and Scope
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}

## User Assumptions
**Category**: Target Users and Behaviors
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}

## Solution Assumptions
**Category**: Solution Approach and Feasibility
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}

## Market Assumptions
**Category**: Market Size and Competitive Landscape
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}

## Business Assumptions
**Category**: Business Model and Success Metrics
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}
```

#### Method 4: Evidence Library Creation

**✅ Choose when**: Quantified data available, need systematic evidence collection
**📊 Evidence Classification Framework**:

```markdown
## Primary Evidence (Highest Confidence)
### {{evidence_type}}: {{evidence_title}}
- **Source**: {{evidence_source}}
- **Collection Date**: {{collection_date}}
- **Sample Size**: {{sample_size}}
- **Methodology**: {{collection_methodology}}
- **Reliability Score**: {{reliability_score}}/10
- **Relevance Score**: {{relevance_score}}/10
- **Recency Score**: {{recency_score}}/10
- **Overall Quality**: {{quality_score}}/10

**Key Data Points**:
- {{data_point_1}}: {{value_1}}
- {{data_point_2}}: {{value_2}}
- {{data_point_3}}: {{value_3}}

**Insights**:
- {{insight_1}}
- {{insight_2}}

**Limitations**:
- {{limitation_1}}
- {{limitation_2}}

## Secondary Evidence (Moderate Confidence)
### {{evidence_type}}: {{evidence_title}}
- **Source**: {{evidence_source}}
- **Publication Date**: {{publication_date}}
- **Relevance Score**: {{relevance_score}}/10
- **Credibility Score**: {{credibility_score}}/10
- **Overall Quality**: {{quality_score}}/10

**Key Insights**:
- {{insight_1}}
- {{insight_2}}

**Supporting Data**:
- {{data_point_1}}
- {{data_point_2}}
```

---

## Business Framework Application Guide

<framework_application>

### Framework Selection by Context Type

#### Problem Context Analysis

Recommended Frameworks:
1. Six Sigma 5 Whys - For root cause identification
2. Impact-Frequency Matrix - For problem prioritization
3. SIPOC Analysis - For process understanding

<instructions>
Framework selection should match the type of analysis needed and available evidence.

Problem Context Guidelines:
- Use 5 Whys when root cause identification is unclear
- Apply Impact-Frequency Matrix when multiple problems need prioritization
- Use SIPOC when process understanding is required

CRITICAL: Document framework rationale and application process for transparency.
</instructions>
</framework_application>

**Application Template**:

```markdown
## Problem Context Analysis

### Framework: {{selected_framework}}
**Why this framework**: {{framework_rationale}}
**Application process**: {{process_steps}}

### Analysis Results:
**Root Cause Identified**: {{root_cause}}
**Problem Priority Score**: {{priority_score}}/25
**Stakeholder Impact Map**: {{stakeholder_impacts}}
**Process Flow Analysis**: {{process_insights}}

### Evidence Supporting Analysis:
- {{evidence_1}} (Confidence: {{confidence_1}}%)
- {{evidence_2}} (Confidence: {{confidence_2}}%)

### Confidence Assessment: {{overall_confidence}}%
```

#### User Context Analysis

**Recommended Frameworks**:
1. **Jobs-to-be-Done** - For user motivation understanding
2. **Design Thinking Empathy Maps** - For user perspective mapping
3. **User Journey Mapping** - For experience flow analysis

**Application Template**:

```markdown
## User Context Analysis

### Framework: {{selected_framework}}
**User Segment**: {{target_user_segment}}
**Job Statement**: When I {{situation}}, I want to {{motivation}}, so I can {{outcome}}

### Analysis Results:
**Functional Job**: {{functional_job}}
**Emotional Job**: {{emotional_job}}
**Social Job**: {{social_job}}
**Job Importance**: {{importance_score}}/10
**Current Satisfaction**: {{satisfaction_score}}/10
**Opportunity Score**: {{opportunity_score}}

### User Journey Stages:
1. **{{stage_1}}**: {{stage_1_analysis}}
2. **{{stage_2}}**: {{stage_2_analysis}}
3. **{{stage_3}}**: {{stage_3_analysis}}

### Pain Points Identified:
- {{pain_point_1}} (Impact: {{impact_1}}/10)
- {{pain_point_2}} (Impact: {{impact_2}}/10)

### Evidence Supporting Analysis:
- {{evidence_1}} (Confidence: {{confidence_1}}%)
- {{evidence_2}} (Confidence: {{confidence_2}}%)

### Confidence Assessment: {{overall_confidence}}%
```

#### Market Context Analysis

**Recommended Frameworks**:
1. **Business Model Canvas** - For value proposition analysis
2. **Porter's Five Forces** - For competitive assessment
3. **TAM/SAM/SOM Analysis** - For market sizing

#### Business Context Analysis

**Recommended Frameworks**:
1. **SWOT Analysis** - For capability and environment assessment
2. **OKR Alignment** - For strategic goal hierarchy
3. **Resource Assessment** - For constraint analysis

---

## Context Synthesis Template

```markdown
# Project Context Summary: {{project_name}}

## Context Collection Summary
- **Methods Used**: {{methods_list}}
- **Timeframe**: {{collection_timeframe}}
- **Stakeholders Involved**: {{stakeholder_count}}
- **Evidence Sources**: {{evidence_source_count}}
- **Overall Confidence**: {{overall_confidence}}%

## Problem Context
### Root Problem Analysis
{{problem_analysis_summary}}

### Problem Evidence
- **Primary Evidence**: {{primary_problem_evidence}}
- **Supporting Data**: {{supporting_problem_data}}
- **Confidence Level**: {{problem_confidence}}%

### Problem Scope and Impact
{{problem_scope_definition}}

## User Context
### Target User Segments
{{user_segment_analysis}}

### User Jobs and Motivations
{{jobs_to_be_done_summary}}

### User Journey and Pain Points
{{user_journey_summary}}

### User Evidence
- **Primary Research**: {{primary_user_evidence}}
- **Supporting Data**: {{supporting_user_data}}
- **Confidence Level**: {{user_confidence}}%

## Market Context
### Market Opportunity
{{market_opportunity_analysis}}

### Competitive Landscape
{{competitive_analysis_summary}}

### Market Evidence
- **Market Data**: {{market_evidence}}
- **Competitive Research**: {{competitive_evidence}}
- **Confidence Level**: {{market_confidence}}%

## Business Context
### Strategic Alignment
{{strategic_alignment_summary}}

### Resource and Constraint Analysis
{{resource_constraint_analysis}}

### Business Evidence
- **Strategic Documentation**: {{strategic_evidence}}
- **Resource Assessment**: {{resource_evidence}}
- **Confidence Level**: {{business_confidence}}%

## Evidence Library Summary
### High Confidence Evidence (80-100%)
{{high_confidence_evidence_list}}

### Medium Confidence Evidence (60-79%)
{{medium_confidence_evidence_list}}

### Low Confidence Evidence (<60%)
{{low_confidence_evidence_list}}

## Assumption Inventory Summary
### High Impact Assumptions (Require Validation)
{{high_impact_assumptions}}

### Medium Impact Assumptions (Monitor)
{{medium_impact_assumptions}}

### Low Impact Assumptions (Accept)
{{low_impact_assumptions}}

## Next Steps for Project Brief Creation
### Ready for Brief Creation
{{ready_areas}}

### Require Additional Context
{{additional_context_needed}}

### Recommended Validation Activities
{{recommended_validations}}

## Quality Assessment
- **Context Completeness**: {{completeness_score}}%
- **Evidence Quality**: {{evidence_quality_score}}%
- **Stakeholder Coverage**: {{stakeholder_coverage_score}}%
- **Framework Application**: {{framework_application_score}}%
- **Overall Context Quality**: {{overall_quality_score}}%
```

---

## Quality Standards and Best Practices

<quality_standards>

### Context Collection Quality Criteria

#### Evidence Quality Standards

- Primary Evidence Target: >70% of key insights from primary sources
- Recency Requirement: Evidence <6 months old for dynamic markets, <12 months for stable markets
- Source Diversity: Minimum 3 different evidence types for key findings
- Confidence Threshold: >60% confidence required for high-impact assumptions

<instructions>
Quality standards ensure context gathering produces reliable, actionable insights.

Evidence Quality Guidelines:
- Prioritize primary evidence (interviews, direct research) over secondary sources
- Weight recent evidence more heavily than historical data
- Validate key findings across multiple source types
- Assign and track confidence levels for all major insights

CRITICAL: Maintain evidence traceability from insights back to original sources.
</instructions>
</quality_standards>

#### Stakeholder Coverage Standards

- **Business Stakeholder**: Minimum 1 executive/decision maker interview
- **User Representative**: Minimum 1 customer-facing role interview
- **Technical Stakeholder**: Minimum 1 technical lead interview
- **Market Expert**: Minimum 1 sales/marketing role interview

#### Framework Application Standards

- **Methodology Documentation**: Clear rationale for framework selection
- **Process Adherence**: Follow framework steps systematically
- **Results Documentation**: Complete analysis with supporting evidence
- **Quality Assessment**: Rate framework application effectiveness

### Context Documentation Standards

#### Structure Requirements

- **Executive Summary**: 1 page maximum context overview
- **Detailed Analysis**: Organized by context type (problem, user, market, business)
- **Evidence Library**: Comprehensive source documentation
- **Assumption Inventory**: Complete assumption tracking

#### Quality Requirements

- **Evidence Attribution**: All insights traced to specific sources
- **Confidence Levels**: All findings rated for reliability
- **Gap Identification**: Areas needing additional context clearly marked
- **Action Items**: Next steps for project brief creation documented

---

## Template Usage Instructions

### When to Use This Guide

- **New Projects**: Starting project brief creation from scratch
- **Context Refresh**: Updating existing project understanding
- **Validation Preparation**: Preparing for business framework validation
- **Stakeholder Alignment**: Building shared understanding of project context

### How to Customize This Guide

- **Select Relevant Methods**: Choose input methods based on available resources
- **Adapt Framework Applications**: Select business frameworks appropriate to project type
- **Customize Evidence Standards**: Adjust quality thresholds based on project risk/importance
- **Tailor Stakeholder Coverage**: Include project-specific stakeholder categories

### Integration with Project Brief Creation

- **Context → Brief**: Direct input for initial project brief creation
- **Evidence → Validation**: Foundation for business framework validation
- **Assumptions → Tracking**: Input for assumption tracker creation
- **Quality → Confidence**: Basis for project brief confidence levels

This guide ensures systematic, high-quality context gathering that enables evidence-based project brief creation with appropriate validation foundation.

==== END FILE ====

==== FILE: .krci-ai/templates/prd-template.md ====
# Product Requirements Document: {{product_name}}

## 1. Problem/Opportunity

<problem_opportunity>
{{problem_statement}}

**Evidence:**
{{supporting_evidence}}

<instructions>
Be crisp and clear about what user or business problem you're solving.

AVOID: "User can't use [solution]" - this is NOT a problem statement
FOCUS: What issues are caused when functionality is missing?

Problem Example: Users frequently forget passwords leading to 15% of support tickets and account lockouts (2,500/month). This creates frustration and consumes 15% of support resources.

Evidence Example: User research shows 65% of users reset passwords monthly. Support ticket analysis reveals 2,500 password-related tickets costing $50K annually.
</instructions>
</problem_opportunity>

---

## 2. Target Users & Use Cases

<target_users_use_cases>
**Primary Users:**
{{primary_users}}

**Key Use Cases:**
{{key_use_cases}}

<instructions>
Always focus on the user - this aligns building with go-to-market.
Be specific about users and use cases, ensure team alignment on definitions.

Primary Users Example: SaaS platform users (10,000+ monthly active) who access platform 3+ times weekly. Primary segment: business professionals aged 25-45 accessing from mobile (60%) and desktop (40%).

Key Use Cases Example:
1. Daily login for work tasks (highest frequency)
2. Password recovery when locked out (highest pain)
3. Multi-device access synchronization (growing need)
</instructions>
</target_users_use_cases>

---

## 3. Current Journeys/Landscape *(Optional)*

<current_journeys_landscape>
**Current User Journey:**
{{current_journey}}

**Competitive Landscape:**
{{competitive_analysis}}

<instructions>
Give context on what users do today or how competitors solve this.
Quick summary + links to detailed materials.

Current Journey Example: Users must remember complex passwords, leading to frequent lockouts. Recovery process takes 5-10 minutes via email verification.

Competitive Analysis: Auth0, Okta provide enterprise solutions. Consumer apps use Face ID/Touch ID. Gap exists for SMB-focused authentication.

Links: [Detailed user journey flow](link) or [Competitive analysis doc](link)
</instructions>
</current_journeys_landscape>

---

## 4. Proposed Solution/Elevator Pitch

<proposed_solution>
**Elevator Pitch:**
{{elevator_pitch}}

**Top 3 MVP Value Props:**

1. {{value_prop_1}}
2. {{value_prop_2}}
3. {{value_prop_3}}

**Conceptual Model:**
{{conceptual_model}}

<instructions>
Standard 2-3 liner in plain English.
Include top 3 MVP value props + conceptual model.

Elevator Pitch Example: Enable users to login using biometric authentication (fingerprint/face) and social login options, reducing password dependency by 80% while maintaining enterprise security standards.

Value Props Example:
1. 3-second biometric login eliminates password frustration
2. Social login reduces new user signup friction
3. Enterprise security maintains compliance requirements

Conceptual Model: [Include simple diagram or description of how users will interact with the solution]
</instructions>
</proposed_solution>

---

## 5. Goals/Measurable Outcomes

<goals_measurable_outcomes>
**Success Metrics:**

1. {{success_metric_1}}
2. {{success_metric_2}}
3. {{success_metric_3}}

<instructions>
Literally 2-3 bullets, no more.
Measurable outcomes defining success or non-failure.

Success Metrics Example:
1. Reduce password-related support tickets by 80% within 3 months
2. Achieve 70% user adoption of new auth methods within 6 months
3. Improve login success rate from 85% to 95%

AVOID vague statements like "improve user experience" or "increase engagement"
</instructions>
</goals_measurable_outcomes>

---

## 6. MVP/Functional Requirements

<mvp_functional_requirements>
### Business Requirements (BR)

**BR1**: {{business_requirement_1}}
**BR2**: {{business_requirement_2}}
**BR3**: {{business_requirement_3}}

### Non-Functional Requirements (NFR)

**NFR1**: {{system_requirement_1}}
**NFR2**: {{system_requirement_2}}
**NFR3**: {{system_requirement_3}}

<instructions>
Critical: Focus on required functionality, save the rest for future phases.
Question: What's the "min-viable" set of functionality for target user adoption?

Format: Focus on functionality, not implementation

DO:
- "First-time user must accept privacy policy to use product"
- "Product team can monitor and visualize user engagement"
- Link to UX sketches for quick visualization
- Include priorities: [P0] [P1] [P2] where P0 = truly required for MVP
- Bucket by use case/user journey for Epic creation
- Consider all critical user journeys (CUJs) - create, maintain, retire, navigate
- Limit to 3 phases/milestones maximum

DON'T:
- Performance metrics unless required for adoption
- Design details like "blue 'Continue' button"
- Technical implementation specifics

Business Requirements (BR) Examples:
BR1 [P0]: User can login using biometric authentication with <3 second response
BR2 [P1]: User can view login history with timestamps and device info
BR3 [P2]: Admin can configure password complexity requirements

Non-Functional Requirements (NFR) Examples:
NFR1 [P0]: System supports 1000 concurrent users with <2 second response time
NFR2 [P1]: System maintains 99.9% uptime during business hours
NFR3 [P2]: System integrates with enterprise SSO solutions

Use Case Buckets for Epic Creation:
Epic 1: Authentication & Security
- BR1: Biometric authentication implementation
- NFR1: Performance and scalability requirements

Epic 2: User Management
- BR2: User history and account features
- NFR2: System reliability requirements

Each bucket should map to an Epic following SDLC naming: {epic_number}-epic-{slug}.md
</instructions>
</mvp_functional_requirements>

==== END FILE ====

# Reference Data

==== FILE: .krci-ai/data/business-frameworks.md ====
# Business Analysis Frameworks and Models

## Requirements Analysis Frameworks

<requirements_frameworks>

### BABOK (Business Analysis Body of Knowledge)

Comprehensive framework for business analysis practices and techniques.

- Knowledge Areas: Business Analysis Planning, Elicitation, Requirements Management, Solution Assessment
- Techniques: Interviews, Workshops, Document Analysis, Observation, Surveys
- Deliverables: Requirements Documentation, Stakeholder Analysis, Solution Assessment
- Application: Use for structured requirements gathering and analysis projects

### MoSCoW Prioritization

Framework for prioritizing requirements based on business importance.

- Must Have: Critical requirements without which the solution fails
- Should Have: Important requirements that add significant value
- Could Have: Desirable requirements that enhance the solution
- Won't Have: Requirements that are out of scope for current iteration

### Kano Model

Framework for understanding customer satisfaction with product features.

- Must-be Quality: Basic expectations that cause dissatisfaction if missing
- One-dimensional Quality: Features that increase satisfaction linearly
- Attractive Quality: Unexpected features that delight customers
- Indifferent Quality: Features that don't significantly impact satisfaction
</requirements_frameworks>

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

## Presentation and Persuasion Frameworks

### Problem-Agitation-Solution (PAS) Framework

Proven copywriting and presentation framework for compelling narrative structure.

- **Problem**: Identify specific pain point audience experiences daily
- **Agitation**: Amplify emotional impact and consequences of the problem
- **Solution**: Present your product as the clear resolution with evidence
- **Application**: Ideal for 3-slide presentations (Problem → Agitation → Solution)

### Before-After-Bridge (BAB) Framework

Transformation-focused framework for demonstrating value proposition.

- **Before**: Current undesirable state with specific pain points
- **After**: Desired future state with quantified benefits
- **Bridge**: Your solution as the path from Before to After
- **Application**: Perfect for solution-focused presentations and demos

### STAR Method for Proof Points

Structured approach for presenting evidence and case studies.

- **Situation**: Context and background of the challenge
- **Task**: Specific objective or goal to be achieved
- **Action**: Steps taken to address the situation
- **Result**: Quantified outcomes and measurable impact
- **Application**: Use for testimonials, case studies, and traction slides

### SCRAP Framework for Concise Pitches

Minimalist framework optimized for executive presentations.

- **Situation**: Market context and opportunity size
- **Complication**: Problem preventing success
- **Resolution**: Your solution approach
- **Action**: Specific next steps required
- **Payoff**: Expected return on investment
- **Application**: Ideal for 4-5 slide investor or executive presentations

### Persuasion Psychology Principles

Core psychological triggers for compelling presentations.

- **Social Proof**: Testimonials, user numbers, industry recognition
- **Authority**: Expert endorsements, certifications, credentials
- **Scarcity**: Limited availability, exclusive access, time sensitivity
- **Reciprocity**: Value provided before asking for commitment
- **Consistency**: Alignment with audience values and commitments
- **Liking**: Similarity, familiarity, and shared interests

### Narrative Arc Framework

Story structure for emotional engagement and retention.

- **Hook**: Attention-grabbing opening (statistic, question, story)
- **Exposition**: Context and background information
- **Rising Action**: Building tension and problem awareness
- **Climax**: Solution reveal and "AHA moment"
- **Resolution**: Outcomes and call to action
- **Application**: Use for longer presentations requiring emotional journey

### Pain-Gains-Reveals Framework

Product-focused framework for customer-centric presentations.

- **Pain**: Contextualize critical problems with emotional resonance
- **Gains**: Articulate tangible benefits and quantifiable value proposition
- **Reveals**: Highlight unique differentiators and competitive advantages
- **Application**: Ideal for product overview decks and customer presentations

### AIDA Framework

Classic marketing framework adapted for presentations.

- **Attention**: Capture audience focus immediately
- **Interest**: Maintain engagement with relevant content
- **Desire**: Create want for your solution
- **Action**: Drive specific next steps
- **Application**: Structure for any persuasive presentation or pitch

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

==== FILE: .krci-ai/data/validation-frameworks.md ====
# Business Framework Validation Library

Purpose: Provide methodologically grounded validation frameworks for project brief enhancement. Each framework uses established business methodologies to validate specific aspects of project briefs with evidence-based approaches.

## Problem Validation Frameworks

<problem_validation>

### Lean Startup Problem-Solution Fit Assessment

Methodology: Eric Ries Lean Startup approach for validating problem-solution alignment.

Application: Use for Problem Statement and Opportunity sections validation.

Validation Process:

1. Problem Hypothesis Validation
   - Define problem hypothesis with specific user segment
   - Identify measurable problem indicators (support tickets, workarounds, time spent)
   - Validate problem urgency using customer interview evidence
   - Quantify problem frequency and impact

2. Solution Hypothesis Validation
   - Define minimum viable solution approach
   - Test solution desirability with target users
   - Validate solution feasibility with technical constraints
   - Assess solution viability with business model fit

3. Evidence Requirements
   - Customer interview data (minimum 10 target users)
   - Quantified problem metrics (costs, time, frequency)
   - Competitive analysis of existing solutions
   - Technical feasibility assessment

Validation Questions:

- Is this problem urgent enough for users to seek solutions?
- Do users currently pay (time/money) for workarounds?
- Does our solution approach address the root cause?
- Is the solution technically and commercially viable?
</problem_validation>

### Six Sigma Root Cause Analysis (5 Whys)

**Methodology**: Six Sigma DMAIC approach using 5 Whys technique for root cause identification.

**Application**: Use for Problem Statement depth validation and constraint analysis.

**Validation Process**:

1. **Problem Statement Definition**
   - State observable problem with quantified impact
   - Ask "Why does this problem occur?" - record first cause
   - Ask "Why does [first cause] occur?" - record second cause
   - Continue for minimum 5 levels of "why" questions
   - Validate each cause level with supporting evidence

2. **Evidence Requirements**
   - Data supporting each "why" level
   - Process documentation showing cause relationships
   - Historical data demonstrating problem persistence
   - Stakeholder confirmation of root causes

**Validation Questions**:

- Is each "why" level supported by evidence?
- Do the root causes align with proposed solution approach?
- Are root causes within our control to address?
- Have similar root causes been successfully addressed elsewhere?

### Impact-Frequency Problem Assessment

**Methodology**: Risk management matrix approach for problem prioritization validation.

**Application**: Use for Problem Statement prioritization and resource allocation validation.

**Validation Process**:

1. **Impact Assessment** (Scale 1-5)
   - User productivity impact
   - Business revenue impact
   - Operational cost impact
   - Strategic goal alignment impact

2. **Frequency Assessment** (Scale 1-5)
   - How often problem occurs
   - Number of users affected
   - Trend analysis (increasing/stable/decreasing)
   - Seasonal or cyclical patterns

3. **Priority Matrix Calculation**
   - Priority Score = Impact × Frequency
   - High Priority: 16-25 (urgent action required)
   - Medium Priority: 6-15 (planned action needed)
   - Low Priority: 1-5 (monitor and defer)

**Evidence Requirements**:

- Quantified impact metrics with historical data
- User activity data showing frequency patterns
- Trend analysis over minimum 6-month period
- Comparative analysis with other organizational priorities

## User Validation Frameworks

<user_validation>

### Jobs-to-be-Done Validation (Clayton Christensen)

**Methodology**: Clayton Christensen's Jobs-to-be-Done framework for user need validation.

**Application**: Use for Target Users section validation and use case prioritization.

**Validation Process**:

1. **Job Statement Formulation**
   - When I [situation], I want to [motivation], so I can [expected outcome]
   - Validate job statement with target user interviews
   - Identify functional, emotional, and social job dimensions
   - Map job executor demographics and context

2. **Job Importance Assessment**
   - Rate job importance (1-10) from user perspective
   - Validate importance with willingness to pay data
   - Assess job frequency and urgency
   - Identify job outcome priorities

3. **Current Satisfaction Assessment**
   - Rate current solution satisfaction (1-10)
   - Identify specific satisfaction gaps
   - Map current solution journey and pain points
   - Validate dissatisfaction with switching behavior

4. **Opportunity Score Calculation**
   - Opportunity = Importance + max(Importance - Satisfaction, 0)
   - High Opportunity: 15+ (strong market opportunity)
   - Medium Opportunity: 12-15 (moderate market opportunity)
   - Low Opportunity: <12 (limited market opportunity)

**Evidence Requirements**:

- User interview data with job statement validation
- Quantified importance and satisfaction ratings
- Current solution usage and switching behavior data
- Demographic and contextual data for job executors

**Validation Questions**:

- Is the job statement validated by target users?
- Do users actively seek better solutions for this job?
- Is there sufficient dissatisfaction with current solutions?
- Does our solution approach align with job outcome priorities?

### Design Thinking User Journey Validation

**Methodology**: Design Thinking approach for user experience validation and empathy mapping.

**Application**: Use for Target Users section validation and user experience requirements.

**Validation Process**:

1. **Journey Stage Mapping**
   - Map complete user journey from awareness to outcome
   - Identify all touchpoints and interactions
   - Document user actions, thoughts, and emotions at each stage
   - Validate journey accuracy with user observation/interviews

2. **Pain Point Identification**
   - Identify specific pain points at each journey stage
   - Quantify pain point impact (time, effort, frustration)
   - Validate pain points with user evidence
   - Prioritize pain points by severity and frequency

3. **Opportunity Identification**
   - Identify improvement opportunities at each stage
   - Assess opportunity impact on user satisfaction
   - Validate opportunity desirability with users
   - Map opportunities to business value potential

4. **Empathy Map Creation**
   - What users Say (quotes and defining words)
   - What users Think (thoughts and beliefs)
   - What users Do (actions and behaviors)
   - What users Feel (emotions and feelings)

**Evidence Requirements**:

- User observation data or detailed user interviews
- Journey stage documentation with user validation
- Quantified pain point impact measurements
- Empathy map elements validated with user research

**Validation Questions**:

- Is the user journey accurately mapped and validated?
- Are pain points quantified and evidenced?
- Do improvement opportunities align with user priorities?
- Does the empathy map reflect authentic user perspectives?

### Persona Demographic Validation

**Methodology**: Data-driven persona validation using demographic and behavioral evidence.

**Application**: Use for Target Users section accuracy and market sizing validation.

**Validation Process**:

1. **Demographic Profile Validation**
   - Age, gender, education, income, role, industry
   - Geographic distribution and cultural factors
   - Technology adoption and digital behavior patterns
   - Validate demographics with market research data

2. **Behavioral Pattern Validation**
   - Usage frequency and session duration patterns
   - Device and platform preferences
   - Feature adoption and engagement metrics
   - Validate patterns with user analytics data

3. **Needs and Goals Validation**
   - Primary goals and success criteria
   - Secondary goals and nice-to-have outcomes
   - Barriers and constraints to goal achievement
   - Validate needs with user interview data

4. **Market Size Validation**
   - Total addressable market (TAM) calculation
   - Serviceable addressable market (SAM) estimation
   - Serviceable obtainable market (SOM) projection
   - Validate sizing with industry reports and data

**Evidence Requirements**:

- Market research data supporting demographic profiles
- User analytics data validating behavioral patterns
- User interview data supporting needs and goals
- Industry data supporting market size calculations

**Validation Questions**:

- Are demographic profiles supported by market data?
- Do behavioral patterns align with analytics evidence?
- Are user needs validated through direct user research?
- Is market sizing based on credible industry data?

</user_validation>

## Success Metrics Validation Frameworks

<metrics_validation>

### SMART Criteria Assessment

**Methodology**: SMART criteria framework for success metrics validation and quality assurance.

**Application**: Use for Success Metrics section validation and goal setting quality control.

**Validation Process**:

1. **Specific Validation**
   - Does the metric clearly define what will be measured?
   - Is the metric scope bounded and unambiguous?
   - Can stakeholders understand the metric without interpretation?
   - Is the metric aligned with project objectives?

2. **Measurable Validation**
   - Can the metric be quantified with specific numbers?
   - Are measurement tools and data sources identified?
   - Is the metric trackable throughout project lifecycle?
   - Are measurement intervals and reporting frequency defined?

3. **Achievable Validation**
   - Is the target realistic given available resources?
   - Are there precedents for similar metric achievements?
   - Have constraints and limitations been considered?
   - Is the metric within team/organization control or influence?

4. **Relevant Validation**
   - Does the metric align with business objectives?
   - Will achieving this metric create meaningful business value?
   - Is the metric important to key stakeholders?
   - Does the metric contribute to strategic goals?

5. **Time-bound Validation**
   - Is there a clear deadline or timeframe?
   - Are there intermediate milestones and checkpoints?
   - Is the timeline realistic for metric achievement?
   - Are dependencies and risks to timeline identified?

**Evidence Requirements**:

- Historical performance data for achievability assessment
- Measurement tool and data source documentation
- Stakeholder confirmation of relevance and importance
- Timeline feasibility analysis with dependency mapping

**Validation Questions**:

- Can this metric be measured accurately and consistently?
- Is the target achievable within specified constraints?
- Will achieving this metric create meaningful value?
- Is the timeline realistic given project dependencies?

### OKR Alignment Validation

**Methodology**: Objectives and Key Results framework for strategic alignment validation.

**Application**: Use for Success Metrics section strategic alignment and goal hierarchy validation.

**Validation Process**:

1. **Objective Alignment Validation**
   - Does project objective align with organizational OKRs?
   - Is the objective inspirational and directionally correct?
   - Does the objective contribute to higher-level strategic goals?
   - Is the objective clear enough to guide decision-making?

2. **Key Result Quality Validation**
   - Are key results measurable and outcome-focused?
   - Do key results collectively indicate objective achievement?
   - Are key results ambitious yet achievable (60-70% confidence)?
   - Do key results avoid activity-based metrics?

3. **OKR Hierarchy Validation**
   - How does project OKR cascade from organizational OKRs?
   - Are there conflicts with other team/department OKRs?
   - Is the OKR appropriately scoped for project level?
   - Does the OKR enable downstream team OKRs?

4. **Success Probability Assessment**
   - What's the confidence level for achieving each key result?
   - What are the key risks to OKR achievement?
   - What resources and dependencies are required?
   - How will progress be measured and reported?

**Evidence Requirements**:

- Organizational OKR documentation and alignment mapping
- Historical performance data for confidence assessment
- Resource availability and dependency analysis
- Risk assessment with mitigation strategies

**Validation Questions**:

- Do project OKRs clearly cascade from organizational strategy?
- Are key results appropriately ambitious and measurable?
- Is there organizational alignment and resource commitment?
- How will OKR progress be tracked and communicated?

### Leading/Lagging Indicator Analysis

**Methodology**: Performance measurement theory for balanced metric portfolio validation.

**Application**: Use for Success Metrics section balance and predictive capability validation.

**Validation Process**:

1. **Leading Indicator Identification**
   - Identify metrics that predict future outcomes
   - Validate leading indicators with historical correlation data
   - Assess leading indicator controllability and actionability
   - Define leading indicator measurement frequency

2. **Lagging Indicator Validation**
   - Identify outcome metrics that measure final results
   - Validate lagging indicators as true measures of success
   - Assess lagging indicator accuracy and reliability
   - Define lagging indicator measurement and reporting process

3. **Balance Assessment**
   - Is there appropriate balance between leading and lagging indicators?
   - Do leading indicators provide early warning capability?
   - Do lagging indicators measure ultimate project success?
   - Are there gaps in the measurement framework?

4. **Predictive Relationship Validation**
   - What's the correlation between leading and lagging indicators?
   - What's the time lag between leading indicator changes and outcome impact?
   - Are there external factors that affect the relationship?
   - How reliable is the predictive relationship?

**Evidence Requirements**:

- Historical data showing leading/lagging indicator relationships
- Correlation analysis and statistical validation
- External factor analysis and impact assessment
- Measurement system reliability and accuracy validation

**Validation Questions**:

- Do leading indicators reliably predict lagging outcomes?
- Is there balanced coverage of controllable and outcome metrics?
- Can the team act on leading indicator signals effectively?
- Are lagging indicators true measures of project success?

</metrics_validation>

## Business Value Validation Frameworks

<business_validation>

### Business Model Canvas Validation

**Methodology**: Alexander Osterwalder's Business Model Canvas for business value validation.

**Application**: Use for Opportunity section business model validation and value creation assessment.

**Validation Process**:

1. **Value Proposition Validation**
   - What specific value does the project create for customers?
   - How does the value proposition address customer jobs, pains, and gains?
   - Is the value proposition differentiated from alternatives?
   - Can the value proposition be delivered sustainably?

2. **Customer Segment Validation**
   - Are customer segments clearly defined and reachable?
   - Is there evidence of customer demand for the value proposition?
   - Are customer segments large enough to justify investment?
   - Do customer segments align with organizational capabilities?

3. **Revenue Model Validation**
   - How will the project generate revenue or cost savings?
   - Are revenue assumptions validated with customer evidence?
   - Is the revenue model sustainable and scalable?
   - What's the timeline to revenue realization?

4. **Cost Structure Analysis**
   - What are the key costs associated with value delivery?
   - Are cost estimates based on reliable data and assumptions?
   - How do costs scale with customer growth?
   - What's the path to positive unit economics?

**Evidence Requirements**:

- Customer interview data validating value proposition
- Market research supporting customer segment assumptions
- Financial model with validated assumptions
- Cost analysis with scaling scenarios

**Validation Questions**:

- Is there validated customer demand for the value proposition?
- Are revenue assumptions based on customer evidence?
- Is the business model economically viable and scalable?
- What are the key risks to business model success?

### Value Proposition Canvas Fit Assessment

**Methodology**: Alexander Osterwalder's Value Proposition Canvas for customer-solution fit validation.

**Application**: Use for Opportunity section customer-solution alignment and value creation validation.

**Validation Process**:

1. **Customer Profile Validation**
   - **Customer Jobs**: Functional, emotional, social jobs customers are trying to accomplish
   - **Pains**: Bad outcomes, obstacles, risks customers experience or fear
   - **Gains**: Benefits customers want, expect, desire, or would be surprised by

2. **Value Map Validation**
   - **Products & Services**: List of products and services the value proposition builds on
   - **Pain Relievers**: How products/services alleviate specific customer pains
   - **Gain Creators**: How products/services create customer gains

3. **Fit Assessment**
   - **Problem-Solution Fit**: Do pain relievers address important customer pains?
   - **Product-Market Fit**: Do gain creators deliver outcomes customers want?
   - **Value Proposition Fit**: Is there strong alignment between customer needs and solution benefits?

4. **Evidence Collection**
   - Customer interview validation of jobs, pains, and gains
   - Solution testing with target customers
   - Customer willingness to pay evidence
   - Customer switching behavior validation

**Evidence Requirements**:

- Customer research validating jobs, pains, and gains
- Solution prototype testing with customer feedback
- Customer willingness to pay and value perception data
- Competitive analysis of alternative solutions

**Validation Questions**:

- Are customer jobs, pains, and gains validated with evidence?
- Do solution features directly address important customer needs?
- Is there strong product-market fit based on customer evidence?
- What's the strength of the value proposition compared to alternatives?

### ROI Calculation Validation

**Methodology**: Financial analysis framework for return on investment validation and business case assessment.

**Application**: Use for Opportunity section financial justification and business case validation.

**Validation Process**:

1. **Investment Cost Analysis**
   - Development costs (personnel, tools, infrastructure)
   - Implementation costs (deployment, training, change management)
   - Operational costs (maintenance, support, ongoing resources)
   - Opportunity costs (alternative investment options)

2. **Benefit Quantification**
   - Revenue generation (new sales, upselling, market expansion)
   - Cost savings (efficiency, automation, waste reduction)
   - Risk mitigation (avoided costs, compliance, security)
   - Strategic benefits (market position, capability building)

3. **ROI Calculation Methods**
   - **Simple ROI**: (Benefits - Investment) / Investment × 100
   - **Net Present Value (NPV)**: Present value of future cash flows minus initial investment
   - **Internal Rate of Return (IRR)**: Discount rate that makes NPV equal to zero
   - **Payback Period**: Time required to recover initial investment

4. **Sensitivity Analysis**
   - Best case, worst case, most likely scenario analysis
   - Key assumption impact on ROI calculations
   - Break-even analysis and threshold identification
   - Risk factor quantification and mitigation costs

**Evidence Requirements**:

- Detailed cost breakdown with vendor quotes and resource estimates
- Benefit quantification with historical data and industry benchmarks
- Financial projections with assumptions documentation
- Sensitivity analysis with scenario modeling

**Validation Questions**:

- Are cost estimates comprehensive and based on reliable data?
- Are benefit projections conservative and evidenced?
- Does the ROI meet organizational investment criteria?
- What are the key risks to ROI achievement and how can they be mitigated?

</business_validation>

## Framework Selection Guide

<selection_guide>

### Problem Statement Validation

- **Primary**: Lean Startup Problem-Solution Fit Assessment
- **Secondary**: Six Sigma Root Cause Analysis (5 Whys)
- **Situational**: Impact-Frequency Assessment (for prioritization)

### Target Users Validation

- **Primary**: Jobs-to-be-Done Validation
- **Secondary**: Design Thinking User Journey Validation
- **Situational**: Persona Demographic Validation (for market sizing)

### Success Metrics Validation

- **Primary**: SMART Criteria Assessment
- **Secondary**: OKR Alignment Validation
- **Situational**: Leading/Lagging Indicator Analysis (for balanced scorecards)

### Business Value Validation

- **Primary**: Value Proposition Canvas Fit Assessment
- **Secondary**: ROI Calculation Validation
- **Situational**: Business Model Canvas Validation (for new business models)

</selection_guide>

## Validation Confidence Levels

<confidence_levels>

### High Confidence (80-100%)

- Multiple validation frameworks applied
- Quantified evidence from primary sources
- Historical data supporting assumptions
- Expert validation and peer review

### Medium Confidence (60-79%)

- Primary validation framework applied
- Mix of quantified and qualitative evidence
- Industry benchmark data available
- Stakeholder validation completed

### Low Confidence (40-59%)

- Basic validation framework applied
- Limited quantified evidence
- Assumptions based on secondary research
- Expert opinion without validation

### Very Low Confidence (<40%)

- No formal validation applied
- Assumptions without supporting evidence
- No stakeholder validation
- High uncertainty in key assumptions

</confidence_levels>

## Usage Guidelines

<usage_guidelines>

### When to Apply Validation Frameworks

1. **Always Apply**: During initial project brief creation
2. **Apply Selectively**: When assumptions are challenged or evidence is weak
3. **Reapply**: When project scope, market, or constraints change significantly
4. **Skip**: When assumptions are well-validated with strong evidence

### Evidence Standards

- **Primary Evidence**: Customer interviews, user research, first-party data
- **Secondary Evidence**: Industry reports, competitive analysis, expert opinions
- **Quantified Evidence**: Numerical data, measurements, calculations
- **Qualitative Evidence**: User feedback, expert insights, observations

### Validation Documentation

- Document validation method used and rationale for selection
- Record evidence sources and confidence levels
- Track assumptions and validation status
- Update validation as new evidence becomes available
</usage_guidelines>

==== END FILE ====

