# Task: Create Scope of Work (SoW)

## Description

Create a detailed Scope of Work document that defines the work to be performed, deliverables, acceptance criteria, and boundaries for the project. The SoW provides a comprehensive description of project scope, building on the high-level scope defined in the project charter. This document serves as the foundation for detailed project planning and work breakdown structure development.

## Prerequisites

- [ ] Project charter approved and available at `/docs/project-management/project-charter.md`
- [ ] Stakeholder requirements gathered and analyzed
- [ ] Project objectives and constraints clearly defined
- [ ] Initial project team identified
- [ ] Budget and timeline parameters established

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/project-management-methodology.md
- ./.krci-ai/templates/sow-template.md
- /docs/project-management/project-charter.md (approved charter)

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Reference project charter**: Build on scope defined in approved project charter
2. **Follow PMBoK scope management**: Apply PMBoK scope management processes and best practices
3. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for artifact dependencies
4. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
5. **Format output**: Use [sow-template.md](./.krci-ai/templates/sow-template.md) for consistent structure

## Output Format

- **Location**: `/docs/project-management/scope-of-work.md` (EXACT path and filename)
- **Length**: 5-10 pages depending on project complexity
- **Downstream Enable**: Enables detailed project plan and work breakdown structure creation

## Success Criteria

- [ ] **File saved** to `/docs/project-management/scope-of-work.md`
- [ ] **Project scope statement** clearly defined with boundaries
- [ ] **Deliverables** specified with acceptance criteria
- [ ] **Work breakdown structure** (WBS) outlined to appropriate level
- [ ] **Scope exclusions** explicitly documented
- [ ] **Assumptions and constraints** clearly stated
- [ ] **Change control process** defined for scope changes
- [ ] **Acceptance criteria** defined for all major deliverables
- [ ] **Dependencies** on external factors identified

## Execution Checklist

### Analysis Phase

- [ ] **Charter review**: Analyze project charter for scope foundation
- [ ] **Stakeholder input**: Gather detailed requirements from key stakeholders
- [ ] **Scope definition**: Define what is included and excluded from project scope
- [ ] **Deliverable identification**: Identify all project deliverables and outputs

### Planning Phase

- [ ] **WBS development**: Create work breakdown structure to appropriate detail level
- [ ] **Acceptance criteria**: Define clear acceptance criteria for each deliverable
- [ ] **Dependency mapping**: Identify dependencies between work packages and external factors
- [ ] **Constraint analysis**: Document all project constraints and limitations

### Documentation Phase

- [ ] **SoW creation**: Use [sow-template.md](./.krci-ai/templates/sow-template.md) structure
- [ ] **Content validation**: Ensure all sections are complete and consistent
- [ ] **Stakeholder review**: Validate scope with key stakeholders
- [ ] **File placement**: Save to exact location `/docs/project-management/scope-of-work.md`

### Approval Phase

- [ ] **Team review**: Review SoW with project team for feasibility
- [ ] **Stakeholder approval**: Obtain formal approval from key stakeholders
- [ ] **Baseline establishment**: Establish scope baseline for change control
- [ ] **Communication**: Communicate approved scope to all stakeholders

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Scope Clarity**: Define scope boundaries clearly to prevent scope creep
- **Deliverable Specificity**: Specify deliverables with measurable acceptance criteria
- **WBS Completeness**: Ensure work breakdown covers all required work packages
- **Change Control**: Establish clear processes for managing scope changes

### LLM Error Prevention Checklist

- **Avoid**: Vague scope descriptions that allow multiple interpretations
- **Avoid**: Missing acceptance criteria for deliverables
- **Avoid**: Incomplete work breakdown structure missing critical work packages
- **Avoid**: Undefined change control processes
- **Reference**: Use [sow-template.md](./.krci-ai/templates/sow-template.md) for all formatting guidance

### PMBoK Integration Context

The Scope of Work follows PMBoK scope management processes by defining project scope statement, creating work breakdown structure, establishing scope baseline, and defining change control procedures. This SoW enables creation of detailed project plans, schedules, and resource assignments based on the defined work packages. 