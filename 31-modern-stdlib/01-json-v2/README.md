# JSON v2 (encoding/json/v2) - Experimental

This example demonstrates the new experimental `encoding/json/v2` package introduced in Go 1.25. This is a major revision of the standard `encoding/json` package with substantial performance improvements and new features.

## Overview

The `encoding/json/v2` package provides:

- **Better Performance**: Substantially faster decoding, encoding at parity
- **Streaming API**: Direct `MarshalWrite` and `UnmarshalRead` functions
- **Enhanced Security**: Rejects duplicate keys and invalid UTF-8 by default
- **Powerful Struct Tags**: Support for `omitzero`, `inline`, and more
- **Type Safety**: More consistent error handling and type checking

## Requirements

**Go Version**: Go 1.25 or later
**Build Flag**: `GOEXPERIMENT=jsonv2`

## Building and Running

### Run with GOEXPERIMENT flag
```bash
GOEXPERIMENT=jsonv2 go run main.go
```

### Build with GOEXPERIMENT flag
```bash
GOEXPERIMENT=jsonv2 go build -o jsonv2-demo main.go
./jsonv2-demo
```

### Test your existing code with v2
```bash
GOEXPERIMENT=jsonv2 go test ./...
```

## Key Features Demonstrated

### 1. Basic Marshal/Unmarshal
The v2 API is similar to v1 but with better defaults:
```go
data, err := jsonv2.Marshal(user)
err := jsonv2.Unmarshal(data, &user)
```

### 2. Streaming with MarshalWrite
Write directly to an `io.Writer` without creating an Encoder:
```go
var buf bytes.Buffer
err := jsonv2.MarshalWrite(&buf, user)
```

**Advantages**:
- More efficient than `json.NewEncoder(w).Encode(v)`
- No unnecessary intermediate allocations
- Perfect for HTTP handlers and streaming scenarios

### 3. Streaming with UnmarshalRead
Read directly from an `io.Reader` without creating a Decoder:
```go
var user User
err := jsonv2.UnmarshalRead(reader, &user)
```

**Advantages**:
- Reads entire input until `io.EOF`
- Validates that no unexpected bytes follow the JSON value
- Cleaner API for single-value reads

### 4. Enhanced Struct Tags

#### omitzero
Omit fields with zero values (not just empty):
```go
type User struct {
    Age int `json:"age,omitzero"` // Omits if 0
}
```

**Difference from v1's omitempty**:
- `omitempty`: Omits if "empty" (0, "", nil, empty slice/map)
- `omitzero`: Omits only if the exact zero value for the type

#### string
Stringify numbers in JSON:
```go
type User struct {
    Balance float64 `json:"balance,string"` // "123.45" instead of 123.45
}
```

#### inline
Flatten nested structs into the parent object:
```go
type Product struct {
    Name  string  `json:"name"`
    Metadata struct {
        Category string `json:"category"`
    } `json:",inline"` // Fields become top-level
}
// Result: {"name":"Laptop","category":"Electronics"}
```

### 5. Security Improvements

#### Duplicate Keys
v2 rejects duplicate object keys by default (v1 silently uses the last value):
```go
// This will error in v2
data := `{"name":"Alice","name":"Bob"}`
err := jsonv2.Unmarshal([]byte(data), &user)
// Error: duplicate name "name" in object
```

#### Unknown Fields
Can enforce strict validation:
```go
err := jsonv2.Unmarshal(data, &user, jsonv2.RejectUnknownMembers(true))
```

#### UTF-8 Validation
v2 validates UTF-8 by default, rejecting invalid sequences.

### 6. Performance Comparison

The example includes a simple benchmark showing v2's performance characteristics:
- **Encoding**: Similar performance to v1
- **Decoding**: Substantially faster than v1
- **GC Overhead**: 10-40% reduction in programs that heavily use JSON

## API Reference

### Core Functions

```go
// Marshal to []byte
func Marshal(in any, opts ...Options) ([]byte, error)

// Unmarshal from []byte
func Unmarshal(in []byte, out any, opts ...Options) error

// Marshal to io.Writer (new in v2)
func MarshalWrite(out io.Writer, in any, opts ...Options) error

// Unmarshal from io.Reader (new in v2)
func UnmarshalRead(in io.Reader, out any, opts ...Options) error
```

### Options

```go
// Reject duplicate object keys (default: false)
jsonv2.RejectDuplicateNames(true)

// Reject unknown fields
jsonv2.RejectUnknownMembers(true)

// Case-insensitive field matching
jsonv2.MatchCaseInsensitiveNames(true)

// Deterministic output
jsonv2.Deterministic(true)

// Stringify all numbers
jsonv2.StringifyNumbers(true)
```

## Comparison: v1 vs v2

| Feature | encoding/json (v1) | encoding/json/v2 |
|---------|-------------------|------------------|
| Duplicate keys | Silently uses last | Rejects by default |
| UTF-8 validation | Lenient | Strict by default |
| Unknown fields | Ignores | Can reject with option |
| Streaming API | Encoder/Decoder | MarshalWrite/UnmarshalRead |
| Performance | Baseline | 10-40% better decoding |
| Struct tags | omitempty | omitempty + omitzero + inline |

## Best Practices

### When to Use v2

1. **HTTP Handlers**: Use `MarshalWrite` and `UnmarshalRead`
   ```go
   func handler(w http.ResponseWriter, r *http.Request) {
       var req Request
       if err := jsonv2.UnmarshalRead(r.Body, &req); err != nil {
           http.Error(w, err.Error(), http.StatusBadRequest)
           return
       }

       resp := processRequest(req)
       w.Header().Set("Content-Type", "application/json")
       jsonv2.MarshalWrite(w, resp)
   }
   ```

2. **Security-Critical Applications**: Leverage strict validation
   ```go
   err := jsonv2.Unmarshal(data, &config,
       jsonv2.RejectUnknownMembers(true),
       jsonv2.RejectDuplicateNames(true))
   ```

3. **High-Performance Systems**: Benefit from improved decoding speed

### Testing with v2

Before adopting v2 in production:
```bash
# Test your entire codebase
GOEXPERIMENT=jsonv2 go test ./...

# Check for breaking changes
GOEXPERIMENT=jsonv2 go build ./...
```

### Migration Considerations

- v2 may reject JSON that v1 accepted (duplicate keys, invalid UTF-8)
- Error messages may differ
- Some behavioral changes (e.g., empty slice marshaling)
- Test thoroughly before switching

## Important Notes

- **Experimental**: This is an experimental feature that may change
- **Build Flag Required**: Must use `GOEXPERIMENT=jsonv2`
- **Not Stable**: API may evolve before v2 becomes standard
- **Feedback Welcome**: The Go team encourages feedback on [Issue #71497](https://github.com/golang/go/issues/71497)

## Common Issues

### Build Flag Forgotten
```
Error: package encoding/json/v2 is not in std
```
**Solution**: Add `GOEXPERIMENT=jsonv2` to your build command

### Build Tag Required
If you see compile errors about missing packages, ensure your file has:
```go
//go:build goexperiment.jsonv2
```

## Resources

- [Go 1.25 Release Notes](https://go.dev/doc/go1.25)
- [JSON v2 Proposal](https://github.com/golang/go/issues/71497)
- [JSON v2 Documentation](https://pkg.go.dev/encoding/json/v2)
- [Performance Benchmarks](https://github.com/go-json-experiment/jsonbench)
- [JSON Evolution Article](https://antonz.org/go-json-v2/)

## Next Steps

After exploring JSON v2:
1. Test with your existing code using `GOEXPERIMENT=jsonv2`
2. Identify performance-critical JSON operations in your codebase
3. Experiment with new features like `inline` and `omitzero`
4. Provide feedback to the Go team
5. Monitor the proposal for updates before v2 becomes standard
