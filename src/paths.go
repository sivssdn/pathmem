package main

package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Path struct {
	Alias 		string 		`json:"alias"`
	Path		string		`json:"path"`
}


//reads save shortcuts from file
func getSavedAbsPath() {
	_, filename, _, status := runtime.Caller(0)
	if !status {
		panic("Unable to read .env file.")
	}
	//path separator changed as per platform by path/filepath
	aliasFilePath, _ := filepath.Abs(filepath.Dir(filepath.Dir(filename)) + "/data/paths.json")
	fmt.Println(aliasFilePath)
}
