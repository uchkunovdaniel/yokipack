package engine

import (
	"testing"

	"github.com/bits-and-blooms/bloom/v3"
)

func TestDetectContainer(t *testing.T) {
	prevFiles := files
	files = []string{
		"src/main.go",
		"Dockerfile",
		"infra/Containerfile.dev",
	}
	t.Cleanup(func() {
		files = prevFiles
	})

	if got := DetectContainer(); got != "Dockerfile" {
		t.Fatalf("unexpected detected container file: got %q, want %q", got, "Dockerfile")
	}
}

func TestDetectContainerNoFile(t *testing.T) {
	prevFiles := files
	files = []string{"src/main.go", "README.md"}
	t.Cleanup(func() {
		files = prevFiles
	})

	if got := DetectContainer(); got != "" {
		t.Fatalf("expected no detected container file, got %q", got)
	}
}

func TestBloomFilterLookup(t *testing.T) {
	prevLanguages := languages
	prevBloom := bloomFilter
	languages = []Languages{
		{Name: "go", Extensions: []string{"go"}},
		{Name: "python", Extensions: []string{"py"}},
	}
	bloomFilter = bloom.New(1000000, 1)

	t.Cleanup(func() {
		languages = prevLanguages
		bloomFilter = prevBloom
	})

	initBloomFilter()

	if !isExtensionInBloomFilter("main.go") {
		t.Fatalf("expected go extension to be present in bloom filter")
	}
	if !isExtensionInBloomFilter("script.py") {
		t.Fatalf("expected py extension to be present in bloom filter")
	}
}
