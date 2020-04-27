


### Sync.Once

In `main.go:main()` one of the first things we do is initialise a registry of connected monitors: this is a fancy way of saying a `map` which uses a UUID as a key, and a `Monitor` struct as the value. Obviously if this initialisation function was to be called twice then we'd have a problem: the state of all connected monitors would be lost, and our service would cease to function!

```go
var initialiser sync.Once

func InitialiseRegistry() {
	initialiser.Do(func(){
		mtx = &sync.Mutex{}
		monitors = make(map[string]Monitor)
	})
}
```

That's why `monitor/monitor.go:InitialiseRegistry()` utilises a `sync.Once`: no matter how many times `InitialiseRegistry` is called, it will only ever run the initialising code *once*.

### Mutex

In `monitor/monitor.go:Run(context.Context, Monitor)` you'll see the following:

```go
	mtx.Lock()
	monitors[mon.uuid] = mon
	mtx.Unlock()
```

Here we utilise a `sync.Mutex` to ensure that there are no concurrent writes to our registry.

### Context

A common issue when controlling concurrent tasks is that of *cancellation*: how do we communicate with a concurrent task that it should stop? We *could* use channels, but that can become quite confusing - especially if you need to *cascade* a cancellation through the call stack. 

The Go `context` package provides mechanisms to 

```go
	for {
		select {
		case update := <-mon.Update:
			if err := send(mon, update); err != nil {
				return
			}
		case <-ctx.Done():
			return
		}
	}
```

Further down `monitor/monitor.go:Run(context.Context, Monitor)` we utilise