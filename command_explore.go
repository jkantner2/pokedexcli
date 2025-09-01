package main

import(
	"fmt"
)

func commandExplore(cfg *config, argument []string) error {
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(argument)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon")
	for _, exp := range exploreResp.PokemonEncounters {
		fmt.Printf("- %v\n",exp.Pokemon.Name)
	}
	
	return nil
}
