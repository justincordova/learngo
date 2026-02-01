# If Statement Notes

## Control Flow

Control flow means deciding which statements are executed based on conditions.

- You cannot change top-to-bottom execution order
- You cannot change left-to-right expression evaluation
- You CAN decide which code blocks run based on boolean conditions

```go
if condition {
    // This code runs if condition is true
}
```

## If Statement Syntax

### Parentheses are Optional

In Go, you don't need parentheses around conditions:

```go
// Both work, but Go style omits parentheses:
if (mood == "perfect") {  // Works but unnecessary
    // code
}

if mood == "perfect" {    // Preferred Go style
    // code
}
```

Invalid syntax:
- `if {mood == "perfect"}` - Wrong brackets
- `if [mood == "perfect"]` - Wrong brackets
- `if mood = "perfect"` - Wrong operator (= is assignment, not comparison)

### Condition Must Be Boolean

The condition must evaluate to a `bool` value:

```go
mood := "happy"

// Correct - comparison yields bool:
if mood == "happy" {
    fmt.Println("cool")
}

// Wrong - string is not bool:
if "happy" {  // ERROR
    fmt.Println("cool")
}

// Wrong - variable is string, not bool:
if mood {  // ERROR
    fmt.Println("cool")
}
```

## Simplifying Boolean Conditions

### Redundant Comparisons to true/false

When a variable is already boolean, don't compare it to `true` or `false`:

```go
happy := true

// Unnecessary:
if happy == true {
    fmt.Println("cool!")
}

// Simplified:
if happy {
    fmt.Println("cool!")
}
```

For negative conditions:

```go
happy := false

// Unnecessary:
if happy == !true {
    fmt.Println("why not?")
}

// Also unnecessary:
if happy == false {
    fmt.Println("why not?")
}

// Simplified:
if !happy {
    fmt.Println("why not?")
}
```

## Else and Else If

### Structure Rules

An if statement can have multiple branches:

```go
if condition1 {
    // Branch 1
} else if condition2 {
    // Branch 2
} else {
    // Default branch
}
```

Rules:
- Only ONE `else` branch allowed
- `else` must be the LAST branch
- Multiple `else if` branches are allowed
- `else` and `else if` are optional

### Common Errors

```go
// ERROR - Two else branches:
if happy {
    fmt.Println("cool!")
} else if !happy {
    fmt.Println("why not?")
} else {
    fmt.Println("why not?")
} else {  // ERROR - duplicate else
    fmt.Println("why not?")
}
```

Fix: Remove one of the `else` branches.

### Unreachable Code

```go
happy := true
energic := happy

if happy {
    fmt.Println("cool!")
} else if !happy {
    fmt.Println("why not?")
} else if energic {  // NEVER EXECUTES
    fmt.Println("working out?")
}
```

Problem: If `happy` is `true`, first branch executes. If `happy` is `false`, second branch executes. The third branch never runs because all cases are handled.

## Simplifying If Chains

When `else` handles all remaining cases, `else if` may be redundant:

```go
// Before simplification:
happy := false

if happy {
    fmt.Println("cool!")
} else if happy != true {  // Redundant
    fmt.Println("why not?")
} else {
    fmt.Println("why not?")
}

// Simplified - remove else if:
if happy {
    fmt.Println("cool!")
} else {  // Handles all "not happy" cases
    fmt.Println("why not?")
}
```

The `else` branch already handles the case where `happy` is `false`, so the `else if happy != true` is unnecessary.

## Key Points

1. Control flow decides which code executes based on conditions
2. Parentheses around conditions are optional (and not idiomatic)
3. Condition must be `bool` type or evaluate to `bool`
4. Don't compare bool variables to `true` or `false` - use them directly
5. Use `!variable` instead of `variable == false`
6. Only one `else` branch allowed per if statement
7. `else` must be the last branch
8. Multiple `else if` branches are allowed
9. Watch for unreachable code in if-else chains
10. Simplify by removing redundant conditions
11. `else` branch is optional
