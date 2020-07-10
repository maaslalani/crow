package test

import (
	"io/ioutil"
	"os"
	"testing"
)

// Mock creates a temp file and mocks stdout
func Mock(t *testing.T) (*os.File, func()) {
	stdout := os.Stdout

	tf, err := ioutil.TempFile("", "crow")
	if err != nil {
		t.Fatal()
	}

	os.Stdout = tf

	return tf, func() {
		tf.Close()
		os.Remove(tf.Name())
		os.Stdout = stdout
	}
}
