// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

// User represents a user in our system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	fmt.Println("CSRF Protection with CrossOriginProtection")
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("Go 1.25 introduces http.CrossOriginProtection middleware")
	fmt.Println("for protecting against Cross-Site Request Forgery attacks.")
	fmt.Println()

	// Example 1: Basic CSRF protection
	example1()

	// Example 2: Custom deny handler
	example2()

	// Example 3: Trusted origins
	example3()

	// Example 4: Bypass patterns
	example4()

	// Example 5: Real-world API example
	example5()

	fmt.Println("Note: Run with a real server to test with browser Fetch API")
	fmt.Println("Example server code is provided in the README.md")
}

func example1() {
	fmt.Println("1. Basic CSRF Protection")
	fmt.Println("------------------------")

	// Create a simple handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Response{
			Success: true,
			Message: "Request processed successfully",
		})
	})

	// Wrap with CSRF protection
	protected := http.NewCrossOriginProtection().Handler(handler)

	// Test 1: Same-origin request (allowed)
	req := httptest.NewRequest("POST", "/api/users", nil)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	w := httptest.NewRecorder()
	protected.ServeHTTP(w, req)
	fmt.Printf("Same-origin POST: %d %s\n", w.Code, w.Body.String())

	// Test 2: Cross-origin request (blocked)
	req = httptest.NewRequest("POST", "/api/users", nil)
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Origin", "https://evil.com")
	w = httptest.NewRecorder()
	protected.ServeHTTP(w, req)
	fmt.Printf("Cross-origin POST: %d (blocked)\n\n", w.Code)
}

func example2() {
	fmt.Println("2. Custom Deny Handler")
	fmt.Println("----------------------")

	// Create a handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Response{Success: true})
	})

	// Create protection with custom deny handler
	cop := http.NewCrossOriginProtection()
	cop.SetDenyHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "CSRF protection triggered: cross-origin request denied",
		})
	}))

	protected := cop.Handler(handler)

	// Test cross-origin request
	req := httptest.NewRequest("POST", "/api/delete", nil)
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	w := httptest.NewRecorder()
	protected.ServeHTTP(w, req)
	fmt.Printf("Status: %d\n", w.Code)
	fmt.Printf("Response: %s\n\n", w.Body.String())
}

func example3() {
	fmt.Println("3. Trusted Origins")
	fmt.Println("------------------")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Response{
			Success: true,
			Message: "Data updated",
		})
	})

	// Create protection with trusted origins
	cop := http.NewCrossOriginProtection()

	// Add trusted origins (e.g., your mobile app domain)
	err := cop.AddTrustedOrigin("https://app.example.com")
	if err != nil {
		log.Printf("Error adding trusted origin: %v\n", err)
	}

	err = cop.AddTrustedOrigin("https://mobile.example.com")
	if err != nil {
		log.Printf("Error adding trusted origin: %v\n", err)
	}

	protected := cop.Handler(handler)

	// Test 1: Request from trusted origin (allowed)
	req := httptest.NewRequest("POST", "/api/update", nil)
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Origin", "https://app.example.com")
	req.Host = "api.example.com"
	w := httptest.NewRecorder()
	protected.ServeHTTP(w, req)
	fmt.Printf("Trusted origin (app.example.com): %d %s\n", w.Code, w.Body.String())

	// Test 2: Request from untrusted origin (blocked)
	req = httptest.NewRequest("POST", "/api/update", nil)
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Origin", "https://untrusted.com")
	req.Host = "api.example.com"
	w = httptest.NewRecorder()
	protected.ServeHTTP(w, req)
	fmt.Printf("Untrusted origin (untrusted.com): %d (blocked)\n\n", w.Code)
}

func example4() {
	fmt.Println("4. Bypass Patterns")
	fmt.Println("------------------")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Response{
			Success: true,
			Message: fmt.Sprintf("Processed: %s", r.URL.Path),
		})
	})

	// Create protection with bypass patterns
	cop := http.NewCrossOriginProtection()

	// Allow webhooks from external services without CSRF protection
	cop.AddInsecureBypassPattern("/webhooks/")
	cop.AddInsecureBypassPattern("/oauth/callback")

	protected := cop.Handler(handler)

	// Test 1: Webhook endpoint (bypasses CSRF protection)
	req := httptest.NewRequest("POST", "/webhooks/github", nil)
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Origin", "https://github.com")
	w := httptest.NewRecorder()
	protected.ServeHTTP(w, req)
	fmt.Printf("Webhook endpoint: %d (bypassed)\n", w.Code)

	// Test 2: OAuth callback (bypasses CSRF protection)
	req = httptest.NewRequest("POST", "/oauth/callback", nil)
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	w = httptest.NewRecorder()
	protected.ServeHTTP(w, req)
	fmt.Printf("OAuth callback: %d (bypassed)\n", w.Code)

	// Test 3: Regular endpoint (CSRF protection applies)
	req = httptest.NewRequest("POST", "/api/users", nil)
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	w = httptest.NewRecorder()
	protected.ServeHTTP(w, req)
	fmt.Printf("Regular endpoint: %d (blocked)\n\n", w.Code)
}

func example5() {
	fmt.Println("5. Real-World API Example")
	fmt.Println("-------------------------")

	// Create a RESTful API mux
	mux := http.NewServeMux()

	// Safe methods (no CSRF protection needed)
	mux.HandleFunc("GET /api/users", handleGetUsers)
	mux.HandleFunc("GET /api/users/{id}", handleGetUser)

	// State-changing methods (need CSRF protection)
	mux.HandleFunc("POST /api/users", handleCreateUser)
	mux.HandleFunc("PUT /api/users/{id}", handleUpdateUser)
	mux.HandleFunc("DELETE /api/users/{id}", handleDeleteUser)

	// Public webhook endpoint
	mux.HandleFunc("POST /webhooks/payment", handleWebhook)

	// Setup CSRF protection
	cop := http.NewCrossOriginProtection()

	// Add trusted frontend domains
	cop.AddTrustedOrigin("https://app.example.com")
	cop.AddTrustedOrigin("https://admin.example.com")

	// Bypass CSRF for webhooks
	cop.AddInsecureBypassPattern("/webhooks/")

	// Custom deny handler
	cop.SetDenyHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "CSRF validation failed",
		})
	}))

	// Wrap the mux with CSRF protection
	protected := cop.Handler(mux)

	// Simulate requests
	testRequests := []struct {
		method string
		path   string
		origin string
		site   string
	}{
		{"GET", "/api/users", "", "same-origin"},
		{"POST", "/api/users", "https://app.example.com", "cross-site"},
		{"POST", "/api/users", "https://evil.com", "cross-site"},
		{"POST", "/webhooks/payment", "https://payment-provider.com", "cross-site"},
	}

	for _, test := range testRequests {
		req := httptest.NewRequest(test.method, test.path, nil)
		if test.site != "" {
			req.Header.Set("Sec-Fetch-Site", test.site)
		}
		if test.origin != "" {
			req.Header.Set("Origin", test.origin)
		}
		w := httptest.NewRecorder()
		protected.ServeHTTP(w, req)

		var resp Response
		json.NewDecoder(w.Body).Decode(&resp)

		status := "allowed"
		if w.Code == 403 {
			status = "blocked"
		}
		fmt.Printf("%s %s (origin: %s): %s\n",
			test.method, test.path,
			truncate(test.origin, 25),
			status)
	}
	fmt.Println()
}

// Handler implementations
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	json.NewEncoder(w).Encode(Response{Success: true, Data: users})
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: fmt.Sprintf("User %s", id),
		Data:    user,
	})
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var user User
	json.Unmarshal(body, &user)
	user.ID = 3
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "User created",
		Data:    user,
	})
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: fmt.Sprintf("User %s updated", id),
	})
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: fmt.Sprintf("User %s deleted", id),
	})
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Webhook received",
	})
}

func truncate(s string, maxLen int) string {
	if s == "" {
		return "none"
	}
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// Example of running a real server (commented out for demo purposes)
func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
<!DOCTYPE html>
<html>
<head><title>CSRF Protection Demo</title></head>
<body>
	<h1>CSRF Protection Demo</h1>
	<button onclick="testSafe()">Test Safe Request (GET)</button>
	<button onclick="testCrossOrigin()">Test Cross-Origin POST</button>
	<div id="result"></div>

	<script>
	async function testSafe() {
		const resp = await fetch('/api/users');
		const data = await resp.json();
		document.getElementById('result').textContent = JSON.stringify(data);
	}

	async function testCrossOrigin() {
		const resp = await fetch('/api/users', {
			method: 'POST',
			headers: {'Content-Type': 'application/json'},
			body: JSON.stringify({name: 'Test', email: 'test@example.com'})
		});
		const data = await resp.json();
		document.getElementById('result').textContent = JSON.stringify(data);
	}
	</script>
</body>
</html>`)
	})

	mux.HandleFunc("GET /api/users", handleGetUsers)
	mux.HandleFunc("POST /api/users", handleCreateUser)

	cop := http.NewCrossOriginProtection()
	cop.AddTrustedOrigin("http://localhost:8080")

	protected := cop.Handler(mux)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", protected))
}
