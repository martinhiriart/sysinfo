package commands

import (
	"fmt"
	"github.com/martinhiriart/sysinfo/internal/os_scripts"
	"strings"
)

func GetAccounts() []os_scripts.AccountInfo {
	return os_scripts.ListAllAccounts()

}

func PrintAccountData(allAccounts []os_scripts.AccountInfo, printType string) {
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
