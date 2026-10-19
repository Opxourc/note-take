package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Opxourc/note-take/utility"
)

const (
	listFileFlag = "f"
	listAllFlag  = "a"
)

// RunList runs the 'list' command by either listing all files or notes from a file.
func RunList(s *bufio.Scanner) {
	flag := utility.GetFlag(s)
	args := utility.GetArguements(s) // Arguements are dependent on the flag

	notesDir, err := utility.CreateNotesDir()
	if err != nil {
		fmt.Printf("Error creating notes directory: %v\n", err)
		return
	}

	switch flag {
	case listFileFlag:
		listFile(args, notesDir)
	case listAllFlag:
		listAll(notesDir)
	default:
		fmt.Println("Error: invalid flag provided.")
	}
}

// listFile prints all notes that are in a file.
func listFile(args []string, notesDir string) {
	if len(args) == 0 {
		fmt.Println("Error: no arguements were provided.")
		return
	}

	// First arguement should be the name that's provided
	fileName := args[0]
	if fileName == "" {
		fmt.Println("Error: invalid file name provided.")
		return
	}

	// Find and open the file by joining paths
	file, err := os.Open(filepath.Join(notesDir, fileName))
	if err != nil {
		fmt.Printf("Error opening file %q: %v\n", fileName, err)
		return
	}
	defer file.Close()

	// Scan every line of text in the text
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", lineNumber, scanner.Text())
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %q: %v\n", fileName, err)
		return
	}
}

// listAll prints all file names that exist in the notes folder.
func listAll(notesDir string) {
	// Read all files contained in the directory
	entries, err := os.ReadDir(notesDir)
	if err != nil {
		fmt.Printf("Unable to read notes folder: %s", err)
		return
	}

	// Print the names of all files
	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}
}
