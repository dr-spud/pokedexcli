package main

import (
	"os"
	"time"

	"pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	if err := replLoop(cfg); err != nil {
		os.Exit(1)
	}
}
