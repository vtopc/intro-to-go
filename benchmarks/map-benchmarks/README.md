# Struct vs. Maps

## Results
```shell
go test -bench=. -cpu=4 -benchmem
goos: darwin
goarch: amd64
pkg: map-benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkAllocMapIface-4                 	   1922929	       589.3 ns/op	     918 B/op	       3 allocs/op
BenchmarkAllocMapInt-4                   	   3376296	       355.4 ns/op	     468 B/op	       2 allocs/op
BenchmarkAllocSlice-4                    	  16424552	        64.67 ns/op	     240 B/op	       1 allocs/op
BenchmarkAllocStruct-4                   	 300585793	         3.666 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteMapIface-4                 	 127716265	         9.557 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteMapInt-4                   	 128018139	         9.198 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteStruct-4                   	1000000000	         0.2521 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsMapIfaceAvg-4      	 100000000	        11.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsMapIntAvg-4        	 125568327	         8.537 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSwitchIntAvg-4     	 219495506	         5.451 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSwitchIntWorst-4   	 152993018	         7.824 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSliceAvg-4         	 286470931	         4.119 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsSliceWorst-4       	 203158540	         5.836 ns/op	       0 B/op	       0 allocs/op
BenchmarkSearch10ElemsStruct-4           	1000000000	         0.2730 ns/op	       0 B/op	       0 allocs/op
BenchmarkMarshalMapIface-4               	    451089	      2517 ns/op	    1288 B/op	      27 allocs/op
BenchmarkMarshalMapInt-4                 	    557138	      2161 ns/op	    1208 B/op	      27 allocs/op
BenchmarkMarshalStruct-4                 	   3088924	       393.4 ns/op	     144 B/op	       2 allocs/op
BenchmarkUnmarshalMapInt-4               	    281619	      4121 ns/op	    1252 B/op	      28 allocs/op
BenchmarkUnmarshalMapIface-4             	    219394	      4584 ns/op	    1582 B/op	      38 allocs/op
BenchmarkUnmarshalStruct-4               	    571075	      2012 ns/op	     304 B/op	       6 allocs/op
BenchmarkSetWriteIface-4                 	  51906891	        22.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkSetWriteString-4                	 126055957	         9.299 ns/op	       0 B/op	       0 allocs/op
```
