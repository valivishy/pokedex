package commands

import (
	"fmt"
)

func CommandHelp(_ *Config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)

	return nil
}
