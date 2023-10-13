//go:build darwin
// +build darwin

package os_scripts

import (
	"bytes"
	"fmt"
	"github.com/martinhiriart/sysinfo/internal/styling"
	"os/exec"
	"strings"
)

type AccountInfo struct {
	Name        string
	Uid         string
	Gid         string
	HomeDir     string
	Shell       string
	DisplayName string
}

func ListAllAccounts() []AccountInfo {
	usersCmd := exec.Command("dscacheutil", "-q", "user")
	var stdout, stderr bytes.Buffer
	usersCmd.Stdout = &stdout
	usersCmd.Stderr = &stderr
	err := usersCmd.Run()
	if err != nil {
		styling.StyleErrors(err, "Log")
	}
	usersStr, _ := string(stdout.Bytes()), string(stderr.Bytes())

	newString := strings.Split(usersStr, "\n\n")

	userList := []AccountInfo{}
	for _, entry := range newString {
		var aInfo AccountInfo
		newString2 := strings.Split(entry, "\n")

		for _, val := range newString2 {
			valString := strings.Split(val, ": ")

			switch valString[0] {
			case "name":
				aInfo.Name = valString[1]
			case "uid":
				aInfo.Uid = valString[1]
			case "gid":
				aInfo.Gid = valString[1]
			case "dir":
				aInfo.HomeDir = valString[1]
			case "shell":
				aInfo.Shell = valString[1]
			case "gecos":
				aInfo.DisplayName = valString[1]
			}

		}

		userList = append(userList, aInfo)
	}

	return userList
}

func PrintAccountData(allAccounts []AccountInfo, printType string) {
	printType = strings.ToLower(printType)
	switch printType {
	case "user", "users":
		for _, user := range allAccounts {
			if strings.Contains(user.HomeDir, "/Users") {
				fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", user.Name, user.Uid, user.Gid, user.HomeDir, user.Shell, user.DisplayName)
			}
		}
	case "all":
		for _, user := range allAccounts {
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", user.Name, user.Uid, user.Gid, user.HomeDir, user.Shell, user.DisplayName)
		}
	}

}
