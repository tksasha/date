.PHONY: all
all: vet fix fmt lint test

.PHONY: vet
vet:
	@echo "go vet"
	@go vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@go fix ./...

.PHONY: fmt
fmt:
	@echo "go fmt"
	@gofumpt -l -w .

.PHONY: lint
lint:
	@echo "go lint"
	@golangci-lint run

.PHONY: test
test:
	@echo "go test"
	@go test ./...

.PHONY: prepare
prepare:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go get mvdan.cc/gofumpt@latest
	go install mvdan.cc/gofumpt@latest
	go mod tidy
