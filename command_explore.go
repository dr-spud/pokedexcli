package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Errorf("you must provide a location name")
	}
	fmt.Printf("Exploring %s\n", args[0])
	encounters, err := cfg.pokeapiClient.GetExplore(args[0])
	if err != nil {
		return err
	}
	pokemon := cfg.pokeapiClient.ListPokemonEncounters(encounters)
	fmt.Println("Found pokemon:")
	for _, p := range pokemon {
		fmt.Println(p)
	}

	return nil
}
