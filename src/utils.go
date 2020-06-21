package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/atotto/clipboard"
)

func checkErr(message string, err error) {
	if err != nil {
		panic(message)
	}
}

//ExecCmd executes any terminal command
func execCmd(command ...string) string {
	execCommand := exec.Command(command[0], command[1:]...)
	commandOutput, err := execCommand.Output()
	if err != nil {
		fmt.Println("Output from terminal : ", err)
	}
	return string(commandOutput)
}

//opens the directory path & execute the command
func openDirectoryAndExecute(dirPath string, command string) {
	cmd := exec.Command("osascript", "-e", `tell application "Terminal" to do script "cd `+dirPath+command+`"`)
	cmd.Run()
}

//copies input string to clipboard
func copyToClipboard(toSave string) {
	if len(toSave) > 0 {
		clipboard.WriteAll(toSave)
	}
}

//gets length of largest string in an array
func getMaxLength(slice []string) int {
	maxLen := 0
	for _, value := range slice {
		if len(value) > maxLen {
			maxLen = len(value)
		}
	}
	return maxLen
}

//removes new line character from string
func removeNewLine(s string) string {
	return strings.Replace(s, "\n", "", -1)
}
