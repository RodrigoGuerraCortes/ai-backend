.PHONY: tests run tidy integration-test

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

# Run tests
tests-integration:
	@echo "🔧 Running unit tests..."
	@if [ -z "$(GEMINI_API_KEY)" ]; then \
		echo "❌ Error: GEMINI_API_KEY is missing in .env file"; \
		exit 1; \
	fi
	GEMINI_API_KEY=$(GEMINI_API_KEY) go test ./tests/integration/. -v -tags=integration


# Run app
run:
	@echo "🚀 Starting API..."
	GEMINI_API_KEY=$(GEMINI_API_KEY) go run cmd/api/main.go

# Tidy dependencies
tidy:
	go mod tidy


# ===== Coverage =====
COVER_OUT := coverage.out
COVER_HTML := coverage.html

# Detect OS opener
UNAME_S := $(shell uname)
ifeq ($(UNAME_S),Darwin)
  OPEN := open
else
  OPEN := xdg-open
endif

## Unit test coverage (excluye tests con tag 'integration')
coverage:
	@echo "🧪 Running unit tests with coverage..."
	go test ./... -tags='' -covermode=atomic -coverpkg=./... -coverprofile=$(COVER_OUT)

## Genera HTML y lo abre
coverage-html: coverage
	@echo "🧮 Generating HTML report..."
	go tool cover -html=$(COVER_OUT) -o $(COVER_HTML)
	@echo "📄 Report: $(COVER_HTML)"
	-@$(OPEN) $(COVER_HTML) >/dev/null 2>&1 || true

## Solo abre el HTML si ya existe
coverage-open:
	@echo "📄 Opening $(COVER_HTML)..."
	-@$(OPEN) $(COVER_HTML) >/dev/null 2>&1 || (echo "Run 'make coverage-html' first"; exit 1)

## Limpia artefactos de coverage
coverage-clean:
	@rm -f $(COVER_OUT) $(COVER_HTML)
	@echo "🧹 Coverage artifacts removed."

## (Opcional) Incluir también tests con tag 'integration' (requiere GEMINI_API_KEY)
coverage-all:
	@echo "🧪 Running ALL tests (unit + integration) with coverage..."
	GEMINI_API_KEY=$(GEMINI_API_KEY) go test ./... -tags=integration -covermode=atomic -coverpkg=./... -coverprofile=$(COVER_OUT)
	@echo "🧮 Generating HTML report..."
	go tool cover -html=$(COVER_OUT) -o $(COVER_HTML)
	-@$(OPEN) $(COVER_HTML) >/dev/null 2>&1 || true

# ===== Coverage threshold =====
COVER_MIN := 60.0

coverage-check: coverage
	@total=$$(go tool cover -func=$(COVER_OUT) | grep total: | awk '{print $$3}' | sed 's/%//'); \
	echo "Total coverage: $$total%"; \
	awk 'BEGIN{if ('"$$total"' < '"$(COVER_MIN)"') {print "❌ Coverage below $(COVER_MIN)%"; exit 1} else {print "✅ Coverage OK"} }'


# Test de integrations
integration-test:
	@echo "🔗 Running integration tests..."
	@if [ -z "$(GEMINI_API_KEY)" ]; then \
		echo "❌ Error: GEMINI_API_KEY is missing in .env file"; \
		exit 1; \
	fi
	GEMINI_API_KEY=$(GEMINI_API_KEY) go test ./tests/integration -tags=integration -v