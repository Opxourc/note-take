package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Opxourc/note-take/utility"
)

const (
	deleteFileFlag = "f"
	deleteNoteFlag = "n"
)

// RunDelete deletes a file or a note depending on the flag.
func RunDelete(input []string) {
	flag, args := utility.ParseFlagArgs(input)
	dir, err := utility.CreateNotesDir()
	if err != nil {
		fmt.Printf("Unable to create notes directory: %v\n", err)
		return
	}

	// Route the request to the matching file or note deletion operation.
	switch flag {
	case deleteFileFlag:
		deleteFile(args, dir)
	case deleteNoteFlag:
		deleteNote(args, dir)
	default:
		fmt.Println("Invalid flag provided.")
	}
}

// deleteFile deletes the existing provided file.
func deleteFile(args []string, dir string) {
	// File deletion needs exactly one argument: the file's name
	if len(args) != 1 {
		fmt.Println("Expected a file name.")
		return
	}

	fileName := args[0]

	// Build the target path inside the notes folder and make sure it is a file
	path := filepath.Join(dir, fileName)
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Unable to find '%s': %v\n", fileName, err)
		return
	}
	if info.IsDir() {
		fmt.Printf("'%s' is a directory, not a note file.\n", fileName)
		return
	}

	// Remove the file only after validating the target
	if err := os.Remove(path); err != nil {
		fmt.Printf("Unable to delete '%s': %v\n", fileName, err)
		return
	}

	fmt.Printf("Deleted file '%s'.\n", fileName)
}

// deleteNote removes the one-based note index shown by the list command.
func deleteNote(args []string, dir string) {
	// A note deletion identifies both the containing file and the note's displayed index
	if len(args) != 2 {
		fmt.Println("Expected a file name and note index.")
		return
	}

	// Convert the supplied index to an integer; list output starts at 1
	index, err := strconv.Atoi(args[1])
	if err != nil || index < 1 {
		fmt.Println("Note index must be a positive number.")
		return
	}

	fileName := args[0]

	// Read each line as one note so its position matches the list command's numbering
	path := filepath.Join(dir, fileName)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Unable to open '%s': %v\n", fileName, err)
		return
	}

	scanner := bufio.NewScanner(file)
	var notes []string
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	// Capture read and close errors before continuing with the collected notes
	scanErr := scanner.Err()
	closeErr := file.Close()
	if scanErr != nil {
		fmt.Printf("Unable to read '%s': %v\n", fileName, scanErr)
		return
	}

	if closeErr != nil {
		fmt.Printf("Unable to close '%s': %v\n", fileName, closeErr)
		return
	}

	// Reject indexes that do not correspond to a note in this file
	if index > len(notes) {
		fmt.Printf("Note index %d is out of range; '%s' has %d notes.\n", index, fileName, len(notes))
		return
	}

	// Convert the one-based display index to a zero-based slice position and remove it
	notes = append(notes[:index-1], notes[index:]...)

	// Rebuild the file with the usual trailing newline for non-empty note lists
	content := strings.Join(notes, "\n")
	if len(notes) > 0 {
		content += "\n"
	}

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		fmt.Printf("Unable to update '%s': %v\n", fileName, err)
		return
	}

	fmt.Printf("Deleted note %d from %s.\n", index, fileName)
}
