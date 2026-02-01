# Comparison Operators Notes

## Types of Comparison Operators

### Equality Operators

Test if values are equal or not equal:

- `==` (equal to)
- `!=` (not equal to)

NOT equality operators:
- `>` - This is an ordering operator (greater than)

### Ordering Operators

Test the relative order of values:

- `>` (greater than)
- `<` (less than)
- `>=` (greater than or equal to)
- `<=` (less than or equal to)

NOT ordering operators:
- `==` - This is an equality operator

## Return Type

All comparison operators return a boolean value:

```go
result := 5 > 3  // type: bool, value: true
```

The type is always `bool` (not `int`, `byte`, or `float64`).

## Operand Requirements

### For Ordering Operators (`>`, `<`, `>=`, `<=`)

Can be used with ordered values:
- `int` values
- `byte` values
- `string` values (strings are sequences of numbers)
- `float64` values

Cannot be used with:
- `bool` values (not ordered)

```go
5 > 3          // OK
"abc" < "def"  // OK
true >= false  // ERROR - bool is not ordered
```

### For Equality Operators (`==`, `!=`)

Can be used with any comparable value:
- `int` values
- `byte` values
- `string` values
- `bool` values
- And more...

```go
5 == 3              // OK
"go" != "go!"       // OK
true == false       // OK
```

## String Comparison Examples

Strings are compared lexicographically (alphabetically):

```go
fmt.Println("go" != "go!")  // true (not equal)
fmt.Println("go" == "go!")  // false (not equal due to !)
```

The exclamation mark makes them different strings.

## Type Compatibility

You cannot compare values of incompatible types:

```go
fmt.Println(1 == true)  // ERROR - can't compare int to bool
```

In Go, `1` does not equal `true`. Go is not like C-based languages where numeric values can represent boolean values.

## Ordering with Floats

```go
fmt.Println(2.9 > 2.9)   // false (not greater, they're equal)
fmt.Println(2.9 <= 2.9)  // true (less than or equal, equal counts)
```

### Boolean Ordering Error

```go
fmt.Println(false >= true)  // ERROR - bool values not ordered
fmt.Println(true <= false)  // ERROR - bool values not ordered
```

Bool values are comparable (with `==` and `!=`) but not ordered (cannot use `>`, `<`, `>=`, `<=`).

## Type Conversion for Mixed Types

When working with different numeric types, convert to preserve precision:

```go
weight, factor := 500, 1.5  // weight is int, factor is float64
// weight *= factor  // ERROR - type mismatch

// Fix without losing precision:
weight = int(float64(weight) * factor)  // Result: 750
```

Why this works:
- `float64(weight)` → `500.0`
- `500.0 * 1.5` → `750.0`
- `int(750.0)` → `750`

Why other approaches don't work:
- `weight *= float64(factor)` - Type mismatch (can't assign float64 to int)
- `weight *= int(factor)` - Loses precision (`int(1.5)` becomes `1`)
- `weight = float64(weight) * factor` - Type mismatch (can't assign float64 to int)

## Key Points

1. Equality operators: `==`, `!=`
2. Ordering operators: `>`, `<`, `>=`, `<=`
3. All comparison operators return `bool`
4. Ordering operators require ordered types (not bool)
5. Equality operators work with any comparable type
6. Strings are ordered and compared lexicographically
7. Go doesn't treat `1` as `true` - types must match
8. Bool values are comparable but not ordered
9. Convert types carefully to preserve precision in calculations
