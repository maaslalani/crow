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
		ignore      []string
		nestedFiles []string
		extensions  []string
	}{
		{[]string{"foo/"}, []string{"foo.text", "foo.md", "foo.go"}, []string{"text"}},
		{[]string{"bar/"}, []string{"bar.text", "bar.md", "bar.go"}, []string{"md"}},
		{[]string{"baz/"}, []string{"baz.text", "baz.md", "baz.go"}, []string{"go"}},
	}

	for _, tc := range tt {
		for _, i := range tc.ignore {
			err := os.Mkdir(i, os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}

			for _, f := range tc.nestedFiles {
				os.Create(i + f)
			}
		}
	}

	defer func() {
		for _, tc := range tt {
			for _, i := range tc.ignore {
				err := os.RemoveAll(i)
				if err != nil {
					t.Fatal(err)
				}
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

			config.IgnoredPaths = tc.ignore
			err = filepath.Walk(pwd, Traverse(tc.extensions, add))
			if err != nil {
				t.Fatal(err)
			}

			for _, i := range tc.ignore {
				for _, f := range files {
					if strings.Contains(f, i) {
						t.Log(f)
						t.Log(i)
						t.Fatal("incorrect ignore")
					}
				}
			}
		})
	}
}
