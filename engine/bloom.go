package engine

import (
	"slices"
	"strings"

	"github.com/bits-and-blooms/bloom/v3"
)

var bloomFilter = bloom.New(1000000, 1)

func initBloomFilter() {
	for _, lang := range languages {
		bloomFilter.Add([]byte(lang.Extensions[0]))
	}
}

func isLanguageInProject(file string) bool {
	fileArray := strings.Split(file, ".")
	last := len(fileArray) - 1
	extension := "." + fileArray[last]

	return bloomFilter.Test([]byte(extension))
}

func isLanguageKnown() bool {
	initBloomFilter()

	return slices.ContainsFunc(files, isLanguageInProject)
}
