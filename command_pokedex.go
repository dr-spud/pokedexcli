package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Printf("Your pokedex:\n")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
