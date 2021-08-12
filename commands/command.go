package commands

type Command interface {
	CanHandle(message string) bool
	Process(message string) (string, error)
}
