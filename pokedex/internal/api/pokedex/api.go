package pokedex

import (
	"encoding/json"
	"fmt"
	"github.com/valivishy/pokedex/internal/api"
	"math/rand"
)

var GetBaseUrl = "https://pokeapi.co/api/v2/pokemon"

var pokedex = make(map[string]Pokemon)

func Catch(name string) (bool, error) {
	body, err := api.Get(fmt.Sprintf("%s/%s", GetBaseUrl, name))
	if err != nil {
		return false, err
	}

	pokemon := Pokemon{}

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return false, err
	}

	caught := rand.Intn(pokemon.BaseExperience) > (pokemon.BaseExperience / 2)
	if caught {
		pokedex[name] = pokemon
	}

	return caught, nil
}
