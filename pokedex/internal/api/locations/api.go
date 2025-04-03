package locations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valivishy/pokedex/internal/api"
)

var GetBaseUrl = "https://pokeapi.co/api/v2/location-area"

func Get(name string) (*LocationDetails, error) {
	body, err := api.Get(fmt.Sprintf("%s/%s", GetBaseUrl, name))
	if err != nil {
		return nil, err
	}

	response := LocationDetails{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func List(url *string) (*LocationList, error) {
	if url == nil {
		return nil, errors.New("no more locations in this direction")
	}

	body, err := api.Get(*url)
	if err != nil {
		return nil, err
	}

	response := LocationList{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
