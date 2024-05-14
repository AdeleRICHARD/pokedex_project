package main

import (
	"fmt"

	"github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

func commandInspect(pokeApi *pokeapi.Client, name *string) error {
	if name == nil {
		return fmt.Errorf("no pokemon name given")
	}

	pokemonInfo := pokeApi.Config.Pokedex[*name]
	if pokemonInfo == nil {
		fmt.Println("You have not caught that Pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", *name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stat {
		fmt.Printf("  -%s: %d\n", *stat.Name, pokemonInfo.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonInfo.Types {
		fmt.Printf("  - %s\n", *t.Type.Name)
	}

	return nil
}
