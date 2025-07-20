# Task: implement-feature

**Description:**
Implement user stories by following a structured development process with clear breakdown, user approval, and status tracking

## Prerequisites

- Access to project codebase and development environment
- Understanding of story acceptance criteria and implementation plan
- Development tools and dependencies configured
- Testing framework available

## Story Selection (HALT AFTER EACH STEP)

1. **Check User Input**: Determine if user has provided a specific story number/name
2. **Direct Story Selection** (if story provided by user):
   - Locate the story file using provided story number/name
   - Skip to Story Validation step
3. **Story Discovery** (if no story specified):
   - Check `{project_root}/docs/stories/` folder for existing stories (format: `{epic_number}.{story_number}-story-{slug}.md`)
   - List all available stories with their titles, status, and priority
   - Display stories in organized format: `{epic_number}.{story_number}: {title} [{status}] - {priority}`
   - Group by epic for better organization
   - **HALT**: Present story list to user and wait for selection
4. **Story Validation**: 
   - Verify story exists and is readable
   - Check story status (should be "Ready for Development" or "Draft")
   - Ensure story has clear acceptance criteria and implementation plan
5. **HALT for Confirmation**: Present selected story details and ask user to confirm implementation should proceed

## Story Location

Target story files in: `{project_root}/docs/stories/{epic_number}.{story_number}-story-{slug}.md`

## Process (EXECUTE ONLY AFTER STORY CONFIRMATION)

### Phase 1: Story Preparation

1. **Update Story Status**
   - Change story status from current state to "In Progress"
   - Update story file with new status
   - Add implementation start timestamp

2. **Analyze Story Requirements**
   - Review acceptance criteria thoroughly
   - Identify all functional and non-functional requirements
   - Parse existing implementation plan from story
   - Note any dependencies or prerequisites

### Phase 2: Implementation Planning

3. **Break Down Implementation**
   - Decompose story into small, manageable development tasks
   - Each task should be completable in 1-2 hours
   - Order tasks by logical dependency and priority
   - Include specific file changes, functions to create/modify, and testing steps

4. **Create Detailed Implementation Plan**
   - List all files to be created, modified, or deleted
   - Specify code changes with function/method signatures
   - Include database schema changes if applicable
   - Define integration points and API changes
   - Plan testing approach for each component

5. **HALT for User Approval**: Present complete implementation plan to user and wait for approval before proceeding

### Phase 3: Implementation Execution

6. **Development Loop** (Iterate until story complete):
   - **Task Selection**: Pick next task from implementation plan
   - **Implementation**: Write/modify code following coding standards
   - **Unit Testing**: Create/update unit tests for new functionality
   - **Integration Testing**: Test integration with existing components
   - **Documentation**: Update inline documentation and README if needed
   - **Validation**: Verify task meets acceptance criteria
   - **Status Update**: Mark task as complete and move to next

7. **Code Quality Assurance**
   - Run linting and formatting tools
   - Ensure all tests pass
   - Check code coverage meets requirements
   - Verify no security vulnerabilities introduced

### Phase 4: Story Completion

8. **Final Validation**
   - Test all acceptance criteria are met
   - Perform end-to-end testing scenarios
   - Verify integration with existing features
   - Check performance and security requirements

9. **Update Story Status**
   - Change story status to "Done"
   - Update implementation results section with:
     - Summary of changes made
     - Files modified/created
     - Testing results
     - Any technical debt or future improvements noted
   - Add completion timestamp

10. **Documentation and Cleanup**
    - Update relevant documentation
    - Clean up temporary files or debug code
    - Prepare summary of implementation for stakeholders

## Implementation Standards

- Follow project coding standards and conventions
- Write comprehensive unit tests (minimum 80% coverage)
- Update documentation for new features
- Ensure backward compatibility unless explicitly breaking change
- Handle error cases and edge conditions
- Implement proper logging and monitoring

## Success Criteria

- Story status updated from initial state to "In Progress" then "Done"
- All acceptance criteria validated and met
- Implementation plan approved by user before execution
- Code changes follow project standards and pass all tests
- Story file updated with detailed implementation results
- No regression in existing functionality
- All tasks in implementation plan completed successfully
- Final implementation matches approved plan or deviations are documented

## Error Handling

- If story file not found, provide helpful error message and list available stories
- If story status is inappropriate (e.g., "Done"), ask user to confirm re-implementation
- If implementation fails, revert changes and update story with failure details
- If user rejects implementation plan, return to planning phase for revision 