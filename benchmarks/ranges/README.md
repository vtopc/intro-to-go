# Benchmarks

Range benchmarks.

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: arm64
pkg: ranges
Benchmark_sumElements-4   	   74632	     15916 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumByIndex-4    	 4594131	       257.9 ns/op	       0 B/op	       0 allocs/op
```
