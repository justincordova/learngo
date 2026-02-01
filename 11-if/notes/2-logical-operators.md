# Logical Operators Notes

## Logical Operators in Go

The logical operators are:

- `&&` (AND)
- `||` (OR)
- `!` (NOT)

NOT a logical operator:
- `!=` - This is a comparison operator ("not equal")

## Return Type

All logical operators return a boolean value:

```go
result := true && false  // type: bool, value: false
```

The type is always `bool` (not `int`, `byte`, or `float64`).

## Operand Requirements

Logical operators require `bool` operands:

```go
true && false   // OK - both are bool
5 > 3 || 2 < 1  // OK - expressions yield bool
```

Cannot use with other types:
- `int`
- `byte`
- `float64`
- `string`

```go
1 && 0     // ERROR - int values, not bool
"hi" || "" // ERROR - string values, not bool
```

## Logical Expressions

### AND Operator (&&)

Translating "and" conditions:

```
"age is equal or above 15 and hair color is yellow"
```

```go
age >= 15 && hairColor == "yellow"
```

Components:
- `age >= 15` - First condition (bool)
- `&&` - AND operator
- `hairColor == "yellow"` - Second condition (bool)

### NOT Operator (!)

Negates a boolean value:

```go
var (
    on  = true
    off = !on   // false
)

fmt.Println(!on && !off)  // false && true = false
fmt.Println(!on || !off)  // false || true = true
```

## Type Compatibility

Go does not treat numeric values as booleans:

```go
on := 1
fmt.Println(on == true)  // ERROR - can't compare int to bool
```

In Go, `1` is not `true`. This is different from C-based languages.

## String Comparison with Logical Operators

Strings can be compared, and results can be combined with logical operators:

```go
// Note: "a" comes before "b" alphabetically
a := "a" > "b"   // false
b := "b" <= "c"  // true
fmt.Println(a || b)  // false || true = true
```

Result is `true` (a bool), not a string like `"a"` or `"b"`.

## Short-Circuit Evaluation

Logical operators in Go short-circuit:

### AND Short-Circuit

If the left side is `false`, the right side is not evaluated:

```go
isOff() && isOn()  // If isOff() returns false, isOn() is NOT called
```

### OR Short-Circuit

If the left side is `true`, the right side is not evaluated:

```go
isOn() || isOff()  // If isOn() returns true, isOff() is NOT called
```

### Complete Example

```go
func isOn() bool {
    fmt.Print("on ")
    return true
}

func isOff() bool {
    fmt.Print("off ")
    return false
}

func main() {
    _ = isOff() && isOn()  // Prints: "off " (isOn not called)
    _ = isOn() || isOff()  // Prints: "on " (isOff not called)
}
// Total output: "off on "
```

Explanation:
1. `isOff() && isOn()`:
   - `isOff()` prints "off " and returns `false`
   - AND short-circuits, `isOn()` is not called
   - Result: prints only "off "

2. `isOn() || isOff()`:
   - `isOn()` prints "on " and returns `true`
   - OR short-circuits, `isOff()` is not called
   - Result: prints only "on "

## Operator Truth Tables

### AND (&&)

| Left | Right | Result |
|------|-------|--------|
| true | true  | true   |
| true | false | false  |
| false | (not evaluated) | false |

### OR (||)

| Left | Right | Result |
|------|-------|--------|
| true | (not evaluated) | true |
| false | true | true |
| false | false | false |

### NOT (!)

| Value | Result |
|-------|--------|
| true  | false  |
| false | true   |

## Key Points

1. Logical operators: `&&` (AND), `||` (OR), `!` (NOT)
2. All logical operators return `bool`
3. Operands must be `bool` type
4. `!=` is a comparison operator, not a logical operator
5. Go doesn't treat `1` as `true` - no implicit conversion
6. Logical operators short-circuit:
   - `&&` stops if left side is `false`
   - `||` stops if left side is `true`
7. Short-circuiting can prevent side effects (function calls)
8. Use comparison operators to create bool expressions for logical operators
