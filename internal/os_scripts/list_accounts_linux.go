//go:build linux
// +build linux

package os_scripts

import (
	"bytes"
	"fmt"
	"github.com/martinhiriart/sysinfo/internal/styling"
	"os/exec"
	"strconv"
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
		entry = strings.Replace(entry, ":", ";", -1)

		individualAccountStrings := strings.Split(entry, ";")

		if len(individualAccountStrings) > 1 {
			aInfo.Name = individualAccountStrings[0]
			aInfo.Uid = individualAccountStrings[2]
			aInfo.Gid = individualAccountStrings[3]
			aInfo.DisplayName = individualAccountStrings[4]
			aInfo.HomeDir = individualAccountStrings[5]
			aInfo.Shell = individualAccountStrings[6]
		}

		accountsList = append(accountsList, aInfo)
	}
	return accountsList
}

func PrintAccountData(allAccounts []AccountInfo, printType string) {
	printType = strings.ToLower(printType)
	switch printType {
	case "user", "users":
		for _, user := range allAccounts {
			if user.Uid != "" {
				uidInt, err := strconv.Atoi(user.Uid)
				if err != nil {
					styling.StyleErrors(err, "log")
				}
				if uidInt > 500 && !strings.Contains(user.HomeDir, "/nonexistent") {
					fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", user.Name, user.Uid, user.Gid, user.HomeDir, user.Shell, user.DisplayName)
				}
			}
		}
	case "all":
		for _, user := range allAccounts {
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", user.Name, user.Uid, user.Gid, user.HomeDir, user.Shell, user.DisplayName)
		}
	}

}
