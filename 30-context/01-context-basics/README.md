# Context Basics

This example introduces the fundamental concepts of Go's `context` package, including `context.Background()`, `context.TODO()`, and the proper way to pass context through function calls.

## Concepts Covered

### 1. context.Background()

The root context for your application:
- Use in `main()`, initialization, and tests
- Never cancelled, no deadline, no values
- The starting point for all context trees

```go
ctx := context.Background()
```

### 2. context.TODO()

A placeholder context:
- Use when you're unsure which context to use
- Signals that context handling needs to be added later
- Useful during refactoring

```go
ctx := context.TODO()
```

### 3. Why Context Exists

Context solves critical problems in concurrent programs:

1. **Cancellation Propagation**: Stop goroutines when parent operations complete
2. **Deadlines and Timeouts**: Prevent operations from running too long
3. **Request-Scoped Values**: Pass request IDs, auth tokens, etc.
4. **Resource Management**: Ensure proper cleanup of resources

### 4. Function Signatures with Context

The idiomatic pattern in Go:

```go
func DoWork(ctx context.Context, args ...any) error {
    // Context is always the first parameter
    // Named 'ctx' by convention
}
```

**Convention Rules:**
- Context is the **first** parameter
- Named `ctx` by convention
- Not stored in structs (with rare exceptions)
- Pass through the entire call chain

### 5. Context is Immutable

Creating child contexts doesn't modify parents:

```go
parent := context.Background()
child := context.WithValue(parent, "key", "value")

parent.Value("key") // nil - parent unchanged
child.Value("key")  // "value" - child has the value
```

## Running the Example

```bash
go run main.go
```

## Expected Output

The program demonstrates:
1. Creating root contexts with Background() and TODO()
2. Passing context through function call chains
3. Proper function signatures with context
4. Inspecting context state
5. Context immutability
6. Real-world HTTP handler pattern

## Key Takeaways

1. **Always use context.Background() or context.TODO()** - Never pass nil context
2. **Context is the first parameter** - This is a strong Go convention
3. **Pass context explicitly** - Context flows through function calls
4. **Context is immutable** - Child contexts don't affect parents
5. **Don't store context** - Pass it as a parameter, not a struct field

## Common Mistakes to Avoid

```go
// ❌ WRONG: Storing context in a struct
type Server struct {
    ctx context.Context // DON'T DO THIS
}

// ✅ CORRECT: Pass context as parameter
func (s *Server) HandleRequest(ctx context.Context) error {
    // Use ctx as parameter
}

// ❌ WRONG: Passing nil context
result := doWork(nil, data) // DON'T DO THIS

// ✅ CORRECT: Use Background or TODO
result := doWork(context.Background(), data)
result := doWork(context.TODO(), data) // If refactoring

// ❌ WRONG: Context not first parameter
func doWork(data string, ctx context.Context) error

// ✅ CORRECT: Context is first parameter
func doWork(ctx context.Context, data string) error
```

## Real-World Usage

### HTTP Handler Example

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Get context from request
    ctx := r.Context()

    // Pass to business logic
    result, err := processRequest(ctx, r.FormValue("id"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(result)
}
```

### Database Query Example

```go
func fetchUser(ctx context.Context, db *sql.DB, userID int) (*User, error) {
    // Context passed to QueryRowContext for cancellation support
    row := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", userID)

    var user User
    if err := row.Scan(&user.ID, &user.Name); err != nil {
        return nil, err
    }
    return &user, nil
}
```

## Next Steps

Now that you understand context basics, move on to:
- **Context Cancellation**: Learn to cancel goroutines gracefully
- **Context Timeout**: Implement automatic timeouts
- **Context Values**: Pass request-scoped data

## Resources

- [Go Blog: Context](https://go.dev/blog/context)
- [context package documentation](https://pkg.go.dev/context)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
