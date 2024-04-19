package main

import (
	"bufio"
	"fmt"
	"os"

	pokeapi "github.com/AdeleRICHARD/pokedexcli/pokeApi"
)

type CliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	command := commands["map"]
	if command.config == nil {
		command.config = pokeapi.NewConfig()
		commands["map"] = command
	}

	locations := command.config.GetNext()

	for _, location := range locations {
		fmt.Println(*location.Name)
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if command, ok := commands[line]; ok {
			if err := command.callback(); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
