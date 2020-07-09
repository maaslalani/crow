package command

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
	Sync(c)
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
	Sync(c)
	c.Run()
}

// Sync synchronizes the commands stdin, stdout, and stderr
func Sync(c *exec.Cmd) {
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
}
