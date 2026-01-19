package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	commands := getCommands()
	helpOrder := []string{"help", "exit", "map", "mapb"}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, commandName := range helpOrder {
		command := commands[commandName]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
