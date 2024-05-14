package main

import (
	"fmt"
	"math/rand"

	"github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

func commandCatch(pokeApi *pokeapi.Client, name *string) error {
	if name == nil {
		return fmt.Errorf("name is empty")
	}

	pokemonInfos, err := pokeApi.GetPokemonInfo(name)
	if err != nil {
		return err
	}

	if pokemonInfos == nil {
		return fmt.Errorf("no pokemon found with name %s", *name)
	}
	const threshold = 50
	randNum := rand.Intn(pokemonInfos.Exp)

	fmt.Printf("Throwing a Pokeball at %s\n", *name)
	if randNum < threshold {
		fmt.Printf("%s was caught\n", *name)
		pokeApi.Config.Pokedex[*name] = pokemonInfos
		return nil
	}

	fmt.Printf("%s escaped!\n", *name)
	return nil

}
