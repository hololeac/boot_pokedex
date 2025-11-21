package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type PokemonsStruct struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExit(config *Config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, param string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *Config, param string) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if config != nil && config.Next != "" {
		url = config.Next
	}

	res, ok := cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		res, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cache.Add(url, res)
	}

	var pokeApiStruct PokeApiResponse

	decoder := json.NewDecoder(bytes.NewReader(res))
	err := decoder.Decode(&pokeApiStruct)
	if err != nil {
		return err
	}

	for _, location := range pokeApiStruct.Results {
		fmt.Println(location.Name)
	}

	config.Next = pokeApiStruct.Next
	config.Prev = pokeApiStruct.Previous
	return nil
}

func commandMapb(config *Config, param string) error {
	var url string
	if config != nil && config.Prev != "" {
		url = config.Prev
	} else {
		fmt.Println("Mapb is impossible, no previous locations available")
		return nil
	}

	res, ok := cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		res, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		cache.Add(url, res)
	}

	var pokeApiStruct PokeApiResponse

	err := json.NewDecoder(bytes.NewReader(res)).Decode(&pokeApiStruct)
	if err != nil {
		return err
	}

	for _, location := range pokeApiStruct.Results {
		fmt.Println(location.Name)
	}

	config.Next = pokeApiStruct.Next
	config.Prev = pokeApiStruct.Previous
	return nil
}

func commandExplore(config *Config, param string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + param

	res, ok := cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		res, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		cache.Add(url, res)
	}

	var pokemons PokemonsStruct
	err := json.NewDecoder(bytes.NewReader(res)).Decode(&pokemons)
	if err != nil {
		return err
	}

	fmt.Println("Exploring" + param + "...")

	if len(pokemons.PokemonEncounters) == 0 {
		fmt.Println("Pokemons not found!")
	} else {
		fmt.Println("Found Pokemon:")
		for _, pokemon := range pokemons.PokemonEncounters {
			fmt.Println(" - " + pokemon.Pokemon.Name)
		}
	}
	return nil
}
