package users

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type UserInfo []struct {
	Description string `json:"description"`
	Directory   string `json:"directory"`
	Shell       string `json:"shell"`
	Type        string `json:"type"`
	Username    string `json:"username"`
	UUID        string `json:"uuid"`
}

func GetUserInfo() UserInfo {
	queryText := "echo \"SELECT username, description, directory, shell, uuid, type FROM users WHERE shell NOT LIKE '%false' AND username NOT LIKE '/_%' ESCAPE '/';\" | . \"C:\\Program Files\\osquery\\osqueryi.exe\" --json"
	query, err := exec.Command("bash", "-c", queryText).Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	var userInfo UserInfo
	if err := json.Unmarshal(query, &userInfo); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return userInfo
}
