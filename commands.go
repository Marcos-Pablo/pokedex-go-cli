package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getAvailableCommands() map[string]cliCommand {
	commandRegistry := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commandRegistry
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getAvailableCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	cfg.previousLocationURL = locationResp.Previous
	cfg.nextLocationURL = locationResp.Next

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationURL)
	if err != nil {
		return err
	}
	cfg.previousLocationURL = locationResp.Previous
	cfg.nextLocationURL = locationResp.Next

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	areaName := args[0]

	loc, err := cfg.pokeapiClient.ExploreLocation(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", loc.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range loc.PokemonEncounters {
		fmt.Printf("- %s\n", enc.Pokemon.Name)
	}

	return nil
}
