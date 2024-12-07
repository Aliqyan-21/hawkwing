// Package middleware provides a collection of reusable middleware functions for common HTTP request processing tasks.
// These middleware can be used to log requests, authenticate users, validate content types, and more.
// By using middleware, you can modularize your application's logic and improve its security, performance, and overall robustness.
package middleware

import (
	"log"
	"net/http"
)

// Middleware represents a function that takes an HTTP handler function and returns a new handler function.
// Middlewares are used to modify or extend the behavior of the original handler.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logger is a middleware that logs incoming HTTP requests to the standard logger.
// It logs the request method and URL path. It is the part of handler function by default
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request received: %s %s", req.Method, req.URL.Path)
		next(w, req)
	}
}

// Auth is a middleware that checks for a valid authorization token in the request header.
// If the token is missing or invalid, it returns a 403 Forbidden response.
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			log.Println("Unauthorized request")
			return
		}

		next(w, req)
	}
}

// ContentType is a middleware that checks the request's Content-Type header.
// It ensures that the Content-Type matches the specified value. If not, it returns a 400 Bad Request response.
func ContentType(contentTypes []string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")
		for _, ct := range contentTypes {
			if contentType == ct {
				next(w, req)
				return
			}
		}
		http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
		log.Printf("Invalid Content-Type: %s", contentType)
	}
}
