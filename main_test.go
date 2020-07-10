package main

import (
	"log"
	"os"
	"testing"
)

func TestCrow(t *testing.T) {
	pwd, _ := os.Getwd()
	app := App(pwd)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
