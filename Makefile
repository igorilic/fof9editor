.PHONY: build run test clean

# Build the application
build:
	@echo "Building fof9editor..."
	@mkdir -p bin
	@go build -o bin/fof9editor.exe ./cmd/fof9editor
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
