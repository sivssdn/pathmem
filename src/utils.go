package main

import (
	"fmt"
	"os/exec"
)

func checkErr(message string, err error) {
	if err != nil {
		panic(message)
	}
}

//ExecCmd executes any terminal command
func execCmd(command ...string) string {
	execCmd := exec.Command(command[0], command...)
	commandOutput, err := execCmd.Output()
	if err != nil {
		fmt.Println("Output from terminal : ", err)
	}
	return string(commandOutput)
}

func openDirectory(dirPath string) {
	cmd := exec.Command("osascript", "-s", "h", "-e", `tell application "Terminal" to do script "cd `+dirPath+`"`)
	cmd.Run()
}
