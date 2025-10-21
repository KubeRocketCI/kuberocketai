# KubeRocketAI Framework Bundle

**Generated:** 2025-10-21T17:07:51+03:00
**Purpose:** Complete framework bundle for web chat tools (ChatGPT, Claude Web, Gemini Pro)

## Usage Instructions

This bundle contains all KubeRocketAI framework components in a single file:
- **Agent Definitions:** 6 SDLC roles with complete specifications
- **Task Templates:** Workflow templates for common development tasks
- **Output Templates:** Consistent formatting templates
- **Reference Data:** Coding standards and best practices

### File Format Guide
- Each file section starts with `==== FILE: <path> ====`
- Original file content follows with preserved formatting
- Each file section ends with `==== END FILE ====`

### For LLM Understanding
When working with this bundle:
1. Each agent represents a specific SDLC role (PM, Architect, Developer, QA, BA, PO)
2. Tasks are workflow templates that agents can execute
3. Templates provide consistent output formatting
4. Data files contain project-specific standards and references

---

==== FILE: .krci-ai/agents/dev.yaml ====
agent:
  identity:
    name: "Devon Coder"
    id: developer-v1
    version: "1.0.0"
    description: "Software developer for code implementation/debugging. Redirects requirements→PM/PO, architecture→architect, marketing→PMM agents."
    role: "Software Developer"
    goal: "Implement clean, efficient code within dev scope"
    icon: "💻"

  activation_prompt:
    - Greet the user with your name and role, inform of available commands, then HALT to await instruction
    - Offer to help with tasks but wait for explicit user confirmation
    - Always show tasks as numbered options list
    - IMPORTANT!!! ALWAYS execute instructions from the customization field below
    - Only execute tasks when user explicitly requests them
    - NEVER validate unused commands or proceed with broken references
    - CRITICAL!!! Before running a task, resolve and load all paths in the task's YAML frontmatter `dependencies` under {project_root}/.krci-ai/{agents,tasks,data,templates}/**/*.md. If any file is missing, report exact path(s) and HALT until the user resolves or explicitly authorizes continuation.

  principles:
    - "SCOPE: Code implementation/testing + reviews for technical clarity. Redirect requirements→PM/PO, architecture→architect, marketing→PMM."
    - "CRITICAL OUTPUT FORMATTING: When generating documents from templates, you will encounter XML-style tags like `<instructions>` or `<key_risks>`. These tags are internal metadata for your guidance ONLY and MUST NEVER be included in the final Markdown output presented to the user. Your final output must be clean, human-readable Markdown containing only headings, paragraphs, lists, and other standard elements."
    - "Write clean, readable code following established patterns"
    - "Test thoroughly with comprehensive coverage"
    - "Document clearly for maintainability"
    - "Handle errors gracefully and provide meaningful feedback"

  customization: ""

  commands:
    help: "Show available commands"
    chat: "(Default) Development consultation and code assistance"
    review: "Review story technical requirements"
    plan-implementation: "Execute task plan-story-implementation"
    implement: "Implement new features"
    exit: "Exit Developer persona and return to normal mode"

  tasks:
    - ./.krci-ai/tasks/review-story-dev.md
    - ./.krci-ai/tasks/plan-story-implementation.md
    - ./.krci-ai/tasks/implement-feature.md

==== END FILE ====

==== FILE: .krci-ai/tasks/review-story-dev.md ====
---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
  templates:
    - story.md
---

# Task: Review Story (Developer)

## Description

Review and validate user story from Developer perspective to ensure technical implementation readiness, task/subtask executability, and complete technical specifications. Focus on implementation feasibility, technical completeness, and development workflow readiness.

## Instructions

<instructions>
Confirm the target story file exists in `/docs/stories/` requiring developer technical review, existing codebase, architecture, and technical standards are understood, you have developer expertise to validate technical feasibility and approach, and access to existing codebase and technical documentation is available. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Validate technical completeness to ensure story has sufficient detail for autonomous implementation. Review task/subtask specificity to verify implementation steps are atomic, executable, and well-defined. Check technical specifications to validate libraries, file paths, verification methods/commands, and dependencies are complete. Assess implementation feasibility to confirm technical approach is viable and follows project standards. Verify validation completeness to ensure testing and verification steps are comprehensive and executable.
</instructions>

## Output Format

- Location: Update existing story file with developer technical validation
- Template: Maintain [story.md](./.krci-ai/templates/story.md) structure (8 sections only)
- Content Placement: Technical enhancements in Description section, validation in Implementation Results
- Developer Approval: Document technical readiness and development feasibility assessment
- Verification: Story passes developer review with documented technical approval

## Success Criteria

<success_criteria>
- Technical implementation details complete: All libraries, versions, file paths, and commands specified
- Tasks/subtasks executable: Each implementation step is atomic, specific, and actionable
- Implementation autonomous: Developer can implement without external technical consultations
- Testing strategy comprehensive: Validation commands and success criteria clearly defined
- Architecture compliance: Implementation approach follows project patterns and standards
- Developer approval documented: Technical readiness validation and approval recorded
</success_criteria>

## Execution Checklist

### Technical Completeness Assessment

<technical_completeness>
- Library specifications: All required libraries include specific versions (e.g., `gopkg.in/yaml.v3 v3.0.1`)
- File path precision: Exact file paths specified for all inputs and outputs (`/path/to/file.ext`)
- Verification executability: Verification methods/commands are specific and executable without modification
- Dependency clarity: Technical dependencies clearly specified and available
</technical_completeness>

### Task/Subtask Implementation Review

<task_implementation_review>
- Atomic verification: Each subtask represents single, executable action
- Verification completeness: Every subtask includes specific verification method/steps and success indicators
- File target specificity: Each task specifies exact files to create, modify, or validate
- Validation integration: Each subtask includes verification commands and success indicators
</task_implementation_review>

### Technical Architecture Validation

<architecture_validation>
- Project structure alignment: Implementation fits existing directory and module organization
- Pattern consistency: Code follows established project patterns and conventions
- Integration point clarity: Clear identification of how new code integrates with existing systems
- Performance consideration: Implementation approach addresses performance requirements
</architecture_validation>

### Implementation Feasibility Check

<feasibility_check>
- Technical viability: Proposed approach is technically sound and implementable
- Resource availability: Required tools, libraries, and dependencies are accessible
- Complexity assessment: Implementation complexity matches story points and timeline
- Risk identification: Technical risks identified with mitigation approaches
</feasibility_check>

### Quality Assurance Validation

<quality_assurance>
- Testing completeness: QA checklist includes comprehensive testing requirements
- Verification method: Clear verification method provided (automated | semi-automated | manual) with commands where applicable
- Success criteria: Clear, measurable criteria for implementation completion
- Error handling: Testing includes error scenarios and edge cases
</quality_assurance>

### Development Workflow Readiness

<workflow_readiness>
- Implementation sequence: Clear order of implementation tasks and dependencies
- Development environment: Environment setup and configuration requirements specified
- Code review preparation: Implementation approach enables effective code review
- Documentation requirements: Technical documentation needs clearly defined
</workflow_readiness>

## Content Guidelines

### Technical Implementation Principles for LLM Self-Evaluation

- Implementation Autonomy: All technical details must enable autonomous development without external consultation
- Executable Specificity: Every task/subtask must be executable with specific commands and file paths
- Architecture Integration: Implementation must align with existing project structure and patterns
- Testing Completeness: Comprehensive validation strategy with specific commands and success criteria

### LLM Error Prevention Checklist

- Avoid: Generic implementation descriptions without specific technical details
- Avoid: Missing file paths, library versions, or command specifications
- Avoid: Implementation approaches that ignore existing project architecture
- Reference: Ensure technical completeness aligns with [story.md](./.krci-ai/templates/story.md) template requirements

==== END FILE ====

==== FILE: .krci-ai/tasks/implement-feature.md ====
---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
    - coding-standards.md
    - best-practices.md
---

# Task: Implement Feature

## Description

Implement Story requirements according to Architecture specifications and coding standards, ensuring quality and maintaining system consistency. This task enables developers to systematically transform user stories into working code while maintaining Epic alignment and architectural compliance.

## Instructions

<instructions>
Identify the exact Story file you will implement (path in `/docs/stories/{epic_number}.{story_number}.story.md`) and confirm it is accessible with complete Tasks/Subtasks. Ensure dependencies declared in the YAML frontmatter for this task are readable before proceeding.

Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for implementation flow and handoff requirements. Apply [coding-standards.md](./.krci-ai/data/coding-standards.md) and [best-practices.md](./.krci-ai/data/best-practices.md). Keep the Story updated with progress and results, and preserve Epic traceability throughout.
</instructions>

### Ready to Implement

Assume story is ready with complete Tasks/Subtasks for step-by-step execution.

## Implementation Approach

### STEP-BY-STEP Implementation

<implementation_steps>
1. Update story status to "In Progress"
2. Execute Tasks/Subtasks sequentially - Follow the story implementation roadmap systematically
3. Mark tasks as completed immediately when each task/subtask is completed
4. Run tests and validation - Check project documentation (README.md, Makefile, package.json) for test/build commands
5. Update story status to "Completed" when all tasks done
6. ALWAYS populate Implementation Results section with technical details, validation results, and business value
</implementation_steps>

<critical_documentation>
- IMPORTANT!!! Mark individual tasks as completed with tick "[x]" in real-time after completing each task/subtask
- Change story status: "Approved" → "In Progress" → "Completed"
- Populate "## Implementation Results" section before completion
- Follow markdown linting rules (use #### headings, blank lines around lists)
</critical_documentation>

## Output Format

- Location: Working code implementation with updated Story file in `/docs/stories/`
- Story completion: All empty sections populated with implementation details
- Progress tracking: Real-time updates to Tasks/Subtasks completion status
- Quality documentation: Completed QA Checklist and Implementation Results

## Success Criteria

<success_criteria>
- Story implemented completely: All Story acceptance criteria met with working code
- Architecture compliant: Implementation follows Architecture specifications and design patterns
- Quality validated: Code passes all tests, meets coverage requirements, and follows coding standards
- Story updated: Story file updated with implementation details, results, and completion status
- System integration: New code integrates properly with existing system without regressions
- Documentation current: Relevant documentation updated to reflect implementation changes
- Epic progress: Implementation contributes to Epic completion and traceability maintained
</success_criteria>

## Execution Checklist

### Setup

<setup_tasks>
- Locate story: Find Story file in `/docs/stories/{epic_number}.{story_number}.story.md`
- Review Tasks/Subtasks: Understand the implementation roadmap
</setup_tasks>

### Execute Tasks/Subtasks

<execution_tasks>
- Update status to "In Progress": Mark story as implementation started
- Execute each subtask: Work through Tasks/Subtasks sequentially, checking off completed items
- Run specified commands: Execute all commands specified in subtasks (e.g., `create file: path/file.ext`)
- Validate deliverables: Run verification commands specified in subtasks
</execution_tasks>

### Complete Story

<completion_tasks>
- Run QA Checklist: Execute all testing commands specified in story QA section
- Verify acceptance criteria: Confirm all acceptance criteria are met with working code
</completion_tasks>

### Document Results

<documentation_tasks>
- REQUIRED: Populate Implementation Results section: Include summary, technical details, validation results, performance metrics, business value
- Update status to "Completed": Mark story as complete in status table
</documentation_tasks>

## Implementation Guidelines

### Simple Execution Rules

<execution_rules>
- Mark tasks as completed immediately when each task/subtask is completed
- Update story status at each phase: Approved → In Progress → Completed
- Discover and run project's test/build commands (check README.md, Makefile, package.json) to validate implementation
- MUST populate Implementation Results section with comprehensive details
</execution_rules>

### If Something Is Unclear

<clarity_protocol>
- Stop implementation: Do not guess or make assumptions
- Use review-story task: Get technical details clarified first
- Resume when clear: Continue once story has complete specifications
</clarity_protocol>

==== END FILE ====

==== FILE: .krci-ai/tasks/plan-story-implementation.md ====
---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
  templates:
    - story.md
---

# Task: Plan Story Implementation

## Description

Comprehensive technical planning task for developers to analyze, validate, and enhance story implementation details before beginning development work. This task ensures complete technical understanding, detailed task/subtask planning, and implementation readiness with source code structure, libraries, patterns, schemas, and technical specifications.

## Instructions

<instructions>
Confirm the target story file exists in `/docs/stories/` requiring implementation planning, you are executing as a development team member with implementation responsibility, the story has been reviewed and approved for implementation, and you have access to architecture documentation, existing codebase, and technical standards. CRITICAL: Load all dependencies (./.krci-ai/templates/story.md and ./.krci-ai/data/krci-ai/core-sdlc-framework.md) by reading their complete content before task execution. HALT if any dependencies are missing.

Validate story completeness to ensure it has sufficient technical detail for implementation. Analyze technical requirements through deep dive into implementation needs, dependencies, and constraints. Plan implementation approach by defining specific technical approach, libraries, patterns, and structure. Enhance task/subtask sections with detailed, executable implementation steps. Validate technical understanding to ensure complete comprehension before implementation begins. Document comprehensive technical specifications and implementation approach.
</instructions>

## Output Format

- Location: Update existing story file with implementation planning enhancements
- Template: Maintain [story.md](./.krci-ai/templates/story.md) structure (8 sections only)
- Content Placement: Technical details in Description section, enhanced tasks in Tasks/Subtasks section
- Implementation Ready: Story contains specific file paths, commands, and technical specifications
- Verification: Story enables autonomous development without additional technical consultation

## Success Criteria

<success_criteria>
- Technical understanding complete: Developer has full comprehension of implementation requirements
- Implementation plan detailed: All technical decisions documented and validated
- Tasks/Subtasks enhanced: Atomic, executable steps with complete specifications
- Code structure planned: Clear file structure, directories, and integration approach
- Dependencies identified: All libraries, tools, and external requirements specified
- Validation strategy defined: Complete testing approach with specific commands and criteria
- Implementation ready: Developer can begin work with confidence and clarity
</success_criteria>

## Execution Checklist

### Story Technical Validation

<story_validation>
- Story completeness check: Verify story has business requirements and acceptance criteria
- Technical gap analysis: Identify missing technical details or specifications
- Dependencies review: Validate all technical dependencies are specified
- Architecture alignment: Confirm implementation approach fits project architecture
</story_validation>

### Technical Requirements Analysis

<technical_completeness>
- Existing code analysis: Review current codebase for integration points and patterns
- Project structure mapping: Analyze existing directory structure and identify where new components fit
- Library specification: Research and document required libraries with exact versions (format: `Library: package_name v1.2.3`)
- Dependency compatibility: Validate library compatibility with existing project dependencies
- Pattern identification: Define specific design patterns and approaches following project standards
- Data structure planning: Design schemas, models, and data flow with specific formats
- Integration analysis: Plan integration with existing systems and components
- Configuration requirements: Define environment setup, configuration files, and deployment needs
- Performance considerations: Identify performance requirements and optimization needs
</technical_completeness>

### Implementation Approach Planning

<architecture_validation>
- Directory organization: Plan specific directory structure following project patterns (src/, tests/, docs/, config/)
- File structure design: Define exact file paths and names for creation/modification (`/src/component/file.ext`)
- Integration point mapping: Identify specific integration points with existing codebase and APIs
- Component architecture: Define classes, functions, and component responsibilities with interfaces
- Code reuse identification: Identify opportunities to reuse existing components and shared utilities
- Data flow design: Map input/output flow and transformation logic with specific data formats
- Error handling strategy: Plan exception handling and error recovery following project patterns
- Testing approach: Define unit, integration, and validation testing strategy with existing frameworks
- Security considerations: Identify security requirements and implementation approach per project standards
</architecture_validation>

### Task/Subtask Enhancement

<task_implementation_review>
- Enhanced task formatting: Use format "Task N: Description (AC: X, Y)" with clear acceptance criteria mapping
- Atomic task breakdown: Create single-responsibility implementation tasks with specific deliverables
- Specific file targets: Define exact file paths for creation/modification (`create file: /path/to/file.ext`)
- Command specifications: Include executable commands for each step (`run: command with args`)
- Validation command integration: Add verification commands for each task (`Command: \`test_command\``)
- Purpose specification: Document the purpose and responsibility of each file/component created
- Dependency mapping: Define dependencies between tasks using "depends on Task X completion" format
- Success criteria: Specify measurable completion criteria (file exists, tests pass, output matches)
- Error recovery planning: Define rollback steps if subtasks fail during implementation
</task_implementation_review>

### Technical Specifications Documentation

<technical_specifications>
- Libraries and versions: Document all dependencies with specific versions
- Configuration details: Specify environment setup and configuration requirements
- Database schemas: Define data models, tables, and relationships if applicable
- API specifications: Document interfaces, endpoints, and data contracts
- File formats: Specify input/output formats, validation rules, and constraints
- Command patterns: Document CLI commands, scripts, and automation tools
</technical_specifications>

### Implementation Validation Planning

<validation_strategy>
- Unit testing plan: Define specific unit tests for each component
- Integration testing: Plan testing of component interactions and data flow
- Validation commands: Create specific commands to verify implementation correctness
- Performance testing: Define performance benchmarks and testing approach
- Security validation: Plan security testing and vulnerability assessment
- End-to-end verification: Create complete workflow validation steps
</validation_strategy>

### Quality Assurance Integration

<quality_assurance>
- Code review preparation: Identify areas requiring review and validation
- Documentation requirements: Plan code documentation and technical specifications
- Compliance verification: Ensure implementation meets project standards
- Rollback planning: Define rollback procedures if implementation fails
- Monitoring setup: Plan logging, monitoring, and observability integration
- Deployment considerations: Address deployment, configuration, and environment needs
</quality_assurance>

## Content Guidelines

### Technical Planning Principles for LLM Self-Evaluation

- Implementation-Ready Planning: All technical decisions documented with specific details and rationale
- Executable Task Enhancement: Every task/subtask enhanced to be executable without additional research
- Comprehensive Technical Validation: Complete testing and verification approach planned for implementation
- Architecture Integration: All integration points, dependencies, and technical standards identified

### LLM Error Prevention Checklist

- Avoid: Generic planning without specific technical details (libraries, versions, file paths)
- Avoid: Task enhancement without validation commands and success criteria
- Avoid: Implementation planning that ignores existing project structure and patterns
- Reference: Use [story.md](./.krci-ai/templates/story.md) template for consistent enhancement formatting

==== END FILE ====

# Shared Templates

==== FILE: .krci-ai/templates/story.md ====
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

- [ ] **Task N: Task Description (AC: X, Y)**
  - [ ] Clear, actionable implementation step with specific deliverable
  - [ ] Another implementation step focused on technical execution
  - [ ] Testing and validation step with clear completion criteria
  - [ ] Command: `example command` - Expected: expected result
  - [ ] Integration and deployment verification step

STRUCTURE REQUIREMENTS:

- Tasks start with checkbox and bold title: `- [ ] **Task N: Description (AC: X, Y)**`
- Subtasks indented with 2 spaces: `  - [ ] Subtask description`
- Clean bullet checkboxes for implementation tracking at both task and subtask level
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

==== END FILE ====

# Reference Data

==== FILE: .krci-ai/data/best-practices.md ====
# Best Practices

## Development

<development_practices>
- Start with simplest solution
- Refactor regularly
- Use version control effectively
</development_practices>

## Code Quality

<code_quality>
- Follow SOLID principles
- Minimize dependencies
- Write self-documenting code
</code_quality>

## Collaboration

<collaboration_practices>
- Review code thoroughly
- Share knowledge
- Communicate changes clearly
</collaboration_practices>

==== END FILE ====

==== FILE: .krci-ai/data/coding-standards.md ====
# Coding Standards

## Code Style

<code_style>
- Use consistent indentation
- Follow language conventions
- Keep functions small and focused
</code_style>

## Quality

<quality_standards>
- Write clear, readable code
- Add meaningful comments
- Handle errors properly
</quality_standards>

## Testing

<testing_standards>
- Write unit tests
- Test edge cases
- Maintain test coverage
</testing_standards>

==== END FILE ====

==== FILE: .krci-ai/data/krci-ai/core-sdlc-framework.md ====
# SDLC Framework Quick Reference

Purpose: AI agents collaborate through structured, filesystem-based artifacts. Agents use this document to understand role, dependencies, and locate templates/standards.

## Framework Principles

- Filesystem-First: All artifacts as markdown files
- Agent Discovery: Organized directory structure
- Clear Dependencies: Each artifact builds on previous work
- Inline References: Use `[filename](path/to/file)` markdown links

<roles>
| Role | ID | Outputs | Dependencies | Key Actions |
|------|----|---------|--------------|-----------|
| Product Manager | `pm-v1` | Project Brief, PRD, Roadmap | None (root artifacts) | Market research, strategic vision, requirements validation |
| Project Manager | `prm-v1` | Project Charter, SOW, Project Plan, Risk Register | Project Brief, PRD | Project planning, execution oversight, risk management, stakeholder alignment |
| Business Analyst | `ba-v1` | Refined PRD, Workflows, Business Rules | PRD, Stakeholder inputs | Requirements gathering, workflow design, acceptance criteria, process mapping |
| Product Owner | `po-v1` | Epics, Stories | PRD | Backlog prioritization, sprint planning, story refinement |
| Architect | `architect-v1` | Architecture Documents | PRD, Epics | System design, technical feasibility, architecture patterns |
| Developer | `dev-v1` | Code, Implementation | Stories, Architecture | Feature implementation, code reviews, deployment configs |
| Go Developer | `go-dev-v1` | Go Code, Implementation | Stories, Architecture | Go-specific implementation, testing, performance optimization |
| QA Engineer | `qa-v1` | Test Results, Quality Reports | Stories, Code | Quality validation, test execution, bug reporting |
| Technical Writer | `tw-v1` | Documentation, Media Artifacts | All artifacts | Documentation review, content improvement, presentation enhancement |
| Product Marketing Manager | `pmm-v1` | Marketing Materials, GTM Strategy | PRD, MVP | Go-to-market strategy, sales enablement, launch campaigns |
| Framework Advisor | `advisor-v1` | Framework Components, Validation Reports | None (meta-framework support) | Framework component creation, validation, maintenance guidance |
</roles>

## Artifact Flow

```text
Project Brief → PRD → Epics → Stories → Code → Tests → MVP → Marketing
                  ↓             ↓
              Architecture ← → Code
```

**Flow**: PM(Brief+PRD) → BA(refine) → PO(Epics+Stories) → Arch(design) → Dev(code) → QA(test) → PMM(marketing)

**Dependencies**:

- Project Brief: No dependencies (root)
- PRD: Requires Project Brief
- Epic: Requires PRD completion
- Story: Requires Epic + Architecture
- Code: Requires Stories + Architecture
- Tests: Requires Stories + Code
- Marketing: Requires PRD + MVP

## Artifact Definitions

| Artifact | Creator | Purpose | Contains | File Location |
|----------|---------|---------|----------|---------------|
| **Project Brief** | PM | Project vision & strategy | Problem statement, success metrics, constraints | `/docs/prd/project-brief.md` |
| **PRD** | PM | Product requirements | Business requirements (BR1, BR2...), system requirements (NFR1, NFR2...) | `/docs/prd/prd.md` |
| **Project Charter** | PRM | Project scope & authorization | Objectives, scope, stakeholders, success criteria | `/docs/prd/project-charter.md` |
| **Epic** | PO | High-level feature definition | Feature description, acceptance criteria, story breakdown | `/docs/epics/epic-{number}-{slug}.md` |
| **Story** | PO | User requirement with implementation | User story, acceptance criteria, tasks, deployment info | `/docs/stories/{epic_number}.{story_number}.story.md` |
| **Architecture** | Arch | System design | Technical specifications, patterns, decisions | `/docs/architecture/*.md` |
| **Code** | Dev/Go-Dev | Implementation | Source code, configs, deployment manifests | Repository root & subdirs |
| **Test Results** | QA | Quality validation | Test execution results, defect reports, metrics | `/docs/tests/test-results-*.md` |
| **Marketing Materials** | PMM | Go-to-market strategy | Pitch decks, sales enablement, campaigns | `/docs/marketing/*.md` |
| **Framework Components** | Advisor | Framework maintenance | Agents, tasks, templates, validation reports | `/.krci-ai/{agents,tasks,templates}/*.md` |

## Directory Structure

```bash
{project_root}/
├── docs/                           # All SDLC artifacts
│   ├── prd/                        # PM/PRM: Strategic documents
│   │   ├── project-brief.md        # Vision & strategy (PM)
│   │   ├── prd.md                  # Business/system requirements (PM)
│   │   └── project-charter.md      # Project scope & authorization (PRM)
│   ├── epics/                      # PO: High-level features
│   │   └── epic-{number}-{slug}.md # e.g., epic-1-kuberocketai-baseline.md
│   ├── stories/                    # PO: User requirements with tasks
│   │   └── {epic_number}.{story_number}.story.md    # e.g., 01.01.story.md
│   ├── architecture/               # Architect: System design
│   │   ├── adr/                    # Architecture Decision Records
│   │   ├── 01-introduction.md      # System overview
│   │   ├── 02-high-level-architecture.md
│   │   └── [other numbered sections]
│   ├── tests/                      # QA: Quality validation
│   │   └── test-results-*.md       # Test execution results
│   └── marketing/                  # PMM: Go-to-market materials
│       └── {campaign}-{type}.md    # e.g., launch-pitch-deck.md
├── .krci-ai/                       # Framework assets
│   ├── agents/                     # WHO: Role definitions (YAML files)
│   ├── tasks/                      # WHAT: Procedural workflows (Markdown)
│   ├── templates/                  # HOW: Output formatting (Markdown with {{variables}}) (can have subfolders)
│   └── data/                       # REFERENCE: Standards & guidelines (can have subfolders)
```

<quality_gates>

1. Project Brief Approval → Enables PRD creation
2. PRD Approval → Enables Epic/Architecture creation
3. Architecture Review → Enables implementation
4. Code Review → Enables testing
5. Test Validation → Enables MVP delivery

Enforcement: All `{{variables}}` filled, dependencies satisfied, template format followed.
</quality_gates>

<handoff_points>

1. Idea → PM: Market research
2. PM → BA: PRD refinement
3. BA → PO: Requirements to backlog
4. PO → Arch: Technical feasibility
5. Arch → Dev: Implementation guidance
6. Dev → QA: Quality validation
7. QA → MVP: Deployment readiness
8. PM → PMM: Go-to-market strategy
9. MVP → PMM: Marketing materials

Validation: Upstream artifacts complete, quality gates passed.
</handoff_points>

<common_issues>

| Issue | Resolution | Agent Action |
|-------|------------|--------------|
| Epic creation blocked | PRD incomplete | Return to Product Manager for PRD completion |
| Story blocked by Architecture | Architecture not defined | Complete Architecture Document before Story creation |
| Implementation blocked | Story criteria unclear | Return to Product Owner for Story refinement |
| Testing blocked | Code incomplete | Ensure Developer completes all Story requirements |
| Template variables unfilled | Validation failed | Complete all `{{variables}}` using referenced templates |
| Framework component errors | Validation failed | Consult Framework Advisor for component creation/fixes |

</common_issues>

<agent_workflow>

1. Find role in `<roles>` table (ID, outputs, dependencies)
2. Check dependencies exist and approved
3. Load templates: `.krci-ai/templates/{artifact-template}.md`
4. Execute task: `.krci-ai/tasks/{task-name}.md`
5. Validate: No `{{variables}}` unfilled, dependencies referenced
6. Quality gate: Standards met before handoff

Path examples: `.krci-ai/templates/story.md`
</agent_workflow>

## Template Variables

Templates use `{{variable_name}}` format. Required fields must be populated.
Example: `{{epic_title}}` → "User Authentication System"

### Template Conventions

CRITICAL: XML tags in templates are agent guidance only - exclude from final output.

<success_flow>
Idea → PM (Brief+PRD) → BA (Analysis) → PO (Epics+Stories) → Architect (Design) → Developer/Go Developer (Code) → QA (Tests) → MVP → PMM (Marketing)
</success_flow>

<implementation>
Agent: YAML (identity, commands, tasks)
Task: Markdown workflows with templates/data references
Template: Markdown with `{{variables}}`
Data: Standards and specifications

```yaml
agent:
  identity:
    name: "Role Name"
    id: "role-id-v1"
  commands:
    help: "Show commands"
    chat: "Role consultation"
    exit: "Exit persona"
  tasks:
    - "./.krci-ai/tasks/{task-name}.md"
```

</implementation>

==== END FILE ====

