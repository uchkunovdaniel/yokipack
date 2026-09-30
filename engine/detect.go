package engine

import (
	"fmt"
	"strings"
)

var fs = NewOSFileSystem(".")

func detectLookupTable() {
	fmt.Println(fs.WalkDir(".", &WalkDirOptions{
		SkipDirs:  []string{".git", ".idea", ".vscode", ".md", "LICENSE"},
		FilesOnly: true,
	}))
}

func detectContainer() string {
	dir := fs.WalkDir(".", &WalkDirOptions{
		SkipDirs:  []string{".git", ".idea", ".vscode", ".md", "LICENSE"},
		FilesOnly: true,
	})

	for idx, d := range dir {
		if strings.Contains(d, "Dockerfile") || strings.Contains(d, "Containerfile") {
			return dir[idx]
		}
	}
	return ""
}
