# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkAllocMapIface-4                 	 2028284	       591.9 ns/op	     918 B/op	       3 allocs/op
BenchmarkAllocMapInt-4                   	 3235831	       363.0 ns/op	     468 B/op	       2 allocs/op
BenchmarkAllocSlice-4                    	19449904	        61.20 ns/op	     240 B/op	       1 allocs/op
BenchmarkAllocPtrSlice-4                 	 4365212	       271.3 ns/op	     320 B/op	      11 allocs/op
BenchmarkAllocStruct-4                   	313006070	         3.815 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteMapIface-4                 	125949531	         9.533 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteMapInt-4                   	132700627	         9.014 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteStruct-4                   	1000000000	         0.2554 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsMapIfaceAvg-4      	100000000	        10.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsMapIntAvg-4        	133037317	         7.807 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSwitchIntAvg-4     	218128281	         5.502 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSwitchIntWorst-4   	172038668	         6.974 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSliceAvg-4         	271148046	         4.219 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSliceWorst-4       	199760920	         5.978 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsStruct-4           	1000000000	         0.2777 ns/op	       0 B/op	       0 allocs/op
BenchmarkMarshalMapIface-4               	  421107	      2595 ns/op	    1288 B/op	      27 allocs/op
BenchmarkMarshalMapInt-4                 	  511644	      2160 ns/op	    1208 B/op	      27 allocs/op
BenchmarkMarshalStruct-4                 	 3142605	       394.1 ns/op	     144 B/op	       2 allocs/op
BenchmarkUnmarshalMapIntPrealloc-4       	  302758	      3887 ns/op	    1044 B/op	      27 allocs/op
BenchmarkUnmarshalMapInt-4               	  272281	      4204 ns/op	    1252 B/op	      28 allocs/op
BenchmarkUnmarshalMapIface-4             	  241986	      4573 ns/op	    1582 B/op	      38 allocs/op
BenchmarkUnmarshalStruct-4               	  582028	      2042 ns/op	     304 B/op	       6 allocs/op
BenchmarkSetWriteIface-4                 	50338532	        23.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkSetWriteString-4                	130943379	         9.245 ns/op	       0 B/op	       0 allocs/op
```
