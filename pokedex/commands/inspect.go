package commands

import (
	"errors"
	"github.com/valivishy/pokedex/internal/api/pokedex"
)

func CommandInspect(_ *Config, params []string) error {
	return inspectPokemon(params)
}

// region PRIVATE
func inspectPokemon(params []string) error {
	if len(params) == 0 {
		return errors.New("no pokemon provided")
	}

	if len(params) > 1 {
		return errors.New("only one pokemon can be inspected at a time")
	}

	pokedex.Inspect(params[0])

	return nil
}

//endregion PRIVATE
