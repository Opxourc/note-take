package utility_test

import (
	"reflect"
	"testing"

	"github.com/Opxourc/note-take/utility"
)

func TestParseFlagArgs(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		wantFlag string
		wantArgs []string
	}{
		{
			name:     "flag and arguments",
			input:    []string{"-n", "sample", "hello", "world"},
			wantFlag: "n",
			wantArgs: []string{"sample", "hello", "world"},
		},
		{
			name:     "flag only",
			input:    []string{"-f"},
			wantFlag: "f",
			wantArgs: []string{},
		},
		{
			name:     "no flag",
			input:    []string{"sample", "hello"},
			wantFlag: "",
			wantArgs: nil,
		},
		{
			name:     "empty input",
			input:    []string{},
			wantFlag: "",
			wantArgs: nil,
		},
		{
			name:     "empty flag",
			input:    []string{"-", "sample"},
			wantFlag: "",
			wantArgs: []string{"sample"},
		},
		{
			name:     "only first argument is treated as flag",
			input:    []string{"-n", "-f", "sample"},
			wantFlag: "n",
			wantArgs: []string{"-f", "sample"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotFlag, gotArgs := utility.ParseFlagArgs(test.input)

			if gotFlag != test.wantFlag {
				t.Errorf("ParseFlagArgs() flag = %q, want %q", gotFlag, test.wantFlag)
			}

			if !reflect.DeepEqual(gotArgs, test.wantArgs) {
				t.Errorf("ParseFlagArgs() args = %#v, want %#v", gotArgs, test.wantArgs)
			}
		})
	}
}
