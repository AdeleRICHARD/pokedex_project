package main

import (
	"fmt"

	pokeapi "github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

func commandMap(pokeapiClient *pokeapi.Client, name *string) error {
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

func commandMapB(pokeapi *pokeapi.Client, name *string) error {
	if pokeapi.Config.PrevUrl == nil {
		fmt.Println("You are at the beginning of the map")
		return nil
	}
	resp, err := pokeapi.GetLocationAreas(pokeapi.Config.PrevUrl)
	if err != nil {
		return err
	}

	for _, location := range resp.Result {
		fmt.Println(*location.Name)
	}

	pokeapi.Config.NextUrl = pokeapi.Config.PrevUrl
	if resp.PrevUrl != nil {
		pokeapi.Config.PrevUrl = resp.PrevUrl
	}

	return nil
}
