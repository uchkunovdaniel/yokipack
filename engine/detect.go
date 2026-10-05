package engine

import (
	"fmt"
	"strings"
)

var fs = NewOSFileSystem(".")
var files = fs.walkDir(".", &WalkDirOptions{
	SkipDirs:  fs.readIgnoreFile(),
	FilesOnly: true,
})
var languages = parseLanguages("engine/languages.yml")

func DetectLookupTable() {
	for _, f := range files {
		fmt.Println(isExtensionInBloomFilter(f), f)
	}
}

func DetectContainer() string {
	for idx, d := range files {
		if strings.Contains(d, "Dockerfile") || strings.Contains(d, "Containerfile") {
			return files[idx]
		}
	}
	return ""
}
