package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCollectFiles(t *testing.T) {
	// Setup temp dirs and files
	tmpDir, err := ioutil.TempDir("", "aip_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	subDir := filepath.Join(tmpDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatal(err)
	}

	files := []string{
		filepath.Join(tmpDir, "main.go"),
		filepath.Join(subDir, "test.go"),
		filepath.Join(subDir, "test.txt"),
	}

	for _, f := range files {
		if err := ioutil.WriteFile(f, []byte("content"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Test
	got, err := CollectFiles([]string{filepath.Join(tmpDir, "*.go"), filepath.Join(tmpDir, "sub", "*.go")})
	if err != nil {
		t.Fatalf("CollectFiles() error: %v", err)
	}

	want := []string{
		filepath.Join(tmpDir, "main.go"),
		filepath.Join(subDir, "test.go"),
	}

	// Sort is unnecessary here if we don't rely on order, but let's just check by set
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CollectFiles() = %v, want %v", got, want)
	}
}
