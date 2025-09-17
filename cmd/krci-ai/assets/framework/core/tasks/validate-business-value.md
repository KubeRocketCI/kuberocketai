---
dependencies:
  data:
    - business-frameworks.md
    - validation-frameworks.md
  templates:
    - validation-report-template.md
---
# Task: Validate Business Value

## Description

Apply Value Proposition Canvas and ROI calculation frameworks to validate business value proposition, financial justification, and market positioning. This validation ensures the project creates meaningful customer and business value with credible return on investment.

This validation uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and outputs results using the [validation report template](./.krci-ai/templates/validation-report-template.md).

## Prerequisites

<prerequisites>
- Project brief with opportunity/business value section exists
- Access to customer research or feedback data
- Financial data for cost and benefit estimation
- Competitive analysis or market positioning data available
</prerequisites>

## Instructions

<instructions>
Apply Value Proposition Canvas and ROI calculation frameworks to validate business value proposition, financial justification, and market positioning.

### Phase 1: Business Value Hypothesis Extraction

Extract and structure current value proposition from project brief:

#### Value Proposition Analysis

1. Customer Value: Benefits delivered to target customers
2. Business Value: Financial and strategic benefits to organization
3. Market Value: Competitive advantage and market positioning
4. Solution Differentiation: Key differentiators from alternatives

#### Value Hypothesis Structure

```markdown
Business Value Hypothesis: {{project_name}} creates value by delivering {{customer_benefit}} to {{target_customers}}, resulting in {{business_outcome}} for the organization through {{value_mechanism}}.

Specific Claims to Validate:
- Customer value: {{customer_value_claims}}
- Business impact: {{business_impact_claims}}
- Financial returns: {{financial_return_claims}}
- Competitive advantage: {{competitive_advantage_claims}}
```

### Phase 2: Value Proposition Canvas Framework Application

#### Customer Profile Validation

Customer Jobs Analysis:
- Functional Jobs: Practical tasks customers are trying to accomplish
- Emotional Jobs: Feelings customers want to achieve or avoid
- Social Jobs: How customers want to be perceived by others

Customer Pains Identification:
- Problem Pains: Undesired outcomes, problems, and characteristics
- Obstacle Pains: Things preventing customers from getting jobs done
- Risk Pains: Potential negative consequences customers fear

Customer Gains Discovery:
- Required Gains: Basic benefits customers expect
- Expected Gains: Benefits customers expect but would function without
- Desired Gains: Benefits customers want but aren't expecting
- Unexpected Gains: Benefits that go beyond expectations

#### Value Map Validation

Products and Services:
- List specific products/services the value proposition builds on
- Validate alignment with customer jobs, pains, and gains
- Assess completeness and competitive differentiation

Pain Relievers:
- How products/services alleviate specific customer pains
- Validate pain relief effectiveness with customer evidence
- Assess pain relief priority and impact

Gain Creators:
- How products/services create customer gains
- Validate gain creation effectiveness with customer evidence
- Assess gain importance and differentiation

#### Fit Assessment Framework

Problem-Solution Fit:

```markdown
Customer Pain: {{identified_pain}}
Pain Intensity: {{pain_score}}/10 (Evidence: {{pain_evidence}})
Solution Pain Relief: {{pain_reliever_description}}
Relief Effectiveness: {{relief_score}}/10 (Evidence: {{relief_evidence}})
Pain-Solution Fit: {{fit_score}}/10

Overall Problem-Solution Fit: {{average_fit_score}}/10
```

Product-Market Fit:

```markdown
Customer Gain: {{desired_gain}}
Gain Importance: {{gain_score}}/10 (Evidence: {{gain_evidence}})
Solution Gain Creation: {{gain_creator_description}}
Creation Effectiveness: {{creation_score}}/10 (Evidence: {{creation_evidence}})
Gain-Solution Fit: {{fit_score}}/10

Overall Product-Market Fit: {{average_fit_score}}/10
```

### Phase 3: ROI Calculation Validation Framework

#### Investment Cost Analysis

Development Costs:
- Personnel costs (developers, designers, product managers)
- Technology costs (tools, infrastructure, licenses)
- Third-party costs (contractors, consultants, services)
- Opportunity costs (alternative investments foregone)

Implementation Costs:
- Deployment and rollout expenses
- Training and change management costs
- Integration and setup costs
- Testing and quality assurance costs

Operational Costs:
- Ongoing maintenance and support costs
- Infrastructure and hosting costs
- Personnel costs for ongoing operations
- Upgrade and enhancement costs

#### Benefit Quantification

Revenue Benefits:
- New revenue from new customers or markets
- Increased revenue from existing customers
- Revenue protection from competitive threats
- Premium pricing from differentiation

Cost Savings Benefits:
- Process efficiency improvements
- Automation and labor savings
- Reduced error and rework costs
- Vendor consolidation savings

Strategic Benefits:
- Market share gains
- Customer retention improvements
- Brand value enhancement
- Competitive positioning advantages

#### ROI Calculation Methods

Simple ROI Calculation:

```markdown
Total Investment: ${{total_investment}}
Annual Benefits: ${{annual_benefits}}
Simple ROI: {{((annual_benefits - annual_costs) / total_investment * 100)}}%
Payback Period: {{total_investment / (annual_benefits - annual_costs)}} years
```

Net Present Value (NPV):

```markdown
Discount Rate: {{discount_rate}}%
Time Horizon: {{time_horizon}} years
Year 0 Investment: ${{initial_investment}}
Year 1-N Cash Flows: ${{cash_flow_by_year}}
NPV: ${{calculated_npv}}
NPV Interpretation: {{npv_positive_negative_assessment}}
```

Internal Rate of Return (IRR):

```markdown
IRR: {{calculated_irr}}%
Company Hurdle Rate: {{hurdle_rate}}%
IRR Assessment: {{irr_vs_hurdle_assessment}}
```

### Phase 4: Evidence Collection and Validation

#### Customer Evidence Collection

Primary Customer Research:
- Customer interviews validating jobs, pains, and gains
- User testing validating pain relief and gain creation
- Customer willingness-to-pay research
- Customer switching behavior analysis

Secondary Market Research:
- Industry reports on similar value propositions
- Competitive analysis of alternative solutions
- Market research on customer segments and needs
- Expert opinions on market trends and opportunities

#### Financial Evidence Collection

Internal Financial Data:
- Historical cost data for similar projects
- Revenue data for existing customer segments
- Operational cost data for baseline comparison
- Investment criteria and hurdle rates

External Benchmark Data:
- Industry benchmark data for similar initiatives
- Vendor quotes and cost estimates
- Market pricing data for competitive solutions
- Economic data affecting discount rates and projections

### Phase 5: Validation Results Analysis

#### Value Proposition Fit Scoring

Customer Profile Validation Score:

```markdown
Jobs Validation: {{jobs_score}}/10
- Evidence quality: {{jobs_evidence_quality}}
- Completeness: {{jobs_completeness}}
- Customer confirmation: {{jobs_customer_confirmation}}

Pains Validation: {{pains_score}}/10
- Evidence quality: {{pains_evidence_quality}}
- Pain intensity: {{pains_intensity_average}}
- Customer confirmation: {{pains_customer_confirmation}}

Gains Validation: {{gains_score}}/10
- Evidence quality: {{gains_evidence_quality}}
- Gain importance: {{gains_importance_average}}
- Customer confirmation: {{gains_customer_confirmation}}
```

Value Map Validation Score:

```markdown
Products/Services Fit: {{products_score}}/10
Pain Relievers Fit: {{pain_relievers_score}}/10
Gain Creators Fit: {{gain_creators_score}}/10

Overall Value Proposition Fit: {{overall_vp_fit_score}}/10
```

#### Financial Validation Assessment

ROI Validation Quality:

```markdown
Cost Estimation Quality: {{cost_quality_score}}/10
- Estimation method: {{cost_estimation_method}}
- Data sources: {{cost_data_sources}}
- Confidence level: {{cost_confidence_level}}%

Benefit Estimation Quality: {{benefit_quality_score}}/10
- Estimation method: {{benefit_estimation_method}}
- Data sources: {{benefit_data_sources}}
- Confidence level: {{benefit_confidence_level}}%

Overall ROI Confidence: {{roi_confidence_level}}%
Investment Decision Recommendation: {{investment_recommendation}}
```

</instructions>

## Output Format

- Primary Output: `/docs/prd/brief-validation-value.md` using validation report template
- Secondary Outputs:
  - Updated `/docs/prd/project-brief.md` with validated value proposition
  - Updated `/docs/prd/brief-assumptions.md` with value assumptions
  - `/docs/prd/value-proposition-canvas.md` with detailed analysis
  - `/docs/prd/roi-analysis.md` with financial projections
- Length: Validation report 2-3 pages, supporting documents as needed

## Success Criteria

<success_criteria>

### Value Proposition Validation

- Customer profile validated with evidence from customer research
- Value map confirmed with customer feedback on pain relief and gain creation
- Problem-solution fit assessed with quantified fit scores >7/10
- Product-market fit evaluated with evidence of customer demand

### Financial Validation

- ROI calculations completed using multiple methods (Simple, NPV, IRR)
- Cost estimates validated with vendor quotes and historical data
- Benefit projections supported with customer evidence and benchmarks
- Investment recommendation provided based on validated analysis

### Evidence Quality

- Customer evidence collected from primary research with target users
- Financial evidence validated with internal and external benchmark data
- Competitive evidence gathered for market positioning validation
- Assumption tracking updated with value proposition validation results

### Strategic Alignment

- Business case validated with quantified value and credible ROI
- Competitive differentiation confirmed with market analysis
- Strategic fit assessed with organizational goals and capabilities
- Risk assessment completed with value realization dependencies
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Validation Preparation

- Current value analysis: Review existing opportunity/business value sections
- Customer access planning: Arrange customer interviews and research access
- Financial data gathering: Collect internal cost and revenue data
- Competitive research: Analyze alternative solutions and market positioning

### Value Proposition Canvas Application

- Customer profile mapping: Define jobs, pains, and gains with customer input
- Value map creation: Detail products/services, pain relievers, gain creators
- Fit assessment: Evaluate problem-solution and product-market fit
- Evidence validation: Confirm fit assessment with customer feedback

### ROI Analysis

- Cost estimation: Calculate development, implementation, and operational costs
- Benefit quantification: Estimate revenue, cost savings, and strategic benefits
- ROI calculation: Apply multiple calculation methods for validation
- Sensitivity analysis: Test ROI under different scenarios and assumptions

### Evidence Collection and Analysis

- Customer research: Conduct interviews and surveys with target users
- Financial validation: Verify costs and benefits with internal and external data
- Competitive analysis: Assess differentiation and market positioning
- Assumption testing: Validate key value proposition assumptions

### Documentation and Communication

- Validation report creation: Complete comprehensive value validation report
- Brief updates: Update project brief with validated value proposition
- Supporting documentation: Create detailed value proposition canvas and ROI analysis
- Stakeholder presentation: Communicate validation results and investment recommendation

</execution_checklist>

## Content Guidelines

### Value Validation Quality Standards

- Customer-Centric: Ground all value claims in validated customer research
- Evidence-Based: Support all value propositions with quantified evidence
- Financially Rigorous: Use multiple ROI methods with conservative assumptions
- Competitively Aware: Position value relative to alternative solutions
- Risk Transparent: Clearly document assumptions and uncertainty factors

### Customer Research Standards

- Representative Sampling: Include diverse customer perspectives within target segments
- Unbiased Methodology: Use open-ended questions avoiding confirmation bias
- Quantified Insights: Collect measurable data on pain intensity and gain importance
- Switching Behavior: Understand customer willingness to change from current solutions
- Willingness to Pay: Validate value with economic decision-making behavior

### Financial Analysis Standards

- Conservative Assumptions: Use realistic rather than optimistic projections
- Multiple Scenarios: Include best case, worst case, and most likely scenarios
- Time Value Consideration: Apply appropriate discount rates for NPV analysis
- Risk Adjustment: Consider implementation risk and market uncertainty
- Peer Benchmarking: Compare ROI expectations with similar initiatives

## Framework Integration Notes

- SDLC Integration: Validated value proposition informs Epic prioritization and Story value statements
- Business Framework Usage: Leverages [business frameworks](./.krci-ai/data/business-frameworks.md) including Value Proposition Canvas and financial analysis methodologies
- Evidence Standards: Maintains customer-validated and financially rigorous approach
- Quality Assurance: Built-in scoring ensures credible value proposition and ROI analysis
- Professional Output: Investment-grade analysis suitable for executive decision-making
