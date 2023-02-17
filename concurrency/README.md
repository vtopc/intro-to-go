# Concurrency

## Goroutines

Every function spawned with `go` keyword creates new goroutine, e.g.:
```go
go fn(x, y, z)
```
[Check a tour of Go](https://go.dev/tour/concurrency/1)

### Before spawning goroutine(checklist):

1. Make sure, that there would be infinite amount of running goroutines use one of next:
- use [semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)
- use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) with [SetLimit()](https://pkg.go.dev/golang.org/x/sync/errgroup#Group.SetLimit)
- create [worker pool](https://gobyexample.com/worker-pools)
- etc.

    The semaphore/errgroup is preferable in most cases, since with the pool there would be hanging workers, that doing nothing and spawning a new goroutine is quite cheap.

2. Check that it's possible to stop spawned goroutines, e.g. on service shutdown:
- use `context.Context` for cancellation

3. **Panic.** Recover could catch panic only in current goroutine, so make sure, that [panic is handled in goroutine](https://medium.com/codex/handle-panic-in-go-routine-54b82d6013d3).
