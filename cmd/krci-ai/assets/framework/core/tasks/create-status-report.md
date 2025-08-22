# Task: Create Status Report

## Description

Create a comprehensive project status report that communicates project progress, performance, and health to stakeholders. Following PMBoK monitoring and controlling processes, the status report provides regular updates on project performance against baselines, issues, risks, and upcoming activities. This document enables informed decision-making and maintains stakeholder engagement throughout the project lifecycle.

## Prerequisites

- [ ] Project plan approved and available at `/docs/project-management/project-plan.md`
- [ ] Performance measurement baseline established
- [ ] Current project performance data collected
- [ ] Work progress and deliverable status tracked
- [ ] Risk register updated at `/docs/project-management/risk-register.md`

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/project-management-methodology.md
- ./.krci-ai/templates/status-report-template.md
- /docs/project-management/project-plan.md (approved plan)
- /docs/project-management/risk-register.md (current register)
- /docs/project-management/scope-of-work.md (approved SoW)

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Reference performance baselines**: Compare current progress against scope, schedule, and cost baselines
2. **Follow PMBoK monitoring processes**: Apply PMBoK monitor and control project work processes
3. **Collect current data**: Gather actual performance data from project execution
4. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for artifact dependencies
5. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
6. **Format output**: Use [status-report-template.md](./.krci-ai/templates/status-report-template.md) for consistent structure

## Output Format

- **Location**: `/docs/project-management/status-reports/status-report-YYYY-MM-DD.md` (EXACT path and filename)
- **Length**: 3-6 pages for comprehensive status communication
- **Frequency**: Regular reporting cadence as defined in communication management plan

## Success Criteria

- [ ] **File saved** to `/docs/project-management/status-reports/status-report-YYYY-MM-DD.md`
- [ ] **Executive summary** provides high-level project health assessment
- [ ] **Schedule performance** compared against baseline with variance analysis
- [ ] **Cost performance** compared against budget with earned value metrics
- [ ] **Scope progress** tracked with deliverable completion status
- [ ] **Quality metrics** reported with trend analysis
- [ ] **Risk status** updated with new risks and response effectiveness
- [ ] **Issues status** tracked with resolution progress
- [ ] **Upcoming activities** outlined for next reporting period
- [ ] **Stakeholder actions** clearly identified and assigned

## Execution Checklist

### Data Collection Phase

- [ ] **Schedule data**: Collect actual start/finish dates, % complete, and remaining duration
- [ ] **Cost data**: Gather actual costs, committed costs, and budget remaining
- [ ] **Scope data**: Track deliverable completion status and quality metrics
- [ ] **Resource data**: Monitor resource utilization and availability
- [ ] **Risk data**: Review current risk status from risk register
- [ ] **Issue data**: Track current issues and resolution progress

### Performance Analysis Phase

- [ ] **Schedule variance analysis**: Calculate schedule variance (SV) and schedule performance index (SPI)
- [ ] **Cost variance analysis**: Calculate cost variance (CV) and cost performance index (CPI)
- [ ] **Earned value analysis**: Calculate earned value (EV), planned value (PV), and actual cost (AC)
- [ ] **Critical path analysis**: Review critical path and schedule risks
- [ ] **Quality analysis**: Assess quality metrics and trends
- [ ] **Resource analysis**: Evaluate resource performance and constraints

### Health Assessment Phase

- [ ] **Overall health rating**: Assign overall project health status (green/yellow/red)
- [ ] **Trend analysis**: Identify performance trends and patterns
- [ ] **Forecast analysis**: Project future performance based on current trends
- [ ] **Risk assessment**: Evaluate current risk exposure and mitigation effectiveness
- [ ] **Issue impact**: Assess impact of current issues on project objectives

### Documentation Phase

- [ ] **Report creation**: Use [status-report-template.md](./.krci-ai/templates/status-report-template.md) structure
- [ ] **Visualization**: Include charts and graphs for key performance indicators
- [ ] **Content validation**: Ensure all required sections are complete and accurate
- [ ] **Review process**: Review with project team before stakeholder distribution

### Communication Phase

- [ ] **Stakeholder distribution**: Distribute to stakeholders per communication plan
- [ ] **Presentation preparation**: Prepare presentation version if required
- [ ] **Follow-up planning**: Plan follow-up discussions for issues requiring attention
- [ ] **File placement**: Save to exact location `/docs/project-management/status-reports/`

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Accuracy Focus**: Ensure all data and metrics are accurate and current
- **Transparency**: Present both positive and negative performance honestly
- **Actionable Information**: Provide clear actions and recommendations
- **Stakeholder Relevance**: Tailor content to stakeholder needs and interests

### LLM Error Prevention Checklist

- **Avoid**: Outdated or inaccurate performance data
- **Avoid**: Overly technical details that obscure key messages
- **Avoid**: Presenting problems without proposed solutions
- **Avoid**: Missing critical performance indicators or trends
- **Reference**: Use [status-report-template.md](./.krci-ai/templates/status-report-template.md) for all formatting guidance

### PMBoK Integration Context

The Status Report follows PMBoK monitoring and controlling processes by measuring project performance against baselines, tracking progress toward objectives, and communicating project status to stakeholders. This report enables informed decision-making, supports project control activities, and maintains stakeholder engagement throughout the project lifecycle. 