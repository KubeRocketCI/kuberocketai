# Task: Review Story (Architect)

## Description

Review and validate user story from Architect perspective to ensure system design alignment, component boundaries, performance implications, and architectural standards compliance. Focus on system integration, scalability, security, and technical architecture consistency.

## Prerequisites

- [ ] **Story exists**: Target story file exists in `/docs/stories/` requiring architectural review
- [ ] **Architecture context**: Understanding of system architecture, design patterns, and technical standards
- [ ] **Design authority**: Architect expertise to validate system design and integration approach
- [ ] **System knowledge**: Familiarity with existing system components, interfaces, and constraints

### Reference Assets

Dependencies:

- ./.krci-ai/templates/story.md
- ./.krci-ai/data/common/sdlc-framework.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Validate system design alignment**: Ensure implementation approach aligns with overall system architecture
2. **Review component boundaries**: Verify proposed components have clear responsibilities and interfaces
3. **Assess integration patterns**: Validate integration approach follows established architectural patterns
4. **Check performance implications**: Evaluate performance, scalability, and resource considerations
5. **Verify compliance standards**: Ensure approach meets security, compliance, and technical standards; confirm AC include appropriate guardrails and that Verification Method/Evidence are adequate where commands aren't feasible

## Output Format

- **Location**: Update existing story file with architectural validation
- **Template**: Maintain [story.md](./.krci-ai/templates/story.md) structure (8 sections only)
- **Content Placement**: Architecture guidance in Description section, approval in Implementation Results
- **Architecture Compliance**: Document system design alignment and integration approach validation
- **Verification**: Story passes architect review with documented design approval

## Success Criteria

- [ ] **System architecture alignment**: Implementation approach consistent with overall system design
- [ ] **Component boundaries clear**: Clear separation of concerns and component responsibilities
- [ ] **Integration patterns appropriate**: Integration follows established architectural patterns
- [ ] **Performance considerations addressed**: Scalability, performance, and resource implications evaluated
- [ ] **Security standards compliance**: Implementation approach meets security and compliance requirements
- [ ] **Architect approval documented**: Architectural validation and design approval recorded

## Execution Checklist

### System Architecture Alignment

- [ ] **Design pattern consistency**: Implementation follows established system design patterns
- [ ] **Component integration**: New components integrate properly with existing system architecture
- [ ] **Service boundary respect**: Implementation respects established service and component boundaries
- [ ] **Data flow alignment**: Data flow and transformation aligns with system data architecture

### Component Design Validation

- [ ] **Separation of concerns**: Components have clear, single responsibilities
- [ ] **Interface definition**: Component interfaces are well-defined and consistent
- [ ] **Dependency management**: Component dependencies are minimal and well-structured
- [ ] **Reusability consideration**: Components designed for reusability where appropriate

### Integration Pattern Review

- [ ] **Communication patterns**: Inter-component communication follows established patterns
- [ ] **Event handling**: Event-driven patterns implemented correctly if applicable
- [ ] **Error propagation**: Error handling and propagation follows architectural standards
- [ ] **Transaction boundaries**: Transaction management aligns with system transaction patterns

### Performance and Scalability Assessment

- [ ] **Performance impact**: Implementation approach evaluated for performance implications
- [ ] **Scalability considerations**: Design supports system scalability requirements
- [ ] **Resource utilization**: Resource usage (memory, CPU, storage) appropriately planned
- [ ] **Bottleneck identification**: Potential performance bottlenecks identified and addressed

### Security and Compliance Validation

- [ ] **Security standards**: Implementation follows established security patterns and standards
- [ ] **Data protection**: Data handling and protection requirements appropriately addressed
- [ ] **Access control**: Authentication and authorization patterns correctly implemented
- [ ] **Compliance requirements**: Implementation meets relevant compliance and regulatory standards

### Technical Standards Compliance

- [ ] **Coding standards**: Technical approach aligns with established coding and design standards
- [ ] **Documentation patterns**: Technical documentation follows architectural documentation standards
- [ ] **Testing approach**: Testing strategy aligns with architectural testing patterns
- [ ] **Deployment considerations**: Implementation supports established deployment and operations patterns

### System Integration Readiness

- [ ] **API consistency**: New APIs follow established API design patterns and standards
- [ ] **Configuration management**: Configuration approach aligns with system configuration patterns
- [ ] **Monitoring integration**: Implementation supports established monitoring and observability patterns
- [ ] **Operational readiness**: Implementation approach supports operational requirements

## Content Guidelines

### Architectural Validation Principles for LLM Self-Evaluation

- **System Consistency**: All architectural decisions must align with overall system design and patterns
- **Component Clarity**: Component boundaries and responsibilities must be clear and well-defined
- **Integration Alignment**: Integration approaches must follow established architectural patterns
- **Standards Compliance**: Implementation must meet security, performance, and compliance standards

### LLM Error Prevention Checklist

- **Avoid**: Architectural decisions that conflict with established system design patterns
- **Avoid**: Component designs that violate separation of concerns or create tight coupling
- **Avoid**: Integration approaches that bypass established architectural patterns
- **Reference**: Ensure architectural alignment with system design standards and [story.md](./.krci-ai/templates/story.md) template
