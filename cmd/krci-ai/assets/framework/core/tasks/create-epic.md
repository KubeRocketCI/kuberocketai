# Task: Create Epic

## Description

Create clear epic with problem statement, goal, scope, and implementation approach that breaks down PRD requirements into manageable high-level features. This epic enables Story creation and provides a clear feature grouping for development teams.

## Prerequisites

- [ ] **Completed PRD**: PRD exists at `/docs/prd/prd.md` with BR/NFR requirements defined
- [ ] **Epic priority identified**: Clear understanding of which PRD requirements this Epic addresses
- [ ] **User context available**: Target users and use cases from PRD understood
- [ ] **Epic scope defined**: Boundaries of what this Epic includes and excludes

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/prioritization-frameworks.md
- ./.krci-ai/templates/epic.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for Epic dependencies and artifact flow
2. **Apply prioritization**: Use methods from [prioritization-frameworks.md](./.krci-ai/data/prioritization-frameworks.md)
3. **Format output**: Use [epic.md](./.krci-ai/templates/epic.md) template with all variables populated
4. **Ensure PRD traceability**: Reference specific BR/NFR requirements from PRD
5. **Enforce agent-optimization**: Follow enhanced template requirements for consistent agent-parseable structure

## Agent-Optimized Template Enforcement

### Required Format Validation

**CRITICAL**: All epics must follow these agent-optimized patterns:

- **Goal Format**: MUST follow "Enable [specific outcome] for [target %] of [user type] within [timeline]"
- **Dependencies Structure**: MUST group using exact headers "Epic Dependencies:", "System Dependencies:", "External Dependencies:"
- **Acceptance Criteria**: MUST include validation method AND verification command for each criterion
- **User Stories**: MUST use phase-based grouping with "**Phase X: Name**" headers and dependency notation

### Quality Gates

Before epic creation, verify:

1. **Goal Measurability**: Goal contains specific outcome, target percentage, user type, and timeline
2. **Dependencies Grouping**: All dependencies properly categorized with required headers
3. **Acceptance Criteria Commands**: Each criterion includes both validation method and testable command
4. **Agent-Parseable Structure**: Consistent formatting enables automated processing

### Validation Commands

Post-creation verification:

- Goal format: `echo "[goal_text]" | grep -E "Enable .* for [0-9]+% .* within .*"`
- Dependencies structure: `grep -E "^(Epic|System|External) Dependencies:" [epic_file]`
- Acceptance commands: `grep -E "Command: \`.*\`" [epic_file]`

## Output Format

- **Location**: `/docs/epics/{epic_number}-epic-{slug}.md` (EXACT naming pattern)
- **Epic numbering**: Sequential number based on existing epics (check `/docs/epics/` folder)
- **Slug format**: Lowercase, hyphenated description (e.g., "ide-integration", "core-engine")
- **Downstream Enable**: Enables Story creation at `/docs/stories/`

## Success Criteria

- [ ] **File saved** to `/docs/epics/{epic_number}-epic-{slug}.md` with correct naming
- [ ] **PRD traceability** clear connection to specific BR/NFR requirements
- [ ] **Problem clarity** epic solves a specific user problem with defined scope
- [ ] **Goal measurability** epic completion criteria are specific and testable
- [ ] **Story readiness** epic provides enough detail for Story breakdown
- [ ] **Template compliance** all template variables populated correctly
- [ ] **Agent-optimization** epic follows enhanced template requirements for consistent structure
- [ ] **Goal format compliance** goal follows "Enable X for Y% of [users] within [timeline]" pattern
- [ ] **Dependencies structure** all dependencies grouped by Epic/System/External with exact headers
- [ ] **Acceptance criteria commands** each criterion includes validation method and verification command
- [ ] **Quality gates passed** all validation commands execute successfully

## Execution Checklist

### Discovery Phase

- [ ] **PRD analysis**: Review BR/NFR requirements to identify Epic scope
- [ ] **Epic numbering**: Check existing epics in `/docs/epics/` for next sequential number
- [ ] **Problem definition**: Identify specific user problem this Epic addresses
- [ ] **Scope boundaries**: Define what's included and excluded from this Epic
- [ ] **Goal format planning**: Draft goal using "Enable [outcome] for [%] of [users] within [timeline]" pattern
- [ ] **Dependencies analysis**: Identify and categorize by Epic/System/External groups

### Planning Phase

- [ ] **Goal definition**: Define clear, measurable Epic completion criteria following required format
- [ ] **User identification**: Specify target users from PRD who benefit from this Epic
- [ ] **Dependencies mapping**: Group dependencies by Epic/System/External with exact headers
- [ ] **Solution approach**: Define high-level implementation strategy
- [ ] **Acceptance criteria planning**: Identify testable validation methods and commands for each criterion
- [ ] **User stories structure**: Plan phase-based grouping with dependency notation

### Documentation Phase

- [ ] **Epic creation**: Use [epic.md](./.krci-ai/templates/epic.md) template structure
- [ ] **Variable population**: Complete all template variables ({{epic_number}}, {{epic_title}}, etc.)
- [ ] **Goal format validation**: Verify goal follows "Enable X for Y% within Z timeline" pattern
- [ ] **Dependencies grouping**: Apply Epic/System/External structure with exact headers
- [ ] **Acceptance criteria commands**: Add validation method AND verification command for each criterion
- [ ] **User stories formatting**: Use phase-based grouping with "**Phase X: Name**" headers
- [ ] **Template compliance check**: Verify all agent-optimized requirements are met
- [ ] **Quality gates validation**: Run validation commands to ensure structure compliance
- [ ] **File placement**: Save to exact location `/docs/epics/epic-{number}-{slug}.md`

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Problem-Focused**: Epic addresses specific user problems from PRD with quantifiable impact
- **Measurable Goal**: Epic completion criteria are specific, testable, and time-bound
- **PRD Traceable**: Maintain clear connection to specific BR/NFR requirements from PRD
- **Story-Ready**: Provide sufficient context and scope for immediate Story breakdown

### LLM Error Prevention Checklist

- **Avoid**: Technical implementation details (save for Stories and Architecture documents)
- **Avoid**: Vague problem statements without user context and quantifiable impact
- **Avoid**: Unmeasurable completion criteria that cannot be validated
- **Reference**: Use [epic.md](./.krci-ai/templates/epic.md) for all formatting guidance and examples

### SDLC Integration Context

This Epic enables immediate Story creation by providing clear problem context, defined user personas for "As a user" scenarios, measurable outcomes that become Story acceptance criteria, and scope boundaries that guide Story prioritization within the Epic.
