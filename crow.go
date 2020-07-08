package main

import (
	"github.com/maaslalani/crow/cmd"
	"github.com/maaslalani/crow/watcher"
	"github.com/urfave/cli"
)

func crow(c *cli.Context) error {
	dir := c.String("watch")
	kill := cmd.Run(c.Args())

	w := watcher.New()
	defer w.Close()

	done := make(chan bool)

	go watcher.Watch(w, func() {
		kill()
		kill = cmd.Run(c.Args())
	})

	err := w.Add(dir)
	if err != nil {
		return err
	}

	<-done

	return nil
}
