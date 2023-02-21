# Concurrency

## Goroutines

Every function spawned with `go` keyword creates new goroutine, e.g.:
```go
go fn(x, y, z)
```
[Check a tour of Go](https://go.dev/tour/concurrency/1)

### Before spawning a new goroutine(the checklist):

1. Consider to write sync code. Concurrency in Go is not a silver bullet, it wouldn't make code faster in all cases.

    > Async code harder to understand and debug is. Hrrmmm.

    Master Yoda

    e.g. there is no speed benefit in [this example](./examples/README.md) for **one** CPU,
    and async solution consumes 50% more memory.


2. Make sure, that there would be a limited amount of running goroutines. Use one of next:
- use [semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)
- use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) with [SetLimit()](https://pkg.go.dev/golang.org/x/sync/errgroup#Group.SetLimit)
- create [worker pool](https://gobyexample.com/worker-pools) 
- use pool from [conc](https://github.com/sourcegraph/conc) with `WithMaxGoroutines`
- etc.

    The semaphore/errgroup is preferable in most cases, since with the worker pool there would be hanging workers, that doing nothing and spawning a new goroutine is quite cheap.

3. It should be possible to stop spawned goroutines, e.g. on service shutdown or HTTP timeout. Use one of next for cancellation:
- `context.Context`(preferable)
- done channel

4. [Never start a goroutine without knowing when it will stop](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_never_start_a_goroutine_without_knowing_when_it_will_stop). 
On shutdown app should wait for all goroutines to stop.

5. **Panics.** Recover could catch panic only in current goroutine, so make sure, that [panic is handled in goroutine](https://medium.com/codex/handle-panic-in-go-routine-54b82d6013d3).
[Bad example](https://play.golang.com/p/lVfDUZTz4ji).

#### To sum up. The easiest ways.

- either use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) ...

- or [conc](https://github.com/sourcegraph/conc)

_TBA examples_

### Tips and tricks

- Add [profiler labels](https://rakyll.org/profiler-labels/), this would help to debug and read stacktrace.
- Use the race detector (`-race` flag) in unit tests.
- Use [uber-go/goleak](https://github.com/uber-go/goleak) to detect goroutine leaks in your tests.
- You can [update pre-allocated slice concurrently](https://stackoverflow.com/questions/49879322/can-i-concurrently-write-different-slice-elements).

## Channels

Channels are a typed conduit through which you can send and receive messages.

[Check a tour of Go](https://go.dev/tour/concurrency/2)

### Possible channel issues

1. Consider using shared memory instead of message passing:

    >  Our study found that message passing does not necessarily make multithreaded programs less error-prone than shared memory.
       In fact, message passing is the main cause of blocking bugs.
       To make it worse, when combined with traditional synchronization primitives or with other new language features
       and libraries, message passing can cause blocking bugs that
       are very hard to detect. Message passing causes less nonblocking bugs than shared memory synchronization and surprisingly, was even used to fix bugs that are caused by wrong
       shared memory synchronization. We believe that message
       passing offers a clean form of inter-thread communication
       and can be useful in passing data and signals. But they are
       only useful if used correctly, which requires programmers
       to not only understand message passing  mechanisms well
       but also other synchronization mechanisms of Go.

    [Understanding Real-World Concurrency Bugs in Go](https://songlh.github.io/paper/go-study.pdf)

2. Don't make huge buffered channels. Channel is just a [data buffer](https://en.wikipedia.org/wiki/Data_buffer),
don't try to feet all results there.

3. The channel consumer should be spawned before channel producer(writer) and write values into DB/cache/file/socket/map/slice/other data structures. 

4. Channel should be closed once either by the producer(if it's one)
or with the help of `sync.WaitGroup`/`sync.Once`(if there are many producers).

5. Do not write into closed channel. However, reading from closed channel is OK.

6. Do not send to or receive from a nil channel it will block forever.
