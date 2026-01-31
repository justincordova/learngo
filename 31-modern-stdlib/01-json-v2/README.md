# JSON v2 Package (Go 1.25 Experimental)

This example introduces the new `encoding/json/v2` package, which is experimental in Go 1.25.

## What's New in json/v2?

### Performance Improvements
- **Substantially faster decoding** - Up to 2x faster than v1
- **Encoding at parity** - Similar performance to v1
- **Lower allocations** - Reduced memory overhead

### Better API Design
- More explicit error handling
- Better control over encoding/decoding behavior
- Streaming support with `encoding/json/jsontext`

## Building with json/v2

Since json/v2 is experimental, you need to enable it:

```bash
# Build with experimental json/v2
GOEXPERIMENT=jsonv2 go build

# Run with experimental json/v2
GOEXPERIMENT=jsonv2 go run main.go
```

## Current Example

This example uses the standard `encoding/json` (v1) package since json/v2 is experimental.
It demonstrates the traditional JSON operations and explains how to enable json/v2.

## When json/v2 Becomes Stable

Once json/v2 graduates from experimental status, the import will change:

```go
// Instead of
import "encoding/json"

// Use
import "encoding/json/v2"
```

## Learn More

- [Go 1.25 Release Notes](https://go.dev/doc/go1.25)
- [encoding/json/v2 Proposal](https://github.com/golang/go/discussions/63397)

## Running This Example

```bash
go run main.go
```

This shows the current v1 API and explains how to enable v2 when ready.
