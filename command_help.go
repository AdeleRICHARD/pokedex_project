package main

import (
	"fmt"

	"github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

func commandHelp(pokeapi *pokeapi.Client, name *string) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
