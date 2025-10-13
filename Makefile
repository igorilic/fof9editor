.PHONY: build run test clean

# Version information
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT_HASH := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS := -ldflags="-s -w \
	-X github.com/igorilic/fof9editor/internal/version.Version=$(VERSION) \
	-X github.com/igorilic/fof9editor/internal/version.CommitHash=$(COMMIT_HASH) \
	-X github.com/igorilic/fof9editor/internal/version.BuildDate=$(BUILD_DATE)"

# Build the application
build:
	@echo "Building fof9editor v$(VERSION)..."
	@mkdir -p bin
	@go build $(LDFLAGS) -o bin/fof9editor.exe ./cmd/fof9editor
	@echo "Build complete: bin/fof9editor.exe"

# Run the application
run: build
	@echo "Running fof9editor..."
	@./bin/fof9editor.exe

# Run tests
test:
	@echo "Running tests..."
	@go test ./... -v

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f fof9editor fof9editor.exe
	@echo "Clean complete"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies updated"
