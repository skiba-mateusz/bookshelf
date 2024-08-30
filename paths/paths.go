package paths

import (
	"os"
	"path/filepath"
)

// Returns the filepath to books.json file and creates it if doesn't exist
func GetBooksJsonFile() (string, error) {
	homerDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDirPath := filepath.Join(homerDir, "/.bookshelf")

	err = os.MkdirAll(configDirPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	jsonFilePath := filepath.Join(configDirPath, "books.json")
	
	_, err = os.Stat(jsonFilePath);
	if err != nil {
		switch {
		case os.IsNotExist(err):
			file, err := os.Create(jsonFilePath)
			if err != nil {
				return "", err
			}
			file.Close()
		default:
			return "", err
		}
	} 

	return jsonFilePath, nil
}