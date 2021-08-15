package commands

import (
	"bytes"
	"fmt"
	"strings"

	mentors "github.com/CaravanaCloud/bino/mentors"
)

type Mentorship struct{}

func (m Mentorship) CanRun(message string) bool {
	return strings.TrimSpace(message) == "mentorias"
}

func (m Mentorship) Run(message string) (string, error) {
	if !m.CanRun(message) {
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
