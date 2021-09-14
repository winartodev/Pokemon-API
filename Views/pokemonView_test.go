package view

import (
	"Pokemon-API/Controllers/mocks"
	pokemon "Pokemon-API/Pokemon"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPokemonView_GetPokemons(t *testing.T) {
	tests := []struct {
		Name         string
		MockPokemon  *[]pokemon.Entity
		ErrorMessage error
		StatusCode   int
		Expected     string
	}{
		{
			Name: "Get Pokemons Success",
			MockPokemon: &[]pokemon.Entity{
				{
					ID:      1,
					Name:    "Pikachu",
					Species: "Mouse",
				},
				{
					ID:      2,
					Name:    "Charmander",
					Species: "Lizard Pokémon",
				},
				{
					ID:      3,
					Name:    "Squirtle",
					Species: "Tiny Turtle Pokémon",
				},
			},
			ErrorMessage: nil,
			StatusCode:   http.StatusOK,
			Expected:     `[{"id":1,"name":"Pikachu","species":"Mouse"},{"id":2,"name":"Charmander","species":"Lizard Pokémon"},{"id":3,"name":"Squirtle","species":"Tiny Turtle Pokémon"}]`,
		},
		{
			Name: "Get Pokemon Failed",
			MockPokemon: &[]pokemon.Entity{
				{
					ID:      1,
					Name:    "Pikachu",
					Species: "Mouse",
				},
			},
			ErrorMessage: errors.New("Internal Server Error"),
			StatusCode:   http.StatusInternalServerError,
			Expected:     `{"error":"Internal Server Error"}[{"id":1,"name":"Pikachu","species":"Mouse"}]`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			mockListPokemon := make([]pokemon.Entity, 0)
			mockListPokemon = append(mockListPokemon, *tc.MockPokemon...)

			mockPokemonController := &mocks.PokemonControllerMock{Mock: mock.Mock{}}
			mockPokemonController.Mock.On("GetPokemons").Return(mockListPokemon, tc.ErrorMessage)

			req, err := http.NewRequest(http.MethodGet, "/pokemons", nil)
			assert.NoError(t, err)

			pokemonView := PokemonView{pokemonController: mockPokemonController}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(pokemonView.GetPokemons)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.StatusCode, rr.Code)
			assert.Equal(t, tc.Expected, rr.Body.String())
		})
	}
}

func TestPokemonView_GetPokemonByID(t *testing.T) {
	tests := []struct {
		Name         string
		MockPokemon  *pokemon.Entity
		PokemonID    int
		ErrorMessage error
		StatusCode   int
		Expected     string
	}{
		{
			Name: "Get Pokemon Id Success",
			MockPokemon: &pokemon.Entity{
				ID:      1,
				Name:    "Pikachu",
				Species: "Mouse",
			},
			PokemonID:    1,
			ErrorMessage: nil,
			StatusCode:   http.StatusOK,
			Expected:     `{"id":1,"name":"Pikachu","species":"Mouse"}`,
		},
		{
			Name:         "Get Pokemon Id Not Found",
			MockPokemon:  nil,
			PokemonID:    2,
			ErrorMessage: errors.New("Id 2 Not Found"),
			StatusCode:   http.StatusNotFound,
			Expected:     `{"error":"Id 2 Not Found"}`,
		},
		{
			Name: "Get Pokemon Id Internal Server Error",
			MockPokemon: &pokemon.Entity{
				ID:      1,
				Name:    "Pikachu",
				Species: "Mouse",
			},
			PokemonID:    1,
			ErrorMessage: errors.New("Internal Server Error"),
			StatusCode:   http.StatusInternalServerError,
			Expected:     `{"error":"Internal Server Error"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			mockPokemonController := &mocks.PokemonControllerMock{Mock: mock.Mock{}}
			mockPokemonController.Mock.On("GetPokemonByID", tc.PokemonID).Return(tc.MockPokemon, tc.ErrorMessage)

			req, err := http.NewRequest(http.MethodGet, "/pokemon?id="+strconv.Itoa(tc.PokemonID), nil)
			assert.NoError(t, err)

			pokemonView := PokemonView{pokemonController: mockPokemonController}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(pokemonView.GetPokemonByID)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.StatusCode, rr.Code)
			assert.Equal(t, tc.Expected, rr.Body.String())
		})
	}
}

func TestPokemonView_AddPokemon(t *testing.T) {
	tests := []struct {
		Name         string
		MockPokemon  *pokemon.Entity
		ErrorMessage error
		StatusCode   int
		Expected     string
	}{
		{
			Name: "Add Pokemon Success",
			MockPokemon: &pokemon.Entity{
				ID:      1,
				Name:    "Pikachu",
				Species: "Mouse",
			},
			ErrorMessage: nil,
			StatusCode:   http.StatusOK,
			Expected:     `"Id Pokemon 1 Created"`,
		},
		{
			Name:         "Add Pokemon Internal Server Error",
			MockPokemon:  nil,
			ErrorMessage: errors.New("Internal Server Error"),
			StatusCode:   http.StatusInternalServerError,
			Expected:     `{"error":"Internal Server Error"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			mockPokemonController := &mocks.PokemonControllerMock{Mock: mock.Mock{}}
			mockPokemonController.Mock.On("AddPokemon", mock.AnythingOfType("*pokemon.Entity")).Return(tc.MockPokemon, tc.ErrorMessage)

			pokemonMarshal, _ := json.Marshal(tc.MockPokemon)
			req, err := http.NewRequest(http.MethodPost, "/pokemon", bytes.NewBuffer(pokemonMarshal))
			assert.NoError(t, err)

			pokemonView := PokemonView{pokemonController: mockPokemonController}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(pokemonView.AddPokemon)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.StatusCode, rr.Code)
			assert.Equal(t, tc.Expected, rr.Body.String())
		})
	}
}