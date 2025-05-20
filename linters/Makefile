GOPATH=$(shell go env GOPATH)
PACKAGE := linters

# linter:
GOLINT = $(GOPATH)/bin/golangci-lint
$(GOLINT):
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.64.8

.PHONY: lint
lint: $(GOLINT)
	$(GOLINT) run

NILAWAY = $(GOPATH)/bin/nilaway
$(NILAWAY):
	go install go.uber.org/nilaway/cmd/nilaway@latest

.PHONY: nilaway
nilaway: $(NILAWAY)
	$(NILAWAY) -include-pkgs="$(PACKAGE)" -test=true -exclude-errors-in-files=mocks/ ./...
