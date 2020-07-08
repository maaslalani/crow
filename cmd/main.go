package cmd

import (
	"log"
	"os"
	"os/exec"
)

// Run executes commands
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
