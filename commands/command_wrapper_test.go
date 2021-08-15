package commands

import (
	"testing"
)

func TestRunFailsIfCommandCantRun(t *testing.T) {
	message := "any message"

	wrapper := CommandWrapper{
		Command: command{
			CanRun: func(message string) bool {
				return false
			},
			Run: func(message string) (string, error) {
				return "", nil
			},
		},
	}

	_, err := wrapper.Run(message)

	if err == nil {
		t.Fatal("Expected run to error but it didn't")
	}
}
