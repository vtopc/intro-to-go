# concurrency examples

## Benchmarks

```shell
go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem ./...
?   	concurrency	[no test files]
goos: darwin
goarch: amd64
pkg: concurrency/examples/1-sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSync       	       4	 267712150 ns/op	   11450 B/op	     502 allocs/op
BenchmarkSync-2     	       4	 268384962 ns/op	   11482 B/op	     502 allocs/op
BenchmarkSync-4     	       4	 266603885 ns/op	   11614 B/op	     503 allocs/op
BenchmarkSync-8     	       4	 273017443 ns/op	   11674 B/op	     502 allocs/op
BenchmarkSync-12    	       4	 265594343 ns/op	   11870 B/op	     503 allocs/op
BenchmarkSync-16    	       4	 272306723 ns/op	   11998 B/op	     503 allocs/op
BenchmarkSync-24    	       4	 272016326 ns/op	   12186 B/op	     502 allocs/op
BenchmarkSync-32    	       4	 272076404 ns/op	   12510 B/op	     503 allocs/op
PASS
ok  	concurrency/examples/1-sync	19.093s
goos: darwin
goarch: amd64
pkg: concurrency/examples/2-workerpool
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkWorkerPool       	      39	 270364189 ns/op	   14809 B/op	     524 allocs/op
BenchmarkWorkerPool-2     	       8	 137484829 ns/op	   16421 B/op	     529 allocs/op
BenchmarkWorkerPool-4     	      16	  67589705 ns/op	   15884 B/op	     527 allocs/op
BenchmarkWorkerPool-8     	      28	  38939084 ns/op	   16246 B/op	     527 allocs/op
BenchmarkWorkerPool-12    	      37	  30364115 ns/op	   16203 B/op	     527 allocs/op
BenchmarkWorkerPool-16    	      37	  30437343 ns/op	   16031 B/op	     527 allocs/op
BenchmarkWorkerPool-24    	      34	  30367533 ns/op	   15237 B/op	     525 allocs/op
BenchmarkWorkerPool-32    	      34	  30333497 ns/op	   15476 B/op	     526 allocs/op
PASS
ok  	concurrency/examples/2-workerpool	20.338s
goos: darwin
goarch: amd64
pkg: concurrency/examples/3.1-group
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkErrGroup       	      39	 272960016 ns/op	   22762 B/op	     713 allocs/op
BenchmarkErrGroup-2     	       8	 135569210 ns/op	   24741 B/op	     719 allocs/op
BenchmarkErrGroup-4     	      16	  67354721 ns/op	   24366 B/op	     719 allocs/op
BenchmarkErrGroup-8     	      30	  38784837 ns/op	   24255 B/op	     718 allocs/op
BenchmarkErrGroup-12    	      37	  31463890 ns/op	   23664 B/op	     715 allocs/op
BenchmarkErrGroup-16    	      37	  31585565 ns/op	   23798 B/op	     715 allocs/op
BenchmarkErrGroup-24    	      34	  31752179 ns/op	   23712 B/op	     715 allocs/op
BenchmarkErrGroup-32    	      33	  32376797 ns/op	   23094 B/op	     714 allocs/op
PASS
ok  	concurrency/examples/3.1-group	20.730s
goos: darwin
goarch: amd64
pkg: concurrency/examples/3.2-group-prealloc
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkErrGroupPrealloc       	      39	 270523236 ns/op	   18873 B/op	     506 allocs/op
BenchmarkErrGroupPrealloc-2     	       8	 135225771 ns/op	   20641 B/op	     510 allocs/op
BenchmarkErrGroupPrealloc-4     	      16	  67233415 ns/op	   20322 B/op	     509 allocs/op
BenchmarkErrGroupPrealloc-8     	      28	  38779372 ns/op	   19888 B/op	     508 allocs/op
BenchmarkErrGroupPrealloc-12    	      37	  31443647 ns/op	   20185 B/op	     508 allocs/op
BenchmarkErrGroupPrealloc-16    	      36	  31443748 ns/op	   19641 B/op	     507 allocs/op
BenchmarkErrGroupPrealloc-24    	      36	  32273943 ns/op	   19090 B/op	     506 allocs/op
BenchmarkErrGroupPrealloc-32    	      34	  32488216 ns/op	   19423 B/op	     507 allocs/op
PASS
ok  	concurrency/examples/3.2-group-prealloc	20.642s
goos: darwin
goarch: amd64
pkg: concurrency/examples/4-conc
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcPool       	      38	 280465047 ns/op	   25175 B/op	     839 allocs/op
BenchmarkConcPool-2     	       8	 133606918 ns/op	   25345 B/op	     839 allocs/op
BenchmarkConcPool-4     	      16	  67298141 ns/op	   25542 B/op	     840 allocs/op
BenchmarkConcPool-8     	      30	  39037080 ns/op	   25407 B/op	     840 allocs/op
BenchmarkConcPool-12    	      36	  30749583 ns/op	   25259 B/op	     839 allocs/op
BenchmarkConcPool-16    	      34	  31092331 ns/op	   25442 B/op	     840 allocs/op
BenchmarkConcPool-24    	      34	  31961767 ns/op	   25625 B/op	     840 allocs/op
BenchmarkConcPool-32    	      36	  31643832 ns/op	   25497 B/op	     839 allocs/op
PASS
ok  	concurrency/examples/4-conc	20.906s
```
