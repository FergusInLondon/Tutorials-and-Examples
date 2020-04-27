
So far we've explored mechanisms for:

- reducing global state (i.e *clojures*);
- reducing race conditions and controlling concurrent access (i.e the humble *mutex*)
- implementing singular use actions (i.e Go's *sync.Once*);
- communicating across routines without sharing memory (i.e via *channels*);

Now we're going to look at how we can break concurrent operations down in to smaller chunks via "*pipelining*". Once again, this is going to be very heavy with `Go`; as with the rest of these posts, this will be our language of choice. Where possible though, I'll talk about alternatives in languages like Elixir/Erlang, and Rust. (It's also worth noting that *Part 5* is heavily focused around exploring languages other than Go, their idioms, and what we can learn from them.)

## What is a pipeline?

The official Go site has a pretty informative post about *what* a pipeline is; but to summarise, *it's a series of steps - in the form of functions - that are chained together, most often via channels*. Here's an example, below we next 5 function calls:

```go
    output := step5(step4(step3(step2(step1(input)))))
```

Needless to say, that's a little... erhhh... "*scruffy*"; and in a code review I'd be the first person to put an obligatory "*No*" underneath it!

When dealing with input streams - quite a common scenario in concurrent environments - this can be expressed in a much more readable way:

```go
    toStage2 := make(chan []byte)
    toStage3 := make(chan []byte)
    toStage4 := make(chan []byte)
    toStage5 := make(chan []byte)
    output := make(chan []byte)

    // Each step has roughly the same function signature: fn(in, out chan)
    go step1(input, toStage2)
    go step2(toStage2, toStage3)
    go step3(toStage3, toStage4)
    go step4(toStage4, toStage5)
    go step5(toStage5, output)

    out := <-output:
    /* do whatever */
```

With real world function names (i.e `verifySignature()`, `verifyPayload()`, `populateAttributes()`, `persistRecord()`...) it becomes even *more* easy to digest. 

### "That sure looks like a lot of goroutines though..."

At first glance it may seem quite wasteful to have 5 goroutines constantly waiting on input; although this doesn't have to be a performance concern. The overhead of a goroutine by itself isn't a concern; but the possibility of leaking goroutines is a valid concern. This is where context cancellation comes in to play.


```go
    toStage2 := make(chan []byte)
    toStage3 := make(chan []byte)
    toStage4 := make(chan []byte)
    toStage5 := make(chan []byte)
    output := make(chan []byte)

    ctx, cancel := context.WithCancel(context.Background())

    // Each step has roughly the same function signature: fn(context.Context, chan, chan)
    // The context is used in each stage, if they receive a `ctx.Done()` then they return
    go step1(ctx, input, toStage2)
    go step2(ctx, toStage2, toStage3)
    go step3(ctx, toStage3, toStage4)
    go step4(ctx, toStage4, toStage5)
    go step5(ctx, toStage5, output)

    select {
        case <- time.After(5 * time.Second):
            cancel()
        case out := <-output:
        /* do whatever */
    }
```

Here we wait for the first event to occur: receiving output or receiving a timeout. In the case of a timeout then we cancel the context.

There's still a problem here though; our above solution still assumes we want to handle the output in some way: often we don't. We want to process the stream and be done with it. To handle this, we can wait on the goroutines using `sync.WaitGroup` to similar (blocking) effect.

Using `sync.WaitGroup` *in conjunction* with `context.Context` and a timeout isn't quite as intuitive as it could be though:

```
    toStage2 := make(chan []byte)
    toStage3 := make(chan []byte)
    toStage4 := make(chan []byte)
    toStage5 := make(chan []byte)

    ctx, cancel := context.WithCancel(context.Background())
    wg := sync.WaitGroup{}

    // Each step has roughly the same function signature: fn(context.Context, *sync.WaitGroup, in, out chan)
    // Each step also `defer`s a call to `wg.Done()`; ensuring `wg.Wait()` blocks until all are finished.
    wg.Add(5)
    go step1(ctx, &wg, input, toStage2)
    go step2(ctx, &wg, toStage2, toStage3)
    go step3(ctx, &wg, toStage3, toStage4)
    go step4(ctx, &wg, toStage4, toStage5)
    go step5(ctx, &wg, toStage5)

    done := make(chan struct{})
    go func(){
        wg.Wait()
        close(done)
    }()

    select {
        case <-time.After(5 * time.Second):
            cancel()
            /* timeout has occured */
        case <-done:
            /* pipeline successfully executed */
    }
```

Both `context.Context` and `sync.WaitGroup` are incredibly useful when dealing with goroutines; and can be used to provide a lot of control over concurrent operations - be they used individually or in conjunction with eachother.

## Use cases

There are two scenarios in which pipelines can be *very* useful:

1. **Sequential processing of streamed data**; consider a chat application, whereby there's a continuous stream of messages that need to be (a) parsed for blacklisted/censored words, (b) have meta-data populated (i.e updated with a timestamp and user identifier), (c) persisted to a data store, and (d) distributed to other users of that chat room. A pipeline makes perfect sense, as each of the above tasks can be delegated to a specific stage of the pipeline.
2. **Proxying data between services**; suppose you're "*gluing*" two services together, and you need to do some basic data conversions before decorating the input so that the destination service is able to consume the input. You may need (a) accept input from one interface, (b) convert it's structure, (c) update the new structure with metadata, and (d) dispatch to the destination via an API adapter. Like the above example, breaking down each task in to a specific stage will increase the readability of the code, and ensure isolation between tasks.

### An example

In *Part 6* we're going to look at a very rudimentary example that ties together all of the concepts discussed in this series; we're going to build a toy Security Information and Event Management (SIEM) system. It will expose a `POST` API endpoint, and distribute inbound events to any clients connected via a websocket interface.

So breaking the requirements down we will need to:

1. validate that the reporting account is permited to access the API, populate the report with account details;
2. populate the event with metadata - i.e time and status;
3. publish alerts via sms/email if the severity necessitates it;
4. persist the alert to a database;
5. distribute the alert to connected clients.

In `Part 4/pipeline.go` there's a toy implementation of this; with things like user management and database storage and event distribution stubbed out.
