package commands

import (
	"fmt"
	"github.com/valivishy/pokedex/internal/api/locations"
)

func CommandMap(cfg *Config) error {
	return printLocationsAndUpdateUrls(cfg, cfg.Next)
}

func CommandMapBack(cfg *Config) error {
	return printLocationsAndUpdateUrls(cfg, cfg.Previous)
}

// region PRIVATE
func printLocationsAndUpdateUrls(cfg *Config, url *string) error {
	response, err := locations.List(url)
	if err != nil {
		return err
	}

	cfg.Next = response.Next
	cfg.Previous = response.Previous

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}

//endregion PRIVATE
