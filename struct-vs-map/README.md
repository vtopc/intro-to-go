# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=1,4 -benchmem
goos: darwin
goarch: amd64
pkg: struct-vs-map
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMapAlloc        	   2316925	       496.1 ns/op	         582 B/op	       1 allocs/op
BenchmarkMapAlloc-4      	   2268127	       489.9 ns/op	         582 B/op	       1 allocs/op
BenchmarkStructAlloc     	1000000000	         0.2636 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructAlloc-4   	1000000000	         0.2715 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	struct-vs-map	4.513s
```
