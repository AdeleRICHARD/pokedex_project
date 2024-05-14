package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/AdeleRICHARD/pokedexcli/internal/pokeapi"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Client, *string) error
	config      *pokeapi.Config
}

var commands map[string]CliCommand

func init() {
	commands = map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the map",
			callback:    commandMap,
			config:      &pokeapi.Config{},
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous map",
			callback:    commandMapB,
			config:      &pokeapi.Config{},
		},
		"explore": {
			name:        "explore",
			description: "Explore the map",
			callback:    commandeExplore,
			config:      &pokeapi.Config{},
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon in your pokedex",
			callback:    commandCatch,
			config:      &pokeapi.Config{},
		},
	}
}

func startRepl(client *pokeapi.Client) {
	var name *string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		if len(lineSplit) > 1 {
			name = &lineSplit[1]
		}

		if command, ok := commands[lineSplit[0]]; ok {
			if err := command.callback(client, name); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
