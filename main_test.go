package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/maaslalani/crow/test"
)

func TestCommand(t *testing.T) {
	t.Run("command execution", func(t *testing.T) {
		pwd, _ := os.Getwd()
		app := App(pwd)

		tf, teardown := test.Mock(t)
		defer teardown()

		time.AfterFunc(50*time.Millisecond, func() {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		})

		err := app.RunContext(nil, []string{"crow", "echo", "foo"})
		if err != nil {
			t.Fatal(err)
		}

		fc, err := ioutil.ReadFile(tf.Name())
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(fc), "foo") {
			t.Fatal("command did not run correctly")
		}
	})
}

func TestChanges(t *testing.T) {
	t.Run("file changes", func(t *testing.T) {
		pwd, _ := os.Getwd()
		app := App(pwd)

		tf, teardown := test.Mock(t)
		defer teardown()

		f, err := os.Create("foo.text")
		defer f.Close()
		defer os.Remove(f.Name())

		if err != nil {
			t.Fatal(err)
		}

		time.AfterFunc(50*time.Millisecond, func() {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		})

		err = app.RunContext(nil, []string{"crow", "echo", "bar"})
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(f, "change")

		fc, err := ioutil.ReadFile(tf.Name())
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(fc), "bar") {
			t.Fatal("command did not run correctly")
		}
	})
}

func TestMultipleChanges(t *testing.T) {
	t.Run("file changes", func(t *testing.T) {
		pwd, _ := os.Getwd()
		app := App(pwd)

		tf, teardown := test.Mock(t)
		defer teardown()

		f, err := os.Create("foo.text")
		defer f.Close()
		defer os.Remove(f.Name())

		if err != nil {
			t.Fatal(err)
		}

		time.AfterFunc(50*time.Millisecond, func() {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		})

		err = app.RunContext(nil, []string{"crow", "cat", f.Name()})
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(f, "foo")

		fc, err := ioutil.ReadFile(tf.Name())
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(fc), "foo") {
			t.Log(string(fc))
			t.Fatal("command did not run correctly initially")
		}

		fmt.Println(f, "bar")

		fc, err = ioutil.ReadFile(tf.Name())
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(fc), "foo") {
			t.Log(string(fc))
			t.Fatal("command did not run correctly after two changes")
		}

		if !strings.Contains(string(fc), "bar") {
			t.Log(string(fc))
			t.Fatal("command did not run correctly after two changes")
		}
	})
}
