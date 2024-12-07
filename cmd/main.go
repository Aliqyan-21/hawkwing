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

	router.Start(":8080", app)
}
