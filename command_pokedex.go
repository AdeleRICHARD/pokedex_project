package main

import (
	"fmt"

	"github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

func commandPokedex(pokeApi *pokeapi.Client, _ *string) error {
	fmt.Println("Pokedex:")
	for name := range pokeApi.Config.Pokedex {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
