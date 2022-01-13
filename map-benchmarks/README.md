# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIfaceMapAlloc-4                 	   2439028	       471.4 ns/op	     582 B/op	       1 allocs/op
BenchmarkIntMapAlloc-4                   	   3829736	       309.4 ns/op	     420 B/op	       1 allocs/op
BenchmarkSliceAlloc-4                    	1000000000	         0.2533 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructAlloc-4                   	1000000000	         0.2505 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapWrite-4                 	 123392436	         9.822 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapWrite-4                   	 124733403	         9.671 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructWrite-4                   	1000000000	         0.2508 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapAvgSearch10Elems-4      	 124899188	         8.189 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapAvgSearch10Elems-4        	 132753459	         8.870 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntSwitchAvgSearch10Elems-4     	 209465904	         5.828 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntSwitchWorstSearch10Elems-4   	 176948509	         6.800 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch10Elems-4           	1000000000	         0.2488 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceAvgSearch10Elems-4         	 344528163	         3.480 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceWorstSearch10Elems-4       	 160719057	         7.459 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapMarshal-4               	    421802	      2619 ns/op	    1288 B/op	      27 allocs/op
BenchmarkIntMapMarshal-4                 	    457372	      2248 ns/op	    1208 B/op	      27 allocs/op
BenchmarkStructMarshal-4                 	   2937532	       385.7 ns/op	     144 B/op	       2 allocs/op
BenchmarkIntMapUnmarshal-4               	    239440	      4250 ns/op	    1252 B/op	      28 allocs/op
BenchmarkIfaceMapUnmarshal-4             	    248637	      4616 ns/op	    1582 B/op	      38 allocs/op
BenchmarkStructUnmarshal-4               	    581743	      2106 ns/op	     304 B/op	       6 allocs/op
BenchmarkIfaceSetWrite-4                 	  52333282	        22.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringSetWrite-4                	 100000000	        10.30 ns/op	       0 B/op	       0 allocs/op
```
