.PHONY: generate
generate: ## generate mock
	rm -rf mock/*
	go generate ./...

.PHONY: test
test: ## run test
	go test -v -cover ./...
