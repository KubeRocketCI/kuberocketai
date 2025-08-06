# ADR-005: Agent Command Interface Design

## Status

Accepted

## Context

The KubeRocketAI framework needed a user-friendly interface for discovering and activating agent capabilities. The challenge was providing intuitive discovery without complex setup or documentation maintenance.

### Forces

- **Discoverability**: Users should easily find available agent capabilities
- **Simplicity**: Interface should be simple and familiar
- **Self-Documentation**: Agents should document their own capabilities
- **Consistency**: All agents should provide the same interface pattern
- **Minimal Overhead**: Interface shouldn't add significant complexity

### Alternative Approaches Considered

1. **External Documentation**: Maintain separate capability documentation
2. **Complex Discovery API**: Rich discovery mechanism with metadata
3. **Command-Based Interface**: Simple command mappings (chosen)
4. **Natural Language Only**: Pure conversational interface

## Decision

We implemented a simple command-based interface inspired by makefile help patterns:

### Interface Pattern

```yaml
commands:                         # Available command mappings
  help: string                    # Always present - shows available commands
  chat: string                    # Always present - default conversational mode
  [command_name]: string → uses: task_name  # Clean one-liner format
```

### Standard Commands

- **`help`**: Always present - shows numbered list of available commands
- **`chat`**: Always present - default mode for open conversation

### User Experience Flow

```bash
User: "*help"
Agent: Shows numbered command list from commands data

User: "*design" or "2"
Agent: Maps to task identifier and executes with agent behavioral context
```

### Example Implementation

```yaml
agent:
  name: software-architect
  id: software-architect-v1
  commands:
    help: "Show available commands"
    chat: "(Default) Architectural consultation"
    analyze: "Analyze requirements" → uses: analyze-requirements
    design: "Create system design" → uses: create-system-design
    review: "Review architecture" → uses: review-architecture
  tasks:
    - analyze-requirements
    - create-system-design
    - review-architecture
```

## Consequences

### Positive

- **Familiar Pattern**: Similar to `make help` - developers understand immediately
- **Self-Documenting**: Each agent contains its own capability definition
- **Natural Discovery**: Users can explore agent capabilities interactively
- **Consistent Experience**: All agents provide the same interface pattern
- **No External Docs**: Eliminates need for separate capability documentation
- **Scalable**: Framework scales through simple agent addition

### Negative

- **Command Namespace**: Need to avoid conflicts between command names
- **Learning Curve**: Users must learn the `*command` syntax
- **Limited Metadata**: Simple interface doesn't support rich command descriptions

### Neutral

- **Agent-Centric**: Commands defined at agent level, not globally
- **Task Mapping**: Commands map to existing task infrastructure
- **Validation Required**: Framework must validate command-to-task mappings

## Implementation

The interface is implemented through:

1. **Required Commands**: Every agent must define `help` and `chat`
2. **Simple Syntax**: `command: "description" → uses: task-name`
3. **Command Discovery**: Framework provides automatic help generation
4. **Task Integration**: Commands map to existing task definitions

### Benefits for Different Roles

| Role | Command Examples | Benefit |
|------|------------------|---------|
| **Architect** | `design`, `review`, `analyze` | Structured architectural workflows |
| **Developer** | `implement`, `test`, `refactor` | Development task automation |
| **Analyst** | `requirements`, `validate`, `document` | Business analysis workflows |

### Framework Integration

- Commands provide user interface layer over task execution
- Same task can be accessed through different agent commands
- Agent behavioral context influences task execution results
- Clean separation between interface (commands) and implementation (tasks)

This design provides the optimal balance of discoverability, simplicity, and consistency for agent interaction in the framework.

## Related ADRs

- [ADR-001: Agent Data Model Enhancement](001-agent-data-model-enhancement.md)
- [ADR-003: Four Component Separation Model](003-component-separation-model.md)
