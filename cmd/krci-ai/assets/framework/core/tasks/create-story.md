---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
    - prioritization-frameworks.md
  templates:
    - story.md
---

# Task: Create Story

## Description

Create user-focused requirements with implementation tasks and acceptance criteria that break down Epic features into actionable development work. This story provides specific user value and clear implementation guidance for development teams.

<prerequisites>
- Epic exists: Target Epic is defined and available in `/docs/epics/`
- Epic context understood: Epic problem, goal, and scope are clear
- User persona identified: Target user from Epic is specified
- Story scope defined: Specific functionality this Story will deliver
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-sdlc-framework.md
- ./.krci-ai/data/prioritization-frameworks.md
- ./.krci-ai/templates/story.md

CRITICAL: Load all dependencies by reading their complete content before task execution. HALT if any missing.

<instructions>
1. Follow SDLC workflow: Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for Story dependencies and naming conventions
2. Apply prioritization: Use methods from [prioritization-frameworks.md](./.krci-ai/data/prioritization-frameworks.md)
3. Format output: Use [story.md](./.krci-ai/templates/story.md) template with proper naming
4. Ensure Epic traceability: Reference parent Epic and map to specific Epic deliverables
5. Create comprehensive content: Generate rich, detailed stories with extensive technical context and strategic reasoning
6. Provide technical depth: Include detailed architectural background, implementation guidance, and quality considerations
7. Integrate strategic context: Explain the purpose, significance, and system integration aspects for all story elements
8. Template compliance: Follow story.md template structure and populate all variables correctly
</instructions>

## Output Format

- **Location**: `/docs/stories/{epic_number}.{story_number}.story.md` (EXACT naming pattern)
- **Story numbering**: Sequential number within Epic (e.g., 01.01, 01.02, 01.03 for Epic 1)
- **Epic reference**: Clear connection to parent Epic in format "Epic {number}: {title}"
- **Implementation Ready**: Story contains sufficient detail for autonomous development
- **Testing Ready**: Acceptance criteria provide clear validation steps for QA

<success_criteria>
- File saved to `/docs/stories/{epic_number}.{story_number}.story.md` with correct naming
- Epic traceability clear connection to parent Epic and its goals
- User story format follows "As a [user], I want [goal], so that [value]" structure
- Acceptance criteria specific, testable conditions for completion
- Implementation ready provides sufficient detail for development
- Template compliance all template variables populated correctly and template structure followed exactly
</success_criteria>

## Execution Checklist

<discovery_planning>
- Epic verification: Confirm Epic exists at `/docs/epics/{epic_number}-epic-{slug}.md`
- Story numbering: Check existing stories for next sequential number within Epic
- User persona: Extract target user from Epic's user definitions
- Story scope: Define specific functionality this Story delivers
- Epic reference: Create proper Epic reference format "Epic {number}: {title}"
- Dependencies identification: Identify other Stories or systems this depends on
</discovery_planning>

### Story Definition

- [ ] **User story creation**: Write "As a [user], I want [goal], so that [value]" aligned with Epic features
- [ ] **Story points estimation**: Estimate complexity (1, 2, 3, 5, 8, 13) using Epic context
- [ ] **Business value validation**: Ensure Story delivers measurable user value
- [ ] **Status format**: Use table format as defined in template (| Field | Value |)

### Requirements Specification

- [ ] **Acceptance criteria**: Define specific, testable conditions with measurable outcomes and file deliverables
- [ ] **Technical approach**: Define overall implementation approach and strategy with specific technologies
- [ ] **Architecture references**: Include direct links to specific architecture sections needed (format: `[filename](path)`)

<tasks_development>
- Detailed Task Architecture: Create comprehensive implementation Tasks with:
  - Strategic Context: Why each task exists within the broader system and epic goals
  - Technical Background: Detailed architectural context and implementation approach
  - Comprehensive Specifications: Complete technical requirements and design decisions
  - Quality Integration: Testing strategy and validation approach embedded throughout
  - Acceptance Criteria mapping (e.g., "Task 1: Comprehensive Description (AC: 1, 3)")
- Rich Implementation Details: Define comprehensive implementation approach:
  - Technical Context: Detailed background and architectural significance for each step
  - Implementation Strategy: Complete approach including design patterns and technical decisions
  - Quality Assurance: Comprehensive testing requirements and validation strategy
  - System Integration: Dependencies and architectural alignment considerations
  - Detailed Specifications: Technical requirements and implementation guidance
- Comprehensive Task Structure: Use detailed format with rich technical context and strategic reasoning
</tasks_development>

### Quality Assurance Planning

- [ ] **QA checklist**: Define testing requirements with clear verification steps/commands and expected outputs
- [ ] **Validation plan**: Create testing checklist with:
  - **Verification methods** LLMs can run (commands where applicable)
  - **Expected outputs** and success indicators
  - **Rollback steps** if testing fails

### Output Formatting

- [ ] **Story creation**: Use [story.md](./.krci-ai/templates/story.md) template structure
- [ ] **Variable population**: Complete all template variables ({{story_number}}, {{story_title}}, etc.)
- [ ] **Content validation**: Ensure user story, acceptance criteria, and Tasks/Subtasks are complete
- [ ] **File placement**: Save to exact location `/docs/stories/{epic_number}.{story_number}.story.md`
