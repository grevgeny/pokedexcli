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

var commands map[string]replCommand

func init() {
	commands = map[string]replCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex\nUsage:\n\n")
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		c, ok := commands[input]
		if !ok {
			fmt.Print("\nCommand not supported! Type 'help' to see available commands.\n\n")
			continue
		}

		if err := c.callback(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command: %s", err)
		}
	}
}
