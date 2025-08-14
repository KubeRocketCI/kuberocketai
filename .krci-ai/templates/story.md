# Story {{story_number}}: {{story_title}}

## Status

| Field                  | Value                       |
|------------------------|-----------------------------|
| Status                 | {{status}}                  |
| Epic                   | {{epic_reference}}          |
| Priority               | {{priority}}                |
| Estimated Story Points | {{story_points}}            |
| Jira                   | {{jira_ticket}}             |

<!-- Status tracking and Epic/PRD traceability -->
<!-- Enables progress monitoring and dependency validation -->

<!-- Template Guidance:
Status Options: Draft -> Approved -> In Progress -> Done -> Completed
Epic Reference Example: "Epic 1 - Unified Agent Activation"
Priority Example: Critical, High, Medium, Low
Story Points Example: 1, 2, 3, 5, 8, 13 (Fibonacci scale)
Jira Example: "[EPMDEDP-15497](https://jira.example.com/browse/EPMDEDP-15497)"
 -->
## Dependencies

**Blocking:**
{{blocking_dependencies}}

**Blocked By:**
{{blocked_by_dependencies}}

**System/Test Dependencies:**
{{system_test_dependencies}}

<!-- Define precise dependencies for execution order and validation readiness -->

<!-- Template Guidance:
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
-->

## Story

**As a** {{persona}},
**I want** {{goal}},
**so that** {{business_value}}.

<!-- Standard story format focusing on persona, goal, and value -->
<!-- Must align with Epic Target Users and provide specific value -->

<!-- Template Guidance:
Story Example:
"As a Software Engineer,
I want a unified agent activation command in the IDE,
so that I can start using an agent without extra setup."

DO:
- Use persona from Epic (PRD-defined), not generic labels
- Make the goal specific and implementation-agnostic
- State tangible business/user value

DON'T:
- Invent personas not defined in PRD/Epic
- Use vague goals like "improve experience"
- Include solution details here
-->

## Acceptance Criteria

{{acceptance_criteria}}

<!-- Specific, testable conditions that define completion at story level -->
<!-- Story ACs SHOULD include verification commands and expected outputs -->

<!-- Required Structure for each criterion:
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
-->

<!-- Template Guidance:
Acceptance Criteria Example:
"1. Unified activation command works in Cursor
   - Scenario: Given a valid agent config, when the user runs the activation command in Cursor, then the agent is activated
   - Expected Behavior: IDE shows active agent; no extra prompts
   - Verification Method: automated; run `python tools/activate_agent.py --ide=cursor --agent=architect` (Expect: exit 0, "Activated" in stdout)
   - Evidence: Command output artifact attached
   - Non-automatable at story-level: false
   - Files Created/Changed: `.ide/agents/architect.state`
   - Guardrails: Command completes in <= 3s on staging; no P0/P1 errors
   - Test Data/Fixtures: `fixtures/agents/architect.yaml`
   - Environment/Flags: `FEATURE_UNIFIED_ACTIVATION=1`
   - Traceability: Epic AC #1; BR2, BR3

2. Session continuity during workday
   - Scenario: Given a valid SSO session, when the user switches from Tool A to Tool B, then no re-auth is required within policy window
   - Expected Behavior: Continuous access without credential prompts
   - Verification Method: manual; steps 1–4 in staging with local IDP stub; observe uninterrupted access and audit log entry
   - Evidence: Screen recording + audit log excerpt
   - Non-automatable at story-level: true (follow-up: automate via Playwright against stub)
   - Files Created/Changed: None
   - Guardrails: Logout invalidates sessions in <= 60s
   - Test Data/Fixtures: Test user `sso_user@test`, seeded session
   - Environment/Flags: `SESSION_TTL_MINUTES=480`; local IDP stub enabled
   - Traceability: Epic AC #2; BR5, NFR-SEC

3. Failure handling
   - Scenario: Given the dependency is unavailable, when activation is attempted, then a clear, actionable error is shown and retry succeeds after recovery
   - Expected Behavior: Friendly error; subsequent retry succeeds
   - Verification Method: semi-automated; curl request to stubbed endpoint (Expect: HTTP 503 with error JSON), then after recovery (Expect: HTTP 200)
   - Evidence: Captured curl responses + log lines
   - Non-automatable at story-level: false
   - Files Created/Changed: `logs/activation.log`
   - Guardrails: No corrupted state; no secrets in logs
   - Test Data/Fixtures: Outage flag in stub config
   - Environment/Flags: `ACTIVATION_RETRY=1`
   - Traceability: Epic AC #3; NFR-UX1, NFR-REL"

Edge cases to cover (example list):
- Invalid agent name
- Missing permissions
- Network interruption mid-activation

DO:
- Make AC executable and non-interactive where possible
- Include expected outputs/exit codes
- Map each AC to Epic/PRD

DON'T:
- Use subjective language (e.g., "works well")
- Omit outputs or acceptance thresholds
- Depend on third-party production services for validation
-->

## Description

{{description}}

<!-- Context for why this story exists and its strategic importance -->
<!-- Provide background for implementation and Epic alignment -->

<!-- Template Guidance:
Content Focus:
- Why this story exists within the Epic
- Strategic importance and user/business context
- Relationships to other stories and architectural decisions
- Implementation philosophy or approach rationale (brief)
-->

## Tasks/Subtasks

{{tasks_subtasks}}

<!-- LLM-executable implementation plan with atomic tasks and validation -->
<!-- Each task maps to acceptance criteria with specific commands and file paths -->

<!-- Template Guidance:
Structure:
- [ ] Task 1: Short imperative description (AC: 1, 2)
  - [ ] Create/Edit file: `path/to/file`
  - [ ] Run: `non_interactive_command --args`
  - [ ] Verify: `assert_command/grep/exit 0`
  - [ ] Rollback Plan: Note revert steps if config/flags change
- [ ] Task 2: Short imperative description (AC: 3)

DO:
- Number tasks sequentially (1, 2, ...) and do not number subtasks
- Reference acceptance criteria ids for every task
- Use atomic subtasks with precise commands
- Mirror AC verification in tasks

DON'T:
- Skip numbering or mix numbering formats across tasks
- Create tasks that require human interpretation only
- Omit validation or rollback steps when needed
- Use vague action words like "handle" or "manage"
-->

## Implementation Results

{{implementation_results}}

<!-- Concrete outcomes populated AFTER completion (evidence-first) -->

<!-- Template Guidance:
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
-->

## QA Checklist

{{qa_checklist}}

<!-- Specific verification steps with commands and expected outputs -->
<!-- Enable automated testing and quality validation -->

<!-- Template Guidance:
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
-->

<!-- Definition of Ready (DoR) – optional checklist (commented)
- Persona exists in Epic/PRD and matches this story
- Dependencies resolved or test doubles defined
- Test data/fixtures identified and accessible
- AC are executable, non-interactive, and map to tasks
- Feature flags/rollout strategy identified
- Small enough to complete within one sprint (INVEST)
-->

<!-- Definition of Done (DoD) – optional checklist (commented)
- All AC pass with evidence linked in Implementation Results
- Unit/integration tests updated/added
- Docs/README updated if applicable
- Feature flags default state decided; rollback plan noted
- Security/privacy checks completed; no P0/P1 issues
-->
