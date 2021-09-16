package controllers

import (
	pokemon "Pokemon-API/Pokemon"
)

type PokemonController struct {
	pokemonModel pokemon.ModelInterface
}

func NewPokemonController(pokemonModel pokemon.ModelInterface) pokemon.ControllerInterface {
	return &PokemonController{
		pokemonModel: pokemonModel,
	}
}

// GetPokemon method to return all data pokemon
func (c *PokemonController) GetPokemons() ([]pokemon.Entity, error) {
	rows, err := c.pokemonModel.GetPokemons()
	if err != nil {
		return nil, err
	}
	return rows, err
}

// GetPokemonByID method to return specified data pokemon by ID
func (c *PokemonController) GetPokemonByID(id int) (*pokemon.Entity, error) {
	row, _ := c.pokemonModel.GetPokemonByID(id)
	return row, nil
}

// Add pokemon method to return pokemon.Entity by data parameter
func (c *PokemonController) AddPokemon(data *pokemon.Entity) (*pokemon.Entity, error) {
	p, err := c.pokemonModel.AddPokemon(data)
	return p, err
}
