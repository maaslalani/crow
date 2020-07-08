package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:    "Sentry",
		Usage:   "Run arbitrary commands on file changes",
		Version: "0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "watch, w",
				Value: ".",
				Usage: "Directory to watch",
			},
		},
		Action: cmd,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
