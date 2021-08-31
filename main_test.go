package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var c Controller

func TestMain(m *testing.M) {
	c = Controller{}
	c.db = c.connect()
	defer c.db.Close()

	c.db.Exec("TRUNCATE pokemon")

	m.Run()
}

func TestEmptyTable(t *testing.T) {
	request, err := http.NewRequest("GET", "/pokemons", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetAllPokemons)
	handler.ServeHTTP(response, request)

	if actual := response.Code; actual != http.StatusOK {
		t.Errorf("Expect %v Got %v", actual, http.StatusOK)
	}

	expected := `[]`

	if response.Body.String() != expected {
		t.Errorf("Expect %v Got %v", expected, response.Body.String())
	}
}

func TestAddNewPokemon(t *testing.T) {
	newPokemon := []byte(`{"Id": 1,"Name": "Rattata","Species": "Mouse Pokémon"}`)

	request, err := http.NewRequest("POST", "/pokemon", bytes.NewBuffer(newPokemon))

	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(c.AddNewPokemon)
	handler.ServeHTTP(response, request)

	if actual := response.Code; actual != http.StatusOK {
		t.Errorf("Expect %v Got %v", http.StatusOK, actual)
	}

	expected := `"Id Pokemon 1 Created"`
	
	if response.Body.String() != expected {
		t.Errorf("Expected %v Got %v", expected, response.Body.String())
	}
}

func TestGetPokemonById(t *testing.T) {
	request, err := http.NewRequest("GET", "/pokemon", nil)

	if err != nil {
		t.Fatal(err)
	}

	pokemonId := request.URL.Query()
	pokemonId.Add("id", "1")
	request.URL.RawQuery = pokemonId.Encode()

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetPokemonByID)
	handler.ServeHTTP(response, request)

	if actual := response.Code; actual != http.StatusOK {
		t.Errorf("Expect %v Got %v", http.StatusOK, actual)
	}

	expected :=`{"id":1,"name":"Rattata","species":"Mouse Pokémon"}`
	
	if response.Body.String() != expected {
		t.Errorf("Expect %v Got %v", expected, response.Body.String())
	}
}


