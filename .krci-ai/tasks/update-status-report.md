# Task: Update Status Report

## Description

Create an updated project status report that builds on previous reports to show project progress, performance trends, and current health status. Following PMBoK continuous monitoring practices, updated status reports provide ongoing communication to stakeholders with focus on changes since the last report, performance trends, and corrective actions taken or needed.

## Prerequisites

- [ ] Previous status report available for comparison and trend analysis
- [ ] Current project performance data collected since last report
- [ ] Updated project baselines if changes were approved
- [ ] Risk register updated at `/docs/project-management/risk-register.md`
- [ ] Issues log current with latest status updates

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/project-management-methodology.md
- ./.krci-ai/templates/status-report-template.md
- /docs/project-management/project-plan.md (approved plan)
- /docs/project-management/risk-register.md (current register)
- /docs/project-management/status-reports/ (previous reports for trend analysis)

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Review previous report**: Analyze last status report for baseline comparison
2. **Follow PMBoK monitoring processes**: Apply continuous monitoring and controlling practices
3. **Trend analysis focus**: Emphasize changes and trends since last reporting period
4. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for artifact dependencies
5. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
6. **Format output**: Use [status-report-template.md](./.krci-ai/templates/status-report-template.md) for consistent structure

## Output Format

- **Location**: `/docs/project-management/status-reports/status-report-YYYY-MM-DD.md` (EXACT path and filename)
- **Naming**: Use current date for new report filename
- **Change focus**: Emphasize changes and trends since previous report

## Success Criteria

- [ ] **File saved** to `/docs/project-management/status-reports/status-report-YYYY-MM-DD.md`
- [ ] **Previous report reviewed** for baseline and trend comparison
- [ ] **Change summary** highlights key changes since last report
- [ ] **Performance trends** analyzed and documented
- [ ] **Variance trends** tracked for schedule and cost performance
- [ ] **Risk changes** documented with new risks and closed risks
- [ ] **Issue progression** tracked with resolution updates
- [ ] **Corrective actions** documented with implementation status
- [ ] **Forecast updates** based on current trends and performance
- [ ] **Stakeholder actions** updated with new assignments and completed items

## Execution Checklist

### Previous Report Analysis Phase

- [ ] **Baseline comparison**: Compare current performance against previous report
- [ ] **Trend identification**: Identify performance trends since last report
- [ ] **Action item review**: Review status of action items from previous report
- [ ] **Forecast accuracy**: Evaluate accuracy of previous forecasts

### Current Data Collection Phase

- [ ] **New schedule data**: Collect schedule progress since last report
- [ ] **New cost data**: Gather cost performance data for current period
- [ ] **New scope progress**: Track deliverable completion since last report
- [ ] **Updated resource data**: Monitor resource changes and utilization
- [ ] **Risk register changes**: Review risk updates since last report
- [ ] **Issue log updates**: Track issue status changes and new issues

### Change Analysis Phase

- [ ] **Performance delta**: Calculate changes in key performance indicators
- [ ] **Variance trends**: Analyze trends in schedule and cost variances
- [ ] **Risk profile changes**: Assess changes in overall risk exposure
- [ ] **Quality trends**: Evaluate quality metric changes and trends
- [ ] **Milestone progress**: Track milestone achievement since last report

### Forecast Update Phase

- [ ] **Trend projection**: Project future performance based on current trends
- [ ] **Completion forecasting**: Update estimated completion date and cost
- [ ] **Risk impact forecasting**: Assess potential future impact of current risks
- [ ] **Resource forecasting**: Project future resource needs and availability

### Documentation Phase

- [ ] **Report creation**: Use [status-report-template.md](./.krci-ai/templates/status-report-template.md) structure
- [ ] **Change highlighting**: Clearly highlight changes since previous report
- [ ] **Trend visualization**: Include trend charts and performance graphs
- [ ] **Content validation**: Ensure accuracy and completeness of all updates

### Communication Phase

- [ ] **Stakeholder distribution**: Distribute to stakeholders per communication plan
- [ ] **Change communication**: Clearly communicate significant changes to stakeholders
- [ ] **Action item follow-up**: Follow up on action items from previous report
- [ ] **File placement**: Save to exact location `/docs/project-management/status-reports/`

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Change Focus**: Emphasize what has changed since the last report
- **Trend Analysis**: Provide meaningful analysis of performance trends
- **Continuity**: Maintain consistency with previous report format and metrics
- **Forward Looking**: Include updated forecasts and future outlook

### LLM Error Prevention Checklist

- **Avoid**: Repeating identical content from previous reports without updates
- **Avoid**: Missing trend analysis or change identification
- **Avoid**: Inconsistent metrics or reporting formats between reports
- **Avoid**: Outdated action items or incomplete follow-up status
- **Reference**: Use [status-report-template.md](./.krci-ai/templates/status-report-template.md) and previous reports

### PMBoK Integration Context

Updated status reports follow PMBoK continuous monitoring by tracking performance trends, measuring progress against baselines, evaluating corrective action effectiveness, and maintaining ongoing stakeholder communication. These reports support project control decisions and enable proactive management of project performance and risks. 