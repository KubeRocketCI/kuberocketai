# Task: create-epic

**Description:**
Create clear epic with problem statement, goal, scope, and implementation approach

## Prerequisites

- Clear understanding of what needs to be built
- Basic user context identified

## Output Location

Create epic files in: `{project_root}/docs/epics/{epic_number}-epic-{slug}.md`

## Process

### Phase 1: Epic Foundation

1. **Assign Epic Identity**
   - Generate sequential epic number (check existing epics for next number)
   - Create clear, descriptive title
   - Generate descriptive slug from title (lowercase, hyphenated)

2. **Define the Problem and Goal**
   - Articulate what problem this epic solves
   - Define clear goal for what we want to achieve
   - Identify target users who will benefit

### Phase 2: Define Scope

3. **Establish What's Included**
   - List main features and functionality in scope
   - Be specific about what will be built

4. **Define What's Not Included**
   - Explicitly state what is out of scope
   - Identify future considerations

5. **Identify Dependencies**
   - List other epics, features, or systems this depends on
   - Mark as "None" if independent

### Phase 3: Solution and Implementation

6. **Describe Solution Approach**
   - Outline high-level approach to solving the problem
   - Keep it simple and clear

7. **Create Acceptance Criteria**
   - Define when this epic will be considered complete
   - Make criteria testable and clear

8. **Plan User Stories**
   - Identify main user stories that will make up this epic
   - Keep high-level, detailed stories come later

### Phase 4: Implementation Planning

9. **Create Implementation Plan**
   - Break down into main phases or steps
   - Include key milestones
   - Keep it simple and actionable

10. **Set Epic Status**
    - Set initial status (Draft/Planning/In Progress/Done)
    - Assign priority (Critical/High/Medium/Low)
    - Set owner and rough timeline

## Output Format

Use the epic template (`{project_root}/.krci-ai/templates/epic.md`) and populate ALL variables:

- `{{epic_number}}`, `{{epic_title}}`, `{{status}}`, `{{priority}}`
- `{{owner}}`, `{{timeline}}`
- `{{problem_statement}}`, `{{goal}}`, `{{target_users}}`
- `{{in_scope}}`, `{{out_of_scope}}`, `{{dependencies}}`
- `{{solution_approach}}`, `{{acceptance_criteria}}`
- `{{planned_stories}}`, `{{implementation_plan}}`
- `{{implementation_results}}`

## Success Criteria

- Epic clearly describes the problem and solution
- Scope is well-defined with clear boundaries
- Implementation approach is practical and actionable
- Acceptance criteria are clear and testable
- User stories are identified at high level
- Epic file created with descriptive name: `{project_root}/docs/epics/{epic_number}-epic-{slug}.md`
