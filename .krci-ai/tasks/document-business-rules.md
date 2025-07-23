# Task: Document Business Rules

## Description

Systematically document business rules and constraints that govern system behavior, supporting PRD requirements and Epic implementation with clear business logic specifications. This task translates business policies, regulatory requirements, and operational constraints into structured rules that guide system design and development decisions, ensuring compliance and consistency in rule application across all Epic features.

## Prerequisites

- [ ] **PRD foundation**: PRD exists at `/docs/prd/prd.md` with basic business and system requirements
- [ ] **Business process understanding**: Current workflows and decision points identified
- [ ] **Stakeholder access**: Subject matter experts and decision makers available for rule validation
- [ ] **Regulatory context**: Compliance requirements and organizational policies understood

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/analysis-methodologies.md
- ./.krci-ai/templates/business-rules.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for business rule documentation dependencies and workflow
2. **Apply analysis methodologies**: Use approaches from [analysis-methodologies.md](./.krci-ai/data/analysis-methodologies.md)
3. **Format output**: Use [business-rules.md](./.krci-ai/templates/business-rules.md) template for structured documentation
4. **Link to PRD**: Ensure business rules support and clarify PRD requirements (BR/NFR)

## Output Format

**Business Rules Documentation** - Create comprehensive rules repository:

- [ ] **Primary documentation**: `/docs/business-rules.md` with structured rule catalog following [business-rules.md](./.krci-ai/templates/business-rules.md) template
- [ ] **Rules traceability**: Clear mapping from business policies to system rules to PRD requirements
- [ ] **Epic implementation guidance**: Rules structured to support Epic feature development
- [ ] **Governance framework**: Rule management and approval processes documented

## Success Criteria

- [ ] **Rules catalog completed** - All business rules documented with standard structure and clear logic
- [ ] **PRD alignment established** - Business rules support and clarify PRD BR/NFR requirements
- [ ] **Epic enablement provided** - Rules structured to guide Epic feature implementation
- [ ] **Governance framework defined** - Rule ownership, approval, and change management processes established
- [ ] **Compliance validated** - All regulatory and policy requirements addressed in rule documentation
- [ ] **Stakeholder approval obtained** - Business rules reviewed and approved by subject matter experts

## Execution Checklist

### Rule Discovery Phase

- [ ] **Decision point analysis**: Review business processes to identify all decision points and rule applications
- [ ] **Policy documentation review**: Examine existing organizational policies, procedures, and regulatory requirements
- [ ] **Stakeholder interviews**: Conduct sessions with subject matter experts to extract business logic using [business-rules.md](./.krci-ai/templates/business-rules.md) format
- [ ] **System constraint identification**: Analyze current system logic and algorithmic rules

### Rule Documentation Phase

- [ ] **Rule categorization**: Organize rules by business domain, rule type (constraints, derivations, action enablers), and complexity
- [ ] **Structured documentation**: Document each rule using standard template with conditions, actions, exceptions, and business rationale
- [ ] **Logic validation**: Test rule logic with business scenarios and edge cases
- [ ] **Business justification**: Document why each rule exists and its business value

### PRD Integration Phase

- [ ] **Requirements mapping**: Link business rules to specific PRD BR/NFR requirements for traceability
- [ ] **Epic guidance creation**: Structure rules to provide clear implementation guidance for Epic features
- [ ] **Compliance verification**: Ensure all regulatory and policy requirements are addressed
- [ ] **Conflict resolution**: Identify and resolve conflicting or contradictory rules

### Governance and Validation Phase

- [ ] **Rule ownership assignment**: Identify business owners and stewards for each rule or rule category
- [ ] **Approval workflows**: Establish processes for rule validation and change management
- [ ] **Stakeholder validation**: Review rules with business stakeholders and obtain formal approval
- [ ] **Implementation readiness**: Ensure rules provide sufficient guidance for Epic and Story development

## Content Guidelines

### 📋 **Business Rule Structure:**

#### **Rule Documentation Format:**

```
**Rule ID**: BR-[NUMBER] (e.g., BR-001)
**Rule Name**: [Descriptive business rule name]
**Rule Type**: [Constraint/Derivation/Action Enabler]
**Business Domain**: [Functional area this rule applies to]

**Rule Statement**: [Clear, unambiguous rule logic]
**Business Rationale**: [Why this rule exists and its business value]
**Conditions**: [Specific circumstances when rule applies]
**Actions**: [What happens when conditions are met]
**Exceptions**: [Circumstances that override the rule]
```

#### **PRD Integration Points:**

- **BR Requirements**: How business rules support specific Business Requirements (BR1, BR2...)
- **NFR Requirements**: How rules address Non-Functional Requirements (NFR1, NFR2...)
- **Epic Features**: How rules guide Epic implementation and acceptance criteria

### ✅ **Quality Standards:**

- **Business Validated**: All rules reviewed and approved by subject matter experts
- **PRD Aligned**: Rules clearly support and clarify PRD requirements
- **Epic Enabled**: Rules provide clear implementation guidance for Epic features
- **Logically Consistent**: No conflicting or contradictory rules exist
- **Compliance Focused**: All regulatory and policy requirements addressed
- **Governable**: Clear ownership and change management processes established

### ❌ **Common Pitfalls to Avoid:**

- Documenting solutions instead of business rules
- Creating overly complex or ambiguous rule statements
- Missing business context and rationale for rules
- Inadequate stakeholder validation and approval
- Poor rule organization and accessibility
- Conflicting or contradictory rules without resolution

### 🎯 **Epic Implementation Enablement:**

This business rules documentation should enable immediate Epic development by providing:

- **Clear constraints** that become Epic acceptance criteria and system requirements
- **Business logic** that guides Epic feature behavior and decision flows
- **Compliance requirements** that inform Epic implementation and validation approaches
- **Rule traceability** that connects Epic features back to business policies and regulatory requirements
