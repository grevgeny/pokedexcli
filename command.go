package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type replCommand struct {
	callback    func(*config, ...string) error
	name        string
	description string
}

func getCommands() map[string]replCommand {
	return map[string]replCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "List of all the Pokemon in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try to catch the Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Get info about a catched Pokemon",
			callback:    commandInspect,
		},
	}
}

func commandHelp(cfg *config, params ...string) error {
	fmt.Print("\nWelcome to the Pokedex\nUsage:\n\n")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, params ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, params ...string) error {
	locations, err := cfg.pokeapiClient.FetchLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, result := range locations.LocationList {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(cfg *config, params ...string) error {
	locations, err := cfg.pokeapiClient.FetchLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range locations.LocationList {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(cfg *config, params ...string) error {
	if len(params) != 1 {
		return errors.New("you must provide a Location name")
	}

	locationName := params[0]
	fmt.Printf("Exploring %s...\n", locationName)

	location, err := cfg.pokeapiClient.FetchOneLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, p := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, params ...string) error {
	if len(params) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := params[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.FetchPokemonInfo(name)
	if err != nil {
		return err
	}

	const maxBaseExp = 608.0
	catchProb := 1 - pokemon.BaseExp/maxBaseExp

	if rand.Float64() >= catchProb {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		cfg.catchedPokemons[name] = pokemon
	}

	return nil
}

func commandInspect(cfg *config, params ...string) error {
	if len(params) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := params[0]

	pokemon, ok := cfg.catchedPokemons[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
