package main

import (
	"fmt"
	"github.com/ncruces/zenity"
	"log"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sysinfo/internal/apps"
	"sysinfo/internal/networking"
	"sysinfo/internal/operatingSystem"
	"sysinfo/internal/styling"
	"sysinfo/internal/system"
	"sysinfo/internal/users"
)

func displaySystemInfo(outputString string) {
	if err := zenity.Info(outputString,
		zenity.Title("System Information"),
		zenity.Icon("internal/assets/dialogIcon.png")); err != nil {
		log.Fatal(err)
	}
}

func getSpecificAppInfo(apps apps.AppInfo) string {
	var appString string
	for _, app := range apps {
		switch runtime.GOOS {
		case "darwin":
			if app.Name == "Cider.app" {
				appString = fmt.Sprintf("Cider version: \t%s\n", app.Version)
				break
			}
		case "windows":
			if app.Name == "Cider" {
				appString = fmt.Sprintf("Cider version: \t%s\n", app.Version)
				break
			}
		}
	}
	return appString
}

func getIPv4Addresses(netInfo networking.NetInfo) []string {
	var ipv4Addresses []string
	for _, ip := range netInfo {
		if !strings.Contains(ip.Address, ":") {
			ipv4Addresses = append(ipv4Addresses, ip.Address)
		}
	}
	return ipv4Addresses
}

func getUsernames(info users.UserInfo) []string {
	var userNameStrings []string
	for _, user := range info {
		if user.Username != "root" {
			userNameStrings = append(userNameStrings, user.Username)
		}
	}
	return userNameStrings
}

func getOSVersionInfo(osInfo operatingSystem.OSInfo) string {
	var osString string
	for _, opSys := range osInfo {
		osString = fmt.Sprintf("OS: \t\t\t\t%s %s", opSys.Name, opSys.Version)
	}
	return osString
}

func getOSArchInfo(osInfo operatingSystem.OSInfo) string {
	var archString string
	for _, opSys := range osInfo {
		archString = fmt.Sprintf("Arch: \t\t\t%s", opSys.Arch)
	}
	return archString
}

func getTZInfo(tzInfo system.TimeZone) string {
	var tzString string
	for _, tz := range tzInfo {
		tzString = fmt.Sprintf("Time Zone: \t\t%v", tz.LocalTimezone)
	}
	return tzString
}

func getHardwareInfo(sysInfo system.SysInfo) (string, string, string, string, string, string, string) {
	var (
		modelString          string
		makeString           string
		hardwareSerialString string
		boardSerialString    string
		memoryString         string
		compNameString       string
		cpuBrandString       string
	)
	for _, sys := range sysInfo {
		modelString = sys.HardwareModel
		makeString = sys.HardwareVendor
		hardwareSerialString = sys.HardwareSerial
		boardSerialString = sys.BoardSerial
		memoryString = sys.PhysicalMemory
		compNameString = sys.Hostname
		cpuBrandString = sys.CPUBrand
	}

	return modelString, makeString, hardwareSerialString, boardSerialString, memoryString, compNameString, cpuBrandString
}

func calculateMemory(mem string) float64 {
	res := math.Pow(1024, 3)
	memInt, err := strconv.Atoi(mem)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return float64(memInt) / res
}

func main() {
	// Collecting the installed applications on the system via osquery
	installedApps := apps.GetInstalledApps()
	appString := getSpecificAppInfo(installedApps)

	// Collecting the network interface information on the system via osquery
	networkInterfaceInfo := networking.GetNetworkInterfaceInfo()
	IPv4Info := getIPv4Addresses(networkInterfaceInfo)
	netInfoString := fmt.Sprintf("IPv4 Addresses: \t%v", styling.FormatMultiLineStrings(IPv4Info))

	usersInfo := users.GetUserInfo()
	usernames := getUsernames(usersInfo)
	usernameString := fmt.Sprintf("Users on system: \t%v", styling.FormatMultiLineStrings(usernames))

	osInfo := operatingSystem.GetOSInfo()
	osInfoString := getOSVersionInfo(osInfo)
	osArchString := getOSArchInfo(osInfo)

	sysInfo := system.GetSystemInfo()
	model, oem, hwSerial, biosSerial, memory, compName, cpuBrand := getHardwareInfo(sysInfo)

	memoryGB := calculateMemory(memory)

	modelString := fmt.Sprintf("Model: \t\t\t%v", model)
	makeString := fmt.Sprintf("Manufacturer: \t%v", oem)
	var serialString string
	if hwSerial != "" {
		serialString = fmt.Sprintf("Serial Number: \t%v", hwSerial)
	} else {
		serialString = fmt.Sprintf("Serial Number: \t%v", biosSerial)
	}
	memString := fmt.Sprintf("Memory: \t\t%v GiB", memoryGB)
	hostnameString := fmt.Sprintf("Hostname: \t\t%v", compName)

	tzInfo := system.GetTimeZone()
	tzString := getTZInfo(tzInfo)

	cpuString := fmt.Sprintf("CPU: \t\t\t%v", cpuBrand)

	// Aggregating the strings needed to create the zenity info dialog
	var systemInfoStrings []string
	systemInfoStrings = append(systemInfoStrings, hostnameString)
	systemInfoStrings = append(systemInfoStrings, makeString)
	systemInfoStrings = append(systemInfoStrings, modelString)
	systemInfoStrings = append(systemInfoStrings, serialString)
	systemInfoStrings = append(systemInfoStrings, "")
	systemInfoStrings = append(systemInfoStrings, cpuString)
	systemInfoStrings = append(systemInfoStrings, memString)
	systemInfoStrings = append(systemInfoStrings, osInfoString)
	systemInfoStrings = append(systemInfoStrings, osArchString)
	systemInfoStrings = append(systemInfoStrings, "")
	systemInfoStrings = append(systemInfoStrings, usernameString)
	systemInfoStrings = append(systemInfoStrings, tzString)
	systemInfoStrings = append(systemInfoStrings, netInfoString)
	systemInfoStrings = append(systemInfoStrings, "")
	systemInfoStrings = append(systemInfoStrings, appString)

	// Formatting the aggregated strings
	stringToOutput := styling.FormatString(systemInfoStrings)
	// Displaying the strings via a zenity info dialog box
	displaySystemInfo(stringToOutput)
}
