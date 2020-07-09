package start

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/maaslalani/crow/command"
	"github.com/maaslalani/crow/config"
	"github.com/maaslalani/crow/watcher"
	"github.com/urfave/cli"
)

// Start begins crow
func Start(cli *cli.Context) error {
	if len(cli.Args()) < 1 {
		log.Fatal("No command provided.")
	}

	dir := cli.String("watch")
	c := command.Run(cli.Args())

	w := watcher.New()
	defer w.Close()

	done := make(chan bool)

	go func() {
		for event := range w.Events {
			if event.Op&fsnotify.Write == fsnotify.Write {
				if config.Restart {
					syscall.Kill(-c.Process.Pid, syscall.SIGKILL)
					c.Process.Wait()
				}
				c = command.Run(cli.Args())
			}
		}
	}()

	go func() {
		for err := range w.Errors {
			log.Println("error:", err)
		}
	}()

	err := filepath.Walk(dir, Traverse(w))
	if err != nil {
		return err
	}

	<-done

	return nil
}

// Traverse returns a WalkFunc which adds paths to watch
func Traverse(w *fsnotify.Watcher) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		for _, i := range config.IgnoredPaths {
			if strings.Contains(path, i) {
				return nil
			}
		}

		return w.Add(path)
	}
}
