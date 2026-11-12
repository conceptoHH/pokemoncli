package main

import (
	"fmt"
)

func commandHelp(cmd map[string]cliCommand) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, c := range cmd {
		fmt.Printf("%v: %v",c.name, c.description)
	}
	return nil
}
