# Assignment Notes

## How Assignment Works

When you assign a variable to another variable, the assignee variable's value is changed to the assigned variable's value. The variables remain separate:

```go
x := 10
y := 5
x = y  // x is now 5, but they are still separate variables
```

## Assignment vs Declaration

Assignment uses `=` to change the value of an existing variable. Declaration uses `:=` or `var` to create a new variable:

```go
opened := true   // Declaration
opened = false   // Assignment - changes existing variable
```

## Multiple Assignment

You can assign multiple variables at once:

```go
c, d = true, false
```

The number of variables on the left must match the number of values on the right.

## Type Compatibility

Assignments must respect type compatibility:

```go
var (
  n = 3        // int
  m int        // int
)

n = 10         // OK - assigning int to int
// n = true    // ERROR - can't assign bool to int
// m = "4"     // ERROR - can't assign string to int
```

## Complex Assignments

You can use expressions on the right side of assignments:

```go
var (
  n = 3
  m int
  f float64
)

n, m, f = m + n, n + 5, 0.5  // All evaluated, then assigned
```

## The Blank Identifier

Use `_` to discard values you don't need:

```go
var (
  a int
  c bool
)

_, _ = a, c  // Discard both values
```

## Functions Returning Multiple Values

When a function returns multiple values, you can assign them or discard unwanted ones:

```go
var c, f string

// path.Split returns two string values
_, f = path.Split("assets/profile.png")  // Keep only second value

// These would be errors:
// f = path.Split("assets/profile.png")            // Can't assign 2 values to 1 variable
// _, _, c = path.Split("assets/profile.png")      // Wrong number of values
// _ = path.Split("assets/profile.png")            // Can't assign 2 values to 1 variable
```
