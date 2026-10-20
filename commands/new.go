package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Opxourc/note-take/utility"
)

const (
	newFileFlag = "f"
	newNoteFlag = "n"
)

// RunNew creates a file or adds a note, depending on the flag.
func RunNew(input []string) {
	flag, args := utility.ParseFlagArgs(input)
	dir, err := utility.CreateNotesDir()
	if err != nil {
		fmt.Printf("Unable to create notes directory: %v\n", err)
		return
	}

	switch flag {
	case newFileFlag:
		newFile(args, dir)
	case newNoteFlag:
		newNote(args, dir)
	default:
		fmt.Println("Invalid flag provided.")
	}
}

// newFile creates a new notes file.
func newFile(args []string, dir string) {
	// There should be 1 arguement
	if len(args) != 1 {
		fmt.Println("Expected a file name.")
		return
	}

	fileName := args[0]

	// Check if the file already exists
	// Don't truncate an existing file which would delete it's contents
	path := filepath.Join(dir, fileName)
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("'%s' already exists.\n", fileName)
		return
	} else if !errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Unable to check '%s': %v\n", fileName, err)
		return
	}

	// Create the desired file
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Unable to create '%s': %v\n", fileName, err)
		return
	}

	// Close the file explicity to not waste resources
	if err := file.Close(); err != nil {
		fmt.Printf("Unable to close '%s': %v\n", fileName, err)
		return
	}

	fmt.Printf("Created file '%s'\n", fileName)
}

// newNote creates a new note text and appends it to an existing text file.
func newNote(args []string, notesDir string) {
	// There should be 2 arguements
	if len(args) < 2 {
		fmt.Println("Expected a file name and note string.")
		return
	}

	fileName := args[0]
	path := filepath.Join(notesDir, fileName)

	// Check if the file exists before writing to it
	// For simplisiy, WriteFile, which creates a file and writes to it, won't be used
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Cannot write to file if it doesn't exist.")
		return
	}

	note := strings.Trim(strings.Join(args[1:], " "), `"`) // Drop the quotes and join all strings with space seperators

	// Open the file and write the string into the file
	// Use OpenFile over Open since we don't want to truncate the already existing contents
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		fmt.Printf("Unable to open '%s': %v\n", fileName, err)
		return
	}
	defer file.Close()

	// Fprintln is conventient as it will write to the file and add a newline at the end of the string
	if _, err := fmt.Fprintln(file, note); err != nil {
		fmt.Printf("Error writing note to %q: %v\n", fileName, err)
		return
	}

	fmt.Printf("Added note to %s.\n", fileName)
}
