## Why use comments?

Comments provide explanations for your code and enable automatic documentation generation.

## Comment syntax

**Single-line comments:** Start with `//`

```go
// This is a single-line comment
```

**Multi-line comments:** Enclosed in `/* ... */`

```go
/*
This is a multi-line comment
that spans multiple lines
*/
```

## Valid comment usage

```go
package main

// main function is an entry point /*
func main() {
    fmt.Println(/* this will print Hi! */ "Hi")
}
```

This works because:
- `//` comments skip the entire rest of the line
- `/* ... */` comments can be placed almost anywhere and are completely ignored by Go

## Invalid comment usage

```go
package main

func main() {
    fmt.Println(// "this will print Hi!")
}
```

This doesn't work because `//` skips everything after it on the line, including the closing `)` that Go needs.

## Automatic documentation

To generate documentation from your code, start comments with the name of the declared item:

```go
// Calculate returns the sum of two numbers
func Calculate(a, b int) int {
    return a + b
}
```

## go doc command

Use `go doc` to print documentation from the command line:

```bash
go doc PackageName
go doc PackageName.FunctionName
```

`go doc` uses the `godoc` tool behind the scenes — it's a simplified interface to `godoc`.
