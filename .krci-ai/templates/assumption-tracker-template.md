# Assumption Tracker: {{project_name}}

> **File Location**: `/docs/prd/brief-assumptions.md`
> **Last Updated**: {{last_updated_date}}
> **Total Assumptions**: {{total_assumptions}} | **Validated**: {{validated_count}} | **Pending**: {{pending_count}}

---

## Assumption Summary Dashboard

<summary_dashboard>

| Category | Total | High Risk | Medium Risk | Low Risk | Validated | Pending |
|----------|-------|-----------|-------------|----------|-----------|----------|
| Problem | {{problem_total}} | {{problem_high}} | {{problem_medium}} | {{problem_low}} | {{problem_validated}} | {{problem_pending}} |
| Users | {{users_total}} | {{users_high}} | {{users_medium}} | {{users_low}} | {{users_validated}} | {{users_pending}} |
| Solution | {{solution_total}} | {{solution_high}} | {{solution_medium}} | {{solution_low}} | {{solution_validated}} | {{solution_pending}} |
| Market | {{market_total}} | {{market_high}} | {{market_medium}} | {{market_low}} | {{market_validated}} | {{market_pending}} |
| Business | {{business_total}} | {{business_high}} | {{business_medium}} | {{business_low}} | {{business_validated}} | {{business_pending}} |

<instructions>
Summary dashboard provides high-level metrics across all assumption categories. Track total assumptions, risk distribution, and validation progress for each category.

Categories: Problem, Users, Solution, Market, Business
Risk Levels: High (impact HIGH + confidence <70%), Medium, Low
Status: Validated (confirmed with evidence), Pending (requires validation)
</instructions>
</summary_dashboard>

---

## Problem Assumptions

### High Impact Assumptions

<high_impact_assumptions>
{{#each problem_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

<instructions>
High impact assumptions require immediate attention and detailed validation plans. These assumptions significantly affect project success if proven false.

Required fields for high impact: rationale, evidence, validation method, timeline, owner, risk assessment, validation plan, success criteria.

CRITICAL: High impact assumptions must have confidence levels tracked and validation deadlines enforced.
</instructions>
</high_impact_assumptions>

---
{{/each}}

### Medium Impact Assumptions

<medium_impact_assumptions>
{{#each problem_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

<instructions>
Medium impact assumptions cause delays or scope changes but project remains viable if they fail. Focus on basic validation without extensive risk planning.

Required fields: evidence description, validation method, owner assignment.
</instructions>
</medium_impact_assumptions>

---
{{/each}}

### Low Impact Assumptions

<low_impact_assumptions>
{{#each problem_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
Low impact assumptions have minimal effect on project outcomes. Monitor only, validate when convenient.

Required fields: basic evidence description, confidence level tracking.
</instructions>
</low_impact_assumptions>

---
{{/each}}

## User Assumptions

<user_assumptions>

### High Impact Assumptions

{{#each user_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

### Medium Impact Assumptions

{{#each user_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

### Low Impact Assumptions

{{#each user_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
User assumptions focus on target user behavior, needs, and characteristics. Critical for product-market fit validation.

Validation Methods: user interviews, surveys, behavioral analytics, usability testing, persona validation.

CRITICAL: User assumptions should be validated through direct user contact whenever possible.
</instructions>
</user_assumptions>

## Solution Assumptions

<solution_assumptions>

### High Impact Assumptions

{{#each solution_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

### Medium Impact Assumptions

{{#each solution_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

### Low Impact Assumptions

{{#each solution_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
Solution assumptions cover technical feasibility, implementation approach, and solution effectiveness.

Validation Methods: prototyping, technical spikes, proof of concepts, architecture review, performance testing.

CRITICAL: Solution assumptions should be validated through hands-on technical validation before full implementation.
</instructions>
</solution_assumptions>

## Market Assumptions

<market_assumptions>

### High Impact Assumptions

{{#each market_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

### Medium Impact Assumptions

{{#each market_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

### Low Impact Assumptions

{{#each market_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
Market assumptions cover market size, competition, timing, and market conditions that affect product success.

Validation Methods: market research, competitive analysis, industry reports, customer development, market testing.

CRITICAL: Market assumptions should be validated through external market data and competitive intelligence.
</instructions>
</market_assumptions>

## Business Assumptions

<business_assumptions>

### High Impact Assumptions

{{#each business_high_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: HIGH | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Rationale: {{rationale}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Validation Timeline: {{validation_deadline}}
- Owner: {{assumption_owner}}
- Last Updated: {{last_updated}}

Risk if False: {{risk_if_false}}
Validation Plan: {{validation_plan}}
Success Criteria: {{success_criteria}}

---
{{/each}}

### Medium Impact Assumptions

{{#each business_medium_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: MEDIUM | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}
- Validation Method: {{validation_method}}
- Owner: {{assumption_owner}}

---
{{/each}}

### Low Impact Assumptions

{{#each business_low_assumptions}}

#### {{id}}. {{assumption_statement}}

- Impact: LOW | Confidence: {{confidence_level}}% | Status: {{validation_status}}
- Evidence: {{evidence_description}}

---
{{/each}}

<instructions>
Business assumptions cover revenue model, business viability, resource requirements, and strategic alignment.

Validation Methods: financial modeling, stakeholder interviews, business case analysis, ROI projections, pilot programs.

CRITICAL: Business assumptions should be validated through quantitative financial analysis and stakeholder validation.
</instructions>
</business_assumptions>

## Validation Pipeline

<validation_pipeline>

### High Priority (Immediate Validation Required)

{{#each high_priority_validations}}
- {{assumption_id}}: {{assumption_statement}}
  - Method: {{validation_method}}
  - Timeline: {{validation_deadline}}
  - Owner: {{assumption_owner}}
  - Status: {{validation_status}}
{{/each}}

### Medium Priority (Planned Validation)

{{#each medium_priority_validations}}
- {{assumption_id}}: {{assumption_statement}}
  - Method: {{validation_method}}
  - Timeline: {{validation_deadline}}
  - Owner: {{assumption_owner}}
  - Status: {{validation_status}}
{{/each}}

### Low Priority (Monitor Only)

{{#each low_priority_validations}}
- {{assumption_id}}: {{assumption_statement}}
  - Method: {{validation_method}}
  - Status: {{validation_status}}
{{/each}}

<instructions>
Validation pipeline organizes assumptions by urgency and validation requirements. High priority requires immediate action with deadlines.

Priority Levels:
- High: Immediate validation required, critical to project success
- Medium: Planned validation, affects timeline but not core viability
- Low: Monitor only, minimal project impact

CRITICAL: High priority items must have assigned owners and specific deadlines.
</instructions>
</validation_pipeline>

## Validation Results Log

### Recently Validated Assumptions

{{#each recent_validations}}

#### {{assumption_id}}. {{assumption_statement}}

- **Validation Date**: {{validation_date}}
- **Method Used**: {{validation_method}}
- **Result**: {{validation_result}}
- **Confidence Level**: {{new_confidence_level}}%
- **Evidence**: {{validation_evidence}}
- **Impact on Project Brief**: {{brief_impact}}

**Key Findings**: {{key_findings}}
**Action Items**: {{action_items}}

---
{{/each}}

### Validation Failures/Challenges

{{#each validation_failures}}

#### {{assumption_id}}. {{assumption_statement}}

- **Validation Date**: {{validation_date}}
- **Method Used**: {{validation_method}}
- **Issue**: {{validation_issue}}
- **Next Steps**: {{next_steps}}
- **Revised Timeline**: {{revised_timeline}}

---
{{/each}}

## Assumption Risk Analysis

### High Risk Assumptions (Impact: HIGH, Confidence: <70%)

{{#each high_risk_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Current Confidence**: {{confidence_level}}%
- **Risk Level**: {{risk_level}}
- **Mitigation Plan**: {{mitigation_plan}}
{{/each}}

### Assumptions Requiring Immediate Attention

{{#each immediate_attention_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Issue**: {{attention_reason}}
- **Deadline**: {{attention_deadline}}
- **Owner**: {{assumption_owner}}
{{/each}}

## Evidence Quality Assessment

### Strong Evidence (Confidence 80-100%)

{{#each strong_evidence_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Evidence Type**: {{evidence_type}}
- **Source**: {{evidence_source}}
- **Confidence**: {{confidence_level}}%
{{/each}}

### Moderate Evidence (Confidence 60-79%)

{{#each moderate_evidence_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Evidence Type**: {{evidence_type}}
- **Source**: {{evidence_source}}
- **Confidence**: {{confidence_level}}%
- **Improvement Needed**: {{improvement_plan}}
{{/each}}

### Weak Evidence (Confidence <60%)

{{#each weak_evidence_assumptions}}
- **{{assumption_id}}**: {{assumption_statement}}
- **Evidence Type**: {{evidence_type}}
- **Source**: {{evidence_source}}
- **Confidence**: {{confidence_level}}%
- **Validation Priority**: HIGH
- **Validation Plan**: {{validation_plan}}
{{/each}}

## Next Steps and Action Items

<action_items>

### Immediate Actions (This Week)

{{#each immediate_actions}}
- {{action_item}}
  - Owner: {{action_owner}}
  - Due Date: {{due_date}}
  - Related Assumptions: {{related_assumptions}}
{{/each}}

### Short-term Actions (Next 2 Weeks)

{{#each short_term_actions}}
- {{action_item}}
  - Owner: {{action_owner}}
  - Due Date: {{due_date}}
  - Related Assumptions: {{related_assumptions}}
{{/each}}

### Long-term Actions (Next Month)

{{#each long_term_actions}}
- {{action_item}}
  - Owner: {{action_owner}}
  - Due Date: {{due_date}}
  - Related Assumptions: {{related_assumptions}}
{{/each}}

<instructions>
Action items drive assumption validation forward. Organize by timeline urgency.

Timeline Categories:
- Immediate (This Week): Critical validation activities
- Short-term (Next 2 Weeks): Planned validation work
- Long-term (Next Month): Strategic validation initiatives

Required fields: clear action description, assigned owner, specific due date, related assumption IDs.
</instructions>
</action_items>

---

## Template Usage Notes

<usage_notes>

### Assumption Impact Levels

- HIGH: Assumption failure significantly affects project success or direction
- MEDIUM: Assumption failure causes delays or scope changes but project remains viable
- LOW: Assumption failure has minimal impact on project outcomes

### Confidence Levels

- 80-100%: Strong evidence supporting assumption
- 60-79%: Moderate evidence, some validation needed
- 40-59%: Limited evidence, validation required
- 0-39%: Weak or no evidence, immediate validation critical

### Validation Status

- VALIDATED: Assumption confirmed with strong evidence
- IN_PROGRESS: Validation activities underway
- PLANNED: Validation scheduled but not started
- DEFERRED: Validation postponed due to low priority/impact
- FAILED: Assumption proven false, project brief update needed

### Evidence Types

- PRIMARY: Direct user research, customer interviews, first-party data
- SECONDARY: Industry reports, competitive analysis, expert opinions
- QUANTIFIED: Numerical data, measurements, statistical analysis
- QUALITATIVE: User feedback, expert insights, observational data

<instructions>
Template usage notes provide consistent standards for assumption tracking and validation.

CRITICAL: Use these definitions consistently across all assumption entries to ensure accurate risk assessment and prioritization.

Impact levels determine validation urgency. Confidence levels guide evidence collection needs. Status tracking enables progress monitoring.
</instructions>
</usage_notes>
