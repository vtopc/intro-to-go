# Concurrency

## Goroutines

Every function spawned with `go` keyword creates new goroutine, e.g.:
```go
go fn(a, b, c)
```
[Check a tour of Go](https://go.dev/tour/concurrency/1)

### Before spawning a new goroutine(the checklist):

0. Avoid Concurrency.

    Concurrency in Go is not a silver bullet, it wouldn't make code faster in all cases.

    > Async code harder to understand and debug is. Hrrmmm.

    Master Yoda

    e.g. there is no speed benefit in [this example](examples/2-workerpool/README.md) for **one** CPU,
    and async solution consumes [50%](examples/2-workerpool/README.md#benchmarks)–[100%](examples/3.1-group/README.md#benchmarks) more memory.

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
   
    Also, [Leave concurrency to the caller](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_leave_concurrency_to_the_caller)

1. Make sure, that there would be a limited amount of running goroutines. Use one of next:
   - use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) with [SetLimit()](https://pkg.go.dev/golang.org/x/sync/errgroup#Group.SetLimit)
   - use [semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore) or just some `make(chan struct{}, N)` before spawning a new goroutine. 
   - use pool from [conc](https://github.com/sourcegraph/conc) with `WithMaxGoroutines`
   - create [worker pool](https://gobyexample.com/worker-pools). Shouldn't be used in most cases since there would be hanging workers, that are doing nothing, and spawning a new goroutine is quite cheap.
   - etc., e.g. HTTP server has connections limit.

1. It should be possible to stop running goroutines, e.g. on service shutdown or HTTP timeout. Use one of next for cancellation:
    - `context.Context`(Context is cancelled or the deadline is exceeded. Preferable.)
    - stop channel(kinda legacy solution and hard to use with other APIs, e.g. with some DB ORM or HTTP clients)

    Also write [context aware code](https://www.storj.io/blog/production-concurrency).

   HTTP timeout example:
    ```go
    func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context() // http.Server should have timeouts
    
        var entity Entity
        _ = json.NewDecoder(r.Body).Decode(&entity) // Don’t forget to check the error and close the body
    
        group, _ := errgroup.WithContext(ctx)
        group.Go(func() (err error) {
            return h.db1.CreateEntity(ctx, entity) // a lot of methods takes context, not stop chan 
        })
        group.Go(func() (err error) {
            return h.db2.TransactionLog(ctx, entity)
        })
    
        err := group.Wait() // blocking
        if err != nil {
            http.Error(w, "db error", http.StatusInternalServerError)
            return
        }
    
        w.WriteHeader(http.StatusNoContent)
    }
    ```

1. [Never start a goroutine without knowing when it will stop](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_never_start_a_goroutine_without_knowing_when_it_will_stop).

   e.g. on shutdown app should wait for all goroutines to stop.

   Use one of next:
   - `sync.WaitGroup`
   - `context.Context` - returned by goroutine spawner, e.g. [errgroup.WithContext](https://pkg.go.dev/golang.org/x/sync/errgroup#WithContext)
   - done channel - [example](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_never_start_a_goroutine_without_knowing_when_it_will_stop)

   See also: [when you spawn goroutines, make it clear when or whether they exit.](https://google.github.io/styleguide/go/decisions#goroutine-lifetimes)

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
   
    Also, [panic could block goroutine forever](https://play.golang.com/p/FCvgacdAEuw), 
so call `mu.Unlock`, `wg.Wait` or `close(ch)` in defers.

1. `wg.Add(...)` should be called before running goroutine.

#### To sum up. The easiest ways.
- either use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) with panic recovery, [e.g.](examples/3.1-group/main.go);

- or use [sourcegraph/conc.WaitGroup](https://pkg.go.dev/github.com/sourcegraph/conc@v0.3.0#WaitGroup.WaitAndRecover) with semaphore.

### Tips and tricks

- Add [profiler labels](https://rakyll.org/profiler-labels/), this would help to debug and read stacktrace.
- Use the race detector (`-race` flag) and `t.Parallel()` in unit tests and subtests.
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

1. [Do not send into closed channel](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_a_send_to_a_closed_channel_panics). However, reading from closed channel is OK.

1. [Do not send to or receive from a nil channel it will block forever.](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_channel_axioms)

1. Don't make huge buffered channels. Channel is just a [data buffer](https://en.wikipedia.org/wiki/Data_buffer),
don't try to feet all results there(you would either make it too small and block on writing, or 
make it too big and use redundant memory).
[Prefer channels with a size of zero or one. Most other sizes are guesses.](https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_prefer_channels_with_a_size_of_zero_or_one)

1. Channel consumer should write values into DB/cache/file/socket/map/slice/other data structures.

1. If channel producer(writer) is not running in goroutine the channel consumer should be spawned before it.

1. Channel should be closed once either by the producer(if it's one)
or with the help of `sync.WaitGroup`/`sync.Once`(if there are many producers).

## More

- Egon Elbre - Production Ready Concurrency([read](https://www.storj.io/blog/production-concurrency) and [watch](https://www.youtube.com/watch?v=qq3gu0JQ0yU))
