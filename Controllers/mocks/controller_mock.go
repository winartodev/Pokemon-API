package mocks

import (
	pokemon "Pokemon-API/Pokemon"

	"github.com/stretchr/testify/mock"
)

type PokemonControllerMock struct {
	Mock mock.Mock
}

// GetPokemons this mock can return all data pokemon
func (controller *PokemonControllerMock) GetPokemons() ([]pokemon.Entity, error) {
	args := controller.Mock.Called()
	return args.Get(0).([]pokemon.Entity), args.Error(1)
}

// GetPokemonByID this mock can return specified data pokemon by ID
func (controller *PokemonControllerMock) GetPokemonByID(id int) (*[]pokemon.Entity, error) {
	args := controller.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(0)
	}
	return args.Get(0).(*[]pokemon.Entity), args.Error(1)
}

// AddPokemon this mock can return new pokemon
func (controller *PokemonControllerMock) AddPokemon(data *pokemon.Entity) (*pokemon.Entity, error) {
	args := controller.Mock.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(0)
	}
	return args.Get(0).(*pokemon.Entity), args.Error(1)
}
