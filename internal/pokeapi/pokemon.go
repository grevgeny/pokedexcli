package pokeapi

type Pokemon struct {
	Name  string `json:"name"`
	Stats []struct {
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
		BaseStat int `json:"base_stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	BaseExp float64 `json:"base_experience"`
	Height  int     `json:"height"`
	Weight  int     `json:"weight"`
}

func (client *Client) FetchPokemonInfo(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	var p Pokemon
	if err := client.fetchData(url, &p, true); err != nil {
		return Pokemon{}, err
	}

	return p, nil
}
