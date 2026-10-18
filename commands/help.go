package commands

import "fmt"

func Help() {
	fmt.Println("Commands:")
	fmt.Println("  help - Display this help message")
	fmt.Println("  exit - Exit the program")
	fmt.Println("  new -f (filepath) - Create a new note file")
	fmt.Println("  new -n (filepath) (notestring) - Create a new note for a file")
	fmt.Println("  list -f (filepath) - List notes for a file")
	fmt.Println("  list -a - List all notes")
}
