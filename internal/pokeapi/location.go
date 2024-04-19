package pokeapi

import (
	"encoding/json"
	"io"
)

type Locations struct {
	Next     *string  `json:"next,omitempty"`
	Previous *string  `json:"previous,omitempty"`
	Results  []Result `json:"results,omitempty"`
}

type Result struct {
	Name string `json:"name,omitempty"`
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
		return Locations{}, nil
	}

	if err := json.Unmarshal(data, &loc); err != nil {
		return Locations{}, err
	}

	client.cache.Add(url, data)

	return loc, nil

}
