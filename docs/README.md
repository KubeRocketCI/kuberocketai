# KubeRocketAI Documentation

**AI-as-Code framework for enterprise development teams** - Define, version-control, and scale AI agents like infrastructure.

## Get Started Now

### [Quick Start Guide](quick-start.md)

**3 minutes to working AI agents** - Install, configure, and run your first AI-assisted workflow.

*Perfect for:* First-time users, proof-of-concept, immediate productivity

---

## Core Documentation

### [Core Concepts](concepts.md)

**Understand the framework** - AI-as-Code principles, value propositions, and how it solves scaling problems.

*Perfect for:* Understanding the "why", strategic decision making, team alignment

### [Architecture Overview](architecture.md)

**How it works under the hood** - System design, components, SDLC integration, and deployment modes.

*Perfect for:* Technical understanding, integration planning, architecture decisions

---

## Key Benefits

✅ **Project Context Awareness** - AI agents understand your architecture, standards, and dependencies
✅ **Multi-Platform Deployment** - Same agents work in IDE, CI/CD, and web chat tools
✅ **Team Scalability** - Version-control and share agent configurations like infrastructure
✅ **Quality Assurance** - Built-in validation for agent configurations
✅ **Offline Operation** - Complete framework embedded in CLI, no network dependencies

## Framework Overview

### 6 Production-Ready Agents

| Agent | Role | Primary Use Cases |
|-------|------|-------------------|
| **PM** | Product Manager | Strategy, requirements, PRDs, roadmaps |
| **PO** | Product Owner | Backlog management, story creation, sprint planning |
| **BA** | Business Analyst | Requirements analysis, workflows, documentation |
| **Architect** | System Architect | Technical design, architecture decisions, reviews |
| **Developer** | Software Engineer | Code implementation, feature development, reviews |
| **QA** | Quality Engineer | Testing strategies, quality assurance, validation |

### Two Deployment Modes

**IDE Integration** - Lightweight, development-focused:

```bash
krci-ai install --ide=cursor    # Focused on coding tasks
```

**Web Chat Bundles** - Complete project context:

```bash
krci-ai bundle --all --output brainstorm.md    # Strategic discussions
```

## Installation Options

```bash
# macOS - Homebrew (Recommended)
brew tap KubeRocketCI/homebrew-tap
brew install krci-ai

# Linux - Direct Download
curl -L "https://github.com/KubeRocketCI/kuberocketai/releases/latest/download/krci-ai_Linux_x86_64.tar.gz" | tar -xz

# Windows - Download from releases
# https://github.com/KubeRocketCI/kuberocketai/releases/latest
```

## Quick Validation

```bash
# Verify everything works
krci-ai validate

# See available agents
krci-ai list agents

# Get version info
krci-ai version
```

## Community & Support

- **GitHub Repository**: [KubeRocketCI/kuberocketai](https://github.com/KubeRocketCI/kuberocketai)
- **Issues & Bug Reports**: [GitHub Issues](https://github.com/KubeRocketCI/kuberocketai/issues)
- **Feature Requests**: [GitHub Discussions](https://github.com/KubeRocketCI/kuberocketai/discussions)
- **License**: Apache-2.0

---

## Navigation Guide

**New to KubeRocketAI?** → Start with [Quick Start](quick-start.md)
**Need to understand the value prop?** → Read [Core Concepts](concepts.md)
**Want technical details?** → Check [Architecture](architecture.md)

**Questions?** Check our [GitHub Issues](https://github.com/KubeRocketCI/kuberocketai/issues) or create a new one.
