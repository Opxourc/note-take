package utility

import (
	"strings"
)

const (
	flagChar = "-"
)

// ParseFlagArgs separates the first flag from the remaining command arguments.
func ParseFlagArgs(input []string) (string, []string) {
	if len(input) == 0 || !strings.HasPrefix(input[0], flagChar) {
		return "", nil
	}

	return strings.TrimPrefix(input[0], flagChar), input[1:]
}
