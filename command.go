package main

import (
	"errors"
	"fmt"
	"os"
)

type replCommand struct {
	callback    func(*config, []string) error
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
			name:        "explore",
			description: "List of all the Pokemon in a given area",
			callback:    commandExplore,
		},
	}
}

func commandHelp(cfg *config, params []string) error {
	fmt.Print("\nWelcome to the Pokedex\nUsage:\n\n")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, params []string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, params []string) error {
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

func commandMapb(cfg *config, params []string) error {
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

func commandExplore(cfg *config, params []string) error {
	if len(params) != 1 {
		return errors.New("invalid number of arguements")
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
