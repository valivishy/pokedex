package tests

import (
	"encoding/json"
	"github.com/valivishy/pokedex/commands"
	"github.com/valivishy/pokedex/internal/api/pokedex"
	"net/http"
	"net/http/httptest"
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
