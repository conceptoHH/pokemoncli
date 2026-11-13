package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)


type cliCommand struct {
	name string
	description string
	callback func() error
}


func cleanInput(text string) []string {
	txt_space_trim := strings.TrimSpace(text)
	return strings.Split(strings.ToLower(txt_space_trim)," ")
	
}



func startRepl() {
	commands := map[string]cliCommand{
		"help": {
			name: "help",
			description: "prints the usage & commands",
			callback: commandHelp(commands),
		},
		"exit": {
			name: "exit",
			description: "exits the Pokedex program",
			callback: commandExit,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")
		scanner.Scan() 
		input_cln := cleanInput(scanner.Text())
		
		if len(input_cln) == 0 {
			continue
		}
		
		if cmd, ok := commands[input_cln[0]]; ok {
			cmd.callback()
		} else {

			fmt.Printf("Unknown command: %v\n",input_cln[0])
		}

		
	}
}
