package commands

import (
	"fmt"
	"os"
)

func CommandExit(_ *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
