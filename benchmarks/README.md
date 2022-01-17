# Benchmarks

Just a random benchmarks.

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIndexSlice-4    	  887233	      1345 ns/op	    8192 B/op	       1 allocs/op
BenchmarkAppendSlice-4   	  629828	      1591 ns/op	    8192 B/op	       1 allocs/op
```
