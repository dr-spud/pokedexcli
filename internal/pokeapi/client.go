// Package pokeapi
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	newCache := pokecache.NewCache(cacheInterval)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: newCache,
	}
}

func (c *Client) GetAndUnmarshal(url string, dest any) error {
	data, exists := c.cache.Get(url)
	if exists {
		if err := json.Unmarshal(data, dest); err != nil {
			return err
		}
		return nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	newData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	c.cache.Add(url, newData)
	if err := json.Unmarshal(newData, dest); err != nil {
		return err
	}
	return nil
}
