package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//reding command line args
	cmdArgs := os.Args[1:]
	if len(cmdArgs) < 2 {
		fmt.Println("Please specify program arguments")
		return
	}
	router(cmdArgs...)
}

func router(args ...string) {
	var pathsDataFilePath string = getSavedAbsPath()

	switch args[0] {
	case "save", "-s":
		currentDirPath := execCmd("pwd")
		savePath(pathsDataFilePath, args[1], currentDirPath)
	case "open", "-o":
		dirPath := getPath(pathsDataFilePath, args[1])
		if len(dirPath) > 0 {
			openDirectory(dirPath)
			return
		}
		fmt.Println("Alias not found")
	default:
		fmt.Println("Unknown command.")
	}
}
