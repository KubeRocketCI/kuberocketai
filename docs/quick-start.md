# Quick Start Guide

Get KubeRocketAI running with your first AI agent in 3 minutes.

## The Problem We Solve

You've figured out how to make AI agents work, but now you need to **scale that success** across your team and projects. KubeRocketAI brings AI-as-Code principles to AI agent management.

## Install KubeRocketAI

### macOS (Recommended)

```bash
# Add the KubeRocketCI tap and install
brew tap KubeRocketCI/homebrew-tap
brew install krci-ai
```

### Linux/Windows

```bash
# Linux - Direct download
curl -L "https://github.com/KubeRocketCI/kuberocketai/releases/latest/download/krci-ai_Linux_x86_64.tar.gz" | tar -xz
chmod +x krci-ai && sudo mv krci-ai /usr/local/bin/

# Windows - Download from releases page
# https://github.com/KubeRocketCI/kuberocketai/releases/latest
```

## Your First 3 Minutes

### 1. Install Framework (20 seconds)

```bash
# Install with Claude Code integration
krci-ai install --ide=claude
```

**Output:**

```bash
✅ Claude Code integration installed successfully!
✅ Framework installation completed successfully!
Next steps:
  • Run 'krci-ai list agents' to see available agents
```

### 2. Meet Your Team (10 seconds)

```bash
krci-ai list agents
```

**Output:**

```bash
✅ Found 6 agent(s):
pm    | Senior Product Manager    | Product strategy, requirements
po    | Senior Product Owner      | User story creation, backlog
ba    | Senior Business Analyst  | Requirements gathering
architect | Senior Software Architect | System design, architecture
dev   | Software Developer        | Implementation, code assistance
qa    | Senior QA Engineer        | Testing strategy, quality assurance
```

### 3. Start Your First Project (2 minutes)

**In Claude Code, type:**

```bash
/pm
```

**Agent responds:**

```bash
Hello! I'm Peter Manager, your Senior Product Manager 📈

Available Commands:
- help - Show available commands
- create-project-brief - Create project brief
- create-prd - Create comprehensive PRD
- exit - Exit Product Manager persona

What would you like to work on today?
```

**Select task by name:**

```bash
create-project-brief
```

**Provide simple context:**

```bash
React SPA company visit page:
- Single page website for tech startup
- Hero section, services overview, contact form
- Modern design, mobile responsive
- Replace current outdated WordPress site
```

**Result:** Professional 2-page project brief created at `docs/prd/project-brief.md`

### 4. Validate Everything Works (10 seconds)

```bash
krci-ai validate
```

**Output:**

```
✅ FRAMEWORK VALID
📊 Overview: 6 agents, 24 tasks, 14 templates, 11 data files
⚡ Validation completed in 0.0s
```

### 5. Create Web Chat Bundle (30 seconds)

```bash
# Bundle agents for ChatGPT/Claude Web
krci-ai bundle --all --output project-context.md

# Now you can upload project-context.md to ChatGPT for strategic discussions
```

## Next Steps

🎯 **You're Ready!** Your AI agents understand your project context and organizational standards.

**What to do next:**

- Review [Core Concepts](concepts.md) to understand the framework
- Check [Architecture](architecture.md) to see how it all works
- Visit [Contributing](contributing.md) to customize and extend

## Troubleshooting

**Common Issues:**

```bash
# Permission denied
sudo krci-ai install --ide=cursor

# Validate your installation
krci-ai validate -v  # Verbose validation

# Update to latest version
brew upgrade krci-ai  # macOS
```

**Need Help?**

- [GitHub Issues](https://github.com/KubeRocketCI/kuberocketai/issues)
- [Documentation Hub](README.md)

---

⏱️ **That's it!** You now have a complete AI-as-Code framework running locally. Your agents are ready to help with everything from project planning to code implementation.
