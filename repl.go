package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/grevgeny/pokedexcli/internal/pokeapi"
)

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	catchedPokemons  map[string]pokeapi.Pokemon
	pokeapiClient    pokeapi.Client
}

func initREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		processedInput := processInput(scanner.Text())

		cmdName := processedInput[0]
		var params []string
		if len(processedInput) > 1 {
			params = processedInput[1:]
		}

		command, ok := getCommands()[cmdName]
		if !ok {
			fmt.Print("\nCommand not supported! Type 'help' to see available commands.\n\n")
			continue
		}

		if err := command.callback(cfg, params...); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command: %s\n", err)
		}
	}
}

func processInput(input string) []string {
	loweredInput := strings.ToLower(input)
	return strings.Fields(loweredInput)
}
