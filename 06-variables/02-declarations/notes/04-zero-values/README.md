## Zero values by type

When you declare a variable without initializing it, Go assigns it a zero value based on its type:

**Numeric types** (int, float64, etc.): `0`

```go
var count int  // count = 0
```

**Bool**: `false`

```go
var isActive bool  // isActive = false
```

**String**: `""` (empty string)

```go
var name string  // name = ""
```

**Pointer types**: `nil`

```go
var ptr *int  // ptr = nil
```
