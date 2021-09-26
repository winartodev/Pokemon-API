package view

import (
	errorhandler "Pokemon-API/Error"
	pokemon "Pokemon-API/Pokemon"
	response "Pokemon-API/Response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type PokemonView struct {
	pokemonController pokemon.ControllerInterface
}

// EndpointsHnadler handle all endpoints
func EndpointsHandler(pokemonController pokemon.ControllerInterface) {
	handler := &PokemonView{
		pokemonController: pokemonController,
	}

	http.HandleFunc("/", route)
	http.HandleFunc("/pokemons", handler.GetPokemons)
	http.HandleFunc("/pokemon", handler.GetPokemonByID)
}

func route(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "POKEMON - API WITH CLEAN ARCHITECTURE")
	}
}

// GetPokemons is used to get all pokemon
func (v *PokemonView) GetPokemons(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rows, err := v.pokemonController.GetPokemons()
		if err != nil {
			respond := response.NewResponseError(errorhandler.GetErrorCode(http.StatusInternalServerError, err))
			response.RespondWithJSON(w, respond, http.StatusInternalServerError)
			return
		}
		respond := response.NewResponseSuccess(rows)
		response.RespondWithJSON(w, respond, http.StatusOK)
	}
}

// GetPokemonByID is used to get pokemon by id
func (v *PokemonView) GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, _ := strconv.Atoi(r.FormValue("id"))

		row, err := v.pokemonController.GetPokemonByID(id)

		if row != nil && err != nil {
			respond := response.NewResponseError(errorhandler.GetErrorCode(http.StatusNotFound, err))
			response.RespondWithJSON(w, respond, http.StatusNotFound)
			return
		}

		if err != nil {
			respond := response.NewResponseError(errorhandler.GetErrorCode(http.StatusInternalServerError, err))
			response.RespondWithJSON(w, respond, http.StatusInternalServerError)
			return
		}

		respond := response.NewResponseSuccess(row)
		response.RespondWithJSON(w, respond, http.StatusOK)
	} else {
		v.AddPokemon(w, r)
	}
}

// AddPokemon is used to create new pokemon
func (v *PokemonView) AddPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		newPoke := json.NewDecoder(r.Body)
		var pokemon pokemon.Entity
		err := newPoke.Decode(&pokemon)

		if err != nil {
			respond := response.NewResponseError(errorhandler.GetErrorCode(http.StatusInternalServerError, err))
			response.RespondWithJSON(w, respond, http.StatusInternalServerError)
			return
		}

		row, err := v.pokemonController.AddPokemon(&pokemon)

		if err != nil {
			respond := response.NewResponseError(errorhandler.GetErrorCode(http.StatusBadRequest, err))
			response.RespondWithJSON(w, respond, http.StatusBadRequest)
			return
		}

		respond := response.NewPokemonCreateSuccess(row.ID)
		response.RespondWithJSON(w, respond, http.StatusOK)
	}
}
