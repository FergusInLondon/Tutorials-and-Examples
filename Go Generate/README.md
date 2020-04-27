# Go Generate: Example Usage

This is meant as an illustration for a blog post I wrote - here. 

## Usage

By executing `go generate` prior to `go build`, `./generators/status_gen.go` will get built - parsing the `statuses.yaml` file, and building another source file from the contents.

```
➜  generate-example git:(master) ls
    go.mod  go.sum  main.go  statuses.yaml  status.go

➜  generate-example git:(master) go generate 

➜  generate-example git:(master) ls
    go.mod  go.sum  main.go  statuses.go  statuses.yaml  status.go

➜  generate-example git:(master) go build

➜  generate-example git:(master) ./genexample 1
    Received Status '1'.
    Message: 'Operation Accepted'.

➜  generate-example git:(master) ./genexample 2
    Received Status '2'.
    Message: 'Operation Queued'.
```