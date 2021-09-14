package controllers

import (
	pokemon "Pokemon-API/Pokemon"
)

type PokemonController struct {
	pokemonInf pokemon.ModelInterface
}

func NewPokemonController(pokemonInf pokemon.ModelInterface) pokemon.ControllerInterface {
	return &PokemonController {
		pokemonInf: pokemonInf,
	}
}

func (c *PokemonController) GetPokemons() ([]pokemon.Entity, error){
	rows, err := c.pokemonInf.GetPokemons()
	if err != nil {
		return nil, err
	}
	return rows, err
}

func (c *PokemonController) GetPokemonByID(id int) (*pokemon.Entity, error) {
	row, _ := c.pokemonInf.GetPokemonByID(id)
	return row, nil
}

func (c *PokemonController) AddPokemon(data *pokemon.Entity) (*pokemon.Entity, error) {
	p, err := c.pokemonInf.AddPokemon(data)
	return p, err
}