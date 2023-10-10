package main

import (
	"fmt"
	"github.com/elastic/go-sysinfo"
	"github.com/ncruces/zenity"
	"net"
	"os"
	"os/user"
	"strings"
)

func main() {
	host, err := sysinfo.Host()
	if err != nil {
		panic(err)
	}

	ifIPs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	hostInfo := host.Info()
	userInfo, err := user.Current()
	if err != nil {
		panic(err)
	}
	var v4IPs []string
	for _, ipAddr := range ifIPs {
		if !strings.Contains(ipAddr.String(), "::") && !strings.Contains(ipAddr.String(), ":") && !strings.Contains(ipAddr.String(), "127.0.0.1") && !strings.Contains(ipAddr.String(), "169.254") {
			v4IPs = append(v4IPs, ipAddr.String())
		}
	}

	v4IPStr := strings.Join(v4IPs, ";\n\t\t\t\t ")

	hostname, err := os.Hostname()
	usrInfo, err := user.LookupId(userInfo.Uid)
	if err != nil {
		panic(err)
	}

	hostnameStr := fmt.Sprintf("Hostname:\t\t %v\n", hostname)
	osStr := fmt.Sprintf("Operating System: %v %v\n", hostInfo.OS.Name, hostInfo.OS.Version)
	archStr := fmt.Sprintf("Architecture:\t\t %v\n", hostInfo.Architecture)
	tmzStr := fmt.Sprintf("Time Zone:\t\t %v\n", hostInfo.Timezone)
	ipStr := fmt.Sprintf("IP Addresses:\t\t %v\n", v4IPStr)
	usrStr := fmt.Sprintf("Current User:\t\t %v\n", usrInfo.Username)

	err = zenity.Info(hostnameStr+usrStr+osStr+archStr+tmzStr+ipStr,
		zenity.Title("System Information"),
		zenity.Icon("internal/diagIcon.png"))

	if err != nil {
		panic(err)
	}
}
