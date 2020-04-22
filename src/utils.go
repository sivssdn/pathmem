package main

import (
	"fmt"
	"os/exec"

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

func openDirectory(dirPath string) {
	cmd := exec.Command("osascript", "-s", "h", "-e", `tell application "Terminal" to do script "cd `+dirPath+`"`)
	cmd.Run()
}

//copies input string to clipboard
func copyToClipboard(toSave string) {
	if len(toSave) > 0 {
		clipboard.WriteAll(toSave)
	}
}
