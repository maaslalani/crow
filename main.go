package main

import (
	"log"
	"os"

	"github.com/maaslalani/crow/start"
	"github.com/urfave/cli/v2"
)

func main() {
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	app := App(pwd)
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// App returns a *cli.App for crow
func App(pwd string) *cli.App {
	return &cli.App{
		Name:    "Crow",
		Usage:   "Run arbitrary commands on file changes",
		Version: "0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "watch",
				Aliases: []string{"w"},
				Value:   pwd,
				Usage:   "Directory to watch",
			},
			&cli.StringSliceFlag{
				Name:    "ext",
				Aliases: []string{"e"},
				Value:   cli.NewStringSlice(""),
				Usage:   "File extensions to watch",
			},
		},
		Action: start.Start,
	}
}
