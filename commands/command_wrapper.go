package commands

type commandRunner interface {
	CanRun(message string) bool
	Run(message string) (string, error)
}

type CommandWrapper struct {
	Command commandRunner
}

func (w CommandWrapper) CanRun(message string) bool {
	return w.Command.CanRun(message)
}

func (w CommandWrapper) Run(message string) (string, error) {
	return w.Command.Run(message)
}
