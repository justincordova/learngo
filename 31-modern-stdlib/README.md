# Modern Standard Library Features (Go 1.25+)

This section explores new features introduced in Go 1.25 and later versions. Go 1.25 brings significant improvements to the standard library, including experimental JSON v2 support and built-in CSRF protection.

## Overview

Go 1.25 introduces several modern features that enhance security, performance, and developer experience:

- **JSON v2 (Experimental)**: A major revision of the JSON package with better performance and new capabilities
- **CSRF Protection**: Built-in middleware for protecting against Cross-Site Request Forgery attacks
- **Performance Improvements**: 10-40% reduction in garbage collection overhead for JSON-heavy applications

## Prerequisites

Before starting this section, you should be comfortable with:

- Go modules and package management
- HTTP servers and handlers
- JSON marshaling and unmarshaling
- Middleware patterns in Go
- Environment variables and build flags

## Requirements

**Go Version**: Go 1.25 or later is required for the examples in this section.

Check your Go version:
```bash
go version
```

## Section Contents

1. **[JSON v2 (Experimental)](01-json-v2/)** - Learn about the new `encoding/json/v2` package
   - Performance improvements
   - New streaming API
   - Enhanced customization
   - Requires `GOEXPERIMENT=jsonv2` build flag

2. **[CSRF Protection](02-csrf-protection/)** - Use the new `CrossOriginProtection` middleware
   - Modern Fetch metadata-based protection
   - No token management required
   - Simple integration with existing handlers

## Key Features in Go 1.25

### encoding/json/v2 (Experimental)

The new JSON implementation provides:

- **Better Performance**: Substantially faster decoding, encoding at parity
- **Streaming API**: Direct `MarshalWrite` and `UnmarshalRead` functions
- **Enhanced Customization**: Generic marshalers and unmarshalers
- **Powerful Struct Tags**: Support for inline fields, custom formatting, and unknown field handling

**Note**: This is an experimental feature requiring the `jsonv2` GOEXPERIMENT flag.

### net/http.CrossOriginProtection

Built-in CSRF protection that:

- Uses modern browser Fetch metadata headers
- Requires no token generation or validation
- Supports origin-based and pattern-based bypasses
- Works automatically with modern browsers

## Build Flags

Some features in this section require experimental build flags:

### JSON v2
```bash
GOEXPERIMENT=jsonv2 go run main.go
GOEXPERIMENT=jsonv2 go build
GOEXPERIMENT=jsonv2 go test
```

## Best Practices

### JSON v2
- Test your existing code with `GOEXPERIMENT=jsonv2` before adopting
- Use the new streaming API for better performance with large datasets
- Take advantage of enhanced struct tags for complex marshaling needs
- Monitor the proposal for changes before v2 becomes standard

### CSRF Protection
- Apply `CrossOriginProtection` to all state-changing endpoints
- Understand that this requires modern browser support
- Consider fallbacks for older browsers if needed
- Use pattern-based bypasses for legitimate cross-origin requests (e.g., webhooks)

## Common Patterns

### Applying CSRF Protection to a Mux
```go
mux := http.NewServeMux()
mux.HandleFunc("/api/users", handleUsers)
mux.HandleFunc("/api/posts", handlePosts)

// Wrap the entire mux with CSRF protection
protected := http.CrossOriginProtection(mux)
http.ListenAndServe(":8080", protected)
```

### Using JSON v2 Streaming
```go
// With GOEXPERIMENT=jsonv2
var w bytes.Buffer
err := json.MarshalWrite(&w, data)

var result MyType
err := json.UnmarshalRead(&r, &result)
```

## Resources

- [Go 1.25 Release Notes](https://go.dev/doc/go1.25)
- [JSON v2 Proposal](https://github.com/golang/go/issues/71497)
- [JSON v2 Benchmarks](https://github.com/go-json-experiment/jsonbench)
- [Fetch Metadata Specification](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Site)

## Next Steps

After completing this section, you'll be ready to:
- Use modern JSON APIs for better performance
- Implement secure CSRF protection without tokens
- Take advantage of Go 1.25's performance improvements
- Prepare for future standard library evolution

## Important Notes

- **Experimental Features**: JSON v2 is experimental and may change before becoming standard
- **Browser Support**: CSRF protection requires modern browsers with Fetch metadata support
- **Testing**: Always test experimental features thoroughly before production use
- **Feedback**: The Go team encourages feedback on experimental features
