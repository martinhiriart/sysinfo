package styling

import (
	"strings"
)

func FormatString(stringsToFormat []string) string {
	outputString := strings.Join(stringsToFormat, "\n")
	return outputString
}

func FormatIPStrings(IPs []string) string {
	IPstring := strings.Join(IPs, ";\n\t\t\t\t ")
	return IPstring
}
