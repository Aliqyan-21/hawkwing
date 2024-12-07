package main

import (
	"github.com/aliqyan-21/hawkwing/internal/render"
	"github.com/aliqyan-21/hawkwing/internal/router"
	"net/http"
)

func main() {
	app := router.Init()

	render.LoadTemplates("./cmd/templates")

	app.LoadStatic("/static/", "./cmd/static")

	app.AddRoute("GET", "/", func(w http.ResponseWriter, req *http.Request) {
		data := map[string]interface{}{
			"Title": "Home Page",
			"Body":  "Welcome to Hawkwing Framework",
		}
		render.RenderHTML(w, "home.html", data)
	})

	// router.Start("localhost", "8080", app)

	// public
	router.Start("0.0.0.0", "8080", app)
}
