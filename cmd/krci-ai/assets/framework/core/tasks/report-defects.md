# Task: Report Defects

## Description

Create comprehensive defect reports and quality assessments based on testing execution results, ensuring systematic documentation of issues and quality observations. This task translates testing findings into actionable defect reports with clear reproduction steps, impact assessment, and quality recommendations that enable development teams to address issues and stakeholders to make informed release decisions.

## Prerequisites

- [ ] **Test execution completed**: Testing execution finished with documented results and identified issues
- [ ] **Testing evidence**: Screenshots, logs, and supporting documentation collected during test execution
- [ ] **Quality standards**: Understanding of defect classification and severity standards from [quality-metrics.md](./.krci-ai/data/quality-metrics.md)
- [ ] **Reporting tools**: Access to defect tracking system and reporting templates

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/quality-metrics.md
- ./.krci-ai/templates/defect-report.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for defect reporting workflow and quality gates
2. **Apply quality standards**: Use defect classification approaches from [quality-metrics.md](./.krci-ai/data/quality-metrics.md)
3. **Format output**: Use [defect-report.md](./.krci-ai/templates/defect-report.md) for structured defect documentation
4. **Ensure traceability**: Link defects to specific test cases, Story acceptance criteria, and Epic requirements

## Output Format

**Defect Reports and Quality Assessment** - Create comprehensive quality documentation:

- [ ] **Defect reports**: Individual defect reports using [defect-report.md](./.krci-ai/templates/defect-report.md) template
- [ ] **Quality assessment**: Overall quality evaluation and release readiness recommendation
- [ ] **Traceability matrix**: Mapping between defects, test cases, and Story acceptance criteria
- [ ] **Priority recommendations**: Defect prioritization and resolution timeline suggestions

## Success Criteria

- [ ] **Defects documented** - All identified issues reported with comprehensive reproduction steps and evidence
- [ ] **Quality assessed** - Overall Story/Epic quality evaluated with clear release readiness recommendation
- [ ] **Priority assigned** - Defects classified by severity and priority with resolution recommendations
- [ ] **Traceability established** - Clear links between defects and affected Story acceptance criteria
- [ ] **Stakeholder informed** - Quality assessment and defect reports communicated to development and product teams
- [ ] **Action items defined** - Clear next steps and resolution timeline recommendations provided

## Execution Checklist

### Defect Analysis Phase

- [ ] **Issue identification**: Review test execution results and identify all defects and quality issues
- [ ] **Impact assessment**: Evaluate how each defect affects Story acceptance criteria and Epic functionality
- [ ] **Severity classification**: Assign severity levels (Critical, High, Medium, Low) based on business impact
- [ ] **Evidence compilation**: Organize screenshots, logs, and supporting documentation for each defect

### Defect Documentation Phase

- [ ] **Defect report creation**: Document each defect using [defect-report.md](./.krci-ai/templates/defect-report.md) format
- [ ] **Reproduction steps**: Provide detailed steps to reproduce each defect with specific test data
- [ ] **Expected vs actual results**: Clearly document expected behavior versus observed behavior
- [ ] **Environment details**: Include environment configuration, browser, device, and system information

### Quality Assessment Phase

- [ ] **Story impact evaluation**: Assess how defects affect Story acceptance criteria completion
- [ ] **Epic functionality review**: Evaluate overall Epic quality and feature readiness
- [ ] **Risk analysis**: Identify potential risks if defects are not resolved before release
- [ ] **Release readiness determination**: Make quality-based recommendation for Story/Epic release

### Communication and Follow-up Phase

- [ ] **Stakeholder notification**: Communicate quality assessment and defect findings to development and product teams
- [ ] **Priority recommendations**: Suggest defect resolution priority and timeline based on business impact
- [ ] **Resolution tracking**: Establish process for tracking defect resolution and verification
- [ ] **Follow-up planning**: Plan re-testing activities once defects are resolved

## Content Guidelines

### 🎯 **Defect Reporting Focus Areas:**

#### **Defect Documentation (Primary Focus):**

- **Clear Reproduction**: Step-by-step instructions that enable consistent defect reproduction
- **Impact Assessment**: Business impact evaluation and effect on Story acceptance criteria
- **Evidence Support**: Screenshots, logs, and data supporting defect identification
- **Environment Context**: Specific environment, browser, device, and configuration details

#### **Quality Assessment (Decision Support):**

- **Story Readiness**: Evaluation of Story completion readiness based on defect impact
- **Risk Analysis**: Assessment of risks associated with releasing despite known defects
- **Priority Guidance**: Recommendations for defect resolution priority and timeline
- **Release Decision Support**: Quality-based recommendations for stakeholder decision making

### ✅ **Quality Standards:**

- **Defect Reproducible**: All defects include sufficient detail for consistent reproduction
- **Impact Assessed**: Business and technical impact clearly evaluated for each defect
- **Evidence Supported**: Screenshots, logs, and supporting documentation included
- **Priority Classified**: Appropriate severity and priority assigned based on standards
- **Traceability Maintained**: Clear links to affected test cases and acceptance criteria
- **Stakeholder Communicated**: Quality findings clearly communicated to relevant teams

### ❌ **Common Pitfalls to Avoid:**

- Reporting defects without sufficient reproduction steps or supporting evidence
- Missing impact assessment and business context for identified issues
- Poor defect classification and priority assignment leading to confusion
- Inadequate communication of quality findings to development and product teams
- Reporting defects without linking to specific Story acceptance criteria or test cases
- Missing follow-up planning for defect resolution verification

### 🎯 **Story/Epic Quality Integration:**

This defect reporting should enable informed decision making by providing:

- **Story completion assessment** based on acceptance criteria validation and defect impact
- **Epic quality evaluation** considering overall functionality and user experience
- **Release readiness recommendation** based on comprehensive quality analysis
- **Action plan guidance** for defect resolution and quality improvement
