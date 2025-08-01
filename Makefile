# KubeRocketAI CLI Makefile

.DEFAULT_GOAL := help

# Variables
APP_NAME ?= krci-ai
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "v0.0.0-dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILT_BY ?= $(shell whoami)

# Go build variables
GO_VERSION ?= 1.24.4
GO_BUILD_FLAGS ?= -trimpath
GO_LDFLAGS = -s -w \
	-X main.version=$(VERSION) \
	-X main.commit=$(COMMIT) \
	-X main.date=$(DATE) \
	-X main.builtBy=$(BUILT_BY)

# Directories
BUILD_DIR := dist
BIN_DIR := bin
DIST_DIR := dist

# Tools - pinned to latest stable versions as of 2025-07-06
GOLANGCI_LINT_VERSION ?= v1.64.8
GORELEASER_VERSION ?= v2.10.2
STATICCHECK_VERSION ?= 2025.1.1
GIT_CHGLOG_VERSION ?= v0.15.4

# Cross-platform builds
PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

.PHONY: help
help: ## Display available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: all
all: fmt lint test build ## Run all quality checks and build

.PHONY: build
build: ## Build the CLI binary
	@echo "Building $(APP_NAME) $(VERSION)..."
	@mkdir -p $(DIST_DIR)
	go build $(GO_BUILD_FLAGS) -ldflags "$(GO_LDFLAGS)" -o $(DIST_DIR)/$(APP_NAME) ./cmd/krci-ai

.PHONY: build-all
build-all: ## Build for all platforms
	@echo "Building for all platforms..."
	@for platform in $(PLATFORMS); do \
		echo "Building for $$platform..."; \
		os=$$(echo $$platform | cut -d'/' -f1); \
		arch=$$(echo $$platform | cut -d'/' -f2); \
		output="$(BUILD_DIR)/$(APP_NAME)-$$os-$$arch"; \
		if [ "$$os" = "windows" ]; then output="$$output.exe"; fi; \
		GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build $(GO_BUILD_FLAGS) -ldflags "$(GO_LDFLAGS)" -o $$output ./cmd/krci-ai; \
	done

.PHONY: test
test: ## Run unit tests
	@echo "Running tests..."
	go test -race -coverprofile=coverage.out ./...

.PHONY: test-verbose
test-verbose: ## Run unit tests with verbose output
	@echo "Running tests (verbose)..."
	go test -v -race -coverprofile=coverage.out ./...

.PHONY: test-coverage
test-coverage: test ## Run tests and show coverage report
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: fmt
fmt: $(BIN_DIR)/golangci-lint ## Format Go code
	@echo "Formatting code..."
	go fmt ./...
	@echo "Fixing imports and formatting..."
	$(BIN_DIR)/golangci-lint run --fix

.PHONY: lint
lint: $(BIN_DIR)/golangci-lint ## Run linter
	@echo "Running linter..."
	$(BIN_DIR)/golangci-lint run

.PHONY: vet
vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...

.PHONY: staticcheck
staticcheck: $(BIN_DIR)/staticcheck ## Run staticcheck
	@echo "Running staticcheck..."
	$(BIN_DIR)/staticcheck ./...

.PHONY: deps
deps: ## Download Go modules
	@echo "Downloading dependencies..."
	go mod download
	go mod verify

.PHONY: tidy
tidy: ## Tidy Go modules
	@echo "Tidying modules..."
	go mod tidy

.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -rf $(DIST_DIR) $(BIN_DIR) coverage.out coverage.html

.PHONY: install
install: build ## Install the binary to GOPATH/bin
	@echo "Installing $(APP_NAME) to GOPATH/bin..."
	cp $(DIST_DIR)/$(APP_NAME) $(shell go env GOPATH)/bin/

.PHONY: release-snapshot
release-snapshot: ## Build snapshot release
	@echo "Building snapshot release..."
	goreleaser release --snapshot --clean

.PHONY: release-test
release-test: ## Test release configuration
	@echo "Testing release configuration..."
	goreleaser check

.PHONY: changelog
changelog: $(BIN_DIR)/git-chglog ## Generate changelog
	@echo "Generating changelog..."
	$(BIN_DIR)/git-chglog -o CHANGELOG.md

.PHONY: changelog-validate
changelog-validate: $(BIN_DIR)/git-chglog ## Validate changelog is up-to-date
	@echo "Validating changelog is up-to-date..."
	@$(MAKE) changelog
	@if git diff --exit-code CHANGELOG.md >/dev/null 2>&1; then \
		echo "✅ Changelog is up-to-date"; \
	else \
		echo "❌ Changelog needs to be regenerated. Run 'make changelog' and commit the changes."; \
		git diff CHANGELOG.md; \
		exit 1; \
	fi

.PHONY: update-readme-badge
update-readme-badge: ## Update framework badge in README.md
	@echo "Updating framework badge in README.md..."
	@AGENTS=$$(find assets/framework/core/agents -name "*.yaml" -o -name "*.yml" -o -name "*.md" 2>/dev/null | sort | uniq | wc -l | tr -d ' '); \
	TASKS=$$(find assets/framework/core/tasks -name "*.md" 2>/dev/null | sort | uniq | wc -l | tr -d ' '); \
	TEMPLATES=$$(find assets/framework/core/templates -name "*.md" 2>/dev/null | sort | uniq | wc -l | tr -d ' '); \
	DATA=$$(find assets/framework/core/data -name "*.md" -o -name "*.json" -o -name "*.yaml" -o -name "*.yml" 2>/dev/null | sort | uniq | wc -l | tr -d ' '); \
	NEW_BADGE="[![Framework Overview](https://img.shields.io/badge/Framework-$${AGENTS}%20agents%20%7C%20$${TASKS}%20tasks%20%7C%20$${TEMPLATES}%20templates%20%7C%20$${DATA}%20data%20files-purple?style=flat-square&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTEyIDJMMTMuMDkgOC4yNkwyMCA5TDEzLjA5IDE1Ljc0TDEyIDIyTDEwLjkxIDE1Ljc0TDQgOUwxMC45MSA4LjI2TDEyIDJaIiBmaWxsPSJ3aGl0ZSIvPgo8L3N2Zz4K)](.krci-ai)"; \
	if [ -f README.md ]; then \
		awk -v new_badge="$$NEW_BADGE" ' \
			/^[![]Framework Overview/ { print new_badge; next } \
			{ print } \
		' README.md > README.md.tmp && mv README.md.tmp README.md; \
		echo "✅ Updated framework badge in README.md"; \
		echo "New counts: $$AGENTS agents, $$TASKS tasks, $$TEMPLATES templates, $$DATA data files"; \
	else \
		echo "❌ README.md not found"; \
		exit 1; \
	fi

.PHONY: version
version: ## Show version information
	@echo "App Name: $(APP_NAME)"
	@echo "Version:  $(VERSION)"
	@echo "Commit:   $(COMMIT)"
	@echo "Date:     $(DATE)"
	@echo "Built By: $(BUILT_BY)"

.PHONY: run
run: build ## Build and run the CLI
	@echo "Running $(APP_NAME)..."
	./$(DIST_DIR)/$(APP_NAME)

.PHONY: ci
ci: deps fmt vet lint staticcheck test build ## Run CI pipeline locally

# Tool installation helpers
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(BIN_DIR)/golangci-lint: $(BIN_DIR)
	@echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	GOBIN=$(PWD)/$(BIN_DIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

$(BIN_DIR)/goreleaser: $(BIN_DIR)
	@echo "Installing goreleaser $(GORELEASER_VERSION)..."
	GOBIN=$(PWD)/$(BIN_DIR) go install github.com/goreleaser/goreleaser/v2@$(GORELEASER_VERSION)

$(BIN_DIR)/staticcheck: $(BIN_DIR)
	@echo "Installing staticcheck $(STATICCHECK_VERSION)..."
	GOBIN=$(PWD)/$(BIN_DIR) go install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)

$(BIN_DIR)/git-chglog: $(BIN_DIR)
	@echo "Installing git-chglog $(GIT_CHGLOG_VERSION)..."
	GOBIN=$(PWD)/$(BIN_DIR) go install github.com/git-chglog/git-chglog/cmd/git-chglog@$(GIT_CHGLOG_VERSION)

.PHONY: tools
tools: $(BIN_DIR)/golangci-lint $(BIN_DIR)/goreleaser $(BIN_DIR)/staticcheck $(BIN_DIR)/git-chglog ## Install all development tools

.PHONY: check-tools
check-tools: ## Check if required tools are installed
	@echo "Checking required tools..."
	@which go >/dev/null || (echo "Go is not installed" && exit 1)
	@echo "Go version: $$(go version)"
	@which git >/dev/null || (echo "Git is not installed" && exit 1)
	@echo "Git version: $$(git --version)"
	@echo "All required tools are available"

# Include additional makefiles if they exist
-include Makefile.local
