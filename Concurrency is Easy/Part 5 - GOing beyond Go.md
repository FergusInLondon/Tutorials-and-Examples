

So far we've only really looked at `Go`; this was a specific decision made as (a) Go is very easy to read, and (b) the simplicity of it's concurrency model is very hard to beat. That's not to say that other languages don't have pretty

### Rust

Rust boasts of "*fearless concurrency*"; and it manages this via a similar mechanism to Go: `channels`. Furthermore, the standard library also provides a mutex implementation.

Concurrency is made slightly more complex due to Rust' ownership model though.



### Erlang (& Elixir)