package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	type cliCommand struct {
		name        string
		description string
		callback    func(*config) error
	}

	m := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help menu",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows map locations in the pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows previous list of map locations in the pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "displays information for a particular map",
			callback:    commandExplore,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanBool := scanner.Scan()
		text := scanner.Text()
		if !scanBool {
			fmt.Println("fail")
		}
		clean := cleanInput(text)
		input := clean[0]
		cmd, exists := m[input]
		if exists {
			cmd.callback(cfg)
		} else {
			fmt.Println("Unkown Command")
		}
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Locations areas:")
	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

// map back
func commandMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("no previous page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Locations areas:")
	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func commandExplore(cfg *config) error {
	return nil
}
