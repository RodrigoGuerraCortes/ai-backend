.PHONY: tests run tidy

# ==== Environment Variables ====
export GEMINI_API_KEY := $(shell grep GEMINI_API_KEY .env | cut -d '=' -f2)

# ==== Commands ====

# Run tests
tests:
	@echo "🔧 Running unit tests..."
	@if [ -z "$(GEMINI_API_KEY)" ]; then \
		echo "❌ Error: GEMINI_API_KEY is missing in .env file"; \
		exit 1; \
	fi
	GEMINI_API_KEY=$(GEMINI_API_KEY) go test ./... -v

# Run app
run:
	@echo "🚀 Starting API..."
	GEMINI_API_KEY=$(GEMINI_API_KEY) go run cmd/api/main.go

# Tidy dependencies
tidy:
	go mod tidy