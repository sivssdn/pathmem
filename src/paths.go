package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

type path struct {
	Alias string `json:"alias"`
	Path  string `json:"path"`
}

//reads save shortcuts from file
func getSavedAbsPath() string {
	_, filename, _, status := runtime.Caller(0)
	if !status {
		panic("Unable to read .env file.")
	}
	//path separator changed as per platform by path/filepath
	aliasFilePath, _ := filepath.Abs(filepath.Dir(filepath.Dir(filename)) + "/data/paths.json")
	return aliasFilePath
}

//gets path from file matchng the alias name
func getPath(filePath string, alias string) string {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("Couldn't find saved paths file")
	}
	var paths []path
	err = json.Unmarshal(file, &paths)
	for _, path := range paths {
		if strings.EqualFold(path.Alias, alias) {
			return path.Path
		}
	}
	return "not found"
}
