package main

import (
	"github.com/aliqyan-21/hawkwing/internal/render"
	"github.com/aliqyan-21/hawkwing/internal/router"
	"net/http"
)

func main() {
	app := router.New()

	// this is to specify in which directory the html files are
	render.LoadTemplates("./cmd/templates")

	app.AddRoute("GET", "/", func(w http.ResponseWriter, req *http.Request) {
		data := map[string]interface{}{
			"Title": "Home Page",
			"Body":  "Welcome to Hawkwing Framework",
		}
		render.RenderHTML(w, "home.html", data)
	})

	app.AddRoute("GET", "/about", func(w http.ResponseWriter, req *http.Request) {
		data := map[string]interface{}{
			"Title": "About Us",
			"Body":  "A lightweight user friendly web framework",
		}
		render.RenderHTML(w, "about.html", data)
	})

	router.Start(":8080", app)
}
