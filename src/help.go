package main

import (
	"fmt"
	"strings"
)

func printAliases(filePath string) {
	paths := getJSONContents(filePath)
	aliases := []string{}
	for _, path := range paths {
		aliases = append(aliases, path.Alias)
	}
	maxAliasLen := getMaxLength(aliases) + 2
	padding := 0
	paddingStart := strings.Join(make([]string, 4), " ")
	var aliasRow string
	fmt.Println("\n Saved Paths:")
	//formatting & printing the table
	for _, path := range paths {
		padding = maxAliasLen - len(path.Alias)
		aliasRow = paddingStart + path.Alias + strings.Join(make([]string, padding), " ") + " : " + path.Path + "\n"
		fmt.Printf(aliasRow)
	}
	fmt.Println("\nYou can always add more in paths.json")
}
