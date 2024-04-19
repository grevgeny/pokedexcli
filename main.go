package main

import (
	"bufio"
	"fmt"
	"os"
)

type replCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	panic("")
}

func commandExit() error {
	panic("")
}

func getCommands() map[string]replCommand {
	return map[string]replCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		}}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Print("Pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s", err)
	}
}
