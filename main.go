package main

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	"github.com/martinhiriart/sysinfo/internal/styling"
	"github.com/ncruces/zenity"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

type UserInfo struct {
	name     string
	password string
	uid      string
	gid      string
	dir      string
	shell    string
	gecos    string
}

func getHostInfo() types.HostInfo {
	host, err := sysinfo.Host()
	if err != nil {
		handleError(err, "Panic")
	}
	return host.Info()
}

func getIPAddresses() []net.Addr {
	interfIPs, err := net.InterfaceAddrs()
	if err != nil {
		handleError(err, "log")
	}
	return interfIPs
}

func getCurrentUser() *user.User {
	userInfo, err := user.Current()
	if err != nil {
		handleError(err, "log")
	}
	return userInfo
}

func handleError(err error, errType string) {
	errType = strings.ToLower(errType)
	switch errType {
	case "panic":
		panic(err)
	default:
		log.Fatalf("ERROR: %v\n", err)
	}
}

func listAllUsers() []UserInfo {
	usersCmd := exec.Command("dscacheutil", "-q", "user")
	var stdout, stderr bytes.Buffer
	usersCmd.Stdout = &stdout
	usersCmd.Stderr = &stderr
	err := usersCmd.Run()
	if err != nil {
		handleError(err, "Log")
	}
	usersStr, _ := string(stdout.Bytes()), string(stderr.Bytes())

	newString := strings.Split(usersStr, "\n\n")

	userList := []UserInfo{}
	for _, entry := range newString {
		var uInfo UserInfo
		newString2 := strings.Split(entry, "\n")

		for _, val := range newString2 {
			valString := strings.Split(val, ": ")

			switch valString[0] {
			case "name":
				uInfo.name = valString[1]
			case "password":
				uInfo.password = valString[1]
			case "uid":
				uInfo.uid = valString[1]
			case "gid":
				uInfo.gid = valString[1]
			case "dir":
				uInfo.dir = valString[1]
			case "shell":
				uInfo.shell = valString[1]
			case "gecos":
				uInfo.gecos = valString[1]
			}

		}

		if strings.Contains(uInfo.dir, "/Users") {
			userList = append(userList, uInfo)
		}
	}

	return userList
}

func main() {

	hostInfo := getHostInfo()
	ifIPs := getIPAddresses()
	userInfo := getCurrentUser()
	hostname, err := os.Hostname()
	if err != nil {
		handleError(err, "log")
	}
	usrInfo, err := user.LookupId(userInfo.Uid)
	if err != nil {
		handleError(err, "log")
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
		handleError(err, "Panic")
	}

	//listAllUsers()

}
