# Operator Precedence Notes

## Precedence Rules

Multiplication and division have higher precedence than addition and subtraction. Operations of equal precedence are evaluated left to right.

Precedence order (highest to lowest):
1. Parentheses `()`
2. Unary minus `-`
3. Multiplication `*`, Division `/`, Remainder `%`
4. Addition `+`, Subtraction `-`

## Examples

### Expression: 5 - 2 * 5 + 7

```go
5 - 2 * 5 + 7
// Step 1: 2 * 5 = 10 (multiplication first)
// Step 2: 5 - 10 + 7
// Step 3: -5 + 7 = 2
// Result: 2
```

### Expression: 5 - (2 * 5) + 7

```go
5 - (2 * 5) + 7
// Parentheses don't change anything here since multiplication
// already has higher precedence
// Step 1: 2 * 5 = 10
// Step 2: 5 - 10 + 7 = 2
// Result: 2
```

This is equivalent to `5 - 2 * 5 + 7`.

### Expression: 5 - 2 * (5 + 7)

```go
5 - 2 * (5 + 7)
// Step 1: 5 + 7 = 12 (parentheses first)
// Step 2: 2 * 12 = 24
// Step 3: 5 - 24 = -19
// Result: -19
```

Parentheses force the addition to happen before multiplication.

### Expression: 5. -(2 * 5 + 7)

```go
5. -(2 * 5 + 7)
// Step 1: 2 * 5 = 10
// Step 2: 10 + 7 = 17
// Step 3: -(17) = -17
// Step 4: 5.0 - 17 = -12.0
// Result: -12.0 (float64)
```

The `5.` makes this a float expression, so the result is `-12.0`, not `-12`.

## Type in Precedence

When floats and integers are mixed, the result is always a float:

```go
5. -(2 * 5 + 7)  // Result: -12.0 (float64)
5 -(2 * 5 + 7)   // Result: -12 (int)
```

## Key Points

1. Multiplication and division before addition and subtraction
2. Parentheses override default precedence
3. Operations of equal precedence go left to right
4. Float operands make the result a float type
5. Use parentheses to make complex expressions clearer