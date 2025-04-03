package repl

import (
	"bufio"
	"fmt"
	"github.com/valivishy/pokedex/commands"
	"github.com/valivishy/pokedex/util"
	"os"
)

func StartRepl() {
	startURL := "https://pokeapi.co/api/v2/location-area/"
	cfg := &commands.Config{
		Next: &startURL,
	}

	availableCommands := map[string]commands.CLICommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    func(params []string) error { return commands.CommandExit(cfg) },
		},
		"help": {
			Name:        "help",
			Description: "Help",
			Callback:    func(params []string) error { return commands.CommandHelp(cfg) },
		},
		"map": {
			Name:        "map",
			Description: "Show the map",
			Callback:    func(params []string) error { return commands.CommandMap(cfg) },
		},
		"mapb": {
			Name:        "mapb",
			Description: "Show the map in binary format",
			Callback:    func(params []string) error { return commands.CommandMapBack(cfg) },
		},
		"explore": {
			Name:        "explore",
			Description: "Explore the location passed as argument",
			Callback:    func(params []string) error { return commands.CommandExplore(cfg, params) },
		},
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := util.CleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := availableCommands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.Callback(words[1:])
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
