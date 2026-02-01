# Modern Standard Library Notes

## JSON v2 (Go 1.25, Experimental)

The new `encoding/json/v2` package provides improved performance and better API design.

### Enabling JSON v2

```bash
GOEXPERIMENT=jsonv2 go run main.go
```

### Key Improvements

- **Better performance**: Faster encoding/decoding
- **Streaming support**: More efficient for large JSON documents
- **Better error messages**: More detailed parsing errors
- **Consistent API**: Clearer semantics for edge cases

### Basic Usage

```go
import "encoding/json/v2"

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// Encoding
user := User{ID: 1, Name: "Alice"}
data, err := json.Marshal(user)

// Decoding
var user User
err := json.Unmarshal(data, &user)
```

### Differences from v1

- More explicit error handling
- Better streaming API
- Improved handling of special values (NaN, Infinity)
- More predictable struct tag behavior

## CSRF Protection (Go 1.25)

Built-in CSRF protection middleware in `net/http`.

### Using CrossOriginProtection

```go
import "net/http"

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/data", dataHandler)

    // Wrap with CSRF protection
    protected := http.CrossOriginProtection()(mux)

    http.ListenAndServe(":8080", protected)
}
```

### How It Works

The middleware:
1. Checks for valid CSRF tokens on state-changing requests (POST, PUT, DELETE)
2. Validates origin and referer headers
3. Automatically includes CSRF tokens in responses
4. Rejects unauthorized cross-origin requests

### Custom Configuration

```go
protectionMiddleware := http.CrossOriginProtection(
    http.WithCSRFTokenHeader("X-CSRF-Token"),
    http.WithAllowedOrigins("https://example.com"),
)

protected := protectionMiddleware(handler)
```

## Zero-Allocation Reflection (Go 1.25)

New `reflect.TypeAssert()` for performance-critical reflection code.

### Traditional Type Assertion

```go
func process(val interface{}) {
    if str, ok := val.(string); ok {
        fmt.Println(str)
    }
}
```

### Zero-Allocation Type Assertion

```go
import "reflect"

func process(val any) {
    var str string
    if reflect.TypeAssert(&str, val) {
        fmt.Println(str)  // No allocation
    }
}
```

### When to Use

Use `reflect.TypeAssert` when:
- Performance is critical
- You're doing many type assertions in a hot path
- You want to avoid allocations

Use traditional type assertions when:
- Code clarity is more important than performance
- Not in a performance-critical path
- The simpler syntax is preferred

### Performance Comparison

```go
// Traditional (may allocate)
if v, ok := val.(string); ok {
    // ...
}

// Zero-allocation (no heap allocation)
var v string
if reflect.TypeAssert(&v, val) {
    // ...
}
```

## Other Modern Go Features

### Structured Logging (Go 1.21+)

```go
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("user login", "userID", 123, "ip", "192.168.1.1")
```

### Enhanced Testing (Go 1.20+)

```go
func TestSomething(t *testing.T) {
    t.Cleanup(func() {
        // Cleanup runs even if test panics
    })
}
```

### go:embed Directive (Go 1.16+)

```go
import _ "embed"

//go:embed config.json
var configData []byte

//go:embed templates/*
var templateFS embed.FS
```

## Best Practices

### JSON v2

- Use streaming for large documents
- Handle errors explicitly
- Validate JSON schema separately if needed

### CSRF Protection

- Always enable for state-changing operations
- Use HTTPS in production
- Configure allowed origins explicitly
- Don't disable protection without understanding risks

### Zero-Allocation Reflection

- Profile before optimizing
- Use only in hot paths
- Keep code readable
- Document why zero-allocation is necessary

## Migration Guide

### From JSON v1 to v2

Most code works the same, but check:
- Error handling behavior
- Special value handling (NaN, Infinity)
- Custom marshalers/unmarshalers

### Adding CSRF Protection

1. Wrap your handler with `CrossOriginProtection()`
2. Update frontend to include CSRF tokens
3. Test with various origins
4. Configure allowed origins for production

## Resources

- [Go 1.25 Release Notes](https://go.dev/doc/go1.25)
- [encoding/json/v2 documentation](https://pkg.go.dev/encoding/json/v2)
- [CSRF Protection Guide](https://go.dev/blog/csrf-protection)
- [Reflection Performance](https://go.dev/blog/reflection-performance)
