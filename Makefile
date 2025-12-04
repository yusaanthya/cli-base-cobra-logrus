.PHONY: help build run test clean docker-build docker-run lint fmt vet

# Variables
BINARY_NAME=cli-base
DOCKER_IMAGE=cli-base-cobra-logrus
DOCKER_TAG=latest
GO_FILES=$(shell find . -name '*.go' -type f)

# Version info
VERSION=$(shell cat VERSION 2>/dev/null || echo "dev")
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags "-X github.com/yusaanthya/cli-base-cobra-logrus/pkg/config.Version=$(VERSION) \
                  -X github.com/yusaanthya/cli-base-cobra-logrus/pkg/config.GitCommit=$(GIT_COMMIT) \
                  -X github.com/yusaanthya/cli-base-cobra-logrus/pkg/config.BuildDate=$(BUILD_DATE)"

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the binary
	@echo "Building $(BINARY_NAME) version $(VERSION)..."
	@mkdir -p bin
	@go build $(LDFLAGS) -o bin/$(BINARY_NAME) cmd/main.go
	@echo "Build complete: bin/$(BINARY_NAME)"

run: build ## Build and run the application
	@echo "Running $(BINARY_NAME)..."
	@./bin/$(BINARY_NAME)

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -rf dist/
	@rm -f log_output.txt
	@echo "Clean complete"

docker-build: ## Build Docker image
	@echo "Building Docker image $(DOCKER_IMAGE):$(DOCKER_TAG)..."
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "Docker image built: $(DOCKER_IMAGE):$(DOCKER_TAG)"

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	@docker run --rm $(DOCKER_IMAGE):$(DOCKER_TAG)

lint: ## Run linter (requires golangci-lint)
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install it from https://golangci-lint.run/usage/install/"; \
		exit 1; \
	fi

fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

tidy: ## Tidy go modules
	@echo "Tidying go modules..."
	@go mod tidy

all: clean fmt vet test build ## Run all checks and build
