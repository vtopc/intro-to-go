# Concurrency

## Goroutines

Every function spawned with `go` keyword creates new goroutine, e.g.:
```go
go fn(x, y, z)
```
[Check a tour of Go](https://go.dev/tour/concurrency/1)

### Before spawning goroutine(checklist):

1. Consider to write sync code. Concurrency in Go is not a silver bullet, it wouldn't make code faster in all cases.

    > Async code harder to understand and debug is. Hrrmmm.

    Master Yoda

    Also, there is no speed benefit in [this example](./examples/README.md) for **one** CPU. 
    And async solution consumes 50% more memory.


2. Make sure, that there would be limited amount of running goroutines. Use one of next:
- use [semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)
- use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) with [SetLimit()](https://pkg.go.dev/golang.org/x/sync/errgroup#Group.SetLimit)
- create [worker pool](https://gobyexample.com/worker-pools) 
- use pool from [conc](https://github.com/sourcegraph/conc) with `WithMaxGoroutines`
- etc.

    The semaphore/errgroup is preferable in most cases, since with the worker pool there would be hanging workers, that doing nothing and spawning a new goroutine is quite cheap.

3. It should be possible to stop spawned goroutines, e.g. on service shutdown. Use one of next for cancellation:
- `context.Context`(preferable)
- done channel

4. [Never start a goroutine without knowing when it will stop](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_never_start_a_goroutine_without_knowing_when_it_will_stop). 
On shutdown app should wait for all goroutines to stop.

5. **Panics.** Recover could catch panic only in current goroutine, so make sure, that [panic is handled in goroutine](https://medium.com/codex/handle-panic-in-go-routine-54b82d6013d3).

#### To sum up

- either use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) ...

- or [conc](https://github.com/sourcegraph/conc)

_TBA examples_

### Tips and tricks

- Add [profiler labels](https://rakyll.org/profiler-labels/), this would help to debug and read stacktrace.

## Channels

Channels are a typed conduit through which you can send and receive values.

[Check a tour of Go](https://go.dev/tour/concurrency/2)

### Channels usage checklist

1. Don't make huge buffered channels. Channel is just a [data buffer](https://en.wikipedia.org/wiki/Data_buffer),
don't try to feet all results there.

2. The channel consumer should be spawned before channel producer(writer) and write values into DB/cache/file/socket/map/slice/other data structures. 
