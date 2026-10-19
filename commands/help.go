package commands

import "fmt"

// RunHelp runs the 'help' command and prints commands and usages.
func RunHelp() {
	fmt.Println("Commands:")
	fmt.Println("If a component of a command has a question mark, it means that it may be optional in certain usages.")
	fmt.Println()

	fmt.Println("  	help - Display this help message.")
	fmt.Println("  	exit - Exit the program.")
	fmt.Println()

	fmt.Println("	new (flag) (file name) (note string)?")
	fmt.Println("		Flags:")
	fmt.Println("  			-f - File, creates a new file with the given name. Note string isn't needed.")
	fmt.Println("			-n - Note, creates a new note to a file name. The file should exist before writing to it.")
	fmt.Println()

	fmt.Println("	list (flag) (file name)?")
	fmt.Println("		Flags:")
	fmt.Println("  			-f - File, prints all notes from a given file name.")
	fmt.Println("  			-a - All, prints all note file names, not their contents. No file name is needed.")
	fmt.Println()
}
