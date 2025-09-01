# Story {{story_number}}: {{story_title}}

<instructions>
STORY STRUCTURE: Follow this exact section ordering:
1. Status → 2. Dependencies → 3. Story → 4. Acceptance Criteria → 5. Description → 6. Technical Context → 7. Tasks/Subtasks → 8. Implementation Results → 9. QA Checklist
</instructions>

## Status

<status>
| Field                  | Value                       |
|------------------------|-----------------------------|
| Status                 | {{status}}                  |
| Epic                   | {{epic_reference}}          |
| Priority               | {{priority}}                |
| Estimated Story Points | {{story_points}}            |
| Jira                   | {{jira_ticket}}             |
</status>

<instructions>
Status tracking and Epic/PRD traceability. Enables progress monitoring and dependency validation.

Status Options: Draft -> Approved -> In Progress -> Done -> Completed
Epic Reference Example: Epic 1 - Unified Agent Activation
Priority Example: Critical, High, Medium, Low
Story Points Example: 1, 2, 3, 5, 8, 13 (Fibonacci scale)
Jira Example: [EPMDEDP-15497](https://jira.example.com/browse/EPMDEDP-15497) or "None" if not assigned

CRITICAL: Status section contains ONLY these 5 fields. Do NOT add Dependencies or other fields here.
</instructions>

## Dependencies

<dependencies>
**Blocking:**
{{blocking_dependencies}}

**Blocked By:**
{{blocked_by_dependencies}}

**System/Test Dependencies:**
{{system_test_dependencies}}

<instructions>
Define precise dependencies for execution order and validation readiness.

CRITICAL: Dependencies section comes immediately after Status section.

Format:

- Blocking: Items that depend on this story
- Blocked By: Items this story depends on (stories, approvals)
- System/Test: Environments, versions, stubs/mocks, fixtures

Examples:

- Blocking: Story 02.03 (depends on config file produced here)
- Blocked By: Story 02.01 (API contract), Security approval
- System/Test: Local IDP stub v1.2; Env var FEATURE_FLAG=on

DO:

- Use exact story numbers/links and system versions
- Include test doubles (stubs/mocks) to avoid external blockers
- State "None" explicitly if empty

DON'T:

- List vague dependencies (e.g., "backend work")
- Omit versions/links where relevant
- Depend on production-only services for acceptance
- Put dependencies in Status section
</instructions>

</dependencies>

## Story

<user_story>
**As a** {{persona}},
**I want** {{goal}},
**so that** {{business_value}}.

<instructions>
Standard story format focusing on persona, goal, and value. Must align with Epic Target Users and provide specific value.

Story Example:
As a Software Engineer,
I want a unified agent activation command in the IDE,
so that I can start using an agent without extra setup.

DO:

- Use persona from Epic (PRD-defined), not generic labels
- Make the goal specific and implementation-agnostic
- State tangible business/user value

DON'T:

- Invent personas not defined in PRD/Epic
- Use vague goals like "improve experience"
- Include solution details here
</instructions>

</user_story>

## Acceptance Criteria

<acceptance_criteria>
{{acceptance_criteria}}

<instructions>
CRITICAL: Acceptance Criteria section comes immediately after Story section. Create specific, testable conditions that define completion at story level. Story ACs SHOULD include verification commands and expected outputs.

Required Structure for each criterion:

1. Scenario (Given/When/Then): Brief user/system flow to validate
2. Expected Behavior: Clear, observable outcome
3. Verification Command: Non-interactive command(s) with expected exit code/output
4. Files Created/Changed: Exact paths
5. Guardrails: NFR constraints applicable to this story (e.g., perf, security)
6. Test Data/Fixtures: Test users, payloads, seeds
7. Environment/Flags: Env vars, feature flags/toggles, mock/stub switches
8. Evidence: Required artifacts (logs, screenshots, run output)
9. Non-automatable at story-level: true/false (and rationale if true)
10. Traceability: Epic AC id(s), PRD BR/NFR
11. Out of Scope (optional): Clarify non-goals

Acceptance Criteria Example:

1. **OAuth Integration**: System successfully connects to OAuth service and processes authentication requests
2. **Token Processing**: Validate, decode, and process JWT tokens with proper error handling
3. **Session Management**: Implement session timeout, auto-renewal, and security mechanisms
4. **API Endpoints**: Develop comprehensive authentication endpoints for complete workflow
5. **Testing & Validation**: Implement comprehensive test suite with security validation and coverage requirements

Format Guidelines:

- Simple numbered list with bold titles and clear descriptions
- One line per acceptance criterion
- Focus on specific, testable outcomes
- Avoid complex sub-bullets or excessive detail
- Keep descriptions concise but complete

DO:

- Use simple, clear numbered format
- Focus on specific, testable outcomes
- Keep descriptions concise and actionable
- Map each AC to Epic requirements

DON'T:

- Add complex sub-bullet structures
- Use excessive technical detail in AC descriptions
- Create overly long or complex acceptance criteria
- Mix implementation details with acceptance criteria
</instructions>

</acceptance_criteria>

## Description

<description>
{{description}}

<instructions>
Comprehensive context for strategic understanding and implementation guidance. Provide detailed background, technical context, and architectural alignment.

Comprehensive Context Requirements:

- Strategic purpose within the Epic and broader system goals
- Detailed technical background and architectural significance
- Implementation philosophy and approach rationale
- System integration considerations and dependencies
- Quality and architectural compliance requirements
- Technical constraints and design decision context
</instructions>

</description>

## Technical Context

<technical_context>
{{technical_context}}

<instructions>
Detailed technical background and architectural considerations. Provide comprehensive implementation context and design guidance.

Technical Context Content:

- Architectural significance and system integration points
- Technical approach and design pattern rationale
- Implementation constraints and technical dependencies
- Quality considerations and testing strategy overview
- System design principles and architectural alignment
- Performance, security, and scalability considerations
</instructions>

</technical_context>

## Tasks/Subtasks

<tasks_subtasks>
{{tasks_subtasks}}

<instructions>
Create LLM-executable implementation plan with atomic tasks and validation. Each task maps to acceptance criteria with specific commands and file paths.

TASKS/SUBTASKS STRUCTURE REQUIREMENTS:

**Task N: Task Description (AC: X, Y)**

- [ ] Clear, actionable implementation step with specific deliverable
- [ ] Another implementation step focused on technical execution
- [ ] Testing and validation step with clear completion criteria
- [ ] Integration and deployment verification step

STRUCTURE REQUIREMENTS:

- Simple task titles with acceptance criteria mapping (AC: X, Y)
- Clean bullet checkboxes for implementation tracking
- Focused on actionable implementation steps
- Clear, concise descriptions without excessive formatting
- Each checkbox represents a specific, completable work item
- Map tasks to acceptance criteria clearly
- Include testing and validation steps
- Avoid excessive formatting or complex structure
- Do not mix strategic planning with implementation checkboxes
- Essential implementation checkboxes are mandatory
</instructions>

</tasks_subtasks>

## Implementation Results

<implementation_results>
{{implementation_results}}

<instructions>
Concrete outcomes populated AFTER completion (evidence-first).

Include:

- Evidence per AC (links to logs, screenshots, artifacts)
- Files created/updated with paths
- Commands executed with outputs
- Final validation summary

DO:

- Use past tense (Created, Implemented, Validated)
- Link evidence to specific AC ids
- Include actual command results/exit codes

DON'T:

- Populate before implementation
- Use future/planning language
- Omit evidence for any AC
</instructions>

</implementation_results>

## QA Checklist

<qa_checklist>
{{qa_checklist}}

<instructions>
Create specific verification steps with commands and expected outputs. Enable automated testing and quality validation.

QA Checklist Example:

### Functional

- [ ] Schema Validation: `python hack/validate-agents.py` (Expect: exit 0)
- [ ] File Existence: `ls -la assets/agents/architect.yaml` (Expect: file exists)
- [ ] Content Validation: `grep -q "identity" assets/agents/architect.yaml` (Expect: pattern found)

### Integration

- [ ] IDE Testing: Activate agent in Cursor (Expect: response within 5s)
- [ ] Cross-Platform: Validate on macOS/Linux/Windows (Expect: consistent behavior)

### Security & Privacy

- [ ] Secrets: No tokens/secrets in logs
- [ ] Auth: No P0/P1 security findings in changed scope

### Accessibility (if UI)

- [ ] Keyboard navigation and focus order
- [ ] Contrast ratio meets baseline

DO:

- Group tests by category; include expected outputs
- Keep checks non-interactive where possible
- Align with AC guardrails

DON'T:

- Use subjective testing criteria
- Omit expected outputs or success indicators
- Depend on production-only services
</instructions>

</qa_checklist>
