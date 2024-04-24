package pokeapi

import (
	"encoding/json"
	"io"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (client *Client) fetchData(url string, v interface{}, useCache bool) error {
	if useCache {
		if value, ok := client.cache.Get(url); ok {
			err := json.Unmarshal(value, v)
			if err != nil {
				return err
			}
			return nil
		}
	}

	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	if useCache {
		client.cache.Add(url, data)
	}

	return nil
}
