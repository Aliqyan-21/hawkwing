// Package router provides a simple and efficient HTTP router for building web applications.
package router

import (
	"context"
	"fmt"
	"github.com/aliqyan-21/hawkwing/internal/middleware"
	"github.com/aliqyan-21/hawkwing/internal/static"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
)

// route represents a single route in the router, including its pattern, handler function, and associated middlewares.
type route struct {
	pattern     *regexp.Regexp          // Compiled regex storing so that bottleneck is handled and for matching the route
	handler     http.HandlerFunc        // Function to handle requests for this route
	middlewares []middleware.Middleware // List of middlewares for this route
}

// Router is the main structure responsible for holding all registered routes organized by HTTP method.
type Router struct {
	routes map[string][]route
}

// Init initializes a new Router instance with an empty route map.
func Init() *Router {
	return &Router{
		routes: make(map[string][]route),
	}
}

// normalizePath trims trailing slashes from the path to ensure consistent routing behavior.
func normalizePath(path string) string {
	return strings.TrimSuffix(path, "/")
}

// AddRoute registers a new route with the specified method, path, handler function, and optional middlewares.
func (r *Router) AddRoute(method, path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	path = normalizePath(path)

	// Convert path parameters (e.g., :id, :name, :hawk, :wing) into named regex groups for matching
	regexStr := "^" + regexp.MustCompile(`:([a-zA-Z_][a-zA-Z0-9_]*)`).ReplaceAllString(path, `(?P<$1>[^/]+)`) + "$"
	compiledRegex := regexp.MustCompile(regexStr)

	// Prepend logger middleware to the list of middlewares as it is default
	middlewares = append([]middleware.Middleware{middleware.Logger, middleware.ErrorHandler}, middlewares...)

	r.routes[method] = append(r.routes[method], route{
		pattern:     compiledRegex,
		handler:     handler,
		middlewares: middlewares,
	})
}

// GetRouteParams retrieves dynamic path parameters from the request context.
// This assumes that the parameters were previously extracted and stored in the context.
func (r *Router) GetRouteParams(req *http.Request) map[string]string {
	if params, ok := req.Context().Value("params").(map[string]string); ok {
		return params
	}
	return map[string]string{}
}

// dynamicParams extracts dynamic parameters (e.g., :id, :name, :hawk) from a path based on a given route pattern.
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

// ServeHTTP handles incoming HTTP requests by matching them against registered routes and executing the corresponding handler function.
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

// extractParams extracts dynamic parameters from a matched path based on the route's regex pattern.
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

// LoadStatic registers a route to serve static files from a specified directory.
func (r *Router) LoadStatic(routePath, dir string) {
	r.AddRoute("GET", routePath+"(.*)", static.LoadStatic(routePath, dir))
}

// Start initializes and starts an HTTP server on the specified host and port, using the router instance to handle incoming requests.
func Start(host, port string, r *Router) {
	address := fmt.Sprintf("%s:%s", host, port)

	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	fmt.Printf("Hawkwing server is running on %s\n", address)

	if host == "0.0.0.0" {
		fmt.Println("Server is publicly accessible.")
		// getting the local IP address of the machine😁
		ip, err := getLocalIP()
		if err != nil {
			fmt.Println("Could not determine local IP address.")
		} else {
			fmt.Printf("Access the server at: http://%s:%s\n", ip, port)
		}
	} else {
		fmt.Printf("URL: http://%s:%s\n", host, port)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %s\n", err)
		}
	}()

	<-stop
	fmt.Println("\nShutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %s\n", err)
	}

	fmt.Println("Hawkwing landed gracefully🦅 : Server stopped")
}

// getLocalIP retrieves the first non-loopback IP address of the host machine for displaying in terminal.
func getLocalIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if iface.Flags&(net.FlagUp|net.FlagLoopback) == net.FlagLoopback {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip := v.IP
				if ip.IsLoopback() || ip.To4() == nil {
					continue
				}
				return ip.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no non-loopback address found")
}
