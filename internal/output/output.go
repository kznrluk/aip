package output

import (
	"fmt"
	"strings"
)

func FormatFileContent(filename string, lines []string, numbered bool) string {
	var sb strings.Builder
	sb.WriteString(filename + "\n")

	if !numbered {
		for _, l := range lines {
			sb.WriteString(l + "\n")
		}
		return strings.TrimRight(sb.String(), "\n")
	}

	lineCount := len(lines)
	maxDigits := len(fmt.Sprintf("%d", lineCount))
	if maxDigits < 3 {
		maxDigits = 3
	}

	for i, l := range lines {
		sb.WriteString(fmt.Sprintf("%*d | %s\n", maxDigits, i+1, l))
	}

	return strings.TrimRight(sb.String(), "\n")
}
