package commands

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	fileFlag    = "f"
	noteFlag    = "n"
	notesFolder = "notes"
)

// New finds a flag in the supplied input or keeps reading lines until it finds one.
// The rest of that line contains either a filename or a filename followed by a quoted note.
func New(input string, scanner *bufio.Scanner) {
	for {
		flag, args, found := flagAndArgs(input)
		if found {
			if flag != fileFlag && flag != noteFlag {
				fmt.Printf("Error: Unknown flag -%s (use -%s or -%s)\n\n", flag, fileFlag, noteFlag)
				return
			}

			filename, note, err := parseNewArgs(args, flag == noteFlag)
			if err != nil {
				fmt.Printf("Error: %v\n\n", err)
				return
			}

			// Flag determines whether to create a file or a note
			switch flag {
			case fileFlag:
				newFile(filename)
			case noteFlag:
				newNote(filename, note)
			}

			return
		}

		if scanner == nil || !scanner.Scan() {
			if scanner != nil && scanner.Err() != nil {
				fmt.Printf("Error: failed to read input: %v\n", scanner.Err())
			} else {
				fmt.Println("Error: expected a flag (-f or -n)")
			}
			return
		}
		input = scanner.Text()
	}
}

// flagAndArgs returns the first dash-prefixed token and everything after it.
func flagAndArgs(input string) (flag, args string, found bool) {
	fields := strings.Fields(input)
	position := 0

	for _, field := range fields {
		offset := strings.Index(input[position:], field)
		if offset < 0 {
			return "", "", false
		}

		start := position + offset
		position = start + len(field)
		if len(field) > 1 && field[0] == '-' {
			return field[1:], strings.TrimSpace(input[position:]), true
		}
	}

	return "", "", false
}

// parseNewArgs returns the filename and note from the given args.
func parseNewArgs(args string, wantsNote bool) (filename, note string, err error) {
	args = strings.TrimSpace(args)
	if args == "" {
		return "", "", fmt.Errorf("expected a filename")
	}

	fields := strings.Fields(args)
	if len(fields) == 0 {
		return "", "", fmt.Errorf("expected a filename")
	}

	filename = fields[0]
	remainder := strings.TrimSpace(args[len(filename):])
	if strings.ContainsAny(filename, "\\/") || filename == "." || filename == ".." {
		return "", "", fmt.Errorf("filename must be a single name")
	}

	if !wantsNote {
		if strings.TrimSpace(remainder) != "" {
			return "", "", fmt.Errorf("-f accepts only a filename")
		}
		return filename, "", nil
	}

	remainder = strings.TrimSpace(remainder)
	if len(remainder) < 2 || remainder[0] != '"' || remainder[len(remainder)-1] != '"' {
		return "", "", fmt.Errorf("-n requires a note string in double quotes")
	}

	note, err = strconv.Unquote(remainder)
	if err != nil {
		return "", "", fmt.Errorf("invalid quoted note: %w", err)
	}

	return filename, note, nil
}

// createPath returns the filesystem path for a given filename.
func createPath(filename string) string {
	return filepath.Join(notesFolder, filename) // Join the notes directory and filename into a filesystem path
}

// newNote creates a new note file with the given filename and note.
func newNote(filename, note string) {
	path := createPath(filename)
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error: file %q does not exist\n", filename)
		} else {
			fmt.Printf("Error checking file %q: %v\n", filename, err)
		}
		return
	}
	if !info.Mode().IsRegular() {
		fmt.Printf("Error: %q is not a regular file\n", filename)
		return
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		fmt.Printf("Error opening file %q: %v\n", filename, err)
		return
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, note); err != nil {
		fmt.Printf("Error writing note to %q: %v\n", filename, err)
		return
	}
	fmt.Printf("Added note to %q\n", filename)
}

// newFile creates a new file with the given filename.
func newFile(filename string) {
	directory := filepath.Clean(notesFolder)
	if err := os.MkdirAll(directory, 0o755); err != nil {
		fmt.Printf("Error creating notes directory: %v\n", err)
		return
	}

	path := createPath(filename)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		if errors.Is(err, fs.ErrExist) {
			fmt.Printf("Error: file %q already exists\n", filename)
		} else {
			fmt.Printf("Error creating file %q: %v\n", filename, err)
		}
		return
	}

	if err := file.Close(); err != nil {
		fmt.Printf("Error closing file %q: %v\n\n", filename, err)
		return
	}

	fmt.Printf("Created file %q\n\n", filename)
}
