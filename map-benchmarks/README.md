# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=1,4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIfaceMapAlloc     	   2492761	       480.2 ns/op	         582 B/op	       1 allocs/op
BenchmarkIfaceMapAlloc-4   	   2543923	       460.1 ns/op	         582 B/op	       1 allocs/op
BenchmarkIntMapAlloc       	   3820152	       303.2 ns/op	         420 B/op	       1 allocs/op
BenchmarkIntMapAlloc-4     	   3842608	       306.1 ns/op	         420 B/op	       1 allocs/op
BenchmarkStructAlloc       	1000000000	         0.2535 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructAlloc-4     	1000000000	         0.2535 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapSearch      	 126815578	         8.457 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapSearch-4    	 134849901	         8.874 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchSearch      	 208705524	         5.706 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchSearch-4    	 210539812	         5.712 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch      	1000000000	         0.2531 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch-4    	1000000000	         0.2532 ns/op	       0 B/op	       0 allocs/op
```
