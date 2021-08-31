package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Controller struct {
	db  *sql.DB
	err error
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

func (c *Controller) connect() *sql.DB {
	c.db, c.err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db_pokemon")

	if c.err != nil {
		panic(c.err.Error())
	}

	return c.db	
}

func (c *Controller) GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		pokemons, err := getPokemons(c.db)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, pokemons)
	}
}

func (c *Controller) GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		var pokemonId, _ = strconv.Atoi(r.FormValue("id"))
		var pokemon = Pokemon{Id: pokemonId}

		err := pokemon.getPokemonById(c.db)

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
		c.AddNewPokemon(w, r)
	}
}

func (c *Controller) AddNewPokemon(w http.ResponseWriter, r *http.Request) {

	newPoke := json.NewDecoder(r.Body)
	
	var pokemon Pokemon
	c.err = newPoke.Decode(&pokemon)

	if c.err != nil {
		respondWithError(w, http.StatusInternalServerError, c.err.Error())
		return
	}

	err := pokemon.addNewPokemon(c.db)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, fmt.Sprintf("Id Pokemon %v Created", pokemon.Id))
}

func route(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pokemon API")
}

func (c *Controller) endpointsHandler() {
	http.HandleFunc("/", route)
	http.HandleFunc("/pokemons", c.GetAllPokemons)
	http.HandleFunc("/pokemon", c.GetPokemonByID)
}

func run() {
	http.ListenAndServe(":8080", nil)
}
