package discord

import (
	"log"
	"strings"
)

type Runner interface {
	Run(string) (string, error)
}

type bot struct {
	mentorship Runner
}

const unknownCommandMessage = "Desculpa, não entendi o seu comando"
const commandErrorMessage = "Desculpa, houve um erro no com seu comando"

func Init(mentorship Runner) bot {
	return bot{
		mentorship: mentorship,
	}
}

func (b bot) Process(message string) string {
	if strings.EqualFold(message, "mentorias") {
		response, err := b.mentorship.Run(message)
		if err != nil {
			log.Fatalf("Error while responding to membership message: %e", err)
			return commandErrorMessage
		}
		return response
	}

	return unknownCommandMessage
}
