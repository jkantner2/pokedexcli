package main

import (
	"time"
	"github.com/jkantner2/pokedexcli"
)

func main() {
	pokeClient := pokedexcli.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}

