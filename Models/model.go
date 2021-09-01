package models

import (
	pokemon "Pokemon-API/Pokemon"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db  *sql.DB
var err error

func Connect() *sql.DB {
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db_pokemon")

	if err != nil {
		panic(err.Error())
	}

	return db	
}

func GetPokemons() ([]pokemon.PokemonFiled, error) {
	rows, err := db.Query("SELECT * FROM pokemon")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pokemons = []pokemon.PokemonFiled{} 

	for rows.Next() {
		var p pokemon.PokemonFiled

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

func GetPokemonById(p *pokemon.PokemonFiled) error {
	return db.QueryRow("SELECT * FROM pokemon WHERE id = ?", p.Id).Scan(&p.Id, &p.Name, &p.Species)
}

func AddNewPokemon(p  *pokemon.PokemonFiled) error {
	statement, err := db.Prepare("INSERT INTO pokemon(id, name, species) VALUES(?, ?, ?)")

	if err != nil {
		return err
	}

	result, err := statement.Exec(&p.Id, &p.Name, &p.Species)

	if err != nil {
		return err
	}

	if rowAffected, _ := result.RowsAffected(); rowAffected == 1 {
		return nil
	}

	return nil
} 