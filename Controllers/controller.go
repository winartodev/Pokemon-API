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
	// refer to pokemonModel and call GetPokemons function
	// which is used to querying all data into db
	// and return 2 values
	rows, err := c.pokemonModel.GetPokemons()
	if err != nil {
		return nil, err
	}
	return rows, err
}

// GetPokemonByID method to return specified data pokemon by ID
func (c *PokemonController) GetPokemonByID(id int) (*pokemon.Entity, error) {
	// refer to pokemonModel and call GetPokemonByID function
	// which is used to querying pokemon data by ID into db
	// and return 2 values
	row, _ := c.pokemonModel.GetPokemonByID(id)
	return row, nil
}

// Add pokemon method to return pokemon.Entity by data parameter
func (c *PokemonController) AddPokemon(data *pokemon.Entity) (*pokemon.Entity, error) {
	// refer to pokemonModel and call AddPokemon function
	// which is used to querying new pokemon into db
	// return 2 values
	p, err := c.pokemonModel.AddPokemon(data)
	return p, err
}
