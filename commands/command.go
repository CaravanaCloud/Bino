package commands

type Command interface {
	CanProcess(message string) bool
	Process(message string) (string, error)
}
