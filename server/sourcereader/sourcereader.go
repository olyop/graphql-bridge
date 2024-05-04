package sourcereader

import (
	"os"
	"path/filepath"
)

// ReadFolder reads all .graphqls files in a folder
func ReadFolder(folderPath string) map[string]string {
	files := map[string]string{}

	err := filepath.Walk(folderPath, readFile(files))
	if err != nil {
		panic(err)
	}

	return files
}

func readFile(files map[string]string) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".graphqls" {
			return nil
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		files[path] = string(contents)

		return nil
	}
}
