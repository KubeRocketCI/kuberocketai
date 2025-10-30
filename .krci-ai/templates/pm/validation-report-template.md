# Validation Report: {{project_name}}

> File Location: `/docs/prd/brief-validation.md`
> Validation Date: {{validation_date}}
> Report Type: {{report_type}} | Framework Used: {{validation_framework}}
> Overall Confidence: {{overall_confidence}}% | Status: {{validation_status}}

---

## Validation Summary

<validation_summary>

### Validation Scope

- Section Validated: {{section_name}}
- Business Framework: {{business_framework}}
- Validation Method: {{validation_method}}
- Evidence Sources: {{evidence_sources}}
- Stakeholders Involved: {{stakeholders_involved}}

### Key Findings

{{#each key_findings}}
- {{finding_type}}: {{finding_description}}
{{/each}}

### Overall Assessment

- Pre-Validation Confidence: {{pre_validation_confidence}}%
- Post-Validation Confidence: {{post_validation_confidence}}%
- Confidence Change: {{confidence_change}} ({{confidence_direction}})
- Validation Status: {{final_status}}

<instructions>
This section provides a concise overview of the validation process and results.

Validation Scope: Document exactly what was validated and how
Key Findings: List the most important discoveries from the validation
Overall Assessment: Show before/after confidence levels and final status

CRITICAL: Keep summary to 3-4 bullet points maximum for executive readability.
</instructions>
</validation_summary>

---

## Detailed Validation Results

<detailed_results>

### {{validation_framework}} Framework Application

#### Framework Methodology

Framework: {{framework_name}}
Source: {{framework_source}}
Application Context: {{application_context}}

#### Validation Process Followed

{{#each validation_steps}}
{{step_number}}. {{step_name}}
- Method: {{step_method}}
- Evidence Collected: {{step_evidence}}
- Findings: {{step_findings}}
- Confidence Level: {{step_confidence}}%
{{/each}}

<instructions>
Document the specific business framework used and how it was applied systematically.

Framework Application: Show which methodology was used and why it was chosen
Process Steps: Document each step taken with evidence and confidence levels

CRITICAL: Include confidence percentages for each step to track validation quality.
</instructions>
</detailed_results>

#### Framework-Specific Results

{{#if is_problem_validation}}

##### Lean Startup Problem-Solution Fit Results

- Problem Hypothesis: {{problem_hypothesis}}
- Problem Validation Score: {{problem_score}}/10
- Solution Hypothesis: {{solution_hypothesis}}
- Solution Validation Score: {{solution_score}}/10
- Overall Problem-Solution Fit: {{problem_solution_fit}}

Evidence Summary:
- Customer Interview Data: {{customer_interviews}} interviews ({{interview_confidence}}% confidence)
- Quantified Problem Metrics: {{problem_metrics}} ({{metrics_confidence}}% confidence)
- Competitive Analysis: {{competitive_analysis}} ({{competitive_confidence}}% confidence)
- Technical Feasibility: {{technical_feasibility}} ({{technical_confidence}}% confidence)
{{/if}}

{{#if is_user_validation}}

##### Jobs-to-be-Done Validation Results

- Primary Job Statement: {{job_statement}}
- Job Importance Score: {{job_importance}}/10
- Current Satisfaction Score: {{job_satisfaction}}/10
- Opportunity Score: {{opportunity_score}}

Job Analysis:
- Functional Job: {{functional_job}}
- Emotional Job: {{emotional_job}}
- Social Job: {{social_job}}

Evidence Summary:
- User Interview Data: {{user_interviews}} interviews ({{user_confidence}}% confidence)
- Usage Analytics: {{usage_analytics}} (Period: {{analytics_period}}, {{analytics_confidence}}% confidence)
- Market Research: {{market_research}} (Source: {{market_source}}, {{market_confidence}}% confidence)
{{/if}}

{{#if is_metrics_validation}}

##### SMART Criteria Assessment Results

{{#each smart_metrics}}
Metric: {{metric_name}}
- Specific: {{specific_score}}/5 - {{specific_assessment}}
- Measurable: {{measurable_score}}/5 - {{measurable_assessment}}
- Achievable: {{achievable_score}}/5 - {{achievable_assessment}}
- Relevant: {{relevant_score}}/5 - {{relevant_assessment}}
- Time-bound: {{timebound_score}}/5 - {{timebound_assessment}}
- Overall SMART Score: {{smart_total_score}}/25

Evidence Summary:
- Baseline Data: {{baseline_data}} ({{baseline_confidence}}% confidence)
- Industry Benchmarks: {{industry_benchmarks}} ({{benchmark_confidence}}% confidence)
- Historical Performance: {{historical_data}} ({{historical_confidence}}% confidence)
{{/each}}
{{/if}}

{{#if is_value_validation}}

##### Value Proposition Canvas Results

Customer Profile Validation:
- Customer Jobs: {{customer_jobs}} (Validation: {{jobs_validation}}%)
- Customer Pains: {{customer_pains}} (Validation: {{pains_validation}}%)
- Customer Gains: {{customer_gains}} (Validation: {{gains_validation}}%)

Value Map Validation:
- Products & Services: {{products_services}}
- Pain Relievers: {{pain_relievers}} (Fit: {{pain_fit}}%)
- Gain Creators: {{gain_creators}} (Fit: {{gain_fit}}%)

Fit Assessment:
- Problem-Solution Fit: {{problem_solution_fit}}%
- Product-Market Fit: {{product_market_fit}}%
- Overall Value Proposition Fit: {{value_prop_fit}}%

Evidence Summary:
- Customer Research: {{customer_research}} ({{research_confidence}}% confidence)
- Solution Testing: {{solution_testing}} ({{testing_confidence}}% confidence)
- Willingness to Pay: {{willingness_to_pay}} ({{wtp_confidence}}% confidence)
{{/if}}

---

## Evidence Analysis

<evidence_analysis>

### Primary Evidence Sources

{{#each primary_evidence}}

#### {{evidence_type}}: {{evidence_name}}

- Source: {{evidence_source}}
- Collection Method: {{collection_method}}
- Sample Size: {{sample_size}}
- Time Period: {{time_period}}
- Reliability: {{reliability_score}}/10
- Relevance: {{relevance_score}}/10
- Recency: {{recency_score}}/10
- Overall Quality Score: {{quality_score}}/10

Key Insights:
{{#each insights}}
- {{insight_description}}
{{/each}}

Supporting Quotes/Data:
{{#each supporting_data}}
- {{data_point}}
{{/each}}
{{/each}}

<instructions>
Document all evidence sources with quality assessments and key findings.

Primary Evidence: Direct research, customer interviews, first-party data
Quality Scoring: Rate each source on reliability, relevance, and recency (1-10)
Key Insights: Extract the most important findings from each evidence source

CRITICAL: Primary evidence should be weighted more heavily than secondary sources.
</instructions>
</evidence_analysis>

### Secondary Evidence Sources

{{#each secondary_evidence}}

#### {{evidence_type}}: {{evidence_name}}

- Source: {{evidence_source}}
- Publication Date: {{publication_date}}
- Relevance: {{relevance_score}}/10
- Credibility: {{credibility_score}}/10
- Overall Quality Score: {{quality_score}}/10

Key Insights:
{{#each insights}}
- {{insight_description}}
{{/each}}
{{/each}}

### Evidence Gaps Identified

{{#each evidence_gaps}}
- Gap: {{gap_description}}
- Impact on Validation: {{gap_impact}}
- Recommended Action: {{gap_recommendation}}
- Priority: {{gap_priority}}
{{/each}}

---

## Assumption Updates

<assumption_updates>

### Assumptions Validated

{{#each validated_assumptions}}

#### {{assumption_id}}: {{assumption_statement}}

- Pre-Validation Confidence: {{pre_confidence}}%
- Post-Validation Confidence: {{post_confidence}}%
- Status Change: {{old_status}} → {{new_status}}
- Supporting Evidence: {{validation_evidence}}
{{/each}}

### Assumptions Challenged/Modified

{{#each challenged_assumptions}}

#### {{assumption_id}}: {{original_assumption}}

- Issue Identified: {{challenge_description}}
- Revised Assumption: {{revised_assumption}}
- Confidence Change: {{confidence_change}}
- Impact on Project Brief: {{brief_impact}}
{{/each}}

<instructions>
Track changes to project assumptions based on validation findings.

Validated Assumptions: Show confidence level improvements with supporting evidence
Challenged Assumptions: Document what was wrong and how assumptions were revised
Status Changes: Use clear before/after format to show assumption evolution

CRITICAL: All assumption changes must link back to specific evidence sources.
</instructions>
</assumption_updates>

### New Assumptions Identified

{{#each new_assumptions}}

#### {{assumption_id}}: {{assumption_statement}}

- Source: {{assumption_source}}
- Initial Confidence: {{initial_confidence}}%
- Impact Level: {{impact_level}}
- Validation Priority: {{validation_priority}}
- Recommended Validation Method: {{recommended_method}}
{{/each}}

---

## Validation Quality Assessment

### Methodology Quality

- Framework Appropriateness: {{methodology_appropriateness}}/10
- Process Adherence: {{process_adherence}}/10
- Evidence Rigor: {{evidence_rigor}}/10
- Stakeholder Involvement: {{stakeholder_involvement}}/10
- Overall Methodology Score: {{methodology_total}}/40

### Evidence Quality

- Source Diversity: {{source_diversity}}/10
- Sample Size Adequacy: {{sample_adequacy}}/10
- Data Recency: {{data_recency}}/10
- Evidence Relevance: {{evidence_relevance}}/10
- Overall Evidence Score: {{evidence_total}}/40

### Results Reliability

- Internal Consistency: {{internal_consistency}}/10
- Cross-Source Validation: {{cross_validation}}/10
- Stakeholder Agreement: {{stakeholder_agreement}}/10
- Bias Mitigation: {{bias_mitigation}}/10
- Overall Reliability Score: {{reliability_total}}/40

Combined Validation Quality Score: {{total_quality_score}}/120

---

## Impact on Project Brief

### Required Updates

{{#each brief_updates}}

#### {{section_name}} Section

- Current Content: {{current_content}}
- Recommended Change: {{recommended_change}}
- Rationale: {{change_rationale}}
- Priority: {{update_priority}}
{{/each}}

### Confidence Level Changes

{{#each confidence_changes}}
- Section: {{section_name}}
- Previous Confidence: {{previous_confidence}}%
- Updated Confidence: {{updated_confidence}}%
- Change Reason: {{change_reason}}
{{/each}}

### New Risks/Considerations Identified

{{#each new_risks}}
- Risk: {{risk_description}}
- Impact: {{risk_impact}}
- Probability: {{risk_probability}}
- Mitigation: {{risk_mitigation}}
- Brief Section Affected: {{affected_section}}
{{/each}}

---

## Recommendations and Next Steps

<recommendations>
### Immediate Actions Required
{{#each immediate_actions}}
- {{action_item}}
  - Rationale: {{action_rationale}}
  - Owner: {{action_owner}}
  - Due Date: {{due_date}}
  - Impact if Not Done: {{inaction_impact}}
{{/each}}

### Additional Validation Recommended

{{#each additional_validations}}
- Area: {{validation_area}}
- Recommended Framework: {{recommended_framework}}
- Priority: {{validation_priority}}
- Timeline: {{validation_timeline}}
- Expected Outcome: {{expected_outcome}}
{{/each}}

<instructions>
Provide clear, actionable recommendations based on validation findings.

Immediate Actions: List specific tasks that must be completed with owners and dates
Additional Validation: Identify areas needing further validation with recommended methods
Priority Levels: Use High/Medium/Low to help stakeholders focus efforts

CRITICAL: Every recommendation must have a clear rationale tied to validation findings.
</instructions>
</recommendations>

### Project Brief Refinement Suggestions

{{#each refinement_suggestions}}
- Section: {{section_name}}
- Suggestion: {{suggestion_description}}
- Benefit: {{suggestion_benefit}}
- Effort Required: {{effort_level}}
{{/each}}

---

## Validation Timeline and Planning

### Completed Validation Activities

{{#each completed_activities}}
- {{activity_date}}: {{activity_description}}
  - Outcome: {{activity_outcome}}
  - Time Invested: {{time_invested}}
  - Quality Score: {{activity_quality}}/10
{{/each}}

### Planned Follow-up Validations

{{#each planned_validations}}
- Activity: {{planned_activity}}
- Scheduled Date: {{scheduled_date}}
- Expected Duration: {{expected_duration}}
- Responsible Party: {{responsible_party}}
- Success Criteria: {{success_criteria}}
{{/each}}

### Validation Backlog

{{#each validation_backlog}}
- Priority {{priority_level}}: {{backlog_item}}
  - Estimated Effort: {{estimated_effort}}
  - Expected Value: {{expected_value}}
  - Dependencies: {{dependencies}}
{{/each}}

---

## Stakeholder Feedback

### Validation Process Feedback

{{#each process_feedback}}
{{stakeholder_name}} ({{stakeholder_role}}):
> "{{feedback_quote}}"

- Process Rating: {{process_rating}}/10
- Value Rating: {{value_rating}}/10
- Suggestions: {{stakeholder_suggestions}}
{{/each}}

### Results Feedback

{{#each results_feedback}}
{{stakeholder_name}} ({{stakeholder_role}}):
- Agreement with Findings: {{agreement_level}}/10
- Confidence in Results: {{confidence_level}}/10
- Key Concerns: {{stakeholder_concerns}}
- Additional Insights: {{additional_insights}}
{{/each}}

---

## Lessons Learned

### What Worked Well

{{#each what_worked}}
- {{success_area}}: {{success_description}}
{{/each}}

### What Could Be Improved

{{#each improvements}}
- {{improvement_area}}: {{improvement_description}}
  - Suggested Solution: {{suggested_solution}}
{{/each}}

### Recommendations for Future Validations

{{#each future_recommendations}}
- {{recommendation_category}}: {{recommendation_description}}
{{/each}}

---

## Appendices

### A. Detailed Evidence Library

{{#each detailed_evidence}}

#### {{evidence_id}}: {{evidence_title}}

Source: {{source_details}}
Collection Date: {{collection_date}}
Methodology: {{collection_methodology}}

Full Evidence:
{{evidence_content}}

Analysis Notes:
{{analysis_notes}}
{{/each}}

### B. Stakeholder Interview Transcripts

{{#each interview_transcripts}}

#### Interview: {{interviewee_name}} ({{interviewee_role}})

Date: {{interview_date}}
Duration: {{interview_duration}}
Interviewer: {{interviewer_name}}

Key Quotes:
{{#each key_quotes}}
- "{{quote_text}}" (Context: {{quote_context}})
{{/each}}

Full Transcript: [Link to {{transcript_file}}]
{{/each}}

### C. Validation Framework Resources

- Framework Documentation: {{framework_documentation_link}}
- Methodology References: {{methodology_references}}
- Best Practices: {{best_practices_link}}
- Training Materials: {{training_materials_link}}

---

<!-- REPORT METADATA
Report ID: {{report_id}}
Generated By: {{report_generator}}
Framework Version: {{framework_version}}
Template Version: {{template_version}}
Last Modified: {{last_modified}}
Next Review Date: {{next_review_date}}
-->

## Template Usage Guidelines

### When to Use This Template

- After completing any business framework validation
- When validation results significantly impact project brief confidence
- For stakeholder communication about validation outcomes
- As input for project brief refinement decisions

### Report Types

- **INITIAL**: First validation of a project brief section
- **FOLLOW_UP**: Additional validation to address gaps or concerns
- **COMPREHENSIVE**: Complete validation using multiple frameworks
- **TARGETED**: Focused validation of specific assumptions or evidence gaps

### Quality Scoring

- **1-3**: Poor quality, unreliable for decision making
- **4-6**: Moderate quality, useable with caveats
- **7-8**: Good quality, reliable for most decisions
- **9-10**: Excellent quality, high confidence for all decisions

### Evidence Classification

- **PRIMARY**: Direct research, first-party data, customer feedback
- **SECONDARY**: Industry reports, competitive analysis, expert opinions
- **QUANTIFIED**: Numerical data, measurements, statistical analysis
- **QUALITATIVE**: Descriptive data, observations, subjective assessments
