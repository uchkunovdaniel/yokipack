package engine

import (
	"fmt"
	"log"
	"strings"

	"github.com/bits-and-blooms/bloom/v3"
	"go.yaml.in/yaml/v4"
)

type Languages struct {
	Name       string   `yaml:"-"`
	Extensions []string `yaml:"extensions"`
}

var fs = NewOSFileSystem(".")

var files = fs.WalkDir(".", &WalkDirOptions{
	SkipDirs:  []string{},
	FilesOnly: true,
})

var bloomFilter = bloom.New(1000000, 1)

var langs []Languages = parseLanguages()

func initBloomFilter() {
	for _, l := range langs {
		bloomFilter.Add([]byte(l.Extensions[0]))
	}
}
func isExtensionInBloomFilter(file string) bool {
	extension := strings.Split(file, ".")[len(strings.Split(file, "."))-1]
	return bloomFilter.Test([]byte(extension))
}

func parseLanguages() []Languages {
	var raw map[string]Languages
	if err := yaml.Unmarshal([]byte(fs.ReadFile("engine/languages.yml")), &raw); err != nil {
		log.Fatal(err)
	}

	langs := make([]Languages, 0, len(raw))
	for name, lang := range raw {
		lang.Name = name
		langs = append(langs, lang)
	}
	return langs
}

func DetectLookupTable() {
	initBloomFilter()

	for _, f := range files {
		fmt.Println(isExtensionInBloomFilter(f), f)
	}
}

func detectContainer() string {
	for idx, d := range files {
		if strings.Contains(d, "Dockerfile") || strings.Contains(d, "Containerfile") {
			return files[idx]
		}
	}
	return ""
}
