package tests

import (
	"encoding/json"
	"fmt"
	"github.com/valivishy/pokedex/commands"
	"github.com/valivishy/pokedex/internal/api/locations"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNextLocations(t *testing.T) {
	mockResponse := locations.LocationList{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results: []locations.Location{
			{Name: "test-location", URL: "https://pokeapi.co/test"},
		},
	}

	server := prepareServer(mockResponse)
	defer server.Close()

	cfg := &commands.Config{Next: &server.URL}

	locationResponse, err := locations.List(cfg.Next)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	locationsList := locationResponse.Results

	if len(locationsList) != 1 || locationsList[0].Name != "test-location" {
		t.Fatalf("Unexpected result: %v", locationsList)
	}
}

func TestGetPreviousLocations(t *testing.T) {
	mockResponse := locations.LocationList{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results: []locations.Location{
			{Name: "previous-location", URL: "https://pokeapi.co/previous"},
		},
	}

	server := prepareServer(mockResponse)
	defer server.Close()

	cfg := &commands.Config{Previous: &server.URL}

	locationResponse, err := locations.List(cfg.Previous)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	locationsList := locationResponse.Results

	if len(locationsList) != 1 || locationsList[0].Name != "previous-location" {
		t.Fatalf("Unexpected result: %v", locationsList)
	}
}

func TestCommandMap(t *testing.T) {
	mockResponse := locations.LocationList{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results: []locations.Location{
			{Name: "command-location", URL: "https://pokeapi.co/command"},
		},
	}

	server := prepareServer(mockResponse)
	defer server.Close()

	cfg := &commands.Config{Next: &server.URL}

	err := commands.CommandMap(cfg)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
}

func prepareServer(mockResponse locations.LocationList) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(mockResponse)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
	}))
	return server
}
