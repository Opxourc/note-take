package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Opxourc/note-take/commands"
)

func main() {
	title()
	scanner := bufio.NewScanner(os.Stdin) // Want to scan from a terminal

	exit := false
	for !exit {
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}
		if input == "" {
			continue
		}

		// Check the first token before the first space
		first := input[:strings.Index(input, " ")]
		switch first {
		case "help":
			commands.Help()
		}
	}
}

func title() {
	fmt.Println("--------------------------")
	fmt.Println("|	   Note-Take CLI 	  |")
	fmt.Println("--------------------------")
	fmt.Print("\nType 'help' to see commands.\n\n")
}
