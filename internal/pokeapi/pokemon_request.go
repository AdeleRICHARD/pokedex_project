package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name *string) (*PokemonInfos, error) {
	var pokemonInfo PokemonInfos

	if dataCache, ok := c.Cache.Get(*name); ok {
		err := json.Unmarshal(dataCache, &pokemonInfo)
		if err != nil {
			return nil, err
		}
		return &pokemonInfo, nil
	}

	endpoint := "pokemon/" + *name
	fullURL := baseUrl + endpoint

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 399 {
		return nil, fmt.Errorf("an error occured on pokemon request %s with status code %d", *name, res.StatusCode)
	}

	dataPokemon, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c.Cache.Add(*name, dataPokemon)

	err = json.Unmarshal(dataPokemon, &pokemonInfo)
	if err != nil {
		return nil, err
	}

	return &pokemonInfo, nil
}
