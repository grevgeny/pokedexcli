package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/grevgeny/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		command, ok := getCommands()[input]
		if !ok {
			fmt.Print("\nCommand not supported! Type 'help' to see available commands.\n\n")
			continue
		}

		if err := command.callback(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command: %s", err)
		}
	}
}
