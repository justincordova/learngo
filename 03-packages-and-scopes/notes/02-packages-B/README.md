## Package types in Go

Go has two types of packages:

1. **Executable packages** - Can be run directly
2. **Library packages** - Provide reusable code for import

## Executable packages

An executable package requires:
- `package main`
- `func main`

Only executable packages can be run with `go run`.

**Example:**
```go
package main

func main() {
    // entry point
}
```

## Library packages

Any package that's not `package main` is a library package. These provide reusable functionality that can be imported by other packages.

**Example:**
```go
package helpers

func Calculate() int {
    return 42
}
```

## go build

`go build` can compile both executable and library packages.

## go run

`go run` can only execute executable packages (`package main` with `func main`).
