package commands

import (
	"errors"
	"fmt"
	"github.com/valivishy/pokedex/internal/api/pokedex"
)

func CommandCatch(_ *Config, params []string) error {
	return catchPokemon(params)
}

// region PRIVATE
func catchPokemon(params []string) error {
	if len(params) == 0 {
		return errors.New("no pokemon provided")
	}

	if len(params) > 1 {
		return errors.New("only one pokemon can be caught at a time")
	}

	pokemon := params[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	caught, err := pokedex.Catch(pokemon)
	if err != nil {
		return err
	}

	if caught {
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}

//endregion PRIVATE
