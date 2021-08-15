package discord

import (
	"fmt"
	"strings"
)

type Runner interface {
	Run(string) (string, error)
}

type bot struct {
	memtorship Runner
}

const unknownCommandMessage = "Desculpa, não entendi o seu comando"
const commandErrorMessage = "Desculpa, houve um erro no com seu comando"

func Init(memtorship Runner) bot {
	return bot{
		memtorship: memtorship,
	}
}

func (b bot) Process(message string) string {
	if strings.EqualFold(message, "mentorias") {
		response, err := b.memtorship.Run(message)
		if err != nil {
			fmt.Printf("Error while responding to membership message: %e", err)
			return commandErrorMessage
		}
		return response
	}

	return unknownCommandMessage
}
