package main

import (
	"fmt"
	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	"github.com/martinhiriart/sysinfo/internal/commands"
	"github.com/martinhiriart/sysinfo/internal/os_scripts"
	"github.com/martinhiriart/sysinfo/internal/styling"
	"github.com/ncruces/zenity"
	"net"
	"os"
	"os/user"
	"strings"
)

func getHostInfo() types.HostInfo {
	host, err := sysinfo.Host()
	if err != nil {
		styling.StyleErrors(err, "Panic")
	}
	return host.Info()
}

func getIPAddresses() []net.Addr {
	interfIPs, err := net.InterfaceAddrs()
	if err != nil {
		styling.StyleErrors(err, "log")
	}
	return interfIPs
}

func getCurrentUser() *user.User {
	userInfo, err := user.Current()
	if err != nil {
		styling.StyleErrors(err, "log")
	}
	return userInfo
}

func diplayQuickInfoDialog() {
	hostInfo := getHostInfo()
	ifIPs := getIPAddresses()
	userInfo := getCurrentUser()
	hostname, err := os.Hostname()
	if err != nil {
		styling.StyleErrors(err, "log")
	}
	usrInfo, err := user.LookupId(userInfo.Uid)
	if err != nil {
		styling.StyleErrors(err, "log")
	}

	var v4IPs []string
	for _, ipAddr := range ifIPs {
		if !strings.Contains(ipAddr.String(), "::") && !strings.Contains(ipAddr.String(), ":") && !strings.Contains(ipAddr.String(), "127.0.0.1") && !strings.Contains(ipAddr.String(), "169.254") {
			v4IPs = append(v4IPs, ipAddr.String())
		}
	}

	v4IPStr := styling.FormatIPStrings(v4IPs)

	var stringsToOutput []string

	hostnameStr := fmt.Sprintf("Hostname:\t\t %v", hostname)
	osStr := fmt.Sprintf("Operating System: %v %v", hostInfo.OS.Name, hostInfo.OS.Version)
	archStr := fmt.Sprintf("Architecture:\t\t %v", hostInfo.Architecture)
	tmzStr := fmt.Sprintf("Time Zone:\t\t %v", hostInfo.Timezone)
	ipStr := fmt.Sprintf("IP Addresses:\t\t %v", v4IPStr)
	usrStr := fmt.Sprintf("Current User:\t\t %v", usrInfo.Username)

	stringsToOutput = append(stringsToOutput, hostnameStr)
	stringsToOutput = append(stringsToOutput, osStr)
	stringsToOutput = append(stringsToOutput, archStr)
	stringsToOutput = append(stringsToOutput, tmzStr)
	stringsToOutput = append(stringsToOutput, ipStr)
	stringsToOutput = append(stringsToOutput, usrStr)

	outputString := styling.FormatString(stringsToOutput)

	err = zenity.Info(outputString,
		zenity.Title("System Information"),
		zenity.Icon("internal/assets/diagIcon.png"))

	if err != nil {
		styling.StyleErrors(err, "Panic")
	}
}

func main() {

	allAccounts := commands.GetAccounts()
	os_scripts.PrintAccountData(allAccounts, "Users")
}
