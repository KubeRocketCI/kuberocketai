---
dependencies:
  templates:
    - story.md
---

# Task: Review Story (Product Owner)

## Description

Review and validate user story from Product Owner perspective to ensure business value clarity, acceptance criteria completeness, and epic alignment. Focus on user value, story format correctness, and implementation readiness from business requirements standpoint.

## Instructions

<instructions>
Confirm the target story file exists in `/docs/stories/` requiring product owner review, parent Epic's business goals and user value are understood, PRD requirements and user personas are familiar, and you have Product Owner approval rights for story advancement. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Validate business requirements to ensure story delivers clear user value aligned with Epic goals. Check story format to verify proper "As a/I want/so that" structure with clear user benefit. Review acceptance criteria to validate they are testable, specific, and business-focused, allowing Verification Method (manual/semi-automated) with required Evidence when commands aren't feasible. Confirm epic alignment to ensure story supports parent Epic objectives and user outcomes. Assess user value to validate story provides measurable business value to target users.
</instructions>

## Output Format

- Location: Update existing story file with business validation
- Template: Maintain [story.md](./.krci-ai/templates/story.md) structure (8 sections only)
- Content Placement: Business context in Description section, approval in Implementation Results
- Business Validation: Document user value confirmation and Epic alignment verification
- Verification: Story passes PO review with documented business approval

## Success Criteria

<success_criteria>
- Story format correct: "As a [user], I want [goal], so that [value]" properly structured
- Business value clear: User benefit and business rationale obvious and measurable
- Acceptance criteria business-focused: Criteria validate user value delivery
- Epic alignment confirmed: Story supports parent Epic goals and user outcomes
- User persona validation: Target user aligns with Epic user definitions
- PO approval documented: Business validation and approval recorded
</success_criteria>

## Execution Checklist

### Business Requirements Validation

<business_validation>
- User value assessment: Confirm story delivers clear, measurable user benefit
- Business justification: Validate business need and priority for this functionality
- User persona alignment: Verify target user matches Epic persona definitions
- Value proposition clarity: Ensure "so that" clause provides clear business benefit
</business_validation>

### Story Format Review

<story_format_review>
- Structure validation: Verify "As a [user], I want [goal], so that [value]" format
- User specification: Confirm specific user role/persona rather than generic "user"
- Goal clarity: Validate goal is specific, actionable, and user-focused
- Value articulation: Ensure business value and user benefit are explicit
</story_format_review>

### Acceptance Criteria Business Validation

<acceptance_criteria_validation>
- Business testability: Confirm criteria can be validated by business stakeholders
- User value measurement: Verify criteria measure actual user benefit delivery
- Completeness assessment: Ensure criteria cover all business validation requirements
- Success metrics alignment: Confirm criteria support Epic success measurements
</acceptance_criteria_validation>

### Epic Alignment Verification

<epic_alignment_verification>
- Goal consistency: Story goal supports parent Epic objectives
- User alignment: Story user matches Epic target user definitions
- Scope compliance: Story scope fits within Epic boundaries
- Priority validation: Story priority aligns with Epic and business priorities
</epic_alignment_verification>

### Implementation Readiness (Business Perspective)

<implementation_readiness>
- Requirements completeness: All business requirements clearly specified
- User acceptance preparation: Story ready for user acceptance validation
- Business validation plan: Clear approach for validating business value delivery
- Stakeholder communication: Key stakeholders informed and aligned
</implementation_readiness>

## Content Guidelines

### Business Validation Principles for LLM Self-Evaluation

- User-Centered Focus: Every validation centers on user value and business benefit
- Clear Value Articulation: Business value and user benefit must be explicit and measurable
- Epic Consistency: All story elements must align with parent Epic goals and users
- Business Testability: Acceptance criteria must be validatable by business stakeholders

### LLM Error Prevention Checklist

- Avoid: Technical implementation details outside PO scope and authority
- Avoid: Validation that requires technical expertise beyond business requirements
- Avoid: Accepting vague business value statements without specific user benefit
- Reference: Focus on business requirements alignment with [story.md](./.krci-ai/templates/story.md) template
