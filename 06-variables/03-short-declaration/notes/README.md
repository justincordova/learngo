# Short Declaration Notes

## Syntax

Short declaration uses the `:=` operator to declare and initialize variables in one step:

```go
safe := true
speed := 50
```

The type is automatically inferred from the value on the right side.

## Multiple Variables

You can declare multiple variables at once:

```go
y, x, p := 5, "hi", 1.5
```

This declares:
- `y` as `int` with value `5`
- `x` as `string` with value `"hi"`
- `p` as `float64` with value `1.5`

The number of variables on the left must match the number of values on the right.

## Type Inference

Go automatically determines the type based on the assigned value:

```go
s := "hmm..."    // type: string
b := true        // type: bool
i := 42          // type: int
f := 6.28        // type: float64
```

## Equivalence with var

Short declaration is often equivalent to longer `var` declarations:

```go
// These are equivalent:
s := "hi"
var s string = "hi"

// These are equivalent:
n := 10
var n int = 10
```

## Redeclaration

When using short declaration with multiple variables, at least one variable must be new. You can redeclare an existing variable if you're also declaring a new one:

```go
y, x := false, 20  // x is 20
x, z := 10, "hi"   // x is now 10, z is new
```

In the second line, `x` is reassigned and `z` is declared.

## Scope Restrictions

Short declarations can only be used inside functions, not at package level:

```go
// Package scope - NOT allowed:
// x := 10
// y, x := 10, 5

// Package scope - allowed:
var x, y = 5, 10

// Inside function - allowed:
func main() {
    x := 10        // OK
    y, z := 5, 20  // OK
}
```
