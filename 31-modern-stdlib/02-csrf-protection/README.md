# CSRF Protection (Go 1.25)

Cross-Site Request Forgery (CSRF) protection using Go 1.25's new `net/http.CrossOriginProtection()` middleware.

## What is CSRF?

CSRF attacks trick users into performing unwanted actions on a web application where they're authenticated. 
An attacker can forge requests that appear to come from the legitimate user.

## Go 1.25 Solution

Go 1.25 introduces `http.CrossOriginProtection()` middleware that automatically protects against CSRF:

```go
handler := http.CrossOriginProtection()(mux)
```

### How It Works

The middleware:
1. Checks `Origin` and `Referer` headers on state-changing requests (POST, PUT, DELETE)
2. Ensures requests come from the same origin
3. Blocks cross-origin requests that could be CSRF attacks
4. Allows safe methods (GET, HEAD, OPTIONS) through

## This Example

This example demonstrates CSRF protection concepts. In production with Go 1.25+, use:

```go
import "net/http"

mux := http.NewServeMux()
// ... setup routes ...

// Apply built-in CSRF protection
protected := http.CrossOriginProtection()(mux)

http.ListenAndServe(":8080", protected)
```

## Testing CSRF Protection

### Same-Origin Request (Allowed)
```bash
curl -X POST http://localhost:8080/api/update \
  -H "Origin: http://localhost:8080"
```

### Cross-Origin Request (Blocked)
```bash
curl -X POST http://localhost:8080/api/update \
  -H "Origin: http://evil.com"
```

## Best Practices

1. **Always use CSRF protection** for state-changing endpoints
2. **Don't rely on cookies alone** for authentication
3. **Use HTTPS** in production
4. **Combine with other security measures** (auth tokens, rate limiting)

## Learn More

- [Go 1.25 Release Notes](https://go.dev/doc/go1.25)
- [OWASP CSRF Prevention](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html)

## Running This Example

```bash
go run main.go
```

Then in another terminal:
```bash
curl http://localhost:8080/
curl http://localhost:8080/api/data
curl -X POST http://localhost:8080/api/update
```
