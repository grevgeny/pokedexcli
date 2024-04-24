package pokeapi

type Pokemon struct {
	BaseExp float64 `json:"base_experience"`
}

func (client *Client) FetchPokemonInfo(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	var p Pokemon
	if err := client.fetchData(url, &p, true); err != nil {
		return Pokemon{}, err
	}

	return p, nil
}
