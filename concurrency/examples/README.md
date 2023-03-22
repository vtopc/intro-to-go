# concurrency examples

## Benchmarks

```shell
go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem ./...
?   	concurrency	[no test files]
goos: darwin
goarch: amd64
pkg: concurrency/examples/1sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSync       	       8	 137688731 ns/op	    5737 B/op	     252 allocs/op
BenchmarkSync-2     	       8	 138795384 ns/op	    5787 B/op	     252 allocs/op
BenchmarkSync-4     	       8	 139212277 ns/op	    5819 B/op	     252 allocs/op
BenchmarkSync-8     	       8	 137957018 ns/op	    5883 B/op	     252 allocs/op
BenchmarkSync-12    	       8	 133918635 ns/op	    5947 B/op	     252 allocs/op
BenchmarkSync-16    	       7	 145093983 ns/op	    6052 B/op	     252 allocs/op
BenchmarkSync-24    	       7	 146673252 ns/op	    6198 B/op	     252 allocs/op
BenchmarkSync-32    	       8	 136650012 ns/op	    6267 B/op	     252 allocs/op
PASS
ok  	concurrency/examples/1sync	12.082s
goos: darwin
goarch: amd64
pkg: concurrency/examples/2workerpool
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkWorkerPool       	      42	 135979918 ns/op	    7567 B/op	     267 allocs/op
BenchmarkWorkerPool-2     	      15	  67660095 ns/op	    7975 B/op	     268 allocs/op
BenchmarkWorkerPool-4     	      30	  34906790 ns/op	    7717 B/op	     267 allocs/op
BenchmarkWorkerPool-8     	      39	  27227578 ns/op	    8114 B/op	     268 allocs/op
BenchmarkWorkerPool-12    	      42	  27211110 ns/op	    8197 B/op	     268 allocs/op
BenchmarkWorkerPool-16    	      42	  27080923 ns/op	    7971 B/op	     268 allocs/op
BenchmarkWorkerPool-24    	      43	  28651142 ns/op	    7716 B/op	     267 allocs/op
BenchmarkWorkerPool-32    	      42	  27059909 ns/op	    7757 B/op	     267 allocs/op
PASS
ok  	concurrency/examples/2workerpool	15.192s
goos: darwin
goarch: amd64
pkg: concurrency/examples/3group
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkErrGroup       	      40	 136421771 ns/op	   11498 B/op	     361 allocs/op
BenchmarkErrGroup-2     	      15	  67934691 ns/op	   11952 B/op	     362 allocs/op
BenchmarkErrGroup-4     	      32	  34894670 ns/op	   11774 B/op	     361 allocs/op
BenchmarkErrGroup-8     	      38	  27239070 ns/op	   11835 B/op	     362 allocs/op
BenchmarkErrGroup-12    	      42	  27089336 ns/op	   11687 B/op	     361 allocs/op
BenchmarkErrGroup-16    	      43	  27202439 ns/op	   11618 B/op	     361 allocs/op
BenchmarkErrGroup-24    	      42	  27173673 ns/op	   12027 B/op	     362 allocs/op
BenchmarkErrGroup-32    	      40	  27306124 ns/op	   11789 B/op	     361 allocs/op
PASS
ok  	concurrency/examples/3group	15.928s
goos: darwin
goarch: amd64
pkg: concurrency/examples/4conc
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcPool       	      40	 138652350 ns/op	   12453 B/op	     374 allocs/op
BenchmarkConcPool-2     	      15	  67425883 ns/op	   12501 B/op	     374 allocs/op
BenchmarkConcPool-4     	      32	  35138203 ns/op	   12492 B/op	     374 allocs/op
BenchmarkConcPool-8     	      40	  27173999 ns/op	   12472 B/op	     374 allocs/op
BenchmarkConcPool-12    	      42	  27769039 ns/op	   12459 B/op	     374 allocs/op
BenchmarkConcPool-16    	      43	  27181072 ns/op	   12489 B/op	     374 allocs/op
BenchmarkConcPool-24    	      40	  27079576 ns/op	   12531 B/op	     374 allocs/op
BenchmarkConcPool-32    	      40	  27531937 ns/op	   12548 B/op	     374 allocs/op
PASS
ok  	concurrency/examples/4conc	14.809s
```