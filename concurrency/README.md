# Concurrency

## Before spawning goroutine(checklist):

1. Make sure, that there would be infinite amount of running goroutines use one of next:
- [semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)
- [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [worker pool](https://gobyexample.com/worker-pools)
- etc.

    The semaphore is preferable in most cases, since with the pool there would be hanging workers, that doing nothing. 

2. Check that it's possible to stop spawned goroutines.
- use `context.Context` for cancellation

3. **Panic.** Recover could catch panic only in current goroutine, so make sure, that [panic is handled in goroutine](https://medium.com/codex/handle-panic-in-go-routine-54b82d6013d3).