package watcher

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// New returns a fsnotify file watcher
func New() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	return watcher
}

// Watch watches for file changes and performs an action
func Watch(w *fsnotify.Watcher, f func()) {
	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				f()
			}
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
