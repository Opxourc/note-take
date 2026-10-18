package utility

import (
	"bufio"
	"strings"
)

const (
	flagChar = "-"
)

// GetFlag returns the flag string from the provided scanner.
// If there's no flag, the zero value of the string will be returned instead.
func GetFlag(s *bufio.Scanner) string {
	if s == nil {
		return ""
	}

	// Keep reading through the characters to find the flagChar
	for s.Scan() {
		text := s.Text()

		// Mark the flagStart where the flagChar is
		flagStart := strings.Index(text, flagChar)
		if flagStart < 0 {
			continue
		}

		flag := text[flagStart+1:] // Extract everything after the flagChar

		// Mark the flagEnd where the first space appears in the flag
		// IndexByte is used instead of Index since it's better suited for search for one exactly byte character
		if flagEnd := strings.IndexByte(flag, ' '); flagEnd >= 0 {
			return flag[:flagEnd] // Everything after the flagChar and before the space is now the flag
		}

		return flag
	}

	return ""
}

// GetArguements returns a slice of strings which will be arguements.
// The order in which the strings are placed in the slice is the order of which the arguements came in.
//
// This function should be called after the flag has been processed by the scanner so the flag won't be considered an arguement.
func GetArguements(s *bufio.Scanner) []string {
	var arguments []string
	if s == nil {
		return arguments
	}

	// Keep reading the lines until EOF is reached
	for s.Scan() {
		// Split the arguements into their respective strings
		// There might be multiple arguements so we use a '...' afterwards
		arguments = append(arguments, strings.Fields(s.Text())...)
	}

	return arguments
}
