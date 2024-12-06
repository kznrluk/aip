package option

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantOpt Options
		wantErr bool
	}{
		{
			name:    "no args",
			args:    []string{},
			wantOpt: Options{UsageRequested: true},
			wantErr: false,
		},
		{
			name:    "only pattern",
			args:    []string{"*.go"},
			wantOpt: Options{Patterns: []string{"*.go"}, PrintLineNumber: false, ToStdout: false, ToClipboard: false, Verbose: false},
			wantErr: false,
		},
		{
			name:    "n option with pattern",
			args:    []string{"n", "*.go"},
			wantOpt: Options{Patterns: []string{"*.go"}, PrintLineNumber: true, ToStdout: false, ToClipboard: false, Verbose: false},
			wantErr: false,
		},
		{
			name:    "ns option with pattern",
			args:    []string{"ns", "*.go"},
			wantOpt: Options{Patterns: []string{"*.go"}, PrintLineNumber: true, ToStdout: true, ToClipboard: false, Verbose: false},
			wantErr: false,
		},
		{
			name:    "nc option with pattern",
			args:    []string{"nc", "*.go"},
			wantOpt: Options{Patterns: []string{"*.go"}, PrintLineNumber: true, ToStdout: false, ToClipboard: true, Verbose: false},
			wantErr: false,
		},
		{
			name:    "c option with multiple patterns",
			args:    []string{"c", "*.go", "cmd/*.go"},
			wantOpt: Options{Patterns: []string{"*.go", "cmd/*.go"}, PrintLineNumber: false, ToStdout: false, ToClipboard: true, Verbose: false},
			wantErr: false,
		},
		{
			name:    "invalid option",
			args:    []string{"x", "*.go"},
			wantErr: true,
		},
		{
			name:    "no patterns after option",
			args:    []string{"n"},
			wantErr: true,
		},
		{
			name:    "v option with pattern",
			args:    []string{"v", "*.go"},
			wantOpt: Options{Patterns: []string{"*.go"}, PrintLineNumber: false, ToStdout: false, ToClipboard: false, Verbose: true},
			wantErr: false,
		},
		{
			name:    "vn option with pattern",
			args:    []string{"vn", "*.go"},
			wantOpt: Options{Patterns: []string{"*.go"}, PrintLineNumber: true, ToStdout: false, ToClipboard: false, Verbose: true},
			wantErr: false,
		},
		{
			name:    "vns option with pattern",
			args:    []string{"vns", "*.go"},
			wantOpt: Options{Patterns: []string{"*.go"}, PrintLineNumber: true, ToStdout: true, ToClipboard: false, Verbose: true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpt, err := Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(gotOpt, tt.wantOpt) {
				t.Errorf("Parse() got = %v, want %v", gotOpt, tt.wantOpt)
			}
		})
	}
}
