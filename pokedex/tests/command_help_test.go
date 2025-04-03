package tests

import (
	"github.com/valivishy/pokedex/commands"
	"testing"
)

func TestCommandHelp(t *testing.T) {
	err := commands.CommandHelp(&commands.Config{})
	if err != nil {
		t.Errorf("commandHelp() returned an unexpected error: %v", err)
	}
}
