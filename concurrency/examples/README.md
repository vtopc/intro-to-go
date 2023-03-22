# concurrency examples

## Benchmarks

```shell
go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem ./...
?   	concurrency	[no test files]
goos: darwin
goarch: amd64
pkg: concurrency/examples/conc
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcPool       	      40	 138190589 ns/op	   10564 B/op	     271 allocs/op
BenchmarkConcPool-2     	      16	  67967610 ns/op	   10796 B/op	     272 allocs/op
BenchmarkConcPool-4     	      30	  34918573 ns/op	   10609 B/op	     271 allocs/op
BenchmarkConcPool-8     	      42	  27190270 ns/op	   10613 B/op	     271 allocs/op
BenchmarkConcPool-12    	      42	  28033095 ns/op	   10600 B/op	     271 allocs/op
BenchmarkConcPool-16    	      40	  27211275 ns/op	   10629 B/op	     271 allocs/op
BenchmarkConcPool-24    	      40	  28864739 ns/op	   10637 B/op	     271 allocs/op
BenchmarkConcPool-32    	      42	  27131906 ns/op	   10711 B/op	     271 allocs/op
PASS
ok  	concurrency/examples/conc	16.034s
goos: darwin
goarch: amd64
pkg: concurrency/examples/group
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkErrGroup       	      43	 136600393 ns/op	   11486 B/op	     361 allocs/op
BenchmarkErrGroup-2     	      16	  67556515 ns/op	   12690 B/op	     364 allocs/op
BenchmarkErrGroup-4     	      32	  34806718 ns/op	   11826 B/op	     362 allocs/op
BenchmarkErrGroup-8     	      39	  27301936 ns/op	   11795 B/op	     361 allocs/op
BenchmarkErrGroup-12    	      42	  27729244 ns/op	   11686 B/op	     361 allocs/op
BenchmarkErrGroup-16    	      42	  27320862 ns/op	   11753 B/op	     361 allocs/op
BenchmarkErrGroup-24    	      42	  27940932 ns/op	   11655 B/op	     361 allocs/op
BenchmarkErrGroup-32    	      42	  27056194 ns/op	   11824 B/op	     361 allocs/op
PASS
ok  	concurrency/examples/group	15.255s
goos: darwin
goarch: amd64
pkg: concurrency/examples/sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSync       	       8	 142392084 ns/op	    5737 B/op	     252 allocs/op
BenchmarkSync-2     	       7	 143014854 ns/op	    5796 B/op	     252 allocs/op
BenchmarkSync-4     	       8	 134557020 ns/op	    5785 B/op	     252 allocs/op
BenchmarkSync-8     	       8	 137267906 ns/op	    5883 B/op	     252 allocs/op
BenchmarkSync-12    	       8	 136586178 ns/op	    5947 B/op	     252 allocs/op
BenchmarkSync-16    	       8	 134556868 ns/op	    5977 B/op	     252 allocs/op
BenchmarkSync-24    	       8	 139400376 ns/op	    6139 B/op	     252 allocs/op
BenchmarkSync-32    	       8	 134481762 ns/op	    6267 B/op	     252 allocs/op
PASS
ok  	concurrency/examples/sync	14.078s
goos: darwin
goarch: amd64
pkg: concurrency/examples/workerpool
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkWorkerPool       	      40	 137403527 ns/op	    7569 B/op	     267 allocs/op
BenchmarkWorkerPool-2     	      16	  67681861 ns/op	    7988 B/op	     268 allocs/op
BenchmarkWorkerPool-4     	      31	  34890644 ns/op	    7749 B/op	     267 allocs/op
BenchmarkWorkerPool-8     	      42	  29870561 ns/op	    8114 B/op	     268 allocs/op
BenchmarkWorkerPool-12    	      43	  27218258 ns/op	    7741 B/op	     267 allocs/op
BenchmarkWorkerPool-16    	      40	  27222840 ns/op	    7636 B/op	     267 allocs/op
BenchmarkWorkerPool-24    	      42	  27389846 ns/op	    7992 B/op	     268 allocs/op
BenchmarkWorkerPool-32    	      42	  27209868 ns/op	    7937 B/op	     268 allocs/op
PASS
ok  	concurrency/examples/workerpool	16.637s
```