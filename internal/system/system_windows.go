package system

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
