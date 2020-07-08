package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func sentry(c *cli.Context) error {
	dir := c.String("watch")
	fmt.Println(dir)

	args := c.Args()
	cmd := strings.Join(args, " ")

	fmt.Println(cmd)
	return nil
}
