.PHONY: run
run:
	go run ./*.go

.PHONY: bench
bench:
	go test -bench=. -cpu=2,4,8 -benchmem
