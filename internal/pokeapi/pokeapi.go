package pokeapi

import (
	"io"
	"net/http"

	"github.com/Sveta-Rozhok/pokedex/internal/pokecache"
)

func FetchFromPokeAPI(url string, cache *pokecache.Cache) ([]byte, error) {
	if data, found := cache.Get(url); found {
		return data, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cache.Add(url, body)
	return body, nil
}
