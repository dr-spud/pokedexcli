package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
	prevURL       *string
	nextURL       *string
}

func replLoop(cfg *config) error {
	supportedCommands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	cfg.pokedex = make(map[string]pokeapi.Pokemon)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleanUserInput := cleanInput(scanner.Text())
		if len(cleanUserInput) == 0 {
			continue
		}

		command, exists := supportedCommands[cleanUserInput[0]]
		args := cleanUserInput[1:]

		if !exists {
			fmt.Println("Unknown command:", cleanUserInput[0])
			continue
		}

		if err := command.callback(cfg, args...); err != nil {
			fmt.Println(fmt.Errorf("error: %w", err))
			continue
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays locations / goes to next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays locations / goes to previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "displays pokemon in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "tries to catch a pokemon",
			callback:    commandCatch,
		},
	}
	return supportedCommands
}
