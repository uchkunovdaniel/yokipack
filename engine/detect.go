package engine

import (
	"fmt"
	"strings"
)

func DetectLookupTable() {
	initBloomFilter()
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
