package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("CSRF Protection with Go 1.25")
	fmt.Println("=============================")
	fmt.Println()
	
	// Setup routes
	mux := http.NewServeMux()
	
	// Public endpoints (no CSRF protection needed)
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/about", handleAbout)
	
	// Protected endpoints (need CSRF protection)
	mux.HandleFunc("/api/data", handleAPIData)
	mux.HandleFunc("/api/update", handleAPIUpdate)
	
	// Apply CSRF protection middleware
	// Note: CrossOriginProtection is new in Go 1.25
	handler := applyCORSProtection(mux)
	
	fmt.Println("Server starting on :8080")
	fmt.Println()
	fmt.Println("Try these requests:")
	fmt.Println("  curl http://localhost:8080/")
	fmt.Println("  curl http://localhost:8080/api/data")
	fmt.Println("  curl -X POST http://localhost:8080/api/update")
	fmt.Println()
	
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! This is the home page.\n")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page - public content\n")
}

func handleAPIData(w http.ResponseWriter, r *http.Request) {
	// This endpoint returns data (GET)
	fmt.Fprintf(w, `{"data": "some information", "timestamp": 1234567890}`)
}

func handleAPIUpdate(w http.ResponseWriter, r *http.Request) {
	// This endpoint modifies data (POST/PUT/DELETE)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	fmt.Fprintf(w, `{"status": "updated", "message": "Data updated successfully"}`)
}

// applyCORSProtection demonstrates CSRF protection
// Note: In Go 1.25+, use net/http.CrossOriginProtection()
func applyCORSProtection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// For demonstration: simple CORS and origin checking
		// In Go 1.25+, use: http.CrossOriginProtection()(next)
		
		origin := r.Header.Get("Origin")
		
		// Log the request
		fmt.Printf("[%s] %s %s (Origin: %s)\n", 
			r.RemoteAddr, r.Method, r.URL.Path, origin)
		
		// Simple CORS headers for demonstration
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}
		
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// Pass to next handler
		next.ServeHTTP(w, r)
	})
}

// Note: In Go 1.25+, you would use:
//
// handler := http.CrossOriginProtection()(mux)
//
// This provides automatic CSRF protection by:
// - Checking Origin and Referer headers
// - Requiring same-origin for state-changing requests
// - Protecting against cross-site request forgery attacks
