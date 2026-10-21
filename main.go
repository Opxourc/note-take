package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Opxourc/note-take/commands"
)

const (
	deleteCommand = "delete"
	helpCommand   = "help"
	exitCommand   = "exit"
	listCommand   = "list"
	newCommand    = "new"
)

func main() {
	fmt.Println("Note-Take CLI")
	fmt.Println("Type 'help' to see commands.")

	scanner := bufio.NewScanner(os.Stdin) // Should be a scanner for the terminal
	exiting := false                      // For when the application needs to quit

	// Keep accepting commands until the user exits
	for !exiting {
		fmt.Print("> ")

		// Stop if Stdin closes or scanning fails
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()             // Input will have components (command) (flag) (args) ...
		components := strings.Fields(input) // Break the input into those respective components

		// Ignore empty or whitespace-only input
		if len(components) < 1 {
			continue
		}

		command := components[0]       // Grab the command
		dependencies := components[1:] // Everything after are dependencies

		switch command {
		case exitCommand:
			exiting = true
			fmt.Println("Exiting Note-Take CLI")

		case helpCommand:
			commands.RunHelp()

		case listCommand:
			commands.RunList(dependencies)

		case newCommand:
			commands.RunNew(dependencies)

		case deleteCommand:
			commands.RunDelete(dependencies)

		default:
			fmt.Printf("No command called '%s'\n\n", command)
		}
	}

	// Report any error encountered while reading Stdin
	if err := scanner.Err(); err != nil {
		fmt.Println("Error occured while reading inputs.")
	}
}
