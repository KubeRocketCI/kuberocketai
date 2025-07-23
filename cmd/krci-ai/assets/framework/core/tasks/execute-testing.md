# Task: Execute Testing

## Description

Execute test cases systematically to validate Story acceptance criteria and Epic functionality, ensuring comprehensive quality verification through structured testing procedures. This task translates test cases into actual testing execution with documented results, defect identification, and quality assessment that supports Story completion and release readiness decisions.

## Prerequisites

- [ ] **Test cases available**: Approved test cases exist with detailed execution steps and validation criteria
- [ ] **Implementation ready**: Story implementation completed with code deployed to testing environment
- [ ] **Test environment**: Testing environment configured with required test data and dependencies
- [ ] **Testing tools**: Access to testing tools, browsers, and validation resources

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/test-methodologies.md
- ./.krci-ai/templates/test-report.md
- ./.krci-ai/templates/defect-report.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for testing execution workflow and quality gates
2. **Apply testing methodologies**: Use execution practices from [test-methodologies.md](./.krci-ai/data/test-methodologies.md)
3. **Format output**: Use [test-report.md](./.krci-ai/templates/test-report.md) for test execution documentation
4. **Document results**: Record all test results, defects, and quality observations for stakeholder review

## Output Format

**Test Execution Results** - Create comprehensive testing documentation:

- [ ] **Test execution report**: Complete test results using [test-report.md](./.krci-ai/templates/test-report.md) template
- [ ] **Test case results**: Pass/fail status for each executed test case with detailed observations
- [ ] **Defect reports**: Documented defects found during testing using [defect-report.md](./.krci-ai/templates/defect-report.md)
- [ ] **Quality assessment**: Overall quality evaluation and recommendation for Story completion

## Success Criteria

- [ ] **Test execution completed** - All planned test cases executed with documented results
- [ ] **Coverage verified** - All Story acceptance criteria validated through test execution
- [ ] **Results documented** - Clear pass/fail status recorded for each test case with supporting evidence
- [ ] **Defects reported** - All identified issues documented with detailed reproduction steps
- [ ] **Quality assessed** - Overall quality evaluation completed with release readiness recommendation
- [ ] **Stakeholder informed** - Test results communicated to development team and product stakeholders

## Execution Checklist

### Test Execution Preparation Phase

- [ ] **Environment verification**: Confirm testing environment is configured correctly with required test data
- [ ] **Test case review**: Review test cases to be executed and understand validation criteria
- [ ] **Tool preparation**: Set up testing tools, browsers, and recording capabilities for evidence collection
- [ ] **Baseline establishment**: Document initial system state and environment configuration

### Functional Testing Execution Phase

- [ ] **Test case execution**: Execute each test case following documented steps systematically
- [ ] **Result documentation**: Record pass/fail status with detailed observations and evidence
- [ ] **Defect identification**: Identify and document any defects or deviations from expected results
- [ ] **Coverage tracking**: Track execution progress against Story acceptance criteria

### Non-Functional Testing Execution Phase

- [ ] **Performance testing**: Execute performance test cases and measure response times, load handling
- [ ] **Security testing**: Validate authentication, authorization, and data protection measures
- [ ] **Usability testing**: Assess user experience, accessibility, and interface design quality
- [ ] **Compatibility testing**: Test across different browsers, devices, and platform configurations

### Results Analysis and Reporting Phase

- [ ] **Test results compilation**: Compile all test results into comprehensive test report using [test-report.md](./.krci-ai/templates/test-report.md)
- [ ] **Defect reporting**: Document all defects using [defect-report.md](./.krci-ai/templates/defect-report.md) format
- [ ] **Quality assessment**: Evaluate overall quality and provide release readiness recommendation
- [ ] **Stakeholder communication**: Present results to development team and product stakeholders

## Content Guidelines

### 🎯 **Test Execution Focus Areas:**

#### **Story Acceptance Criteria Validation (Primary Focus):**

- **Systematic Execution**: Execute test cases in logical order following documented procedures
- **Evidence Collection**: Capture screenshots, logs, and data to support test results
- **Validation Verification**: Confirm each acceptance criterion is met through successful test execution
- **Defect Documentation**: Record any failures or deviations with detailed reproduction steps

#### **Quality Assessment (Evaluation Focus):**

- **Coverage Analysis**: Verify all planned test cases executed and acceptance criteria covered
- **Risk Evaluation**: Assess severity and impact of any identified defects
- **Release Readiness**: Provide quality-based recommendation for Story completion
- **Improvement Recommendations**: Suggest areas for quality improvement or additional testing

### ✅ **Quality Standards:**

- **Systematic Execution**: All test cases executed following documented procedures
- **Complete Documentation**: Test results recorded with clear pass/fail status and evidence
- **Defect Quality**: All defects documented with sufficient detail for reproduction and resolution
- **Coverage Verified**: All Story acceptance criteria validated through test execution
- **Quality Assessed**: Overall quality evaluation completed with stakeholder communication
- **Standards Compliant**: Test execution follows testing standards and quality procedures

### ❌ **Common Pitfalls to Avoid:**

- Executing tests without proper environment verification and setup
- Recording test results without sufficient detail or supporting evidence
- Missing defect documentation or inadequate reproduction steps
- Rushing through test execution without thorough validation
- Poor communication of test results and quality assessment to stakeholders
- Executing tests without understanding acceptance criteria and validation requirements

### 🎯 **Story Completion Integration:**

This test execution should enable Story completion decisions by providing:

- **Acceptance criteria validation** through systematic test case execution
- **Quality evidence** supporting Story completion and release readiness
- **Defect identification** enabling development team to address issues before release
- **Stakeholder confidence** through transparent test results and quality assessment
