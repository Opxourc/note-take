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

	"github.com/Opxourc/note-take/config"
	"github.com/Opxourc/note-take/utility"
)

const (
	fileFlag = "f"
	noteFlag = "n"
)

// New reads a flag and its arguments from the supplied scanner.
func RunNew(s *bufio.Scanner) {
	for s != nil && s.Scan() {
		line := s.Text()
		flagScanner := bufio.NewScanner(strings.NewReader(line))
		flag := utility.GetFlag(flagScanner)
		if flag == "" {
			continue
		}

		if flag != fileFlag && flag != noteFlag {
			fmt.Printf("Error: Unknown flag -%s (use -%s or -%s)\n\n", flag, fileFlag, noteFlag)
			return
		}

		flagStart := strings.Index(line, "-"+flag)
		argsScanner := bufio.NewScanner(strings.NewReader(line[flagStart+len(flag)+1:]))
		args := utility.GetArguements(argsScanner)
		if len(args) == 0 {
			fmt.Println("Error: expected a filename")
			return
		}

		filename := args[0]
		if strings.ContainsAny(filename, "\\/") || filename == "." || filename == ".." {
			fmt.Println("Error: filename must be a single name")
			return
		}

		switch flag {
		case fileFlag:
			if len(args) != 1 {
				fmt.Println("Error: -f accepts only a filename")
				return
			}
			newFile(filename)
		case noteFlag:
			remainder := strings.Join(args[1:], " ")
			if len(remainder) < 2 || remainder[0] != '"' || remainder[len(remainder)-1] != '"' {
				fmt.Println("Error: -n requires a note string in double quotes")
				return
			}

			note, err := strconv.Unquote(remainder)
			if err != nil {
				fmt.Printf("Error: invalid quoted note: %v\n", err)
				return
			}
			newNote(filename, note)
		}
		return
	}

	if s != nil && s.Err() != nil {
		fmt.Printf("Error: failed to read input: %v\n", s.Err())
	} else {
		fmt.Println("Error: expected a flag (-f or -n)")
	}
}

// createPath returns the filesystem path for a given filename.
func createPath(filename string) string {
	return filepath.Join(config.NoteFolderPath, filename)
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
	directory := filepath.Clean(config.NoteFolderPath)
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
