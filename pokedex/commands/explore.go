package commands

import (
	"errors"
	"fmt"
	"github.com/valivishy/pokedex/internal/api/locations"
)

func CommandExplore(_ *Config, params []string) error {
	return exploreLocation(params)
}

// region PRIVATE
func exploreLocation(params []string) error {
	if len(params) == 0 {
		return errors.New("no location provided")
	}

	if len(params) > 1 {
		return errors.New("too many location provided")
	}

	area := params[0]
	response, err := locations.Get(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")

	for _, encounter := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

//endregion PRIVATE
