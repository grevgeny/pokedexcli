package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/grevgeny/pokedexcli/internal/command"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		c, ok := command.Commands[input]
		if !ok {
			fmt.Print("\nCommand not supported! Type 'help' to see available commands.\n\n")
			continue
		}

		if err := c.Callback(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command: %s", err)
		}
	}
}
