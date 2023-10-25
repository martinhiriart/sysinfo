package apps

type AppInfo []struct {
	Version string `json:"version"`
	Name    string `json:"name"`
	Path    string `json:"install_location"`
}

func GetInstalledApps() AppInfo {
	query, err := exec.Command("bash", "-c", "echo \"SELECT name, version, install_location FROM programs;\" | osqueryi --json").Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	var appInfo AppInfo
	if err := json.Unmarshal(query, &appInfo); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return appInfo
}
