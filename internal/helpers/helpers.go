package helpers

import (
	"io/ioutil"
	"path/filepath"
)

// Function for helping loading files (especially used for loading html files)

func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Function for deciding if given data is empty
func IsEmpty(data string) bool {
	return len(data) == 0
}

func JoinPaths(paths ...string) string {
	return filepath.Join(paths...)
}

func GetTemplate(path string) string {
	return filepath.Join("web", "templates", path)
}
