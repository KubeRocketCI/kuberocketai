# Contributing to KubeRocketAI

## Development Setup

### Prerequisites

- Go 1.24.4+
- Git
- Make

### Getting Started

```bash
git clone https://github.com/KubeRocketCI/kuberocketai.git
cd kuberocketai
make tools
make test
make build
```

## Making Changes

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes
4. Run tests: `make ci`
5. Commit using conventional commits: `feat(cli): add new command`
6. Push and create a pull request

## Contribution Areas

- **CLI Development**: Add commands, improve validation, enhance UX
- **Agent Definitions**: Create agent templates for SDLC roles
- **Documentation**: Improve guides, add examples
- **CI/CD Integration**: Add platform support, improve pipeline awareness

## Code Style

- Follow Go conventions
- Use `gofmt` for formatting
- Add tests for new functionality
- Ensure all CI checks pass

## Commit Messages

We use [Conventional Commits](https://www.conventionalcommits.org/) for automated changelog generation and semantic versioning.

### Format

```bash
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `test:` - Adding tests
- `chore:` - Maintenance
- `refactor:` - Code refactoring
- `perf:` - Performance improvements
- `ci:` - CI/CD changes

### Examples

```bash
feat: add changelog generation command
fix(cli): improve validate command warning display
docs: update installation instructions
chore: update dependencies to latest versions
```

### Breaking Changes

For breaking changes, add `BREAKING CHANGE:` in the footer:

```bash
feat: change API endpoint structure

BREAKING CHANGE: The /api/v1/users endpoint has been moved to /api/v2/users
```

### Complete Guide

See [docs/CONVENTIONAL_COMMITS.md](./docs/CONVENTIONAL_COMMITS.md) for the complete conventional commits guide with detailed examples and best practices.

## Changelog

The changelog is automatically generated from conventional commit messages using git-chglog. **Do not manually edit CHANGELOG.md**.

### Changelog Commands

```bash
# Generate changelog (requires git-chglog)
make changelog
```

### Release Process

1. Commits using conventional format are merged to main
2. During release, `make changelog` generates CHANGELOG.md from commits
3. Changelog is embedded in CLI binary for offline access
4. Release is created with generated changelog content

## Quality Assurance

### SonarCloud Integration

The project uses [SonarCloud](https://sonarcloud.io/project/overview?id=kuberocketci_kuberocketai) for continuous code quality analysis:

- **Quality Gate**: Automatic analysis of code quality, security vulnerabilities, and technical debt
- **Coverage Reports**: Go test coverage is automatically uploaded and tracked
- **Configuration**: See `sonar-project.properties` for SonarCloud configuration
- **Badges**: Quality gate status and coverage are displayed in README.md

The SonarCloud analysis runs automatically on every push to main/develop branches and pull requests.

## License

By contributing, you agree to license your work under Apache-2.0.
