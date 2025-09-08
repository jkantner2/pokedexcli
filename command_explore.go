package main

import(
	"fmt"
)

func commandExplore(cfg *config, location []string) error {
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon")
	for _, exp := range exploreResp.PokemonEncounters {
		fmt.Printf("- %v\n",exp.Pokemon.Name)
	}
	
	return nil
}
