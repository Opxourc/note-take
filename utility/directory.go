package utility

import (
	"os"
	"path/filepath"

	"github.com/Opxourc/note-take/config"
)

// CreateNotesDir creates the notes directory at the project root and returns its path.
func CreateNotesDir() (string, error) {
	root, err := projectRoot()
	if err != nil {
		return "", err
	}

	path := filepath.Join(root, config.NoteFolderPath)
	if err := os.MkdirAll(path, 0o755); err != nil {
		return "", err
	}
	return path, nil
}

func projectRoot() (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(root, "go.mod")); err == nil {
			return root, nil
		} else if !os.IsNotExist(err) {
			return "", err
		}

		parent := filepath.Dir(root)
		if parent == root {
			return os.Getwd()
		}
		root = parent
	}
}
