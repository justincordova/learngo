# Modern Standard Library Examples

Learn about modern Go standard library features introduced in recent versions.

## Topics Covered

1. **JSON v2** (Go 1.25, Experimental) - New `encoding/json/v2` package
2. **CSRF Protection** (Go 1.25) - `net/http.CrossOriginProtection()` middleware
3. **Zero-Allocation Reflection** (Go 1.25) - `reflect.TypeAssert()` for performance

## Prerequisites

- Go 1.25+
- Understanding of JSON, HTTP servers, and reflection basics

## Note on Experimental Features

Some examples use experimental features (like `encoding/json/v2`) that require:
```bash
GOEXPERIMENT=jsonv2 go run main.go
```

Check each example's README for specific requirements.
