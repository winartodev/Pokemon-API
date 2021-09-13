package models

import (
	pokemon "Pokemon-API/Pokemon"
	"database/sql"
)

type PokemonMysql struct {
	DB  *sql.DB
}

func Connect(db *sql.DB) pokemon.ModelInterface {
	return &PokemonMysql {
		DB: db,
	}
}

func (m *PokemonMysql) GetPokemons() ([]pokemon.Entity, error) {
	rows, err := m.DB.Query("SELECT * FROM pokemon")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pokemons = []pokemon.Entity{} 

	for rows.Next() {
		var p pokemon.Entity

		err := rows.Scan(&p.ID, &p.Name, &p.Species)

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

func (m *PokemonMysql) GetPokemonByID(id int) (*pokemon.Entity, error) {
	var p pokemon.Entity
	err := m.DB.QueryRow("SELECT * FROM pokemon WHERE id = ?", id).Scan(&p.ID, &p.Name, &p.Species)
	if err != nil {
		if err == sql.ErrNoRows{ 
			return nil, err
		}
		return nil, err
	}
	
	return &p, nil
}

func (m *PokemonMysql) AddPokemon(p *pokemon.Entity) (*pokemon.Entity, error) {
	statement, err := m.DB.Prepare("INSERT INTO pokemon (id, name, species) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	result, err := statement.Exec(&p.ID, &p.Name, &p.Species)

	if err != nil {
		return nil, err
	}

	if rowAffected, _ := result.RowsAffected(); rowAffected == 1 {
		return p, nil
	}

	return p, nil
} 