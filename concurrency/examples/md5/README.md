# wg

This is an example of wait groups and done channel.

## Benchmarks

```shell
# make bench
go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem
2023/02/27 14:09:16 generated requests
goos: darwin
goarch: amd64
pkg: concurrency
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkDoAsync       	      40	 129234027 ns/op	    7479 B/op	     267 allocs/op
BenchmarkDoAsync-2     	      16	  65426950 ns/op	    7614 B/op	     267 allocs/op
BenchmarkDoAsync-4     	      33	  34988374 ns/op	    7825 B/op	     268 allocs/op
BenchmarkDoAsync-8     	      40	  28144555 ns/op	    8011 B/op	     268 allocs/op
BenchmarkDoAsync-12    	      36	  29294738 ns/op	    8108 B/op	     268 allocs/op
BenchmarkDoAsync-16    	      38	  30252928 ns/op	    7737 B/op	     267 allocs/op
BenchmarkDoAsync-24    	      42	  27465128 ns/op	    7642 B/op	     267 allocs/op
BenchmarkDoAsync-32    	      43	  26973481 ns/op	    7915 B/op	     268 allocs/op
BenchmarkDoSync        	       8	 131417420 ns/op	    5737 B/op	     252 allocs/op
BenchmarkDoSync-2      	       8	 127920070 ns/op	    5787 B/op	     252 allocs/op
BenchmarkDoSync-4      	       8	 138157063 ns/op	    5785 B/op	     252 allocs/op
BenchmarkDoSync-8      	       8	 132623499 ns/op	    5883 B/op	     252 allocs/op
BenchmarkDoSync-12     	       8	 130274766 ns/op	    5947 B/op	     252 allocs/op
BenchmarkDoSync-16     	       8	 131227420 ns/op	    6011 B/op	     252 allocs/op
BenchmarkDoSync-24     	       8	 130575618 ns/op	    6139 B/op	     252 allocs/op
BenchmarkDoSync-32     	       8	 131656592 ns/op	    6267 B/op	     252 allocs/op
PASS
ok  	concurrency	24.768s
```