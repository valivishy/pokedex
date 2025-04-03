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

var BaseUrl = "https://pokeapi.co/api/v2/location-area"
var localCache = cache.NewCache(time.Second * 5)

func List(url *string) (*LocationList, error) {
	if url == nil {
		return nil, errors.New("no more locations in this direction")
	}

	body, err := get(*url)
	if err != nil {
		return nil, err
	}

	response := LocationList{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func Get(name string) (*LocationDetails, error) {
	body, err := get(fmt.Sprintf("%s/%s", BaseUrl, name))
	if err != nil {
		return nil, err
	}

	response := LocationDetails{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// region PRIVATE
func get(url string) ([]byte, error) {
	if value, ok := localCache.Get(url); ok {
		fmt.Printf("%s found in cache, returning early\n", url)
		return value, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s returned status code %d", url, resp.StatusCode)
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	localCache.Add(url, result)

	return result, nil
}

//endregion PRIVATE
