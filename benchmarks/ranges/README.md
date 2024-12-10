# Benchmarks

Range benchmarks.

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: arm64
pkg: ranges
Benchmark_sumValueElements-4     	   64627	     15898 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumValuesByIndex-4     	 4653972	       257.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumPointerElements-4   	 2309469	       517.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumPointersByIndex-4   	 2313625	       518.6 ns/op	       0 B/op	       0 allocs/op
```

### More optimizations
https://sourcegraph.com/blog/slow-to-simd