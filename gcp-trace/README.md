# GCP Cloud Trace

based on https://cloud.google.com/trace/docs/setup/go

## Permissions
https://cloud.google.com/trace/docs/iam#roles

## Setup
Set `GOOGLE_APPLICATION_CREDENTIALS` and `GOOGLE_CLOUD_PROJECT` env vars.

## Benchmarks
```bash
# make bench
go test -bench=. -cpu=4 -benchmem -benchtime=100x
goos: darwin
goarch: amd64
pkg: gcp-trace
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkWithoutTrace-4   	     100	  53391358 ns/op	   26127 B/op	     127 allocs/op
--- BENCH: BenchmarkWithoutTrace-4
BenchmarkWithTrace-4      	     100	  53548751 ns/op	   82529 B/op	    1118 allocs/op
--- BENCH: BenchmarkWithTrace-4
PASS
ok  	gcp-trace	11.617s
```
