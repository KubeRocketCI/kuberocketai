# Task: create-story

**Description:**
Create INVEST-compliant user story with clear acceptance criteria, implementation plan, and comprehensive QA checklist

## Prerequisites

- Epic context and requirements
- Target user persona identified
- Business value proposition clear

## Epic Selection (HALT AFTER EACH STEP)

1. **Review Available Epics**: Check `{project_root}/docs/epics/` folder for existing epics (format: `{epic_number}-epic-{slug}.md`)
2. **Epic Discovery** (if no epics available or unclear):
   - Ask user to describe the broader feature/initiative this story supports
   - Identify the business goal and user value theme
   - Determine if this creates a new epic or fits existing product roadmap
   - Suggest creating an epic first if none exists
3. **Propose Target Epic**: Identify which epic this story belongs to based on requirements and discovery
4. **Determine Epic Number**: Extract epic number from chosen epic filename (e.g., from `1-epic-ide-integration.md` get `1`)
5. **HALT for Confirmation**: Present epic proposal and epic number to user and wait for confirmation before proceeding

## Output Location

Create story files in: `{project_root}/docs/stories/{epic_number}.{story_number}-story-{slug}.md`

## Process (EXECUTE ONLY AFTER EPIC CONFIRMATION)

### Phase 1: Story Foundation

1. **Assign Story Identity**
   - Generate sequential story number within epic (check existing stories for next number)
   - Create clear, descriptive title following pattern: "{Action} {Feature} for {User}"
   - Generate descriptive slug from title (lowercase, hyphenated)

2. **Define User Story Core**
   - Identify target persona from epic's target users
   - Articulate user goal (what they want to achieve)
   - Define business value (why it matters)
   - Format as: "As a {persona}, I want {goal}, so that {business_value}"

### Phase 2: Story Requirements

3. **Establish Story Context**
   - Set initial status (Draft/Ready for Development/In Progress)
   - Create epic reference in format: "Epic {epic_number}: {epic_title}"
   - Assign priority level (Critical/High/Medium/Low)
   - Estimate story points (1, 2, 3, 5, 8, 13) using complexity assessment

4. **Define Dependencies**
   - List all prerequisite stories, features, or technical components
   - Mark as "None" if independent

### Phase 3: Story Details

5. **Create Acceptance Criteria**
   - Write specific, testable conditions using Given-When-Then format where applicable
   - Ensure each criterion is measurable and covers happy path and edge cases

6. **Develop Detailed Implementation Plan**
   - Outline technical approach and architecture decisions
   - List development steps in logical order
   - Identify key components to modify/create
   - Define integration points with existing systems
   - Include code structure and key methods/functions

### Phase 4: Quality Assurance

7. **Design QA Checklist**
   - Include functional testing scenarios
   - Add non-functional requirements (performance, security, usability)
   - Specify test data requirements and validation criteria

8. **Complete Story Metadata**
   - Add Jira ticket reference (or mark "TBD")
   - Write comprehensive description with context and rationale
   - Set implementation results placeholder

## Output Format

Use the story template (`{project_root}/.krci-ai/templates/story.md`) and populate ALL variables:

- `{{story_number}}`, `{{story_title}}`, `{{status}}`
- `{{epic_reference}}`, `{{priority}}`, `{{story_points}}`, `{{jira_ticket}}`
- `{{dependencies}}`
- `{{persona}}`, `{{goal}}`, `{{business_value}}`
- `{{acceptance_criteria}}`, `{{description}}`
- `{{implementation_plan}}`, `{{implementation_results}}`
- `{{qa_checklist}}`

## Success Criteria

- Story follows INVEST principles (Independent, Negotiable, Valuable, Estimable, Small, Testable)
- All template variables populated with meaningful content
- Acceptance criteria are clear and testable
- Implementation plan provides detailed development guidance
- QA checklist ensures comprehensive validation
- Story file created with descriptive name: `{project_root}/docs/stories/{epic_number}.{story_number}-story-{slug}.md`
- Epic properly referenced as "Epic {epic_number}: {epic_title}"
