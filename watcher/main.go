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
