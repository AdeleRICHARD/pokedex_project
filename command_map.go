package main

import (
	"fmt"

	pokeapi "github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

func commandMap(pokeapiClient *pokeapi.Client) error {
	resp, err := pokeapiClient.GetLocationAreas(pokeapiClient.Config.NextUrl)
	if err != nil {
		return err
	}

	for _, location := range resp.Result {
		fmt.Println(*location.Name)
	}
	pokeapiClient.Config.NextUrl = resp.NextUrl
	if resp.PrevUrl != nil {
		pokeapiClient.Config.PrevUrl = resp.PrevUrl
	}

	return nil
}

func commandMapB(pokeapi *pokeapi.Client) error {
	command := commands["mapb"]
	if commands["map"].config != nil {
		command.config = commands["map"].config
	}
	/* previousLocations := command.config.GetPrev()

	if previousLocations == nil {
		return fmt.Errorf("no previous locations")
	}

	for _, location := range previousLocations {
		fmt.Println(*location.Name)
	} */

	return nil
}
