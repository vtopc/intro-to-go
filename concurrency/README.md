# Concurrency

## Goroutines

Every function spawned with `go` keyword creates new goroutine, e.g.:
```go
go fn(x, y, z)
```
[Check a tour of Go](https://go.dev/tour/concurrency/1)

### Before spawning a new goroutine(the checklist):

0. Avoid Concurrency.

    Concurrency in Go is not a silver bullet, it wouldn't make code faster in all cases.

    > Async code harder to understand and debug is. Hrrmmm.

    Master Yoda

    e.g. there is no speed benefit in [this example](examples/workerpool/README.md) for **one** CPU,
    and async solution consumes [50%](examples/workerpool/README.md#benchmarks)–[100%](examples/errgroup/README.md#benchmarks) more memory.

    Another example:
    ```go
    var wg sync.WaitGroup

    wg.Add(1)
    go serve(&wg)
    wg.Wait()
    ```
   
    vs.:
    ```go
    serve()
    ```

1. Make sure, that there would be a limited amount of running goroutines. Use one of next:
   - use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) with [SetLimit()](https://pkg.go.dev/golang.org/x/sync/errgroup#Group.SetLimit)
   - use [semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore) or just some `make(chan struct{}, N)` before spawning a new goroutine. 
   - use pool from [conc](https://github.com/sourcegraph/conc) with `WithMaxGoroutines`
   - create [worker pool](https://gobyexample.com/worker-pools). Shouldn't be used in most cases since there would be hanging workers, that are doing nothing and spawning a new goroutine is quite cheap.
   - etc.

1. It should be possible to stop running goroutines, e.g. on service shutdown or HTTP timeout. Use one of next for cancellation:
    - `context.Context`(preferable)
    - ~~done channel~~(kinda legacy and hard to use with APIs)

    Also write context aware code.

1. [Never start a goroutine without knowing when it will stop](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_never_start_a_goroutine_without_knowing_when_it_will_stop).
e.g. on shutdown app should wait for all goroutines to stop.
Use one of next:
   - `sync.WaitGroup`
   - `context.Context` returned by goroutine spawner, e.g. [errgroup.WithContext](https://pkg.go.dev/golang.org/x/sync/errgroup#WithContext)
   - ~~done channel~~(could wait only for one goroutine to stop)

1. **Panics.** Recover could catch panic only in current goroutine, so make sure, that [panic is handled in goroutine](https://medium.com/codex/handle-panic-in-go-routine-54b82d6013d3).

   [Non-catchable example](https://play.golang.com/p/lVfDUZTz4ji):
    ```go
    func f1() {
        if true {
            panic("f1")
        }
    
        fmt.Println("f1 done")
    }
    
    func main() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("recovered panic:", r)
            }
        }()
    
        go f1()
    
        for i := 0; i < 10; i++ {
            fmt.Println("tick", i)
            time.Sleep(time.Second)
        }
    }
    ```

1. `wg.Add(...)` should be called before spawning goroutine.

#### To sum up. The easiest ways.
- either use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) with panic recovery, [e.g.](examples/errgroup/main.go);

- or use [sourcegraph/conc.WaitGroup](https://pkg.go.dev/github.com/sourcegraph/conc@v0.3.0#WaitGroup.WaitAndRecover) with semaphore.

### Tips and tricks

- Add [profiler labels](https://rakyll.org/profiler-labels/), this would help to debug and read stacktrace.
- Use the race detector (`-race` flag) and `t.Parallel()` in unit tests.
- Use [uber-go/goleak](https://github.com/uber-go/goleak) to detect goroutine leaks in your tests.
- You can [update pre-allocated slice concurrently](https://stackoverflow.com/questions/49879322/can-i-concurrently-write-different-slice-elements).

## Channels

Channels are a typed conduit through which you can send and receive messages.

[Check a tour of Go](https://go.dev/tour/concurrency/2)

### Commandments

1. Consider using shared memory instead of message passing:

    >  Our study found that message passing does not necessarily make 
       multithreaded programs less error-prone than shared memory.
       In fact, message passing is the main cause of blocking bugs.
       To make it worse, when combined with traditional synchronization 
       primitives or with other new language features
       and libraries, message passing can cause blocking bugs that
       are very hard to detect. Message passing causes less nonblocking bugs 
       than shared memory synchronization and surprisingly, 
       was even used to fix bugs that are caused by wrong
       shared memory synchronization. We believe that message
       passing offers a clean form of inter-thread communication
       and can be useful in passing data and signals. But they are
       only useful if used correctly, which requires programmers
       to not only understand message passing  mechanisms well
       but also other synchronization mechanisms of Go.

    [Understanding Real-World Concurrency Bugs in Go](https://songlh.github.io/paper/go-study.pdf)

1. Don't make huge buffered channels. Channel is just a [data buffer](https://en.wikipedia.org/wiki/Data_buffer),
don't try to feet all results there.

1. Channel consumer should write values into DB/cache/file/socket/map/slice/other data structures.

1. If channel producer(writer) is not spawned in goroutine the channel consumer should be spawned before it.

1. Channel should be closed once either by the producer(if it's one)
or with the help of `sync.WaitGroup`/`sync.Once`(if there are many producers).

1. Do not write into closed channel. However, reading from closed channel is OK.

1. Do not send to or receive from a nil channel it will block forever.

## More

- Egon Elbre - Production Ready Concurrency([read](https://www.storj.io/blog/production-concurrency) and [watch](https://www.youtube.com/watch?v=qq3gu0JQ0yU))
