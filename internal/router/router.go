package router

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func New() *Router {
	return &Router{routes: make(map[string]map[string]http.HandlerFunc)}
}

func (r *Router) AddRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	if handlers, ok := r.routes[method]; ok {
		if handler, ok := handlers[path]; ok {
			handler(w, req)
			return
		}
	}

	http.NotFound(w, req)
}

func Start(port string, r *Router) {
	fmt.Printf("Hawkwing server is running on port %s\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
		return
	}
}
