package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	var url string
	if cfg.nextURL == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *cfg.nextURL
	}

	locations, err := cfg.pokeapiClient.GetMap(url)
	if err != nil {
		return err
	}

	locationNames, err := cfg.pokeapiClient.GetMapNames(locations)
	if err != nil {
		return err
	}

	for _, name := range locationNames {
		fmt.Println(name)
	}

	cfg.nextURL = locations.Next
	cfg.prevURL = locations.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	var url string
	if cfg.prevURL == nil {
		return errors.New("you're on the first page")
	} else {
		url = *cfg.prevURL
	}

	locations, err := cfg.pokeapiClient.GetMap(url)
	if err != nil {
		return err
	}

	locationNames, err := cfg.pokeapiClient.GetMapNames(locations)
	if err != nil {
		return err
	}

	for _, name := range locationNames {
		fmt.Println(name)
	}

	cfg.nextURL = locations.Next
	cfg.prevURL = locations.Previous

	return nil
}
