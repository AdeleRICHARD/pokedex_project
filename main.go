package main

import (
	pokeapi "github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()
	startRepl(pokeapiClient)
}
