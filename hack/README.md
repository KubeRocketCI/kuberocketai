# KubeRocketAI Validation Tools

This directory contains development and validation tools for the KubeRocketAI framework.

## Agent Validation Tool

### Usage

```bash
# Validate a single agent
./hack/validate-agents.py assets/framework/core/agents/architect.yaml

# Validate all agents
./hack/validate-agents.py --all
```

### Features

- **Automatic Environment Setup**: Creates `.venv` and installs dependencies automatically
- **JSON Schema Validation**: Validates against `assets/schemas/agent-schema.json`
- **Framework Rule Checking**: Validates KubeRocketAI-specific rules
- **Cross-Platform**: Works on macOS, Linux, and Windows

### Dependencies

The tool automatically manages its dependencies using Python virtual environments:

- `PyYAML==6.0.1` - YAML file parsing
- `jsonschema==4.17.3` - JSON Schema validation

### How It Works

1. **First Run**: Creates `.venv` directory and installs requirements
2. **Subsequent Runs**: Uses existing virtual environment
3. **Validation**: Runs comprehensive validation checks:
   - JSON Schema compliance (with identity grouping)
   - Required command validation
   - Task reference integrity
   - Framework rule compliance

### Output Examples

**Valid Agent**:
```
✅ VALID: architect.yaml passes schema validation
```

**Invalid Agent**:
```
❌ INVALID: architect.yaml - Missing required 'help' command
```

**Multiple Agents**:
```
🔍 Validating 3 agent files...
✅ VALID: architect.yaml passes schema validation
✅ VALID: developer.yaml passes schema validation
❌ INVALID: qa.yaml - Missing required 'chat' command

❌ Some agents failed validation
```