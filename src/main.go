package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//reding command line args
	cmdArgs := os.Args[1:]
	if len(cmdArgs) < 1 {
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
		toExecute := ""
		if len(dirPath) > 0 {
			if len(args) >= 4 && args[2] == "-e" {
				toExecute = strings.Join(args[3:], " ")
			}
			openDirectoryAndExecute(dirPath, toExecute)
			return
		}
		fmt.Println("Alias not found")
	case "copy", "-c":
		dirPath := getPath(pathsDataFilePath, args[1])
		copyToClipboard(dirPath)
		fmt.Println(dirPath, " - copied to clipboard")
	case "paths", "-p":
		printAliases(pathsDataFilePath)
	default:
		fmt.Println("Unknown command.")
	}
}
