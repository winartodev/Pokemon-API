package models_test

import (
	model "Pokemon-API/Models"
	pokemon "Pokemon-API/Pokemon"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPokemonMysql_GetPokemons(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("database %v connection not connect", err)
	}
	defer db.Close()

	t.Run("Get Pokemons Success Retrieve All Data", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "species"}).
				AddRow(1, "Bulbasaur", "Seed Pokémon").
				AddRow(2, "Charmander", "Lizard Pokémon").
				AddRow(3, "Squirtle", "Tiny Turtle Pokémon").
				AddRow(4, "Raticate", "Mouse Pokémon")

		query := "SELECT (.+) FROM pokemon"

		mock.ExpectQuery(query).WillReturnRows(rows)

		conn := model.Connect(db)
		pokemon, err := conn.GetPokemons()

		assert.NoError(t, err)
		assert.NotNil(t, pokemon)
	})

	t.Run("Get Pokemons Success Retrieve Empty Data", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "species"})

		query := "SELECT (.+) FROM pokemon"

		mock.ExpectQuery(query).WillReturnRows(rows)

		conn := model.Connect(db)
		pokemon, err := conn.GetPokemons()

		assert.NoError(t, err)
		assert.NotNil(t, pokemon)
	})

	t.Run("Get Pokemons Failed Query Error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "species"})

		query := "SELECTS (.+) FROMS pokemons"

		mock.ExpectQuery(query).WillReturnRows(rows)

		conn := model.Connect(db)
		_, err := conn.GetPokemons()

		assert.Error(t, err)
	})
}

func TestPokemonMysql_GetPokemonById(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("database %v connection not connect", err)
	}
	defer db.Close()

	var id = int(1)

	t.Run("Get Pokemon By Success Row Found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "species"}).
				AddRow(1, "Bulbasaur", "Seed Pokémon")

		query := "SELECT (.+) FROM pokemon WHERE id = \\?"

		mock.ExpectQuery(query).WillReturnRows(rows)

		conn := model.Connect(db)
		row, err := conn.GetPokemonByID(id)

		assert.NoError(t, err)
		assert.NotNil(t, row)
	})

	t.Run("Get Pokemon By Id Failed Row Not Found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "species"})

		query := "SELECT (.+) FROM pokemon WHERE id = \\?"

		mock.ExpectQuery(query).WillReturnRows(rows)

		conn := model.Connect(db)
		row, err := conn.GetPokemonByID(id)

		assert.Error(t, err)
		assert.Nil(t, row)
	})
	
	t.Run("Get Pokemon By Id Failed Query Error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "species"}).
						AddRow(1, "Bulbasaur", "Seed Pokémon")

		query := "SELECTS (.+) FROMS pokemons WHERES ids = \\?"

		mock.ExpectQuery(query).WillReturnRows(rows)

		conn := model.Connect(db)
		row, err := conn.GetPokemonByID(id)

		assert.Error(t, err)
		assert.Nil(t, row)
	})
}

func TestPokemonMysql_AddPokemon(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("database %v connection not connect", err)
	}
	defer db.Close()

	t.Run("Add Pokemon Success", func(t *testing.T) {
		newPokemon := &pokemon.Entity{
			ID: 1,
			Name: "Ivysaur",
			Species: "Seed Pokémon",
		}

		query := "INSERT INTO pokemon \\(id, name, species\\) VALUES \\(\\?, \\?, \\?\\)"
		
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(newPokemon.ID, newPokemon.Name, newPokemon.Species).WillReturnResult(sqlmock.NewResult(0, 0))

		conn := model.Connect(db)
		_, err := conn.AddPokemon(newPokemon)	

		assert.NoError(t, err)
	})

	t.Run("Add Pokemon Failed Query Error", func(t *testing.T) {
		newPokemon := &pokemon.Entity{
			ID: 1,
			Name: "Ivysaur",
			Species: "Seed Pokémon",
		}

		query := "INSERT INTO pokemons \\(id, name, species\\) VALUES \\(\\?, \\?, \\?\\)"
		
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(newPokemon.ID, newPokemon.Name, newPokemon.Species).WillReturnResult(sqlmock.NewResult(0, 0))

		conn := model.Connect(db)
		_, err = conn.AddPokemon(newPokemon)	

		assert.Error(t, err)
	})

	t.Run("Add Pokemon Failed Filed Empty", func(t *testing.T) {
		newPokemon := &pokemon.Entity{}

		query := "INSERT INTO pokemons \\(id, name, species\\) VALUES \\(\\?, \\?, \\?\\)"
		
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(newPokemon.ID, newPokemon.Name).WillReturnResult(sqlmock.NewResult(0, 0))

		conn := model.Connect(db)
		_, err = conn.AddPokemon(newPokemon)	
		
		assert.Error(t, err)
	})
}
