.DEFAULT_GOAL := help

.PHONY: install
install: install-goimports install-mockgen ## install dev tools

.PHONY: install-goimports
install-goimports: ## install goimports
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: install-mockgen
install-mockgen: ## install mockgen
	go install github.com/golang/mock/mockgen@latest

.PHONY: generate
generate: ## generate mock
	rm -rf mock/*
	go generate ./...

.PHONY: test
test: ## run test
	go test -v -cover `go list ./... | grep -v mock_*` -coverprofile=coverage.out
	go tool cover -html=coverage.out -o ./testdata/coverage.html

.PHONY: help
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
