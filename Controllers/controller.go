package controllers

import (
	pokemon "Pokemon-API/Pokemon"
)

type PokemonController struct {
	pokemonInf pokemon.PokemonInterface
}

func NewPokemonController(pokemonInf pokemon.PokemonInterface) pokemon.PokemonController {
	return &PokemonController {
		pokemonInf: pokemonInf,
	}
}

func (c *PokemonController) GetPokemons() ([]pokemon.PokemonFiled, error){
	rows, err := c.pokemonInf.GetPokemons()
	if err != nil {
		return nil, err
	}
	return rows, err
}

func (c *PokemonController) GetPokemonById(id int) (*pokemon.PokemonFiled, error) {
	row, _ := c.pokemonInf.GetPokemonById(id)
	return row, nil
}

func (c *PokemonController) AddPokemon(data *pokemon.PokemonFiled) (*pokemon.PokemonFiled, error) {
	p, err := c.pokemonInf.AddPokemon(data)
	return p, err
}