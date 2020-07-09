package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/maaslalani/crow/cmd"
	"github.com/maaslalani/crow/watcher"
	"github.com/urfave/cli"
)

var pid int

func crow(c *cli.Context) error {
	dir := c.String("watch")
	p := cmd.Run(c.Args()).Process

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
					p.Kill()
					p.Wait()
					p = cmd.Run(c.Args()).Process
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
