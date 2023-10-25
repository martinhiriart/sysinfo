package apps

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type AppInfo []struct {
	Version string `json:"bundle_version"`
	Name    string `json:"name"`
	Path    string `json:"path"`
}

func GetInstalledApps() AppInfo {
	query, err := exec.Command("bash", "-c", "echo \"SELECT name, bundle_version, path FROM apps;\" | osqueryi --json").Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	var appInfo AppInfo
	if err := json.Unmarshal(query, &appInfo); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return appInfo
}
