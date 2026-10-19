package commands

import (
	"bufio"
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
func RunNew(s *bufio.Scanner) {
	flag := utility.GetFlag(s)
	args := utility.GetArguements(s)

	notesDir, err := utility.CreateNotesDir()
	if err != nil {
		fmt.Printf("Error creating notes directory: %v\n", err)
		return
	}

	switch flag {
	case newFileFlag:
		newFile(args, notesDir)
	case newNoteFlag:
		newNote(args, notesDir)
	default:
		fmt.Println("Error: invalid flag provided.")
	}
}

func newFile(args []string, notesDir string) {
	// There should be 1 arguement
	if len(args) != 1 {
		fmt.Println("Error: expected a filename")
		return
	}

	fileName := args[0]
	file, err := os.Create(filepath.Join(notesDir, fileName))
	if err != nil {
		fmt.Printf("Error creating file %q: %v\n", fileName, err)
		return
	}
	if err := file.Close(); err != nil {
		fmt.Printf("Error closing file %q: %v\n", fileName, err)
		return
	}
	fmt.Printf("Created file %q\n", fileName)
}

func newNote(args []string, notesDir string) {
	// There should be 2 arguements
	if len(args) < 2 {
		fmt.Println("Error: expected a filename and note")
		return
	}

	// Get the file name and notes
	// Note strings are defined with having quotes around them
	fileName := args[0]
	note := strings.Trim(strings.Join(args[1:], " "), `"`)

	// Open the file and write the note into the file
	file, err := os.OpenFile(filepath.Join(notesDir, fileName), os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		fmt.Printf("Error opening file %q: %v\n", fileName, err)
		return
	}

	if _, err := fmt.Fprintln(file, note); err != nil {
		fmt.Printf("Error writing note to %q: %v\n", fileName, err)
		file.Close()
		return
	}

	if err := file.Close(); err != nil {
		fmt.Printf("Error closing file %q: %v\n", fileName, err)
		return
	}

	fmt.Printf("Added note to %q\n", fileName)
}
