package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/maaslalani/crow/watcher"
	"github.com/urfave/cli"
)

func cmd(c *cli.Context) error {
	dir := c.String("watch")
	args := c.Args()
	cmd := strings.Join(args, " ")
	fmt.Println(cmd)

	w := watcher.New()
	defer w.Close()

	done := make(chan bool)
	go watcher.Watch(w, func() {
		fmt.Println(cmd)
	})

	err := w.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	<-done

	return nil
}
