package engine

import (
	"fmt"
	"strings"
)

func DetectLookupTable() {
	fmt.Println(isLanguageKnown())
}

func DetectContainer() string {
	for idx, d := range files {
		if strings.Contains(d, "Dockerfile") || strings.Contains(d, "Containerfile") {
			return files[idx]
		}
	}
	return ""
}
