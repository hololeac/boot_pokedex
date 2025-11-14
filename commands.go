package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokeApiResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location/"

	if config != nil && config.Next != "" {
		url = config.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// body, err := io.ReadAll(res.Body)
	// bodyString := string(body)
	// if bodyString == "" {
	// 	return nil
	// }

	var pokeApiStruct PokeApiResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeApiStruct)
	if err != nil {
		return err
	}

	for _, location := range pokeApiStruct.Results {
		fmt.Println(location.Name + "-area")
	}

	config.Next = pokeApiStruct.Next
	config.Prev = pokeApiStruct.Previous
	return nil
}

func commandMapb(config *Config) error {
	var url string
	if config != nil && config.Prev != "" {
		url = config.Prev
	} else {
		fmt.Println("Mapb is impossible, no previous locations available")
		return nil
	}

	var pokeApiStruct PokeApiResponse

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&pokeApiStruct)
	if err != nil {
		return err
	}

	for _, location := range pokeApiStruct.Results {
		fmt.Println(location.Name + "-area")
	}

	config.Next = pokeApiStruct.Next
	config.Prev = pokeApiStruct.Previous
	return nil
}
