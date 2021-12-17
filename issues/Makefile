# linter:
GOLINT = $(GOPATH)/bin/golangci-lint
$(GOLINT):
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.43.0

.PHONY: lint
lint: $(GOLINT)
	$(GOLINT) run

.PHONE: govet
govet:
	go vet ./...
