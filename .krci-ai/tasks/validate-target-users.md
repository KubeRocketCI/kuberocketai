# Task: Validate Target Users

## Description

Apply Jobs-to-be-Done framework to validate target user segments, their motivations, and the value proposition alignment. This validation ensures user segments are accurately defined, their needs are properly understood, and the solution approach resonates with their actual jobs, pains, and gains.

This validation uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and outputs results using the [validation report template](./.krci-ai/templates/validation-report-template.md).

## Prerequisites

<prerequisites>
- Project brief with target users section exists
- Access to target users for interviews or surveys
- User analytics or behavioral data available
- Market segmentation data or competitive user research
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/validation-frameworks.md (Jobs-to-be-Done section)
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/validation-report-template.md
- /docs/prd/project-brief.md (source document)

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

<instructions>
Apply Jobs-to-be-Done framework from [validation frameworks](./.krci-ai/data/validation-frameworks.md) to validate target user segments, their motivations, and the value proposition alignment.

### Phase 1: User Segment Hypothesis Extraction

Extract and structure current user segment definitions from project brief:

#### Current User Segment Analysis

1. **Primary User Segment**: Extract demographics, behaviors, needs
2. **Secondary User Segments**: Identify additional user groups mentioned
3. **User Context**: Document when and where users encounter the problem
4. **User Goals**: List stated user objectives and success criteria

#### User Hypothesis Structure

```markdown
**Primary User Segment Hypothesis**: {{user_demographic}} who {{user_behavior}} because they need to {{user_goal}}.

**Specific Claims to Validate**:
- User segment size: {{segment_size_claim}}
- User behavior patterns: {{behavior_pattern_claims}}
- User needs and pain points: {{needs_pain_claims}}
- User willingness to adopt solutions: {{adoption_claims}}
```

### Phase 2: Jobs-to-be-Done Framework Application

#### Step 1: Job Statement Definition and Validation

**Job Statement Construction**:
For each user segment, create job statements in the format:
"When I [situation], I want to [motivation], so I can [expected outcome]"

**Job Dimensions Analysis**:
- **Functional Job**: The practical task the user is trying to accomplish
- **Emotional Job**: How the user wants to feel or avoid feeling
- **Social Job**: How the user wants to be perceived by others

**Job Statement Validation Process**:
- Conduct user interviews to validate job statements
- Assess job importance (1-10) through user ranking exercises
- Evaluate current satisfaction (1-10) with existing solutions
- Calculate opportunity score: Importance + max(Importance - Satisfaction, 0)

#### Step 2: User Journey and Touchpoint Validation

**User Journey Mapping**:
- Map complete user journey from problem awareness to solution adoption
- Identify all touchpoints where users interact with current solutions
- Document user actions, thoughts, emotions at each stage
- Validate journey accuracy through user observation or interviews

**Pain Point Identification and Validation**:
- Identify specific pain points at each journey stage
- Quantify pain point impact (time lost, money spent, frustration level)
- Validate pain point significance through user prioritization exercises
- Assess pain point frequency and universality across user segment

**Opportunity Identification**:
- Identify improvement opportunities at each journey stage
- Assess opportunity impact on user satisfaction and business value
- Validate opportunity desirability through user feedback
- Prioritize opportunities by impact and feasibility

#### Step 3: User Segment Validation

**Demographic and Behavioral Validation**:
- Validate user segment demographics through market research data
- Confirm behavioral patterns through analytics and user observation
- Assess segment size and growth trends through industry data
- Validate segment accessibility and reachability for marketing

**Needs and Goals Validation**:
- Confirm user needs through structured interviews and surveys
- Validate goal priorities through user ranking and trade-off exercises
- Assess need intensity and urgency through user behavior analysis
- Confirm alignment between stated and revealed preferences

### Phase 3: Evidence Collection and Analysis

#### Primary Evidence Collection

**User Interview Program**:
- Conduct minimum 8-10 interviews per primary user segment
- Use structured interview guide based on Jobs-to-be-Done methodology
- Document specific examples and quantified impacts
- Record user quotes and insights for validation report

**User Survey Deployment**:
- Design survey to quantify job importance and satisfaction ratings
- Include demographic and behavioral questions for segmentation
- Achieve statistically significant sample size (minimum 100 responses)
- Analyze results for segment patterns and validation insights

**User Analytics Analysis**:
- Analyze existing user behavior data to validate segment assumptions
- Identify usage patterns that support or challenge segment definitions
- Quantify user engagement and retention metrics by segment
- Validate user journey stages through behavioral flow analysis

#### Secondary Evidence Collection

**Market Research Integration**:
- Leverage existing market segmentation studies and reports
- Identify industry benchmarks for similar user segments
- Validate segment size estimates through multiple data sources
- Assess competitive positioning targeting similar segments

**Competitive User Analysis**:
- Analyze competitor user bases and targeting strategies
- Identify overlapping segments and positioning differences
- Assess competitive user satisfaction and switching behavior
- Validate user needs through competitive feature analysis

### Phase 4: Jobs-to-be-Done Analysis Framework

#### Job Importance and Satisfaction Scoring

**Importance Assessment (1-10 scale)**:
- How important is this job for the user's success?
- How frequently does the user need to complete this job?
- What's the impact on the user if this job isn't completed well?
- How much time/resources does the user invest in this job?

**Satisfaction Assessment (1-10 scale)**:
- How satisfied is the user with current solution options?
- How well do current solutions complete the job?
- What gaps exist in current solution capabilities?
- How much effort is required to use current solutions?

**Opportunity Score Calculation**:

```
Opportunity Score = Importance + max(Importance - Satisfaction, 0)

Interpretation:
- 15+: High opportunity (underserved job with high importance)
- 12-15: Moderate opportunity (room for improvement exists)
- <12: Low opportunity (adequately served or low importance)
```

#### User Segment Validation Assessment

**Segment Definition Quality**:
- **Specificity**: Can you clearly identify who belongs to this segment?
- **Measurability**: Can you quantify segment size and characteristics?
- **Accessibility**: Can you reach this segment through marketing channels?
- **Actionability**: Can you create distinct value propositions for this segment?

**Evidence Quality Assessment**:
- **Primary Research Percentage**: How much validation comes from direct user research?
- **Sample Size Adequacy**: Is the research sample representative and significant?
- **Evidence Recency**: How current is the supporting evidence?
- **Cross-Source Validation**: Are findings consistent across multiple evidence types?

### Phase 5: Validation Results and Recommendations

#### User Segment Validation Results

**Validated Segments**:
- Segments supported by strong evidence across multiple sources
- Clear job definitions with high importance and low satisfaction scores
- Measurable segment characteristics and accessible through known channels
- Strong alignment between user needs and proposed solution approach

**Challenged Segments**:
- Segments with weak or conflicting evidence
- Unclear or low-importance job definitions
- Difficult to measure or access through available channels
- Poor alignment between stated needs and solution approach

**New Segments Identified**:
- Additional user groups discovered during validation research
- Segments with high opportunity scores not previously considered
- Adjacent segments with similar jobs but different contexts
- Segments that could be served with minor solution modifications

#### Project Brief Update Recommendations

**User Section Enhancements**:
- Update user segment definitions with validated evidence
- Add quantified user data (segment sizes, behavior patterns, satisfaction levels)
- Include validated job statements and opportunity scores
- Enhance user journey understanding with pain point and opportunity details

**Confidence Level Updates**:
- Assign confidence levels to all user segment claims based on evidence quality
- Identify areas requiring additional validation or research
- Update assumption tracker with validated, challenged, and new assumptions
- Prioritize remaining user validation activities
</instructions>

## Output Format

<output_format>
- Primary Output: `/docs/prd/brief-validation-users.md` using validation report template
- Secondary Outputs:
  - Updated `/docs/prd/project-brief.md` with validated user information
  - Updated `/docs/prd/brief-assumptions.md` with user assumption status
  - `/docs/prd/user-research-summary.md` with detailed research findings
- Length: Validation report 2-3 pages, supporting documents as needed
</output_format>

## Success Criteria

<success_criteria>

### Phase 1: User Hypothesis Definition

- User segment hypotheses extracted from current project brief
- Specific validation claims identified with measurable criteria
- User context and goals documented for validation activities
- Primary and secondary segments prioritized for validation focus

### Phase 2: Jobs-to-be-Done Framework Application

- Job statements created for each user segment using proper format
- Job dimensions identified (functional, emotional, social) with evidence
- User journey mapped with touchpoints, pain points, and opportunities
- Opportunity scores calculated using framework methodology

### Phase 3: Evidence Collection

- User interviews completed with minimum 8-10 per primary segment
- User survey deployed with statistically significant sample size
- Analytics analysis conducted validating behavioral assumptions
- Market research integrated for segment sizing and validation

### Phase 4: Jobs-to-be-Done Analysis

- Importance and satisfaction scored for all job statements
- Opportunity scores calculated and prioritized by potential impact
- Segment quality assessed using specificity, measurability, accessibility, actionability
- Evidence quality evaluated with confidence levels assigned

### Phase 5: Results and Recommendations

- Validation report completed using standardized template
- Project brief updated with validated user segment information
- Assumption tracker updated with user validation results
- Recommendations provided for user-focused solution development
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Validation Preparation

- [ ] **Current user analysis**: Review existing user segment definitions and evidence
- [ ] **Research planning**: Design interview guide, survey, and analytics analysis plan
- [ ] **Participant recruitment**: Identify and recruit representative users for research
- [ ] **Tools preparation**: Set up interview recording, survey platform, analytics access

### User Research Execution

- [ ] **Interview conduct**: Complete structured interviews with diverse user representatives
- [ ] **Survey deployment**: Launch survey and monitor response rates and quality
- [ ] **Analytics analysis**: Extract relevant user behavior data and patterns
- [ ] **Competitive research**: Analyze competitor user bases and positioning

### Jobs-to-be-Done Analysis

- [ ] **Job statement validation**: Confirm job statements with user feedback
- [ ] **Importance/satisfaction scoring**: Collect quantified ratings from users
- [ ] **Opportunity calculation**: Apply framework methodology to identify high-opportunity jobs
- [ ] **User journey validation**: Confirm journey accuracy through user input

### Results Synthesis and Documentation

- [ ] **Evidence synthesis**: Combine findings from multiple research methods
- [ ] **Validation scoring**: Apply segment quality and evidence quality assessments
- [ ] **Report creation**: Complete validation report with findings and recommendations
- [ ] **Stakeholder communication**: Present results to product team and stakeholders

</execution_checklist>

## Content Guidelines

### User Research Quality Standards

- **Representative Sampling**: Ensure research participants represent target segment diversity
- **Unbiased Research**: Use open-ended questions and avoid leading participants toward expected answers
- **Quantified Insights**: Collect measurable data wherever possible (time, frequency, cost, satisfaction)
- **Multiple Method Validation**: Validate key findings through multiple research methods
- **Evidence Documentation**: Properly document and cite all research sources and methodologies

### Jobs-to-be-Done Best Practices

- **Outcome Focus**: Focus on outcomes users want to achieve, not features they want
- **Context Importance**: Document when, where, and why users have specific jobs
- **Emotional Dimensions**: Don't ignore emotional and social aspects of user jobs
- **Job Evolution**: Understand how user jobs change over time or in different contexts
- **Solution Agnostic**: Define jobs independent of current solution approaches

### User Segment Definition Standards

- **Actionable Segments**: Define segments that enable different marketing or product approaches
- **Measurable Characteristics**: Include quantifiable attributes for segment tracking
- **Accessible Segments**: Ensure segments can be reached through available marketing channels
- **Stable Definitions**: Create segment definitions that remain valid over time
- **Evidence-Based**: Ground all segment characteristics in validated research findings

## Framework Integration Notes

- **SDLC Integration**: User validation results inform Epic prioritization and Story creation
- **Business Framework Usage**: Leverages Jobs-to-be-Done methodology for systematic user understanding
- **Evidence Standards**: Maintains quantified, multi-source validation approach
- **Quality Assurance**: Built-in scoring and validation ensures reliable user insights
- **Professional Output**: Structured documentation supports product and marketing decisions
