package main

import (
	"database/sql"
)

func getPokemons(db *sql.DB) ([]Pokemon, error) {
	rows, err := db.Query("SELECT * FROM pokemon")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pokemons = []Pokemon{} 

	for rows.Next() {
		var p Pokemon

		err := rows.Scan(&p.Id, &p.Name, &p.Species)

		if err != nil {
			panic(err.Error())
		}

		pokemons = append(pokemons, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (p *Pokemon) getPokemonById(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM pokemon WHERE id = ?", p.Id).Scan(&p.Id, &p.Name, &p.Species)
}

func (p *Pokemon) addNewPokemon(db *sql.DB) error {
	statement, err := db.Prepare("INSERT INTO pokemon VALUES(?, ?, ?)")

	if err != nil {
		
		return err
	}

	result, err := statement.Exec(p.Id, p.Name, p.Species)

	if err != nil {
		return err
	}

	if rowAffected, _ := result.RowsAffected(); rowAffected == 1 {
		return nil
	}

	return nil
} 