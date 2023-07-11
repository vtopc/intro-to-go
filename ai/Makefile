.PHONY: test
test:
	go test `go list ./... | grep -v '/mocks'` -cover -count=1

.PHONY: run
run:
	go run main.go
