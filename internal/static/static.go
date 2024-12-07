// Package static provides utilities for serving static files efficiently.
package static

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"net/http"
	"path/filepath"
)

// LoadStatic registers a handler that serves static files from a specified directory.
// It accepts a route path for prefixing the URLs and the actual directory containing the static files.
// Additionally, it starts a goroutine to watch for changes in the directory and logs them.
func LoadStatic(routePath, dir string) http.HandlerFunc {
	go watcher(dir)
	fs := http.FileServer(http.Dir(filepath.Clean(dir)))
	return func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(routePath, fs).ServeHTTP(w, r)
	}
}

// watcher is a helper function that watches the static directory for changes.
// It monitors for various file operations (write, create, remove) and logs them.
func watcher(staticDir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Error creating file watcher:", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add(staticDir)
	if err != nil {
		log.Println("Error adding static directory to watcher:", err)
		return
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) != 0 {
				log.Printf("Static file change detected: %s", event.Name)
			}
		case err := <-watcher.Errors:
			log.Println("Error watching static files:", err)
		}
	}
}
