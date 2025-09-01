# Task: Review Story (Project Manager)

## Description

Review and enhance user story against the story template to ensure completeness and implementation readiness. This task can be used by Product Owners, Developers, and Architects to collaboratively shape stories with appropriate level of detail. Focus on validation, clarification, and enhancement - NO IMPLEMENTATION.

## Prerequisites

- [ ] **Story exists**: Target story file exists in `/docs/stories/` requiring review and enhancement
- [ ] **Template available**: Story template at `./.krci-ai/templates/story.md` for reference
- [ ] **Role context**: Understanding appropriate to your role (PO: business clarity, Dev: technical details, Architect: system design)
- [ ] **Story knowledge**: Familiarity with story structure and requirements for implementation readiness

### Reference Assets

Dependencies:

- ./.krci-ai/templates/story.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Validate against template**: Check story against [story.md](./.krci-ai/templates/story.md) template structure
2. **Role-based review**: Assess story from your role perspective (PO: business clarity, Dev: technical details, Architect: design)
3. **Identify gaps**: Find missing details needed for implementation readiness
4. **Provide feedback**: Document what's unclear, missing, or needs enhancement
5. **Enhance story**: Add appropriate details while preserving business requirements and Epic alignment

### Template Compliance Check

Validate story against template sections:

1. **Status Table**: Status, Epic, Priority, Story Points properly filled
2. **Dependencies**: Technical dependencies clearly specified
3. **Story Format**: "As a/I want/so that" correctly structured
4. **Acceptance Criteria**: Testable, specific, with validation commands
5. **Description**: Clear context for implementation
6. **Tasks/Subtasks**: Specific, executable implementation steps
7. **Implementation Results**: Section ready for developer documentation
8. **QA Checklist**: Testing requirements with specific commands

## Multi-Role Review Focus

### STORY Completeness (All Roles)

**Template Section Validation:**

- **Status table complete** - Status, Epic, Priority, Story Points filled
- **Dependencies clear** - Technical and business dependencies specified
- **Story format correct** - "As a/I want/so that" properly structured
- **Acceptance criteria specific** - Testable and measurable conditions

### TECHNICAL Clarity (Developer/Architect Focus)

**Implementation Details:**

- **File paths specified** - Exact paths for files to create/modify
- **Libraries/versions defined** - Specific packages with versions
- **Commands specified** - Exact commands to execute
- **Architecture patterns** - Design approaches and system integration

### BUSINESS Clarity (Product Owner Focus)

**Business Requirements:**

- **User value clear** - Business value and user benefit obvious
- **Acceptance criteria testable** - Criteria can be validated
- **How to test it** - Clear from QA Checklist
- **No assumptions needed** - All technical decisions pre-defined

## Output Format

- **Location**: Update existing story file in `/docs/stories/` with enhanced technical details
- **Technical enhancement**: Comprehensive Tasks/Subtasks with implementation specifications
- **Implementation guidance**: All technical details required for autonomous development
- **Quality validation**: Enhanced QA Checklist with specific testing requirements

## Success Criteria

- [ ] **Story technically complete** - All technical implementation details specified
- [ ] **Tasks/Subtasks enhanced** - Atomic, executable steps with specific commands and file paths
- [ ] **Implementation autonomous** - Developer can implement without external consultations
- [ ] **Quality requirements defined** - Specific testing commands and validation steps included
- [ ] **Repository specifications clear** - Exact file paths, directory structures, and dependencies specified
- [ ] **Technical decisions documented** - Implementation approaches and patterns pre-defined

## Execution Checklist

### Template Validation

- [ ] **Check template sections**: Verify all 8 story template sections are present and populated
- [ ] **Status table review**: Confirm Status, Epic, Priority, Story Points are filled
- [ ] **Story format check**: Validate "As a/I want/so that" structure is correct
- [ ] **Dependencies clarity**: Ensure technical dependencies are specified

### Technical Clarity Assessment

- [ ] **File paths validation**: Check all file paths are exact and specific (e.g., `/path/to/file.ext`)
- [ ] **Library specifications**: Verify libraries include versions (e.g., `gopkg.in/yaml.v3 v3.0.1`)
- [ ] **Command verification**: Ensure commands are specific and executable
- [ ] **Project structure check**: Validate directory structure is clearly defined

### Tasks/Subtasks Review

- [ ] **Atomic verification**: Confirm each subtask is single, executable action
- [ ] **Command completeness**: Check each subtask has specific commands
- [ ] **Validation steps**: Ensure verification commands are included
- [ ] **Dependency mapping**: Verify task dependencies are clear

### Implementation Readiness

- [ ] **Clarity test**: Confirm developer can implement without assumptions
- [ ] **Gap identification**: Document any missing technical details
- [ ] **QA completeness**: Verify QA checklist has specific commands and expected outputs
- [ ] **Feedback documentation**: Record any areas needing clarification

## Review Guidelines (Multi-Role)

### Validation Rules

- **Check template compliance** - Verify all template sections are complete
- **Role-appropriate review** - Focus on your domain expertise while considering other perspectives
- **Enhance clarity** - Add details that improve implementation readiness
- **Preserve requirements** - Don't change business requirements or Epic alignment

### Required Technical Specifications

**Tasks/Subtasks Must Include:**

- **Exact file paths** - `/path/to/file.ext` for all deliverables
- **Specific commands** - Executable commands (e.g., `go mod tidy`, `pytest tests/`)
- **Library versions** - Package names with versions (e.g., `gopkg.in/yaml.v3 v3.0.1`)
- **Validation steps** - How to verify each deliverable

**QA Checklist Must Include:**

- **Testing commands** - Specific commands with expected outputs
- **Performance criteria** - Measurable success metrics
- **Verification steps** - How to validate story completion

### Feedback Documentation

**If Technical Details Missing:**

- **Document gaps** - List what's unclear or missing
- **Suggest specifics** - Propose exact file paths, commands, dependencies
- **Ensure clarity** - Make implementation autonomous and unambiguous
