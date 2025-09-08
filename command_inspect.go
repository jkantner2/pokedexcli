package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemonName []string) error{
	pokemonInfo, ok := cfg.capturedPokemon[pokemonName[0]]
	if !ok {
		fmt.Println("you have not caught this pokemon")
		return nil
	}
	fmt.Printf("Height: %v\n", pokemonInfo.Height)
	fmt.Printf("Wieght: %v\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, val := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %v\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range pokemonInfo.Types {
		fmt.Printf("- %s\n", val.Type.Name)
	}
	return nil
}
