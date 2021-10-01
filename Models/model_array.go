package models

import (
	pokemon "Pokemon-API/Pokemon"
	"fmt"
)

var (
	pokemonArray []pokemon.Entity
)

type pokemonStruct struct {
	pokemons *pokemon.Entity
}

// Initialize sql to modelInterface
func WithStruct(pokemons *pokemon.Entity) pokemon.ModelInterface {
	return &pokemonStruct{
		pokemons: pokemons,
	}
}

// GetPokemons return all data pokemon
func (m *pokemonStruct) GetPokemons() ([]pokemon.Entity, error) {
	pokemons := pokemonArray
	return pokemons, nil
}

// GetPokemonByID return specific pokemon by ID
func (m *pokemonStruct) GetPokemonByID(id int) (*[]pokemon.Entity, error) {
	pokemon := []pokemon.Entity{}
	for _, v := range pokemonArray {
		if id == v.ID {
			pokemon = append(pokemon, v)
			return &pokemon, nil
		}
	}
	return &pokemon, nil
}

// AddPokemon return new pokemon after insert query success
func (m *pokemonStruct) AddPokemon(p *pokemon.Entity) (*pokemon.Entity, error) {
	data := pokemon.Entity{
		ID:      p.ID,
		Name:    p.Name,
		Species: p.Species,
	}

	for _, v := range pokemonArray {
		if data.ID == v.ID {
			return nil, fmt.Errorf("duplicate entry %d for pokemon id", data.ID)
		}
	}

	pokemonArray = append(pokemonArray, data)
	return &pokemonArray[len(pokemonArray)-1], nil
}
