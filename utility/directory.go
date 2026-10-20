package utility

import (
	"os"
	"path/filepath"

	"github.com/Opxourc/note-take/config"
)

// CreateNotesDir creates the notes directory at the project root and returns its path.
func CreateNotesDir() (string, error) {
	root, err := getProjectRoot()
	if err != nil {
		return "", err
	}

	// Get the path of where the notes will be stored
	// Create the path it doesn't already exist
	path := filepath.Join(root, config.NoteFolderPath)
	if err := os.MkdirAll(path, 0o755); err != nil {
		return "", err
	}

	return path, nil
}

// getProjectRoot returns the project's root where go.mod is located.
func getProjectRoot() (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// Check if go.mod exists in the current directory
		// If there's a error that is not that go.mod exists, then just return
		if _, err := os.Stat(filepath.Join(root, "go.mod")); err == nil {
			return root, nil
		} else if !os.IsNotExist(err) {
			return "", err
		}

		parent := filepath.Dir(root) // Get the parent of the root
		if parent == root {
			return os.Getwd() // Fallback to just to the current directory
		}

		root = parent // Move up the tree
	}
}
