package output

import (
	"reflect"
	"strings"
	"testing"
)

func TestFormatFileContent(t *testing.T) {
	content := []string{
		"package main",
		"",
		"func main() {}",
	}
	wantNumbered := []string{
		"  1 | package main",
		"  2 | ",
		"  3 | func main() {}",
	}
	wantNoNumber := []string{
		"package main",
		"",
		"func main() {}",
	}

	t.Run("numbered", func(t *testing.T) {
		got := FormatFileContent("main.go", content, true)
		lines := strings.Split(got, "\n")
		if lines[0] != "main.go" {
			t.Errorf("got first line = %q, want %q", lines[0], "main.go")
		}
		if !reflect.DeepEqual(lines[1:], wantNumbered) {
			t.Errorf("got body = %v, want %v", lines[1:], wantNumbered)
		}
	})

	t.Run("no number", func(t *testing.T) {
		got := FormatFileContent("main.go", content, false)
		lines := strings.Split(got, "\n")
		if lines[0] != "main.go" {
			t.Errorf("got first line = %q, want %q", lines[0], "main.go")
		}
		if !reflect.DeepEqual(lines[1:], wantNoNumber) {
			t.Errorf("got body = %v, want %v", lines[1:], wantNoNumber)
		}
	})
}
