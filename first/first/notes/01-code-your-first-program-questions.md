# Coding Your First Program

## The `package` Keyword

The `package` keyword defines which package a Go file belongs to:

```go
package main

func main() {
}
```

Every Go file must start with a `package` declaration. This organizes your code and determines how it can be used.

## The `func` Keyword

The `func` keyword is used to declare a new function:

```go
func main() {
    // function body
}
```

A function is like a mini-program—a reusable and executable block of code that performs a specific task.

## The `main` Package

```go
package main

func main() {
}
```

The `main` package is special: it allows you to create an **executable Go program**. When you compile a program with `package main`, Go creates a binary that can be run directly.

Other package names (like `package utils` or `package models`) create libraries that can be imported by other programs, but cannot be executed on their own.

## The `main` Function

```go
package main

func main() {
}
```

The `main` function is the entry point of your program. Go automatically calls the `main` function when your program starts executing. You don't need to call it yourself.

**Important:** Go only automatically calls `func main`. All other functions must be called explicitly by your code.

## Calling Functions

While Go calls `main()` automatically, you must call all other functions yourself:

```go
func greet() {
    fmt.Println("Hello!")
}

func main() {
    greet()  // You must call greet() explicitly
}
```

If you don't call a function, it will never execute.

## The `import` Keyword

The `import` keyword brings external packages into your program so you can use their functionality:

```go
package main
import "fmt"

func main() {
    fmt.Println("Hi!")
}
```

`import "fmt"` imports the `fmt` package, which provides formatted I/O functions like `Println`. After importing, you can call functions from that package using the package name as a prefix (e.g., `fmt.Println`).

## A Minimal Program

The simplest valid Go program:

```go
package main

func main() {
}
```

This program is **correct** but doesn't do anything visible. It will compile and run successfully, but produces no output because there's no code inside `main()`.

## Common Mistakes

### Missing Double Quotes

```go
package main

func main() {
    fmt.Println(Hi! I want to be a Gopher!)  // Error!
}
```

**Problems:**
1. The message is not wrapped in double quotes
2. The `fmt` package is not imported

**Correct version:**
```go
package main
import "fmt"

func main() {
    fmt.Println("Hi! I want to be a Gopher!")
}
```

## A Complete Working Program

```go
package main
import "fmt"

func main() {
    fmt.Println("Hi there!")
}
```

**This program:**
1. Declares itself as part of the `main` package
2. Imports the `fmt` package for formatted output
3. Defines a `main` function that Go will call automatically
4. Prints "Hi there!" to the console

**Output:** `Hi there!`
