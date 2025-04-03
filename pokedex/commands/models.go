package commands

type Config struct {
	Next     *string
	Previous *string
}
type CLICommand struct {
	Name        string
	Description string
	Callback    func([]string) error
}
