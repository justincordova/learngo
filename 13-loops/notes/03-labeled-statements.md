# Labeled Statements

## Label Scope

Labels belong to the **function scope**, not the statement scope where they are declared. This means:

- Labels can be used anywhere within the function
- Labels can even be referenced before their declaration in the function
- This behavior enables `goto` statements to jump to any label within a function

Unlike variables or constants that are scoped to their block, labels have function-wide scope.

## Which Statement Does a Label Label?

A label can only label **one statement** at a time—the statement immediately following it:

```go
for range words {
words:              // This label labels the nested loop, not the outer loop
    for range letters {
        // ...
    }
}
```

In this example, the `words` label labels the second (nested) loop, not the first loop.

## Break Without Labels

An unlabeled `break` terminates the closest enclosing loop or switch:

```go
package main

func main() {
main:
    for {
        switch "A" {
        case "A":
            break  // Only breaks the switch, not the loop
        case "B":
            continue main
        }
    }
}
```

The `break` statement terminates the switch, but the loop continues running indefinitely. Since the break doesn't reference the `main` label, it only affects the switch statement.

## Break With Labels

A labeled `break` can terminate an outer statement:

```go
package main

func main() {
    flag := "A"

main:
    for {
        switch flag {
        case "A":
            flag = "B"
            break          // Breaks the switch
        case "B":
            break main     // Breaks the loop
        }
    }
}
```

**Execution flow:**
1. First iteration: `flag` is "A", first case matches, sets `flag` to "B", breaks from switch
2. Loop continues to next iteration
3. Second iteration: `flag` is "B", second case matches, `break main` terminates the loop

The loop terminates because `break main` references the label on the loop.

## Nested Switches and Labels

Labels help you control which statement to break from in deeply nested code:

```go
package main

func main() {
    for {
    switcher:
        switch 1 {
        case 1:
            switch 2 {
            case 2:
                break switcher  // Breaks from the first (labeled) switch
            }
        }
        break  // This break terminates the loop
    }
}
```

**Execution flow:**
1. `break switcher` breaks from the first switch (the one labeled `switcher`)
2. Execution continues to the second `break`
3. The second `break` terminates the loop
4. Program ends

## Continue With Labels

The `continue` statement can **only** be used with loop labels, not switch labels:

```go
package main

func main() {
    for {
    switcher:
        switch {
        case true:
            switch {
            case false:
                continue switcher  // Error!
            }
        }
    }
}
```

**This fails** because `switcher` labels a switch statement, and `continue` requires a loop label. The `continue` statement is designed to skip to the next iteration of a loop—it has no meaning for switch statements.

## Goto Statements and Labels

The `goto` statement jumps execution to a labeled statement. Be careful with `goto`—it can create infinite loops:

### Infinite Loop Example 1

```go
func main() {
    start: goto exit
    exit : fmt.Println("exiting")
    goto start
}
```

- `goto exit` jumps to the exit label
- `goto start` jumps back to the start label
- Infinite loop between start and exit

### Infinite Loop Example 2

```go
func main() {
    exit: fmt.Println("exiting")
    goto exit
}
```

- The exit label prints "exiting"
- `goto exit` jumps back to the exit label
- Infinite loop printing "exiting"

### Terminating Example

```go
func main() {
    goto exit
    start : goto getout
    exit  : goto start
    getout: fmt.Println("exiting")
}
```

**Execution flow:**
1. `goto exit` jumps to the exit label
2. `goto start` jumps to the start label
3. `goto getout` jumps to the getout label
4. The getout label is the last statement, so the function ends
5. Program terminates successfully

**Note:** While `goto` is available in Go, it's rarely the best solution. Prefer structured control flow with loops, functions, and early returns.
