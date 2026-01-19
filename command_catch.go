package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	const CATCH_MAX = 500.00
	const MIN_CHANCE = 0.05
	const CATCH_DENOM = 600.00 // increase to lower catch chances, decrease to increase catch chances

	if len(args) == 0 {
		return errors.New("please enter a pokemon to catch")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pData, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	catchChance := (CATCH_MAX - float64(pData.BaseExperience)) / CATCH_DENOM
	if catchChance < MIN_CHANCE {
		catchChance = MIN_CHANCE
	}
	if rand.Float64() < catchChance {
		fmt.Printf("Caught %s\n", name)
		cfg.pokedex[name] = pData
	} else {
		fmt.Printf("%s escaped...\n", name)
	}

	return nil
}
