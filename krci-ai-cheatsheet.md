# KubeRocketAI CLI Cheatsheet

> **🚀 KubeRocketAI** - AI-as-Code framework providing a centralized library of AI agents for software development workflows

## 🎯 Quick Start (3 Minutes)

### Installation

```bash
# macOS (Recommended)
brew tap KubeRocketCI/homebrew-tap
brew install krci-ai

# Linux - Direct download
curl -L "https://github.com/KubeRocketCI/kuberocketai/releases/latest/download/krci-ai_Linux_x86_64.tar.gz" | tar -xz
chmod +x krci-ai && sudo mv krci-ai /usr/local/bin/

# Windows - Download from releases
# https://github.com/KubeRocketCI/kuberocketai/releases/latest
```

### First Setup

```bash
# Install framework with IDE integration
krci-ai install --ide=claude    # For Claude Code
krci-ai install --ide=cursor    # For Cursor IDE
krci-ai install --all           # All IDE integrations

# Verify installation
krci-ai validate
```

---

## 📋 Command Reference

### `krci-ai install` - Framework Installation

Install AI-as-Code framework components to your project.

| Command | Description |
|---------|-------------|
| `krci-ai install` | Install core framework components |
| `krci-ai install --ide=cursor` | Install with Cursor IDE integration |
| `krci-ai install --ide=claude` | Install with Claude Code integration |
| `krci-ai install --ide=vscode` | Install with VS Code integration |
| `krci-ai install --ide=windsurf` | Install with Windsurf IDE integration |
| `krci-ai install --ide=all` | Install with all IDE integrations |
| `krci-ai install --all` | Install core + all IDE integrations (shortcut) |
| `krci-ai install --force` | Force installation (overwrite existing) |

**What Gets Installed:**

- `.krci-ai/agents/` - 6 role-based agent definitions (PM, Architect, Developer, QA, BA, PO)
- `.krci-ai/tasks/` - Common workflow templates
- `.krci-ai/templates/` - Output formatting templates
- `.krci-ai/data/` - Reference data and standards
- IDE-specific integration files (`.cursor/rules/`, `.claude/commands/`, etc.)

---

### `krci-ai list` - Component Discovery

Discover available framework components.

| Command | Description |
|---------|-------------|
| `krci-ai list agents` | List all available agents with roles |
| `krci-ai list agents -v` | List agents with dependency details |

**Available Agents:**

- **pm** - Senior Product Manager (Product strategy, requirements)
- **po** - Senior Product Owner (User story creation, backlog)
- **ba** - Senior Business Analyst (Requirements gathering)
- **architect** - Senior Software Architect (System design, architecture)
- **dev** - Software Developer (Implementation, code assistance)
- **qa** - Senior QA Engineer (Testing strategy, quality assurance)

---

### `krci-ai validate` - Framework Validation

Validate framework components and detect issues.

| Command | Description |
|---------|-------------|
| `krci-ai validate` | Standard validation with summary |
| `krci-ai validate -v` | Verbose validation with detailed insights |
| `krci-ai validate -q` | Quiet mode - minimal output |

**Validation Checks:**

- ✅ Agent YAML schema compliance
- ✅ Task path link validation
- ✅ Template file structure
- ✅ Markdown links to framework files
- ✅ Cross-platform file accessibility
- ✅ Dependency analysis and circular dependency detection
- ✅ Orphaned file detection

---

### `krci-ai bundle` - Web Chat Integration

Generate agent bundles for web chat tools (ChatGPT, Gemini Pro, Claude Web).

| Command | Description |
|---------|-------------|
| `krci-ai bundle --all` | Generate complete bundle with all agents |
| `krci-ai bundle --agent pm,architect` | Generate targeted bundle with specific agents |
| `krci-ai bundle --agent pm --task create-prd` | Generate minimal bundle for specific task |
| `krci-ai bundle --all --output my-bundle.md` | Custom output filename |
| `krci-ai bundle --all --dry-run` | Preview bundle scope without generating |

**Bundle Contents:**

- All selected agent definitions with complete YAML
- Agent dependencies, tasks, and templates
- Project-specific context from `.krci-ai/data/` and `.krci-ai/templates/`
- System prompt structure with role definitions

---

### `krci-ai version` - Version Information

Display version and build information.

```bash
krci-ai version
# Output: version, commit, build date, Go version, platform
```

---

### `krci-ai check-updates` - Update Management

Check for available CLI updates.

```bash
krci-ai check-updates
# Queries GitHub API for latest releases
```

---

## 🎭 Agent Usage in IDEs

### Claude Code Integration

After installing with `--ide=claude`:

```bash
/pm           # Activate Product Manager
/architect    # Activate Software Architect
/dev          # Activate Developer
/qa           # Activate QA Engineer
/ba           # Activate Business Analyst
/po           # Activate Product Owner
```

### Agent Commands

Each agent provides these standard commands:

- `help` - Show available commands
- `chat` - (Default) Consultation and assistance
- `exit` - Exit agent persona

**Agent-Specific Commands:**

- **PM**: `create-project-brief`, `create-prd`
- **Architect**: `design-system`, `review-architecture`
- **Developer**: `review`, `plan-implementation`, `implement`
- **QA**: `create-test-plan`, `review-test-coverage`

---

## 📁 Directory Structure

```bash
project-root/
├── .krci-ai/
│   ├── agents/          # 6 role-based agent definitions
│   │   ├── pm.yaml
│   │   ├── architect.yaml
│   │   ├── dev.yaml
│   │   ├── qa.yaml
│   │   ├── ba.yaml
│   │   └── po.yaml
│   ├── tasks/           # Common workflow templates
│   ├── templates/       # Output formatting templates
│   ├── data/           # Reference data and standards
│   └── bundle/         # Generated bundles (from bundle command)
├── .cursor/rules/      # Cursor IDE integration (if --ide=cursor)
├── .claude/commands/   # Claude Code integration (if --ide=claude)
└── .github/chatmodes/ # VS Code integration (if --ide=vscode)
```

---

## 🔧 Common Workflows

### 1. Project Setup

```bash
# Install and setup
krci-ai install --ide=claude
krci-ai validate

# List available agents
krci-ai list agents
```

### 2. Create Project Brief

```bash
# In Claude Code
/pm
create-project-brief
# Provide context about your project
```

### 3. Architecture Design

```bash
# In Claude Code
/architect
design-system
# Describe technical requirements
```

### 4. Implementation

```bash
# In Claude Code
/dev
implement
# Describe features to implement
```

### 5. Web Chat Bundle

```bash
# Generate bundle for ChatGPT/Claude Web
krci-ai bundle --all --output project-context.md
# Upload project-context.md to web chat tools
```

---

## 🛠️ Troubleshooting

### Common Issues

```bash
# Permission denied
sudo krci-ai install --ide=cursor

# Validate installation
krci-ai validate -v

# Force reinstall
krci-ai install --force

# Check for updates
krci-ai check-updates
```

### Validation Errors

- **Broken links**: Check `.krci-ai/` file paths in agent YAML
- **Missing tasks**: Ensure task files exist in `.krci-ai/tasks/`
- **Schema errors**: Validate agent YAML structure

### IDE Integration Issues

- Ensure IDE integration was installed: `krci-ai install --ide=your-ide`
- Check for IDE-specific files in project root
- Restart IDE after installation

---

## 📚 Additional Resources

- **Documentation**: [GitHub Repository](https://github.com/KubeRocketCI/kuberocketai)
- **Issues & Support**: [GitHub Issues](https://github.com/KubeRocketCI/kuberocketai/issues)
- **Architecture**: See `docs/architecture.md` in repository
- **Quick Start**: See `docs/quick-start.md` in repository

---

## 🚀 Advanced Usage

### Development Commands

```bash
# If you're developing KubeRocketAI itself
make build          # Build CLI binary
make test           # Run unit tests with coverage
make lint           # Run linting (golangci-lint)
make fmt            # Format code
make ci             # Run full CI pipeline locally
```

### Schema Validation

KubeRocketAI uses JSON schemas for validation:

- Agent YAML files follow strict schema definitions
- Template and task files are validated for structure
- Cross-references are verified for integrity
