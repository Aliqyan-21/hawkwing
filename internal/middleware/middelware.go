package middleware

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request received: %s %s", req.Method, req.URL.Path)
		next(w, req)
	}
}

func Authenticator(next http.HandlerFunc) http.HandlerFunc {
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

func CheckContentType(contentType string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Content-Type") != contentType {
			http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
			log.Printf("Invalid Content-Type: %s", req.Header.Get("Content-Type"))
			return
		}
		next(w, req)
	}
}
