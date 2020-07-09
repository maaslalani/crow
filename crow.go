package main

import (
	"log"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/maaslalani/crow/cmd"
	"github.com/maaslalani/crow/watcher"
	"github.com/urfave/cli"
)

var pid int

func crow(cli *cli.Context) error {
	if len(cli.Args()) < 1 {
		log.Fatal("No command provided.")
	}

	dir := cli.String("watch")
	c := cmd.Run(cli.Args())

	w := watcher.New()
	defer w.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-w.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					syscall.Kill(-c.Process.Pid, syscall.SIGKILL)
					c.Process.Wait()
					c = cmd.Run(cli.Args())
				}
			case err, ok := <-w.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err := w.Add(dir)
	if err != nil {
		return err
	}

	<-done

	return nil
}
