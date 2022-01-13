# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIfaceMapAlloc-4       	   2497278	       468.4 ns/op	         582 B/op	       1 allocs/op
BenchmarkIntMapAlloc-4         	   3617047	       335.2 ns/op	         420 B/op	       1 allocs/op
BenchmarkStructAlloc-4         	1000000000	         0.2664 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapWrite-4       	 100000000	        10.54 ns/op	           0 B/op	       0 allocs/op
BenchmarkIntMapWrite-4         	 121685557	         9.683 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructWrite-4         	1000000000	         0.2520 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapSearch-4      	 138029095	         8.754 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapSearch-4        	 137737975	         8.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchSearch-4        	 210887877	         5.638 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch-4        	1000000000	         0.4963 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapMarshal-4     	    446660	      2590 ns/op	        1288 B/op	      27 allocs/op
BenchmarkIntMapMarshal-4       	    540987	      2159 ns/op	        1208 B/op	      27 allocs/op
BenchmarkStructMarshal-4       	   3077613	       384.8 ns/op	         144 B/op	       2 allocs/op
BenchmarkIntMapUnmarshal-4     	    267614	      4173 ns/op	        1252 B/op	      28 allocs/op
BenchmarkIfaceMapUnmarshal-4   	    243628	      4617 ns/op	        1582 B/op	      38 allocs/op
BenchmarkSliceUnmarshal-4      	    579432	      2016 ns/op	         304 B/op	       6 allocs/op
BenchmarkIfaceSetWrite-4       	  53491489	        22.21 ns/op	           0 B/op	       0 allocs/op
BenchmarkStringSetWrite-4      	 125699049	        10.18 ns/op	           0 B/op	       0 allocs/op
```
