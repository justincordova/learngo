# Context Values and Best Practices

This example demonstrates how to use `context.WithValue()` to store and retrieve request-scoped data, along with best practices and common anti-patterns to avoid.

## Key Concepts

### context.WithValue()

Stores a key-value pair in a context:

```go
ctx := context.WithValue(parentCtx, key, value)
```

### Retrieving Values

```go
value := ctx.Value(key)
if value != nil {
    // Type assert to expected type
    str, ok := value.(string)
    if ok {
        // Use str
    }
}
```

## Type-Safe Keys (CRITICAL)

**Always use custom types for context keys**, never plain strings:

### Wrong (Don't Do This)

```go
// BAD: String keys can collide with other packages
ctx := context.WithValue(ctx, "userID", "123")
value := ctx.Value("userID")
```

### Correct Approach

```go
// GOOD: Custom type prevents collisions
type contextKey string

const userIDKey contextKey = "userID"

ctx := context.WithValue(ctx, userIDKey, "123")
value := ctx.Value(userIDKey)
```

### Even Better - Unexported Types

```go
// BEST: Unexported struct type for maximum type safety
type userKey struct{}

ctx := context.WithValue(ctx, userKey{}, user)
value := ctx.Value(userKey{})
```

**Why this matters:** Two packages using string key "userID" will collide. Custom types prevent this entirely.

## When to Use Context Values

### Good Use Cases ✓

#### 1. Request-Scoped Identifiers

```go
// Request IDs for tracing
ctx := context.WithValue(ctx, requestIDKey, "req-12345")

// Correlation IDs for distributed systems
ctx = context.WithValue(ctx, correlationIDKey, "corr-abc-xyz")
```

**Use for:** Data that needs to flow through the entire request lifecycle for logging, tracing, and debugging.

#### 2. Authentication/Authorization Data

```go
type User struct {
    ID       string
    Username string
    Role     string
}

// After authentication, store user in context
ctx := context.WithValue(ctx, userKey{}, user)

// Later, extract for authorization checks
user, ok := ctx.Value(userKey{}).(*User)
```

**Use for:** Authenticated user information that multiple layers need without passing explicitly.

#### 3. Request Metadata

```go
// Client information
ctx = context.WithValue(ctx, clientIPKey, "192.168.1.1")
ctx = context.WithValue(ctx, userAgentKey, "Mozilla/5.0...")

// Localization
ctx = context.WithValue(ctx, localeKey, "en-US")
```

**Use for:** Request-specific data that would otherwise require threading through many function calls.

### Anti-Patterns ✗ (Don't Do These)

#### 1. Optional Function Parameters

```go
// BAD: Hiding parameters in context
ctx := context.WithValue(ctx, "timeout", 5*time.Second)
ctx = context.WithValue(ctx, "retries", 3)
fetchData(ctx, url)

// GOOD: Explicit parameters or options struct
type Options struct {
    Timeout time.Duration
    Retries int
}
fetchData(ctx, url, Options{Timeout: 5*time.Second, Retries: 3})
```

**Why it's bad:**
- Function signature doesn't show available options
- No type safety or compile-time checking
- Hard to discover and test
- Violates principle of explicit interfaces

#### 2. Application Configuration

```go
// BAD: Configuration in context
ctx := context.WithValue(ctx, "dbHost", "localhost")
ctx = context.WithValue(ctx, "dbPort", 5432)

// GOOD: Config struct and dependency injection
type Config struct {
    DBHost string
    DBPort int
}
db := NewDatabase(config)
```

**Why it's bad:**
- Configuration is application-wide, not request-scoped
- Makes dependencies invisible
- Impossible to reason about without runtime inspection
- Testing becomes much harder

#### 3. Passing Dependencies

```go
// BAD: Hiding dependencies in context
ctx := context.WithValue(ctx, "logger", logger)
ctx = context.WithValue(ctx, "database", db)

// GOOD: Explicit dependency injection
type Service struct {
    logger *Logger
    db     *Database
}
```

**Why it's bad:**
- Dependencies should be explicit and visible
- No compile-time checking
- Violates dependency inversion principle
- Makes code harder to understand and test

## Best Practices

### 1. Use Custom Types for Keys

```go
// Define in package with related functionality
type contextKey string

const (
    requestIDKey contextKey = "requestID"
    userIDKey    contextKey = "userID"
)
```

Or even better, use unexported struct types:

```go
type requestIDKey struct{}
type userKey struct{}

// Keys can't collide with anything else
ctx := context.WithValue(ctx, requestIDKey{}, "req-123")
```

### 2. Document What Values Are Expected

```go
// RequestID returns the request ID from the context, or empty string if not set.
func RequestID(ctx context.Context) string {
    id, _ := ctx.Value(requestIDKey).(string)
    return id
}

// WithRequestID returns a new context with the given request ID.
func WithRequestID(ctx context.Context, id string) context.Context {
    return context.WithValue(ctx, requestIDKey, id)
}
```

### 3. Always Type Assert Safely

```go
// Check if value exists and has correct type
user, ok := ctx.Value(userKey{}).(*User)
if !ok {
    // Handle missing or wrong type
    return errors.New("no user in context")
}
```

### 4. Don't Overuse Context Values

If data is required for a function to work, it should be a parameter:

```go
// BAD: Required data in context
func processOrder(ctx context.Context) error {
    orderID := ctx.Value(orderIDKey).(string) // Required but hidden!
    // ...
}

// GOOD: Required data as parameters
func processOrder(ctx context.Context, orderID string) error {
    // Clear and explicit
}
```

### 5. Keep Values Immutable

Values stored in context should be read-only:

```go
// BAD: Storing mutable data
type Config struct {
    Setting string
}
ctx := context.WithValue(ctx, configKey{}, &Config{}) // Mutable pointer

// GOOD: Store immutable data or copies
type Config struct {
    Setting string
}
config := Config{Setting: "value"}
ctx := context.WithValue(ctx, configKey{}, config) // Value copy
```

## Examples in This Program

1. **Basic context values** - Simple WithValue usage (with warning about string keys)
2. **Type-safe keys** - Recommended approach using custom types
3. **Request-scoped values** - Request IDs and correlation IDs for tracing
4. **Authentication data** - Storing and retrieving user information
5. **Value propagation** - How values flow through context hierarchy
6. **Anti-pattern: Optional parameters** - What NOT to do
7. **Anti-pattern: Configuration** - Another common mistake
8. **Good use cases summary** - When to and when not to use context values

## Common Patterns

### Logging with Request Context

```go
func logWithContext(ctx context.Context, message string) {
    reqID := ctx.Value(requestIDKey)
    userID := ctx.Value(userIDKey)

    log.Printf("[ReqID: %v] [UserID: %v] %s", reqID, userID, message)
}
```

### Middleware Adding Context Values

```go
func RequestIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := generateRequestID()
        ctx := context.WithValue(r.Context(), requestIDKey, requestID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

### Helper Functions for Type Safety

```go
// Helper to add user to context
func WithUser(ctx context.Context, user *User) context.Context {
    return context.WithValue(ctx, userKey{}, user)
}

// Helper to extract user from context
func GetUser(ctx context.Context) (*User, error) {
    user, ok := ctx.Value(userKey{}).(*User)
    if !ok {
        return nil, errors.New("no user in context")
    }
    return user, nil
}
```

## Running the Example

```bash
go run main.go
```

The program demonstrates correct usage patterns and highlights common anti-patterns to avoid.

## Key Takeaways

1. **Always use custom types for keys** - Never use plain strings
2. **Use for request-scoped data only** - Not for configuration or dependencies
3. **Document expected values** - Create helper functions for type safety
4. **Don't hide required parameters** - If a function needs it, make it explicit
5. **Type assert safely** - Always check the type assertion result
6. **Keep values immutable** - Don't store mutable data in contexts
7. **Use sparingly** - Context values are for cross-cutting concerns, not general data passing

## The Golden Rule

**Use context.Value for request-scoped data that crosses API boundaries and would otherwise require threading through many function signatures. Don't use it to hide function parameters or dependencies.**

If you're unsure whether to use context.Value, err on the side of explicit function parameters.
