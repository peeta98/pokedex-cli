package main

import (
	"errors"
	"fmt"

	"github.com/peeta98/pokedex-cli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, exists := cfg.caughtPokemon[name]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}
	printPokemonInf(pokemon)
	return nil
}

func printPokemonInf(pokemon pokeapi.Pokemon) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, pokemonStat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}
}
