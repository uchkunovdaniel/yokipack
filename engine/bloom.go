package engine

import (
	"strings"

	"github.com/bits-and-blooms/bloom/v3"
)

var bloomFilter = bloom.New(1000000, 1)

func initBloomFilter() {
	for _, l := range languages {
		bloomFilter.Add([]byte(l.Extensions[0]))
	}
}

func isExtensionInBloomFilter(file string) bool {
	extension := strings.Split(file, ".")[len(strings.Split(file, "."))-1]
	return bloomFilter.Test([]byte(extension))
}
