package main

import (
	"fmt"
	"math/rand"
)
func commandCatch(cfg *config, pokemonName []string) error{
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	
	pokemonInfo, _ := cfg.pokeapiClient.GetPokemonInfo(pokemonName)

	minChance := 10
	maxChance := 90
	normalizedBaseExperience := (pokemonInfo.BaseExperience * 1000) / 608
	escapeThreshold := maxChance - (((maxChance - minChance) * normalizedBaseExperience) / 1000)
	captureValue := rand.Intn(90)
	if captureValue > escapeThreshold {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonInfo.Name)
	cfg.capturedPokemon[pokemonInfo.Name] = pokemonInfo
	return nil
}
