package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Opxourc/note-take/config"
	"github.com/Opxourc/note-take/utility"
)

var flags = map[string]string{
	"File": "f",
	"All":  "a",
}

// RunList runs the List command.
func RunList(s *bufio.Scanner) {
	flag := utility.GetFlag(s)
	found := false

	// Check if the flag is valid
	for _, f := range flags {
		if flag == f {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Error: provided flag is invalid. Use 'help' to get flags for this command.")
		return
	}

	args := utility.GetArguements(s) // Arguements are dependent on the flag

	switch flag {
	case flags["File"]:
		listFile(args)
	case flags["All"]:
		listAll()

	}
}

// listFile prints all notes that are in a file.
func listFile(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: invalid file name provided.")
		return
	}

	fileName := args[0]
	if fileName == "" {
		fmt.Println("Error: invalid file name provided.")
		return
	}

	file, err := os.Open(filepath.Join(config.NoteFolderPath, fileName))
	if err != nil {
		fmt.Printf("Error opening file %q: %v\n", fileName, err)
		return
	}
	defer file.Close()

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
func listAll() {
	entries, err := os.ReadDir(config.NoteFolderPath)
	if err != nil {
		fmt.Printf("Error reading notes directory: %v\n", err)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}
}
