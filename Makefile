# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Generate mocks
mocks:
	@echo "Generating mocks..."
	@echo "Installing mockgen..."
	@go install go.uber.org/mock/mockgen@latest
	@mockgen -source=internal/ports/logger.go -destination="mocks/mock_logger.go" -package=mocks
	@echo "Mock generation complete"

# Run the application
run:
	@export $$(cat development.env | xargs) && go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

.PHONY: all build run test clean watch docker-run docker-down itest
