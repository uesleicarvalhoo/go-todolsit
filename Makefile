-include: .env

GO_ENTRYPOINT=cmd/api.go
COVERAGE_OUTPUT=coverage.output
COVERAGE_HTML=coverage.html

## @ Help
.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "\033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)


## @ Application
.PHONY: run setup setdown
run:  ## Start application
	@go run $(GO_ENTRYPOINT)

setup: ## Start app dependencies
	@docker-compose up -d

setdown:  ## Stop application dependencies
	@docker-compose down


## @ Create docs
.PHONY: doc
doc:
	swag init -g $(GO_ENTRYPOINT)
	swag fmt -g $(GO_ENTRYPOINT)


## @ Tests
.PHONY: test coverage
test:  ## Run tests of project
	@go test ./... -cover -coverprofile=$(COVERAGE_OUTPUT)

coverage: test ## Run tests, make report and open into browser
	@go tool cover -html=$(COVERAGE_OUTPUT) -o $(COVERAGE_HTML)
	@wslview ./$(COVERAGE_HTML) || xdg-open ./$(COVERAGE_HTML) || powershell.exe Invoke-Expression ./$(COVERAGE_HTML)

## @ Clean
.PHONY: clean clean_coverage_cache
clean_coverage_cache: ## Remove coverage cache files
	@rm -rf $(COVERAGE_OUTPUT)
	@rm -rf $(COVERAGE_HTML)

clean: clean_coverage_cache ## Remove Cache files
