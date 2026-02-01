# Assignment Operations Notes

## Basic Increment/Decrement

### Incrementing by 1

```go
var n float64
n = n + 1  // Correct: increases n by 1
```

Incorrect attempts:
- `n = +1` - Just assigns 1 to n
- `n = n++` - ERROR: ++ is a statement, not an operator
- `++n` - ERROR: Go doesn't support prefix notation

### Decrementing by 1

```go
var n int
n = n - 1  // Correct: decreases n by 1
```

Incorrect attempts:
- `n = -1` - Just assigns -1 to n
- `n = n--` - ERROR: -- is a statement, not an operator
- `--n` - ERROR: Go doesn't support prefix notation

## Increment/Decrement Statements

Go provides `++` and `--` as statements (not operators):

```go
n++  // Equivalent to: n = n + 1
n--  // Equivalent to: n = n - 1
```

Important restrictions:
- Only postfix notation: `n++` (not `++n`)
- Only as statements: `n++` (not `x = n++`)
- Only with addressable values: `n++` (not `1++`)

### Examples

```go
// Equivalent expressions for n = n + 1:
n++
n += 1

// Equivalent expressions for n = n - 1 (or n -= 1):
n--
```

## Compound Assignment Operators

Go supports compound assignment for all arithmetic operators:

### Division Assignment

```go
length /= 10  // Equivalent to: length = length / 10
```

Incorrect:
- `length = length // 10` - ERROR: `//` is not an operator in Go
- `length //= 10` - ERROR: `//=` doesn't exist

### Remainder Assignment

```go
x %= 2  // Equivalent to: x = x % 2
```

Incorrect:
- `x = x / 2` - This is division, not remainder
- `x =% 2` - ERROR: operator on wrong side

### Other Compound Operators

```go
x += 5   // x = x + 5
x -= 3   // x = x - 3
x *= 2   // x = x * 2
x /= 4   // x = x / 4
x %= 7   // x = x % 7
```

## String to Number Conversion

### Using strconv.ParseFloat

The correct function is in the `strconv` package:

```go
strconv.ParseFloat("10", 64)  // Correct
```

Function signature:
```go
func ParseFloat(s string, bitSize int) (float64, error)
```

Parameters:
- `s`: string to convert
- `bitSize`: 32 or 64 (for float32 or float64)

Incorrect:
- `fmtconv.ToFloat` - no such package
- `conv.ParseFloat` - wrong package name
- `strconv.ToFloat` - no such function
- `strconv.ParseFloat("10", 128)` - no 128-bit floats (at runtime)
- `strconv.ParseFloat("10", "64")` - bitSize must be int, not string
- `strconv.ParseFloat(10, 64)` - first arg must be string, not int

## Key Points

1. `++` and `--` are statements, not operators
2. Go only supports postfix: `n++`, not `++n`
3. Compound operators: `+=`, `-=`, `*=`, `/=`, `%=`
4. Use `strconv.ParseFloat` for string to float conversion
5. bitSize for ParseFloat is 32 or 64