package models

import (
	pokemon "Pokemon-API/Pokemon"
	"database/sql"
)

type PokemonMysql struct {
	DB  *sql.DB
}

func Connect(db *sql.DB) pokemon.PokemonInterface {
	return &PokemonMysql {
		DB: db,
	}
}

func (m *PokemonMysql) GetPokemons() ([]pokemon.PokemonFiled, error) {
	rows, err := m.DB.Query("SELECT * FROM pokemon")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pokemons = []pokemon.PokemonFiled{} 

	for rows.Next() {
		var p pokemon.PokemonFiled

		err := rows.Scan(&p.Id, &p.Name, &p.Species)

		if err != nil {
			return nil, err
		}

		pokemons = append(pokemons, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (m *PokemonMysql) GetPokemonById(id int) (*pokemon.PokemonFiled, error) {
	var p pokemon.PokemonFiled
	err := m.DB.QueryRow("SELECT * FROM pokemon WHERE id = ?", id).Scan(&p.Id, &p.Name, &p.Species)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, err
		}
	}
	
	return &p, nil
}

func (m *PokemonMysql) AddPokemon(p *pokemon.PokemonFiled) error {
	statement, err := m.DB.Prepare("INSERT INTO pokemon (id, name, species) VALUES (?, ?, ?)")

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