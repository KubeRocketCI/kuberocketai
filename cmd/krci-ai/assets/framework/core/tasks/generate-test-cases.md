# Task: Generate Test Cases

## Description

Generate detailed test cases and scenarios based on test plan strategy and Story acceptance criteria, ensuring comprehensive coverage of functional and non-functional requirements. This task translates test plan scenarios into executable test cases with clear steps, expected results, and validation criteria that enable systematic testing execution and quality validation.

## Prerequisites

<prerequisites>
- Test plan available: Approved test plan exists with defined test scenarios and strategy
- Story clarity: Stories with well-defined acceptance criteria available for test case generation
- Testing standards: Understanding of test case writing standards from [testing-standards.md](./.krci-ai/data/testing-standards.md)
- Quality metrics: Familiarity with test coverage requirements from [quality-metrics.md](./.krci-ai/data/quality-metrics.md)

### Reference Assets

Dependencies:

- ./.krci-ai/data/testing-standards.md
- ./.krci-ai/data/quality-metrics.md
- ./.krci-ai/data/krci-ai/core-sdlc-framework.md
- ./.krci-ai/data/test-methodologies.md
- ./.krci-ai/templates/test-cases.md

CRITICAL: Load all dependencies by reading their complete content before task execution. HALT if any missing.
</prerequisites>

## Instructions

<instructions>
1. Follow SDLC workflow: Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for test case generation workflow and quality gates
2. Apply testing methodologies: Use test case design techniques from [test-methodologies.md](./.krci-ai/data/test-methodologies.md)
3. Format output: Use [test-cases.md](./.krci-ai/templates/test-cases.md) for structured test case documentation
4. Ensure traceability: Map each test case to specific Story acceptance criteria and test plan scenarios
</instructions>

## Output Format

<output_format>
**Test Cases Documentation** - Create executable test specifications:

- Test case document: Complete test cases using [test-cases.md](./.krci-ai/templates/test-cases.md) template
- Functional test cases: Detailed test cases covering all Story acceptance criteria
- Non-functional test cases: Performance, security, and usability test cases based on requirements
- Test data specifications: Required test data and environment setup for test execution
</output_format>

## Success Criteria

<success_criteria>
- Test cases completed: All test scenarios from test plan converted to detailed executable test cases
- Coverage achieved: Every Story acceptance criterion covered by at least one test case
- Quality validated: Test cases follow testing standards and include clear validation criteria
- Traceability established: Clear mapping from test cases to Story acceptance criteria and test plan scenarios
- Execution ready: Test cases include sufficient detail for independent execution by team members
- Review approved: Test cases reviewed and approved by development team and QA stakeholders
</success_criteria>

## Execution Checklist

### Test Case Planning Phase

<test_case_planning>
- Test plan analysis: Review approved test plan scenarios and testing strategy
- Story acceptance criteria mapping: Identify all acceptance criteria requiring test case coverage
- Test case prioritization: Prioritize test case creation based on risk assessment and critical functionality
- Test design approach: Select appropriate test design techniques (equivalence partitioning, boundary value analysis, etc.)
</test_case_planning>

### Functional Test Case Development Phase

<functional_test_development>
- Positive test cases: Create test cases validating normal functionality and happy path scenarios
- Negative test cases: Design test cases for error conditions, invalid inputs, and edge cases
- Boundary testing: Generate test cases for boundary conditions and limit values
- User workflow testing: Create end-to-end test cases following user journey scenarios
</functional_test_development>

### Non-Functional Test Case Development Phase

<non_functional_test_development>
- Performance test cases: Design test cases for load, stress, and performance requirements
- Security test cases: Create test cases for authentication, authorization, and data protection
- Usability test cases: Generate test cases for user experience and accessibility requirements
- Compatibility testing: Design test cases for browser, device, and platform compatibility
</non_functional_test_development>

### Documentation and Validation Phase

<documentation_validation>
- Test case documentation: Document all test cases using [test-cases.md](./.krci-ai/templates/test-cases.md) format
- Test data preparation: Define required test data, user accounts, and environment configurations
- Traceability matrix: Create mapping between test cases and Story acceptance criteria
- Peer review: Conduct test case review with development team and obtain approval
</documentation_validation>

## Content Guidelines

### 🎯 **Test Case Generation Focus Areas:**

<test_case_focus_areas>

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
</test_case_focus_areas>

### ✅ **Quality Standards:**

<quality_standards>
- Requirements Traceable: Every test case maps to specific Story acceptance criteria
- Execution Ready: Test cases contain sufficient detail for independent execution
- Standards Compliant: Test cases follow testing standards and format guidelines
- Coverage Complete: All acceptance criteria covered with appropriate test scenarios
- Review Approved: Test cases validated by development team and QA stakeholders
- Maintainable: Test cases are structured for easy maintenance and updates
</quality_standards>

### ❌ **Common Pitfalls to Avoid:**

<common_pitfalls>
- Writing test cases without referencing specific Story acceptance criteria
- Creating overly complex test cases that are difficult to execute and maintain
- Missing negative test cases and edge condition scenarios
- Inadequate test data specification and environment requirements
- Poor traceability between test cases and requirements
- Test cases that cannot be executed independently
</common_pitfalls>

### 🎯 **Story Testing Integration:**

<story_testing_integration>
This test case generation should enable comprehensive quality validation by providing:

- Acceptance criteria validation through systematic test case coverage
- Story completion verification through executable test scenarios
- Quality gate enablement through clear pass/fail criteria
- Risk mitigation through comprehensive positive and negative test scenarios
</story_testing_integration>
