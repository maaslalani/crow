package main

import (
	"io/ioutil"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	t.Run("command execution", func(t *testing.T) {
		pwd, _ := os.Getwd()
		app := App(pwd)

		tf, teardown := MockOs(t)
		defer teardown()

		time.AfterFunc(100*time.Millisecond, func() {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		})

		err := app.RunContext(nil, []string{"crow", "echo", "foo"})
		if err != nil {
			t.Fatal(err)
		}

		fc, err := ioutil.ReadFile(tf.Name())

		if !strings.Contains(string(fc), "foo") {
			t.Fatal("command did not run correctly")
		}
	})
}

func TestChanges(t *testing.T) {
	t.Run("file changes", func(t *testing.T) {
		pwd, _ := os.Getwd()
		app := App(pwd)

		tf, teardown := MockOs(t)
		defer teardown()

		f, err := os.Create("foo.text")
		defer f.Close()
		defer os.Remove(f.Name())

		if err != nil {
			t.Fatal(err)
		}

		time.AfterFunc(100*time.Millisecond, func() {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		})

		err = app.RunContext(nil, []string{"crow", "echo", "bar"})
		if err != nil {
			t.Fatal(err)
		}

		_, err = f.WriteString("change")
		if err != nil {
			t.Fatal(err)
		}

		fc, err := ioutil.ReadFile(tf.Name())
		if !strings.Contains(string(fc), "bar") {
			t.Fatal("command did not run multiple times")
		}
	})
}

func MockOs(t *testing.T) (*os.File, func()) {
	stdin := os.Stdin
	stdout := os.Stdout

	tf, err := ioutil.TempFile("", "crow")
	if err != nil {
		t.Fatal()
	}

	os.Stdin = tf
	os.Stdout = tf

	return tf, func() {
		tf.Close()
		os.Remove(tf.Name())
		os.Stdin = stdin
		os.Stdout = stdout
	}
}
