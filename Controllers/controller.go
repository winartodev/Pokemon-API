package controllers

import (
	models "Pokemon-API/Models"
	pokemon "Pokemon-API/Pokemon"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pokemons, err := models.GetPokemons()

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, pokemons)
	}
}

func GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var pokemonId, _ = strconv.Atoi(r.FormValue("id"))
		var pokemon = pokemon.PokemonFiled{Id: pokemonId}

		err := models.GetPokemonById(&pokemon)

		if err != nil {
			switch {
			case err == sql.ErrNoRows:
				respondWithError(w, http.StatusBadRequest, fmt.Sprintf("ID Pokemon %v Not Found", pokemonId))
			default:
				respondWithError(w, http.StatusInternalServerError, err.Error()) 
			}
			return
		}
	
		respondWithJSON(w, http.StatusOK, pokemon)
	} else if r.Method == "POST" {
		AddNewPokemon(w, r)
	}
}

func AddNewPokemon(w http.ResponseWriter, r *http.Request) {
	newPoke := json.NewDecoder(r.Body)
	var pokemon pokemon.PokemonFiled
	err := newPoke.Decode(&pokemon)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = models.AddNewPokemon(&pokemon)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, fmt.Sprintf("Id Pokemon %v Created", pokemon.Id))
}

func route(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pokemon API")
}

func EndpointsHandler() {
	http.HandleFunc("/", route)
	http.HandleFunc("/pokemons", GetAllPokemons)
	http.HandleFunc("/pokemon", GetPokemonByID)
}

func Run() {
	http.ListenAndServe(":8080", nil)
}
