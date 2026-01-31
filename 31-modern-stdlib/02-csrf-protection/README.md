# CSRF Protection with CrossOriginProtection

This example demonstrates Go 1.25's new `http.CrossOriginProtection` middleware for protecting against Cross-Site Request Forgery (CSRF) attacks without requiring token management.

## Overview

The `CrossOriginProtection` middleware provides:

- **Modern Approach**: Uses browser Fetch metadata headers (Sec-Fetch-Site)
- **No Token Management**: No need to generate, store, or validate CSRF tokens
- **Simple Integration**: Works as standard middleware with any `http.Handler`
- **Flexible Configuration**: Support for trusted origins and bypass patterns
- **Secure by Default**: Automatically protects state-changing requests

## Requirements

**Go Version**: Go 1.25 or later

## Building and Running

### Run the examples
```bash
go run main.go
```

### Run a real server (uncomment runServer in main.go)
```bash
go run main.go
# Visit http://localhost:8080 in your browser
```

## How It Works

### Detection Methods

CrossOriginProtection detects cross-origin requests using:

1. **Sec-Fetch-Site header** (available in all modern browsers since 2023)
   - `same-origin`: Request from the same origin (allowed)
   - `none`: Direct navigation/bookmark (allowed)
   - `cross-site`: Request from different origin (blocked unless trusted)

2. **Origin header** comparison with Host header
   - Compares the requesting origin with the server's host
   - Blocks if origins don't match (unless trusted)

### Safe Methods

GET, HEAD, and OPTIONS are always allowed because they should be read-only operations. Applications must not perform state-changing actions with these methods.

### Non-Browser Requests

If no Sec-Fetch-Site or Origin headers are present, the middleware assumes:
- The request is from a non-browser client (API client, curl, etc.)
- No CSRF protection is needed (these tools don't execute malicious scripts)

This allows API clients, mobile apps, and command-line tools to work without issues.

## Key Features Demonstrated

### 1. Basic Protection

The simplest usage wraps your handler:

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Your handler code
})

protected := http.NewCrossOriginProtection().Handler(handler)
```

**Behavior**:
- Same-origin POST/PUT/DELETE: Allowed
- Cross-origin POST/PUT/DELETE: Blocked (403 Forbidden)
- All GET/HEAD/OPTIONS: Allowed

### 2. Custom Deny Handler

Customize the response when requests are blocked:

```go
cop := http.NewCrossOriginProtection()
cop.SetDenyHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusForbidden)
    json.NewEncoder(w).Encode(map[string]string{
        "error": "CSRF protection triggered",
    })
}))

protected := cop.Handler(handler)
```

### 3. Trusted Origins

Allow specific cross-origin requests:

```go
cop := http.NewCrossOriginProtection()

// Add your trusted domains
cop.AddTrustedOrigin("https://app.example.com")
cop.AddTrustedOrigin("https://mobile.example.com")
cop.AddTrustedOrigin("https://admin.example.com:8443")

protected := cop.Handler(handler)
```

**Use Cases**:
- Mobile app communicating with your API
- Multiple subdomains of your application
- Partner applications with explicit trust

**Origin Format**: `"scheme://host[:port]"` (no path, query, or fragment)

### 4. Bypass Patterns

Disable CSRF protection for specific paths:

```go
cop := http.NewCrossOriginProtection()

// Allow webhooks from external services
cop.AddInsecureBypassPattern("/webhooks/")
cop.AddInsecureBypassPattern("/oauth/callback")

protected := cop.Handler(handler)
```

**Use Cases**:
- Webhook endpoints (GitHub, Stripe, etc.)
- OAuth callback URLs
- Public APIs that don't require CSRF protection

**Pattern Syntax**: Matches `http.ServeMux` pattern rules

**Security Note**: Use bypass patterns carefully! These endpoints will accept any cross-origin request.

### 5. Real-World API Example

Complete example of protecting a RESTful API:

```go
mux := http.NewServeMux()

// Safe read operations (no CSRF protection needed)
mux.HandleFunc("GET /api/users", getUsers)
mux.HandleFunc("GET /api/users/{id}", getUser)

// State-changing operations (CSRF protection needed)
mux.HandleFunc("POST /api/users", createUser)
mux.HandleFunc("PUT /api/users/{id}", updateUser)
mux.HandleFunc("DELETE /api/users/{id}", deleteUser)

// Setup protection
cop := http.NewCrossOriginProtection()
cop.AddTrustedOrigin("https://app.example.com")
cop.AddInsecureBypassPattern("/webhooks/")

protected := cop.Handler(mux)
http.ListenAndServe(":8080", protected)
```

## API Reference

### NewCrossOriginProtection

```go
func NewCrossOriginProtection() *CrossOriginProtection
```

Creates a new CrossOriginProtection value. The zero value is valid with no trusted origins or bypass patterns.

### AddTrustedOrigin

```go
func (c *CrossOriginProtection) AddTrustedOrigin(origin string) error
```

Allows all requests with an Origin header matching the given value.

**Parameters**:
- `origin`: String in format `"scheme://host[:port]"`

**Returns**: Error if origin is invalid

**Concurrency**: Safe to call concurrently

**Example**:
```go
cop.AddTrustedOrigin("https://example.com")
cop.AddTrustedOrigin("https://app.example.com:8443")
```

### AddInsecureBypassPattern

```go
func (c *CrossOriginProtection) AddInsecureBypassPattern(pattern string)
```

Permits all requests matching the given pattern, regardless of origin.

**Parameters**:
- `pattern`: Pattern matching `http.ServeMux` syntax

**Concurrency**: Safe to call concurrently

**Security**: Use with caution! Bypasses all CSRF protection.

**Example**:
```go
cop.AddInsecureBypassPattern("/webhooks/")
cop.AddInsecureBypassPattern("/public/api/")
```

### SetDenyHandler

```go
func (c *CrossOriginProtection) SetDenyHandler(h Handler)
```

Sets a custom handler to invoke when a request is rejected.

**Default**: Returns 403 Forbidden

**Example**:
```go
cop.SetDenyHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusForbidden)
    json.NewEncoder(w).Encode(map[string]string{
        "error": "CSRF validation failed",
    })
}))
```

### Handler

```go
func (c *CrossOriginProtection) Handler(h Handler) Handler
```

Returns a handler that applies cross-origin checks before invoking handler `h`.

**Example**:
```go
protected := cop.Handler(myHandler)
http.Handle("/api/", protected)
```

### Check

```go
func (c *CrossOriginProtection) Check(req *Request) error
```

Applies cross-origin checks to a request without handling it.

**Returns**: Error if the request should be rejected

**Use Case**: Manual checking in custom middleware

## Best Practices

### 1. Apply to State-Changing Endpoints

```go
// Good: Protect POST, PUT, DELETE
cop := http.NewCrossOriginProtection()
protected := cop.Handler(apiMux)

// GET endpoints are automatically safe
// (assuming they don't change state)
```

### 2. Use Trusted Origins for Your Domains

```go
cop := http.NewCrossOriginProtection()

// Add all your legitimate frontend domains
cop.AddTrustedOrigin("https://example.com")
cop.AddTrustedOrigin("https://www.example.com")
cop.AddTrustedOrigin("https://app.example.com")
cop.AddTrustedOrigin("https://mobile.example.com")
```

### 3. Bypass Only When Necessary

```go
// Good: Specific webhook endpoints
cop.AddInsecureBypassPattern("/webhooks/github")
cop.AddInsecureBypassPattern("/webhooks/stripe")

// Bad: Broad bypass (dangerous!)
// cop.AddInsecureBypassPattern("/api/")
```

### 4. Implement Custom Error Responses

```go
cop.SetDenyHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("CSRF blocked: %s %s from %s",
        r.Method, r.URL.Path, r.Header.Get("Origin"))

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusForbidden)
    json.NewEncoder(w).Encode(map[string]string{
        "error": "Request forbidden",
        "code":  "CSRF_VALIDATION_FAILED",
    })
}))
```

### 5. Never Modify State with GET

```go
// Bad: State change with GET
mux.HandleFunc("GET /api/users/{id}/delete", deleteUser)

// Good: Use DELETE method
mux.HandleFunc("DELETE /api/users/{id}", deleteUser)
```

## Comparison: Token-Based vs CrossOriginProtection

| Aspect | Token-Based CSRF | CrossOriginProtection |
|--------|------------------|------------------------|
| Token Generation | Required | Not needed |
| Token Storage | Required (session/cookie) | Not needed |
| Token Validation | Manual implementation | Automatic |
| Browser Support | All browsers | Modern browsers (2023+) |
| Implementation | Complex | Simple |
| State Management | Required | Stateless |
| API Clients | Need special handling | Work automatically |

## Browser Support

CrossOriginProtection requires browsers that support Sec-Fetch-Site headers:
- Chrome 76+ (July 2019)
- Edge 79+ (January 2020)
- Firefox 90+ (July 2021)
- Safari 15.5+ (May 2022)

**All major browsers since 2023** support this feature.

## Testing with Browser

Create a test HTML page:

```html
<!DOCTYPE html>
<html>
<head><title>CSRF Test</title></head>
<body>
    <h1>CSRF Protection Test</h1>

    <button onclick="testSameOrigin()">
        Test Same-Origin POST (Should Work)
    </button>

    <button onclick="testCrossOrigin()">
        Test Cross-Origin POST (Should Fail)
    </button>

    <div id="result"></div>

    <script>
    async function testSameOrigin() {
        const resp = await fetch('/api/users', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({name: 'Alice'})
        });
        showResult(resp);
    }

    async function testCrossOrigin() {
        // This would need to be from a different origin
        // Open DevTools to see Sec-Fetch-Site header
        const resp = await fetch('http://different-origin.com/api/users', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({name: 'Bob'})
        });
        showResult(resp);
    }

    async function showResult(resp) {
        const data = await resp.json();
        document.getElementById('result').textContent =
            `Status: ${resp.status}\nBody: ${JSON.stringify(data, null, 2)}`;
    }
    </script>
</body>
</html>
```

## Security Considerations

### What CrossOriginProtection Protects Against

- Traditional CSRF attacks from malicious websites
- Cross-origin form submissions
- Cross-origin fetch/AJAX requests
- Cross-origin iframe exploits

### What It Doesn't Protect Against

- Same-origin XSS attacks (use Content Security Policy)
- Clickjacking (use X-Frame-Options or CSP frame-ancestors)
- Token theft via XSS (use HttpOnly cookies)
- Replay attacks (implement nonces or short-lived tokens)

### Defense in Depth

CrossOriginProtection should be part of a comprehensive security strategy:

1. **Input Validation**: Validate and sanitize all inputs
2. **Output Encoding**: Prevent XSS attacks
3. **Authentication**: Verify user identity
4. **Authorization**: Check user permissions
5. **HTTPS**: Encrypt all traffic
6. **CSP**: Content Security Policy headers
7. **CSRF**: CrossOriginProtection middleware

## Common Issues

### 403 Forbidden on Legitimate Requests

**Cause**: Frontend domain not in trusted origins

**Solution**:
```go
cop.AddTrustedOrigin("https://your-frontend.com")
```

### API Clients Failing

**Cause**: API client might be setting unexpected headers

**Solution**: API clients without Sec-Fetch-Site headers work automatically. If issues persist, check for custom header injection.

### Webhooks Not Working

**Cause**: CSRF protection blocking external services

**Solution**:
```go
cop.AddInsecureBypassPattern("/webhooks/")
```

## Resources

- [Go 1.25 Release Notes](https://go.dev/doc/go1.25)
- [CSRF Protection Source Code](https://go.dev/src/net/http/csrf.go)
- [Sec-Fetch-Site Specification](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Site)
- [OWASP CSRF Prevention](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html)

## Next Steps

After exploring CSRF protection:
1. Identify state-changing endpoints in your application
2. Apply CrossOriginProtection to your HTTP handlers
3. Add trusted origins for your frontend domains
4. Test with browser DevTools to verify headers
5. Configure bypass patterns for webhooks/public APIs
6. Implement custom deny handlers for better error messages
