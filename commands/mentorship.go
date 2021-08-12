package commands

import (
	"bytes"
	"fmt"
	"strings"

	mentors "github.com/CaravanaCloud/bino/mentors"
)

func CanHandle(message string) bool {
	return strings.TrimSpace(message) == "mentorias"
}

func Process(message string) (string, error) {
	if !CanHandle(message) {
		return "", fmt.Errorf("Can't process message '%s'", message)
	}

	return asTextList(mentors.List()), nil
}

func asTextList(mentors []mentors.Mentor) string {
	var stringBuffer bytes.Buffer
	for _, mentor := range mentors {
		stringBuffer.WriteString("- " + mentor.Name + "\n")
	}
	return stringBuffer.String()
}
