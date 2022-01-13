# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIfaceMapAlloc-4                 	   2560513	       466.5 ns/op	         582 B/op	       1 allocs/op
BenchmarkIntMapAlloc-4                   	   3874237	       302.0 ns/op	         420 B/op	       1 allocs/op
BenchmarkSliceAlloc-4                    	1000000000	         0.2517 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructAlloc-4                   	1000000000	         0.2551 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapWrite-4                 	 100000000	        10.23 ns/op	           0 B/op	       0 allocs/op
BenchmarkIntMapWrite-4                   	 123532088	         9.790 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructWrite-4                   	1000000000	         0.2529 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapAvgSearch10Elems-4      	 122859554	         9.625 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapAvgSearch10Elems-4        	 126457390	        10.24 ns/op	           0 B/op	       0 allocs/op
BenchmarkIntSwitchAvgSearch10Elems-4     	 213067028	         5.585 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntSwitchWorstSearch10Elems-4   	 179469534	         6.648 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch10Elems-4           	1000000000	         0.4985 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAvgSearch10Elems-4         	 373249536	         3.238 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceWorstSearch10Elems-4       	 189094639	         6.356 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapMarshal-4               	    432094	      2571 ns/op	        1288 B/op	      27 allocs/op
BenchmarkIntMapMarshal-4                 	    535476	      2233 ns/op	        1208 B/op	      27 allocs/op
BenchmarkStructMarshal-4                 	   3099561	       390.6 ns/op	         144 B/op	       2 allocs/op
BenchmarkIntMapUnmarshal-4               	    252367	      6713 ns/op	        1252 B/op	      28 allocs/op
BenchmarkIfaceMapUnmarshal-4             	    123322	      8808 ns/op	        1582 B/op	      38 allocs/op
BenchmarkStructUnmarshal-4               	    339848	      3135 ns/op	         304 B/op	       6 allocs/op
BenchmarkIfaceSetWrite-4                 	  52636830	        22.66 ns/op	           0 B/op	       0 allocs/op
BenchmarkStringSetWrite-4                	 100000000	        10.56 ns/op	           0 B/op	       0 allocs/o
```
