# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=1,4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIfaceMapAlloc         	   2306450	       479.7 ns/op	         582 B/op	       1 allocs/op
BenchmarkIfaceMapAlloc-4       	   2554461	       485.7 ns/op	         582 B/op	       1 allocs/op
BenchmarkIntMapAlloc           	   3820524	       305.3 ns/op	         420 B/op	       1 allocs/op
BenchmarkIntMapAlloc-4         	   3841504	       305.1 ns/op	         420 B/op	       1 allocs/op
BenchmarkStructAlloc           	1000000000	         0.2453 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructAlloc-4         	1000000000	         0.2490 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapSearch        	 136473198	         8.520 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapSearch-4      	 150598351	         8.211 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntMapSearch          	 158145339	        10.01 ns/op	           0 B/op	       0 allocs/op
BenchmarkIntMapSearch-4        	 147792490	         8.554 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchSearch          	 204488901	         5.753 ns/op	       0 B/op	       0 allocs/op
BenchmarkSwitchSearch-4        	 208450507	         5.731 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch          	1000000000	         0.2635 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructSearch-4        	1000000000	         0.2468 ns/op	       0 B/op	       0 allocs/op
BenchmarkIfaceMapMarshal       	    440292	      2570 ns/op	        1288 B/op	      27 allocs/op
BenchmarkIfaceMapMarshal-4     	    448000	      2603 ns/op	        1288 B/op	      27 allocs/op
BenchmarkIntMapMarshal         	    522098	      2290 ns/op	        1208 B/op	      27 allocs/op
BenchmarkIntMapMarshal-4       	    497221	      2199 ns/op	        1208 B/op	      27 allocs/op
BenchmarkStructMarshal         	   2908034	       404.6 ns/op	         144 B/op	       2 allocs/op
BenchmarkStructMarshal-4       	   3072655	       383.9 ns/op	         144 B/op	       2 allocs/op
BenchmarkIntMapUnmarshal       	    282876	      4256 ns/op	        1252 B/op	      28 allocs/op
BenchmarkIntMapUnmarshal-4     	    272235	      4208 ns/op	        1252 B/op	      28 allocs/op
BenchmarkIfaceMapUnmarshal     	    249858	      4667 ns/op	        1582 B/op	      38 allocs/op
BenchmarkIfaceMapUnmarshal-4   	    233522	      4745 ns/op	        1582 B/op	      38 allocs/op
BenchmarkSliceUnmarshal        	    600792	      2072 ns/op	         304 B/op	       6 allocs/op
BenchmarkSliceUnmarshal-4      	    606438	      2066 ns/op	         304 B/op	       6 allocs/op
```
