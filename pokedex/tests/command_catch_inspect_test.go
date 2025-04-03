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

func TestCommandCatch_NoParams(t *testing.T) {
	err := commands.CommandCatch(nil, []string{})
	if err == nil || err.Error() != "no pokemon provided" {
		t.Fatalf("expected error for missing param, got: %v", err)
	}
}

func TestCommandCatch_TooManyParams(t *testing.T) {
	err := commands.CommandCatch(nil, []string{"pikachu", "bulbasaur"})
	if err == nil || err.Error() != "only one pokemon can be caught at a time" {
		t.Fatalf("expected error for too many params, got: %v", err)
	}
}

func TestCommandCatch_EasyToCatch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := pokedex.Pokemon{BaseExperience: 1}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(payload)
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	pokedex.GetBaseUrl = ts.URL
	err := commands.CommandCatch(nil, []string{"magikarp"})
	if err != nil {
		t.Fatalf("expected success catching easy pokemon, got: %v", err)
	}
}

func TestCommandCatch_HardToCatch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := pokedex.Pokemon{BaseExperience: 1000}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(payload)
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	pokedex.GetBaseUrl = ts.URL
	err := commands.CommandCatch(nil, []string{"mewtwo"})
	if err != nil {
		t.Fatalf("expected no error even if pokemon escapes, got: %v", err)
	}
}

func TestCommandCatch_ThenInspect(t *testing.T) {
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

	// Catch the Pokémon
	err := commands.CommandCatch(nil, []string{"magikarp"})
	if err != nil {
		t.Fatalf("expected to catch pokemon, got error: %v", err)
	}

	// Capture output of Inspect
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
