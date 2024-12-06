package main

import (
	"fmt"
	"github.com/aliqyan-21/hawkwing/internal/middleware"
	"github.com/aliqyan-21/hawkwing/internal/router"
	"net/http"
)

func main() {
	r := router.New()

	// Static route with middleware
	r.AddRoute("GET", "/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}, middleware.Logger)

	// Dynamic route with middleware
	r.AddRoute("GET", "/users/:name", func(w http.ResponseWriter, req *http.Request) {
		params := req.Context().Value("params").(map[string]string)
		name := params["name"]
		fmt.Fprintf(w, "Hello, %s!", name)
	}, middleware.Logger)

	// Another static route
	r.AddRoute("GET", "/about", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "This is the about page.")
	})

	router.Start(":8080", r)
}
