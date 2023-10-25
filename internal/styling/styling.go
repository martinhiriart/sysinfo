package styling

import (
	"strings"
)

func FormatString(stringsToFormat []string) string {
	outputString := strings.Join(stringsToFormat, "\n")
	return outputString
}

func FormatMultiLineStrings(IPs []string) string {
	multiString := strings.Join(IPs, ";\n\t\t\t\t")
	return multiString
}
