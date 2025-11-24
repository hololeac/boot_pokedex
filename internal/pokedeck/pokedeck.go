package pokedeck

import "fmt"

type PokemonApiStruct struct {
	BaseExperience int    `json:"base_experience"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  struct {
		HP             int
		Attack         int
		Defence        int
		SpecialAttack  int
		SpecialDefence int
		Speed          int
	}
	Types []string
}

var Deck map[string]Pokemon

func translateApiToDeckStruct(pokemonApiResponse *PokemonApiStruct, pokemonStruct *Pokemon) {
	pokemonStruct.Name = pokemonApiResponse.Name
	pokemonStruct.Height = pokemonApiResponse.Height
	pokemonStruct.Weight = pokemonApiResponse.Weight

	for _, st := range pokemonApiResponse.Stats {
		switch st.Stat.Name {
		case "hp":
			pokemonStruct.Stats.HP = st.BaseStat
		case "attack":
			pokemonStruct.Stats.Attack = st.BaseStat
		case "defence":
			pokemonStruct.Stats.Defence = st.BaseStat
		case "special-attack":
			pokemonStruct.Stats.SpecialAttack = st.BaseStat
		case "special-defence":
			pokemonStruct.Stats.SpecialDefence = st.BaseStat
		case "speed":
			pokemonStruct.Stats.Speed = st.BaseStat
		}
	}

	for _, ty := range pokemonApiResponse.Types {
		pokemonStruct.Types = append(pokemonStruct.Types, ty.Type.Name)
	}
}

func AddToDeck(pokemonApiResponse *PokemonApiStruct, deck map[string]Pokemon) error {
	var pokemon Pokemon
	_, alreadyInDeck := deck[pokemonApiResponse.Name]
	if alreadyInDeck {
		fmt.Printf("%v already beed caught! releasing...\n", pokemonApiResponse.Name)
		return nil
	}

	translateApiToDeckStruct(pokemonApiResponse, &pokemon)
	deck[pokemon.Name] = pokemon
	fmt.Printf("%v was added to a deck!\n", pokemon.Name)
	return nil
}

func init() {
	Deck = make(map[string]Pokemon)
}
