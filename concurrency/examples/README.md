# concurrency examples

## Benchmarks

```shell
go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem ./...
?   	concurrency	[no test files]
goos: darwin
goarch: amd64
pkg: concurrency/examples/1-sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSync       	       4	 276720742 ns/op	   11450 B/op	     502 allocs/op
BenchmarkSync-2     	       4	 267847693 ns/op	   11550 B/op	     503 allocs/op
BenchmarkSync-4     	       4	 274887652 ns/op	   11614 B/op	     503 allocs/op
BenchmarkSync-8     	       4	 272057458 ns/op	   11674 B/op	     502 allocs/op
BenchmarkSync-12    	       4	 262413964 ns/op	   11870 B/op	     503 allocs/op
BenchmarkSync-16    	       4	 267116446 ns/op	   11998 B/op	     503 allocs/op
BenchmarkSync-24    	       4	 296667514 ns/op	   12254 B/op	     503 allocs/op
BenchmarkSync-32    	       4	 292239467 ns/op	   12510 B/op	     503 allocs/op
PASS
ok  	concurrency/examples/1-sync	19.661s
goos: darwin
goarch: amd64
pkg: concurrency/examples/2-workerpool
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkWorkerPool       	      39	 276705599 ns/op	   14830 B/op	     524 allocs/op
BenchmarkWorkerPool-2     	       8	 133726390 ns/op	   15701 B/op	     528 allocs/op
BenchmarkWorkerPool-4     	      16	  68740375 ns/op	   15476 B/op	     527 allocs/op
BenchmarkWorkerPool-8     	      28	  38570249 ns/op	   16033 B/op	     527 allocs/op
BenchmarkWorkerPool-12    	      38	  30222389 ns/op	   15896 B/op	     526 allocs/op
BenchmarkWorkerPool-16    	      34	  30872055 ns/op	   16013 B/op	     528 allocs/op
BenchmarkWorkerPool-24    	      34	  31402787 ns/op	   15386 B/op	     526 allocs/op
BenchmarkWorkerPool-32    	      37	  30620674 ns/op	   16092 B/op	     527 allocs/op
PASS
ok  	concurrency/examples/2-workerpool	20.716s
goos: darwin
goarch: amd64
pkg: concurrency/examples/3.1-group
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkErrGroup       	      39	 270530705 ns/op	   22750 B/op	     713 allocs/op
BenchmarkErrGroup-2     	       8	 133940065 ns/op	   23621 B/op	     717 allocs/op
BenchmarkErrGroup-4     	      16	  67341826 ns/op	   23552 B/op	     717 allocs/op
BenchmarkErrGroup-8     	      30	  39016636 ns/op	   23193 B/op	     714 allocs/op
BenchmarkErrGroup-12    	      37	  30892468 ns/op	   23708 B/op	     715 allocs/op
BenchmarkErrGroup-16    	      34	  35173348 ns/op	   23278 B/op	     714 allocs/op
BenchmarkErrGroup-24    	      31	  34612960 ns/op	   23170 B/op	     714 allocs/op
BenchmarkErrGroup-32    	      32	  35033246 ns/op	   22861 B/op	     713 allocs/op
PASS
ok  	concurrency/examples/3.1-group	22.560s
goos: darwin
goarch: amd64
pkg: concurrency/examples/3.2-group-prealloc
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkErrGroupPrealloc       	      37	 283008617 ns/op	   18874 B/op	     506 allocs/op
BenchmarkErrGroupPrealloc-2     	       8	 144343527 ns/op	   20201 B/op	     509 allocs/op
BenchmarkErrGroupPrealloc-4     	      16	  69865480 ns/op	   20413 B/op	     510 allocs/op
BenchmarkErrGroupPrealloc-8     	      30	  38695941 ns/op	   19334 B/op	     507 allocs/op
BenchmarkErrGroupPrealloc-12    	      37	  31627891 ns/op	   19955 B/op	     508 allocs/op
BenchmarkErrGroupPrealloc-16    	      34	  32150653 ns/op	   19426 B/op	     507 allocs/op
BenchmarkErrGroupPrealloc-24    	      34	  32317635 ns/op	   19711 B/op	     508 allocs/op
BenchmarkErrGroupPrealloc-32    	      34	  33254914 ns/op	   19522 B/op	     507 allocs/op
PASS
ok  	concurrency/examples/3.2-group-prealloc	22.563s
goos: darwin
goarch: amd64
pkg: concurrency/examples/4-conc
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcPool       	      27	 277422255 ns/op	   25188 B/op	     839 allocs/op
BenchmarkConcPool-2     	       8	 134669752 ns/op	   25785 B/op	     840 allocs/op
BenchmarkConcPool-4     	      16	  70695686 ns/op	   25884 B/op	     841 allocs/op
BenchmarkConcPool-8     	      28	  38587511 ns/op	   25409 B/op	     840 allocs/op
BenchmarkConcPool-12    	      36	  30689658 ns/op	   25506 B/op	     840 allocs/op
BenchmarkConcPool-16    	      37	  31437338 ns/op	   25322 B/op	     839 allocs/op
BenchmarkConcPool-24    	      36	  32078184 ns/op	   25432 B/op	     839 allocs/op
BenchmarkConcPool-32    	      34	  32057813 ns/op	   25343 B/op	     839 allocs/op
PASS
ok  	concurrency/examples/4-conc	17.710s
goos: darwin
goarch: amd64
pkg: concurrency/examples/5-longrunning
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkLongRunning       	      39	 273748452 ns/op	   23221 B/op	     717 allocs/op
BenchmarkLongRunning-2     	       8	 135774122 ns/op	   23305 B/op	     718 allocs/op
BenchmarkLongRunning-4     	      16	  67165545 ns/op	   24414 B/op	     720 allocs/op
BenchmarkLongRunning-8     	      28	  38664868 ns/op	   24859 B/op	     723 allocs/op
BenchmarkLongRunning-12    	      38	  30388611 ns/op	   24672 B/op	     720 allocs/op
BenchmarkLongRunning-16    	      33	  31083696 ns/op	   24146 B/op	     719 allocs/op
BenchmarkLongRunning-24    	      36	  32141456 ns/op	   23509 B/op	     717 allocs/op
BenchmarkLongRunning-32    	      33	  32446626 ns/op	   23374 B/op	     717 allocs/op
PASS
ok  	concurrency/examples/5-longrunning	21.564s
```
