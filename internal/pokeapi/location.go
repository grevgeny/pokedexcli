package pokeapi

import (
	"encoding/json"
	"io"
)

type Locations struct {
	Next         *string    `json:"next,omitempty"`
	Previous     *string    `json:"previous,omitempty"`
	LocationList []Location `json:"results,omitempty"`
}

type Location struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (client *Client) FetchOneLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	var l Location

	if value, ok := client.cache.Get(url); ok {
		err := json.Unmarshal(value, &l)
		if err != nil {
			return Location{}, err
		}
		return l, nil
	}

	res, err := client.Get(url)
	if err != nil {
		return Location{}, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	if err := json.Unmarshal(data, &l); err != nil {
		return Location{}, err
	}

	client.cache.Add(url, data)

	return l, nil
}

func (client *Client) FetchLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var loc Locations

	if value, ok := client.cache.Get(url); ok {
		err := json.Unmarshal(value, &loc)
		if err != nil {
			return Locations{}, err
		}
		return loc, nil
	}

	res, err := client.Get(url)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	if err := json.Unmarshal(data, &loc); err != nil {
		return Locations{}, err
	}

	client.cache.Add(url, data)

	return loc, nil
}
