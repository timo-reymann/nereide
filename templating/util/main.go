package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadJSONFilesInFolder(folderPath string) (map[string]map[string]interface{}, error) {
	fileData := make(map[string]map[string]interface{})

	// Get a list of all files in the folder
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		filePath := filepath.Join(folderPath, file.Name())

		// Read the JSON file
		jsonFile, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer jsonFile.Close()

		// Decode the JSON data into an interface{}
		var jsonData map[string]interface{}
		decoder := json.NewDecoder(jsonFile)
		if err := decoder.Decode(&jsonData); err != nil {
			return nil, err
		}

		// Add the data to the map with the file name as the key
		fileData[file.Name()] = jsonData
	}

	return fileData, nil
}
