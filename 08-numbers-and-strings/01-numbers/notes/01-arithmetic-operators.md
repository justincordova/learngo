# Arithmetic Operators Notes

## Arithmetic Operators in Go

The arithmetic operators are:
- `*` (multiplication)
- `/` (division)
- `%` (remainder/modulo)
- `+` (addition)
- `-` (subtraction)

Not arithmetic operators:
- `**` (not in Go)
- `^` (XOR, not exponentiation)
- `!` (logical NOT)
- `++`, `--` (increment/decrement statements, not operators)
- `&`, `|` (bitwise operators)

## Remainder Operator (%)

The remainder operator only works with integer values:

```go
8 % 3  // Result: 2 (8 divided by 3 is 2 with remainder 2)
```

Cannot be used with:
- Floats: `3.54 % 2` - ERROR
- Bools: `true % false` - ERROR
- Strings: `"Try Me!" % "hi"` - ERROR

## Negative Numbers

```go
-(3 * -2)  // Result: 6
// Step 1: 3 * -2 = -6
// Step 2: -(-6) = 6
```

## Integer Division

When dividing integers, the result is always an integer (fractional part is truncated):

```go
var degree float64 = 10 / 4  // degree = 2.0 (not 2.5)
```

Even though `degree` is `float64`, the division `10 / 4` happens with integers first, resulting in `2`, which is then converted to `2.0`.

## Float Division

If any operand is a float, the result is a float:

```go
var degree float64 = 3. / 2  // degree = 1.5
```

The `3.` (with a decimal point) makes the whole expression a float operation.

## Type Inference in Expressions

The type of an expression depends on the operands:

```go
x := 5 * 2.   // x is float64 (because of 2.)
x := 5 * -(2) // x is int (both operands are integers)
```

A single float operand makes the entire expression float:

```go
5 * 2.  // Result type: float64
5 * 2   // Result type: int
```

## Float Precision Issues

Floating-point calculations can be inaccurate due to how computers represent decimal numbers:

```go
0.1 + 0.2  // May not exactly equal 0.3
```

This is a limitation of binary floating-point representation, not a Go-specific issue. Integers do not have this problem.

## Key Points

1. Remainder (%) only works with integers
2. Integer division truncates fractional parts
3. Any float operand makes the whole expression float
4. Floats can have precision issues
5. Use `3.` notation to force float division
