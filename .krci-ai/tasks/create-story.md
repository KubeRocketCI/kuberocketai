# Task: Create Story

## Description

Create user-focused requirements with implementation tasks and acceptance criteria that break down Epic features into actionable development work. This story provides specific user value and clear implementation guidance for development teams.

## Framework Context

**Reference**: [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) - See role responsibilities and artifact flow

## Prerequisites

- [ ] **Epic exists**: Target Epic is defined and available in `/docs/epics/`
- [ ] **Epic context understood**: Epic problem, goal, and scope are clear
- [ ] **User persona identified**: Target user from Epic is specified
- [ ] **Story scope defined**: Specific functionality this Story will deliver

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for Story dependencies and naming conventions
2. **Apply prioritization**: Use methods from [prioritization-frameworks.md](./.krci-ai/data/prioritization-frameworks.md)
3. **Format output**: Use [story.md](./.krci-ai/templates/story.md) template with proper naming
4. **Ensure Epic traceability**: Reference parent Epic and map to specific Epic deliverables

## Output Format

- **Location**: `/docs/stories/{epic_number}.{story_number}.story.md` (EXACT naming pattern)
- **Story numbering**: Sequential number within Epic (e.g., 1.1, 1.2, 1.3 for Epic 1)
- **Epic reference**: Clear connection to parent Epic in format "Epic {number}: {title}"
- **Downstream Enable**: Enables Code implementation and Testing

## Success Criteria

- [ ] **File saved** to `/docs/stories/{epic_number}.{story_number}.story.md` with correct naming
- [ ] **Epic traceability** clear connection to parent Epic and its goals
- [ ] **User story format** follows "As a [user], I want [goal], so that [value]" structure
- [ ] **Acceptance criteria** specific, testable conditions for completion
- [ ] **Implementation ready** provides sufficient detail for development
- [ ] **Template compliance** all template variables populated correctly

## Execution Checklist

### Discovery Phase

- [ ] **Epic selection**: Identify target Epic from `/docs/epics/` folder
- [ ] **Story numbering**: Check existing stories for next sequential number within Epic
- [ ] **User persona**: Extract target user from Epic's user definitions
- [ ] **Story scope**: Define specific functionality this Story delivers

### Planning Phase

- [ ] **User story definition**: Create "As a [user], I want [goal], so that [value]" statement
- [ ] **Epic reference**: Create proper Epic reference format "Epic {number}: {title}"
- [ ] **Story points estimation**: Estimate complexity (1, 2, 3, 5, 8, 13) using Epic context
- [ ] **Dependencies identification**: Identify other Stories or systems this depends on

### Requirements Phase

- [ ] **Acceptance criteria**: Define specific, testable conditions for completion
- [ ] **Implementation plan**: Outline key development tasks and approach
- [ ] **QA checklist**: Define testing requirements and validation steps
- [ ] **Business value validation**: Ensure Story delivers measurable user value

### Documentation Phase

- [ ] **Story creation**: Use [story.md](./.krci-ai/templates/story.md) template structure
- [ ] **Variable population**: Complete all template variables ({{story_number}}, {{story_title}}, etc.)
- [ ] **Content validation**: Ensure user story, acceptance criteria, and plan are complete
- [ ] **File placement**: Save to exact location `/docs/stories/{epic_number}.{story_number}.story.md`

### Epic Integration Phase

- [ ] **Epic verification**: Confirm Epic exists at `/docs/epics/{epic_number}-epic-{slug}.md`
- [ ] **User story creation**: Write user story format "As a... I want... so that..." aligned with Epic features
- [ ] **Acceptance criteria**: Define specific, testable success criteria for Story validation
- [ ] **Tasks/Subtasks plan**: Outline key development tasks and approach as checklist structure

### Story Completion Phase

- [ ] **Story file creation**: Save Story to exact location `/docs/stories/{epic_number}.{story_number}.story.md`
- [ ] **Template compliance**: Use [story.md](./.krci-ai/templates/story.md) template with all required sections
- [ ] **Content validation**: Ensure user story, acceptance criteria, and Tasks/Subtasks are complete

### Story Creation Phase

- [ ] **Story structure**: Create Story using [story.md](./.krci-ai/templates/story.md) template with structured Tasks/Subtasks section
- [ ] **User story definition**: Write clear "As a... I want... so that..." format with persona, goal, and business value
- [ ] **Acceptance criteria**: Define specific, testable criteria that validate Story completion
- [ ] **Tasks/Subtasks planning**: Create checklist structure for systematic LLM execution

### Tasks/Subtasks Development Phase

- [ ] **Technical approach**: Define overall implementation approach and strategy
- [ ] **Architecture references**: Include direct links to specific architecture sections needed
- [ ] **Task breakdown**: Create main implementation Tasks with descriptions and file lists
- [ ] **Subtask checklists**: Define detailed Subtasks as checkboxes for systematic LLM execution
- [ ] **Testing plan**: Create testing checklist items to validate Story functionality

## Content Guidelines

### 📋 **Story Template Sections:**

1. **Status Table**: Story number, title, status, Epic reference, priority, story points, Jira ticket
2. **Dependencies**: Other Stories or systems this Story depends on
3. **Story**: "As a [user], I want [goal], so that [value]" format
4. **Acceptance Criteria**: Specific, testable conditions for completion
5. **Description**: Detailed context and background for the Story
6. **Implementation Plan**: Development approach and key tasks
7. **Implementation Results**: Outcomes and deliverables (populated during development)
8. **QA Checklist**: Testing requirements and validation steps

### ✅ **Quality Standards:**

- **User-Centered**: Story focuses on specific user value and outcomes
- **INVEST Principles**: Independent, Negotiable, Valuable, Estimable, Small, Testable
- **Epic Traceable**: Clear connection to parent Epic goals and scope
- **Implementation Ready**: Sufficient detail for development without being prescriptive
- **Testable**: Acceptance criteria are specific and verifiable

### ❌ **Common Pitfalls to Avoid:**

- Technical implementation details (leave flexibility for development)
- Vague acceptance criteria that can't be tested
- Stories too large for single sprint (consider Story splitting)
- Missing Epic traceability and context
- User stories that don't deliver measurable value

### 🎯 **Implementation Enablement:**

This Story should enable immediate development by providing:

- **Clear user value** that developers can understand and implement
- **Specific acceptance criteria** that define completion
- **Implementation guidance** without being overly prescriptive
- **Testing requirements** that enable quality validation

### ✅ **Story Structure Elements:**

1. **Status**: Status tracking and Epic reference
2. **Dependencies**: Inter-story and external dependencies
3. **Story**: User story in standard format
4. **Acceptance Criteria**: Specific validation requirements
5. **Description**: Additional context and details
6. **Tasks/Subtasks**: Development checklist and approach
