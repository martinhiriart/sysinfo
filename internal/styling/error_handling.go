package styling

import (
	"log"
	"strings"
)

func StyleErrors(err error, errType string) {
	errType = strings.ToLower(errType)
	switch errType {
	case "panic":
		panic(err)
	default:
		log.Fatalf("ERROR: %v\n", err)
	}
}
