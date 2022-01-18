# Benchmarks

Just a random benchmarks.

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceIndex-4            	    826884	      1291 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSliceAppend-4           	    771477	      1455 ns/op	    8192 B/op	       1 allocs/op
BenchmarkAllocStructLiteral-4    	1000000000	         0.2511 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllocStructByFields-4   	1000000000	         0.2489 ns/op	       0 B/op	       0 allocs/op
```
