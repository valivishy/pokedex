package tests

import (
	"github.com/valivishy/pokedex/commands"
	"github.com/valivishy/pokedex/internal/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExploreLocation_InvalidParams(t *testing.T) {
	err := commands.CommandExplore(nil, []string{})
	if err == nil || !strings.Contains(err.Error(), "no location provided") {
		t.Fatalf("expected area error, got: %v", err)
	}
}

func TestExploreLocation_APIFailure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	api.BaseUrl = ts.URL
	err := commands.CommandExplore(nil, []string{"kanto"})
	if err == nil || !strings.Contains(err.Error(), "500") {
		t.Fatalf("expected 500 error, got: %v", err)
	}
}

func TestCommandExplore_Success(t *testing.T) {
	called := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"locations": []}`))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	api.BaseUrl = ts.URL
	err := commands.CommandExplore(nil, []string{"kanto"})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if !called {
		t.Fatal("expected handler to be called")
	}
}
