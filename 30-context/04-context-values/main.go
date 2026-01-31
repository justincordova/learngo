package main

import (
	"context"
	"fmt"
	"time"
)

// Define type-safe keys to avoid collisions
type contextKey string

const (
	requestIDKey  contextKey = "requestID"
	userIDKey     contextKey = "userID"
	correlationID contextKey = "correlationID"
)

// Alternative: use unexported struct type for even stronger type safety
type requestContextKey struct{}

func main() {
	fmt.Println("Context Values and Best Practices")
	fmt.Println("==================================")
	fmt.Println()

	// Example 1: Basic context values with WithValue
	fmt.Println("1. Basic context values:")
	example1BasicValues()
	fmt.Println()

	// Example 2: Type-safe keys using custom types
	fmt.Println("2. Type-safe keys (RECOMMENDED):")
	example2TypeSafeKeys()
	fmt.Println()

	// Example 3: Request-scoped values (request ID, correlation ID)
	fmt.Println("3. Request-scoped values:")
	example3RequestScoped()
	fmt.Println()

	// Example 4: Authentication/authorization data
	fmt.Println("4. Authentication data in context:")
	example4AuthData()
	fmt.Println()

	// Example 5: Value propagation through call chain
	fmt.Println("5. Value propagation:")
	example5ValuePropagation()
	fmt.Println()

	// Example 6: ANTI-PATTERN - Don't use for optional parameters
	fmt.Println("6. ANTI-PATTERN - Optional parameters (DON'T DO THIS):")
	example6AntiPatternOptionalParams()
	fmt.Println()

	// Example 7: ANTI-PATTERN - Don't use for configuration
	fmt.Println("7. ANTI-PATTERN - Configuration (DON'T DO THIS):")
	example7AntiPatternConfig()
	fmt.Println()

	// Example 8: Good use cases summary
	fmt.Println("8. When to use context values (SUMMARY):")
	example8GoodUseCases()
	fmt.Println()
}

// example1BasicValues demonstrates basic context.WithValue usage
func example1BasicValues() {
	// Create context with a value
	ctx := context.WithValue(context.Background(), "key", "value")

	// Retrieve the value
	value := ctx.Value("key")
	if value != nil {
		fmt.Printf("   Found value: %v\n", value)
	}

	// Non-existent key returns nil
	missing := ctx.Value("nonexistent")
	fmt.Printf("   Missing key returns: %v\n", missing)

	fmt.Println()
	fmt.Println("   WARNING: Using string keys is NOT recommended!")
	fmt.Println("   They can collide with keys from other packages.")
	fmt.Println("   See next example for type-safe approach.")
}

// example2TypeSafeKeys demonstrates the recommended approach
func example2TypeSafeKeys() {
	// Using custom type for keys prevents collisions
	ctx := context.WithValue(context.Background(), requestIDKey, "req-12345")
	ctx = context.WithValue(ctx, userIDKey, "user-67890")

	// Retrieve values using type-safe keys
	reqID := ctx.Value(requestIDKey)
	userID := ctx.Value(userIDKey)

	fmt.Printf("   Request ID: %v\n", reqID)
	fmt.Printf("   User ID: %v\n", userID)
	fmt.Println()
	fmt.Println("   Benefits of custom types:")
	fmt.Println("   - No collisions with other packages")
	fmt.Println("   - Clear intent and documentation")
	fmt.Println("   - Type safety at compile time")
}

// example3RequestScoped demonstrates request-scoped metadata
func example3RequestScoped() {
	// Simulating HTTP request with request ID and correlation ID
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "req-abc123")
	ctx = context.WithValue(ctx, correlationID, "corr-xyz789")

	fmt.Println("   Incoming request:")
	logRequest(ctx, "Processing user login")

	// Simulate calling another service
	fmt.Println()
	callExternalService(ctx)
}

// logRequest demonstrates using context values for logging
func logRequest(ctx context.Context, message string) {
	reqID := ctx.Value(requestIDKey)
	corrID := ctx.Value(correlationID)

	fmt.Printf("   [RequestID: %v] [CorrelationID: %v] %s\n", reqID, corrID, message)
}

// callExternalService shows values propagating to external calls
func callExternalService(ctx context.Context) {
	logRequest(ctx, "Calling external authentication service")
	time.Sleep(50 * time.Millisecond)
	logRequest(ctx, "External service responded")
}

// User represents an authenticated user
type User struct {
	ID       string
	Username string
	Role     string
}

// userKey is an unexported type for stronger type safety
type userKey struct{}

// example4AuthData demonstrates storing auth information in context
func example4AuthData() {
	// After authentication, store user in context
	user := &User{
		ID:       "user-123",
		Username: "alice",
		Role:     "admin",
	}

	ctx := context.WithValue(context.Background(), userKey{}, user)

	// Pass context through application layers
	handleAuthenticatedRequest(ctx)
}

// handleAuthenticatedRequest shows using auth data from context
func handleAuthenticatedRequest(ctx context.Context) {
	// Extract user from context
	user, ok := ctx.Value(userKey{}).(*User)
	if !ok {
		fmt.Println("   Error: No authenticated user in context")
		return
	}

	fmt.Printf("   Processing request for user: %s (ID: %s)\n", user.Username, user.ID)

	// Check authorization
	if user.Role == "admin" {
		fmt.Println("   User has admin privileges")
		performAdminAction(ctx)
	} else {
		fmt.Println("   User has standard privileges")
	}
}

// performAdminAction demonstrates accessing user info deep in call chain
func performAdminAction(ctx context.Context) {
	user, ok := ctx.Value(userKey{}).(*User)
	if !ok || user.Role != "admin" {
		fmt.Println("   Unauthorized: Admin access required")
		return
	}

	fmt.Printf("   Admin action performed by: %s\n", user.Username)
}

// example5ValuePropagation shows how values flow through contexts
func example5ValuePropagation() {
	// Parent context with value
	parent := context.WithValue(context.Background(), requestIDKey, "req-parent")

	// Child context inherits parent values
	child := context.WithValue(parent, userIDKey, "user-child")

	// Grandchild context inherits both
	grandchild := context.WithValue(child, correlationID, "corr-grandchild")

	// All values are accessible from grandchild
	fmt.Printf("   From grandchild context:\n")
	fmt.Printf("   - Request ID: %v (from parent)\n", grandchild.Value(requestIDKey))
	fmt.Printf("   - User ID: %v (from child)\n", grandchild.Value(userIDKey))
	fmt.Printf("   - Correlation ID: %v (from grandchild)\n", grandchild.Value(correlationID))
	fmt.Println()
	fmt.Println("   Note: Child contexts inherit all parent values")
}

// ANTI-PATTERN EXAMPLES BELOW
// These demonstrate what NOT to do with context values

// example6AntiPatternOptionalParams shows WRONG use for optional parameters
func example6AntiPatternOptionalParams() {
	fmt.Println("   BAD: Using context for optional function parameters")
	fmt.Println()

	// WRONG: Hiding function parameters in context
	ctx := context.WithValue(context.Background(), "timeout", 5*time.Second)
	ctx = context.WithValue(ctx, "retries", 3)
	badFetchData(ctx, "https://api.example.com")

	fmt.Println()
	fmt.Println("   WHY THIS IS BAD:")
	fmt.Println("   - Function signature doesn't show what options exist")
	fmt.Println("   - No type safety - easy to pass wrong types")
	fmt.Println("   - Hard to discover what parameters are available")
	fmt.Println("   - Testing becomes difficult")
	fmt.Println()
	fmt.Println("   CORRECT APPROACH:")
	fmt.Println("   Use explicit function parameters or option structs:")
	fmt.Println("   func fetchData(ctx context.Context, url string, opts Options) error")
}

// badFetchData demonstrates the anti-pattern
func badFetchData(ctx context.Context, url string) {
	// BAD: Extracting "parameters" from context
	timeout := ctx.Value("timeout").(time.Duration)
	retries := ctx.Value("retries").(int)

	fmt.Printf("   Fetching %s with timeout=%v, retries=%d\n", url, timeout, retries)
	fmt.Println("   ^ This makes the function signature misleading!")
}

// example7AntiPatternConfig shows WRONG use for configuration
func example7AntiPatternConfig() {
	fmt.Println("   BAD: Using context for application configuration")
	fmt.Println()

	// WRONG: Storing configuration in context
	ctx := context.WithValue(context.Background(), "dbHost", "localhost")
	ctx = context.WithValue(ctx, "dbPort", 5432)
	ctx = context.WithValue(ctx, "logLevel", "debug")

	badInitializeApp(ctx)

	fmt.Println()
	fmt.Println("   WHY THIS IS BAD:")
	fmt.Println("   - Configuration is application-wide, not request-scoped")
	fmt.Println("   - Makes dependencies invisible and hard to test")
	fmt.Println("   - Violates principle of explicit dependencies")
	fmt.Println("   - No compile-time type checking")
	fmt.Println()
	fmt.Println("   CORRECT APPROACH:")
	fmt.Println("   Use a Config struct passed to constructors:")
	fmt.Println("   func NewDatabase(config *DBConfig) (*DB, error)")
	fmt.Println("   Or use dependency injection")
}

// badInitializeApp demonstrates the anti-pattern
func badInitializeApp(ctx context.Context) {
	// BAD: Pulling configuration from context
	host := ctx.Value("dbHost").(string)
	port := ctx.Value("dbPort").(int)
	logLevel := ctx.Value("logLevel").(string)

	fmt.Printf("   Initializing app with host=%s, port=%d, log=%s\n", host, port, logLevel)
	fmt.Println("   ^ Configuration should be passed explicitly, not through context!")
}

// example8GoodUseCases summarizes when to use context values
func example8GoodUseCases() {
	fmt.Println("   GOOD use cases for context values:")
	fmt.Println()
	fmt.Println("   ✓ Request-scoped data:")
	fmt.Println("     - Request IDs for tracing")
	fmt.Println("     - Correlation IDs for distributed tracing")
	fmt.Println("     - Request start time")
	fmt.Println()
	fmt.Println("   ✓ Authentication/Authorization:")
	fmt.Println("     - Authenticated user information")
	fmt.Println("     - API tokens (for propagation to downstream services)")
	fmt.Println("     - Session data")
	fmt.Println()
	fmt.Println("   ✓ Request-specific metadata:")
	fmt.Println("     - Client IP address")
	fmt.Println("     - User agent")
	fmt.Println("     - Request locale/language")
	fmt.Println()
	fmt.Println("   ✗ AVOID using context values for:")
	fmt.Println()
	fmt.Println("   ✗ Optional function parameters")
	fmt.Println("     - Use function parameters or option structs instead")
	fmt.Println()
	fmt.Println("   ✗ Application configuration")
	fmt.Println("     - Use config structs and dependency injection")
	fmt.Println()
	fmt.Println("   ✗ Passing dependencies")
	fmt.Println("     - Pass dependencies explicitly as function parameters")
	fmt.Println()
	fmt.Println("   ✗ Data that should be in function signatures")
	fmt.Println("     - If it's required for the function, make it a parameter")
	fmt.Println()
	fmt.Println("   RULE OF THUMB:")
	fmt.Println("   Use context.Value for request-scoped data that crosses")
	fmt.Println("   API boundaries and would otherwise require threading through")
	fmt.Println("   many function signatures. Don't use it to hide function parameters.")
}
