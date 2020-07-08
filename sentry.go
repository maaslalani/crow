package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func sentry(c *cli.Context) error {
	fmt.Println("Hello")
	return nil
}
