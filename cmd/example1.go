package main

import (
	"fmt"
	"github.com/aliqyan-21/hawkwing/internal/router"
	"net/http"
)

func main() {
	r := router.New()

	r.AddRoute("GET", "/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Welcome to Hawkwing!")
	})

	r.AddRoute("POST", "/users", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "User created!")
	})

	router.Start(":5000", r)
}
