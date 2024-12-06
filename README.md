# aip - 

A command-line tool to display file contents matched by glob patterns, with optional line numbering, clipboard copying, and verbose mode for matched files.

[日本語はこちら (README-JA.md)](./README-JA.md)

## Features

- Print file contents to stdout or copy to the clipboard.
- Add line numbers with the `n` option.
- Verbose mode `v` shows matched files before contents.
- Pattern matching with glob syntax.

## Install

```
go install github.com/kznrluk/aip@main
```

## Usage

```
aip [options] 
```

**Options**:
- `n` : Number lines
- `s` : Print to stdout (default if no output option is chosen)
- `c` : Copy to clipboard
- `v` : Verbose mode (print matched files before content)

### Examples

- `aip ns *.go`  
  Print all `.go` files with line numbers to stdout.

- `aip nc *.go`  
  Print all `.go` files with line numbers to clipboard.

- `aip vns */**/*.go`  
  Verbose, numbered, stdout output for all `.go` files in subdirectories.

## License

This software is released under the MIT License.
