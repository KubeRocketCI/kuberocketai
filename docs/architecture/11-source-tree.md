# 11. Source Tree

This section defines the comprehensive source code organization, project structure, and repository layout for the KubeRocketAI project to ensure clear separation of concerns and maintainable codebase.

## 11.1 Project Structure Overview

The KubeRocketAI project follows a **monolithic CLI architecture** with embedded LLM assets, organized for clear separation between:

- **CLI Tool Code** (Go implementation)
- **Framework Assets** (AI-as-Code components)
- **Build & Distribution** (Cross-platform deployment)
- **Documentation & Testing** (Development support)

## 11.2 Root Directory Structure

```bash
kuberocketai/
├── cmd/                           # CLI entry points and command definitions
├── internal/                      # Internal Go packages (not exported)
├── pkg/                           # Public Go packages (if needed)
├── assets -> cmd/krci-ai/assets   # Symlink to framework assets (backward compatibility)
├── configs/                       # Configuration files and templates
├── hack/                          # Build and development scripts
├── tests/                         # Test suites and test data
├── docs/                          # Architecture and user documentation
├── .krci-ai/                      # Framework dogfooding (developing with itself)
├── .cursor/                       # Cursor IDE configuration
├── .claude/                       # Claude Code configuration
├── .vscode/                       # VS Code GitHub Copilot configuration
├── .windsurf/                     # Windsurf IDE configuration
├── .editorconfig                  # Editor configuration for consistent formatting
├── .gitignore                     # Git ignore rules
├── .golangci.yml                  # Go linting configuration
├── .goreleaser.yaml               # Release automation configuration
├── .markdownlint.json             # Markdown linting configuration
├── .prettierrc                    # Prettier configuration for formatting
├── Makefile                       # Build and development tasks
├── go.mod                         # Go module definition
├── go.sum                         # Go module checksums
├── LICENSE                        # Project license
└── README.md                      # Project overview and quick start
```

## 11.3 CLI Tool Code Organization (`cmd/` & `internal/`)

### 11.3.1 Command Structure (`cmd/`)

```bash
cmd/
└── krci-ai/                       # Main CLI entry point
    ├── main.go                    # Application bootstrap with Go embed
    ├── assets/                    # Framework assets (embedded via Go embed)
    │   ├── framework/
    │   │   └── core/              # Core framework components
    │   │       ├── agents/        # Agent definitions (WHO)
    │   │       ├── tasks/         # Task definitions (WHAT)
    │   │       ├── templates/     # Template definitions (HOW)
    │   │       └── data/          # Data definitions (REFERENCE)
    │   └── schemas/               # Component schemas & validation
    └── cmd/                       # CLI commands
        ├── root.go                # Root command definition
        └── install.go             # Install command implementation
```

### 11.3.2 Internal Packages (`internal/`)

```bash
internal/
├── cli/                           # Command Interface component
│   ├── commands/                  # Command implementations
│   ├── ui/                        # User interface helpers
│   └── config/                    # CLI configuration
├── engine/                        # Framework Engine component
│   ├── orchestrator/              # Operation coordination
│   ├── processor/                 # Asset processing
│   └── validator/                 # Validation orchestration
├── assets/                        # Asset Management component
│   ├── embedded/                  # Embedded asset handling
│   ├── loader/                    # Asset loading logic
│   └── installer/                 # Local installation
├── validation/                    # Validation System component
│   ├── static/                    # Static validation
│   ├── runtime/                   # Runtime validation
│   └── schemas/                   # Schema definitions
├── integration/                   # Integration Layer component
│   ├── ide/                       # IDE adapter interfaces
│   └── adapters/                  # Specific IDE implementations
└── utils/                         # Utility Services component
    ├── errors/                    # Error handling
    ├── logging/                   # Logging utilities
    └── filesystem/                # File system operations
```

## 11.4 LLM Framework Assets (`assets/`)

### 11.4.1 Asset Location and Go Embed Strategy

Due to Go's embed directive limitations, the framework assets are physically located at `cmd/krci-ai/assets/` to satisfy the embed path requirements. A symlink is maintained at the root level (`assets -> cmd/krci-ai/assets`) to preserve backward compatibility with existing documentation and references.

**Key Design Decisions:**

- **Physical Location**: `cmd/krci-ai/assets/` - Required for Go embed directive (`//go:embed assets/framework/core`)
- **Symlink Strategy**: `assets -> cmd/krci-ai/assets` - Maintains backward compatibility
- **Embedding**: Assets are embedded at build time via Go embed for distribution in the CLI binary

### 11.4.2 Framework Components Structure

```bash
cmd/krci-ai/assets/                    # Framework assets (embedded via Go embed)
├── framework/                     # Framework asset collections
│   └── core/                      # Core framework (embedded in CLI)
│       ├── agents/                # Agent definitions (WHO)
│       │   ├── architect.yaml     # System architect persona
│       │   ├── developer.yaml     # Developer persona
│       │   ├── pm.yaml            # Product manager persona
│       │   ├── qa.yaml            # QA engineer persona
│       │   └── ba.yaml            # Business analyst persona
│       ├── tasks/                 # Task definitions (WHAT)
│       │   ├── create-prd.md       # Product requirements document
│       │   ├── create-system-design.md # System design task
│       │   ├── review-code.md      # Code review task
│       │   ├── create-test-plan.md # Test planning task
│       │   └── [additional tasks] # Other task definitions
│       ├── templates/             # Template definitions (HOW)
│       │   ├── prd-template.md    # Product requirements template
│       │   ├── system-design.md   # System design template
│       │   ├── code-review.md     # Code review template
│       │   ├── test-plan.md       # Test plan template
│       │   └── [additional templates] # Other template definitions
│       └── data/                  # Data definitions (REFERENCE)
│           ├── coding-standards.md # Coding standards reference
│           ├── architecture-principles.md # Architecture principles
│           ├── testing-standards.md # Testing guidelines
│           ├── product-principles.md # Product management principles
│           └── [additional data]  # Other reference data
└── schemas/                       # Component schemas & validation
    └── agent-schema.json          # Agent validation schema (JSON Schema)
```

### 11.4.3 Progressive Complexity Organization

The framework assets are organized to support progressive complexity levels:

- **Level 1** (Agent Only): `cmd/krci-ai/assets/framework/core/agents/` (accessible via `assets/framework/core/agents/`)
- **Level 2** (+ Tasks): `cmd/krci-ai/assets/framework/core/tasks/` (accessible via `assets/framework/core/tasks/`)
- **Level 3** (+ Templates): `cmd/krci-ai/assets/framework/core/templates/` (accessible via `assets/framework/core/templates/`)
- **Level 4** (+ Data): `cmd/krci-ai/assets/framework/core/data/` (accessible via `assets/framework/core/data/`)

**Framework Categories:**

- **Core Framework**: `cmd/krci-ai/assets/framework/core/` - Essential components embedded in CLI
- **Extensions**: `cmd/krci-ai/assets/framework/extensions/` - Domain-specific framework (future)
- **Community**: `cmd/krci-ai/assets/framework/community/` - Community-contributed framework (future)

**Note**: The symlink at the root level (`assets -> cmd/krci-ai/assets`) ensures existing documentation references remain valid while satisfying Go embed requirements.

## 11.5 Framework Dogfooding (`.krci-ai/`)

The KubeRocketAI project uses **itself** to develop itself - a practice called "dogfooding" that ensures the framework is tested in real-world scenarios and continuously improved based on actual usage.

### 11.5.1 Self-Development Framework

```bash
.krci-ai/
├── agents/                        # Project-specific agents
│   ├── architect.md               # System architect for KubeRocketAI
│   ├── developer.md               # Go developer for CLI
│   ├── pm.md                      # Product manager for features
│   ├── qa.md                      # QA engineer for testing
│   ├── framework-designer.md      # Framework design specialist
│   ├── asset-curator.md           # Asset management specialist
│   └── cli-architect.md           # CLI architecture specialist
├── tasks/                         # Development workflow tasks
│   ├── create-doc.md              # Documentation creation
│   ├── code-review.md             # Code review process
│   ├── release.md                 # Release management
│   ├── design-component.md        # Component design task
│   ├── validate-assets.md         # Asset validation task
│   └── build-cli.md               # CLI build process
├── templates/                     # Output formatting templates
│   ├── architecture-doc.md        # Architecture documentation
│   ├── go-package-doc.md          # Go package documentation
│   ├── test-plan.md               # Test plan template
│   ├── component-spec.md          # Component specification
│   ├── asset-schema.md            # Asset schema template
│   └── cli-help.md                # CLI help template
├── data/                          # Project knowledge base
│   ├── go-standards.md            # Go coding standards
│   ├── testing-guidelines.md      # Testing best practices
│   ├── security-rules.md          # Security requirements
│   ├── tech-stack.md              # Technology decisions
│   ├── architecture-principles.md # Architecture principles
│   └── framework-constraints.md   # Framework limitations
└── config/                        # Framework configuration
    ├── validation-rules.yaml      # Custom validation rules
    ├── asset-policies.yaml        # Asset management policies
    └── ide-settings.json          # IDE-specific settings
```

### 11.5.2 IDE Configuration Generation

The CLI automatically converts framework components to IDE-specific configurations. When developers run `krci-ai install --ide=cursor`, the tool generates `.cursor/` rules from embedded framework assets.

#### 11.5.2.1 Developer Role Example

**Source Assets** (embedded from `cmd/krci-ai/assets/framework/core/`):

```bash
agents/developer.yaml     # WHO: Senior Go developer persona
tasks/code-review.md      # WHAT: Perform code review task
templates/pr-review.md    # HOW: Code review output format
data/go-standards.md      # REFERENCE: Go coding guidelines
```

**Generated Configuration** (`.cursor/rules/developer.mdc`):

```markdown
---
description: Senior Go developer with code review expertise
globs: "**/*.go,**/*.md"
alwaysApply: false
---

# Developer Agent Context
You are a senior Go developer with expertise in CLI tools, embedded assets, and framework design.

## Code Review Task
Perform thorough code reviews focusing on:
- Go idioms and best practices
- Error handling patterns
- Performance considerations
- Test coverage and quality

## Output Format
Use structured PR review format:
- Summary of changes
- Issues found with severity
- Suggestions for improvement
- Approval status

## Reference Standards
Follow Go coding guidelines:
- Use table-driven tests
- Implement proper error handling
- Prefer composition over inheritance
- Document public APIs with examples
```

**CLI Magic**: `krci-ai install --ide=cursor` reads the 4 framework assets and combines them into a single, contextual IDE rule that gives AI assistants complete understanding of the developer role, tasks, output format, and reference standards.

#### 11.5.2.2 Multi-IDE Configuration Generation

The KubeRocketAI CLI automatically generates configurations for multiple AI-powered editors from the same framework assets.

##### Generated IDE Configurations

**Cursor IDE** (`.cursor/`):

```bash
.cursor/
├── rules/
│   ├── developer.mdc
│   ├── architect.mdc
│   ├── pm.mdc
│   ├── qa.mdc
│   ├── framework-designer.mdc
│   ├── asset-curator.mdc
│   └── global.mdc
├── instructions/
│   ├── code-review.md
│   ├── documentation.md
│   └── testing.md
└── .cursorrules
```

**Claude Code** (`.claude/`):

```bash
.claude/
├── project.md
├── agents/
│   ├── developer.md
│   ├── architect.md
│   └── qa.md
├── tasks/
│   ├── code-review.md
│   ├── documentation.md
│   └── testing.md
└── knowledge/
    ├── architecture.md
    ├── patterns.md
    └── standards.md
```

**VS Code GitHub Copilot** (`.vscode/`):

```bash
.vscode/
├── settings.json
├── tasks.json
├── launch.json
├── copilot/
│   ├── workspace.md
│   ├── patterns.md
│   └── instructions.md
└── extensions.json
```

**Windsurf IDE** (`.windsurf/`):

```bash
.windsurf/
├── project.md
├── rules/
│   ├── go-development.md
│   ├── framework-assets.md
│   └── testing.md
└── context/
    ├── architecture.md
    └── patterns.md
```

## 11.6 Configuration and Build (`hack/`)

### 11.6.1 Build Scripts (`hack/`)

```bash
hack/
├── build.sh                       # Local build script
├── test.sh                        # Test execution script
├── release.sh                     # Release preparation script
├── validate-assets.sh             # Asset validation script
├── validate-agents.py             # Professional agent validation tool
├── requirements.txt               # Python dependencies for validation
├── dev-setup.sh                   # Development environment setup
├── gen-schemas.sh                 # Generate component schemas
├── fmt-check.sh                   # Format checking script
└── ci-setup.sh                    # CI/CD setup script
```

## 11.7 Testing Strategy (`tests/`)

### 11.7.1 Test Organization

```bash
tests/
├── unit/                          # Unit tests
│   ├── cli/                       # CLI component tests
│   ├── engine/                    # Engine component tests
│   ├── assets/                    # Asset management tests
│   └── validation/                # Validation system tests
├── integration/                   # Integration tests
│   ├── install-flow/              # Installation flow tests
│   ├── validate-flow/             # Validation flow tests
│   └── ide-integration/           # IDE integration tests
├── fixtures/                      # Test data and fixtures
│   ├── valid-frameworks/          # Valid framework examples
│   ├── invalid-frameworks/        # Invalid framework examples
│   └── sample-projects/           # Sample project structures
└── e2e/                           # End-to-end tests
    ├── cli-workflows/             # Complete CLI workflows
    └── cross-platform/            # Cross-platform compatibility
```

## 11.8 Documentation (`docs/`)

The documentation structure includes comprehensive guides and practical examples:

```bash
docs/
├── architecture/                  # Architecture documentation
├── user-guide/                    # User documentation
├── developer-guide/               # Developer documentation
├── api-reference/                 # API reference (if needed)
└── examples/                      # Usage examples and samples
    ├── basic-setup/               # Basic KubeRocketAI setup examples
    ├── agent-examples/            # Example agent implementations
    ├── task-workflows/            # Example task workflows
    ├── integration-guides/        # IDE and CI/CD integration examples
    └── real-world-scenarios/      # Complete project examples
```

## 11.9 Development and Distribution Files

### 11.9.1 Core Build Files

- **go.mod/go.sum**: Go module management
- **Makefile**: Standardized build tasks
- **.goreleaser.yaml**: Cross-platform release automation
- **LICENSE**: Project licensing
- **README.md**: Project overview and quick start

### 11.9.2 Code Quality & Formatting Configuration

- **.golangci.yml**: Go linting rules and configuration
- **.markdownlint.json**: Markdown linting and style rules
- **.prettierrc**: Code formatting configuration
- **.editorconfig**: Editor-agnostic formatting settings
- **.gitignore**: Git ignore patterns for build artifacts and temporary files

### 11.9.3 GitHub Community Standards & CI/CD Configuration

The `.github/` directory contains essential community health files and automation that ensures KubeRocketAI meets GitHub's community standards checklist while providing robust CI/CD capabilities.

```bash
.github/
├── workflows/                     # GitHub Actions CI/CD pipelines
│   ├── ci.yml                     # Continuous integration & testing
│   ├── release.yml                # Release automation with GoReleaser
│   ├── security.yml               # Security scanning & vulnerability checks
│   ├── codeql-analysis.yml        # CodeQL security analysis
│   ├── dependabot-auto-merge.yml  # Automated dependency updates
│   └── docs.yml                   # Documentation deployment
├── ISSUE_TEMPLATE/                # Issue templates for structured reporting
│   ├── bug_report.yml             # Bug report form template
│   ├── feature_request.yml        # Feature request form template
│   ├── documentation.yml          # Documentation improvement template
│   ├── question.yml               # General question template
│   └── config.yml                 # Issue template chooser configuration
├── PULL_REQUEST_TEMPLATE/         # Pull request templates
│   ├── pull_request_template.md   # Default PR template
│   ├── feature.md                 # Feature development PR template
│   ├── bugfix.md                  # Bug fix PR template
│   └── docs.md                    # Documentation PR template
├── CODEOWNERS                     # Code ownership definitions
├── FUNDING.yml                    # Project funding information
├── dependabot.yml                 # Dependabot configuration
└── SECURITY.md                    # Security policy & vulnerability reporting
```

#### GitHub Actions Workflows

**Continuous Integration (`ci.yml`)**:

- Go build and test across multiple platforms
- Asset validation and schema checking
- Linting with golangci-lint and Super-Linter
- Cross-platform compatibility testing

**Release Automation (`release.yml`)**:

- Triggered on version tags
- GoReleaser for multi-platform binary builds
- Automated changelog generation
- Package manager updates (Homebrew, Chocolatey)

**Security Scanning (`security.yml`)**:

- Secret scanning with GitLeaks
- Dependency vulnerability scanning
- SAST (Static Application Security Testing)
- License compliance checking

#### Issue and PR Templates

**Bug Report Template** (`.github/ISSUE_TEMPLATE/bug_report.yml`):

**Pull Request Template** (`.github/pull_request_template.md`):

#### Community Health Files

**CODEOWNERS**:

```txt
# Global owners
* @kuberocketci/admin
```

**Dependabot Configuration** (`.github/dependabot.yml`):

#### 11.9.3 Community Standards Compliance

This structure ensures KubeRocketAI meets GitHub's community standards checklist:

✅ **README** - Project overview and getting started guide
✅ **LICENSE** - Open source license (MIT recommended)
✅ **CODE_OF_CONDUCT** - Community behavior guidelines
✅ **CONTRIBUTING** - Contribution guidelines and process
✅ **SECURITY** - Security policy and vulnerability reporting
✅ **SUPPORT** - Support resources and help channels
✅ **Issue Templates** - Structured issue reporting forms
✅ **Pull Request Templates** - Consistent PR structure
✅ **CODEOWNERS** - Code review assignments
✅ **Dependabot** - Automated dependency management

This source tree structure ensures clear separation of concerns, and provides a solid foundation for building and distributing the KubeRocketAI CLI tool with embedded LLM assets.
