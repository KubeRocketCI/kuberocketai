# Task: Review Story (Product Owner)

## Description

Review and validate user story from Product Owner perspective to ensure business value clarity, acceptance criteria completeness, and epic alignment. Focus on user value, story format correctness, and implementation readiness from business requirements standpoint.

## Prerequisites

- [ ] **Story exists**: Target story file exists in `/docs/stories/` requiring product owner review
- [ ] **Epic context**: Understanding of parent Epic's business goals and user value
- [ ] **Product requirements**: Familiarity with PRD requirements and user personas
- [ ] **Business validation authority**: Product Owner approval rights for story advancement

### Reference Assets

Dependencies:

- ./.krci-ai/templates/story.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Validate business requirements**: Ensure story delivers clear user value aligned with Epic goals
2. **Check story format**: Verify proper "As a/I want/so that" structure with clear user benefit
3. **Review acceptance criteria**: Validate criteria are testable, specific, and business-focused; allow Verification Method (manual/semi-automated) with required Evidence when commands aren't feasible
4. **Confirm epic alignment**: Ensure story supports parent Epic objectives and user outcomes
5. **Assess user value**: Validate story provides measurable business value to target users

## Output Format

- **Location**: Update existing story file with business validation
- **Template**: Maintain [story.md](./.krci-ai/templates/story.md) structure (8 sections only)
- **Content Placement**: Business context in Description section, approval in Implementation Results
- **Business Validation**: Document user value confirmation and Epic alignment verification
- **Verification**: Story passes PO review with documented business approval

## Success Criteria

- [ ] **Story format correct**: "As a [user], I want [goal], so that [value]" properly structured
- [ ] **Business value clear**: User benefit and business rationale obvious and measurable
- [ ] **Acceptance criteria business-focused**: Criteria validate user value delivery
- [ ] **Epic alignment confirmed**: Story supports parent Epic goals and user outcomes
- [ ] **User persona validation**: Target user aligns with Epic user definitions
- [ ] **PO approval documented**: Business validation and approval recorded

## Execution Checklist

### Business Requirements Validation

- [ ] **User value assessment**: Confirm story delivers clear, measurable user benefit
- [ ] **Business justification**: Validate business need and priority for this functionality
- [ ] **User persona alignment**: Verify target user matches Epic persona definitions
- [ ] **Value proposition clarity**: Ensure "so that" clause provides clear business benefit

### Story Format Review

- [ ] **Structure validation**: Verify "As a [user], I want [goal], so that [value]" format
- [ ] **User specification**: Confirm specific user role/persona rather than generic "user"
- [ ] **Goal clarity**: Validate goal is specific, actionable, and user-focused
- [ ] **Value articulation**: Ensure business value and user benefit are explicit

### Acceptance Criteria Business Validation

- [ ] **Business testability**: Confirm criteria can be validated by business stakeholders
- [ ] **User value measurement**: Verify criteria measure actual user benefit delivery
- [ ] **Completeness assessment**: Ensure criteria cover all business validation requirements
- [ ] **Success metrics alignment**: Confirm criteria support Epic success measurements

### Epic Alignment Verification

- [ ] **Goal consistency**: Story goal supports parent Epic objectives
- [ ] **User alignment**: Story user matches Epic target user definitions
- [ ] **Scope compliance**: Story scope fits within Epic boundaries
- [ ] **Priority validation**: Story priority aligns with Epic and business priorities

### Implementation Readiness (Business Perspective)

- [ ] **Requirements completeness**: All business requirements clearly specified
- [ ] **User acceptance preparation**: Story ready for user acceptance validation
- [ ] **Business validation plan**: Clear approach for validating business value delivery
- [ ] **Stakeholder communication**: Key stakeholders informed and aligned

## Content Guidelines

### Business Validation Principles for LLM Self-Evaluation

- **User-Centered Focus**: Every validation centers on user value and business benefit
- **Clear Value Articulation**: Business value and user benefit must be explicit and measurable
- **Epic Consistency**: All story elements must align with parent Epic goals and users
- **Business Testability**: Acceptance criteria must be validatable by business stakeholders

### LLM Error Prevention Checklist

- **Avoid**: Technical implementation details outside PO scope and authority
- **Avoid**: Validation that requires technical expertise beyond business requirements
- **Avoid**: Accepting vague business value statements without specific user benefit
- **Reference**: Focus on business requirements alignment with [story.md](./.krci-ai/templates/story.md) template
