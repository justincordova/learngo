## Statements

A statement instructs Go to do something. Statements can change the execution flow of your program.

## Execution direction

Go executes code from top to bottom, one statement at a time.

## Expressions

An expression produces a value. Unlike statements, expressions cannot change execution flow or stand alone.

## Operators

Operators combine expressions together to create more complex expressions.

## Expressions need statements

The following doesn't work because expressions can't stand alone:

```go
package main
import "fmt"

func main() {
    "Hello"  // ERROR: expression without statement
}
```

`"Hello"` is an expression that must be part of a statement (like `fmt.Println("Hello")`).

## Semicolons and multiple statements per line

```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.NumCPU()); fmt.Println("cpus"); fmt.Println("the machine")
}
```

This works because Go automatically adds semicolons behind the scenes for every statement. The statements are treated as if they're on separate lines.

## Combining expressions with operators

```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.NumCPU() + 10)
}
```

This works because the `+` operator combines two expressions: `runtime.NumCPU()` and `10`.
