package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kznrluk/aip/internal/clipboard"
	"github.com/kznrluk/aip/internal/file"
	"github.com/kznrluk/aip/internal/option"
	"github.com/kznrluk/aip/internal/output"
)

func main() {
	opts, err := option.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "github.com/kznrluk/aip: %v\n", err)
		os.Exit(1)
	}

	if opts.UsageRequested {
		printUsage()
		os.Exit(0)
	}

	files, err := file.CollectFiles(opts.Patterns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "github.com/kznrluk/aip: %v\n", err)
		os.Exit(1)
	}

	var results []string

	if opts.Verbose {
		for _, f := range files {
			fmt.Println("Appending file:", f)
		}
	}

	for _, f := range files {
		content, err := file.ReadFileLines(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "github.com/kznrluk/aip: %v\n", err)
			os.Exit(1)
		}
		formatted := output.FormatFileContent(f, content, opts.PrintLineNumber)
		results = append(results, formatted)
	}

	all := strings.Join(results, "\n\n")

	// If no output option is chosen, default is stdout.
	if opts.ToStdout || (!opts.ToStdout && !opts.ToClipboard) {
		fmt.Println(all)
	}

	if opts.ToClipboard {
		err := clipboard.WriteClipboard(all)
		if err != nil {
			fmt.Fprintf(os.Stderr, "github.com/kznrluk/aip: %v\n", err)
			os.Exit(1)
		}
	}
}

func printUsage() {
	fmt.Println("Usage: aip [options] <patterns>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  n  Number lines")
	fmt.Println("  s  Print to stdout (default if no output option chosen)")
	fmt.Println("  c  Copy to clipboard")
	fmt.Println("  v  Verbose mode (print matched files before content)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  aip ns *.go                Print all go files with line numbers to stdout")
	fmt.Println("  aip nc *.go                Print all go files with line numbers to clipboard")
	fmt.Println("  aip vns */**/*.go          Verbose, numbered, stdout for all go files in subdirs")
	fmt.Println()
}
