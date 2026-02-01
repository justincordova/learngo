# Input Scanning

## What is an Input Stream?

An **input stream** is any data source that generates contiguous data. The most common example is **Standard Input (stdin)**, but input streams can come from many sources:

- User keyboard input
- File contents
- Network connections
- Website contents
- Output from other programs (via pipes)

Standard Input can be redirected to almost any data source, making it a flexible interface for reading data into your program.

## Using bufio.Scanner

The `bufio.Scanner` type provides a convenient way to read input streams line by line (or token by token):

```go
in := bufio.NewScanner(os.Stdin)
in.Scan()  // user enters: "hi!"
in.Scan()  // user enters: "how are you?"
fmt.Println(in.Text())
// Prints: "how are you?"
```

**Important:** The `Text()` method returns only the **last scanned token**. Each call to `Scan()` advances to the next token and overwrites the previous one.

To process all input, you typically use a loop:

```go
in := bufio.NewScanner(os.Stdin)
for in.Scan() {
    fmt.Println(in.Text())  // Process each line
}
```

## Error Handling with Scanner

Use the `Err()` method to detect errors that occurred during scanning:

```go
in := bufio.NewScanner(os.Stdin)
for in.Scan() {
    fmt.Println(in.Text())
}

if err := in.Err(); err != nil {
    fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
}
```

The `Err()` method returns `nil` if no errors occurred, or an error value describing what went wrong.

## Configuring Split Functions

By default, `bufio.Scanner` splits input into lines. You can change this behavior using the `Split()` method:

```go
in := bufio.NewScanner(os.Stdin)
in.Split(bufio.ScanWords)  // Split by words instead of lines
```

**Available split functions:**
- `bufio.ScanLines` - Split by newlines (default)
- `bufio.ScanWords` - Split by whitespace-separated words
- `bufio.ScanRunes` - Split by individual Unicode characters
- `bufio.ScanBytes` - Split by individual bytes

You can also write custom split functions for more complex parsing needs.

## The "Must" Naming Convention

Functions or methods with the "Must" prefix (like `regexp.MustCompile`) indicate that they **may panic** (crash your program) if they encounter an error:

```go
re := regexp.MustCompile("...")  // Panics if the regex is invalid
```

**Convention meaning:**
- `Must` prefix = function may panic on error
- Use these when you're certain the input is valid (e.g., hard-coded patterns)
- For user input or uncertain data, use the non-Must variant that returns an error instead:

```go
// Safe version that returns an error
re, err := regexp.Compile(userPattern)
if err != nil {
    fmt.Println("Invalid regex:", err)
    return
}
```

The "Must" prefix serves as a clear warning that the function won't return an error—it will panic instead.
