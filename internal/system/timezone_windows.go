package system

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type TimeZone []struct {
	LocalTimezone string `json:"local_timezone"`
}

func GetTimeZone() TimeZone {
	queryText := "echo SELECT \"local_timezone FROM time;\" | osqueryi --json"
	query, err := exec.Command("bash", "-c", queryText).Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	var tzInfo TimeZone
	if err := json.Unmarshal(query, &tzInfo); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return tzInfo
}
