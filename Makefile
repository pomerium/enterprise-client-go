.PHONY: all
all: generate build lint test

.PHONY: build
build:
	@echo "==> $@"
	go build ./...

.PHONY: cover
cover: ## Runs go test with coverage
	@echo "==> $@"
	@go test -race -coverprofile=coverage.txt ./...

.PHONY: generate
generate:
	@echo "==> $@"
	./scripts/generate

.PHONY: lint
lint:
	@echo "==> $@"
	golangci-lint run --fix

.PHONY: test
test:
	@echo "==> $@"
	go test -v ./...

.PHONY: update-pomerium
update-pomerium:
	@echo "==> $@"
	git submodule update --remote deps/github.com/pomerium
