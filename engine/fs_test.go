package engine

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestOSFileSystemHelpers(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmpDir, "visible.txt"), []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := NewOSFileSystem(tmpDir)

	if !fs.exists("visible.txt") {
		t.Fatalf("expected file to exist")
	}
	if fs.exists("missing.txt") {
		t.Fatalf("expected missing file to not exist")
	}

	if got := string(fs.readFile("visible.txt")); got != "hello" {
		t.Fatalf("unexpected file content: %q", got)
	}
}

func TestOSFileSystemGlob(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmpDir, "a.txt"), []byte("a"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "b.md"), []byte("b"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := NewOSFileSystem(tmpDir)
	matches := fs.glob("*.txt")
	if len(matches) != 1 {
		t.Fatalf("expected 1 match, got %d", len(matches))
	}
	if filepath.Base(matches[0]) != "a.txt" {
		t.Fatalf("unexpected match: %s", matches[0])
	}
}

func TestOSFileSystemWalkDirWithSkipAndFilesOnly(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(tmpDir, "nested"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(tmpDir, "skipdir"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "nested", "keep.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "skipdir", "skip.txt"), []byte("y"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := NewOSFileSystem(tmpDir)
	got := fs.walkDir(".", &WalkDirOptions{
		SkipDirs:  []string{"skipdir"},
		FilesOnly: true,
	})

	for _, path := range got {
		info, err := os.Stat(path)
		if err != nil {
			t.Fatal(err)
		}
		if info.IsDir() {
			t.Fatalf("expected files only, got directory: %s", path)
		}
	}

	paths := make([]string, 0, len(got))
	for _, p := range got {
		paths = append(paths, filepath.Base(p))
	}
	if !slices.Contains(paths, "keep.txt") {
		t.Fatalf("expected keep.txt to be returned")
	}
	if slices.Contains(paths, "skip.txt") {
		t.Fatalf("expected skip.txt to be skipped")
	}
}

func TestOSFileSystemReadIgnoreFile(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmpDir, ".yokipackignore"), []byte("node_modules\ndist\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := NewOSFileSystem(tmpDir)
	got := fs.readIgnoreFile()
	want := []string{"node_modules", "dist", ""}

	if len(got) != len(want) {
		t.Fatalf("unexpected number of ignore entries: got %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("entry %d mismatch: got %q, want %q", i, got[i], want[i])
		}
	}
}
