package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// route struct stores compiled regexp so that
// compilation of regex does not become a bottleneck
// when many routes are present
type route struct {
	pattern *regexp.Regexp
	handler http.HandlerFunc
}

type Router struct {
	routes map[string][]route // Method -> []route (pattern + handler)
}

func New() *Router {
	return &Router{
		routes: make(map[string][]route),
	}
}

func normalizePath(path string) string {
	return strings.TrimSuffix(path, "/")
}

func (r *Router) AddRoute(method, path string, handler http.HandlerFunc) {
	// Normalize the path
	path = normalizePath(path)

	// Convert dynamic segments like `:name` to regex groups
	regexStr := "^" + regexp.MustCompile(`:([a-zA-Z_][a-zA-Z0-9_]*)`).ReplaceAllString(path, `(?P<$1>[^/]+)`) + "$"
	compiledRegex := regexp.MustCompile(regexStr)

	// Add the route
	r.routes[method] = append(r.routes[method], route{
		pattern: compiledRegex,
		handler: handler,
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := normalizePath(req.URL.Path)

	// Find routes for the HTTP method
	if routes, ok := r.routes[method]; ok {
		for _, rt := range routes {
			if rt.pattern.MatchString(path) {
				// Extract dynamic parameters if any
				params := extractParams(rt.pattern, path)
				if len(params) > 0 {
					// Add params to context
					ctx := context.WithValue(req.Context(), "params", params)
					req = req.WithContext(ctx)
				}
				// Call the handler
				rt.handler(w, req)
				return
			}
		}
	}

	http.NotFound(w, req)
}

func extractParams(pattern *regexp.Regexp, path string) map[string]string {
	params := make(map[string]string)
	matches := pattern.FindStringSubmatch(path)
	if len(matches) > 0 {
		for i, name := range pattern.SubexpNames() {
			if i > 0 && name != "" {
				params[name] = matches[i]
			}
		}
	}
	return params
}

func Start(port string, r *Router) {
	fmt.Printf("Hawkwing server is running on port %s\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
