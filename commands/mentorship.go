package commands

import "strings"

func CanHandle(message string) bool {
	return strings.TrimSpace(message) == "mentorias"
}
