package commands

import (
	"fmt"
	"os"
)

// RunExit runs the 'exit' command, allowing the program to exit.
func RunExit() {
	fmt.Print("\nExitting Note-Take CLI.")
	os.Exit(0)
}
