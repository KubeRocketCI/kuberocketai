# ADR-007: Agent Rules Hierarchy and Enforcement Strategy

## Status

Accepted

## Context

The KubeRocketAI framework requires clear architectural decisions about where and how to define behavioral rules for AI agents. The initial design question was whether rules should exist at both the Agent level and the Task level, and how to ensure absolute rule compliance "with no exceptions."

### Initial Questions

1. **Rule Location**: Should rules be defined only at the Agent level, or also at the Task level?
2. **Rule Hierarchy**: If both levels have rules, how should conflicts be resolved?
3. **Rule Enforcement**: How can we ensure AI agents follow rules with absolute compliance?
4. **Complexity Management**: How do we balance rule flexibility with cognitive simplicity for LLMs?

### Forces

- **Separation of Concerns**: Agents represent WHO (behavioral identity), Tasks represent WHAT (procedural workflow)
- **Rule Conflicts**: Multiple rule sources can create contradictions and unpredictable behavior
- **LLM Limitations**: AI models work better with simpler, clearer rule structures
- **Flexibility Needs**: Different tasks may require different types of constraints
- **Enforcement Reliability**: Rules must be followed consistently without exceptions

## Decision

**Implement Agent Principles with Task Rules for Clear User Guidance and Loose Coupling**

### Core Architecture

```yaml
# AGENT: Behavioral Principles (Clear User Guidance)
agent:
  principles:
    - "Always consider security implications first"
    - "Communicate with clear explanations and detailed reasoning"
    - "Ask clarifying questions when information is incomplete or ambiguous"
    - "Prefer proven, battle-tested technologies over experimental ones"
    - "Document all decisions with explicit trade-offs and rationale"

# TASK: Procedural Workflow
task: create-secure-api
instructions: |
  1. Analyze and validate input requirements
  2. Design API architecture with appropriate technology stack
  3. Specify exact versions for all technologies
  4. Document all architectural decisions and trade-offs
  5. Include security analysis and considerations
  6. Format using {{ ref:./.krci-ai/templates/api-design-output.md }}
```

### Clear Field Names Architecture for User Guidance

| Component | Field Name | Purpose | User Mental Model |
|-----------|------------|---------|------------------|
| **Agent** | `principles` | Personal beliefs and values | "How I think and what I believe" |
| **Task** | `instructions` | Procedural workflow steps | "What sequence of actions to take" |
| **Data** | `references` | Technical specifications and constraints | "Detailed requirements and validation" |

### Critical Architectural Insights

**Rejected: Task references to agent rules** to achieve:

- **True loose coupling**: Tasks are completely agent-agnostic
- **Maximum reusability**: Same task can be used by different agents
- **Behavioral diversity**: Different agents approach same task differently
- **No coupling dependencies**: Tasks and agents are independently maintainable

**Investigation: First-Person Framing for Agent Principles**

Research suggests first-person statements create better LLM persona embodiment:

| Framing | Example | LLM Behavior | Framework Status |
|---------|---------|--------------|------------------|
| **Imperative (Baseline)** | "Always prioritize security" | Clear directive, consistent behavior | ✅ Framework standard |
| **First-Person (Research)** | "I always prioritize security" | Natural personality expression, authentic character | 🔬 Requires investigation |

**Benefits of First-Person Principles:**

- Enhanced agent character consistency and embodiment
- More natural, personality-driven responses vs. checklist mentality
- Clearer semantic distinction between agent identity and operational constraints
- Better alignment with "principles" as personal beliefs rather than external commands

**Note**: This is **recommended based on investigation, not mandatory** - users can choose their preferred framing style.

### Future Research: First-Person Framing Investigation

**Research Topic**: First-person framing may improve LLM model performance for specific tasks

Research suggests that first-person statements ("I always prioritize security") may create better agent character consistency than imperative statements ("Always prioritize security") for some specific use cases.

**Potential Benefits for Investigation:**

- Enhanced LLM persona embodiment and character consistency
- More natural, personality-driven responses vs. mechanical rule-following
- Clearer distinction between agent identity and operational requirements
- Better alignment with "principles" as personal beliefs rather than external commands

**Investigation Status**: This requires separate research and validation outside the framework baseline. While the framework supports both styles, we are not making first-person framing the baseline approach until further investigation confirms its effectiveness across different models and use cases.

**Framework Position**: The KubeRocketAI framework uses imperative framing ("Always prioritize security") as the baseline, with support for first-person framing as a user choice pending further research validation.

## Consequences

### Positive

- **True Loose Coupling**: Tasks are completely independent of specific agents
- **Maximum Reusability**: Same task can be executed by multiple different agents
- **Behavioral Diversity**: Different agents bring unique perspectives to same tasks
- **Clear Separation**: Agent principles guide HOW, task instructions define WHAT
- **Ecosystem Scalability**: Tasks become shared assets across entire agent ecosystem
- **Reduced Maintenance**: No interdependencies between agents and tasks
- **Enhanced Agent Embodiment**: First-person principles improve character consistency (recommended)
- **User Guidance**: Clear field names eliminate confusion about where to put content

### Negative

- **Coordination Complexity**: Ensuring task quality without explicit agent constraints
- **Consistency Challenges**: Different agents may interpret same task differently
- **Learning Curve**: Users must understand principle vs. instruction distinction

### Neutral

- **Architecture Consistency**: Maintains clean separation between WHO (agent) and WHAT (task)
- **Migration Impact**: Existing agent rules need renaming to principles

## Implementation

### Rule Enforcement Strategies

#### 1. Simple, Absolute Rules

```yaml
# ✅ OPTIMAL: Clear, unconditional rules
rules:
  - "ALWAYS ask for clarification when requirements are unclear"
  - "NEVER suggest technology without version numbers"
  - "ALWAYS document architectural decisions with trade-offs"

# ❌ AVOID: Conditional or ambiguous rules
rules:
  - "Sometimes consider security, but prioritize speed if deadline is tight"
  - "Usually ask clarifying questions unless context seems obvious"
```

#### 2. Built-in Rule Validation

```yaml
task: create-system-design
instructions: |
  1. Design the system architecture
  2. VALIDATE: Check security measures included (Agent Principle compliance)
  3. VALIDATE: Confirm decisions documented with trade-offs
  4. VALIDATE: Verify technologies include version numbers
  5. Output results using template
```

#### 3. Multi-Level Reinforcement

```yaml
# Level 1: Agent Principle
principles:
  - "Always prioritize security first"

# Level 2: Task Instruction (reinforces agent principle)
instructions: |
  1. Apply security-first design approach
  2. Follow security checklist in {{ ref:./.krci-ai/data/validation/security.md }}

# Level 3: Data Reference (detailed requirements)
# File: ./.krci-ai/data/validation/security.md
# Contains: Specific security requirements and validation steps
```

### Framework Integration

- **Agent Definition**: Principles array with 5-7 maximum behavioral guidelines
- **Task Instructions**: Include validation as explicit workflow steps
- **Data References**: Provide detailed constraint specifications and checklists
- **Template Variables**: Can include rule compliance confirmation fields

### Example Implementation

#### Agent Principles (Behavioral Guidance)

```yaml
# Security-Focused Architect
agent:
  name: security-architect
  id: security-architect-v1
  version: "1.0.0"
  principles:
    - "Prioritize security above all other architectural concerns"
    - "Always identify and mitigate potential attack vectors"
    - "Default to defense-in-depth security strategies"
    - "Communicate security implications clearly"

# Performance-Focused Architect
agent:
  name: performance-architect
  id: performance-architect-v1
  version: "1.0.0"
  principles:
    - "Optimize for speed and efficiency in all design decisions"
    - "Always consider scalability and performance implications"
    - "Prefer lightweight, high-performance solutions"
    - "Measure and optimize based on concrete metrics"
```

#### Loosely Coupled Task

```yaml
# Same task usable by both agents above
task: create-microservice-design
description: Design scalable microservice architecture
instructions: |
  1. Analyze and validate input requirements
  2. Review relevant architecture patterns in {{ ref:./.krci-ai/data/docs/patterns.md }}
  3. Design microservice architecture with appropriate technology stack
     - Apply specifications from {{ ref:./.krci-ai/data/specs/microservice-requirements.yaml }}
     - Follow scaling guidelines from {{ ref:./.krci-ai/data/specs/scalability-patterns.yaml }}
  4. Specify exact versions for all technologies
  5. Document all architectural decisions and trade-offs
  6. Format using {{ ref:./.krci-ai/templates/microservice-design.md }}
```

#### Execution Results (Same Task, Different Approaches)

**Security Architect + Microservice Task:**

- Emphasizes secure communication patterns, authentication flows
- Focuses on input validation, encryption, access controls
- Documents threat model and security boundaries

**Performance Architect + Microservice Task:**

- Emphasizes caching strategies, load balancing, optimization
- Focuses on response times, throughput, resource efficiency
- Documents performance benchmarks and scaling thresholds

Both execute the same task but with completely different architectural perspectives!

## Related ADRs

- [ADR-001: Agent Data Model Enhancement](001-agent-data-model-enhancement.md)
- [ADR-003: Component Separation Model](003-component-separation-model.md)
- [ADR-005: Agent Command Interface](005-agent-command-interface.md)
- [ADR-006: Data Consolidation Strategy](006-data-consolidation-strategy.md)

---
*Date: January 20, 2025*
*Decision: Agent Principles Field for Clear User Guidance*
*Architecture: Different field names for different purposes - agent `principles` for behavior, task `rules` for constraints*
*Investigation: First-person framing recommended for better LLM persona embodiment (not mandatory)*
