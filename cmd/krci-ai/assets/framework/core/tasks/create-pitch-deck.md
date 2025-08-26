# Task: Create Pitch Deck

## Description

Create a concise, high-impact pitch deck (3-5 slides) that transforms product features into an emotional story driving immediate action. Focus on maximum "WOW factor" through powerful visuals, compelling problem-solution narrative, and professional design that makes audiences stop and say "I need this product" within minutes, not hours.

## Prerequisites

**Core Requirements (PRD-First + Interactive):**

- [ ] Completed PRD at `/docs/prd/prd.md` containing product information (primary source)
- [ ] Product demo or MVP available for showcasing
- [ ] User available to answer interactive questions for gaps not covered in PRD:
  - Target audience specifics
  - Desired presentation tone
  - Competitive context
  - Presentation objectives

**Optional (For Advanced Marketing Scenarios):**

- [ ] Marketing Brief at `/docs/marketing/marketing-brief.md` (for sophisticated positioning)
- [ ] Competitive analysis and market research (for differentiation focus)
- [ ] Visual assets and brand guidelines (for brand-consistent presentations)

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/pitch-deck-template.md
- /docs/prd/prd.md (primary source for product information)

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing. If PRD is incomplete or missing sections, proceed with available information and gather missing details through user interaction.

## Instructions

**EXECUTION SEQUENCE (Follow in Order):**

1. **Load Context**: Read all reference assets to load framework knowledge and templates
2. **Gather Information**: Collect all prerequisites and analyze PRD/MVP content
3. **Select Framework**: Choose optimal presentation framework based on audience and content strategy
4. **Structure Content**: Apply chosen framework to organize pitch deck slides
5. **Create Output**: Generate final pitch deck using template structure with visual design guidance

**Framework Application Requirements:**

- **Mandatory**: Use presentation frameworks from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
- **Available Frameworks**: Pain-Gains-Reveals, PAS, BAB, SCRAP (all adapt to 3-5 slides)
- **Template Structure**: Use [pitch-deck-template.md](./.krci-ai/templates/pitch-deck-template.md) for formatting

## Output Format

- **Location**: `/docs/marketing/pitch-deck.md` (EXACT path and filename)
- **Length**: 3-5 slides maximum for focused, high-impact presentation
- **Format**: Structured markdown with concise slide descriptions optimized for immediate visual implementation
- **Downstream Enable**: Ready for PowerPoint/presentation tool conversion and executive-level presentations

## Success Criteria

- [ ] File saved to `/docs/marketing/pitch-deck-{number}.md`
- [ ] Slide count is 3-5 slides maximum for optimal attention span
- [ ] Opening hook creates immediate engagement and curiosity within first 30 seconds
- [ ] Problem statement resonates emotionally with target audience pain points
- [ ] Solution reveal has clear "AHA moment" that demonstrates product value
- [ ] Visual design descriptions enable compelling slide creation with modern aesthetics
- [ ] Call to action is specific and drives desired next steps

## Execution Checklist

### Step 1: Load Context

- [ ] Read business-frameworks.md: Load Pain-Gains-Reveals, PAS, BAB, SCRAP framework knowledge
- [ ] Read pitch-deck-template.md: Load template structure and formatting requirements
- [ ] Read sdlc-framework.md: Understand artifact dependencies and workflow context

### Step 2: Gather Information (PRD-First + Interactive)

- [ ] Extract PRD information: Follow template "PRD Integration Guide" for specific extraction instructions
- [ ] Review MVP: Identify key capabilities for demonstration
- [ ] Ask user for gaps: Follow template "PRD Integration Guide" for interactive questions to fill missing information

**Interactive Question Examples:**

- "Who is your specific presentation audience? (e.g., 'Tech executives at SaaS companies' or 'Series A investors')"
- "What tone should this presentation have? (professional, innovative, empathetic, or other)"
- "What are your main competitors in this space, and how do you differentiate?"
- "What's your primary objective for this presentation? (funding, partnerships, sales, etc.)"

### Step 3: Select Framework

- [ ] Choose framework: Follow template "Framework Adaptation Guide" to select optimal framework based on audience and content strategy. **Default: Start with Pain-Gains-Reveals framework unless user specifies preference for PAS, BAB, or SCRAP**
- [ ] Determine slide count: Use template framework structure to adapt to 3-5 slides based on content depth

### Step 4: Structure Content

- [ ] Apply framework: Structure slides according to chosen framework methodology
- [ ] Integrate persuasion: Add Social Proof, Authority, Scarcity elements
- [ ] Add emotional connection: Use Customer Emotion Guidelines for human feelings

**Content Integration Guide:** Start by mapping your framework structure to emotional touchpoints - for Pain-Gains-Reveals, weave frustration/hope emotions into Pain slides, excitement/relief into Gains, and confidence/authority into Reveals. Layer persuasion psychology throughout: use Social Proof in problem validation, Authority in solution presentation, and Scarcity in competitive differentiation. Apply STAR method (Situation-Task-Action-Result) for any proof points, ensuring each slide contributes to the emotional journey from problem awareness to solution commitment. **Follow all template guidance sections for implementation details.**

### Step 5: Create Output

- [ ] Populate template variables: Follow template "Variable Source Guide" to map gathered information to template variables
- [ ] Generate pitch deck: Use template structure with chosen framework content following "Framework Adaptation Guide"
- [ ] Add visual descriptions: Follow template visual design format examples for each slide
- [ ] Save file: Place at exact location `/docs/marketing/pitch-deck-{number}.md`

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Framework Adherence**: MUST follow chosen framework structure exactly (Pain-Gains-Reveals/PAS/BAB/SCRAP)
- **Persuasion Integration**: Apply psychology principles (Social Proof, Authority, Scarcity) from business-frameworks.md
- **Emotional Impact**: Every slide should contribute to emotional journey from problem to solution
- **Visual Storytelling**: Use compelling visuals and design principles to enhance narrative impact
- **Evidence Structure**: Use STAR method for all proof points and testimonials

### Framework Application Reference

**Framework Selection**: Use template "Framework Adaptation Guide" for detailed structure and selection criteria
**Customer Emotion Guidelines**: Follow template guidance for human emotions and quantifiable impact integration

### LLM Error Prevention Checklist

- **Avoid**: Feature lists without clear customer benefit and emotional connection
- **Avoid**: Generic template slides that lack specific visual design guidance
- **Avoid**: Overwhelming information density that dilutes core message impact
- **Avoid**: Mixing frameworks - stick to one chosen framework throughout
- **Reference**: Use [pitch-deck-template.md](./.krci-ai/templates/pitch-deck-template.md) for proven slide structure

### WOW Factor Design Principles for 3-5 Slides

- **Slide 1 - Hook**: Open with surprising statistic, compelling question, or bold problem statement that audience immediately recognizes
- **Slide 2-3 - Solution**: Present clear before/after transformation with visual impact and concrete value demonstration
- **Slide 4-5 - Action**: Combine proof (traction/testimonial) with specific call to action that audience can take immediately
- **Visual Impact**: Each slide must standalone - no gradual build-up, immediate clarity and impact
- **Message Density**: Pack maximum emotional and logical impact into minimal content - every word counts

### Framework Validation Checklist

Before finalizing the pitch deck, verify strategic framework application:

- [ ] Framework selection rationale: Chosen framework aligns with audience type and content strategy
- [ ] Framework consistency: All slides follow chosen framework structure and logical flow
- [ ] Slide count optimization: Framework adapted appropriately for chosen slide count (3-5)
- [ ] Psychology principles: At least 2 persuasion principles integrated (Social Proof, Authority, Scarcity)
- [ ] STAR method: All proof points follow Situation-Task-Action-Result structure
- [ ] Narrative coherence: Content tells compelling story according to framework methodology
- [ ] Audience alignment: Framework choice serves target audience needs and decision criteria

### SDLC Integration Context

This Pitch Deck transforms PRD requirements and MVP demonstrations into compelling presentations for sales enablement, investor relations, stakeholder alignment, and customer acquisition. Uses proven presentation frameworks from business-frameworks.md to structure problem-solution narratives directly from product requirements, with optional marketing brief enhancement for advanced positioning scenarios.
