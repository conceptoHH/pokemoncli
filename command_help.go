package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}
