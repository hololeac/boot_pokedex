package pokeconfig

import (
	"github.com/hololeac/boot_pokedex/internal/pokecache"
)

type Config struct {
	Next  string
	Prev  string
	Cache pokecache.Cache
}

var config Config

func getMap()
