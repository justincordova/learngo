# Predeclared Types Notes

## What are Predeclared Types?

Predeclared types are built-in data types that come with Go. You can use them anywhere without importing any package.

Examples of predeclared types:
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `complex64`, `complex128`
- `byte` (alias for `uint8`)
- `rune` (alias for `int32`)
- `bool`
- `string`

Not predeclared (must import):
- `duration` (from `time` package)

## Bits and Bytes

### Bits to States

With n bits, you can represent 2^n different states:

```
8 bits = 2^8 = 256 different states/numbers
```

### Bytes to Bits

1 byte = 8 bits

```
2 bytes = 16 bits
4 bytes = 32 bits
8 bytes = 64 bits
```

## Binary Representation

Binary numbers are read from right to left, with each position representing 2^(position-1):

```go
fmt.Printf("%08b = %d", 2, 2)
// Prints: 00000010 = 2

// Position from right (1-indexed):
// Position 1: 2^0 = 1
// Position 2: 2^1 = 2  ← the 1 is here
// Position 3: 2^2 = 4
// Position 4: 2^3 = 8
```

Examples:
- `00000001` = 1 (2^0)
- `00000010` = 2 (2^1)
- `00000100` = 4 (2^2)
- `00001000` = 8 (2^3)

## Integer Type Sizes

### Bytes Per Type

Type sizes in bytes:

```
int8, uint8   = 1 byte  (8 bits)
int16, uint16 = 2 bytes (16 bits)
int32, uint32 = 4 bytes (32 bits)
int64, uint64 = 8 bytes (64 bits)
```

Formula: bits ÷ 8 = bytes

Examples:
- `int64` uses 64 bits ÷ 8 = 8 bytes
- `uint32` uses 32 bits ÷ 8 = 4 bytes

### Platform-Dependent int

The `int` and `uint` types have variable size:

```go
// Size depends on target architecture:
// 32-bit system: int = 32 bits
// 64-bit system: int = 64 bits
```

Go determines the size at compile-time based on the target platform.

## Choosing the Right Type

### For English Letters

English letters fit in the range 0-255 (ASCII):

```go
var letter byte  // byte can represent 0-255
// Perfect for English letters and basic ASCII
```

Why `byte` is best:
- `byte` (uint8): 0-255 range, exactly what's needed
- `rune` (int32): Too large (4 bytes for 255 values)
- `int64`: Way too large (8 bytes)
- `float64`: Wrong type (floats are for decimals)

## Integer Overflow/Wraparound

### Unsigned Integer Wraparound

When an unsigned integer exceeds its maximum, it wraps to 0:

```go
var letter uint8 = 255  // Max value for uint8
fmt.Print(letter + 5)   // Prints: 4

// Explanation:
// 255 + 1 = 0 (wraps around)
// 255 + 2 = 1
// 255 + 3 = 2
// 255 + 4 = 3
// 255 + 5 = 4
```

### Signed Integer Wraparound

When a signed integer goes below its minimum, it wraps to the maximum:

```go
var num int8 = -128     // Min value for int8
fmt.Print(num - 3)      // Prints: 125

// Explanation:
// int8 range: -128 to 127
// -128 - 1 = 127 (wraps around)
// -128 - 2 = 126
// -128 - 3 = 125
```

## Type Ranges

Common type ranges:

| Type | Bytes | Range |
|------|-------|-------|
| `int8` | 1 | -128 to 127 |
| `uint8` / `byte` | 1 | 0 to 255 |
| `int16` | 2 | -32,768 to 32,767 |
| `uint16` | 2 | 0 to 65,535 |
| `int32` / `rune` | 4 | -2^31 to 2^31-1 |
| `uint32` | 4 | 0 to 2^32-1 |
| `int64` | 8 | -2^63 to 2^63-1 |
| `uint64` | 8 | 0 to 2^64-1 |

## Key Points

1. Predeclared types are built-in, no import needed
2. 1 byte = 8 bits, 8 bits = 256 possible states
3. Type size = bits ÷ 8 bytes
4. `int` size varies by platform (32 or 64 bits)
5. `byte` is perfect for ASCII/English characters
6. Integers wrap around when they overflow
7. Choose smallest type that fits your data range
