package tests

import (
	"encoding/json"
	"github.com/valivishy/pokedex/commands"
	"github.com/valivishy/pokedex/internal/api"
	"github.com/valivishy/pokedex/internal/cache"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	c := cache.NewCache(5 * time.Second)

	key := "pikachu"
	value := []byte("static-electric")

	c.Add(key, value)

	got, ok := c.Get(key)
	if !ok {
		t.Fatalf("expected to find key %q in cache", key)
	}

	if string(got) != string(value) {
		t.Errorf("expected value %q, got %q", value, got)
	}
}

func TestReapRemovesExpiredEntries(t *testing.T) {
	c := cache.NewCache(10 * time.Millisecond)

	key := "bulbasaur"
	value := []byte("grass-poison")

	c.Add(key, value)

	time.Sleep(30 * time.Millisecond)

	_, ok := c.Get(key)
	if ok {
		t.Errorf("expected key %q to be removed after reaping", key)
	}
}

func TestCacheUsedForSubsequentCalls(t *testing.T) {
	callCount := 0

	mockResponse := api.LocationList{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results: []api.Location{
			{Name: "cached-location", URL: "https://pokeapi.co/test"},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		err := json.NewEncoder(w).Encode(mockResponse)
		if err != nil {
			t.Fatalf("error writing mock response: %v", err)
		}
	}))
	defer server.Close()

	cfg := &commands.Config{Next: &server.URL}

	_, err := api.List(cfg.Next)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	_, err = api.List(cfg.Next)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected server to be called once, but got %d calls", callCount)
	}
}
