package main

import (
	"github.com/timo-reymann/nereide/templating/util"
	"os"
	"strings"
	"text/template"
)

const TemplateFileRoot = "./templates/"
const RenderedTemplateFilesRoot = "/templates/"

func createTemplate(fileName string) (*template.Template, error) {
	content, err := os.ReadFile(TemplateFileRoot + fileName)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New(fileName).Parse(string(content))
	if err != nil {
		panic(err)
	}

	return tpl, nil
}

func writeTemplate(tpl *template.Template, targetFileName string, data interface{}) error {
	file, err := os.Create(RenderedTemplateFilesRoot + targetFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = tpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	i18nMappings, err := util.ReadJSONFilesInFolder("i18n")
	check(err)

	htmlTpl, err := createTemplate("index.html")
	check(err)

	xmlTpl, err := createTemplate("index.xml")
	check(err)

	jsonTpl, err := createTemplate("index.json")
	check(err)

	for fileName, config := range i18nMappings {
		lang := strings.TrimSuffix(fileName, ".json")

		err := writeTemplate(htmlTpl, "index."+lang+".html", config["html"])
		check(err)

		err = writeTemplate(xmlTpl, "index."+lang+".xml", config["xml"])
		check(err)

		err = writeTemplate(jsonTpl, "index."+lang+".json", config["json"])
		check(err)

	}
}
