GO=/opt/homebrew/bin/go
LINTER=github.com/golangci/golangci-lint/cmd/golangci-lint@latest
FORMATTER=mvdan.cc/gofumpt@latest

.PHONY: default
default: vet fix fmt lint test

.PHONY: vet
vet:
	@echo "go vet"
	@$(GO) vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@$(GO) fix ./...

.PHONY: fmt
fmt:
	@echo "go fmt"
	@$(GO) run $(FORMATTER) -l -w .

.PHONY: lint
lint:
	@echo "go lint"
	@$(GO) run $(LINTER) run

.PHONY: test
test:
	@echo "go test"
	@$(GO) test ./...
