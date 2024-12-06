package hawkwing

import (
	"github.com/aliqyan-21/hawkwing/internal/middleware"
	"github.com/aliqyan-21/hawkwing/internal/render"
	"github.com/aliqyan-21/hawkwing/internal/router"
	"net/http"
)

func New() *router.Router {
	return router.New()
}

func AddRoute(r *router.Router, method, path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	r.AddRoute(method, path, handler, middlewares...)
}

func Start(port string, r *router.Router) {
	Start(port, r)
}

func LoadStatic(r *router.Router, routePath, dir string) {
	r.LoadStatic(routePath, dir)
}

func GetRouteParams(r *http.Request) map[string]string {
	return GetRouteParams(r)
}

func LoadTemplates(templateDir string) {
	render.LoadTemplates(templateDir)
}

func RenderHTML(w http.ResponseWriter, tmpl string, data interface{}) {
	render.RenderHTML(w, tmpl, data)
}

var LoggerMiddleware = middleware.Logger
var AuthMiddleware = middleware.Auth
var ContentTypeMiddleware = middleware.ContentType
