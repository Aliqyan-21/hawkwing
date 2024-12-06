package render

import (
	"github.com/fsnotify/fsnotify"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates *template.Template

func LoadTemplates(templateDir string) {
	templates = template.Must(template.ParseGlob(filepath.Join(templateDir, "*.html")))
	go watcher(templateDir)
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
