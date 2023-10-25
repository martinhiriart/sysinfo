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
	queryText := "echo \"SELECT name, version, install_location FROM programs;\" | . \"C:\\Program Files\\osquery\\osqueryi.exe\" --json"
	query, err := exec.Command("cmd", "/c", queryText).Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	var appInfo AppInfo
	if err := json.Unmarshal(query, &appInfo); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return appInfo
}
