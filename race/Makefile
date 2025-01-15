.PHONY: test
test:
	go test ./... -count=1000000

.PHONY: test-race
test-race:
	go test -race ./...

.PHONY: run
run:
	GOMAXPROCS=16 go run race.go

.PHONY: run-race
run-race:
	go run -race race_map.go
