package pokeapi

func (c *Client) GetMap(url string) (locationStruct, error) {
	var location locationStruct
	if err := c.GetAndUnmarshal(url, &location); err != nil {
		return locationStruct{}, err
	}
	return location, nil
}

func (c *Client) GetMapNames(locations locationStruct) ([]string, error) {
	res := []string{}
	for _, result := range locations.Results {
		res = append(res, result.Name)
	}
	return res, nil
}
