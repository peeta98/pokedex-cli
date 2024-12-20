package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("must provide location name")
	}

	locationArea := args[0]
	exploreLocationResp, err := cfg.pokeapiClient.ExploreLocation(locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationArea)

	if len(exploreLocationResp.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
	} else {
		fmt.Println("No pokemon found in this area.")
	}

	for _, pokemonEncounter := range exploreLocationResp.PokemonEncounters {
		fmt.Println("-", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
