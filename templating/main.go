package main

import (
	"github.com/timo-reymann/nereide/templating/util"
	"html/template"
	"os"
	"strings"
)

const DefaultLanguage = "de"
const TemplateFileRoot = "./templates"
const RenderedTemplateFilesRoot = "/templates"

func main() {
	landingContent, err := os.ReadFile(TemplateFileRoot + "/landing.html")
	if err != nil {
		panic(err)
	}

	i81nMappings, err := util.ReadJSONFilesInFolder("i18n")
	if err != nil {
		panic(err)
	}

	t, err := template.New("landing").Parse(string(landingContent))
	if err != nil {
		panic(err)
	}

	for fileName, config := range i81nMappings {
		lang := strings.TrimSuffix(fileName, ".json")
		file, err := os.Create(RenderedTemplateFilesRoot + "/index." + lang + ".html")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = t.ExecuteTemplate(file, "landing", config)
		if err != nil {
			panic(err)
		}
	}

}
