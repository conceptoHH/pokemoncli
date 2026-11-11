package main

import (
	"os"
	"fmt"
)

func commandExit() error {
	fmt.Print("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
