
.PHONY: install
install: install-goimports install-mockgen

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
	go test -v -cover `go list ./... | grep -v mock_*`
