# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIfaceMapAlloc-4                 	   2004290	       583.7 ns/op	     918 B/op	       3 allocs/op
BenchmarkIntMapAlloc-4                   	   3328993	       347.4 ns/op	     468 B/op	       2 allocs/op
BenchmarkSliceAlloc-4                    	  17106690	        64.45 ns/op	     240 B/op	       1 allocs/op
BenchmarkStructAlloc-4                   	 308639869	         3.810 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapWrite-4                 	 123062497	         9.808 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapWrite-4                   	 121491268	         9.536 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructWrite-4                   	1000000000	         0.2485 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapAvgSearch10Elems-4      	 100000000	        11.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapAvgSearch10Elems-4        	 130678629	         9.158 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntSwitchAvgSearch10Elems-4     	 208487640	         5.625 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntSwitchWorstSearch10Elems-4   	 145732152	         8.131 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAvgSearch10Elems-4         	 283927846	         4.189 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceWorstSearch10Elems-4       	 200096373	         5.970 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch10Elems-4           	1000000000	         0.2813 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapMarshal-4               	    412387	      2651 ns/op	    1288 B/op	      27 allocs/op
BenchmarkIntMapMarshal-4                 	    532070	      2186 ns/op	    1208 B/op	      27 allocs/op
BenchmarkStructMarshal-4                 	   2979066	       385.2 ns/op	     144 B/op	       2 allocs/op
BenchmarkIntMapUnmarshal-4               	    274398	      4298 ns/op	    1252 B/op	      28 allocs/op
BenchmarkIfaceMapUnmarshal-4             	    244969	      4729 ns/op	    1582 B/op	      38 allocs/op
BenchmarkStructUnmarshal-4               	    561007	      2052 ns/op	     304 B/op	       6 allocs/op
BenchmarkIfaceSetWrite-4                 	  48727297	        22.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringSetWrite-4                	 122805477	         9.999 ns/op	       0 B/op	       0 allocs/op
```
