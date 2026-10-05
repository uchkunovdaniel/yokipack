package engine

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseLanguages(t *testing.T) {
	tmpDir := t.TempDir()
	content := []byte("go:\n  extensions:\n    - go\npython:\n  extensions:\n    - py\n")
	if err := os.WriteFile(filepath.Join(tmpDir, "languages.yml"), content, 0o644); err != nil {
		t.Fatal(err)
	}

	prevFS := fs
	fs = NewOSFileSystem(tmpDir)
	t.Cleanup(func() {
		fs = prevFS
	})

	got := parseLanguages("languages.yml")
	if len(got) != 2 {
		t.Fatalf("unexpected number of parsed languages: got %d, want 2", len(got))
	}

	parsed := map[string][]string{}
	for _, lang := range got {
		parsed[lang.Name] = lang.Extensions
	}

	if exts, ok := parsed["go"]; !ok || len(exts) != 1 || exts[0] != "go" {
		t.Fatalf("unexpected go language value: %#v", parsed["go"])
	}
	if exts, ok := parsed["python"]; !ok || len(exts) != 1 || exts[0] != "py" {
		t.Fatalf("unexpected python language value: %#v", parsed["python"])
	}
}
