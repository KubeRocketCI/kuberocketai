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

# Tools - pinned to latest stable versions as of 2025-07-06
GOLANGCI_LINT_VERSION ?= v1.64.8
GORELEASER_VERSION ?= v2.10.2
STATICCHECK_VERSION ?= 2025.1.1

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
	go build $(GO_BUILD_FLAGS) -ldflags "$(GO_LDFLAGS)" -o $(BIN_DIR)/$(APP_NAME) ./cmd/krci-ai

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
	rm -rf $(BUILD_DIR) $(BIN_DIR) coverage.out coverage.html

.PHONY: install
install: build ## Install the binary to GOPATH/bin
	@echo "Installing $(APP_NAME) to GOPATH/bin..."
	cp $(BIN_DIR)/$(APP_NAME) $(shell go env GOPATH)/bin/

.PHONY: release-snapshot
release-snapshot: ## Build snapshot release
	@echo "Building snapshot release..."
	goreleaser release --snapshot --clean

.PHONY: release-test
release-test: ## Test release configuration
	@echo "Testing release configuration..."
	goreleaser check

.PHONY: changelog
changelog: ## Generate changelog
	@echo "Generating changelog..."
	@if command -v git-chglog >/dev/null 2>&1; then \
		git-chglog -o CHANGELOG.md; \
	else \
		echo "git-chglog not found. Install it with: go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest"; \
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
	./$(BIN_DIR)/$(APP_NAME)

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

.PHONY: tools
tools: $(BIN_DIR)/golangci-lint $(BIN_DIR)/goreleaser $(BIN_DIR)/staticcheck ## Install all development tools

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
