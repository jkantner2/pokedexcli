package main

import (
	"time"
	"github.com/jkantner2/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		capturedPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}

