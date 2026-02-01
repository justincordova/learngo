# String Operations Notes

## String Literals and Escape Sequences

### Interpreted vs Raw String Literals

Go treats escape sequences differently in interpreted strings (double quotes) vs raw strings (backticks):

```go
"\"Hello\\"" + ` \"World\"`
// Result: "Hello" \"World\"
```

Breakdown:
- `"\"Hello\\""` → `"Hello"` (escape sequences interpreted)
- ` \"World\"` → `\"World\"` (raw string, no interpretation)

### Multi-line Strings

Raw string literals (backticks) support multi-line content naturally:

```go
// Correct - using raw string literal:
`<xml>
    <items>
        <item>"Teddy Bear"</item>
    </items>
</xml>`
```

Regular string literals cannot span multiple lines:
- Multi-line with `"..."` - ERROR
- No need to escape quotes in raw literals

## String Length with len()

### Basic Length

`len()` counts bytes, not characters:

```go
len("lovely")  // 6 bytes
```

### Escape Sequences in Length

Interpreted strings count escape sequences as their result:

```go
len("very") + len("\"cool\"")  // 4 + 6 = 10
// "\"cool\"" becomes "cool" (6 characters including quotes)
```

Raw strings count everything literally:

```go
len("very") + len(`\"cool\"`)  // 4 + 8 = 12
// `\"cool\"` stays as \"cool\" (8 characters)
```

### Multi-byte Characters

`len()` counts bytes, not Unicode characters:

```go
len("péripatéticien")  // 16 bytes (not 14)
// é is 2 bytes, so: 12 regular letters + (2 × 2 é's) = 16 bytes
```

## Counting Unicode Characters

### Using utf8.RuneCountInString

To count actual characters (runes/codepoints):

```go
utf8.RuneCountInString("péripatéticien")  // 14 characters
```

Correct usage:
- Package name is `utf8` (not `unicode/utf8`)
- Import: `"unicode/utf8"`
- Call: `utf8.RuneCountInString(s)`

Incorrect:
- `len(péripatéticien)` - missing quotes
- `len("péripatéticien")` - counts bytes (16), not runes (14)
- `unicode/utf8.RuneCountInString()` - wrong package prefix

## String Manipulation Package

The `strings` package provides string manipulation functions:

```go
import "strings"
```

Not valid:
- `string` - no such package
- `unicode/strings` - wrong path

### strings.Repeat

Repeats a string n times:

```go
strings.Repeat("*x", 3) + "*"  // "*x*x*x*"
```

Breakdown:
- `Repeat("*x", 3)` → `"*x*x*x"`
- `+ "*"` → `"*x*x*x*"`

### strings.ToUpper

Converts string to uppercase:

```go
strings.ToUpper("bye bye ") + "see you!"  // "BYE BYE see you!"
```

Important: Only the first part is uppercase because ToUpper only applies to its argument:
- `ToUpper("bye bye ")` → `"BYE BYE "`
- `+ "see you!"` → `"BYE BYE see you!"`

## Summary

| Function | Purpose | Example |
|----------|---------|---------|
| `len(s)` | Count bytes | `len("hello")` → `5` |
| `utf8.RuneCountInString(s)` | Count runes | `utf8.RuneCountInString("café")` → `4` |
| `strings.Repeat(s, n)` | Repeat string | `strings.Repeat("x", 3)` → `"xxx"` |
| `strings.ToUpper(s)` | Uppercase | `strings.ToUpper("hi")` → `"HI"` |

## Key Points

1. Raw string literals (``) don't interpret escape sequences
2. `len()` counts bytes, not characters
3. Multi-byte characters (like é) take multiple bytes
4. Use `utf8.RuneCountInString()` for character count
5. String functions only affect their arguments, not concatenated parts
