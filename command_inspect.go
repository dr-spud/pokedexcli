package main

import (
	"fmt"

	"pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify a pokemon to inspect")
	}
	pName := args[0]
	pokemon, exists := cfg.pokedex[pName]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	} else {
		printPokeInfo(pokemon)
	}
	return nil
}

func printPokeInfo(pokemon pokeapi.Pokemon) {
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, s := range pokemon.Stats {
		fmt.Printf("    -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("    -%s\n", t.Type.Name)
	}
}
