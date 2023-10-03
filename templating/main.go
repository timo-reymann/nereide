package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"io"
	"path/filepath"
	"strings"
)

func readJSONFilesInFolder(folderPath string) (map[string]interface{}, error) {
	fileData := make(map[string]interface{})

	// Get a list of all files in the folder
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(folderPath, file.Name())

			// Read the JSON file
			jsonFile, err := os.Open(filePath)
			if err != nil {
				return nil, err
			}
			defer jsonFile.Close()

			// Decode the JSON data into an interface{}
			var jsonData interface{}
			decoder := json.NewDecoder(jsonFile)
			if err := decoder.Decode(&jsonData); err != nil {
				return nil, err
			}

			// Add the data to the map with the file name as the key
			fileData[file.Name()] = jsonData
		}
	}

	return fileData, nil
}

func copyFile(srcPath, destPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	landingContent, err := os.ReadFile("./landing.html")
	if err != nil {
		panic(err)
	}

	i81nMappings, err := readJSONFilesInFolder("i18n")
	if err != nil {
		panic(err)
	}

	t, err := template.New("landing").Parse(string(landingContent))
	if err != nil {
		panic(err)
	}

	for fileName, config := range i81nMappings {
	    lang := strings.TrimSuffix(fileName, ".json")
		file, err := os.Create("/templates/index." + lang + ".html")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = t.ExecuteTemplate(file, "landing", config)
		if err != nil {
			panic(err)
		}

		if lang == "de" {
		    copyFile("/templates/index." + lang + ".html", "/templates/index.html")
		}
	}

}
