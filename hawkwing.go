// Package hawkwing provides a simple and efficient HTTP framework for building web applications.
// It offers functionalities like routing, middleware support, templating, and static file serving
package hawkwing

import (
	"github.com/aliqyan-21/hawkwing/internal/middleware"
	"github.com/aliqyan-21/hawkwing/internal/render"
	"github.com/aliqyan-21/hawkwing/internal/router"
	"github.com/aliqyan-21/hawkwing/internal/static"
	"net/http"
)

// Init initializes a new router instance.
func Init() *router.Router {
	return router.Init()
}

// AddRoute registers a new route with the specified method, path, handler function, and optional middlewares.
func AddRoute(r *router.Router, method, path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	r.AddRoute(method, path, handler, middlewares...)
}

// Start initializes and starts an HTTP server on the specified port, using the provided router instance.
func Start(host, port string, r *router.Router) {
	router.Start(host, port, r)
}

// LoadStatic registers a handler for serving static files from a specified directory.
func LoadStatic(routePath, dir string) {
	static.LoadStatic(routePath, dir)
}

// GetRouteParams retrieves dynamic path parameters from the request context (delegates to internal function).
func GetRouteParams(r *http.Request) map[string]string {
	return GetRouteParams(r)
}

// LoadTemplates parses and prepares HTML templates for rendering.
func LoadTemplates(templateDir string) {
	render.LoadTemplates(templateDir)
}

// RenderHTML renders the specified template with the provided data and writes the output to the HTTP response writer.
func RenderHTML(w http.ResponseWriter, tmpl string, data interface{}) {
	render.RenderHTML(w, tmpl, data)
}

// LoggerMiddleware, AuthMiddleware, ContentTypeMiddleware are pre-defined middleware functions provided by the middleware package.
var LoggerMiddleware = middleware.Logger
var AuthMiddleware = middleware.Auth
var ContentTypeMiddleware = middleware.ContentType
var ErrorHandlerMiddleware = middleware.ErrorHandler
