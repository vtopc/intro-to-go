# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkAllocMapIface-4                 	 1866698	       618.6 ns/op	     918 B/op	       3 allocs/op
BenchmarkAllocMapInt-4                   	 3223738	       357.3 ns/op	     468 B/op	       2 allocs/op
BenchmarkAllocSlice-4                    	16317514	        70.14 ns/op	     240 B/op	       1 allocs/op
BenchmarkAllocPtrSlice-4                 	 4038168	       277.2 ns/op	     320 B/op	      11 allocs/op
BenchmarkAllocHugeSlice-4                	 1000000	      1021 ns/op	   10240 B/op	       1 allocs/op
BenchmarkAllocHugePtrSlice-4             	  739033	      1513 ns/op	   10320 B/op	      11 allocs/op
BenchmarkAllocStruct-4                   	303304056	         3.873 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteMapIface-4                 	128140224	         9.381 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteMapInt-4                   	130739418	         9.069 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteStruct-4                   	1000000000	         0.2537 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsMapIfaceAvg-4      	100000000	        12.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsMapIntAvg-4        	133038070	         9.487 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSwitchIntAvg-4     	212355748	         5.755 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSwitchIntWorst-4   	172580755	         7.020 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSliceAvg-4         	324067020	         3.696 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSliceWorst-4       	153205366	         7.804 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsStruct-4           	1000000000	         0.5054 ns/op	       0 B/op	       0 allocs/op
BenchmarkMarshalMapIface-4               	  432603	      2709 ns/op	    1288 B/op	      27 allocs/op
BenchmarkMarshalMapInt-4                 	  518820	      2206 ns/op	    1208 B/op	      27 allocs/op
BenchmarkMarshalStruct-4                 	 3099295	       388.8 ns/op	     144 B/op	       2 allocs/op
BenchmarkUnmarshalMapIntPrealloc-4       	  283684	      3948 ns/op	    1044 B/op	      27 allocs/op
BenchmarkUnmarshalMapInt-4               	  237858	      4288 ns/op	    1252 B/op	      28 allocs/op
BenchmarkUnmarshalMapIface-4             	  248125	      4669 ns/op	    1582 B/op	      38 allocs/op
BenchmarkUnmarshalStruct-4               	  585271	      2054 ns/op	     304 B/op	       6 allocs/op
BenchmarkSetWriteIface-4                 	45962313	        24.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkSetWriteString-4                	127555977	         9.495 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapLookup2000Elems-4            	194989315	         6.028 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLookup2000ElemsAvg-4       	 4327219	       272.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceLookup2000ElemsWorst-4     	  549072	      2188 ns/op	       0 B/op	       0 allocs/op
```
