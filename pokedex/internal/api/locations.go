package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valivishy/pokedex/internal/cache"
	"io"
	"net/http"
	"time"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Response struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

var localCache = cache.NewCache(time.Second * 5)

func List(url *string) (*Response, error) {
	if url == nil {
		return nil, errors.New("no more locations in this direction")
	}

	body, err := get(url)
	if err != nil {
		return nil, err
	}

	response := Response{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// region PRIVATE
func get(url *string) ([]byte, error) {
	if value, ok := localCache.Get(*url); ok {
		fmt.Printf("%s found in cache, returning early\n", *url)
		return value, nil
	}

	resp, err := http.Get(*url)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	localCache.Add(*url, result)

	return result, nil
}

//endregion PRIVATE
