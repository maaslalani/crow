package command_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/maaslalani/crow/command"
	"github.com/maaslalani/crow/test"
)

func TestRun(t *testing.T) {
	t.Run("run commands", func(t *testing.T) {
		tf, teardown := test.Mock(t)
		defer teardown()

		command.Run([]string{"echo", "foo"})

		time.Sleep(10 * time.Millisecond)

		fc, err := ioutil.ReadFile(tf.Name())
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(fc), "foo") {
			t.Fatal("command was not run")
		}
	})

	t.Run("kill long running commands", func(t *testing.T) {
		_, teardown := test.Mock(t)
		defer teardown()

		cmd := command.Run([]string{"top"})

		p, err := os.FindProcess(cmd.Process.Pid)
		if err != nil {
			t.Fatal("process was not started")
		}
		if p == nil {
			t.Fatal("process was not started")
		}

		command.Kill(cmd)
		cmd.Process.Wait()
	})
}
