//go:build linux
// +build linux

package os_scripts

import (
	"os/exec"
	"strings"
)

type AccountInfo struct {
	Name        string
	Password    string
	Uid         string
	Gid         string
	HomeDir     string
	Shell       string
	DisplayName string
}

func ListAllAccounts() []AccountInfo {
	accountsCmd := exec.Command("cat", "/etc/passwd")
	var stdout, stderr bytes.Buffer
	accountsCmd.Stdout = &stdout
	accountsCmd.Stderr = &stderr
	err := accountsCmd.Run()
	if err != nil {
		styling.StyleErrors(err, "Log")
	}
	accountsStr, _ := string(stdout.Bytes()), string(stderr.Bytes())

	splitAccountsString := strings.Split(accountsStr, "\n")

	accountsList := []AccountInfo{}
	for _, entry := range splitAccountsString {
		var aInfo AccountInfo
		individualAccountStrings := strings.Split(entry, ":")

		for _, val := range individualAccountStrings {
			valString := strings.Split(val, ": ")

			switch valString[0] {
			case "name":
				aInfo.Name = valString[1]
			case "password":
				aInfo.Password = valString[1]
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

		accountsList = append(accountsList, aInfo)
	}

	return accountsList
}
