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
	pokeapiClient    pokeapi.Client
}

func initREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		inputSplit := strings.SplitN(input, " ", 2)

		cmdName := inputSplit[0]
		var params []string
		if len(inputSplit) == 2 {
			params = strings.Split(inputSplit[1], " ")
		}

		command, ok := getCommands()[cmdName]
		if !ok {
			fmt.Print("\nCommand not supported! Type 'help' to see available commands.\n\n")
			continue
		}

		if err := command.callback(cfg, params); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command: %s", err)
		}
	}
}
