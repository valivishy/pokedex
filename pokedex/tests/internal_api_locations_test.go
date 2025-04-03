package tests

import (
	"encoding/json"
	"fmt"
	"github.com/valivishy/pokedex/commands"
	"github.com/valivishy/pokedex/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNextLocations(t *testing.T) {
	mockResponse := api.Response{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results: []api.Location{
			{Name: "test-location", URL: "https://pokeapi.co/test"},
		},
	}

	server := prepareServer(mockResponse)
	defer server.Close()

	cfg := &commands.Config{Next: &server.URL}

	locationResponse, err := api.List(cfg.Next)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	locationsList := locationResponse.Results

	if len(locationsList) != 1 || locationsList[0].Name != "test-location" {
		t.Fatalf("Unexpected result: %v", locationsList)
	}
}

func prepareServer(mockResponse api.Response) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(mockResponse)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
	}))
	return server
}

func TestGetPreviousLocations(t *testing.T) {
	mockResponse := api.Response{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results: []api.Location{
			{Name: "previous-location", URL: "https://pokeapi.co/previous"},
		},
	}

	server := prepareServer(mockResponse)
	defer server.Close()

	cfg := &commands.Config{Previous: &server.URL}

	locationResponse, err := api.List(cfg.Previous)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	locationsList := locationResponse.Results

	if len(locationsList) != 1 || locationsList[0].Name != "previous-location" {
		t.Fatalf("Unexpected result: %v", locationsList)
	}
}

func TestCommandMap(t *testing.T) {
	mockResponse := api.Response{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results: []api.Location{
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
