.DEFAULT_GOAL := help
.PHONY : build

run: ## Run code
	@go run ./cmd/main.go

build: ## Build binary
	@mkdir -p bin
	@go build -o bin/${APP} ./cmd/

test: ## Run tests
	@go test -v -race ./... -coverprofile=coverage.out

clean: ## Cleaning binary
	@rm -f bin/${APP}

# Meant to display a pretty list of commands of this makefile, works on Linux :). Just run `make` to understand what it does
help: ## Show commands availables
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


