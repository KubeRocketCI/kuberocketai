# ADR-003: Four Component Separation Model

## Status

Accepted

## Context

The KubeRocketAI framework required a clear component model that would enable both simplicity for basic use cases and composability for complex workflows. The challenge was defining distinct component types that work together while maintaining clear separation of concerns.

### Forces

- **Simplicity**: Framework should be easy to understand and use
- **Composability**: Components should mix and match for different use cases
- **Separation of Concerns**: Each component type should have a distinct purpose
- **Progressive Complexity**: Support simple use cases without requiring complex setup
- **Reusability**: Components should be shareable across different contexts

## Decision

We established a four-component separation model based on distinct questions each component answers:

| Component | Purpose | Question Answered | Data Representation |
|-----------|---------|------------------|-------------------|
| **Agents** | WHO | "Who is performing this work?" | Persona data with behavioral instructions |
| **Tasks** | WHAT | "What specific work needs to be done?" | Workflow data with procedural instructions |
| **Templates** | HOW | "How should the output be formatted?" | Output formatting structures with variables |
| **Data** | REFERENCE | "What reference information is needed?" | Static data including specs, schemas, guidelines |

### Component Characteristics

#### Agents (WHO)

- **Purpose**: Define personas, expertise, and behavioral approach
- **Contains**: Identity, role, goals, behavioral rules, available commands
- **Reusability**: Domain-specific (architect, developer, analyst)
- **Independence**: Can work alone for conversational mode

#### Tasks (WHAT)

- **Purpose**: Define specific workflows and procedures
- **Contains**: Step-by-step instructions, input/output specifications
- **Reusability**: High - same task used by different agents produces different results
- **Independence**: Cannot work alone - requires agent context

#### Templates (HOW)

- **Purpose**: Define output formatting and structure
- **Contains**: Static text with variable placeholders
- **Reusability**: High - shared across multiple tasks/agents
- **Independence**: Passive - only used when referenced

#### Data (REFERENCE)

- **Purpose**: Provide static reference materials and constraints
- **Contains**: API specs, schemas, documentation, guidelines
- **Reusability**: Highest - referenced by any component as needed
- **Independence**: Passive - only used when referenced

### Usage Patterns

The framework supports progressive complexity:

1. **Agent Only**: Pure conversational mode
2. **Agent + Task**: Structured workflows with natural language output
3. **Agent + Task + Template**: Structured workflows with formatted output
4. **Agent + Task + Template + Data**: Full framework with reference materials

## Consequences

### Positive

- **Clear Mental Model**: Each component has distinct purpose and boundaries
- **Progressive Enhancement**: Start simple, add complexity as needed
- **Maximum Reusability**: Tasks and templates shared across agents
- **Flexible Composition**: 9 different usage patterns supported
- **Natural User Interface**: Agents provide discoverable commands

### Negative

- **Learning Curve**: Users must understand four different component types
- **File Organization**: More files to manage compared to monolithic approach
- **Validation Complexity**: Framework must validate relationships between components

### Neutral

- **Directory Structure**: Requires specific folder organization (`./.krci-ai/agents/`, etc.)
- **Reference Management**: All connections made through file references

## Implementation

The model is implemented through:

1. **Required Directory Structure**:

   ```
   ./.krci-ai/
   ├── agents/     # WHO (personas)
   ├── tasks/      # WHAT (actions)
   ├── templates/  # HOW (formats)
   └── data/       # REFERENCE (static materials)
   ```

2. **Component Schemas**: Each component type has defined structure and validation rules

3. **Reference Patterns**: Clear rules about which components can reference others

4. **Progressive Examples**: Documentation shows all usage patterns from simple to complex

This separation enables the framework's core promise: composable, reusable components that can be mixed and matched for maximum flexibility.

## Related ADRs

- [ADR-002: Template Reference Pattern](002-template-reference-pattern.md)
- [ADR-006: Data Asset Consolidation Strategy](006-data-consolidation-strategy.md)
- [ADR-004: Inline Reference Pattern](004-inline-reference-pattern.md)
