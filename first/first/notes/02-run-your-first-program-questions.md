# Running Your First Program

## Go Build vs Go Run

Go provides two primary commands for working with your programs:

### `go build`
Compiles your program into an executable binary file but does not run it. The compiled binary is saved to disk and can be executed later.

```bash
go build main.go
# Creates an executable file (e.g., main or main.exe)
```

### `go run`
Compiles your program and immediately runs it. The compiled binary is stored in a temporary directory and automatically executed. The temporary binary is cleaned up after execution.

```bash
go run main.go
# Compiles and runs the program in one step
```

**Key difference:** `go run` is convenient for development and testing, while `go build` is used when you want to create a distributable executable.

## Runtime vs Compile-Time

Understanding these two stages is crucial for debugging and development:

### Compile-Time
The stage when your program is being compiled by the Go compiler. During this phase:
- The compiler checks your code for syntax errors
- Type checking occurs
- The source code is translated into machine code
- Your program is "dead"—it cannot execute or interact with the computer
- No output can be printed to the console

### Runtime
The stage when your compiled program is actually running on a computer. During this phase:
- Your program is alive and executing
- It can interact with the operating system
- It can print messages to the console
- It can read files, make network requests, etc.
- Errors that occur are called "runtime errors"

## When Can Your Program Print?

Your program can **only print messages during runtime**, after it has been successfully compiled.

```go
package main
import "fmt"

func main() {
    fmt.Println("This prints at runtime")
}
```

During compile-time, your program exists only as source code being processed by the compiler. It has no ability to execute instructions or produce output. Only after compilation completes and the program starts executing can it interact with the console or any other system resources.
