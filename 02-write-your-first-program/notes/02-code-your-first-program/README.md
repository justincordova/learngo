## Package keyword

The `package` keyword defines which package a Go file belongs to.

```go
package main

func main() {
}
```

## package main

`package main` creates an executable Go program. This is required for programs that can be run directly.

## func main

`func main` is the entry point of an executable program. Go automatically calls this function when the program runs.

## import statement

The `import` statement brings in external packages so you can use their functionality.

```go
package main
import "fmt"

func main() {
    fmt.Println("Hi!")
}
```

This imports the `fmt` package, allowing you to use `fmt.Println` to print messages.

## func keyword

The `func` keyword declares a new function.

## What is a function?

A function is a reusable, executable block of code — like a mini-program within your program.

## Calling functions

**Main function:** Go calls `main()` automatically — you never call it yourself.

**Other functions:** You must call them explicitly. Go doesn't execute functions automatically (except `main` and some special initialization functions).

## Example: Empty program

```go
package main

func main() {
}
```

This is a valid, executable Go program, but it doesn't print anything since there's no `fmt.Println` call.

## Example: Incorrect program

```go
package main

func main() {
    fmt.Println(Hi! I want to be a Gopher!)
}
```

This is incorrect because:
- The `fmt` package is not imported
- The message is not wrapped in double quotes

## Example: Correct program

```go
package main
import "fmt"

func main() {
    fmt.Println("Hi there!")
}
```

Prints: `Hi there!`
