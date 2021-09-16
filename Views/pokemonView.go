package view

import (
	pokemon "Pokemon-API/Pokemon"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type PokemonView struct {
	pokemonController pokemon.ControllerInterface
}

// respondWithError to show error code and error message
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON to show response code and body into client
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// w.Write(response)
	fmt.Fprint(w, string(response))
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
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		respondWithJSON(w, http.StatusOK, rows)
	}
}

// GetPokemonByID is used to get pokemon by id
func (v *PokemonView) GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, _ := strconv.Atoi(r.FormValue("id"))

		row, err := v.pokemonController.GetPokemonByID(id)

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

// AddPokemon is used to create new pokemon
func (v *PokemonView) AddPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		newPoke := json.NewDecoder(r.Body)
		var pokemon pokemon.Entity
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

		respondWithJSON(w, http.StatusOK, fmt.Sprintf("Id Pokemon %v Created", row.ID))
	}
}
