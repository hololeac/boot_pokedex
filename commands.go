package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap() error {
	url := "https://pokeapi.co/api/v2/location/"
	// type location struct {
	// }

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	bodyString := string(body)

	if bodyString == "" {
		return nil
	}
	return nil
}
