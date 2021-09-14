package mocks

import (
	pokemon "Pokemon-API/Pokemon"

	"github.com/stretchr/testify/mock"
)

type PokemonControllerMock struct {
	Mock mock.Mock
}

func (controller *PokemonControllerMock) GetPokemons() ([]pokemon.Entity, error) {
	args := controller.Mock.Called()
	return args.Get(0).([]pokemon.Entity), args.Error(1)
}

func (controller *PokemonControllerMock) GetPokemonByID(id int) (*pokemon.Entity, error) {
	args := controller.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(0)
	}
	return args.Get(0).(*pokemon.Entity), args.Error(1)
}

func (controller *PokemonControllerMock) AddPokemon(data *pokemon.Entity) (*pokemon.Entity, error) {
	args := controller.Mock.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(0)
	}
	return args.Get(0).(*pokemon.Entity), args.Error(1)
}

