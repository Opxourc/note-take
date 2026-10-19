package utility

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateNotesDirAtProjectRoot(t *testing.T) {
	projectRoot := t.TempDir()
	nestedDir := filepath.Join(projectRoot, "commands")
	if err := os.Mkdir(nestedDir, 0o755); err != nil {
		t.Fatalf("failed to create nested directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(projectRoot, "go.mod"), []byte("module example\n"), 0o644); err != nil {
		t.Fatalf("failed to create module file: %v", err)
	}
	t.Chdir(nestedDir)

	got, err := CreateNotesDir()
	if err != nil {
		t.Fatalf("CreateNotesDir() error = %v", err)
	}

	want := filepath.Join(projectRoot, "notes")
	if got != want {
		t.Errorf("CreateNotesDir() = %q, want %q", got, want)
	}
	if info, err := os.Stat(want); err != nil {
		t.Fatalf("notes directory was not created at project root: %v", err)
	} else if !info.IsDir() {
		t.Errorf("notes path is not a directory")
	}
}
