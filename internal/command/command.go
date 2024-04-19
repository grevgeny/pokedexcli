package command

import (
	"fmt"
	"os"

	"github.com/grevgeny/pokedexcli/internal/pokeapi"
)

type ReplCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var Commands map[string]ReplCommand

var apiClient *pokeapi.APIClient

func init() {
	apiClient = pokeapi.NewAPIClient()

	Commands = map[string]ReplCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Fetch next 20 locations",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Fetch previous 20 location",
			Callback:    CommandMapb,
		},
	}
}

func CommandHelp() error {
	fmt.Print("\nWelcome to the Pokedex\nUsage:\n\n")
	for _, c := range Commands {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}
	fmt.Println()
	return nil
}

func CommandExit() error {
	os.Exit(0)
	return nil
}

func CommandMap() error {
	results, err := apiClient.GetNextLocations()
	if err != nil {
		return err
	}

	for _, result := range results {
		fmt.Println(result.Name)
	}

	return nil
}

func CommandMapb() error {
	results, err := apiClient.GetPreviousLocations()
	if err != nil {
		return err
	}

	for _, result := range results {
		fmt.Println(result.Name)
	}

	return nil
}
