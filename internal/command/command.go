package command

import (
	"fmt"
	"os"
)

type ReplCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var Commands map[string]ReplCommand

func init() {
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
