package start

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/maaslalani/crow/config"
)

func TestTraverse(t *testing.T) {
	tt := []struct {
		extensions []string
		files      []string
	}{
		{[]string{"go"}, []string{"start.go", "start_test.go"}},
		{[]string{"md"}, []string{""}},
		{[]string{"text"}, []string{"foo.text"}},
		{[]string{"foo.text"}, []string{"foo.text"}},
		{[]string{"crow/start/foo.text"}, []string{"foo.text"}},
		{[]string{"start.go"}, []string{"start.go"}},
		{[]string{"start_test.go"}, []string{"start_test.go"}},
		{[]string{"start_test.go", "start.go"}, []string{"start.go", "start_test.go"}},
		{[]string{"foo.text", "start.go", "foo.text"}, []string{"foo.text", "start.go", "start_test.go"}},
	}

	_, err := os.Create("foo.text")
	defer os.Remove("foo.text")
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range tt {
		t.Run("add correct extensions", func(t *testing.T) {
			var files []string
			add := func(file string) error {
				files = append(files, file)
				return nil
			}

			pwd, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			err = filepath.Walk(pwd, Traverse(tc.extensions, add))
			if err != nil {
				t.Fatal(err)
			}

			for i, file := range files {
				if !strings.HasSuffix(file, tc.files[i]) {
					t.Log(file)
					t.Log(tc.files[i])
					t.Fatal("incorrect file watched")
				}
			}
		})
	}
}

func TestIgnore(t *testing.T) {
	tt := []struct {
		ignore      string
		nestedFiles []string
		extensions  []string
	}{
		{"foo", []string{"foo/foo.text", "foo/foo.md", "foo/foo.go"}, []string{"text"}},
		{"bar", []string{"bar/bar.text", "bar/bar.md", "bar/bar.go"}, []string{"md"}},
		{"baz", []string{"baz/baz.text", "baz/baz.md", "baz/baz.go"}, []string{"go"}},
	}

	for _, tc := range tt {
		err := os.Mkdir(tc.ignore, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		for _, f := range tc.nestedFiles {
			os.Create(f)
		}
	}

	defer func() {
		for _, tc := range tt {
			err := os.RemoveAll(tc.ignore)
			if err != nil {
				t.Fatal(err)
			}
		}
	}()

	for _, tc := range tt {
		t.Run("ignore intended directories", func(t *testing.T) {
			var files []string
			add := func(file string) error {
				files = append(files, file)
				return nil
			}

			pwd, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			ip := config.IgnoredPaths
			defer func() {
				config.IgnoredPaths = ip
			}()

			config.IgnoredPaths = []string{tc.ignore}
			err = filepath.Walk(pwd, Traverse(tc.extensions, add))
			if err != nil {
				t.Fatal(err)
			}

			for _, f := range files {
				if strings.Contains(f, tc.ignore) {
					t.Log(f)
					t.Log(tc.ignore)
					t.Fatal("incorrect ignore")
				}
			}
		})
	}
}
