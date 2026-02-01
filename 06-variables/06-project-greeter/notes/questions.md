# Command-Line Arguments Notes

## os.Args Variable

The `os.Args` variable stores command-line arguments passed to your program.

### Type

```go
var Args []string  // A slice of strings
```

`os.Args` is a slice of strings, not a string array or single string.

### First Item Special Meaning

The first item (`os.Args[0]`) contains the path to the running program, not the first user argument:

```bash
./myprogram hello world
```

- `os.Args[0]` → `./myprogram` (program path)
- `os.Args[1]` → `hello` (first argument)
- `os.Args[2]` → `world` (second argument)

## Working with os.Args

### Accessing Elements

Use bracket notation with zero-based indexing:

```go
first := os.Args[0]   // Program path
second := os.Args[1]  // First actual argument
```

Incorrect attempts:
- `Args.0` - not valid syntax
- `Args{1}` - wrong brackets
- `Args(1)` - not a function call

### Getting the Length

Use the built-in `len()` function:

```go
numArgs := len(os.Args)
```

Incorrect attempts:
- `length(Args)` - no such function
- `Args.len` - not an object property
- `Args.Length` - not an object property

### Element Type

Each element in `os.Args` is a `string`:

```go
var Args []string

// Each element:
arg := Args[0]  // type: string
```

## Getting User Arguments

To get the first actual argument (not the program path):

```go
firstArg := os.Args[1]  // First user-provided argument
```

Index guide:
- `os.Args[0]` → Program path
- `os.Args[1]` → First user argument
- `os.Args[2]` → Second user argument
- And so on...

## Example Usage

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Check if arguments were provided
    if len(os.Args) < 2 {
        fmt.Println("Please provide a name")
        return
    }

    name := os.Args[1]
    fmt.Println("Hello,", name)
}
```

