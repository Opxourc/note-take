package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewCreatesFileAndAddsNote(t *testing.T) {
	projectRoot := t.TempDir()
	nestedDir := filepath.Join(projectRoot, "commands")
	if err := os.Mkdir(nestedDir, 0o755); err != nil {
		t.Fatalf("failed to create nested directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(projectRoot, "go.mod"), []byte("module example\n"), 0o644); err != nil {
		t.Fatalf("failed to create module file: %v", err)
	}
	t.Chdir(nestedDir)

	RunNew([]string{"-f", "sample"})

	path := filepath.Join(projectRoot, "notes", "sample")
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("New() did not create file: %v", err)
	}
	if _, err := os.Stat(filepath.Join(nestedDir, "notes")); !os.IsNotExist(err) {
		t.Errorf("notes directory was created in nested directory, stat error = %v", err)
	}

	RunNew([]string{"-n", "sample", "hello world"})

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read created file: %v", err)
	}
	if got, want := string(content), "hello world\n"; got != want {
		t.Errorf("file content = %q, want %q", got, want)
	}
}
