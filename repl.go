package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func dexScan() {
	scanner := bufio.NewScanner(os.Stdin)
	type cliCommand struct {
		name        string
		description string
		callback    func() error
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
	}

	for {
		fmt.Print("Pokedex > ")
		scanBool := scanner.Scan()
		text := scanner.Text()
		if scanBool == false {
			fmt.Println("fail")
		}
		clean := cleanInput(text)
		input := clean[0]
		cmd, exists := m[input]
		if exists {
			cmd.callback()
		} else {
			fmt.Println("Unkown Command")
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}
