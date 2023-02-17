# wg

This is an example of wait groups and done channel.

## Benchmarks

```shell
# make bench
go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem
2023/02/17 17:55:51 generated requests
goos: darwin
goarch: amd64
pkg: concurrency
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkDoAsync       	      37	 134790350 ns/op	    7503 B/op	     267 allocs/op
BenchmarkDoAsync-2     	      16	  64639567 ns/op	    7576 B/op	     267 allocs/op
BenchmarkDoAsync-4     	      32	  34670997 ns/op	    7810 B/op	     267 allocs/op
BenchmarkDoAsync-8     	      40	  27121990 ns/op	    8025 B/op	     268 allocs/op
BenchmarkDoAsync-12    	      42	  27080714 ns/op	    7981 B/op	     268 allocs/op
BenchmarkDoAsync-16    	      42	  26861722 ns/op	    7720 B/op	     267 allocs/op
BenchmarkDoAsync-24    	      42	  26913291 ns/op	    7880 B/op	     268 allocs/op
BenchmarkDoAsync-32    	      43	  27082743 ns/op	    7789 B/op	     268 allocs/op
BenchmarkDoSync        	       8	 127653992 ns/op	    5737 B/op	     252 allocs/op
BenchmarkDoSync-2      	       8	 126686749 ns/op	    5753 B/op	     252 allocs/op
BenchmarkDoSync-4      	       8	 126389762 ns/op	    5819 B/op	     252 allocs/op
BenchmarkDoSync-8      	       8	 127686917 ns/op	    5883 B/op	     252 allocs/op
BenchmarkDoSync-12     	       8	 127301666 ns/op	    5947 B/op	     252 allocs/op
BenchmarkDoSync-16     	       8	 126206153 ns/op	    6011 B/op	     252 allocs/op
BenchmarkDoSync-24     	       8	 125129413 ns/op	    6139 B/op	     252 allocs/op
BenchmarkDoSync-32     	       8	 131223197 ns/op	    6267 B/op	     252 allocs/op
PASS
ok  	concurrency	23.333s
```