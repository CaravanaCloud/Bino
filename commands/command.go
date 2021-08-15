package commands

type CommandRunner interface {
	CanRun(message string) bool
	Run(message string) (string, error)
}
