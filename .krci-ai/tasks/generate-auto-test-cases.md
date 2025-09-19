---
dependencies:
  data:
    - test-methodologies.md
    - testing-standards.md
    - quality-metrics.md
    - krci-ai/core-sdlc-framework.md
---

# Task: Generate Test Cases

## Description

Generate detailed test cases and scenarios based on test plan strategy and Story acceptance criteria, ensuring comprehensive coverage of functional and non-functional requirements. This task translates test plan scenarios into executable test cases with clear steps, expected results, and validation criteria that enable systematic testing execution and quality validation.

## Prerequisites

- Test plan available: Approved test plan exists with defined test scenarios and strategy
- Story clarity: Stories with well-defined acceptance criteria available for test case generation
- Testing standards: Understanding of test case writing standards from [testing-standards.md](./.krci-ai/data/testing-standards.md)
- Quality metrics: Familiarity with test coverage requirements from [quality-metrics.md](./.krci-ai/data/quality-metrics.md)

### Reference Assets

Dependencies (BDD-only):

- ./src/main/resources/README.md
- ./src/main/resources/features/
- ./.krci-ai/data/test-methodologies.md (optional)
- ./.krci-ai/data/testing-standards.md (optional)
- ./.krci-ai/data/quality-metrics.md (optional)
- ./.krci-ai/data/krci-ai/core-sdlc-framework.md (optional)

Validation (HALT if missing):
- Require ./src/main/resources/README.md and ./src/main/resources/features/

Prechecks (route user appropriately):
- If `./src/main/resources/features/` exists and `./src/main/resources/README.md` is missing → propose running `onboard-testing` instead; HALT until user confirms.
- If both `./src/main/resources/README.md` and `./src/main/resources/features/` are missing → propose running `setup-testing` instead; HALT until user confirms.

## Instructions

0. Select input source (HALT until resolved):
   - Ask the user: "Where should I get the task text?"
   - Options:
     a) Scan repository for available stories under `docs/stories/` (list with numbers, then HALT to pick)
     b) Use a specific story file (user provides path like `docs/stories/NN.MM.story.md`; validate exists)
     c) User will paste the task context here in chat (do not scan until provided)
1. Intent confirmation (HALT):
   - Summarize planned actions before any search:
     - What will be searched (keywords/themes)
     - Where (directories/namespace priority)
     - Expected action (extend existing vs create new)
     - Open questions/assumptions
   - Proceed only after user confirms or refines the intent.
2. Discovery phase (universal search for candidates):
   - Build normalized keyword variants (hyphen/underscore/space/camelCase; add common prefixes/suffixes like `-remote`, `sast`, `security`).
   - Detect domain/context hints (e.g., UI vs API, module/subsystem names) from the task; prioritize likely directories/namespaces in this repository.
   - ALWAYS PROMPT before search: "Rebuild semantic index now (ALL .feature files) to ensure up-to-date results? [yes/no]".
     - On yes: prefer FAISS vector index (`./.krci-ai/indexes/gherkin-faiss.index` + `.meta.json`) using sentence-transformers; fallback to JSON index under `./.krci-ai/indexes/` if FAISS is unavailable.
     - On no: use existing FAISS/JSON index if present; otherwise scan files directly.
     - If user approves rebuild, propose and run (upon approval) one of the following commands:

```bash
pwsh -NoProfile -Command "python ./.krci-ai/scripts/build_gherkin_faiss.py --root . --model sentence-transformers/all-MiniLM-L6-v2 --out-index ./.krci-ai/indexes/gherkin-faiss.index --out-meta ./.krci-ai/indexes/gherkin-faiss.meta.json"
```

```bash
pwsh -NoProfile -Command "python ./.krci-ai/scripts/build_gherkin_index.py --root . --out-json ./.krci-ai/indexes/gherkin-lex.json --out-sqlite ./.krci-ai/indexes/gherkin-lex.sqlite"
```
   - FAISS-first if available: encode intent → top-K (e.g., 10) → post-filter lexically (anchors: steps/tags/artifacts) and re-rank.
   - Search across: scenario titles, steps, tags, Examples. Include characteristic artifacts discovered in this repo (avoid vendor-specific hardcoding).
   - Present top 3–5 candidates (file path + short snippet) and HALT for user choice (extend vs reject).
3. Read testing README: Review `./src/main/resources/README.md` for process, directory structure, and tagging rules (single source of truth)
4. Scan existing coverage: Analyze `./src/main/resources/features/` per README to determine Covered / Partial / Not covered
5. Decide actions: Use README decision matrix; request confirmation before creating/updating tests
6. Generate Gherkin: Create or extend `.feature` files under `./src/main/resources/features/` with proper tags and structure
7. Ensure traceability: Map each Story acceptance criterion to specific feature files and scenarios
8. Follow SDLC workflow: Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for test case generation workflow and quality gates
9. Apply testing methodologies: Use test case design techniques from [test-methodologies.md](./.krci-ai/data/test-methodologies.md)

### Extension policy (respect user intent)

- If user requests to extend an existing test:
  1) Ask the user to pick the target scenario or confirm best match (by exact/nearest title or a unique anchor step/tag).
  2) Prefer in-place edits to that scenario: insert after a specified/matched step anchor; avoid duplicating the scenario header.
  3) If the change represents a distinct flow, add a new Scenario/Scenario Outline in the same file; do not duplicate existing flows.
  4) Create a new file only if no suitable host file exists or upon explicit user request.
  5) Show a minimal diff preview around the anchor and ask for confirmation.

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

### 🎯 Test Case Generation Focus Areas:

#### Story Acceptance Criteria (Primary Focus):

- Requirement Coverage: Each acceptance criterion must have corresponding test cases
- Positive Scenarios: Test cases validating expected functionality and business rules
- Negative Scenarios: Test cases for error handling and invalid input conditions
- Edge Cases: Test cases for boundary conditions and exceptional scenarios

#### Test Plan Implementation (Execution Focus):

- Scenario Translation: Convert test plan scenarios into detailed, executable test steps
- Test Data Requirements: Specify data needed for test execution and validation
- Environment Setup: Define required test environment configurations and dependencies
- Validation Criteria: Clear expected results and success criteria for each test case

### ✅ Quality Standards:

- Requirements Traceable: Every test case maps to specific Story acceptance criteria
- Execution Ready: Test cases contain sufficient detail for independent execution
- Standards Compliant: Test cases follow testing standards and format guidelines
- Coverage Complete: All acceptance criteria covered with appropriate test scenarios
- Review Approved: Test cases validated by development team and QA stakeholders
- Maintainable: Test cases are structured for easy maintenance and updates

### ❌ Common Pitfalls to Avoid:

- Writing test cases without referencing specific Story acceptance criteria
- Creating overly complex test cases that are difficult to execute and maintain
- Missing negative test cases and edge condition scenarios
- Inadequate test data specification and environment requirements
- Poor traceability between test cases and requirements
- Test cases that cannot be executed independently

### 🎯 Story Testing Integration:

This test case generation should enable comprehensive quality validation by providing:

- Acceptance criteria validation through systematic test case coverage
- Story completion verification through executable test scenarios
- Quality gate enablement through clear pass/fail criteria
- Risk mitigation through comprehensive positive and negative test scenarios
