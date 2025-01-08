package file

import (
	"github.com/bmatcuk/doublestar/v4"
	"io/ioutil"
)

func CollectFiles(patterns []string) ([]string, error) {
	var results []string
	for _, p := range patterns {
		matches, err := doublestar.FilepathGlob(p)
		if err != nil {
			return nil, err
		}
		results = append(results, matches...)
	}
	return results, nil
}

func ReadFileLines(path string) ([]string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := string(b)
	return splitLines(content), nil
}

// splitLines handles \n
func splitLines(text string) []string {
	if text == "" {
		return []string{""}
	}
	lines := []string{}
	start := 0
	for i, ch := range text {
		if ch == '\n' {
			lines = append(lines, text[start:i])
			start = i + 1
		}
	}
	if start <= len(text) {
		lines = append(lines, text[start:])
	}
	return lines
}
