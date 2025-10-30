---
dependencies:
  data:
    - qa/test-methodologies.md
    - qa/testing-standards.md
    - shared/quality-metrics.md
    - krci-ai/core-sdlc-framework.md
---

# Task: Generate Test Cases

## Description

Generate detailed test cases and scenarios based on test plan strategy and Story acceptance criteria, ensuring comprehensive coverage of functional and non-functional requirements. This task translates test plan scenarios into executable test cases with clear steps, expected results, and validation criteria that enable systematic testing execution and quality validation.

## Instructions

<instructions>
Confirm approved test plan exists with defined test scenarios and strategy, Stories with well-defined acceptance criteria are available for test case generation, and BDD dependencies are accessible including `./src/main/resources/README.md` and `./src/main/resources/features/`. HALT if missing. Run prechecks to route user appropriately: if `./src/main/resources/features/` exists but README is missing, propose `onboard-testing`; if both are missing, propose `setup-testing`. Ensure dependencies declared in the YAML frontmatter (testing-standards.md, quality-metrics.md, test-methodologies.md, sdlc-framework.md) are readable before proceeding.

Ask user for input source: available stories under `docs/stories/` to scan and select, specific story file path like `docs/stories/NN.MM.story.md`, or user will paste task context in chat. HALT until input source confirmed and task context obtained.

Summarize planned actions before any search including what will be searched (keywords/themes), where (directories/namespace priority), expected action (extend existing vs create new), and open questions or assumptions. Proceed only after user confirms or refines intent.

Execute discovery phase using universal search for candidates. Build normalized keyword variants (hyphen/underscore/space/camelCase, common prefixes/suffixes like `-remote`, `sast`, `security`), detect domain/context hints (UI vs API, module/subsystem names) from task, and prioritize likely directories/namespaces in repository. ALWAYS PROMPT before search: "Rebuild semantic index now (ALL .feature files) to ensure up-to-date results? [yes/no]". On yes, prefer FAISS vector index (`./.krci-ai/indexes/gherkin-faiss.index` + `.meta.json`) using sentence-transformers with fallback to JSON index under `./.krci-ai/indexes/` if FAISS unavailable. On no, use existing FAISS/JSON index if present, otherwise scan files directly. If user approves rebuild, propose and run (upon approval) appropriate index build command for either FAISS or JSON/SQLite backend.

Use FAISS-first approach if available: encode intent, retrieve top-K candidates (e.g., 10), post-filter lexically using anchors (steps/tags/artifacts), and re-rank. Search across scenario titles, steps, tags, and Examples, including characteristic artifacts discovered in repository. Present top 3-5 candidates with file path and short snippet, then HALT for user choice (extend vs reject).

Review `./src/main/resources/README.md` for process, directory structure, and tagging rules (single source of truth). Analyze `./src/main/resources/features/` per README to determine coverage status (Covered / Partial / Not covered). Use README decision matrix and request confirmation before creating or updating tests.

Generate Gherkin by creating or extending `.feature` files under `./src/main/resources/features/` with proper tags and structure. Ensure traceability by mapping each Story acceptance criterion to specific feature files and scenarios. Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for test case generation workflow and quality gates. Apply test case design techniques from [test-methodologies.md](./.krci-ai/data/test-methodologies.md).

If user requests extending existing test, ask user to pick target scenario or confirm best match by exact/nearest title or unique anchor step/tag. Prefer in-place edits to that scenario by inserting after specified/matched step anchor and avoiding scenario header duplication. If change represents distinct flow, add new Scenario/Scenario Outline in same file without duplicating existing flows. Create new file only if no suitable host file exists or upon explicit user request. Show minimal diff preview around anchor and ask for confirmation before proceeding.

</instructions>
## Output Format

Gherkin Outputs - Create or update BDD feature specifications:

- Feature files: List of created/updated `.feature` files with paths
- Scenarios: New/updated Scenario/Scenario Outline titles and Examples blocks
- Functional test cases: Detailed test cases covering all Story acceptance criteria
- Non-functional test cases: Performance, security, and usability test cases based on requirements
- Test data specifications: Required test data and environment setup for test execution
- UI/API alignment: Which side is covered and any proposed complementary coverage
- Traceability: Story acceptance criteria mapped to feature files and scenarios

## Success Criteria

- Test cases completed - All test scenarios from test plan converted to detailed executable test cases
- Coverage achieved - Every Story acceptance criterion covered by at least one test case
- Quality validated - Test cases follow testing standards and include clear validation criteria
- Traceability established - Clear mapping from test cases to Story acceptance criteria and test plan scenarios
- Execution ready - Test cases include sufficient detail for independent execution by team members
- Review approved - Test cases reviewed and approved by development team and QA stakeholders

## Execution Checklist

### Test Case Planning Phase

- Test plan analysis: Review approved test plan scenarios and testing strategy
- Story acceptance criteria mapping: Identify all acceptance criteria requiring test case coverage
- Test case prioritization: Prioritize test case creation based on risk assessment and critical functionality
- Test design approach: Select appropriate test design techniques (equivalence partitioning, boundary value analysis, etc.)

### Functional Test Case Development Phase

- Positive test cases: Create test cases validating normal functionality and happy path scenarios
- Negative test cases: Design test cases for error conditions, invalid inputs, and edge cases
- Boundary testing: Generate test cases for boundary conditions and limit values
- User workflow testing: Create end-to-end test cases following user journey scenarios

### Non-Functional Test Case Development Phase

- Performance test cases: Design test cases for load, stress, and performance requirements
- Security test cases: Create test cases for authentication, authorization, and data protection
- Usability test cases: Generate test cases for user experience and accessibility requirements
- Compatibility testing: Design test cases for browser, device, and platform compatibility

### Documentation and Validation Phase

- Feature sources of truth: Ensure `.feature` files contain authoritative steps and data tables
- Test data preparation: Define required test data, user accounts, and environment configurations
- Traceability matrix: Map Story acceptance criteria to feature files and scenarios
- Peer review: Review Gherkin with the team and obtain approval

## Content Guidelines

### Test Case Generation Focus Areas

#### Story Acceptance Criteria (Primary Focus)

- Requirement Coverage: Each acceptance criterion must have corresponding test cases
- Positive Scenarios: Test cases validating expected functionality and business rules
- Negative Scenarios: Test cases for error handling and invalid input conditions
- Edge Cases: Test cases for boundary conditions and exceptional scenarios

#### Test Plan Implementation (Execution Focus)

- Scenario Translation: Convert test plan scenarios into detailed, executable test steps
- Test Data Requirements: Specify data needed for test execution and validation
- Environment Setup: Define required test environment configurations and dependencies
- Validation Criteria: Clear expected results and success criteria for each test case

### Quality Standards

- Requirements Traceable: Every test case maps to specific Story acceptance criteria
- Execution Ready: Test cases contain sufficient detail for independent execution
- Standards Compliant: Test cases follow testing standards and format guidelines
- Coverage Complete: All acceptance criteria covered with appropriate test scenarios
- Review Approved: Test cases validated by development team and QA stakeholders
- Maintainable: Test cases are structured for easy maintenance and updates

### Common Pitfalls to Avoid

- Writing test cases without referencing specific Story acceptance criteria
- Creating overly complex test cases that are difficult to execute and maintain
- Missing negative test cases and edge condition scenarios
- Inadequate test data specification and environment requirements
- Poor traceability between test cases and requirements
- Test cases that cannot be executed independently

### Story Testing Integration

This test case generation should enable comprehensive quality validation by providing:

- Acceptance criteria validation through systematic test case coverage
- Story completion verification through executable test scenarios
- Quality gate enablement through clear pass/fail criteria
- Risk mitigation through comprehensive positive and negative test scenarios
