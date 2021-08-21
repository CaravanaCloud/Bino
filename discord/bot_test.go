package discord

import (
	"errors"
	"testing"
)

type testRunner struct {
	run func(string) (string, error)
}

func (f testRunner) Run(m string) (string, error) {
	return f.run(m)
}

func TestProcessMentorship(t *testing.T) {
	message := "mentorias"
	expectedResponse := "Mentorship executed"
	fakeMentorship := testRunner{
		run: func(m string) (string, error) {
			return expectedResponse, nil
		},
	}

	bot := Init(fakeMentorship)

	response := bot.Process(message)

	if response != expectedResponse {
		t.Fatalf("Expected process to responde with '%s', but got: %s", expectedResponse, response)
	}
}

func TestProcessMentorshipIgnoringCase(t *testing.T) {
	message := "MentoRias"
	expectedResponse := "Mentorship executed"
	fakeMentorship := testRunner{
		run: func(m string) (string, error) {
			return expectedResponse, nil
		},
	}

	bot := Init(fakeMentorship)

	response := bot.Process(message)

	if response != expectedResponse {
		t.Fatalf("Expected process to responde with '%s', but got: %s", expectedResponse, response)
	}
}

func TestProcessRespondsWithUserFriendlyMessageIfCaseOfError(t *testing.T) {
	message := "mentorias"
	expectedErrorMessage := "Desculpa, houve um erro no com seu comando"
	fakeMentorship := testRunner{
		run: func(m string) (string, error) {
			return "", errors.New("boom!")
		},
	}

	bot := Init(fakeMentorship)

	response := bot.Process(message)

	if response != expectedErrorMessage {
		t.Fatalf("Expected process to responde with '%s', got: %s", expectedErrorMessage, response)
	}
}

func TestProcessUnknownWithMessage(t *testing.T) {
	message := "somethingElse"
	expectedResponse := "Desculpa, não entendi o seu comando"
	fakeMentorship := testRunner{
		run: func(m string) (string, error) {
			return "", errors.New("Should not have been executed")
		},
	}

	bot := Init(fakeMentorship)

	response := bot.Process(message)

	if response != expectedResponse {
		t.Fatalf("Expected process to responde with '%s' , got: %s", expectedResponse, response)
	}
}
