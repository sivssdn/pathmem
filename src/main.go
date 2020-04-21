package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("---", err)
		}
	}()

	fmt.Println("Project initiated!")
	var absPath string = getSavedAbsPath()
	fmt.Println(getPath(absPath, "hc"))
}
