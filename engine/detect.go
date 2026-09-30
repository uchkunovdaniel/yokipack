package engine

import (
	"fmt"
	"log"
	"strings"

	"go.yaml.in/yaml/v4"
)

var fs = NewOSFileSystem(".")

type Languages struct {
	Name       string   `yaml:"-"`
	Extensions []string `yaml:"extensions"`
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
	//files := fs.WalkDir(".", &WalkDirOptions{
	//	SkipDirs:  []string{".git", ".idea", ".vscode", ".md", "LICENSE"},
	//	FilesOnly: true,
	//})
	for _, lang := range parseLanguages() {
		fmt.Println(lang.Name, lang.Extensions)
	}

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
