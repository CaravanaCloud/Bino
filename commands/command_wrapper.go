package commands

type command struct {
	CanRun func(message string) bool
	Run    func(message string) (string, error)
}

type CommandWrapper struct {
	Command command
}

func (w CommandWrapper) CanRun(message string) bool {
	return w.Command.CanRun(message)
}

func (w CommandWrapper) Run(message string) (string, error) {
	return w.Command.Run(message)
}
