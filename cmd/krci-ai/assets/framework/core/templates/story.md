# Story {{story_number}}: {{story_title}}

## Status

| Field                  | Value                       |
|------------------------|-----------------------------|
| Status                 | {{status}}                  |
| Epic                   | {{epic_reference}}          |
| Priority               | {{priority}}                |
| Estimated Story Points | {{story_points}}            |
| Jira                   | {{jira_ticket}}             |

<!-- Status tracking and Epic traceability -->
<!-- Enables progress monitoring and Epic dependency validation -->

<!-- Template Guidance:
Status Options: Draft -> Approved -> In Progress -> Done -> Completed
Epic Reference Example: "Epic 1 - KubeRocketAI Baseline"
Priority Example: Critical, High, Medium, Low
Story Points Example: 1, 2, 3, 5, 8, 13 (Fibonacci scale - estimate complexity)
Jira Example: "[EPMDEDP-15497](https://jira.example.com/browse/EPMDEDP-15497)"
-->

## Dependencies

**Story Dependencies:**
{{dependencies}}

<!-- Define what this story depends on and what depends on it -->
<!-- Critical for LLM execution order and validation -->

<!-- Template Guidance:
Dependencies Example:
- "Story 01.01: Architect Agent Definition (foundation patterns)"
- "System: Python validation tooling infrastructure"
- "None" if independent story

Format: Use "Story XX.YY: Brief description" or "System: Description"
✅ DO: List specific story numbers and brief context
✅ DO: Include external system dependencies
❌ DON'T: List vague dependencies like "architecture work"
-->

## Story

**As a** {{persona}},
**I want** {{goal}},
**so that** {{business_value}}.

<!-- Standard user story format focusing on persona, goal, and business value -->
<!-- Must align with Epic's target users and provide specific value -->

<!-- Template Guidance:
Story Example:
"As a Software Architect,
I want a complete architect agent framework with system design capabilities,
so that I can immediately use KubeRocketAI for comprehensive architectural analysis."

Persona Example: Use specific persona from Epic (e.g., "Software Architect", "Development Lead", "QA Engineer")
Goal Example: Specific, measurable capability the user wants
Business Value Example: Clear outcome or benefit that explains the "why"

✅ DO: Use persona from Epic's target users
✅ DO: Make goal specific and measurable
✅ DO: Explain clear business value
❌ DON'T: Use generic roles like "user" or "developer"
❌ DON'T: Make goals vague like "improve experience"
-->

## Acceptance Criteria

{{acceptance_criteria}}

<!-- Specific, testable conditions that define completion -->
<!-- Must include file deliverables and verification commands for LLM validation -->

<!-- Template Guidance:
Acceptance Criteria Example:
1. File `assets/agents/architect.yaml` exists and passes schema validation
2. Agent responds to queries within 5 seconds in IDE testing
3. Validation command `python hack/validate-agents.py` returns exit code 0
4. Agent generates output following template format with required variables

Format Structure:
- Use numbered list for clear tracking
- Include specific file paths and expected outputs
- Add verification commands LLMs can execute
- Define measurable success criteria

✅ DO: Include file deliverables with exact paths
✅ DO: Add verification commands and expected results
✅ DO: Make criteria testable and measurable
✅ DO: Reference Epic's BR/NFR requirements where applicable
❌ DON'T: Use subjective criteria like "works well"
❌ DON'T: Omit verification steps for LLM validation
-->

## Description

{{description}}

<!-- Context explaining why this story exists and its strategic importance -->
<!-- Should provide background for LLM understanding and Epic alignment -->

<!-- Template Guidance:
Description Example:
"This foundational story establishes the architect agent framework as the first complete implementation of the KubeRocketAI four-asset-type model. The architect agent serves as the pattern-setter for subsequent agent development, demonstrating progressive complexity from Level 1 through Level 4."

Content Focus:
- WHY this story exists within the Epic
- Strategic importance and business context
- Relationships to other stories and architectural decisions
- Implementation philosophy or approach rationale

✅ DO: Explain the strategic importance within Epic context
✅ DO: Include relationships to other stories
✅ DO: Provide context for implementation decisions
❌ DON'T: Repeat acceptance criteria or implementation details
❌ DON'T: Include generic descriptions
-->

## Tasks/Subtasks

{{tasks_subtasks}}

<!-- LLM-executable implementation plan with atomic tasks and validation -->
<!-- Each task maps to acceptance criteria with specific commands and file paths -->

<!-- Template Guidance:
Tasks/Subtasks Structure Example:

- [ ] Task 1: Create Agent Foundation (AC: 1, 2, 3)
  - [ ] Create file: `assets/agents/architect.yaml` with identity structure
  - [ ] Validate schema: `python hack/validate-agents.py`
  - [ ] Test IDE integration: Confirm @agent functionality in Cursor
- [ ] Task 2: Implementation Validation (AC: 4)
  - [ ] Run verification: `grep -q "expected_pattern" output.md`
  - [ ] Document results: Update Implementation Results section

Task Structure Guidelines:
- Reference acceptance criteria: (AC: X, Y, Z)
- Atomic subtasks: One action per checkbox
- Specific commands: Include exact file paths and commands
- Validation steps: Add verification for each deliverable
- Dependencies: Note task order requirements

✅ DO: Structure as main Tasks with atomic Subtasks
✅ DO: Include specific file paths and commands
✅ DO: Add validation steps for each deliverable
✅ DO: Reference acceptance criteria numbers
✅ DO: Use checkbox format for progress tracking
❌ DON'T: Create tasks that require human interpretation
❌ DON'T: Omit validation or verification steps
❌ DON'T: Use vague action words like "handle" or "manage"

Command Patterns for LLMs:
- File creation: `create file: path/to/file.ext`
- File editing: `edit file: path/to/file.ext`
- Command execution: `run: command with args`
- Validation: `verify: validation_command`
-->

## Implementation Results

{{implementation_results}}

<!-- Concrete outcomes and deliverables populated AFTER story completion -->
<!-- Documents actual files created, commands executed, and validation results -->

<!-- Template Guidance:
Implementation Results Example:

### Completed Deliverables
**Agent Infrastructure:**
- ✅ `architect.yaml` agent implemented with identity structure
- ✅ Schema validation passing: `python hack/validate-agents.py`
- ✅ IDE integration tested: Cursor @agent functionality confirmed

**Validation Results:**
- ✅ All acceptance criteria verified and passing
- ✅ Files created at expected locations
- ✅ Commands executed successfully with expected outputs

Content Guidelines:
- Use past tense: "Created", "Implemented", "Validated"
- Include actual file paths and command results
- Document validation outcomes
- Group by logical categories

✅ DO: Populate AFTER completion with concrete outcomes
✅ DO: Include actual file paths and verification results
✅ DO: Use past tense for completed actions
✅ DO: Group results by logical categories
❌ DON'T: Populate before implementation
❌ DON'T: Use future tense or planning language
-->

## QA Checklist

{{qa_checklist}}

<!-- Specific verification steps with commands and expected outputs -->
<!-- Enables automated testing and quality validation -->

<!-- Template Guidance:
QA Checklist Example:

### Functional Testing
- [ ] **Schema Validation**: Run `python hack/validate-agents.py` - Expected: exit code 0
- [ ] **File Existence**: Verify `ls -la assets/agents/architect.yaml` - Expected: file exists
- [ ] **Content Validation**: Check `grep -q "identity" architect.yaml` - Expected: pattern found

### Integration Testing
- [ ] **IDE Testing**: Test agent in Cursor @agent - Expected: responds within 5 seconds
- [ ] **Cross-platform**: Validate on macOS, Linux, Windows - Expected: consistent behavior

Format Structure:
- Group by testing type (Functional, Integration, Performance)
- Include specific commands with expected outputs
- Add verification steps and success indicators
- Enable automated validation where possible

✅ DO: Include specific verification commands with expected outputs
✅ DO: Group by testing categories
✅ DO: Add measurable success criteria
✅ DO: Enable automated validation
❌ DON'T: Use subjective testing criteria
❌ DON'T: Omit expected outputs or success indicators
-->
