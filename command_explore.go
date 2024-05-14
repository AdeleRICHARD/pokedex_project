package main

import (
	"fmt"

	"github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

func commandeExplore(pokeapi *pokeapi.Client, name *string) error {
	if name == nil {
		return fmt.Errorf("no names given")
	}
	resp, err := pokeapi.GetLocationInfo(name)
	if err != nil {
		return err
	}

	for _, pokemon := range resp.Result {
		fmt.Println(*pokemon.Name)
	}

	return nil
}
