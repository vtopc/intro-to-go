.PHONY: test
test:
	go test ./... -count=100000

.PHONY: run
run:
	GOMAXPROCS=16 go run race.go

.PHONY: run-race
run-race:
	go run -race race.go
