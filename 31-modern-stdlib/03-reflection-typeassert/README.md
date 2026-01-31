# Zero-Allocation Reflection (Go 1.25)

Learn about `reflect.TypeAssert()`, a new Go 1.25 function for zero-allocation reflection.

## The Problem with Traditional Reflection

Traditional reflection uses `Interface()` which allocates memory:

```go
v := reflect.ValueOf(42)
x := v.Interface().(int)  // Allocates on heap!
```

This is fine for occasional use, but in hot paths it creates garbage collection pressure.

## The Solution: reflect.TypeAssert (Go 1.25)

Go 1.25 introduces `reflect.TypeAssert[T]()` for zero-allocation type extraction:

```go
v := reflect.ValueOf(42)
x, ok := reflect.TypeAssert[int](v)  // Zero allocations!
```

### Key Benefits

1. **Zero allocations** - No heap allocations, no GC pressure
2. **Type safe** - Uses Go generics for compile-time safety
3. **Performance** - Significantly faster than `Interface()`
4. **Same API** - Familiar comma-ok pattern like type assertions

## When to Use

Use `reflect.TypeAssert` when:
- Working with reflection in performance-critical code
- Processing large amounts of reflected values
- Avoiding GC pressure is important
- You know the expected type at compile time

Use traditional `Interface()` when:
- Performance isn't critical
- You need the interface{} value itself
- Working with unknown types

## Performance

In benchmarks, `TypeAssert` is typically 5-10x faster than `Interface()` due to:
- No heap allocations
- No interface conversion overhead
- Compiler optimizations with generics

## Example Usage

```go
// Extracting values from reflect.Value
v := reflect.ValueOf(someVar)

// Type-safe extraction
if num, ok := reflect.TypeAssert[int](v); ok {
    // Use num directly, no allocation
}

// Works with any type
if str, ok := reflect.TypeAssert[string](v); ok {
    fmt.Println(str)
}

// Struct types too
type User struct{ Name string }
if user, ok := reflect.TypeAssert[User](v); ok {
    fmt.Println(user.Name)
}
```

## Learn More

- [Go 1.25 Release Notes](https://go.dev/doc/go1.25)
- [reflect package documentation](https://pkg.go.dev/reflect)

## Running This Example

```bash
go run main.go
```

The example demonstrates the performance difference and usage patterns.
