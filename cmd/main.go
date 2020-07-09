package cmd

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// Run executes commands
func Run(cmd []string) *exec.Cmd {
	Clear()

	c := exec.Command(cmd[0], cmd[1:]...)

	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout

	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	err := c.Start()
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// Clear clears the screen
func Clear() {
	c := exec.Command("clear")

	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout

	c.Run()
}
