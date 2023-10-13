package commands

import (
	"github.com/martinhiriart/sysinfo/internal/os_scripts"
)

func GetAccounts() []os_scripts.AccountInfo {
	return os_scripts.ListAllAccounts()

}
