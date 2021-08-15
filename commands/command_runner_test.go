package commands

import (
	"testing"
)

func TestRunFailsIfCommandCantRun(t *testing.T) {
	message := "any message"

	runner := CommandRunner{
		Command: command{
			CanRun: func(message string) bool {
				return false
			},
			Run: func(message string) (string, error) {
				return "", nil
			},
		},
	}

	_, err := runner.Run(message)

	if err == nil {
		t.Fatal("Expected run to error but it didn't")
	}
}

func TestRunIgnoresAnyTraillingWhiteSpacesOnMessage(t *testing.T) {
	message := "testMessage        "

	runner := CommandRunner{
		Command: command{
			CanRun: func(message string) bool {
				return message == "testMessage"
			},
			Run: func(message string) (string, error) {
				return "", nil
			},
		},
	}

	_, err := runner.Run(message)

	if err != nil {
		t.Fatalf("Expected message '%s' to be acccepted", message)
	}
}
