package pokeapi

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var pData Pokemon
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	if err := c.GetAndUnmarshal(url, &pData); err != nil {
		return Pokemon{}, err
	}
	return pData, nil
}
