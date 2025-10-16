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
	@VERSION=$$(go run github.com/mikefarah/yq/v4@v4.34.1 '.jobs.lint.steps[] | select(.uses == "golangci/golangci-lint-action*") | .with.version' .github/workflows/reusable-build.yaml) && \
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@$$VERSION run ./...

.PHONY: test
test:
	@echo "==> $@"
	go test -v ./...

.PHONY: update-pomerium
update-pomerium:
	@echo "==> $@"
	git submodule update --remote deps/github.com/pomerium
