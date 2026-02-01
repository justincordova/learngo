# Short If Statement Notes

## Short If Syntax

A short if statement allows you to declare variables in the if statement itself:

```go
if simpleStatement; condition {
    // code
}
```

Structure:
1. Simple statement (usually a short declaration)
2. Semicolon separator
3. Condition expression
4. Code block

## Syntax Order

The simple statement must come BEFORE the condition:

```go
// Correct:
if d, err := time.ParseDuration("1h10s"); err != nil {
    fmt.Println(d)
}

// Wrong - condition before statement:
if err != nil; d, err := time.ParseDuration("1h10s") {
    fmt.Println(d)
}
```

The semicolon separates the statement from the condition, but the statement must be first.

## Variable Shadowing

Variables declared in a short if shadow outer variables with the same name:

```go
done := false                    // Outer scope
if done := true; done {          // Shadows outer 'done'
    fmt.Println(done)            // Prints: true (inner done)
}
fmt.Println(done)                // Prints: false (outer done)
```

Output: `true` then `false`

Explanation:
- Inner `done` (true) only exists inside the if statement
- Outer `done` (false) is used outside the if statement
- They are different variables with the same name

## Common Shadowing Issue

Shadowing can cause unexpected behavior when you want to use a value outside the if:

```go
var n int  // Declare n in outer scope

if n, err := strconv.Atoi("10"); err != nil {  // Shadows n!
    fmt.Printf("error: %s (n: %d)", err, n)
    return
}

fmt.Println(n)  // Prints: 0 (not 10!)
```

Problem: The short if declares a new `n`, shadowing the outer `n`. The converted value stays in the inner scope.

### Fix: Declare Error Variable First

```go
var (
    n   int    // Declare n
    err error  // Declare err
)

if n, err = strconv.Atoi("10"); err != nil {  // Assignment, not declaration
    fmt.Printf("error: %s (n: %d)", err, n)
    return
}

fmt.Println(n)  // Prints: 10 (correct!)
```

Changes:
- Declare both `n` and `err` in outer scope
- Use `=` (assignment) instead of `:=` (declaration) in the if
- Now `n` is assigned in outer scope, not shadowed

Why other approaches don't work:
- Removing outer `n` declaration - Code after if needs it
- Removing inner `n` declaration - Code needs to set n
- Declaring err outside main - Would be unused (error)

## When to Use Short If

Short if is useful when:
- Variable is only needed for the condition check
- You want to limit variable scope
- Combining declaration and checking in one line

```go
if d, err := time.ParseDuration("1h10s"); err == nil {
    fmt.Println(d)  // d only used here
}
// d not accessible here
```

## When to Avoid Short If

Avoid short if when:
- You need the variable after the if statement
- Shadowing would cause confusion
- The variable already exists in outer scope

Better approach for these cases:

```go
d, err := time.ParseDuration("1h10s")
if err != nil {
    // handle error
}
fmt.Println(d)  // d still accessible
```

## Key Points

1. Short if syntax: `if statement; condition { }`
2. Statement comes before condition, separated by semicolon
3. Variables declared in short if shadow outer variables
4. Shadowed variables only exist inside the if block
5. To avoid shadowing issues, declare variables outside and use `=` instead of `:=`
6. Use short if when variable scope should be limited to if block
7. Avoid short if when you need the variable afterward
8. Common pattern: `if value, err := function(); err != nil { }`
