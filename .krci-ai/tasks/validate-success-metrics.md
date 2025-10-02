---
dependencies:
  data:
    - validation-frameworks.md
    - business-frameworks.md
  templates:
    - validation-report-template.md
---

# Task: Validate Success Metrics

## Description

Apply SMART criteria and OKR alignment frameworks to validate success metrics quality, achievability, and strategic alignment. This validation ensures metrics are specific, measurable, achievable, relevant, and time-bound while supporting organizational objectives.

This validation uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and outputs results using the [validation report template](./.krci-ai/templates/validation-report-template.md).

## Instructions

<instructions>
Confirm project brief with success metrics section exists, baseline data or historical performance data is available, organizational OKR or strategic goals are documented, and industry benchmark data is accessible. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Apply SMART criteria and OKR alignment frameworks from [validation-frameworks.md](./.krci-ai/data/validation-frameworks.md) to validate success metrics quality, achievability, and strategic alignment. Extract current success metrics from project brief including business metrics (revenue, cost savings, market share), user metrics (adoption, engagement, satisfaction), performance metrics (system performance, reliability), and operational metrics (efficiency, productivity). Use [validation-report-template.md](./.krci-ai/templates/validation-report-template.md) for output.
</instructions>

## Metrics Hypothesis Structure

```markdown
Success Metrics Hypothesis: {{project_name}} success measured by achieving:
- Business Impact: {{business_metric}} from {{baseline}} to {{target}} by {{date}}
- User Impact: {{user_metric}} from {{baseline}} to {{target}} by {{date}}
- Performance Impact: {{performance_metric}} from {{baseline}} to {{target}} by {{date}}

Specific Claims to Validate:
- Baseline measurements: {{baseline_claims}}
- Target achievability: {{target_claims}}
- Timeline feasibility: {{timeline_claims}}
- Measurement capability: {{measurement_claims}}
```

### Phase 2: SMART Criteria Validation Framework

#### SMART Assessment Process

For each success metric, evaluate against SMART criteria:

Specific Validation (1-5 scale):
- Is the metric clearly defined without ambiguity?
- Can stakeholders understand what exactly will be measured?
- Is the metric scope bounded and focused?
- Does the metric avoid vague or subjective language?

Measurable Validation (1-5 scale):
- Can the metric be quantified with specific numbers?
- Are measurement tools and data sources identified?
- Is the measurement process documented and reliable?
- Can progress be tracked throughout project lifecycle?

Achievable Validation (1-5 scale):
- Is the target realistic given available resources?
- Are there precedents for similar achievements?
- Have constraints and limitations been considered?
- Is the metric within team/organization control or influence?

Relevant Validation (1-5 scale):
- Does the metric align with business objectives?
- Will achieving this metric create meaningful value?
- Is the metric important to key stakeholders?
- Does the metric contribute to strategic goals?

Time-bound Validation (1-5 scale):
- Is there a clear deadline or timeframe?
- Are there intermediate milestones defined?
- Is the timeline realistic for metric achievement?
- Are dependencies and risks to timeline identified?

#### SMART Score Calculation

```
Total SMART Score = Specific + Measurable + Achievable + Relevant + Time-bound

Interpretation:
- 20-25: Excellent metrics, high confidence to proceed
- 15-19: Good metrics, minor refinement recommended
- 10-14: Moderate metrics, significant improvement needed
- <10: Poor metrics, fundamental revision required
```

### Phase 3: OKR Alignment Validation

#### Organizational Alignment Assessment

Objective Alignment:
- Does project objective align with company/department OKRs?
- Is the objective inspirational and directionally correct?
- Does the objective contribute to higher-level strategic goals?
- Is the objective clear enough to guide decision-making?

Key Result Quality:
- Are key results outcome-focused rather than activity-focused?
- Do key results collectively indicate objective achievement?
- Are key results ambitious yet achievable (60-70% confidence)?
- Do key results avoid measuring inputs or activities?

#### OKR Integration Analysis

```markdown
Company/Department OKR: {{higher_level_okr}}
Project Objective: {{project_objective}}
Alignment Score: {{alignment_score}}/10

Key Results Analysis:
KR1: {{key_result_1}}
- Outcome Focus: {{outcome_score_1}}/5
- Measurability: {{measurability_score_1}}/5
- Ambition Level: {{ambition_score_1}}/5
- Achievement Confidence: {{confidence_score_1}}%

KR2: {{key_result_2}}
- Outcome Focus: {{outcome_score_2}}/5
- Measurability: {{measurability_score_2}}/5
- Ambition Level: {{ambition_score_2}}/5
- Achievement Confidence: {{confidence_score_2}}%

Overall OKR Quality Score: {{okr_total_score}}/50
```

### Phase 4: Leading/Lagging Indicator Analysis

#### Indicator Balance Assessment

Leading Indicators (Predictive metrics):
- Identify metrics that predict future outcomes
- Validate correlation with historical data
- Assess controllability and actionability
- Define measurement frequency and reporting

Lagging Indicators (Outcome metrics):
- Identify metrics that measure final results
- Validate as true measures of success
- Assess measurement accuracy and reliability
- Define measurement and reporting process

#### Balanced Scorecard Creation

```markdown
Leading Indicators (Early Warning System):
1. {{leading_indicator_1}}: {{description}} (Frequency: {{frequency}})
2. {{leading_indicator_2}}: {{description}} (Frequency: {{frequency}})

Lagging Indicators (Success Measurement):
1. {{lagging_indicator_1}}: {{description}} (Frequency: {{frequency}})
2. {{lagging_indicator_2}}: {{description}} (Frequency: {{frequency}})

Predictive Relationships:
- {{leading_indicator}} correlates with {{lagging_indicator}} (R² = {{correlation}})
- Time lag: {{time_lag_period}}
- Confidence in relationship: {{relationship_confidence}}%
```

### Phase 5: Evidence Collection and Validation

#### Baseline Data Validation

Data Collection Requirements:
- Historical performance data for achievability assessment
- Industry benchmark data for target calibration
- Measurement system reliability validation
- Data source accessibility and cost assessment

#### Benchmark Analysis

Industry Benchmarking:
- Identify relevant industry performance benchmarks
- Compare proposed targets with industry standards
- Assess competitive positioning implications
- Validate market expectation alignment

#### Measurement Feasibility Assessment

Technical Validation:
- Verify measurement tools and systems capability
- Assess data collection cost and complexity
- Identify measurement gaps and solutions
- Validate measurement accuracy and reliability

## Output Format

<output_format>
- Primary Output: `/docs/prd/brief-validation-metrics.md` using validation report template
- Secondary Outputs:
  - Updated `/docs/prd/project-brief.md` with validated metrics
  - Updated `/docs/prd/brief-assumptions.md` with metrics assumptions
  - `/docs/prd/metrics-baseline-data.md` with supporting data
- Length: Validation report 2-3 pages, supporting documents as needed
</output_format>

## Success Criteria

<success_criteria>

### SMART Criteria Validation

- All metrics assessed using SMART criteria with scores documented
- SMART scores calculated with improvement recommendations for scores <15
- Baseline data identified for all measurable metrics
- Measurement systems validated for feasibility and reliability

### OKR Alignment Validation

- Strategic alignment confirmed with organizational/departmental OKRs
- Key result quality assessed using outcome-focus criteria
- Ambition level calibrated for 60-70% achievement confidence
- OKR hierarchy documented from company to project level

### Leading/Lagging Balance

- Indicator balance achieved with both leading and lagging metrics
- Predictive relationships validated with historical correlation data
- Measurement frequency defined for each indicator type
- Early warning system established through leading indicators

### Evidence and Benchmarking

- Baseline data collected and quality assessed
- Industry benchmarks identified and targets calibrated
- Measurement feasibility confirmed with cost and complexity assessment
- Success thresholds defined with clear criteria
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Validation Preparation

- Current metrics review: Analyze existing success metrics in project brief
- Data source identification: Identify available baseline and benchmark data
- Stakeholder coordination: Connect with teams responsible for measurement
- OKR documentation: Gather organizational and departmental OKR information

### SMART Criteria Assessment

- Individual metric evaluation: Apply SMART criteria to each metric
- Scoring documentation: Record scores and rationale for each criterion
- Gap identification: Identify metrics failing SMART criteria
- Improvement recommendations: Suggest specific metric enhancements

### OKR Alignment Validation

- Alignment mapping: Document connection to higher-level OKRs
- Key result analysis: Evaluate each key result for quality
- Ambition calibration: Assess and adjust target ambition levels
- Confidence assessment: Rate achievement probability for each metric

### Evidence Collection

- Baseline data gathering: Collect historical performance data
- Benchmark research: Identify relevant industry performance standards
- Measurement validation: Verify measurement system capabilities
- Cost assessment: Evaluate measurement cost and complexity

### Documentation and Reporting

- Validation report creation: Complete comprehensive validation report
- Brief updates: Update project brief with validated metrics
- Assumption tracking: Update assumption tracker with metrics validation results
- Stakeholder communication: Present validation results to project team
</execution_checklist>

## Content Guidelines

### Metrics Quality Standards

- Quantified Focus: Prioritize metrics with specific numerical targets
- Evidence-Based Targets: Ground all targets in baseline data and benchmarks
- Balanced Portfolio: Include both leading (predictive) and lagging (outcome) indicators
- Stakeholder Relevance: Ensure metrics matter to key decision makers
- Measurement Feasibility: Verify realistic measurement capability and cost

### Validation Methodology Standards

- Systematic Assessment: Apply SMART criteria consistently across all metrics
- Evidence Requirements: Support all assessments with documented evidence
- Benchmark Integration: Use industry data to calibrate targets appropriately
- Stakeholder Validation: Confirm metrics relevance with key stakeholders
- Continuous Monitoring: Establish ongoing measurement and review processes

### Documentation Standards

- Clear Scoring: Document all SMART and OKR scores with explicit rationale
- Evidence Citation: Properly attribute all baseline data and benchmarks
- Improvement Recommendations: Provide specific guidance for metric enhancement
- Measurement Planning: Document how each metric will be tracked and reported
- Success Criteria: Define clear thresholds for success and failure

## Framework Integration Notes

- SDLC Integration: Validated metrics inform Epic success criteria and Story acceptance criteria
- Business Framework Usage: Leverages SMART methodology and OKR framework for systematic validation
- Evidence Standards: Maintains quantified, benchmark-driven approach for credible targets
- Quality Assurance: Built-in scoring ensures metrics support effective project management
- Professional Output: Structured documentation supports stakeholder decision-making and tracking
