package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates *template.Template

func LoadTemplates(templateDir string) {
	templates = template.Must(template.ParseGlob(filepath.Join(templateDir, "*.html")))
}

func RenderHTML(w http.ResponseWriter, tmpl string, data interface{}) {
	if templates == nil {
		log.Println("Error: Templates not loaded. Please call LoadTemplates first.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		log.Printf("Error rendering template: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
