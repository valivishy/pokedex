package commands

import (
	"github.com/valivishy/pokedex/internal/api/pokedex"
)

func CommandPokedex(_ *Config, _ []string) error {
	return listCaughtPokemon()
}

// region PRIVATE
func listCaughtPokemon() error {
	pokedex.List()

	return nil
}

//endregion PRIVATE
