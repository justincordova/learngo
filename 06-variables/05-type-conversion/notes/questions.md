# Type Conversion Notes

## Syntax

Type conversion uses the type name as a function:

```go
int(4.)      // Convert float to int
float64(10)  // Convert int to float64
string(65)   // Convert int to string (byte value)
```

Common incorrect attempts:
- `convert(40)` - no such function
- `var("hi")` - var is not a conversion function
- `int[4]` - wrong syntax, square brackets are for indexing

## Float to Integer Conversion

When converting a float to an integer, the fractional part is truncated (not rounded):

```go
age := 6.5
fmt.Print(int(age))  // Prints: 6
```

The `.5` is lost, leaving only `6`.

## Compile-Time Detection

Go detects type conversion errors at compile-time when possible:

```go
fmt.Print(int(6.5))  // Compile-Time Error
```

This is an error because literal conversion issues can be detected before runtime.

## Type Compatibility in Operations

Operations require compatible types. You cannot mix types:

```go
area := 10.5  // float64
fmt.Print(area/2)  // OK - 2 is treated as float64, prints 5.25
```

```go
area := 10.5  // float64
div := 2      // int
fmt.Print(area/div)  // ERROR - can't divide float64 by int
```

## Fixing Type Mismatches

To fix type mismatches, convert one value to match the other's type:

```go
area := 10.5
div := 2
fmt.Print(area/float64(div))  // OK - prints 5.25
```

Other approaches and their results:
- `int(area)/div` → `5` (both integers, integer division)
- `area/int(div)` → Type mismatch error
- `int(area)/int(div)` → `5` (both integers, integer division)
- `area/float64(div)` → `5.25` (both floats, correct result)

## Key Points

1. Type conversions are explicit in Go - no automatic type coercion
2. Float to int conversion truncates the decimal part
3. Operations require matching types
4. Use `Type(value)` syntax for conversions
5. Choose your conversion carefully to preserve precision when needed