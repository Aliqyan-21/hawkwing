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

// responseWriter is a custom HTTP response writer that captures the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Logger is a middleware that logs incoming HTTP requests to the standard logger.
// It logs the request method and URL path. (default for all routes)
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request received: %s %s", req.Method, req.URL.Path)

		ww := &responseWriter{ResponseWriter: w}
		next(ww, req)

		log.Printf("Response: %d %s", ww.statusCode, http.StatusText(ww.statusCode))
	}
}

// WriteHeader captures the status code and passes it to the original WriteHeader method.
// Automatically called through the http.ResponseWriter interface whenever a status code is set.
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
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

// ErrorHandler is a middleware that recovers from panics, logs the error, and returns an internal server error response. (default for all routes)
func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next(w, req)
	}
}
