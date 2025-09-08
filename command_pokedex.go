package main

import (
	"fmt"
)

func commandPokedex (cfg *config, argument [] string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.capturedPokemon) == 0 {
		fmt.Println("You haven't caught any pokemon!")
		return nil
	}
	for _, pokemon := range cfg.capturedPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil

}
