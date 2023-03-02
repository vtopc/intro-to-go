# errgroup

This is an example of errgroup.

## Benchmarks

```shell
go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem
2023/02/27 14:07:14 generated requests
goos: darwin
goarch: amd64
pkg: concurrency
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkDoAsync       	      40	 128901325 ns/op	   10996 B/op	     363 allocs/op
BenchmarkDoAsync-2     	      16	  64511305 ns/op	   11198 B/op	     363 allocs/op
BenchmarkDoAsync-4     	      33	  34434699 ns/op	   11468 B/op	     364 allocs/op
BenchmarkDoAsync-8     	      40	  26960529 ns/op	   11349 B/op	     364 allocs/op
BenchmarkDoAsync-12    	      43	  26935106 ns/op	   11326 B/op	     363 allocs/op
BenchmarkDoAsync-16    	      43	  27025449 ns/op	   11119 B/op	     363 allocs/op
BenchmarkDoAsync-24    	      43	  26867428 ns/op	   11338 B/op	     364 allocs/op
BenchmarkDoAsync-32    	      40	  26872675 ns/op	   11196 B/op	     363 allocs/op
BenchmarkDoSync        	       8	 126294462 ns/op	    5737 B/op	     252 allocs/op
BenchmarkDoSync-2      	       8	 130346871 ns/op	    5753 B/op	     252 allocs/op
BenchmarkDoSync-4      	       8	 126239970 ns/op	    5819 B/op	     252 allocs/op
BenchmarkDoSync-8      	       8	 133230816 ns/op	    5883 B/op	     252 allocs/op
BenchmarkDoSync-12     	       8	 128322215 ns/op	    5947 B/op	     252 allocs/op
BenchmarkDoSync-16     	       8	 128095719 ns/op	    5977 B/op	     252 allocs/op
BenchmarkDoSync-24     	       8	 126188920 ns/op	    6139 B/op	     252 allocs/op
BenchmarkDoSync-32     	       8	 128930582 ns/op	    6267 B/op	     252 allocs/op
PASS
ok  	concurrency	24.609s
```