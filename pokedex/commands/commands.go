package commands

import (
	"fmt"
	"os"
)

type Config struct {
	Next     *string
	Previous *string
}
type CLICommand struct {
	Name        string
	Description string
	Callback    func([]string) error
}

func CommandExit(_ *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func CommandHelp(_ *Config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)

	return nil
}
