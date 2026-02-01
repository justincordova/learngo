## Difference between `go build` and `go run`

**`go build`**: Compiles your program and creates an executable binary in the current directory.

**`go run`**: Compiles your program to a temporary directory and immediately runs it. The compiled binary is discarded after execution.

## Where compiled code is saved

When you run `go build`, the executable is created in the same directory where you ran the command.

When you run `go install`, the binary is installed to `$GOBIN` or `$GOPATH/bin` for global access.

## Runtime vs Compile-time

**Compile-time**: When your program is being compiled into machine code. No program execution happens here.

**Runtime**: When your compiled program actually runs on the computer. This is when your program can interact with the system, print messages, read files, etc.

## When programs print to console

A Go program can only print messages during runtime, after it has been compiled. During compilation, the program is "dead" — it cannot execute any logic or interact with the system.