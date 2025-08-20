# Task: Create Risk Register

## Description

Create a comprehensive risk register that identifies, analyzes, and plans responses for project risks. Following PMBoK risk management processes, the risk register is a repository of information about identified risks, their analysis, and planned responses. This document enables proactive risk management throughout the project lifecycle and serves as a key input for project planning and control activities.

## Prerequisites

- [ ] Project charter approved and available at `/docs/project-management/project-charter.md`
- [ ] Scope of Work completed at `/docs/project-management/scope-of-work.md`
- [ ] Project plan developed at `/docs/project-management/project-plan.md`
- [ ] Risk management plan defined within project plan
- [ ] Stakeholder risk tolerance levels understood

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/project-management-methodology.md
- ./.krci-ai/templates/risk-register-template.md
- /docs/project-management/project-charter.md (approved charter)
- /docs/project-management/scope-of-work.md (approved SoW)
- /docs/project-management/project-plan.md (approved plan)

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Reference project documents**: Build on charter, SoW, and project plan for risk identification
2. **Follow PMBoK risk processes**: Apply identify risks, qualitative/quantitative analysis, and plan risk responses
3. **Use structured identification**: Apply systematic risk identification techniques and categories
4. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for artifact dependencies
5. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
6. **Format output**: Use [risk-register-template.md](./.krci-ai/templates/risk-register-template.md) for consistent structure

## Output Format

- **Location**: `/docs/project-management/risk-register.md` (EXACT path and filename)
- **Length**: 5-15 pages depending on project complexity and risk count
- **Downstream Enable**: Enables risk monitoring, control, and response implementation

## Success Criteria

- [ ] **File saved** to `/docs/project-management/risk-register.md`
- [ ] **Risk identification** completed using multiple techniques and sources
- [ ] **Qualitative analysis** performed with probability and impact assessment
- [ ] **Risk priority ranking** established using risk score matrix
- [ ] **Risk responses planned** for all high and medium priority risks
- [ ] **Risk owners assigned** for each identified risk
- [ ] **Trigger conditions** defined for implementing risk responses
- [ ] **Contingency plans** developed for critical risks
- [ ] **Risk budget** allocated for response implementation
- [ ] **Monitoring approach** defined for ongoing risk tracking

## Execution Checklist

### Risk Identification Phase

- [ ] **Document review**: Analyze charter, SoW, and project plan for risks
- [ ] **Stakeholder interviews**: Gather risk information from project stakeholders
- [ ] **Brainstorming sessions**: Conduct structured risk identification workshops
- [ ] **Checklist analysis**: Use risk checklists and historical data
- [ ] **Expert judgment**: Engage subject matter experts for risk identification
- [ ] **SWOT analysis**: Analyze strengths, weaknesses, opportunities, and threats

### Risk Analysis Phase

- [ ] **Probability assessment**: Evaluate likelihood of each risk occurring
- [ ] **Impact analysis**: Assess potential impact on scope, schedule, cost, and quality
- [ ] **Risk scoring**: Calculate risk scores using probability × impact matrix
- [ ] **Risk prioritization**: Rank risks by score and importance to project
- [ ] **Risk categorization**: Group risks by source, area, or type
- [ ] **Quantitative analysis**: Perform detailed analysis for high-priority risks

### Risk Response Planning Phase

- [ ] **Strategy selection**: Choose appropriate response strategy (avoid, mitigate, transfer, accept)
- [ ] **Response planning**: Develop specific actions for each response strategy
- [ ] **Owner assignment**: Assign risk owners responsible for implementation
- [ ] **Trigger identification**: Define conditions that trigger response implementation
- [ ] **Contingency planning**: Develop backup plans for critical risks
- [ ] **Budget allocation**: Estimate costs for risk response implementation

### Documentation Phase

- [ ] **Register creation**: Use [risk-register-template.md](./.krci-ai/templates/risk-register-template.md) structure
- [ ] **Content validation**: Ensure all risk management components are included
- [ ] **Stakeholder review**: Review register with key stakeholders and team
- [ ] **File placement**: Save to exact location `/docs/project-management/risk-register.md`

### Integration Phase

- [ ] **Project plan integration**: Ensure risk responses are integrated into project plan
- [ ] **Schedule integration**: Include risk response activities in project schedule
- [ ] **Budget integration**: Include risk response costs in project budget
- [ ] **Communication planning**: Plan risk communication with stakeholders

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Comprehensive Identification**: Use multiple techniques to identify all significant risks
- **Objective Analysis**: Base probability and impact assessments on evidence and data
- **Actionable Responses**: Develop specific, measurable, and implementable risk responses
- **Clear Ownership**: Assign clear ownership and accountability for each risk

### LLM Error Prevention Checklist

- **Avoid**: Generic risk lists without project-specific analysis
- **Avoid**: Subjective probability/impact assessments without rationale
- **Avoid**: Vague or unactionable risk responses
- **Avoid**: Missing risk ownership or trigger conditions
- **Reference**: Use [risk-register-template.md](./.krci-ai/templates/risk-register-template.md) for all formatting guidance

### PMBoK Integration Context

The Risk Register follows PMBoK risk management processes by systematically identifying risks, performing qualitative and quantitative analysis, planning appropriate responses, and establishing monitoring procedures. This register integrates with project planning by informing schedule, budget, and resource decisions and enables ongoing risk monitoring and control. 