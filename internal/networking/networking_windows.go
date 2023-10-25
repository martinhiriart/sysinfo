package networking

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type NetInfo []struct {
	Address      string `json:"address"`
	FriendlyName string `json:"friendly_name"`
	Interface    string `json:"interface"`
	SubnetMask   string `json:"mask"`
	PointToPoint string `json:"point_to_point"`
	NetworkType  string `json:"type"`
}

func GetNetworkInterfaceInfo() NetInfo {
	queryText := "echo \"SELECT interface, address, mask,point_to_point, type, friendly_name FROM interface_addresses WHERE interface NOT LIKE 'lo0';\" | osqueryi --json"
	query, err := exec.Command("bash", "-c", queryText).Output()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	var netInfo NetInfo
	if err := json.Unmarshal(query, &netInfo); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return netInfo
}
