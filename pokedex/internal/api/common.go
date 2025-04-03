package api

import (
	"fmt"
	"github.com/valivishy/pokedex/internal/cache"
	"io"
	"net/http"
	"time"
)

var localCache = cache.NewCache(time.Second * 5)

func Get(url string) ([]byte, error) {
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
