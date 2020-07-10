package main

import (
	"io/ioutil"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	pwd, _ := os.Getwd()
	app := App(pwd)

	tf, teardown := MockOs(t)
	defer teardown()

	time.AfterFunc(50*time.Millisecond, func() {
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	})

	err := app.RunContext(nil, []string{"crow", "echo", "hello"})
	if err != nil {
		t.Fatal(err)
	}

	fc, err := ioutil.ReadFile(tf.Name())

	if !strings.Contains(string(fc), "hello") {
		t.Fatal("command not run correctly")
	}
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
