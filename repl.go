package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Marcos-Pablo/pokedex-go-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	previousLocationURL *string
	nextLocationURL     *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	commandRegistry := getAvailableCommands()

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			fmt.Println("Command name not provided")
			continue
		}

		cmdName := words[0]
		cmd, exists := commandRegistry[cmdName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg)
		if err != nil {
			fmt.Printf("Error while executing command: %v", err)
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
