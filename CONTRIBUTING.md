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

Use conventional commits:

- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `test:` - Adding tests
- `chore:` - Maintenance

## License

By contributing, you agree to license your work under Apache-2.0.
