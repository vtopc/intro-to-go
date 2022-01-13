# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIfaceMapAlloc-4       	   2474974	       468.3 ns/op	         582 B/op	       1 allocs/op
BenchmarkIntMapAlloc-4         	   3764715	       318.7 ns/op	         420 B/op	       1 allocs/op
BenchmarkStructAlloc-4         	1000000000	         0.2573 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapWrite-4       	 100000000	        10.35 ns/op	           0 B/op	       0 allocs/op
BenchmarkIntMapWrite-4         	 122433998	         9.639 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructWrite-4         	1000000000	         0.2497 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapSearch-4      	 136487755	         9.535 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapSearch-4        	 147279693	         8.808 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchSearch-4        	 207983079	         5.594 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch-4        	1000000000	         0.2494 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapMarshal-4     	    443677	      2603 ns/op	        1288 B/op	      27 allocs/op
BenchmarkIntMapMarshal-4       	    506270	      2164 ns/op	        1208 B/op	      27 allocs/op
BenchmarkStructMarshal-4       	   3086096	       387.1 ns/op	         144 B/op	       2 allocs/op
BenchmarkIntMapUnmarshal-4     	    252643	      4178 ns/op	        1252 B/op	      28 allocs/op
BenchmarkIfaceMapUnmarshal-4   	    251274	      4605 ns/op	        1582 B/op	      38 allocs/op
BenchmarkSliceUnmarshal-4      	    597054	      2012 ns/op	         304 B/op	       6 allocs/op
BenchmarkIfaceSetWrite-4       	  50754938	        22.68 ns/op	           0 B/op	       0 allocs/op
BenchmarkStringSetWrite-4      	 125685156	         9.540 ns/op	       0 B/op	       0 allocs/op
```
