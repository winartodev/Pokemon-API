package controllers_test

import (
	controllers "Pokemon-API/Controllers"
	"Pokemon-API/Controllers/mocks"
	pokemon "Pokemon-API/Pokemon"
	"errors"
	"fmt"
	"strconv"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Test struct {
	Name 			string
	MockPokemon		*[]pokemon.Entity
	ErrorMessage	error
	Expected 		string
}

func TestPokemonController_GetPokemons(t *testing.T) {
	tests := []Test{
		{
			Name: "Get Pokemon 6 Row Success ",
			MockPokemon: &[]pokemon.Entity{
				{
					ID: 1,
					Name: "Bulbasaur",
					Species: "Seed Pokémon",
				},
				{
					ID: 2,
					Name: "Charmander",
					Species: "Lizard Pokémon",
				},
				{
					ID: 3,
					Name: "Squirtle",
					Species: "Tiny Turtle Pokémon",
				},
				{
					ID: 4,
					Name: "Raticate",
					Species: "Mouse Pokémon",
				},
				{
					ID: 5,
					Name: "Rattata",
					Species: "Mouse Pokémon",
				},
				{
					ID: 6,
					Name: "Rattata",
					Species: "Mouse Pokémon",
				},
			},
			ErrorMessage: nil,
			Expected: `[{1 Bulbasaur Seed Pokémon} {2 Charmander Lizard Pokémon} {3 Squirtle Tiny Turtle Pokémon} {4 Raticate Mouse Pokémon} {5 Rattata Mouse Pokémon} {6 Rattata Mouse Pokémon}]`,
		},
		{
			Name: "Get Pokemon 1 Row Success ",
			MockPokemon: &[]pokemon.Entity{
				{	
					ID: 1,
					Name: "Pikachu",
					Species: "Mouse",
				},
			},
			ErrorMessage: nil,
			Expected: `[{1 Pikachu Mouse}]`,
		},
		{
			Name: "Get Pokemon Empty Row Success",
			MockPokemon: &[]pokemon.Entity{},
			ErrorMessage: nil,
			Expected: `[]`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			mockListPokemon := make([]pokemon.Entity, 0)
			mockListPokemon = append(mockListPokemon, *tc.MockPokemon...)

			mockPokemonController := &mocks.PokemonControllerMock{Mock: mock.Mock{}}
			mockPokemonController.Mock.On("GetPokemons").Return(mockListPokemon, tc.ErrorMessage)

			c := controllers.NewPokemonController(mockPokemonController)
			pokemons, err := c.GetPokemons()

			assert.NoError(t, err)
			assert.NotNil(t, pokemons)
			assert.Equal(t, tc.Expected, fmt.Sprint(pokemons))
		})
	}
}

func TestPokemonView_GetPokemons_Failed(t *testing.T) {
	tests := []Test{
		{
			Name: "Get Pokemon Failed",
			MockPokemon: &[]pokemon.Entity{
				{
					ID: 1,
					Name: "Pikachu",
					Species: "Mouse",
				},
			},
			ErrorMessage: errors.New("Data Unreadable"),
			Expected: `[]`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			mockListPokemon := make([]pokemon.Entity, 0)
			mockListPokemon = append(mockListPokemon, *tc.MockPokemon...)

			mockPokemonController := &mocks.PokemonControllerMock{Mock: mock.Mock{}}
			mockPokemonController.Mock.On("GetPokemons").Return(mockListPokemon, tc.ErrorMessage)

			c := controllers.NewPokemonController(mockPokemonController)
			pokemons, err := c.GetPokemons()

			assert.Error(t, err)
			assert.Nil(t, pokemons)
			assert.Equal(t, tc.Expected, fmt.Sprint(pokemons))
		})
	}
}

func TestPokemonController_GetPokemonByID(t *testing.T) {
	tests := []Test{
		{
			Name: "Get Pokemon By Id Success",
			MockPokemon: &[]pokemon.Entity {
				{	ID: 1,
					Name: "Pikachu",
					Species: "Mouse",
				},
			},
			Expected: `{1 Pikachu Mouse}`,
		},
	}
	
	for _, tc := range tests {
			mockListPokemon := make([]pokemon.Entity, 0)
			mockListPokemon = append(mockListPokemon, *tc.MockPokemon...)

		mockPokemonController := &mocks.PokemonControllerMock{Mock: mock.Mock{}}
		mockPokemonController.Mock.On("GetPokemonByID", 1).Return(&mockListPokemon[0], nil)

		c := controllers.NewPokemonController(mockPokemonController)
		pokemons, err := c.GetPokemonByID(mockListPokemon[0].ID)

		assert.NoError(t, err)
		assert.NotNil(t, *pokemons)
		assert.Equal(t, tc.Expected, fmt.Sprint(*pokemons))
	}
}

func TestPokemonController_AddPokemon(t *testing.T) {
	tests := []Test{
		{
			Name: "Add Pokemon Success",
			MockPokemon: &[]pokemon.Entity{
				{
					ID: 1,
					Name: "Pikachu",
					Species: "Mouse",
				},
			},
			Expected: fmt.Sprint(int(1)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			mockListPokemon := make([]pokemon.Entity, 0)
			mockListPokemon = append(mockListPokemon, *tc.MockPokemon...)

			mockPokemonController := &mocks.PokemonControllerMock{Mock: mock.Mock{}}
			mockPokemonController.Mock.On("AddPokemon", mock.AnythingOfType("*pokemon.Entity")).Return(&mockListPokemon[0], nil)

			c := controllers.NewPokemonController(mockPokemonController)
			pokemons, err := c.AddPokemon(&mockListPokemon[0])

			expected, _ := strconv.Atoi(tc.Expected)

			assert.NoError(t, err)
			assert.NotNil(t, pokemons)
			assert.Equal(t, expected, pokemons.ID)
		})
	}
}