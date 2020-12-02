# Go Learnings

Here is a list of tips and tricks I have learnt while carrying out Advent of Code 2020 in golang.

## Error Handlings

It seems the norm to have functions declared like so

```go
func example() int, error {
  // body
}
```

where we then `return nil, errors.New("")` when an error occurs.

Functions should return an error instead of calling `panic`.

## File Reading

Due to how go runs, you cannot read files relative to a source file. For example, say we have the following file structure

```
project
└───folder1
    │   source.go
    │   input.txt
```

and you run `source.go` from the `project` folder with `go run folder1/source.go`, `source.go` cannot access `input.txt` relatively. It instead would need to access it as `folder1/input.txt`. This is because of how `golang` is compiled.
