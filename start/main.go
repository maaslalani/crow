package start

import (
	"log"
	"os"
	"os/signal"
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
	cmd := command.Run(cli.Args())

	w := watcher.New()
	defer w.Close()

	done := make(chan bool)
	sigs := make(chan os.Signal, 1)

	go func() {
		for {
			select {
			case event, ok := <-w.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					command.Kill(cmd)
					cmd.Process.Wait()
					cmd = command.Run(cli.Args())
				}
			case err, ok := <-w.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		command.Kill(cmd)
		cmd.Process.Wait()
		done <- true
	}()

	err := filepath.Walk(dir, Traverse(cli, w))
	if err != nil {
		return err
	}

	<-done
	return nil
}

// Traverse returns a WalkFunc which adds paths to watch
func Traverse(cli *cli.Context, w *fsnotify.Watcher) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		for _, i := range config.IgnoredPaths {
			if strings.Contains(path, i) {
				return nil
			}
		}

		for _, ext := range cli.StringSlice("ext") {
			if strings.HasSuffix(path, ext) {
				return w.Add(path)
			}
		}

		return nil
	}
}
