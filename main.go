package main

import (
	"log"
	"os"

	"github.com/maaslalani/crow/start"
	"github.com/urfave/cli"
)

func main() {
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:    "Crow",
		Usage:   "Run arbitrary commands on file changes",
		Version: "0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "watch, w",
				Value: pwd,
				Usage: "Directory to watch",
			},
			&cli.StringSliceFlag{
				Name:  "ext, e",
				Value: &cli.StringSlice{""},
				Usage: "File extensions to watch",
			},
		},
		Action: start.Start,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
