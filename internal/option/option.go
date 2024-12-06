package option

import (
	"errors"
	"strings"
)

type Options struct {
	Patterns        []string
	PrintLineNumber bool
	ToStdout        bool
	ToClipboard     bool
	Verbose         bool
	UsageRequested  bool
}

func Parse(args []string) (Options, error) {
	if len(args) == 0 {
		// Instead of returning an error, indicate that usage should be shown
		return Options{UsageRequested: true}, nil
	}

	first := args[0]
	allOptionChars := true
	for _, c := range first {
		if !strings.ContainsRune("nscv", c) {
			allOptionChars = false
			break
		}
	}

	if allOptionChars {
		opts := Options{}
		for _, c := range first {
			switch c {
			case 'n':
				opts.PrintLineNumber = true
			case 's':
				opts.ToStdout = true
			case 'c':
				opts.ToClipboard = true
			case 'v':
				opts.Verbose = true
			}
		}
		if len(args) < 2 {
			return Options{}, errors.New("no patterns provided after options")
		}
		opts.Patterns = args[1:]
		return opts, nil
	} else {
		// If the first argument is one character and not in "nscv", it's invalid
		if len(first) == 1 && !strings.ContainsRune("nscv", rune(first[0])) {
			return Options{}, errors.New("invalid option: " + first)
		}
		// Otherwise, treat them as patterns
		return Options{
			Patterns: args,
		}, nil
	}
}
