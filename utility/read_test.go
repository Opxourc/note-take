package utility

import (
	"bufio"
	"strings"
	"testing"
)

type test[T any] struct {
	name  string
	input string
	want  T
}

func TestGetFlag(t *testing.T) {
	tests := []test[string]{
		{
			name:  "flag after text",
			input: "text before -name rest",
			want:  "name",
		},
		{
			name:  "flag on following line",
			input: "no flag here\n-x value",
			want:  "x",
		},
		{
			name:  "no flag",
			input: "no flag here",
			want:  "",
		},
		{
			name:  "hyphen without characters",
			input: "-",
			want:  "",
		},
		{
			name:  "hyphen followed by space",
			input: "- next",
			want:  "",
		},
		{
			name:  "flag at end of input",
			input: "-name",
			want:  "name",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(test.input))

			// The flag returned should be the same as the wanted flag
			if got := GetFlag(scanner); got != test.want {
				t.Errorf("GetFlag() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestGetArguements(t *testing.T) {
	tests := []test[[]string]{
		{
			name:  "space separated arguments",
			input: "first second third",
			want:  []string{"first", "second", "third"},
		},
		{
			name:  "repeated spaces",
			input: " first  second ",
			want:  []string{"first", "second"},
		},
		{
			name:  "multiple lines",
			input: "first second\nthird",
			want:  []string{"first", "second", "third"},
		},
		{
			name:  "empty input",
			input: "",
			want:  []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(test.input))
			got := GetArguements(scanner)

			// The length returned should be the same as the wanted length
			if len(got) != len(test.want) {
				t.Fatalf("GetArguements() = %q, want %q", got, test.want)
			}

			// The values of every index of the returned should be the same as the wanted values
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("GetArguements()[%d] = %q, want %q", i, got[i], test.want[i])
				}
			}
		})
	}
}
