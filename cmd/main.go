package cmd

import (
	"log"
	"os"
	"os/exec"
)

// Run executes commands
func Run(cmd []string) *exec.Cmd {
	Clear()

	c := exec.Command(cmd[0], cmd[1:]...)

	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout

	err := c.Start()
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// Clear clears the screen
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
