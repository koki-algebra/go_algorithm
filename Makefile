.PHONY: help
.DEFAULT_GOAL := help

fmt: ## Format code
	@go fmt ./...

lint: ## Lint code
	@golangci-lint run

test: ## Run Test
	@go test -race -shuffle=on ./pkg/libs

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
