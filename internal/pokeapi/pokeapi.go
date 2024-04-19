package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type PaginationConfig struct {
	Previous string
	Next     string
}

type Locations struct {
	Next     string   `json:"next,omitempty"`
	Previous string   `json:"previous,omitempty"`
	Results  []Result `json:"results,omitempty"`
}

type Result struct {
	Name string `json:"name,omitempty"`
}

type APIClient struct {
	Config PaginationConfig
}

func NewAPIClient() *APIClient {
	return &APIClient{
		Config: PaginationConfig{
			Next: "https://pokeapi.co/api/v2/location-area",
		},
	}
}

func (client *APIClient) GetNextLocations() ([]Result, error) {
	if client.Config.Next == "" {
		return nil, errors.New("No next location\n")
	}

	return client.fetchData(client.Config.Next)
}

func (client *APIClient) GetPreviousLocations() ([]Result, error) {
	if client.Config.Previous == "" {
		return nil, errors.New("No previous locations\n")
	}
	return client.fetchData(client.Config.Previous)
}

func (client *APIClient) fetchData(url string) ([]Result, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		body, _ := io.ReadAll(res.Body)
		return nil, errors.New("HTTP error: " + string(body))
	}

	var loc Locations
	if err := json.NewDecoder(res.Body).Decode(&loc); err != nil {
		return nil, err
	}

	client.Config.Next = loc.Next
	client.Config.Previous = loc.Previous

	return loc.Results, nil

}
