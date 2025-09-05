# Context Gathering Guide: {{project_name}}

> **Purpose**: Structured guide for systematic project context collection
> **Target Output**: `/docs/prd/project-context.md`
> **Framework Integration**: Business methodology-driven approach
> **Quality Standard**: Evidence-based with confidence assessments

---

## Context Collection Methodology Selection

### Available Input Methods

Select appropriate methods based on available resources and project needs:

#### Method 1: Document Analysis

Choose when: Existing research, competitive analysis, user studies, or PRDs available
Required Resources:
- Market research reports
- Competitive analysis documents
- User research and interview data
- Existing PRDs or product specifications
- Technical feasibility studies
- Customer support tickets and feedback
- Analytics and user behavior data

<instructions>
Document analysis method is most effective when substantial existing research is available.

Use this method for:
- Projects with existing market research or user studies
- Follow-up initiatives with prior analysis
- Competitive landscape analysis projects
- Data-rich environments with analytics available

CRITICAL: Always assess document relevance and recency before including in analysis.
</instructions>

**📊 Analysis Framework**:

```markdown
**Document Type**: {{document_type}}
**Source**: {{document_source}}
**Date**: {{document_date}}
**Relevance Score**: {{relevance_score}}/10
**Reliability Score**: {{reliability_score}}/10

**Key Insights Extracted**:
- {{insight_1}}
- {{insight_2}}
- {{insight_3}}

**Supporting Evidence**:
- {{evidence_1}}
- {{evidence_2}}

**Confidence Level**: {{confidence_percentage}}%
**Gaps Identified**: {{knowledge_gaps}}
```

#### Method 2: Stakeholder Interview Process

Choose when: Key stakeholders available for structured input gathering
Required Stakeholders:
- Business stakeholders (product owners, executives)
- User representatives (customer success, support)
- Technical stakeholders (architects, lead developers)
- Market experts (sales, marketing, customer success)

<instructions>
Stakeholder interviews provide primary evidence and diverse perspectives.

Use this method for:
- New initiatives requiring broad stakeholder input
- Complex projects with multiple stakeholder perspectives
- Situations where assumptions need validation
- Projects lacking existing documentation

CRITICAL: Use structured interview guides to ensure consistent data collection across stakeholders.
</instructions>

**🎯 Interview Guide Structure**:

```markdown
**Stakeholder**: {{stakeholder_name}} ({{role}})
**Interview Date**: {{interview_date}}
**Duration**: {{interview_duration}}
**Framework Used**: {{business_framework}} (e.g., BABOK, Design Thinking)

**Opening Context Questions**:
1. How would you describe the current problem/opportunity?
2. What evidence do you have that this problem exists?
3. Who is most affected by this problem?

**Problem Deep Dive** (using 5 Whys framework):
1. What is the core problem?
   - Why does this problem exist?
   - Why does [previous answer] occur?
   - Continue for 5 levels...

**User Context Questions** (using Jobs-to-be-Done framework):
1. When users encounter this problem, what are they trying to accomplish?
2. What is their desired outcome?
3. What current solutions do they use?
4. What gaps exist in current solutions?

**Business Context Questions**:
1. What business impact does this problem have?
2. How do you measure this impact?
3. What constraints exist for solving this problem?
4. What success would look like?

**Evidence and Validation**:
1. What data supports your perspective?
2. How confident are you in this assessment? (1-10)
3. What would change your mind about this?
4. Who else should we talk to about this?

**Key Quotes**:
- "{{important_quote_1}}"
- "{{important_quote_2}}"

**Action Items**:
- {{follow_up_item_1}}
- {{follow_up_item_2}}
```

#### Method 3: Assumption Inventory Creation

Choose when: Limited existing research, need to identify knowledge gaps
Assumption Categories Framework:

<instructions>
Assumption inventory method helps identify and track knowledge gaps systematically.

Use this method for:
- Early-stage projects with limited validation
- High-risk initiatives with many unknowns
- Innovation projects in new markets
- Situations requiring explicit risk identification

CRITICAL: Assign confidence levels and validation methods to all assumptions for effective tracking.
</instructions>

```markdown
## Problem Assumptions
**Category**: Problem Definition and Scope
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}

## User Assumptions
**Category**: Target Users and Behaviors
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}

## Solution Assumptions
**Category**: Solution Approach and Feasibility
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}

## Market Assumptions
**Category**: Market Size and Competitive Landscape
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}

## Business Assumptions
**Category**: Business Model and Success Metrics
- {{assumption_statement}}
  - **Confidence**: {{confidence_level}}%
  - **Impact if False**: {{impact_description}}
  - **Evidence**: {{supporting_evidence}}
  - **Validation Method**: {{suggested_validation}}
```

#### Method 4: Evidence Library Creation

**✅ Choose when**: Quantified data available, need systematic evidence collection
**📊 Evidence Classification Framework**:

```markdown
## Primary Evidence (Highest Confidence)
### {{evidence_type}}: {{evidence_title}}
- **Source**: {{evidence_source}}
- **Collection Date**: {{collection_date}}
- **Sample Size**: {{sample_size}}
- **Methodology**: {{collection_methodology}}
- **Reliability Score**: {{reliability_score}}/10
- **Relevance Score**: {{relevance_score}}/10
- **Recency Score**: {{recency_score}}/10
- **Overall Quality**: {{quality_score}}/10

**Key Data Points**:
- {{data_point_1}}: {{value_1}}
- {{data_point_2}}: {{value_2}}
- {{data_point_3}}: {{value_3}}

**Insights**:
- {{insight_1}}
- {{insight_2}}

**Limitations**:
- {{limitation_1}}
- {{limitation_2}}

## Secondary Evidence (Moderate Confidence)
### {{evidence_type}}: {{evidence_title}}
- **Source**: {{evidence_source}}
- **Publication Date**: {{publication_date}}
- **Relevance Score**: {{relevance_score}}/10
- **Credibility Score**: {{credibility_score}}/10
- **Overall Quality**: {{quality_score}}/10

**Key Insights**:
- {{insight_1}}
- {{insight_2}}

**Supporting Data**:
- {{data_point_1}}
- {{data_point_2}}
```

---

## Business Framework Application Guide

<framework_application>

### Framework Selection by Context Type

#### Problem Context Analysis

Recommended Frameworks:
1. Six Sigma 5 Whys - For root cause identification
2. Impact-Frequency Matrix - For problem prioritization
3. SIPOC Analysis - For process understanding

<instructions>
Framework selection should match the type of analysis needed and available evidence.

Problem Context Guidelines:
- Use 5 Whys when root cause identification is unclear
- Apply Impact-Frequency Matrix when multiple problems need prioritization
- Use SIPOC when process understanding is required

CRITICAL: Document framework rationale and application process for transparency.
</instructions>
</framework_application>

**Application Template**:

```markdown
## Problem Context Analysis

### Framework: {{selected_framework}}
**Why this framework**: {{framework_rationale}}
**Application process**: {{process_steps}}

### Analysis Results:
**Root Cause Identified**: {{root_cause}}
**Problem Priority Score**: {{priority_score}}/25
**Stakeholder Impact Map**: {{stakeholder_impacts}}
**Process Flow Analysis**: {{process_insights}}

### Evidence Supporting Analysis:
- {{evidence_1}} (Confidence: {{confidence_1}}%)
- {{evidence_2}} (Confidence: {{confidence_2}}%)

### Confidence Assessment: {{overall_confidence}}%
```

#### User Context Analysis

**Recommended Frameworks**:
1. **Jobs-to-be-Done** - For user motivation understanding
2. **Design Thinking Empathy Maps** - For user perspective mapping
3. **User Journey Mapping** - For experience flow analysis

**Application Template**:

```markdown
## User Context Analysis

### Framework: {{selected_framework}}
**User Segment**: {{target_user_segment}}
**Job Statement**: When I {{situation}}, I want to {{motivation}}, so I can {{outcome}}

### Analysis Results:
**Functional Job**: {{functional_job}}
**Emotional Job**: {{emotional_job}}
**Social Job**: {{social_job}}
**Job Importance**: {{importance_score}}/10
**Current Satisfaction**: {{satisfaction_score}}/10
**Opportunity Score**: {{opportunity_score}}

### User Journey Stages:
1. **{{stage_1}}**: {{stage_1_analysis}}
2. **{{stage_2}}**: {{stage_2_analysis}}
3. **{{stage_3}}**: {{stage_3_analysis}}

### Pain Points Identified:
- {{pain_point_1}} (Impact: {{impact_1}}/10)
- {{pain_point_2}} (Impact: {{impact_2}}/10)

### Evidence Supporting Analysis:
- {{evidence_1}} (Confidence: {{confidence_1}}%)
- {{evidence_2}} (Confidence: {{confidence_2}}%)

### Confidence Assessment: {{overall_confidence}}%
```

#### Market Context Analysis

**Recommended Frameworks**:
1. **Business Model Canvas** - For value proposition analysis
2. **Porter's Five Forces** - For competitive assessment
3. **TAM/SAM/SOM Analysis** - For market sizing

#### Business Context Analysis

**Recommended Frameworks**:
1. **SWOT Analysis** - For capability and environment assessment
2. **OKR Alignment** - For strategic goal hierarchy
3. **Resource Assessment** - For constraint analysis

---

## Context Synthesis Template

```markdown
# Project Context Summary: {{project_name}}

## Context Collection Summary
- **Methods Used**: {{methods_list}}
- **Timeframe**: {{collection_timeframe}}
- **Stakeholders Involved**: {{stakeholder_count}}
- **Evidence Sources**: {{evidence_source_count}}
- **Overall Confidence**: {{overall_confidence}}%

## Problem Context
### Root Problem Analysis
{{problem_analysis_summary}}

### Problem Evidence
- **Primary Evidence**: {{primary_problem_evidence}}
- **Supporting Data**: {{supporting_problem_data}}
- **Confidence Level**: {{problem_confidence}}%

### Problem Scope and Impact
{{problem_scope_definition}}

## User Context
### Target User Segments
{{user_segment_analysis}}

### User Jobs and Motivations
{{jobs_to_be_done_summary}}

### User Journey and Pain Points
{{user_journey_summary}}

### User Evidence
- **Primary Research**: {{primary_user_evidence}}
- **Supporting Data**: {{supporting_user_data}}
- **Confidence Level**: {{user_confidence}}%

## Market Context
### Market Opportunity
{{market_opportunity_analysis}}

### Competitive Landscape
{{competitive_analysis_summary}}

### Market Evidence
- **Market Data**: {{market_evidence}}
- **Competitive Research**: {{competitive_evidence}}
- **Confidence Level**: {{market_confidence}}%

## Business Context
### Strategic Alignment
{{strategic_alignment_summary}}

### Resource and Constraint Analysis
{{resource_constraint_analysis}}

### Business Evidence
- **Strategic Documentation**: {{strategic_evidence}}
- **Resource Assessment**: {{resource_evidence}}
- **Confidence Level**: {{business_confidence}}%

## Evidence Library Summary
### High Confidence Evidence (80-100%)
{{high_confidence_evidence_list}}

### Medium Confidence Evidence (60-79%)
{{medium_confidence_evidence_list}}

### Low Confidence Evidence (<60%)
{{low_confidence_evidence_list}}

## Assumption Inventory Summary
### High Impact Assumptions (Require Validation)
{{high_impact_assumptions}}

### Medium Impact Assumptions (Monitor)
{{medium_impact_assumptions}}

### Low Impact Assumptions (Accept)
{{low_impact_assumptions}}

## Next Steps for Project Brief Creation
### Ready for Brief Creation
{{ready_areas}}

### Require Additional Context
{{additional_context_needed}}

### Recommended Validation Activities
{{recommended_validations}}

## Quality Assessment
- **Context Completeness**: {{completeness_score}}%
- **Evidence Quality**: {{evidence_quality_score}}%
- **Stakeholder Coverage**: {{stakeholder_coverage_score}}%
- **Framework Application**: {{framework_application_score}}%
- **Overall Context Quality**: {{overall_quality_score}}%
```

---

## Quality Standards and Best Practices

<quality_standards>

### Context Collection Quality Criteria

#### Evidence Quality Standards

- Primary Evidence Target: >70% of key insights from primary sources
- Recency Requirement: Evidence <6 months old for dynamic markets, <12 months for stable markets
- Source Diversity: Minimum 3 different evidence types for key findings
- Confidence Threshold: >60% confidence required for high-impact assumptions

<instructions>
Quality standards ensure context gathering produces reliable, actionable insights.

Evidence Quality Guidelines:
- Prioritize primary evidence (interviews, direct research) over secondary sources
- Weight recent evidence more heavily than historical data
- Validate key findings across multiple source types
- Assign and track confidence levels for all major insights

CRITICAL: Maintain evidence traceability from insights back to original sources.
</instructions>
</quality_standards>

#### Stakeholder Coverage Standards

- **Business Stakeholder**: Minimum 1 executive/decision maker interview
- **User Representative**: Minimum 1 customer-facing role interview
- **Technical Stakeholder**: Minimum 1 technical lead interview
- **Market Expert**: Minimum 1 sales/marketing role interview

#### Framework Application Standards

- **Methodology Documentation**: Clear rationale for framework selection
- **Process Adherence**: Follow framework steps systematically
- **Results Documentation**: Complete analysis with supporting evidence
- **Quality Assessment**: Rate framework application effectiveness

### Context Documentation Standards

#### Structure Requirements

- **Executive Summary**: 1 page maximum context overview
- **Detailed Analysis**: Organized by context type (problem, user, market, business)
- **Evidence Library**: Comprehensive source documentation
- **Assumption Inventory**: Complete assumption tracking

#### Quality Requirements

- **Evidence Attribution**: All insights traced to specific sources
- **Confidence Levels**: All findings rated for reliability
- **Gap Identification**: Areas needing additional context clearly marked
- **Action Items**: Next steps for project brief creation documented

---

## Template Usage Instructions

### When to Use This Guide

- **New Projects**: Starting project brief creation from scratch
- **Context Refresh**: Updating existing project understanding
- **Validation Preparation**: Preparing for business framework validation
- **Stakeholder Alignment**: Building shared understanding of project context

### How to Customize This Guide

- **Select Relevant Methods**: Choose input methods based on available resources
- **Adapt Framework Applications**: Select business frameworks appropriate to project type
- **Customize Evidence Standards**: Adjust quality thresholds based on project risk/importance
- **Tailor Stakeholder Coverage**: Include project-specific stakeholder categories

### Integration with Project Brief Creation

- **Context → Brief**: Direct input for initial project brief creation
- **Evidence → Validation**: Foundation for business framework validation
- **Assumptions → Tracking**: Input for assumption tracker creation
- **Quality → Confidence**: Basis for project brief confidence levels

This guide ensures systematic, high-quality context gathering that enables evidence-based project brief creation with appropriate validation foundation.
