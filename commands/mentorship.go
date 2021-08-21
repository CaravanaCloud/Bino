package commands

import (
	"fmt"

	mentors "github.com/CaravanaCloud/bino/mentors"
)

var mentorshipCommand = command{
	CanRun: canRun,
	Run:    run,
}

var Mentorship = CommandRunner{
	Command: mentorshipCommand,
}

func canRun(message string) bool {
	return message == "mentorias"
}

func run(message string) (mentorsList string, err error) {
	for _, mentor := range mentors.List() {
		mentorsList = fmt.Sprintf("%s- %s\n", mentorsList, mentor.Name)
	}
	return
}
