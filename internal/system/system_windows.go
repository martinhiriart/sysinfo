package system

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type SysInfo []struct {
	BoardSerial    string `json:"board_serial"`
	ComputerName   string `json:"computer_name"`
	CPUBrand       string `json:"cpu_brand"`
	HardwareModel  string `json:"hardware_model"`
	HardwareSerial string `json:"hardware_serial"`
	HardwareVendor string `json:"hardware_vendor"`
	Hostname       string `json:"hostname"`
	PhysicalMemory string `json:"physical_memory"`
}

func GetSystemInfo() SysInfo {
	queryText := "echo SELECT \"hostname, cpu_brand, physical_memory, hardware_vendor, hardware_model, hardware_serial, board_serial, computer_name FROM system_info;\" | osqueryi --json"
	query, err := exec.Command("bash", "-c", queryText).Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	var systemInfo SysInfo
	if err := json.Unmarshal(query, &systemInfo); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return systemInfo
}
