package commands

import (
	"fmt"
	"strings"
)

func CanHandle(message string) bool {
	return strings.TrimSpace(message) == "mentorias"
}

func Process(message string) (string, error) {
	if !CanHandle(message) {
		return "", fmt.Errorf("Can't process message '%s'", message)
	}

	return `
		- Lucia
		- Julio
		- Marcus
	`, nil
}
