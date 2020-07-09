package start

import (
	"log"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/maaslalani/crow/command"
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
				syscall.Kill(-c.Process.Pid, syscall.SIGKILL)
				c.Process.Wait()
				c = command.Run(cli.Args())
			}
		}
	}()

	go func() {
		for err := range w.Errors {
			log.Println("error:", err)
		}
	}()

	err := w.Add(dir)
	if err != nil {
		return err
	}

	<-done

	return nil
}
