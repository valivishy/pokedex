package pokedex

import (
	"encoding/json"
	"fmt"
	"github.com/valivishy/pokedex/internal/api"
	"math/rand"
)

var GetBaseUrl = "https://pokeapi.co/api/v2/pokemon"
var pokedex = make(map[string]Pokemon)

const maxExperienceForInstantCatch = 10

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

	var caught bool
	if pokemon.BaseExperience <= maxExperienceForInstantCatch {
		caught = true
	} else {
		caught = rand.Intn(pokemon.BaseExperience) > (pokemon.BaseExperience / 2)
	}

	if caught {
		pokedex[name] = pokemon
	}

	return caught, nil
}

func Inspect(name string) {
	pokemon, ok := pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
}

func List() {
	if len(pokedex) == 0 {
		fmt.Println("no Pokemon caught")
		return
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
}
