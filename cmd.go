package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/maaslalani/crow/watcher"
	"github.com/urfave/cli"
)

func cmd(c *cli.Context) error {
	dir := c.String("watch")

	w := watcher.New()
	defer w.Close()

	done := make(chan bool)
	go watcher.Watch(w, func() {
		Run(c.Args())
	})

	err := w.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	<-done

	return nil
}

func Run(cmd []string) {
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Start()
	if err != nil {
		log.Fatal(err)
	}
	c.Wait()
}
