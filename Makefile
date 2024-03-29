MODULE_NAME := "github.com/psyb0t/logrus-configurator"
PKG_LIST := $(shell go list ${MODULE_NAME}/...)
MIN_TEST_COVERAGE := 80

all: dep lint test ## Run dep, lint and test

dep: ## Get project dependencies
	@echo "Getting project dependencies..."
	@go mod tidy

lint: ## Lint all Golang files
	@echo "Linting all Golang files..."
	@golangci-lint run --timeout=30m0s

test: ## Run all tests
	@echo "Running all tests..."
	@go test -race $(PKG_LIST)

test-coverage: ## Run tests with coverage check. Fails if coverage is below the threshold.
	@echo "Running tests with coverage check..."
	@trap 'rm -f coverage.txt' EXIT; \
	go test -race -coverprofile=coverage.txt $(PKG_LIST); \
	if [ $$? -ne 0 ]; then \
		echo "Test failed. Exiting."; \
		exit 1; \
	fi; \
	result=$$(go tool cover -func=coverage.txt | grep -oP 'total:\s+\(statements\)\s+\K\d+' || echo "0"); \
	if [ $$result -eq 0 ]; then \
		echo "No test coverage information available."; \
		exit 0; \
	elif [ $$result -lt $(MIN_TEST_COVERAGE) ]; then \
		echo "FAIL: Coverage $$result% is less than the minimum $(MIN_TEST_COVERAGE)%"; \
		exit 1; \
	fi

help: ## Display this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
