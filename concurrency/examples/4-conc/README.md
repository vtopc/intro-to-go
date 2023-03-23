# conc

This is an example of [sourcegraph/conc](https://github.com/sourcegraph/conc).

## Features

- Use generics.
- Has different [iterators](https://pkg.go.dev/github.com/sourcegraph/conc@v0.3.0/iter).
- Propagates panic to parent goroutines.

## Pros

- Boilerplate.
- Goroutine stack starts at 2Kb only (as of Go 1.19), so it's almost free to start a new goroutine.
- Has context cancellation.
- Limits number of goroutine.

## Cons

- ResultPool uses mutex for appending to the result slice, but could do it [concurrently](https://stackoverflow.com/questions/49879322/can-i-concurrently-write-different-slice-elements).
- Too smart and not so flexible. Doesn't fit for batching or any kind of workers.
- One more API to learn.

## Benchmarks

```shell
go test -bench=. -cpu=1,2,4,8,12,16,24,32 -benchmem ./...
?   	concurrency	[no test files]
goos: darwin
goarch: amd64
pkg: concurrency/examples/conc
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcPool       	      40	 138190589 ns/op	   10564 B/op	     271 allocs/op
BenchmarkConcPool-2     	      16	  67967610 ns/op	   10796 B/op	     272 allocs/op
BenchmarkConcPool-4     	      30	  34918573 ns/op	   10609 B/op	     271 allocs/op
BenchmarkConcPool-8     	      42	  27190270 ns/op	   10613 B/op	     271 allocs/op
BenchmarkConcPool-12    	      42	  28033095 ns/op	   10600 B/op	     271 allocs/op
BenchmarkConcPool-16    	      40	  27211275 ns/op	   10629 B/op	     271 allocs/op
BenchmarkConcPool-24    	      40	  28864739 ns/op	   10637 B/op	     271 allocs/op
BenchmarkConcPool-32    	      42	  27131906 ns/op	   10711 B/op	     271 allocs/op
PASS
ok  	concurrency/examples/conc	16.034s
```
