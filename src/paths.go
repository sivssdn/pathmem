package main

import (
	"encoding/json"
	"fmt"
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
	checkErr("Couldn't find saved paths file", err)
	var paths []path
	err = json.Unmarshal(file, &paths)
	checkErr("Error in saved paths (json file)", err)
	for _, path := range paths {
		if strings.EqualFold(path.Alias, alias) {
			return path.Path
		}
	}
	return ""
}

//saves alias to corresponding absolute path, to json
func savePath(filePath string, alias string, absPath string) {
	file, err := ioutil.ReadFile(filePath)
	checkErr("Couldn't find saved paths file", err)

	var savedPaths []path
	err = json.Unmarshal([]byte(file), &savedPaths)
	checkErr("Error in saved paths (json file)", err)
	newSavedpaths := append(savedPaths, path{Alias: alias, Path: absPath})

	result, err := json.Marshal(newSavedpaths)
	checkErr("Error appending to json", err)
	err = ioutil.WriteFile(filePath, result, 644)
	checkErr("Error appending to json", err)
	fmt.Println("Filepath ", absPath, " saved with alias ", alias)
}
