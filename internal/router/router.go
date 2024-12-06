package router

import (
	"context"
	"fmt"
	"github.com/aliqyan-21/hawkwing/internal/middleware"
	"github.com/aliqyan-21/hawkwing/internal/static"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	pattern     *regexp.Regexp
	handler     http.HandlerFunc
	middlewares []middleware.Middleware
}

type Router struct {
	routes map[string][]route
}

func New() *Router {
	return &Router{
		routes: make(map[string][]route),
	}
}

func normalizePath(path string) string {
	return strings.TrimSuffix(path, "/")
}

func (r *Router) AddRoute(method, path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	path = normalizePath(path)

	regexStr := "^" + regexp.MustCompile(`:([a-zA-Z_][a-zA-Z0-9_]*)`).ReplaceAllString(path, `(?P<$1>[^/]+)`) + "$"
	compiledRegex := regexp.MustCompile(regexStr)

	middlewares = append([]middleware.Middleware{middleware.Logger}, middlewares...)

	r.routes[method] = append(r.routes[method], route{
		pattern:     compiledRegex,
		handler:     handler,
		middlewares: middlewares,
	})
}

func (r *Router) GetRouteParams(req *http.Request) map[string]string {
	if params, ok := req.Context().Value("params").(map[string]string); ok {
		return params
	}
	return nil
}

func (r *Router) dynamicParams(pathPattern, actualPath string) map[string]string {
	params := make(map[string]string)
	re := regexp.MustCompile(pathPattern)
	matches := re.FindStringSubmatch(actualPath)
	if len(matches) > 0 {
		for i, name := range re.SubexpNames() {
			if i > 0 && name != "" {
				params[name] = matches[i]
			}
		}
	}
	return params
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := normalizePath(req.URL.Path)

	if routes, ok := r.routes[method]; ok {
		for _, rt := range routes {
			if rt.pattern.MatchString(path) {
				handler := rt.handler
				for _, mw := range rt.middlewares {
					handler = mw(handler)
				}

				params := extractParams(rt.pattern, path)
				if len(params) > 0 {
					ctx := context.WithValue(req.Context(), "params", params)
					req = req.WithContext(ctx)
				}
				handler(w, req)
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

func (r *Router) LoadStatic(routePath, dir string) {
	r.AddRoute("GET", routePath+"(.*)", static.LoadStatic(routePath, dir))
}

func Start(port string, r *Router) {
	fmt.Printf("Hawkwing server is running on port %s\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
