package commands

import (
	"fmt"
	"os"
)

// Exit ends the program.
func Exit() {
	fmt.Print("\nExitting Note-Take CLI.")
	os.Exit(0)
}
