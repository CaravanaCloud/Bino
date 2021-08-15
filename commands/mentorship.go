package commands

import (
	"bytes"
	"strings"

	mentors "github.com/CaravanaCloud/bino/mentors"
)

var Mentorship = CommandWrapper{
	Command: command{
		CanRun: canRun,
		Run:    run,
	},
}

func canRun(message string) bool {
	return strings.TrimSpace(message) == "mentorias"
}

func run(message string) (string, error) {
	return asTextList(mentors.List()), nil
}

func asTextList(mentors []mentors.Mentor) string {
	var stringBuffer bytes.Buffer
	for _, mentor := range mentors {
		stringBuffer.WriteString("- " + mentor.Name + "\n")
	}
	return stringBuffer.String()
}
