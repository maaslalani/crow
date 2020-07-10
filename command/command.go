package command

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/maaslalani/crow/config"
)

// Run executes commands
func Run(cmd []string) *exec.Cmd {
	if config.Clear {
		Clear()
	}

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

// Sync synchronizes the commands stdin, stdout, and stderr with os
func Sync(c *exec.Cmd) {
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
}

// Kill sends a SIGKILL to the process of the command
func Kill(c *exec.Cmd) {
	syscall.Kill(-c.Process.Pid, syscall.SIGKILL)
}
