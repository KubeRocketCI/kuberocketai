# 99. Appendix: Solution Architecture Document (SAD) Template

This appendix provides the comprehensive Architecture Document template for the SDLC framework. This template synthesizes industry best practices from enterprise architecture standards while providing specific SDLC integration for Epic/Story development.

---

## Recommended File Structure for Large SAD Documents

For complex architecture documents that may grow beyond 1,000 lines, the complete template below can be organized into the following granular file structure while preserving all content, scope, and visual diagrams:

```bash
docs/architecture/
├── 01-executive-summary.md           # Business overview & strategic objectives
├── 02-introduction.md                # Scope, stakeholders, PRD requirements mapping
├── 03-context.md                     # Technology strategy, business & data architecture
├── 04-requirements.md                # Business goals, functional & non-functional requirements
├── 05-baseline-architecture.md      # Current state conceptual, logical & deployment views
├── 06-target-architecture.md        # Future state architecture with C4 diagrams
├── 07-transition-migration.md       # Implementation roadmap & Epic/Story guidance
├── 08-architectural-decisions.md    # ADR format decision log with alternatives
├── 09-cross-cutting-concerns.md     # Security, scalability, observability, fault tolerance
├── 10-quality-assurance.md          # Testing strategy, automation & quality metrics
├── 11-appendices.md                 # Glossary, diagrams index & reference materials
└── README.md                        # Navigation guide linking all sections
```

### **File Content Distribution**

| File | Original Template Sections | Key Content | Mermaid Diagrams |
|------|----------------------------|-------------|------------------|
| `01-executive-summary.md` | Section 1 | Business context, architectural approach, success metrics | Business flow diagrams |
| `02-introduction.md` | Section 2 | Definitions, scope, stakeholders, **PRD requirements mapping** | Stakeholder interaction diagrams |
| `03-context.md` | Section 3 | Technology strategy, business/data/infrastructure/application/security architecture | Context diagrams |
| `04-requirements.md` | Section 4 | Business goals, functional requirements, NFRs, constraints, assumptions | Requirements flow diagrams |
| `05-baseline-architecture.md` | Section 5 | Current state conceptual, logical, integration, physical views | Current state C4 diagrams |
| `06-target-architecture.md` | Section 6 | **Target state C4 diagrams**, quality attributes, solution strategy, risks | Target state C4 diagrams |
| `07-transition-migration.md` | Section 7 | Migration approach, roadmap, **Epic breakdown guidance** | Implementation timeline diagrams |
| `08-architectural-decisions.md` | Section 8 | ADR format decisions with context, alternatives, consequences | Decision flow diagrams |
| `09-cross-cutting-concerns.md` | Section 9 | Security, scalability, observability, fault tolerance approaches | Cross-cutting concern diagrams |
| `10-quality-assurance.md` | Section 10 | Testing strategy, automation approach, quality metrics | Testing workflow diagrams |
| `11-appendices.md` | Section 11 | Glossary, diagram index, reference materials | Diagram catalog |

### **Usage Guidelines**

- **All Projects**: Always use multiple file structure for better organization and maintainability
- **Small Projects**: Use core 5-file structure (sections 1, 2, 6, 7, 8)
- **Medium Projects**: Use 8-file structure (sections 1, 2, 3, 6, 7, 8, 9, 10)
- **Large Projects**: Use full 11-file structure above
- **Flexible Sections**: Add or remove sections based on project requirements and complexity
- **All diagrams preserved**: Every `mermaid` block stays in appropriate file
- **Cross-references maintained**: Files link to each other as needed

### **SDLC Framework Integration**

This SAD template integrates with the [SDLC Framework](./13-sdlc-framework.md) architecture documentation approach:

- **Current Project Compatibility**: This template can extend existing project architecture structure (01-introduction.md, 02-high-level-architecture.md, etc.)
- **Template Flexibility**: Use this template alongside or instead of current architecture organization based on project complexity
- **Enterprise Alignment**: Full template provides enterprise-grade architecture documentation with PRD traceability and Epic breakdown guidance
- **Migration Path**: Projects can start with current structure and selectively adopt SAD template sections as needed

**When to Use This Template**:

- **New Enterprise Projects**: Start with full SAD template structure
- **Existing Projects**: Selectively adopt sections (04-requirements.md for PRD mapping, 07-transition-migration.md for Epic guidance)
- **Complex Systems**: Use complete template for comprehensive stakeholder communication

---

## Template Overview

**Purpose**: Define comprehensive system architecture with SDLC integration, providing technical design guidance that bridges PRD requirements to Epic/Story implementation.

**Framework Synthesis**:

- **Foundation**: Enterprise architecture standards (SAD format)
- **Visual Approach**: C4 Model hierarchy for architectural views
- **SDLC Integration**: Direct PRD traceability and Epic/Story guidance
- **Quality Focus**: Comprehensive quality management and risk assessment

---

## Complete Template Structure

```markdown
# Solution Architecture Document (SAD): {{system_name}}

---

## Table of Contents
1. [Executive Summary](#1-executive-summary)
2. [Introduction](#2-introduction)
3. [Context](#3-context)
4. [Requirements](#4-requirements)
5. [Baseline Architecture](#5-baseline-architecture)
6. [Target Architecture](#6-target-architecture)
7. [Transition/Migration](#7-transitionmigration)
8. [Architectural Decisions](#8-architectural-decisions)
9. [Cross-Cutting Concerns](#9-cross-cutting-concerns)
10. [Quality Assurance](#10-quality-assurance)
11. [Appendices](#11-appendices)

---

## 1. Executive Summary

**Purpose**: Provide business-focused overview connecting architecture to business value and strategic objectives.

{{executive_summary}}

### Key Sections:
- **Business Context**: Link to Project Brief and business drivers
- **Architectural Approach**: High-level solution strategy and technology choices
- **Key Benefits**: Expected outcomes and value proposition
- **Success Metrics**: How architectural success will be measured

---

## 2. Introduction

**Purpose**: Establish foundation and context for all architectural decisions throughout the document.

### 2.1 Definitions, Acronyms, Abbreviations
{{definitions_acronyms_abbreviations}}

### 2.2 Scope
{{scope_boundaries}}

**What's Included:**
- {{included_systems}}

**What's Excluded:**
- {{excluded_systems}}

### 2.3 Stakeholders
{{stakeholders}}

| Stakeholder | Role | Key Concerns |
|-------------|------|--------------|
| {{stakeholder_name}} | {{stakeholder_role}} | {{stakeholder_concerns}} |

### 2.4 PRD Requirements Mapping
{{prd_requirements_mapping}}

**Requirements Traceability:**

| PRD Requirement | Architectural Component | Implementation Approach |
|-----------------|------------------------|------------------------|
| {{requirement_id}} ({{requirement_description}}) | {{component_name}} | {{implementation_approach}} |

**Purpose**: Establish direct traceability from PRD business/system requirements to architectural decisions.

---

## 3. Context

**Purpose**: Establish comprehensive environmental understanding across all architectural domains.

### 3.1 Technology Strategy
{{technology_strategy}}

**Organizational Technology Direction:**
- {{tech_strategy_point_1}}
- {{tech_strategy_point_2}}

**Alignment with Architecture:**
- {{alignment_point_1}}
- {{alignment_point_2}}

### 3.2 Business Architecture
{{business_architecture}}

**Key Business Processes:**
- {{business_process_1}}
- {{business_process_2}}

**Business Objectives:**
- {{business_objective_1}}
- {{business_objective_2}}

### 3.3 Data Architecture
{{data_architecture}}

**Data Models:**
- {{data_model_1}}
- {{data_model_2}}

**Data Flows:**
- {{data_flow_1}}
- {{data_flow_2}}

**Data Governance:**
- {{governance_requirement_1}}
- {{governance_requirement_2}}

### 3.4 Infrastructure Strategy
{{infrastructure_strategy}}

**Organizational Infrastructure Direction:**
- {{org_infrastructure_1}}
- {{org_infrastructure_2}}

**Strategic Platform Decisions:**
- {{strategic_platform_1}}
- {{strategic_platform_2}}

### 3.5 Application Architecture
{{application_architecture}}

**Application Landscape:**
- {{application_1}}
- {{application_2}}

**Integration Points:**
- {{integration_point_1}}
- {{integration_point_2}}

### 3.6 Security Architecture
{{security_architecture}}

**Security Posture:**
- {{security_requirement_1}}
- {{security_requirement_2}}

**Compliance Requirements:**
- {{compliance_requirement_1}}
- {{compliance_requirement_2}}

---

## 4. Requirements

**Purpose**: Document comprehensive driving forces behind all architectural decisions.

### 4.1 Business Goals
{{business_goals}}

**Primary Objectives:**
1. {{business_goal_1}}
2. {{business_goal_2}}
3. {{business_goal_3}}

**Success Criteria:**
- {{success_criterion_1}}
- {{success_criterion_2}}

### 4.2 Functional Requirements
{{functional_requirements}}

**Core Functionality:**

| Requirement ID | Description | Priority | Architectural Impact |
|----------------|-------------|----------|---------------------|
| {{req_id}} | {{req_description}} | {{priority}} | {{arch_impact}} |

### 4.3 Non-Functional Requirements
{{non_functional_requirements}}

**Quality Attributes:**

| Quality Attribute | Requirement | Measurement | Architectural Approach |
|-------------------|-------------|-------------|----------------------|
| Performance | {{performance_req}} | {{performance_metric}} | {{performance_approach}} |
| Scalability | {{scalability_req}} | {{scalability_metric}} | {{scalability_approach}} |
| Security | {{security_req}} | {{security_metric}} | {{security_approach}} |
| Availability | {{availability_req}} | {{availability_metric}} | {{availability_approach}} |

### 4.4 Constraints
{{constraints}}

**Technical Constraints:**
- {{technical_constraint_1}}
- {{technical_constraint_2}}

**Business Constraints:**
- {{business_constraint_1}}
- {{business_constraint_2}}

**Resource Constraints:**
- {{resource_constraint_1}}
- {{resource_constraint_2}}

### 4.5 Assumptions
{{assumptions}}

**Technical Assumptions:**
- {{technical_assumption_1}}
- {{technical_assumption_2}}

**Business Assumptions:**
- {{business_assumption_1}}
- {{business_assumption_2}}

---

## 5. Baseline Architecture

**Purpose**: Document current state architecture to understand transformation scope and complexity.

### 5.1 Conceptual View
{{baseline_conceptual_view}}

**Current System Overview:**
- {{current_system_1}}
- {{current_system_2}}

### 5.2 Logical View
{{baseline_logical_view}}

**Current System Components:**

| Component | Current Purpose | Technology Stack | Condition | Migration Notes |
|-----------|----------------|------------------|-----------|-----------------|
| {{current_component}} | {{current_purpose}} | {{current_tech}} | {{condition}} | {{migration_notes}} |

### 5.3 Integration View
{{baseline_integration_view}}

**Current Integration Points:**
- {{integration_1}}
- {{integration_2}}

### 5.4 Physical/Deployment View
{{baseline_deployment_view}}

**Current Infrastructure:**
- {{infrastructure_1}}
- {{infrastructure_2}}

---

## 6. Target Architecture

**Purpose**: Define comprehensive future state architecture with clear implementation guidance.

### 6.1 Conceptual View (C4 Context Level)
{{target_conceptual_view}}

**System Context Diagram:**

[Diagram showing system boundaries and external interactions using C4 Context level]


**External Systems:**
- {{external_system_1}}: {{relationship_1}}
- {{external_system_2}}: {{relationship_2}}

### 6.2 Logical View (C4 Container/Component)
{{target_logical_view}}

**Container Diagram (C4 Level 2):**

[Diagram showing high-level system decomposition]

**Component Diagram (C4 Level 3):**

[Diagram showing internal component structure]

**Key Components:**

| Component | Responsibility | Technology | Interfaces |
|-----------|---------------|------------|------------|
| {{component_name}} | {{responsibility}} | {{technology}} | {{interfaces}} |

### 6.3 Integration View
{{target_integration_view}}

**API Design:**
- {{api_1}}: {{api_purpose_1}}
- {{api_2}}: {{api_purpose_2}}

**Integration Patterns:**
- {{pattern_1}}: {{pattern_usage_1}}
- {{pattern_2}}: {{pattern_usage_2}}

### 6.4 Data View
{{target_data_view}}

**Data Architecture:**
- {{data_component_1}}: {{data_purpose_1}}
- {{data_component_2}}: {{data_purpose_2}}

**Data Flow:**

[Diagram showing information flow through the system]

### 6.5 Physical/Deployment View
{{target_deployment_view}}

**Deployment Architecture:**
- {{deployment_component_1}}: {{deployment_purpose_1}}
- {{deployment_component_2}}: {{deployment_purpose_2}}

**Infrastructure Mapping:**

[Diagram showing infrastructure deployment topology]

### 6.6 Quality Attributes Implementation
{{quality_implementation}}

**Architecture Quality Approaches:**

| Quality Attribute | Implementation Strategy | Architectural Pattern | Validation Method |
|-------------------|------------------------|---------------------|-------------------|
| Performance | {{performance_strategy}} | {{performance_pattern}} | {{performance_validation}} |
| Scalability | {{scalability_strategy}} | {{scalability_pattern}} | {{scalability_validation}} |
| Security | {{security_strategy}} | {{security_pattern}} | {{security_validation}} |
| Availability | {{availability_strategy}} | {{availability_pattern}} | {{availability_validation}} |

### 6.7 Risks and Mitigations
{{risks_mitigations}}

| Risk | Impact | Probability | Mitigation Strategy |
|------|--------|-------------|-------------------|
| {{risk_1}} | {{impact_1}} | {{probability_1}} | {{mitigation_1}} |
| {{risk_2}} | {{impact_2}} | {{probability_2}} | {{mitigation_2}} |

  ### 6.8 Solution Strategy
{{solution_strategy}}

**Architectural Principles:**
1. {{principle_1}}: {{principle_description_1}}
2. {{principle_2}}: {{principle_description_2}}

**Technology Decisions:**

| Decision Area | Chosen Technology | Rationale | Trade-offs |
|---------------|-------------------|-----------|------------|
| {{decision_area_1}} | {{technology_1}} | {{rationale_1}} | {{tradeoffs_1}} |
| {{decision_area_2}} | {{technology_2}} | {{rationale_2}} | {{tradeoffs_2}} |

**Architecture Patterns:**
- {{pattern_1}}: {{pattern_rationale_1}}
- {{pattern_2}}: {{pattern_rationale_2}}

---

## 7. Transition/Migration

**Purpose**: Define clear implementation roadmap from current state to target architecture.

### 7.1 Migration Approach
{{migration_approach}}

**Migration Strategy:**
- {{strategy_element_1}}
- {{strategy_element_2}}

**Migration Principles:**
- {{principle_1}}
- {{principle_2}}

### 7.2 Migration Roadmap
{{migration_roadmap}}

**Implementation Phases:**

| Phase | Duration | Scope | Dependencies | Success Criteria |
|-------|----------|-------|--------------|------------------|
| {{phase_1}} | {{duration_1}} | {{scope_1}} | {{dependencies_1}} | {{criteria_1}} |
| {{phase_2}} | {{duration_2}} | {{scope_2}} | {{dependencies_2}} | {{criteria_2}} |

### 7.3 Implementation Guidance
{{implementation_guidance}}

**Epic Breakdown Guidance:**

| Architectural Component | Epic Mapping | Story Creation Focus |
|------------------------|--------------|---------------------|
| {{component_1}} | {{epic_guidance_1}} | {{story_focus_1}} |
| {{component_2}} | {{epic_guidance_2}} | {{story_focus_2}} |

**Development Standards:**
- {{dev_standard_1}}
- {{dev_standard_2}}

**API Design Guidelines:**
- {{api_guideline_1}}
- {{api_guideline_2}}

**Testing Alignment:**
- {{testing_alignment_1}}
- {{testing_alignment_2}}

---

## 8. Architectural Decisions

**Purpose**: Record significant decisions with comprehensive rationale and alternatives analysis.

### 8.1 Decision Log (ADR Format)
{{architectural_decisions}}

**Decision Template:**

#### ADR-001: {{decision_title}}

**Status**: {{status}} (Proposed/Accepted/Superseded)
**Date**: {{decision_date}}
**Deciders**: {{decision_makers}}

**Context:**
{{decision_context}}

**Decision:**
{{decision_made}}

**Consequences:**
**Positive:**
- {{positive_consequence_1}}
- {{positive_consequence_2}}

**Negative:**
- {{negative_consequence_1}}
- {{negative_consequence_2}}

**Alternatives Considered:**
- {{alternative_1}}: {{alternative_rationale_1}}
- {{alternative_2}}: {{alternative_rationale_2}}

---

## 9. Cross-Cutting Concerns

**Purpose**: Address system-wide aspects that affect multiple architectural components.

### 9.1 Security
{{security_concerns}}

**Cross-Cutting Security Implementation:**
- {{cross_cutting_security_1}}
- {{cross_cutting_security_2}}

**Security Integration Points:**
- {{security_integration_1}}
- {{security_integration_2}}

### 9.2 Scalability
{{scalability_concerns}}

**Scalability Strategy:**
- {{scalability_strategy_1}}
- {{scalability_strategy_2}}

**Scaling Patterns:**
- {{scaling_pattern_1}}
- {{scaling_pattern_2}}

### 9.3 Observability
{{observability_concerns}}

**Monitoring Strategy:**
- {{monitoring_approach_1}}
- {{monitoring_approach_2}}

**Logging Strategy:**
- {{logging_approach_1}}
- {{logging_approach_2}}

### 9.4 Fault Tolerance
{{fault_tolerance_concerns}}

**Resilience Patterns:**
- {{resilience_pattern_1}}
- {{resilience_pattern_2}}

**Error Handling Strategy:**
- {{error_handling_1}}
- {{error_handling_2}}

---

## 10. Quality Assurance

**Purpose**: Define comprehensive approach to validating architectural decisions and implementation quality.

### 10.1 Testing Strategy
{{testing_strategy}}

**Architecture Testing Approach:**
- {{arch_testing_1}}
- {{arch_testing_2}}

**Quality Gates:**
- {{quality_gate_1}}
- {{quality_gate_2}}

### 10.2 Test Automation Approach
{{test_automation}}

**Automation Strategy:**
- {{automation_approach_1}}
- {{automation_approach_2}}

**Testing Tools:**
- {{testing_tool_1}}: {{tool_purpose_1}}
- {{testing_tool_2}}: {{tool_purpose_2}}

### 10.3 Quality Metrics
{{quality_metrics}}

**Success Metrics:**

| Quality Aspect | Metric | Target | Measurement Method |
|----------------|---------|--------|-------------------|
| {{quality_1}} | {{metric_1}} | {{target_1}} | {{method_1}} |
| {{quality_2}} | {{metric_2}} | {{target_2}} | {{method_2}} |

---

## 11. Appendices

**Purpose**: Provide supplementary materials and detailed references.

### 11.1 Glossary
{{glossary}}

| Term | Definition |
|------|------------|
| {{term_1}} | {{definition_1}} |
| {{term_2}} | {{definition_2}} |

### 11.2 Diagrams
{{diagrams}}

**Diagram Index:**
- {{diagram_1}}: {{diagram_purpose_1}}
- {{diagram_2}}: {{diagram_purpose_2}}

### 11.3 Reference Materials
{{reference_materials}}

**Standards and Guidelines:**
- {{reference_1}}
- {{reference_2}}

**Related Documentation:**
- {{related_doc_1}}
- {{related_doc_2}}
```

---

## Field Definitions and Best Practices

### Template Variable Guidelines

| Variable Pattern | Description | Example |
|-----------------|-------------|---------|
| `{{system_name}}` | Name of the system being architected | `User Authentication System` |
| `{{executive_summary}}` | High-level business and technical overview | `This architecture implements multi-factor authentication to reduce support tickets by 80% while maintaining 99.9% uptime...` |
| `{{requirement_id}}` | PRD requirement identifier | `BR1`, `NFR2` |
| `{{component_name}}` | Target architectural component identifier | `Authentication Service`, `User Database` |
| `{{current_component}}` | Baseline/current system component | `Legacy Login Module`, `Old User Store` |
| `{{quality_implementation}}` | Quality attribute implementation approach | `Performance achieved through caching strategy and load balancing...` |
| `{{cross_cutting_security_X}}` | Security concerns spanning multiple components | `Data encryption at rest and in transit`, `Authentication middleware` |

### SDLC Integration Guidelines

1. **PRD Requirements Mapping (Section 2.4)**: Always reference specific PRD requirements (BR1, BR2, NFR1, NFR2) and show how architectural decisions address them.

2. **Implementation Guidance (Section 7.3)**: Provide specific guidance for Epic and Story creation, including component-to-Epic mapping and development standards.

3. **C4 Visual Hierarchy (Sections 6.1-6.2)**: Use C4 model approach for architectural views to ensure clear visual communication.

4. **ADR Format (Section 8.1)**: Use structured Architecture Decision Records with Context, Decision, and Consequences for better decision tracking.

### Quality Validation Checklist

- [ ] All PRD requirements mapped to architectural components
- [ ] Epic breakdown guidance provided for each major component
- [ ] C4 diagrams included for system context and container views
- [ ] Architecture Decision Records follow ADR format
- [ ] Quality attributes addressed with specific approaches
- [ ] Risk mitigation strategies defined
- [ ] Implementation roadmap includes SDLC artifact dependencies

---

## Usage Instructions

1. **Copy Template**: Use the complete template structure as starting point
2. **Replace Variables**: Fill in all `{{variable}}` placeholders with project-specific content
3. **SDLC Integration**: Pay special attention to sections 2.4, 6.8, 7.3, and 8.1 for SDLC artifact flow
4. **Visual Elements**: Include C4 diagrams in sections 6.1 and 6.2 for clear communication
5. **Validation**: Use the quality checklist to ensure comprehensive coverage

This template provides the foundation for creating professional architecture documentation that seamlessly integrates with the SDLC framework while maintaining enterprise architecture standards.

---

## AI Agent Instructions for Architecture Document Creation

This section provides specific guidance for AI agents (LLMs) to effectively create and work with Solution Architecture Documents using this template.

### Core Agent Responsibilities

**When creating an Architecture Document, the AI agent must:**

1. **Parse Input Requirements**: Extract and understand PRD requirements, Epic definitions, and business context
2. **Apply Template Structure**: Use the complete 11-section template systematically
3. **Ensure SDLC Integration**: Create direct traceability and implementation guidance
4. **Validate Completeness**: Check all required sections and variables are addressed
5. **Maintain Professional Standards**: Follow enterprise architecture best practices

---

### Step-by-Step Agent Workflow

#### **Phase 1: Input Analysis** 🔍

```
1. ANALYZE PRD Requirements
   - Extract BR1, BR2, BR3... (Business Requirements)
   - Extract NFR1, NFR2, NFR3... (Non-Functional Requirements)
   - Identify constraints and assumptions

2. REVIEW Epic Definitions
   - Understand high-level features and business objectives
   - Identify component breakdown needs
   - Note Epic-to-Story guidance requirements

3. UNDERSTAND Business Context
   - Project Brief objectives and success metrics
   - Target users and stakeholder concerns
   - Technology strategy and constraints
```

#### **Phase 2: Template Population** 📝

```
1. START with Executive Summary
   - Link to Project Brief business drivers
   - Summarize architectural approach
   - Define success metrics

2. COMPLETE Introduction Sections (2.1-2.4)
   - 2.4 PRD Requirements Mapping is CRITICAL for SDLC integration
   - Create direct BR/NFR to component traceability

3. ESTABLISH Context (Section 3)
   - Connect to organizational strategy
   - Define data, infrastructure, and security context

4. DOCUMENT Requirements (Section 4)
   - Map PRD requirements to architectural needs
   - Use quality attributes table format

5. DEFINE Architectures (Sections 5-6)
   - Use C4 visual hierarchy approach
   - Section 6.8 Solution Strategy is CRITICAL for SDLC integration

6. PLAN Implementation (Section 7)
   - 7.3 Implementation Guidance is CRITICAL for Epic/Story creation
   - Provide specific Epic breakdown guidance

7. RECORD Decisions (Section 8)
   - Use ADR format for all major decisions

8. ADDRESS Cross-Cutting Concerns (Section 9)
   - Focus on system-wide aspects

9. DEFINE Quality Assurance (Section 10)
   - Connect to testing strategy

10. COMPLETE Appendices (Section 11)
    - Ensure comprehensive reference materials
```

#### **Phase 3: SDLC Integration Validation** ✅

```
1. VERIFY PRD Traceability (Section 2.4)
   - Every BR/NFR requirement mapped to architectural component
   - Clear implementation approach defined

2. VALIDATE Epic Breakdown Guidance (Section 7.3)
   - Each major architectural component has Epic mapping
   - Story creation focus areas defined
   - Development standards specified

3. CONFIRM C4 Visual Approach (Sections 6.1-6.2)
   - System context, container, and component views planned
   - Clear visual communication strategy

4. CHECK ADR Completeness (Section 8.1)
   - All major decisions documented with context, decision, consequences
   - Alternatives considered and rationale provided
```

---

### Critical Variable Replacement Guide

#### **Template Variables - Required Replacements**

| Variable | Agent Action | Example Output |
|----------|--------------|----------------|
| `{{system_name}}` | Extract from PRD/Epic context | `Multi-Factor Authentication System` |
| `{{executive_summary}}` | Synthesize business + technical overview | `This architecture implements biometric authentication to reduce support tickets by 80% while maintaining 99.9% uptime, serving 10,000+ users through Auth0 integration within a 3-month timeline.` |
| `{{requirement_id}}` | Reference specific PRD requirements | `BR1 (Biometric Authentication)`, `NFR2 (99.9% Uptime)` |
| `{{component_name}}` | Define architectural components | `Authentication Service`, `Biometric Verification Module`, `User Session Manager` |
| `{{epic_guidance_X}}` | Create Epic breakdown mapping | `Authentication Epic: Focus on login/logout workflows, biometric integration, session management` |
| `{{story_focus_X}}` | Define Story creation areas | `User biometric enrollment stories, authentication flow stories, error handling stories` |

#### **SDLC-Critical Variables - Mandatory Completion**

| Variable | Purpose | Agent Must Include |
|----------|---------|-------------------|
| `{{prd_requirements_mapping}}` | SDLC artifact traceability | Direct BR/NFR to component mapping table with implementation approach |
| `{{solution_strategy}}` | Bridge requirements to design | Architectural principles, technology decisions, pattern choices with rationale |
| `{{implementation_guidance}}` | Epic/Story creation support | Component-to-Epic mapping table, development standards, API guidelines |
| `{{architectural_decisions}}` | Decision tracking | ADR format with context, decision, consequences, alternatives |
| `{{quality_implementation}}` | Quality attribute implementation | Strategy, pattern, validation method for each quality attribute |

---

### Quality Validation Checklist for Agents

**Before completing the Architecture Document, AI agents must verify:**

#### **Content Completeness** ✅

- [ ] All 11 sections completed with relevant content
- [ ] All template variables replaced with project-specific content
- [ ] No `{{variable}}` placeholders remaining in final document
- [ ] Professional tone and technical accuracy maintained

#### **SDLC Integration** ✅

- [ ] Section 2.4: PRD Requirements Mapping completed with traceability table
- [ ] Section 6.8: Solution Strategy defines architectural approach
- [ ] Section 7.3: Implementation Guidance provides Epic breakdown support
- [ ] Section 8.1: ADR format used for decision documentation
- [ ] C4 visual approach referenced in sections 6.1-6.2

#### **Technical Quality** ✅

- [ ] Architecture addresses all PRD requirements (BR and NFR)
- [ ] Quality attributes have specific implementation approaches
- [ ] Cross-cutting concerns addressed comprehensively
- [ ] Risk mitigation strategies defined
- [ ] Migration approach includes phased implementation

#### **Professional Standards** ✅

- [ ] Consistent terminology and definitions used
- [ ] Clear stakeholder identification and concerns
- [ ] Comprehensive constraints and assumptions documented
- [ ] Reference materials and standards included

---

### Common Agent Pitfalls to Avoid

❌ **DON'T:**

- Leave any `{{variable}}` placeholders unfilled
- Copy/paste generic content without project customization
- Skip SDLC integration sections (2.4, 6.8, 7.3, 8.1)
- Create duplicate content between sections
- Use vague or non-specific language
- Ignore PRD requirements traceability

✅ **DO:**

- Customize all content to specific project context
- Create direct PRD-to-architecture traceability
- Provide actionable Epic/Story guidance
- Use professional architecture terminology
- Follow C4 visual hierarchy approach
- Document decisions with clear rationale

---

### Agent Success Criteria

**A successful Architecture Document created by an AI agent will:**

1. **Enable Epic Creation**: Product Owners can directly create Epics from Section 7.3 guidance
2. **Support Story Development**: Developers understand implementation approach from architecture
3. **Provide Decision Context**: All major technology and design decisions have clear rationale
4. **Ensure Quality**: Quality attributes have specific implementation and validation approaches
5. **Maintain Traceability**: Clear path from PRD requirements → Architecture → Epic → Story

**This instruction set ensures AI agents create professional, SDLC-integrated architecture documentation that drives successful product development.**

---

## SDLC Framework Integration Example

### Sample Agent Task Definition

**File**: `.krci-ai/tasks/create-architecture-document.md`

```markdown
# Task: Create Solution Architecture Document

## Description
Create comprehensive Solution Architecture Document that bridges PRD requirements to Epic/Story implementation using enterprise architecture standards with SDLC integration.

## Input Dependencies
- **Required**: PRD document with BR/NFR requirements defined
- **Required**: Epic definitions and business context
- **Optional**: Project Brief for strategic context

## Instructions

### Phase 1: Input Analysis
1. **Parse PRD Requirements**
   - Extract all Business Requirements (BR1, BR2, BR3...)
   - Extract all Non-Functional Requirements (NFR1, NFR2, NFR3...)
   - Identify constraints, assumptions, and success criteria

2. **Analyze Epic Context**
   - Understand high-level features and business objectives
   - Note component breakdown requirements
   - Identify Epic-to-Story guidance needs

3. **Review Business Context**
   - Project Brief objectives and success metrics
   - Target users and stakeholder concerns
   - Technology strategy and organizational constraints

### Phase 2: Architecture Document Creation
1. **Use Complete Template**
   - Follow [Architecture Document Template](./99-appendix-sad.md)
   - Complete all 11 sections systematically
   - Replace all template variables with project-specific content

2. **Focus on SDLC Integration**
   - **Section 2.4**: Create PRD Requirements Mapping table
   - **Section 6.8**: Define Solution Strategy with technology decisions
   - **Section 7.3**: Provide Epic Breakdown Guidance for development
   - **Section 8.1**: Use ADR format for architectural decisions

3. **Apply Quality Standards**
   - Use C4 visual hierarchy approach for architectural views
   - Address all quality attributes with implementation strategies
   - Ensure comprehensive risk mitigation strategies
   - Provide actionable development guidance

### Phase 3: Validation
1. **Content Completeness Check**
   - All sections completed with relevant content
   - No template variables remaining unfilled
   - Professional tone and technical accuracy maintained

2. **SDLC Integration Verification**
   - PRD requirements fully traced to architectural components
   - Epic breakdown guidance enables immediate Epic creation
   - Development standards and API guidelines provided
   - Quality assurance approach defined

## Output Format
- **File Name**: `architecture-v1.md` (increment version as needed)
- **Location**: `docs/architecture/`
- **Format**: Markdown following SAD template structure
- **Content**: Complete architecture document ready for Epic/Story creation

## Success Criteria
- [ ] Product Owner can create Epics directly from Section 7.3 guidance
- [ ] Developers understand implementation approach from architecture
- [ ] All major technology and design decisions have clear rationale
- [ ] Quality attributes have specific implementation approaches
- [ ] Clear traceability path: PRD → Architecture → Epic → Story

## Quality Gates
- Architecture addresses all PRD requirements (BR and NFR)
- SDLC integration sections (2.4, 6.8, 7.3, 8.1) completed
- Professional architecture standards maintained
- Epic/Story implementation guidance actionable
```

### Sample Agent Command Integration

**File**: `.krci-ai/agents/architect.yaml`

```yaml
agent:
  identity:
    name: "Solutions Architect"
    id: "architect-v1"
    role: "System design and technical architecture"
    goal: "Create comprehensive architecture that enables Epic/Story development"

  commands:
    help: "Show available architecture commands"
    chat: "(Default) Architecture consultation and guidance"
    create-architecture: "Create Solution Architecture Document from PRD and Epics"
    review-architecture: "Review and validate existing architecture document"
    update-architecture: "Update architecture based on new requirements"
    exit: "Exit architect persona"

  tasks:
    - "./.krci-ai/tasks/create-architecture-document.md"
    - "./.krci-ai/tasks/review-architecture-document.md"
    - "./.krci-ai/tasks/update-architecture-document.md"

  templates:
    - "./.krci-ai/templates/architecture-template.md"

  data:
    - "./.krci-ai/data/architecture-principles.yaml"
    - "./.krci-ai/data/technology-standards.yaml"
```

### Usage Example

```bash
# Activate architect agent
krci-ai architect

# Create architecture document from PRD and Epic inputs
architect: create-architecture
# Agent will:
# 1. Read PRD requirements (BR/NFR)
# 2. Analyze Epic definitions
# 3. Use SAD template to create comprehensive architecture
# 4. Generate Epic breakdown guidance
# 5. Output architecture-v1.md ready for development
```
