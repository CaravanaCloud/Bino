package commands

import (
	"fmt"
	"strings"
)

type command struct {
	CanRun func(message string) bool
	Run    func(message string) (string, error)
}

type CommandWrapper struct {
	Command command
}

func (w CommandWrapper) Run(message string) (string, error) {
	if !w.Command.CanRun(strings.TrimSpace(message)) {
		return "Execution failed", fmt.Errorf("Command can not be executed for message: %s", message)
	}
	return w.Command.Run(message)
}
