package commands_test

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Opxourc/note-take/commands"
)

func TestRunList(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "expected input for -f",
			input: "-f test.txt",
		},
		{
			name:  "expected input for -a",
			input: "-a",
		},
		{
			name:  "no flag",
			input: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Chdir(t.TempDir())
			scanner := bufio.NewScanner(strings.NewReader(test.input))
			commands.RunList(scanner)
		})
	}
}

func TestRunListCreatesNotesDirAtProjectRoot(t *testing.T) {
	projectRoot := t.TempDir()
	nestedDir := filepath.Join(projectRoot, "commands")
	if err := os.Mkdir(nestedDir, 0o755); err != nil {
		t.Fatalf("failed to create nested directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(projectRoot, "go.mod"), []byte("module example\n"), 0o644); err != nil {
		t.Fatalf("failed to create module file: %v", err)
	}
	t.Chdir(nestedDir)

	commands.RunList(bufio.NewScanner(strings.NewReader("-a")))

	if info, err := os.Stat(filepath.Join(projectRoot, "notes")); err != nil {
		t.Fatalf("notes directory was not created at project root: %v", err)
	} else if !info.IsDir() {
		t.Error("notes path is not a directory")
	}
	if _, err := os.Stat(filepath.Join(nestedDir, "notes")); !os.IsNotExist(err) {
		t.Errorf("notes directory was created in nested directory, stat error = %v", err)
	}
}
