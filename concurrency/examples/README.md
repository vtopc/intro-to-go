# concurrency examples

## Benchmarks

```shell
go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem ./...
?   	concurrency	[no test files]
goos: darwin
goarch: amd64
pkg: concurrency/examples/group
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkErrGroup       	      40	 133948931 ns/op	   10791 B/op	     379 allocs/op
BenchmarkErrGroup-2     	      16	  65901210 ns/op	   11590 B/op	     381 allocs/op
BenchmarkErrGroup-4     	      32	  34955429 ns/op	   11040 B/op	     379 allocs/op
BenchmarkErrGroup-8     	      40	  26981328 ns/op	   11303 B/op	     380 allocs/op
BenchmarkErrGroup-12    	      40	  27026977 ns/op	   11129 B/op	     380 allocs/op
BenchmarkErrGroup-16    	      42	  28974379 ns/op	   11002 B/op	     379 allocs/op
BenchmarkErrGroup-24    	      39	  27133207 ns/op	   10893 B/op	     379 allocs/op
BenchmarkErrGroup-32    	      42	  28333549 ns/op	   10999 B/op	     379 allocs/op
PASS
ok  	concurrency/examples/group	14.619s
goos: darwin
goarch: amd64
pkg: concurrency/examples/sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSync       	       8	 131166579 ns/op	    5737 B/op	     252 allocs/op
BenchmarkSync-2     	       8	 133288799 ns/op	    5753 B/op	     252 allocs/op
BenchmarkSync-4     	       8	 135637828 ns/op	    5819 B/op	     252 allocs/op
BenchmarkSync-8     	       8	 135954035 ns/op	    5883 B/op	     252 allocs/op
BenchmarkSync-12    	       8	 133825792 ns/op	    5913 B/op	     252 allocs/op
BenchmarkSync-16    	       8	 134349121 ns/op	    5977 B/op	     252 allocs/op
BenchmarkSync-24    	       8	 133118355 ns/op	    6139 B/op	     252 allocs/op
BenchmarkSync-32    	       8	 140350428 ns/op	    6267 B/op	     252 allocs/op
PASS
ok  	concurrency/examples/sync	13.616s
goos: darwin
goarch: amd64
pkg: concurrency/examples/workerpool
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkWorkerPool       	      42	 134638651 ns/op	    6882 B/op	     285 allocs/op
BenchmarkWorkerPool-2     	      16	  68279144 ns/op	    7076 B/op	     285 allocs/op
BenchmarkWorkerPool-4     	      31	  35037120 ns/op	    7298 B/op	     286 allocs/op
BenchmarkWorkerPool-8     	      39	  26921808 ns/op	    7434 B/op	     286 allocs/op
BenchmarkWorkerPool-12    	      40	  27269997 ns/op	    7035 B/op	     285 allocs/op
BenchmarkWorkerPool-16    	      42	  27007829 ns/op	    7115 B/op	     285 allocs/op
BenchmarkWorkerPool-24    	      40	  27222673 ns/op	    7115 B/op	     285 allocs/op
BenchmarkWorkerPool-32    	      40	  26942670 ns/op	    7265 B/op	     286 allocs/op
PASS
ok  	concurrency/examples/workerpool	14.613s
```