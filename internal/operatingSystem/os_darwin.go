package operatingSystem

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type OSInfo []struct {
	Arch    string `json:"arch"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func GetOSInfo() OSInfo {
	queryText := "echo \"SELECT name, version, arch FROM os_version;\" | osqueryi --json"
	query, err := exec.Command("bash", "-c", queryText).Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	var osInfo OSInfo
	if err := json.Unmarshal(query, &osInfo); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return osInfo
}
