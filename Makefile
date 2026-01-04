# Makefile for Notes API Go project
# Provides convenient commands for development

.PHONY: help run build test clean migrate swagger

# Default target
help:
	@echo "Available commands:"
	@echo "  make run      - Run the application"
	@echo "  make build    - Build the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make migrate  - Run database migrations"
	@echo "  make swagger  - Generate Swagger documentation"

# Run the application
run:
	@echo "🚀 Starting server..."
	@go run ./cmd/api

# Build the application
build:
	@echo "🔨 Building application..."
	@go build -o bin/api ./cmd/api
	@echo "✅ Build complete: bin/api"

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	@rm -rf bin/
	@go clean
	@echo "✅ Clean complete"

# Run database migrations (same as running the app once)
migrate:
	@echo "📦 Running migrations..."
	@go run ./cmd/api

# Generate Swagger documentation
swagger:
	@echo "📚 Generating Swagger documentation..."
	@if command -v swag >/dev/null 2>&1; then \
		swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal; \
	elif [ -f $$(go env GOPATH)/bin/swag ]; then \
		$$(go env GOPATH)/bin/swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal; \
	else \
		echo "⚠️  swag tool not found, installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		$$(go env GOPATH)/bin/swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal; \
	fi
	@echo "✅ Swagger documentation generated in docs/ folder"
