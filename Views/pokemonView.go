package view

import (
	pokemon "Pokemon-API/Pokemon"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type PokemonView struct {
	pokemonController pokemon.PokemonController
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func EndpointsHandler(pokemonController pokemon.PokemonController) {
	handler := &PokemonView{
		pokemonController: pokemonController,
	}

	http.HandleFunc("/", route)
	http.HandleFunc("/pokemons", handler.GetPokemons)
	http.HandleFunc("/pokemon", handler.GetPokemonById)
}

func route(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "POKEMON - API WITH CLEAN ARCHITECTURE")
	}
}

func (v *PokemonView) GetPokemons(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		rows, err := v.pokemonController.GetPokemons()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		respondWithJSON(w, http.StatusOK, rows)
	}
}

func (v *PokemonView) GetPokemonById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, _ := strconv.Atoi(r.FormValue("id"))

		row, err := v.pokemonController.GetPokemonById(id)

		if row == nil {
			message := fmt.Sprintf("Id %v Not Found", id)
			respondWithError(w, http.StatusNotFound, message)
			return
		} 
		
		if err != nil { 
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, &row)
		
	} else {
		v.AddPokemon(w, r)
	}
}

func (v *PokemonView) AddPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		newPoke := json.NewDecoder(r.Body)
		var pokemon pokemon.PokemonFiled
		err := newPoke.Decode(&pokemon)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		row, err := v.pokemonController.AddPokemon(&pokemon)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, fmt.Sprintf("Id Pokemon %v Created", row.Id))
	}
}

