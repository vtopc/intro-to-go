.PHONY: bench
bench:
	go test -bench=. -cpu=1,4 -benchmem

BENCHSTAT = $(GOPATH)/bin/benchstat
$(BENCHSTAT):
	go install golang.org/x/perf/cmd/benchstat@latest
