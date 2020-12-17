BIN_PATH = build/bin

.PHONY: all
all: run

.PHONY: run
run:
	go run main.go

.PHONY: bench
bench:
	go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem

.PHONY: test
test:
	go test -count=1 -cover ./...

.PHONY: lint
lint:
	golangci-lint run ./...
