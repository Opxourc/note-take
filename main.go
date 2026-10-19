package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Opxourc/note-take/commands"
)

const (
	helpCommand = "help"
	exitCommand = "exit"
	listCommand = "list"
	newCommand  = "new"
)

func main() {
	fmt.Println("Note-Take CLI")
	fmt.Print("Type 'help' to see commands.\n\n")
	scanner := bufio.NewScanner(os.Stdin) // Want to scan from a terminal

	for {
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}
		if input == "" {
			continue
		}

		// Check the first token before the first space
		com := strings.Fields(input)
		if len(com) == 0 {
			continue
		}
		first := com[0]

		switch first {
		case helpCommand:
			commands.RunHelp()
		case exitCommand:
			commands.RunExit()
		case listCommand:
			commands.RunList(scanner)
		case newCommand:
			commands.RunNew(scanner)
		default:
			fmt.Printf("No command called %s", first)
		}
	}
}
