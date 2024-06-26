package main

import (
	"time"

	"github.com/grevgeny/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Minute)
	cfg := &config{
		pokeapiClient:   pokeClient,
		catchedPokemons: map[string]pokeapi.Pokemon{},
	}

	initREPL(cfg)
}
