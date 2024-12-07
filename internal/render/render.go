// Package render provides functions for loading and rendering HTML templates.
package render

import (
	"github.com/fsnotify/fsnotify"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates *template.Template

// LoadTemplates parses all HTML files within the specified directory and stores them in a global template cache.
// It also starts a goroutine to watch for changes in the template files and automatically reload them.
func LoadTemplates(templateDir string) {
	templates = template.Must(template.ParseGlob(filepath.Join(templateDir, "*.html")))
	go watcher(templateDir)
}

// RenderHTML renders the specified template with the provided data and writes the output to the HTTP response writer.
// It first checks if the templates are loaded, and returns an error if not.
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

// watcher is a helper function that watches the template directory for changes.
// Upon detecting a modification (write event), it reloads the templates using LoadTemplates.
func watcher(templateDir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Error creating file watcher:", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add(templateDir)
	if err != nil {
		log.Println("Error adding directory to watcher:", err)
		return
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("Detected change in template:", event.Name)
				LoadTemplates(templateDir)
			}
		case err := <-watcher.Errors:
			log.Println("Error watching files:", err)
		}
	}
}
