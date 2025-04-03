package tests

import (
	"bytes"
	"encoding/json"
	"github.com/valivishy/pokedex/commands"
	"github.com/valivishy/pokedex/internal/api/pokedex"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCommandCatch_ThenList(t *testing.T) {
	catchPokemon(t)

	output := captureOutput(func() {
		_ = commands.CommandPokedex(nil, []string{})
	})

	if !strings.Contains(output, "Your Pokedex:") || !strings.Contains(output, "- magikarp") {
		t.Fatalf("expected inspect output, got:\n%s", output)
	}
}

func catchPokemon(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := pokedex.Pokemon{
			Name:           "magikarp",
			Height:         9,
			Weight:         100,
			BaseExperience: 1,
			Stats: []pokedex.Stats{
				{BaseStat: 20, Stat: pokedex.Stat{Name: "hp"}},
			},
			Types: []pokedex.Types{
				{Type: pokedex.Type{Name: "water"}},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(payload)
	}))
	defer ts.Close()

	pokedex.GetBaseUrl = ts.URL

	err := commands.CommandCatch(nil, []string{"magikarp"})
	if err != nil {
		t.Fatalf("expected to catch pokemon, got error: %v", err)
	}
}

func TestCommandCatch_ThenInspect(t *testing.T) {
	catchPokemon(t)

	output := captureOutput(func() {
		_ = commands.CommandInspect(nil, []string{"magikarp"})
	})

	if !strings.Contains(output, "Name: magikarp") || !strings.Contains(output, "- water") {
		t.Fatalf("expected inspect output, got:\n%s", output)
	}
}

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)
		outC <- buf.String()
	}()

	f()
	_ = w.Close()
	os.Stdout = stdout
	return <-outC
}
