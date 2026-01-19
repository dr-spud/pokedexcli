package pokeapi

func (c *Client) GetExplore(location string) (exploreStruct, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + location
	var encounters exploreStruct
	if err := c.GetAndUnmarshal(url, &encounters); err != nil {
		return exploreStruct{}, err
	}
	return encounters, nil
}

func (c *Client) ListPokemonEncounters(encounters exploreStruct) []string {
	res := []string{}
	for _, pokemon := range encounters.PokemonEncounters {
		res = append(res, pokemon.Pokemon.Name)
	}
	return res
}
