package cmd

import (
	"log"
	"os"
	"os/exec"
)

// KillFunc kills the child process when call
type KillFunc func() error

// Run executes commands
func Run(cmd []string) KillFunc {
	Clear()

	c := exec.Command(cmd[0], cmd[1:]...)

	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout

	err := c.Start()
	if err != nil {
		log.Fatal(err)
	}

	return c.Process.Kill
}

// Clear clears the screen
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
