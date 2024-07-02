package main

import "github.com/pureit-dev/pokedexcli/internal/pokeapi"

func commandMap() error {
	err := pokeapi.GetPokeLocations()
	if err != nil {
		return err
	}
	return nil
}