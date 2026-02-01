# Defined Types Notes

## Why Define New Types?

Defined types provide multiple benefits:

1. **Declare methods** - Add behavior to types
2. **Type-safety** - Prevent mixing incompatible values
3. **Readability** - Convey meaning and intent
4. **All of the above** - The complete answer

Example:
```go
type Celsius float64
type Fahrenheit float64

// Now you can't accidentally mix temperature units
var c Celsius = 100
var f Fahrenheit = 212
// c = f  // ERROR: type mismatch
```

## What Defined Types Inherit

When you define a new type from an existing type, the new type inherits:

- **Representation** (how it's stored in memory)
- **Size** (number of bytes)
- **Range of values** (min/max)

The new type does NOT inherit:

- **Methods** from the source type

```go
type Millennium time.Duration

// Millennium has:
// ✓ Same representation as time.Duration (int64)
// ✓ Same size (8 bytes)
// ✓ Same value range
// ✗ No methods from time.Duration
```

## Defining a New Type

### Syntax

```go
type radius float32  // Defines radius as a new type based on float32
```

Incorrect syntax:
- `var radius float32` - This declares a variable, not a type
- `radius = type float32` - Invalid syntax
- `type radius = float32` - This creates an alias (same type), not a new type

### Type Definition vs Type Alias

```go
// Type definition (creates NEW type):
type Celsius float64     // Celsius and float64 are different types

// Type alias (creates SAME type):
type Fahrenheit = float64  // Fahrenheit and float64 are the same type
```

## Type Compatibility

Defined types are not compatible with their underlying types without conversion:

```go
type Distance int

var (
    village Distance = 50
    city = 100           // city is int (inferred)
)

// fmt.Print(village + city)  // ERROR: type mismatch

// Fix with type conversion:
fmt.Print(village + Distance(city))  // OK
```

Solutions:
- `village + Distance(city)` - Convert `city` to `Distance`
- NOT `int(village + city)` - Can't add before converting
- NOT `village(int) + city` - Invalid syntax
- NOT `village + int(city)` - Still type mismatch

## Practical Example: Temperature Conversion

For a temperature conversion program:

```go
celsius := 35.
fahrenheit := (9*celsius + 160) / 5
fmt.Printf("%g ºC is %g ºF\n", celsius, fahrenheit)
```

Best type choices:
- Define `Celsius` and `Fahrenheit` types using `float64`
- Temperatures have fractional parts, so use float
- Two distinct units deserve distinct types for safety

```go
type Celsius float64
type Fahrenheit float64
```

Why not alternatives:
- `int64` - Can't represent fractional degrees
- Single `Temperature` type - Doesn't distinguish units
- `uint32` - Can't represent negative temperatures

## Understanding Underlying Types

Go's type system is flat. The underlying type is always a predeclared type:

```go
type (
    Duration int64      // Underlying: int64
    Century Duration    // Underlying: int64 (not Duration)
    Millennium Century  // Underlying: int64 (not Century)
)
```

All three types have `int64` as their underlying type. The chain always resolves to a predeclared type with actual structure.

## Type Aliases (No Conversion Needed)

Some types are aliases to predeclared types and don't require conversion:

```go
byte   // Alias for uint8
rune   // Alias for int32
```

These pairs don't need conversion:

```go
var b byte = 65
var u uint8 = 65
b = u  // OK - same type

var r rune = 'A'
var i int32 = 65
r = i  // OK - same type
```

But these need conversion:

```go
// byte (uint8) and rune (int32) - different types
// byte and uint32 - different types
// byte and int8 - different types (signed vs unsigned)
```

## Key Points

1. Defined types provide type-safety, readability, and method attachment
2. New types inherit representation, size, and range, but NOT methods
3. Use `type Name Type` for new types, `type Name = Type` for aliases
4. Defined types need explicit conversion from their underlying type
5. Go's type system is flat - underlying type is always predeclared
6. `byte` = `uint8` and `rune` = `int32` (aliases)
7. Choose types that match your data and convey meaning
