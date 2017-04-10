package gfile

import (
	"io/ioutil"
	"os"
	"strings"
)

func Getfile(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		os.Exit(1)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		paths = append(paths, file.Name())
	}
	return paths
}
