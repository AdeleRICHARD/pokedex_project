package main

import (
	"os"

	"github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

func commandExit(pokeClient *pokeapi.Client, name *string) error {
	os.Exit(0)
	return nil
}
