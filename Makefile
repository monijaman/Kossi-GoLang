# Kossti - Go Clean Architecture Project
# Makefile for building and running applications

.PHONY: build run-app run-migrate clean test help

# Build all applications
build:
	@echo "Building applications..."
	@go build -o bin/kossti-server ./cmd/app
	@go build -o bin/kossti-migrate ./cmd/migrate
	@echo "✅ Build completed!"

# Run the main application
run-app:
	@echo "Starting Kossti API server..."
	@go run ./cmd/app/main.go

# Run database migration
run-migrate:
	@echo "Running database migration..."
	@go run ./cmd/migrate/main.go -migrate

# Run fresh migration (development only)
migrate-fresh:
	@echo "⚠️  Running fresh migration (DANGEROUS)..."
	@go run ./cmd/migrate/main.go -fresh

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@echo "✅ Clean completed!"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "✅ Dependencies installed!"

# Show help
help:
	@echo "Kossti - Available commands:"
	@echo ""
	@echo "  make build        - Build all applications"
	@echo "  make run-app      - Run the main API server"
	@echo "  make run-migrate  - Run safe database migration"
	@echo "  make migrate-fresh - Run fresh migration (DANGEROUS)"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make test         - Run tests"
	@echo "  make deps         - Install/update dependencies"
	@echo "  make help         - Show this help"
	@echo ""
	@echo "Examples:"
	@echo "  make build && ./bin/kossti-server"
	@echo "  make run-migrate"
	@echo "  make run-app"

# Default target
.DEFAULT_GOAL := help
