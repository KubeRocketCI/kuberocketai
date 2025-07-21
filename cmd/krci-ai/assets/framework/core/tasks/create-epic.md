# Task: Create Epic

## Description

Create clear epic with problem statement, goal, scope, and implementation approach that breaks down PRD requirements into manageable high-level features. This epic enables Story creation and provides a clear feature grouping for development teams.

## Framework Context

**Reference**: [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) - See role responsibilities and artifact flow

## Prerequisites

- [ ] **Completed PRD**: PRD exists at `/docs/prd/prd.md` with BR/NFR requirements defined
- [ ] **Epic priority identified**: Clear understanding of which PRD requirements this Epic addresses
- [ ] **User context available**: Target users and use cases from PRD understood
- [ ] **Epic scope defined**: Boundaries of what this Epic includes and excludes

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for Epic dependencies and artifact flow
2. **Apply prioritization**: Use methods from [prioritization-frameworks.md](./.krci-ai/data/prioritization-frameworks.md)
3. **Format output**: Use [epic.md](./.krci-ai/templates/epic.md) template with all variables populated
4. **Ensure PRD traceability**: Reference specific BR/NFR requirements from PRD

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

## Execution Checklist

### Discovery Phase

- [ ] **PRD analysis**: Review BR/NFR requirements to identify Epic scope
- [ ] **Epic numbering**: Check existing epics in `/docs/epics/` for next sequential number
- [ ] **Problem definition**: Identify specific user problem this Epic addresses
- [ ] **Scope boundaries**: Define what's included and excluded from this Epic

### Planning Phase

- [ ] **Goal definition**: Define clear, measurable Epic completion criteria
- [ ] **User identification**: Specify target users from PRD who benefit from this Epic
- [ ] **Dependencies mapping**: Identify other Epics, systems, or external dependencies
- [ ] **Solution approach**: Define high-level implementation strategy

### Documentation Phase

- [ ] **Epic creation**: Use [epic.md](./.krci-ai/templates/epic.md) template structure
- [ ] **Variable population**: Complete all template variables ({{epic_number}}, {{epic_title}}, etc.)
- [ ] **Content validation**: Ensure problem, goal, scope, and approach are clearly defined
- [ ] **File placement**: Save to exact location `/docs/epics/epic-{number}-{slug}.md`

## Content Guidelines

### 📋 **Epic Template Sections:**

1. **Status Table**: Epic number, title, status, priority, owner, timeline
2. **Overview**: Problem statement, goal, target users
3. **Scope**: What's included, what's excluded, dependencies
4. **Solution Approach**: High-level implementation strategy
5. **Acceptance Criteria**: How to know when Epic is complete
6. **User Stories**: Planned stories that will implement this Epic

### ✅ **Quality Standards:**

- **Problem-Focused**: Epic addresses a specific user problem from PRD
- **Measurable Goal**: Epic completion criteria are specific and testable
- **Clear Scope**: Boundaries are well-defined with included/excluded items
- **PRD Traceable**: Clear connection to specific BR/NFR requirements
- **Story-Ready**: Provides sufficient detail for Story creation

### ❌ **Common Pitfalls to Avoid:**

- Technical implementation details (leave for Stories and Architecture)
- Vague problem statements without user context
- Unmeasurable completion criteria
- Missing PRD traceability to requirements
- Scope too large for single Epic (consider Epic splitting)

### 🎯 **Story Enablement:**

This Epic should enable immediate Story creation by providing:

- **Clear problem context** that Stories can address with specific features
- **Defined user personas** that Stories can target with "As a user" scenarios
- **Measurable outcomes** that become Story acceptance criteria
- **Scope boundaries** that guide Story prioritization within the Epic
