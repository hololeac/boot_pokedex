package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hololeac/boot_pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

// Config holds runtime configuration passed to command callbacks.
type Config struct {
	Next string
	Prev string
}

// package-level config instance (can be populated in init or at runtime)
var config Config

var commands map[string]cliCommand

var cache *pokecache.Cache

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())

		command, ok := commands[words[0]]
		if ok {
			err := command.callback(&config)
			if err != nil {
				fmt.Printf("Error thrown: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	loweredInput := strings.ToLower(text)
	words := strings.Fields(loweredInput)

	return words
}

func init() {
	cache = pokecache.NewCache(5 * time.Second)
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the list of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the list of previous locations",
			callback:    commandMapb,
		},
	}
}
