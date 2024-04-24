package pokeapi

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

	var loc Location
	if err := client.fetchData(url, &loc, true); err != nil {
		return Location{}, err
	}

	return loc, nil
}

func (client *Client) FetchLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locs Locations
	if err := client.fetchData(url, &locs, true); err != nil {
		return Locations{}, err
	}

	return locs, nil
}
