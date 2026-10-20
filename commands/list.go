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
func RunList(input []string) {
	flag, args := utility.ParseFlagArgs(input)
	dir, err := utility.CreateNotesDir()
	if err != nil {
		fmt.Printf("Unable to create notes directory: %v\n", err)
		return
	}

	switch flag {
	case listFileFlag:
		listFile(args, dir)
	case listAllFlag:
		listAll(dir)
	default:
		fmt.Println("Invalid flag provided.")
	}
}

// listFile prints all notes that are in a file.
func listFile(args []string, dir string) {
	// There should be 1 arguement
	if len(args) != 1 {
		fmt.Println("Expected a file name.")
		return
	}

	fileName := args[0]

	// Find and open the file by joining paths
	// Open over OperFile since Open conventiely focuses on read only cases
	file, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		fmt.Printf("Unable to open '%s': %v\n", fileName, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Print if there's no notes to scan
	if !scanner.Scan() {
		fmt.Println("There's no notes to list.")
		return
	}

	// Print all notes from this file
	// Track the line number seperately since there's no iterator value in the loop
	lineNumber := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", lineNumber, scanner.Text())
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error occured while reading '%s': %v\n", fileName, err)
		return
	}
}

// listAll prints all file names that exist in the notes folder.
func listAll(dir string) {
	// Read all files contained in the directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Unable to read notes folder: %s", err)
		return
	}

	// Print if there's no entries
	if len(entries) == 0 {
		fmt.Println("There's no note files to list.")
	}

	// Print the names of all files
	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}
}
