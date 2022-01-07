# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=1,4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIfaceMapAlloc     	   2261497	       494.7 ns/op	         582 B/op	       1 allocs/op
BenchmarkIfaceMapAlloc-4   	   1787796	       589.7 ns/op	         582 B/op	       1 allocs/op
BenchmarkIntMapAlloc       	   3389368	       335.9 ns/op	         420 B/op	       1 allocs/op
BenchmarkIntMapAlloc-4     	   3536082	       326.1 ns/op	         420 B/op	       1 allocs/op
BenchmarkStructAlloc       	1000000000	         0.2663 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructAlloc-4     	1000000000	         0.2690 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapSearch      	 142442265	         9.476 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapSearch-4    	 100000000	        10.35 ns/op	           0 B/op	       0 allocs/op
BenchmarkStructSearch      	1000000000	         0.2625 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch-4    	1000000000	         0.2521 ns/op	       0 B/op	       0 allocs/op
```

## TODO:
- Add map vs switch bench
