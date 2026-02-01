# Strings, Runes and Bytes

## Strings as Byte Slices

A string is internally represented as a sequence of bytes. Each character's Unicode code point is encoded as one or more bytes:

```go
// "keeper" with corresponding code points:
// k => 107
// e => 101
// p => 112
// r => 114

bytes := []byte{107, 101, 101, 112, 101, 114}
str := string(bytes)  // "keeper"
```

The byte slice `[]byte{107, 101, 101, 112, 101, 114}` represents the string "keeper" because each byte corresponds to the UTF-8 encoding of each character.

## Converting Integers to Strings

When you convert an integer to a string, Go interprets it as a Unicode code point and produces the corresponding character:

```go
// Code points:
// g => 103
// o => 111
fmt.Println(string(103), string(111))
// Prints: g o
```

The conversion `string(103)` produces "g" (not "103") because 103 is the Unicode code point for the letter 'g'.

## Rune Count vs Byte Length

For strings containing multi-byte characters, the number of runes (characters) differs from the number of bytes:

```go
const word = "gökyüzü"
bword := []byte(word)

// ö => 2 bytes in UTF-8
// ü => 2 bytes in UTF-8
fmt.Println(utf8.RuneCount(bword), len(word), len(string(word[1])))
// Prints: 7 10 2
```

**Breakdown:**
- `utf8.RuneCount(bword)` = 7 (seven characters/runes: g, ö, k, y, ü, z, ü)
- `len(word)` = 10 (ten bytes total, because ö and ü each take 2 bytes)
- `len(string(word[1]))` = 2 (byte at index 1 is part of ö, which is 2 bytes)

## For Range Loops Over Strings

The `for range` loop iterates over the **runes** (characters) of a string, not the bytes:

```go
str := "gökyüzü"
for i, r := range str {
    fmt.Printf("%d: %c\n", i, r)
}
```

This correctly handles multi-byte characters, giving you each complete character rather than individual bytes. The index `i` represents the byte position where the rune starts, which may not increment by 1 for multi-byte characters.

## Rune Indexing in UTF-8

In a UTF-8 encoded string, runes can occupy multiple bytes, so their **start and end positions differ**:

```go
str := "gökyüzü"
// g: bytes 0-0 (1 byte)
// ö: bytes 1-2 (2 bytes)
// k: bytes 3-3 (1 byte)
// y: bytes 4-4 (1 byte)
// ü: bytes 5-6 (2 bytes)
// z: bytes 7-7 (1 byte)
// ü: bytes 8-9 (2 bytes)
```

ASCII characters occupy 1 byte (start index = end index), but characters outside the ASCII range can occupy 2, 3, or 4 bytes (start index ≠ end index).

## String Immutability

String values in Go are **immutable**—you cannot change the bytes of a string after it's created:

```go
str := "hello"
// str[0] = 'H'  // Compile error!
```

**Why strings are immutable:**

1. **Strings are immutable byte slices** - Unlike slices, string backing arrays cannot be modified
2. **Strings are shared** - Go shares string values behind the scenes for efficiency. If you could modify a string, it would unexpectedly affect all references to that string

To "modify" a string, you must create a new string:

```go
str := "hello"
str = "H" + str[1:]  // Creates a new string "Hello"
```

Or convert to a byte slice, modify it, then convert back:

```go
str := "hello"
bytes := []byte(str)
bytes[0] = 'H'
str = string(bytes)  // "Hello"
```
