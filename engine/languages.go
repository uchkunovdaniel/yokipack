package engine

import (
	"log"

	"go.yaml.in/yaml/v4"
)

type Languages struct {
	Name       string   `yaml:"-"`
	Extensions []string `yaml:"extensions"`
}

func parseLanguages(path string) []Languages {
	var raw map[string]Languages
	if err := yaml.Unmarshal(fs.readFile(path), &raw); err != nil {
		log.Fatal(err)
	}

	langs := make([]Languages, 0, len(raw))
	for name, lang := range raw {
		lang.Name = name
		langs = append(langs, lang)
	}
	return langs
}
