package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Sveta-Rozhok/pokedex/internal/pokecache"
)

const cacheInterval = 5 * time.Minute

type Client struct {
	baseURL string
	cache   *pokecache.Cache
}

type Config struct {
	Next     string
	Previous string
}

type LocationArea struct {
	Name string
	URL  string
}

type LocationAreaResp struct {
	Count    int
	Next     *string
	Previous *string
	Results  []LocationArea
}

func NewClient() *Client {
	return &Client{
		baseURL: "https://pokeapi.co/api/v2",
		cache:   pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := c.baseURL + "/location-area"
	if pageURL != nil {
		endpoint = *pageURL
	}

	if data, ok := c.cache.Get(endpoint); ok {
		locations := LocationAreaResp{}
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locations, nil
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(endpoint, body)

	var locations LocationAreaResp
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locations, nil
}

type LocationAreaInfo struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationArea(locationName string) (LocationAreaInfo, error) {
	endpoint := c.baseURL + "/location-area/" + locationName

	if data, ok := c.cache.Get(endpoint); ok {
		locationInfo := LocationAreaInfo{}
		err := json.Unmarshal(data, &locationInfo)
		if err != nil {
			return LocationAreaInfo{}, err
		}
		return locationInfo, nil
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		return LocationAreaInfo{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaInfo{}, err
	}

	c.cache.Add(endpoint, body)

	var locationInfo LocationAreaInfo
	err = json.Unmarshal(body, &locationInfo)
	if err != nil {
		return LocationAreaInfo{}, err
	}

	return locationInfo, nil
}

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}

type Stat struct {
	BaseStat int  `json:"base_stat"`
	Stat     Name `json:"stat"`
}

type Type struct {
	Type Name `json:"type"`
}

type Name struct {
	Name string `json:"name"`
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := c.baseURL + "/pokemon/" + pokemonName

	if data, ok := c.cache.Get(endpoint); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(endpoint, body)

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
