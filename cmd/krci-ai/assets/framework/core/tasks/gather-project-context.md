---
dependencies:
  data:
    - validation-frameworks.md
    - business-frameworks.md
    - krci-ai/core-sdlc-framework.md
  templates:
    - context-gathering-guide-template.md
---

# Task: Gather Project Context

## Description

Systematically collect and structure project context using business frameworks and established methodologies. This task provides multiple input methods to gather comprehensive information needed for evidence-based project brief creation, ensuring all assumptions are tracked and validated with appropriate evidence.

This task leverages [validation frameworks](./.krci-ai/data/validation-frameworks.md) and the [context gathering guide template](./.krci-ai/templates/context-gathering-guide-template.md) for systematic context collection.

## Prerequisites

<prerequisites>
- Business opportunity or problem area identified
- Stakeholder access available for interviews/input
- Initial project scope or direction defined
- Access to existing research, documentation, or data sources
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/validation-frameworks.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/data/krci-ai/core-sdlc-framework.md
- ./.krci-ai/templates/context-gathering-guide-template.md

CRITICAL: Load all dependencies by reading their complete content before task execution. HALT if any missing.

## Instructions

<instructions>
Systematically collect and structure project context using business frameworks and established methodologies. This provides comprehensive information needed for evidence-based project brief creation.

### Phase 1: Input Method Selection

Choose appropriate input collection methods based on available resources and project context:

#### Method 1: Document Analysis

**When to use**: Existing research, competitive analysis, user studies, or PRDs available
**Process**:
1. Inventory all available documents and research
2. Categorize documents by type (user research, market data, technical specs, competitive analysis)
3. Extract key insights using structured analysis framework
4. Identify evidence gaps and assumption areas

**Document Types to Analyze**:
- Market research reports
- Competitive analysis
- User research and interviews
- Existing PRDs or product specs
- Technical feasibility studies
- Stakeholder feedback
- Customer support data
- Analytics and usage data

#### Method 2: Stakeholder Interview Process

**When to use**: Key stakeholders available for structured input gathering
**Process**:
1. Use BABOK elicitation techniques for structured interviews
2. Apply Design Thinking empathy mapping methods
3. Document interviews using standardized templates
4. Validate findings across multiple stakeholders

**Stakeholder Categories**:
- Business stakeholders (product owners, executives)
- User representatives (customer success, support)
- Technical stakeholders (architects, lead developers)
- Market experts (sales, marketing, customer success)

#### Method 3: Assumption Inventory Creation

**When to use**: Limited existing research, need to identify knowledge gaps
**Process**:
1. Brainstorm all assumptions about problem, users, solution, market
2. Categorize assumptions by type and confidence level
3. Prioritize assumptions by impact and uncertainty
4. Create validation plan for high-priority assumptions

**Assumption Categories**:
- Problem assumptions (what, who, why, when)
- User assumptions (segments, needs, behaviors, preferences)
- Solution assumptions (feasibility, desirability, viability)
- Market assumptions (size, competition, trends)
- Business assumptions (model, metrics, resources)

#### Method 4: Evidence Library Creation

**When to use**: Quantified data available, need systematic evidence collection
**Process**:
1. Identify available quantified data sources
2. Structure data using evidence classification framework
3. Assess data quality and confidence levels
4. Identify additional data needs

**Evidence Types**:
- Usage analytics and user behavior data
- Support ticket analysis and customer feedback
- Market sizing and industry benchmark data
- Technical performance and feasibility data
- Financial data (costs, revenue, pricing)

### Phase 2: Business Framework Application

Apply relevant [business frameworks](./.krci-ai/data/business-frameworks.md) to structure collected information:

#### Problem Context Analysis

**Frameworks to Apply**:
- **Six Sigma 5 Whys**: For root cause analysis of identified problems
- **Impact-Frequency Matrix**: For problem prioritization and urgency assessment
- **SIPOC Analysis**: For process understanding and stakeholder mapping

**Output**: Structured problem context with root causes, stakeholders, and priority assessment

#### User Context Analysis

**Frameworks to Apply**:
- **Jobs-to-be-Done**: For user motivation and outcome identification
- **Design Thinking Empathy Maps**: For user perspective and experience mapping
- **User Journey Mapping**: For touchpoint and pain point identification

**Output**: Comprehensive user context with jobs, pains, gains, and journey insights

#### Market Context Analysis

**Frameworks to Apply**:
- **Business Model Canvas**: For value proposition and business model analysis
- **Porter's Five Forces**: For competitive environment assessment
- **TAM/SAM/SOM Analysis**: For market sizing and opportunity quantification

**Output**: Market context with competitive landscape, opportunity size, and positioning

#### Business Context Analysis

**Frameworks to Apply**:
- **SWOT Analysis**: For internal capabilities and external factors
- **OKR Alignment**: For strategic alignment and goal hierarchy
- **Resource Assessment**: For capability and constraint analysis

**Output**: Business context with strategic alignment, resources, and constraints

### Phase 3: Context Synthesis and Validation

Synthesize collected information into structured project context:

#### Context Synthesis Process

1. **Integrate Multi-Source Findings**: Combine insights from all input methods
2. **Identify Patterns and Themes**: Look for consistent insights across sources
3. **Highlight Contradictions**: Note conflicting information for further validation
4. **Assess Evidence Quality**: Rate confidence levels for each insight

#### Validation Checkpoints

1. **Stakeholder Validation**: Review synthesized context with key stakeholders
2. **Evidence Validation**: Verify quantified claims with additional sources
3. **Assumption Validation**: Test high-impact assumptions with appropriate methods
4. **Internal Consistency**: Ensure context elements align logically

### Phase 4: Context Documentation

Document project context using structured templates:

#### Context Summary Structure

```markdown
# Project Context Summary: {{project_name}}

## Context Collection Methodology
- Input methods used: {{methods_list}}
- Business frameworks applied: {{frameworks_list}}
- Evidence sources: {{sources_list}}
- Confidence assessment: {{overall_confidence_level}}

## Problem Context
- Root problem identification using {{validation_method}}
- Problem scope and boundaries
- Stakeholder impact assessment
- Priority and urgency analysis

## User Context
- Target user segments with validation evidence
- User jobs, pains, and gains analysis
- User journey and touchpoint mapping
- User needs prioritization

## Market Context
- Market opportunity sizing and validation
- Competitive landscape analysis
- Value proposition positioning
- Market trend and timing analysis

## Business Context
- Strategic alignment and goal hierarchy
- Resource and capability assessment
- Constraint and limitation identification
- Success criteria and measurement approach

## Evidence Library
- Primary evidence sources and confidence levels
- Secondary research and benchmark data
- Quantified metrics and baseline measurements
- Evidence gaps and validation needs

## Assumption Inventory
- High-impact assumptions requiring validation
- Medium-impact assumptions for monitoring
- Low-impact assumptions acceptable as-is
- Assumption validation plan and timeline
```

</instructions>

## Output Format

<output_format>

- Location: `/docs/prd/project-context.md` (EXACT path and filename)
- Supporting Files:
  - `/docs/prd/evidence-library.md`
  - `/docs/prd/assumption-inventory.md`
- Length: 3-4 pages maximum for executive consumption
- Downstream Enablement: Enables enhanced project brief creation
</output_format>

## Success Criteria

<success_criteria>

### Phase 1: Input Collection

- Input methods selected based on available resources and project needs
- Document analysis completed with structured insight extraction
- Stakeholder interviews conducted using business framework methods
- Assumption inventory created with prioritization and confidence levels

### Phase 2: Framework Application

- Problem context analyzed using appropriate business frameworks
- User context mapped with jobs, pains, gains, and journey insights
- Market context assessed with competitive and opportunity analysis
- Business context documented with strategic alignment and constraints

### Phase 3: Context Validation

- Multi-source integration completed with pattern identification
- Stakeholder validation conducted with key project stakeholders
- Evidence quality assessed with confidence levels assigned
- Assumption validation plan created for high-impact assumptions

### Phase 4: Documentation

- Context summary created using structured template
- Evidence library documented with sources and confidence levels
- Assumption inventory finalized with validation priorities
- Files saved to exact locations `/docs/prd/project-context.md`
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Context Gathering

- [ ] **Stakeholder identification**: Map all relevant stakeholders and their roles
- [ ] **Resource assessment**: Identify available documents, data, and research
- [ ] **Timeline planning**: Set realistic timeline for context gathering activities
- [ ] **Framework selection**: Choose appropriate business frameworks for analysis

### Context Collection Phase

- [ ] **Document inventory**: Catalog all available research and documentation
- [ ] **Stakeholder interviews**: Conduct structured interviews using BABOK methods
- [ ] **Data collection**: Gather quantified evidence from available sources
- [ ] **Assumption brainstorming**: Identify assumptions across all project dimensions

### Analysis Phase

- [ ] **Framework application**: Apply selected business frameworks to collected data
- [ ] **Pattern recognition**: Identify themes and insights across sources
- [ ] **Gap identification**: Identify evidence gaps and validation needs
- [ ] **Confidence assessment**: Rate evidence quality and assumption confidence

### Validation Phase

- [ ] **Stakeholder review**: Validate synthesized context with key stakeholders
- [ ] **Evidence verification**: Cross-check quantified claims with additional sources
- [ ] **Assumption testing**: Plan validation for high-impact assumptions
- [ ] **Internal consistency**: Ensure logical alignment across context elements

### Documentation Phase

- [ ] **Context summary**: Create comprehensive context documentation
- [ ] **Evidence library**: Document all sources with confidence assessments
- [ ] **Assumption inventory**: List assumptions with validation priorities
- [ ] **Quality review**: Verify completeness and accuracy of documentation
</execution_checklist>

## Content Guidelines

### Quality Principles for Context Gathering

- **Evidence-Based**: Prioritize quantified evidence over opinions or assumptions
- **Multi-Source Validation**: Validate key insights across multiple sources
- **Assumption Transparency**: Clearly identify and track all assumptions
- **Confidence Assessment**: Rate evidence quality and assumption confidence levels
- **Business Framework Grounding**: Use established methodologies for analysis

### Context Collection Best Practices

- **Structure Over Volume**: Focus on structured insights rather than comprehensive data
- **Validation Over Confirmation**: Seek to test assumptions rather than confirm biases
- **Quantification Over Description**: Prioritize measurable insights over qualitative observations
- **Stakeholder Diversity**: Include diverse perspectives to avoid single viewpoint bias
- **Evidence Traceability**: Maintain clear links from insights back to source evidence

### Documentation Standards

- **Executive Summary Focus**: Keep context summary concise and decision-focused
- **Evidence Transparency**: Clearly document sources and confidence levels
- **Assumption Accountability**: Make assumptions explicit and track validation status
- **Action Orientation**: Structure context to enable project brief creation
- **Professional Format**: Use consistent formatting and structure for stakeholder consumption

## Integration with Project Brief Creation

### Context-to-Brief Mapping

**Project Context → Project Brief Sections**:
- **Problem Context** → Problem Statement and Opportunity sections
- **User Context** → Target Users section with validated segments
- **Market Context** → Opportunity section with market sizing and positioning
- **Business Context** → Success Metrics, Constraints, and Risk sections
- **Evidence Library** → Supporting evidence for all brief assertions
- **Assumption Inventory** → Transparent assumption tracking throughout brief

### Enhanced Brief Quality Through Context

- **Evidence-Based Assertions**: All brief statements backed by context evidence
- **Validated Assumptions**: Brief assumptions tested through context gathering
- **Stakeholder Alignment**: Brief content pre-validated with stakeholders during context phase
- **Risk Mitigation**: Context gaps identified and addressed before brief creation
- **Strategic Alignment**: Brief grounded in business context and strategic goals

## Framework Integration Notes

- **SDLC Integration**: Project context serves as foundation for all downstream [SDLC framework](./.krci-ai/data/krci-ai/core-sdlc-framework.md) artifacts
- **Business Framework Usage**: Leverages established methodologies for professional analysis
- **Evidence Standards**: Maintains KubeRocketAI focus on quantified, evidence-based approach
- **Quality Assurance**: Built-in validation checkpoints ensure context accuracy and completeness
- **Professional Output**: Structured documentation suitable for executive and stakeholder consumption
